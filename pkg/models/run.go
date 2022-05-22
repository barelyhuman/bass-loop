package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/vito/bass/pkg/bass"
)

func CreateThunkRun(ctx context.Context, db *sql.DB, thunk bass.Thunk) (*Run, error) {
	sha2, err := thunk.SHA256()
	if err != nil {
		return nil, err
	}

	payload, err := bass.MarshalJSON(thunk)
	if err != nil {
		return nil, err
	}

	dbThunk, err := ThunkByDigest(ctx, db, sha2)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			dbThunk = &Thunk{
				Digest:    sha2,
				JSON:      payload,
				Sensitive: 0,
			}

			err = dbThunk.Save(ctx, db)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	thunkRun := Run{
		ID:          id.String(),
		ThunkDigest: sha2,
		StartTime:   int(time.Now().UnixNano()),
	}

	err = thunkRun.Save(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("save thunk run: %w", err)
	}

	return &thunkRun, nil
}

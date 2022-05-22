package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// User represents a row from 'users'.
type User struct {
	NodeID string `json:"node_id"` // node_id
	Login  string `json:"login"`   // login
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted returns true when the User has been marked for deletion from
// the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) error {
	switch {
	case u._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case u._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO users (` +
		`node_id, login` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	logf(sqlstr, u.NodeID, u.Login)
	if _, err := db.ExecContext(ctx, sqlstr, u.NodeID, u.Login); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case u._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE users SET ` +
		`login = $1 ` +
		`WHERE node_id = $2`
	// run
	logf(sqlstr, u.Login, u.NodeID)
	if _, err := db.ExecContext(ctx, sqlstr, u.Login, u.NodeID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the User to the database.
func (u *User) Save(ctx context.Context, db DB) error {
	if u.Exists() {
		return u.Update(ctx, db)
	}
	return u.Insert(ctx, db)
}

// Upsert performs an upsert for User.
func (u *User) Upsert(ctx context.Context, db DB) error {
	switch {
	case u._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO users (` +
		`node_id, login` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` ON CONFLICT (node_id) DO ` +
		`UPDATE SET ` +
		`login = EXCLUDED.login `
	// run
	logf(sqlstr, u.NodeID, u.Login)
	if _, err := db.ExecContext(ctx, sqlstr, u.NodeID, u.Login); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return nil
	case u._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM users ` +
		`WHERE node_id = $1`
	// run
	logf(sqlstr, u.NodeID)
	if _, err := db.ExecContext(ctx, sqlstr, u.NodeID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	return nil
}

// UserByNodeID retrieves a row from 'users' as a User.
//
// Generated from index 'sqlite_autoindex_users_1'.
func UserByNodeID(ctx context.Context, db DB, nodeID string) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`node_id, login ` +
		`FROM users ` +
		`WHERE node_id = $1`
	// run
	logf(sqlstr, nodeID)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, nodeID).Scan(&u.NodeID, &u.Login); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}
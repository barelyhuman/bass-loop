package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Service represents a row from 'services'.
type Service struct {
	UserID      string `json:"user_id"`      // user_id
	RuntimeName string `json:"runtime_name"` // runtime_name
	Service     string `json:"service"`      // service
	Addr        string `json:"addr"`         // addr
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Service exists in the database.
func (s *Service) Exists() bool {
	return s._exists
}

// Deleted returns true when the Service has been marked for deletion from
// the database.
func (s *Service) Deleted() bool {
	return s._deleted
}

// Insert inserts the Service to the database.
func (s *Service) Insert(ctx context.Context, db DB) error {
	switch {
	case s._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case s._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO services (` +
		`user_id, runtime_name, service, addr` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`
	// run
	logf(sqlstr, s.UserID, s.RuntimeName, s.Service, s.Addr)
	if _, err := db.ExecContext(ctx, sqlstr, s.UserID, s.RuntimeName, s.Service, s.Addr); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Update updates a Service in the database.
func (s *Service) Update(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case s._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE services SET ` +
		`user_id = $1, addr = $2 ` +
		`WHERE runtime_name = $3 AND service = $4`
	// run
	logf(sqlstr, s.UserID, s.Addr, s.RuntimeName, s.Service)
	if _, err := db.ExecContext(ctx, sqlstr, s.UserID, s.Addr, s.RuntimeName, s.Service); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Service to the database.
func (s *Service) Save(ctx context.Context, db DB) error {
	if s.Exists() {
		return s.Update(ctx, db)
	}
	return s.Insert(ctx, db)
}

// Upsert performs an upsert for Service.
func (s *Service) Upsert(ctx context.Context, db DB) error {
	switch {
	case s._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO services (` +
		`user_id, runtime_name, service, addr` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (runtime_name, service) DO ` +
		`UPDATE SET ` +
		`user_id = EXCLUDED.user_id, addr = EXCLUDED.addr `
	// run
	logf(sqlstr, s.UserID, s.RuntimeName, s.Service, s.Addr)
	if _, err := db.ExecContext(ctx, sqlstr, s.UserID, s.RuntimeName, s.Service, s.Addr); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Delete deletes the Service from the database.
func (s *Service) Delete(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return nil
	case s._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM services ` +
		`WHERE runtime_name = $1 AND service = $2`
	// run
	logf(sqlstr, s.RuntimeName, s.Service)
	if _, err := db.ExecContext(ctx, sqlstr, s.RuntimeName, s.Service); err != nil {
		return logerror(err)
	}
	// set deleted
	s._deleted = true
	return nil
}

// ServicesByUserIDRuntimeName retrieves a row from 'services' as a Service.
//
// Generated from index 'idx_services_runtime_name'.
func ServicesByUserIDRuntimeName(ctx context.Context, db DB, userID, runtimeName string) ([]*Service, error) {
	// query
	const sqlstr = `SELECT ` +
		`user_id, runtime_name, service, addr ` +
		`FROM services ` +
		`WHERE user_id = $1 AND runtime_name = $2`
	// run
	logf(sqlstr, userID, runtimeName)
	rows, err := db.QueryContext(ctx, sqlstr, userID, runtimeName)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Service
	for rows.Next() {
		s := Service{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&s.UserID, &s.RuntimeName, &s.Service, &s.Addr); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ServiceByRuntimeNameService retrieves a row from 'services' as a Service.
//
// Generated from index 'sqlite_autoindex_services_1'.
func ServiceByRuntimeNameService(ctx context.Context, db DB, runtimeName, service string) (*Service, error) {
	// query
	const sqlstr = `SELECT ` +
		`user_id, runtime_name, service, addr ` +
		`FROM services ` +
		`WHERE runtime_name = $1 AND service = $2`
	// run
	logf(sqlstr, runtimeName, service)
	s := Service{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, runtimeName, service).Scan(&s.UserID, &s.RuntimeName, &s.Service, &s.Addr); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}

// User returns the User associated with the Service's (UserID).
//
// Generated from foreign key 'services_user_id_fkey'.
func (s *Service) User(ctx context.Context, db DB) (*User, error) {
	return UserByID(ctx, db, s.UserID)
}

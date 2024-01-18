// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: bucket.sql

package client

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addAllowedMimeTypesToBucket = `-- name: AddAllowedMimeTypesToBucket :exec
update storage.buckets
set allowed_mime_types = array_append(allowed_mime_types, $1::text[])
where id = $2
`

type AddAllowedMimeTypesToBucketParams struct {
	MimeType []string
	ID       string
}

func (q *Queries) AddAllowedMimeTypesToBucket(ctx context.Context, arg AddAllowedMimeTypesToBucketParams) error {
	_, err := q.db.Exec(ctx, addAllowedMimeTypesToBucket, arg.MimeType, arg.ID)
	return err
}

const countBuckets = `-- name: CountBuckets :one
select count(1) as count
from storage.buckets
`

func (q *Queries) CountBuckets(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countBuckets)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createBucket = `-- name: CreateBucket :exec
insert into storage.buckets
(id, name, allowed_mime_types, max_allowed_object_size, public, disabled)
values ($1,
        $2,
        $3,
        $4,
        $5,
        $6)
returning id, name, allowed_mime_types, max_allowed_object_size, public, disabled, locked, lock_reason, locked_at, created_at, updated_at
`

type CreateBucketParams struct {
	ID                   string
	Name                 string
	AllowedMimeTypes     []string
	MaxAllowedObjectSize pgtype.Int8
	Public               bool
	Disabled             bool
}

func (q *Queries) CreateBucket(ctx context.Context, arg CreateBucketParams) error {
	_, err := q.db.Exec(ctx, createBucket,
		arg.ID,
		arg.Name,
		arg.AllowedMimeTypes,
		arg.MaxAllowedObjectSize,
		arg.Public,
		arg.Disabled,
	)
	return err
}

const deleteBucket = `-- name: DeleteBucket :exec
delete
from storage.buckets
where id = $1
`

func (q *Queries) DeleteBucket(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteBucket, id)
	return err
}

const disableBucket = `-- name: DisableBucket :exec
update storage.buckets
set disabled = true
where id = $1
`

func (q *Queries) DisableBucket(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, disableBucket, id)
	return err
}

const enableBucket = `-- name: EnableBucket :exec
update storage.buckets
set disabled = false
where id = $1
`

func (q *Queries) EnableBucket(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, enableBucket, id)
	return err
}

const getBucketById = `-- name: GetBucketById :one
select id,
       name,
       allowed_mime_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where id = $1
limit 1
`

func (q *Queries) GetBucketById(ctx context.Context, id string) (StorageBucket, error) {
	row := q.db.QueryRow(ctx, getBucketById, id)
	var i StorageBucket
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AllowedMimeTypes,
		&i.MaxAllowedObjectSize,
		&i.Public,
		&i.Disabled,
		&i.Locked,
		&i.LockReason,
		&i.LockedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBucketByName = `-- name: GetBucketByName :one
select id,
       name,
       allowed_mime_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where name = $1
limit 1
`

func (q *Queries) GetBucketByName(ctx context.Context, name string) (StorageBucket, error) {
	row := q.db.QueryRow(ctx, getBucketByName, name)
	var i StorageBucket
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AllowedMimeTypes,
		&i.MaxAllowedObjectSize,
		&i.Public,
		&i.Disabled,
		&i.Locked,
		&i.LockReason,
		&i.LockedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBucketObjectCountById = `-- name: GetBucketObjectCountById :one
select count(1) as count
from storage.objects
where bucket_id = $1
`

func (q *Queries) GetBucketObjectCountById(ctx context.Context, id string) (int64, error) {
	row := q.db.QueryRow(ctx, getBucketObjectCountById, id)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getBucketObjectCountByName = `-- name: GetBucketObjectCountByName :one
select count(1) as count
from storage.objects
where bucket_id = (select id from storage.buckets where storage.buckets.name = $1)
`

func (q *Queries) GetBucketObjectCountByName(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRow(ctx, getBucketObjectCountByName, name)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getBucketSizeById = `-- name: GetBucketSizeById :one
select sum(size) as size
from storage.objects
where bucket_id = $1
`

func (q *Queries) GetBucketSizeById(ctx context.Context, id string) (int64, error) {
	row := q.db.QueryRow(ctx, getBucketSizeById, id)
	var size int64
	err := row.Scan(&size)
	return size, err
}

const getBucketSizeByName = `-- name: GetBucketSizeByName :one
select sum(size) as size
from storage.objects
where bucket_id = (select id from storage.buckets where storage.buckets.name = $1)
`

func (q *Queries) GetBucketSizeByName(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRow(ctx, getBucketSizeByName, name)
	var size int64
	err := row.Scan(&size)
	return size, err
}

const listAllBuckets = `-- name: ListAllBuckets :many
select id,
       name,
       allowed_mime_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
`

func (q *Queries) ListAllBuckets(ctx context.Context) ([]StorageBucket, error) {
	rows, err := q.db.Query(ctx, listAllBuckets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StorageBucket
	for rows.Next() {
		var i StorageBucket
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AllowedMimeTypes,
			&i.MaxAllowedObjectSize,
			&i.Public,
			&i.Disabled,
			&i.Locked,
			&i.LockReason,
			&i.LockedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBucketsPaged = `-- name: ListBucketsPaged :many
select id,
       name,
       allowed_mime_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
limit $2 offset $1
`

type ListBucketsPagedParams struct {
	Offset pgtype.Int4
	Limit  pgtype.Int4
}

func (q *Queries) ListBucketsPaged(ctx context.Context, arg ListBucketsPagedParams) ([]StorageBucket, error) {
	rows, err := q.db.Query(ctx, listBucketsPaged, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StorageBucket
	for rows.Next() {
		var i StorageBucket
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AllowedMimeTypes,
			&i.MaxAllowedObjectSize,
			&i.Public,
			&i.Disabled,
			&i.Locked,
			&i.LockReason,
			&i.LockedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const lockBucket = `-- name: LockBucket :exec
update storage.buckets
set locked      = true,
    lock_reason = $1,
    locked_at   = now()
where id = $2
`

type LockBucketParams struct {
	LockReason pgtype.Text
	ID         string
}

func (q *Queries) LockBucket(ctx context.Context, arg LockBucketParams) error {
	_, err := q.db.Exec(ctx, lockBucket, arg.LockReason, arg.ID)
	return err
}

const makeBucketPrivate = `-- name: MakeBucketPrivate :exec
update storage.buckets
set public = false
where id = $1
`

func (q *Queries) MakeBucketPrivate(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, makeBucketPrivate, id)
	return err
}

const makeBucketPublic = `-- name: MakeBucketPublic :exec
update storage.buckets
set public = true
where id = $1
`

func (q *Queries) MakeBucketPublic(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, makeBucketPublic, id)
	return err
}

const removeAllowedMimeTypesFromBucket = `-- name: RemoveAllowedMimeTypesFromBucket :exec
update storage.buckets
set allowed_mime_types = array_remove(allowed_mime_types, $1::text[])
where id = $2
`

type RemoveAllowedMimeTypesFromBucketParams struct {
	MimeType []string
	ID       string
}

func (q *Queries) RemoveAllowedMimeTypesFromBucket(ctx context.Context, arg RemoveAllowedMimeTypesFromBucketParams) error {
	_, err := q.db.Exec(ctx, removeAllowedMimeTypesFromBucket, arg.MimeType, arg.ID)
	return err
}

const searchBucketsPaged = `-- name: SearchBucketsPaged :many
select id,
       name,
       allowed_mime_types,
       max_allowed_object_size,
       public,
       disabled,
       locked,
       lock_reason,
       locked_at,
       created_at,
       updated_at
from storage.buckets
where name ilike $1
limit $3 offset $2
`

type SearchBucketsPagedParams struct {
	Name   pgtype.Text
	Offset pgtype.Int4
	Limit  pgtype.Int4
}

func (q *Queries) SearchBucketsPaged(ctx context.Context, arg SearchBucketsPagedParams) ([]StorageBucket, error) {
	rows, err := q.db.Query(ctx, searchBucketsPaged, arg.Name, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StorageBucket
	for rows.Next() {
		var i StorageBucket
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AllowedMimeTypes,
			&i.MaxAllowedObjectSize,
			&i.Public,
			&i.Disabled,
			&i.Locked,
			&i.LockReason,
			&i.LockedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unlockBucket = `-- name: UnlockBucket :exec
update storage.buckets
set locked      = false,
    lock_reason = null,
    locked_at   = null
where id = $1
`

func (q *Queries) UnlockBucket(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, unlockBucket, id)
	return err
}

const updateBucketMaxAllowedObjectSize = `-- name: UpdateBucketMaxAllowedObjectSize :exec
update storage.buckets
set max_allowed_object_size = $1
where id = $2
`

type UpdateBucketMaxAllowedObjectSizeParams struct {
	MaxAllowedObjectSize pgtype.Int8
	ID                   string
}

func (q *Queries) UpdateBucketMaxAllowedObjectSize(ctx context.Context, arg UpdateBucketMaxAllowedObjectSizeParams) error {
	_, err := q.db.Exec(ctx, updateBucketMaxAllowedObjectSize, arg.MaxAllowedObjectSize, arg.ID)
	return err
}

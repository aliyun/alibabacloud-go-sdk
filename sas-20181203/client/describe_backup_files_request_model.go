// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeBackupFilesRequest
	GetCurrentPage() *string
	SetPageSize(v string) *DescribeBackupFilesRequest
	GetPageSize() *string
	SetPath(v string) *DescribeBackupFilesRequest
	GetPath() *string
	SetSnapshotHash(v string) *DescribeBackupFilesRequest
	GetSnapshotHash() *string
	SetUuid(v string) *DescribeBackupFilesRequest
	GetUuid() *string
}

type DescribeBackupFilesRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The path to the backup file.
	//
	// example:
	//
	// “”
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The hash value of the backup file.
	//
	// This parameter is required.
	//
	// example:
	//
	// a7f26223ef3974c6fac324cd37713ab65ab618859d20b4039192a5da44d77b63
	SnapshotHash *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	// The UUID of the server to which an anti-ransomware policy is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6d5b361f-958d-48a8-a9d2-d6e82c1a****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeBackupFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeBackupFilesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeBackupFilesRequest) GetPath() *string {
	return s.Path
}

func (s *DescribeBackupFilesRequest) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *DescribeBackupFilesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeBackupFilesRequest) SetCurrentPage(v string) *DescribeBackupFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetPageSize(v string) *DescribeBackupFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetPath(v string) *DescribeBackupFilesRequest {
	s.Path = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetSnapshotHash(v string) *DescribeBackupFilesRequest {
	s.SnapshotHash = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetUuid(v string) *DescribeBackupFilesRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeBackupFilesRequest) Validate() error {
	return dara.Validate(s)
}

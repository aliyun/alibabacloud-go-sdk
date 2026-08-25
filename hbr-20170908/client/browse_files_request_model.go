// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowseFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbsolutePath(v string) *BrowseFilesRequest
	GetAbsolutePath() *string
	SetClientId(v string) *BrowseFilesRequest
	GetClientId() *string
	SetEdition(v string) *BrowseFilesRequest
	GetEdition() *string
	SetMaxResults(v int32) *BrowseFilesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *BrowseFilesRequest
	GetNextToken() *string
	SetPageNumber(v int32) *BrowseFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *BrowseFilesRequest
	GetPageSize() *int32
	SetPath(v string) *BrowseFilesRequest
	GetPath() *string
	SetRestoreId(v string) *BrowseFilesRequest
	GetRestoreId() *string
	SetSecurityToken(v string) *BrowseFilesRequest
	GetSecurityToken() *string
	SetSnapshotHash(v string) *BrowseFilesRequest
	GetSnapshotHash() *string
	SetStorageClass(v string) *BrowseFilesRequest
	GetStorageClass() *string
	SetToken(v string) *BrowseFilesRequest
	GetToken() *string
	SetVaultId(v string) *BrowseFilesRequest
	GetVaultId() *string
}

type BrowseFilesRequest struct {
	// The absolute path of the directory. Specify `/` to browse the root directory of the backup.
	//
	// example:
	//
	// /data/
	AbsolutePath *string `json:"AbsolutePath,omitempty" xml:"AbsolutePath,omitempty"`
	// The backup client ID.
	//
	// example:
	//
	// c-000***o48
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The Cloud Backup edition. Valid values:
	//
	// - **STANDARD**: Standard Edition. This is the default value.
	//
	// - **BASIC**: Basic Edition. Only ECS file backup is supported in Basic Edition.
	//
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The maximum number of results to return per request.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is required to retrieve the next page of results. If this parameter is not specified, the first page of results is returned.
	//
	// example:
	//
	// eyJ***Q==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Deprecated. Use MaxResults and NextToken for pagination instead.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Deprecated. Use MaxResults and NextToken for pagination instead.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The hash value of the directory. If this parameter is not specified, the root directory of the backup is browsed.
	//
	// example:
	//
	// ef6***46a
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Deprecated.
	//
	// example:
	//
	// r-000***oy9
	RestoreId     *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The hash value of the backup snapshot.
	//
	// example:
	//
	// 971***e9d
	SnapshotHash *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	// The storage class of the backup data. Valid values:
	//
	// - **STANDARD**: Standard.
	//
	// - **ARCHIVE**: Archive.
	//
	// example:
	//
	// STANDARD
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// Deprecated. Do not use.
	//
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The backup vault ID.
	//
	// example:
	//
	// v-000***jtz
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s BrowseFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s BrowseFilesRequest) GoString() string {
	return s.String()
}

func (s *BrowseFilesRequest) GetAbsolutePath() *string {
	return s.AbsolutePath
}

func (s *BrowseFilesRequest) GetClientId() *string {
	return s.ClientId
}

func (s *BrowseFilesRequest) GetEdition() *string {
	return s.Edition
}

func (s *BrowseFilesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *BrowseFilesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *BrowseFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *BrowseFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *BrowseFilesRequest) GetPath() *string {
	return s.Path
}

func (s *BrowseFilesRequest) GetRestoreId() *string {
	return s.RestoreId
}

func (s *BrowseFilesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BrowseFilesRequest) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *BrowseFilesRequest) GetStorageClass() *string {
	return s.StorageClass
}

func (s *BrowseFilesRequest) GetToken() *string {
	return s.Token
}

func (s *BrowseFilesRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *BrowseFilesRequest) SetAbsolutePath(v string) *BrowseFilesRequest {
	s.AbsolutePath = &v
	return s
}

func (s *BrowseFilesRequest) SetClientId(v string) *BrowseFilesRequest {
	s.ClientId = &v
	return s
}

func (s *BrowseFilesRequest) SetEdition(v string) *BrowseFilesRequest {
	s.Edition = &v
	return s
}

func (s *BrowseFilesRequest) SetMaxResults(v int32) *BrowseFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *BrowseFilesRequest) SetNextToken(v string) *BrowseFilesRequest {
	s.NextToken = &v
	return s
}

func (s *BrowseFilesRequest) SetPageNumber(v int32) *BrowseFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *BrowseFilesRequest) SetPageSize(v int32) *BrowseFilesRequest {
	s.PageSize = &v
	return s
}

func (s *BrowseFilesRequest) SetPath(v string) *BrowseFilesRequest {
	s.Path = &v
	return s
}

func (s *BrowseFilesRequest) SetRestoreId(v string) *BrowseFilesRequest {
	s.RestoreId = &v
	return s
}

func (s *BrowseFilesRequest) SetSecurityToken(v string) *BrowseFilesRequest {
	s.SecurityToken = &v
	return s
}

func (s *BrowseFilesRequest) SetSnapshotHash(v string) *BrowseFilesRequest {
	s.SnapshotHash = &v
	return s
}

func (s *BrowseFilesRequest) SetStorageClass(v string) *BrowseFilesRequest {
	s.StorageClass = &v
	return s
}

func (s *BrowseFilesRequest) SetToken(v string) *BrowseFilesRequest {
	s.Token = &v
	return s
}

func (s *BrowseFilesRequest) SetVaultId(v string) *BrowseFilesRequest {
	s.VaultId = &v
	return s
}

func (s *BrowseFilesRequest) Validate() error {
	return dara.Validate(s)
}

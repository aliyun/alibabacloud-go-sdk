// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDNAFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBId(v string) *ListDNAFilesRequest
	GetDBId() *string
	SetNextPageToken(v string) *ListDNAFilesRequest
	GetNextPageToken() *string
	SetOwnerAccount(v string) *ListDNAFilesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListDNAFilesRequest
	GetOwnerId() *int64
	SetPageSize(v int32) *ListDNAFilesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListDNAFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListDNAFilesRequest
	GetResourceOwnerId() *int64
}

type ListDNAFilesRequest struct {
	// The ID of the media fingerprint library.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****
	DBId *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// ae0fd49c0840e14daf0d66a75b83****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListDNAFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDNAFilesRequest) GoString() string {
	return s.String()
}

func (s *ListDNAFilesRequest) GetDBId() *string {
	return s.DBId
}

func (s *ListDNAFilesRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListDNAFilesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListDNAFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDNAFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDNAFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListDNAFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListDNAFilesRequest) SetDBId(v string) *ListDNAFilesRequest {
	s.DBId = &v
	return s
}

func (s *ListDNAFilesRequest) SetNextPageToken(v string) *ListDNAFilesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListDNAFilesRequest) SetOwnerAccount(v string) *ListDNAFilesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListDNAFilesRequest) SetOwnerId(v int64) *ListDNAFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDNAFilesRequest) SetPageSize(v int32) *ListDNAFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDNAFilesRequest) SetResourceOwnerAccount(v string) *ListDNAFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDNAFilesRequest) SetResourceOwnerId(v int64) *ListDNAFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDNAFilesRequest) Validate() error {
	return dara.Validate(s)
}

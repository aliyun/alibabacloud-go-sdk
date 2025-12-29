// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagListPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryTagListPageRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *QueryTagListPageRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryTagListPageRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *QueryTagListPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTagListPageRequest
	GetResourceOwnerId() *int64
}

type QueryTagListPageRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 18
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 66
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryTagListPageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTagListPageRequest) GoString() string {
	return s.String()
}

func (s *QueryTagListPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTagListPageRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryTagListPageRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryTagListPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTagListPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTagListPageRequest) SetOwnerId(v int64) *QueryTagListPageRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTagListPageRequest) SetPageNo(v int64) *QueryTagListPageRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTagListPageRequest) SetPageSize(v int64) *QueryTagListPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTagListPageRequest) SetResourceOwnerAccount(v string) *QueryTagListPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTagListPageRequest) SetResourceOwnerId(v int64) *QueryTagListPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTagListPageRequest) Validate() error {
	return dara.Validate(s)
}

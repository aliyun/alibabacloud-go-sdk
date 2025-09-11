// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsSignListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QuerySmsSignListRequest
	GetOwnerId() *int64
	SetPageIndex(v int32) *QuerySmsSignListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *QuerySmsSignListRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QuerySmsSignListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsSignListRequest
	GetResourceOwnerId() *int64
}

type QuerySmsSignListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of signatures per page. Valid values: **1 to 50**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySmsSignListRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsSignListRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsSignListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsSignListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *QuerySmsSignListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySmsSignListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsSignListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsSignListRequest) SetOwnerId(v int64) *QuerySmsSignListRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsSignListRequest) SetPageIndex(v int32) *QuerySmsSignListRequest {
	s.PageIndex = &v
	return s
}

func (s *QuerySmsSignListRequest) SetPageSize(v int32) *QuerySmsSignListRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySmsSignListRequest) SetResourceOwnerAccount(v string) *QuerySmsSignListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsSignListRequest) SetResourceOwnerId(v int64) *QuerySmsSignListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsSignListRequest) Validate() error {
	return dara.Validate(s)
}

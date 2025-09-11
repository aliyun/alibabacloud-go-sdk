// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QuerySmsTemplateListRequest
	GetOwnerId() *int64
	SetPageIndex(v int32) *QuerySmsTemplateListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *QuerySmsTemplateListRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QuerySmsTemplateListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsTemplateListRequest
	GetResourceOwnerId() *int64
}

type QuerySmsTemplateListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of templates per page. Valid values: **1 to 50**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySmsTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTemplateListRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsTemplateListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *QuerySmsTemplateListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySmsTemplateListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsTemplateListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsTemplateListRequest) SetOwnerId(v int64) *QuerySmsTemplateListRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsTemplateListRequest) SetPageIndex(v int32) *QuerySmsTemplateListRequest {
	s.PageIndex = &v
	return s
}

func (s *QuerySmsTemplateListRequest) SetPageSize(v int32) *QuerySmsTemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySmsTemplateListRequest) SetResourceOwnerAccount(v string) *QuerySmsTemplateListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsTemplateListRequest) SetResourceOwnerId(v int64) *QuerySmsTemplateListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsTemplateListRequest) Validate() error {
	return dara.Validate(s)
}

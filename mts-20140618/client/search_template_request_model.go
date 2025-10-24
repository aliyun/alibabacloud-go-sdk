// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamePrefix(v string) *SearchTemplateRequest
	GetNamePrefix() *string
	SetOwnerAccount(v string) *SearchTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SearchTemplateRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *SearchTemplateRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *SearchTemplateRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *SearchTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SearchTemplateRequest
	GetResourceOwnerId() *int64
	SetState(v string) *SearchTemplateRequest
	GetState() *string
}

type SearchTemplateRequest struct {
	// The name prefix based on which you want to search for templates.
	//
	// example:
	//
	// S00000001
	NamePrefix   *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The size of each page set during the result paging query.
	//
	// - Upper limit: 100.
	//
	// - Default value: 10.
	//
	// example:
	//
	// 10
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the custom transcoding templates that you want to query.
	//
	// 	- **All**: All custom transcoding templates are queried.
	//
	// 	- **Normal**: Normal custom transcoding templates are queried.
	//
	// 	- **Deleted**: Deleted custom transcoding templates are queried.
	//
	// 	- Default value: **All**.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SearchTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateRequest) GoString() string {
	return s.String()
}

func (s *SearchTemplateRequest) GetNamePrefix() *string {
	return s.NamePrefix
}

func (s *SearchTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SearchTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SearchTemplateRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchTemplateRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SearchTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SearchTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SearchTemplateRequest) GetState() *string {
	return s.State
}

func (s *SearchTemplateRequest) SetNamePrefix(v string) *SearchTemplateRequest {
	s.NamePrefix = &v
	return s
}

func (s *SearchTemplateRequest) SetOwnerAccount(v string) *SearchTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SearchTemplateRequest) SetOwnerId(v int64) *SearchTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchTemplateRequest) SetPageNumber(v int64) *SearchTemplateRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTemplateRequest) SetPageSize(v int64) *SearchTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTemplateRequest) SetResourceOwnerAccount(v string) *SearchTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SearchTemplateRequest) SetResourceOwnerId(v int64) *SearchTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SearchTemplateRequest) SetState(v string) *SearchTemplateRequest {
	s.State = &v
	return s
}

func (s *SearchTemplateRequest) Validate() error {
	return dara.Validate(s)
}

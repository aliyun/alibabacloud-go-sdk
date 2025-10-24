// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *SearchMediaWorkflowRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SearchMediaWorkflowRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *SearchMediaWorkflowRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *SearchMediaWorkflowRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *SearchMediaWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SearchMediaWorkflowRequest
	GetResourceOwnerId() *int64
	SetStateList(v string) *SearchMediaWorkflowRequest
	GetStateList() *string
}

type SearchMediaWorkflowRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// 	- A maximum of **100*	- entries can be returned on each page.
	//
	// 	- Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the media workflows that you want to query. You can specify multiple states. Separate multiple states with commas (,). Default value: **Inactive,Active,Deleted**. Valid values:
	//
	// 	- **Inactive**: Deactivated media workflows are queried.
	//
	// 	- **Active**: Activated media workflows are queried.
	//
	// 	- **Deleted**: Deleted media workflows are queried.
	//
	// example:
	//
	// Inactive,Active,Deleted
	StateList *string `json:"StateList,omitempty" xml:"StateList,omitempty"`
}

func (s SearchMediaWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaWorkflowRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaWorkflowRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SearchMediaWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SearchMediaWorkflowRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchMediaWorkflowRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SearchMediaWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SearchMediaWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SearchMediaWorkflowRequest) GetStateList() *string {
	return s.StateList
}

func (s *SearchMediaWorkflowRequest) SetOwnerAccount(v string) *SearchMediaWorkflowRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SearchMediaWorkflowRequest) SetOwnerId(v int64) *SearchMediaWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchMediaWorkflowRequest) SetPageNumber(v int64) *SearchMediaWorkflowRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchMediaWorkflowRequest) SetPageSize(v int64) *SearchMediaWorkflowRequest {
	s.PageSize = &v
	return s
}

func (s *SearchMediaWorkflowRequest) SetResourceOwnerAccount(v string) *SearchMediaWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SearchMediaWorkflowRequest) SetResourceOwnerId(v int64) *SearchMediaWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SearchMediaWorkflowRequest) SetStateList(v string) *SearchMediaWorkflowRequest {
	s.StateList = &v
	return s
}

func (s *SearchMediaWorkflowRequest) Validate() error {
	return dara.Validate(s)
}

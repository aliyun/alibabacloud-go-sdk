// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *SearchPipelineRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SearchPipelineRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *SearchPipelineRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *SearchPipelineRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *SearchPipelineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SearchPipelineRequest
	GetResourceOwnerId() *int64
	SetState(v string) *SearchPipelineRequest
	GetState() *string
}

type SearchPipelineRequest struct {
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
	// The status of the MPS queues that you want to query. If you leave this parameter empty, all MPS queues are queried.
	//
	// 	- **All**: queries all MPS queues.
	//
	// 	- **Active**: queries the MPS queues that are active.
	//
	// 	- **Paused**: queues the MPS queues that are paused.
	//
	// 	- Default value: **All**.
	//
	// example:
	//
	// Paused
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SearchPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchPipelineRequest) GoString() string {
	return s.String()
}

func (s *SearchPipelineRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SearchPipelineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SearchPipelineRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchPipelineRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SearchPipelineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SearchPipelineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SearchPipelineRequest) GetState() *string {
	return s.State
}

func (s *SearchPipelineRequest) SetOwnerAccount(v string) *SearchPipelineRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SearchPipelineRequest) SetOwnerId(v int64) *SearchPipelineRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchPipelineRequest) SetPageNumber(v int64) *SearchPipelineRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchPipelineRequest) SetPageSize(v int64) *SearchPipelineRequest {
	s.PageSize = &v
	return s
}

func (s *SearchPipelineRequest) SetResourceOwnerAccount(v string) *SearchPipelineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SearchPipelineRequest) SetResourceOwnerId(v int64) *SearchPipelineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SearchPipelineRequest) SetState(v string) *SearchPipelineRequest {
	s.State = &v
	return s
}

func (s *SearchPipelineRequest) Validate() error {
	return dara.Validate(s)
}

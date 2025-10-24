// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchWaterMarkTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *SearchWaterMarkTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SearchWaterMarkTemplateRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *SearchWaterMarkTemplateRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *SearchWaterMarkTemplateRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *SearchWaterMarkTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SearchWaterMarkTemplateRequest
	GetResourceOwnerId() *int64
	SetState(v string) *SearchWaterMarkTemplateRequest
	GetState() *string
}

type SearchWaterMarkTemplateRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
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
	// The state of the watermark templates that you want to query. Valid values:
	//
	// 	- **All (default)**
	//
	// 	- **Normal**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SearchWaterMarkTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchWaterMarkTemplateRequest) GoString() string {
	return s.String()
}

func (s *SearchWaterMarkTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SearchWaterMarkTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SearchWaterMarkTemplateRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchWaterMarkTemplateRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SearchWaterMarkTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SearchWaterMarkTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SearchWaterMarkTemplateRequest) GetState() *string {
	return s.State
}

func (s *SearchWaterMarkTemplateRequest) SetOwnerAccount(v string) *SearchWaterMarkTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SearchWaterMarkTemplateRequest) SetOwnerId(v int64) *SearchWaterMarkTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchWaterMarkTemplateRequest) SetPageNumber(v int64) *SearchWaterMarkTemplateRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchWaterMarkTemplateRequest) SetPageSize(v int64) *SearchWaterMarkTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *SearchWaterMarkTemplateRequest) SetResourceOwnerAccount(v string) *SearchWaterMarkTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SearchWaterMarkTemplateRequest) SetResourceOwnerId(v int64) *SearchWaterMarkTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SearchWaterMarkTemplateRequest) SetState(v string) *SearchWaterMarkTemplateRequest {
	s.State = &v
	return s
}

func (s *SearchWaterMarkTemplateRequest) Validate() error {
	return dara.Validate(s)
}

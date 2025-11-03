// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersistentAppInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ListPersistentAppInstancesRequest
	GetAppInstanceGroupId() *string
	SetAppInstancePersistentIds(v []*string) *ListPersistentAppInstancesRequest
	GetAppInstancePersistentIds() []*string
	SetPageNumber(v int32) *ListPersistentAppInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPersistentAppInstancesRequest
	GetPageSize() *int32
	SetProductType(v string) *ListPersistentAppInstancesRequest
	GetProductType() *string
}

type ListPersistentAppInstancesRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-0bxls9m9arax6****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The IDs of the persistent sessions.
	//
	// if can be null:
	// false
	AppInstancePersistentIds []*string `json:"AppInstancePersistentIds,omitempty" xml:"AppInstancePersistentIds,omitempty" type:"Repeated"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The value cannot be greater than **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// Valid values:
	//
	// 	- CloudApp: App Streaming
	//
	// 	- CloudBrowser: Cloud-based Browser
	//
	// 	- AndroidCloud: Cloud Phone
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListPersistentAppInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPersistentAppInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListPersistentAppInstancesRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListPersistentAppInstancesRequest) GetAppInstancePersistentIds() []*string {
	return s.AppInstancePersistentIds
}

func (s *ListPersistentAppInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPersistentAppInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPersistentAppInstancesRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListPersistentAppInstancesRequest) SetAppInstanceGroupId(v string) *ListPersistentAppInstancesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListPersistentAppInstancesRequest) SetAppInstancePersistentIds(v []*string) *ListPersistentAppInstancesRequest {
	s.AppInstancePersistentIds = v
	return s
}

func (s *ListPersistentAppInstancesRequest) SetPageNumber(v int32) *ListPersistentAppInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersistentAppInstancesRequest) SetPageSize(v int32) *ListPersistentAppInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersistentAppInstancesRequest) SetProductType(v string) *ListPersistentAppInstancesRequest {
	s.ProductType = &v
	return s
}

func (s *ListPersistentAppInstancesRequest) Validate() error {
	return dara.Validate(s)
}

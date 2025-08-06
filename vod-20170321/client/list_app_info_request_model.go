// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListAppInfoRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListAppInfoRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListAppInfoRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListAppInfoRequest
	GetStatus() *string
}

type ListAppInfoRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **10**. Maximum value: **100**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID to which the instance belongs.
	//
	// example:
	//
	// rg-aekzko7fsuj****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the application. You can specify the status of the applications that you want to query. After an application is created, it enters the **Normal*	- state. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Disable**
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppInfoRequest) GoString() string {
	return s.String()
}

func (s *ListAppInfoRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListAppInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInfoRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAppInfoRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAppInfoRequest) SetPageNo(v int32) *ListAppInfoRequest {
	s.PageNo = &v
	return s
}

func (s *ListAppInfoRequest) SetPageSize(v int32) *ListAppInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInfoRequest) SetResourceGroupId(v string) *ListAppInfoRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAppInfoRequest) SetStatus(v string) *ListAppInfoRequest {
	s.Status = &v
	return s
}

func (s *ListAppInfoRequest) Validate() error {
	return dara.Validate(s)
}

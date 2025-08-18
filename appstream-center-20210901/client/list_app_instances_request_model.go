// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ListAppInstancesRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceId(v string) *ListAppInstancesRequest
	GetAppInstanceId() *string
	SetAppInstanceIdList(v []*string) *ListAppInstancesRequest
	GetAppInstanceIdList() []*string
	SetIncludeDeleted(v bool) *ListAppInstancesRequest
	GetIncludeDeleted() *bool
	SetPageNumber(v int32) *ListAppInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppInstancesRequest
	GetPageSize() *int32
	SetStatus(v []*string) *ListAppInstancesRequest
	GetStatus() []*string
	SetUserIdList(v []*string) *ListAppInstancesRequest
	GetUserIdList() []*string
}

type ListAppInstancesRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-4p5f8tj16yb8b****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the application instance.
	//
	// example:
	//
	// ai-azn3kmwruh1vl****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The IDs of the application instances. Up to 100 IDs can be specified.
	//
	// if can be null:
	// false
	AppInstanceIdList []*string `json:"AppInstanceIdList,omitempty" xml:"AppInstanceIdList,omitempty" type:"Repeated"`
	// Specifies whether to query the information about deleted app instances. If you set this parameter to true, you must configure AppInstanceIdList. Otherwise, a parameter error is reported.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IncludeDeleted *bool `json:"IncludeDeleted,omitempty" xml:"IncludeDeleted,omitempty"`
	// The page number. Default value: `1`. We recommend that you specify this parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The value cannot be greater than `100`. Default value: `20`. We recommend that you specify this parameter.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the application instances.
	//
	// if can be null:
	// false
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The user IDs. You can specify up to 100 IDs.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s ListAppInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstancesRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAppInstancesRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *ListAppInstancesRequest) GetAppInstanceIdList() []*string {
	return s.AppInstanceIdList
}

func (s *ListAppInstancesRequest) GetIncludeDeleted() *bool {
	return s.IncludeDeleted
}

func (s *ListAppInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstancesRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListAppInstancesRequest) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *ListAppInstancesRequest) SetAppInstanceGroupId(v string) *ListAppInstancesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstancesRequest) SetAppInstanceId(v string) *ListAppInstancesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *ListAppInstancesRequest) SetAppInstanceIdList(v []*string) *ListAppInstancesRequest {
	s.AppInstanceIdList = v
	return s
}

func (s *ListAppInstancesRequest) SetIncludeDeleted(v bool) *ListAppInstancesRequest {
	s.IncludeDeleted = &v
	return s
}

func (s *ListAppInstancesRequest) SetPageNumber(v int32) *ListAppInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstancesRequest) SetPageSize(v int32) *ListAppInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstancesRequest) SetStatus(v []*string) *ListAppInstancesRequest {
	s.Status = v
	return s
}

func (s *ListAppInstancesRequest) SetUserIdList(v []*string) *ListAppInstancesRequest {
	s.UserIdList = v
	return s
}

func (s *ListAppInstancesRequest) Validate() error {
	return dara.Validate(s)
}

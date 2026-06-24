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
	// The delivery group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-4p5f8tj16yb8b****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The application instance ID.
	//
	// example:
	//
	// ai-azn3kmwruh1vl****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The list of application instance IDs. A maximum of 100 IDs can be specified.
	//
	// if can be null:
	// false
	AppInstanceIdList []*string `json:"AppInstanceIdList,omitempty" xml:"AppInstanceIdList,omitempty" type:"Repeated"`
	// Specifies whether to query information about deleted instances. If this parameter is set to true, the AppInstanceIdList parameter is required. Otherwise, a parameter error is returned.
	//
	// example:
	//
	// true
	IncludeDeleted *bool `json:"IncludeDeleted,omitempty" xml:"IncludeDeleted,omitempty"`
	// The page number of the query results to display. Default value: `1`. Specify this parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of query results per page. Maximum value: `100`. Default value: `20`. Specify this parameter.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of application instance statuses.
	//
	// if can be null:
	// false
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The list of user IDs. A maximum of 100 IDs can be specified.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIdList(v []*string) *ListBindInfoRequest
	GetAppIdList() []*string
	SetAppInstanceGroupIdList(v []*string) *ListBindInfoRequest
	GetAppInstanceGroupIdList() []*string
	SetAppInstanceIdList(v []*string) *ListBindInfoRequest
	GetAppInstanceIdList() []*string
	SetPageNumber(v int32) *ListBindInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBindInfoRequest
	GetPageSize() *int32
	SetUserIdList(v []*string) *ListBindInfoRequest
	GetUserIdList() []*string
	SetWyIdList(v []*string) *ListBindInfoRequest
	GetWyIdList() []*string
}

type ListBindInfoRequest struct {
	// The app IDs. You can specify 1 to 100 IDs.
	//
	// >  If you specify this parameter, only the bindings of the specified apps are returned.
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	// The IDs of the delivery groups. You can specify 1 to 100 IDs.
	//
	// >  If you specify this parameter, only the bindings of the specified delivery groups are returned.
	AppInstanceGroupIdList []*string `json:"AppInstanceGroupIdList,omitempty" xml:"AppInstanceGroupIdList,omitempty" type:"Repeated"`
	// The IDs of app instances. You can specify 1 to 100 IDs.
	//
	// >  If you specify this parameter, only the bindings of the specified app instances are returned.
	AppInstanceIdList []*string `json:"AppInstanceIdList,omitempty" xml:"AppInstanceIdList,omitempty" type:"Repeated"`
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The user IDs. You can specify 1 to 100 IDs.
	//
	// >  If you specify this parameter, only the bindings of the specified users are returned.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
	// The IDs of the Alibaba Cloud Workspace users. You can specify 1 to 100 IDs.
	//
	// >  If you specify this parameter, only the bindings of the specified Alibaba Cloud Workspace users are returned.
	WyIdList []*string `json:"WyIdList,omitempty" xml:"WyIdList,omitempty" type:"Repeated"`
}

func (s ListBindInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBindInfoRequest) GoString() string {
	return s.String()
}

func (s *ListBindInfoRequest) GetAppIdList() []*string {
	return s.AppIdList
}

func (s *ListBindInfoRequest) GetAppInstanceGroupIdList() []*string {
	return s.AppInstanceGroupIdList
}

func (s *ListBindInfoRequest) GetAppInstanceIdList() []*string {
	return s.AppInstanceIdList
}

func (s *ListBindInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBindInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBindInfoRequest) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *ListBindInfoRequest) GetWyIdList() []*string {
	return s.WyIdList
}

func (s *ListBindInfoRequest) SetAppIdList(v []*string) *ListBindInfoRequest {
	s.AppIdList = v
	return s
}

func (s *ListBindInfoRequest) SetAppInstanceGroupIdList(v []*string) *ListBindInfoRequest {
	s.AppInstanceGroupIdList = v
	return s
}

func (s *ListBindInfoRequest) SetAppInstanceIdList(v []*string) *ListBindInfoRequest {
	s.AppInstanceIdList = v
	return s
}

func (s *ListBindInfoRequest) SetPageNumber(v int32) *ListBindInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBindInfoRequest) SetPageSize(v int32) *ListBindInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListBindInfoRequest) SetUserIdList(v []*string) *ListBindInfoRequest {
	s.UserIdList = v
	return s
}

func (s *ListBindInfoRequest) SetWyIdList(v []*string) *ListBindInfoRequest {
	s.WyIdList = v
	return s
}

func (s *ListBindInfoRequest) Validate() error {
	return dara.Validate(s)
}

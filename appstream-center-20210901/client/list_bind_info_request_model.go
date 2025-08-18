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
	AppIdList              []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	AppInstanceGroupIdList []*string `json:"AppInstanceGroupIdList,omitempty" xml:"AppInstanceGroupIdList,omitempty" type:"Repeated"`
	AppInstanceIdList      []*string `json:"AppInstanceIdList,omitempty" xml:"AppInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
	WyIdList   []*string `json:"WyIdList,omitempty" xml:"WyIdList,omitempty" type:"Repeated"`
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

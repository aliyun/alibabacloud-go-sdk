// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLdpsComputeGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListLdpsComputeGroupsResponseBody
	GetAccessDeniedDetail() *string
	SetGroupList(v []*ListLdpsComputeGroupsResponseBodyGroupList) *ListLdpsComputeGroupsResponseBody
	GetGroupList() []*ListLdpsComputeGroupsResponseBodyGroupList
	SetRequestId(v string) *ListLdpsComputeGroupsResponseBody
	GetRequestId() *string
}

type ListLdpsComputeGroupsResponseBody struct {
	AccessDeniedDetail *string                                       `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	GroupList          []*ListLdpsComputeGroupsResponseBodyGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	RequestId          *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLdpsComputeGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLdpsComputeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLdpsComputeGroupsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListLdpsComputeGroupsResponseBody) GetGroupList() []*ListLdpsComputeGroupsResponseBodyGroupList {
	return s.GroupList
}

func (s *ListLdpsComputeGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLdpsComputeGroupsResponseBody) SetAccessDeniedDetail(v string) *ListLdpsComputeGroupsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBody) SetGroupList(v []*ListLdpsComputeGroupsResponseBodyGroupList) *ListLdpsComputeGroupsResponseBody {
	s.GroupList = v
	return s
}

func (s *ListLdpsComputeGroupsResponseBody) SetRequestId(v string) *ListLdpsComputeGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBody) Validate() error {
	if s.GroupList != nil {
		for _, item := range s.GroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLdpsComputeGroupsResponseBodyGroupList struct {
	ExceptionInfo *string                `json:"ExceptionInfo,omitempty" xml:"ExceptionInfo,omitempty"`
	GroupName     *string                `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IsDefault     *bool                  `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Properties    map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	State         *string                `json:"State,omitempty" xml:"State,omitempty"`
	WebUI         *string                `json:"WebUI,omitempty" xml:"WebUI,omitempty"`
}

func (s ListLdpsComputeGroupsResponseBodyGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListLdpsComputeGroupsResponseBodyGroupList) GoString() string {
	return s.String()
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) GetExceptionInfo() *string {
	return s.ExceptionInfo
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) GetState() *string {
	return s.State
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) GetWebUI() *string {
	return s.WebUI
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetExceptionInfo(v string) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.ExceptionInfo = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetGroupName(v string) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.GroupName = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetIsDefault(v bool) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.IsDefault = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetProperties(v map[string]interface{}) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.Properties = v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetState(v string) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.State = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetWebUI(v string) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.WebUI = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) Validate() error {
	return dara.Validate(s)
}

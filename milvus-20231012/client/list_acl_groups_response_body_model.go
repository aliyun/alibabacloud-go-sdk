// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAclGroupsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*ListAclGroupsResponseBodyData) *ListAclGroupsResponseBody
	GetData() []*ListAclGroupsResponseBodyData
	SetHttpStatusCode(v int64) *ListAclGroupsResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *ListAclGroupsResponseBody
	GetRequestId() *string
}

type ListAclGroupsResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string                          `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	Data               []*ListAclGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 66A13C31-473A-5B3A-8974-0B07A40649CF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAclGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAclGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclGroupsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAclGroupsResponseBody) GetData() []*ListAclGroupsResponseBodyData {
	return s.Data
}

func (s *ListAclGroupsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListAclGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAclGroupsResponseBody) SetAccessDeniedDetail(v string) *ListAclGroupsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAclGroupsResponseBody) SetData(v []*ListAclGroupsResponseBodyData) *ListAclGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListAclGroupsResponseBody) SetHttpStatusCode(v int64) *ListAclGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAclGroupsResponseBody) SetRequestId(v string) *ListAclGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclGroupsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAclGroupsResponseBodyData struct {
	Cidrs []*string `json:"cidrs,omitempty" xml:"cidrs,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-10-17T13:53:27Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 376774
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// c-xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 50832118
	Uid *int64 `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListAclGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAclGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAclGroupsResponseBodyData) GetCidrs() []*string {
	return s.Cidrs
}

func (s *ListAclGroupsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAclGroupsResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAclGroupsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListAclGroupsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAclGroupsResponseBodyData) GetUid() *int64 {
	return s.Uid
}

func (s *ListAclGroupsResponseBodyData) SetCidrs(v []*string) *ListAclGroupsResponseBodyData {
	s.Cidrs = v
	return s
}

func (s *ListAclGroupsResponseBodyData) SetCreateTime(v string) *ListAclGroupsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAclGroupsResponseBodyData) SetGroupName(v string) *ListAclGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListAclGroupsResponseBodyData) SetId(v int64) *ListAclGroupsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAclGroupsResponseBodyData) SetInstanceId(v string) *ListAclGroupsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListAclGroupsResponseBodyData) SetUid(v int64) *ListAclGroupsResponseBodyData {
	s.Uid = &v
	return s
}

func (s *ListAclGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

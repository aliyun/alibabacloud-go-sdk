// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAllGroupsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListAllGroupsResponseBody
	GetCode() *string
	SetData(v []*ListAllGroupsResponseBodyData) *ListAllGroupsResponseBody
	GetData() []*ListAllGroupsResponseBodyData
	SetMessage(v string) *ListAllGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllGroupsResponseBody
	GetSuccess() *bool
}

type ListAllGroupsResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListAllGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 608F9CCA-B5EB-3D72-XXXXB25D6D75BDEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllGroupsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAllGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllGroupsResponseBody) GetData() []*ListAllGroupsResponseBodyData {
	return s.Data
}

func (s *ListAllGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllGroupsResponseBody) SetAccessDeniedDetail(v string) *ListAllGroupsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAllGroupsResponseBody) SetCode(v string) *ListAllGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllGroupsResponseBody) SetData(v []*ListAllGroupsResponseBodyData) *ListAllGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListAllGroupsResponseBody) SetMessage(v string) *ListAllGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllGroupsResponseBody) SetRequestId(v string) *ListAllGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllGroupsResponseBody) SetSuccess(v bool) *ListAllGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllGroupsResponseBody) Validate() error {
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

type ListAllGroupsResponseBodyData struct {
	// example:
	//
	// 1789000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// aaa
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 5435
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListAllGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllGroupsResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAllGroupsResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAllGroupsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListAllGroupsResponseBodyData) SetGmtCreate(v string) *ListAllGroupsResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListAllGroupsResponseBodyData) SetGroupName(v string) *ListAllGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListAllGroupsResponseBodyData) SetId(v string) *ListAllGroupsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAllGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

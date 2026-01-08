// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListChatGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListChatGroupResponseBody
	GetCode() *string
	SetData(v *ListChatGroupResponseBodyData) *ListChatGroupResponseBody
	GetData() *ListChatGroupResponseBodyData
	SetMessage(v string) *ListChatGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChatGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListChatGroupResponseBody
	GetSuccess() *bool
}

type ListChatGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListChatGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListChatGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListChatGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatGroupResponseBody) GetData() *ListChatGroupResponseBodyData {
	return s.Data
}

func (s *ListChatGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChatGroupResponseBody) SetAccessDeniedDetail(v string) *ListChatGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListChatGroupResponseBody) SetCode(v string) *ListChatGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatGroupResponseBody) SetData(v *ListChatGroupResponseBodyData) *ListChatGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListChatGroupResponseBody) SetMessage(v string) *ListChatGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatGroupResponseBody) SetRequestId(v string) *ListChatGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatGroupResponseBody) SetSuccess(v bool) *ListChatGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListChatGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChatGroupResponseBodyData struct {
	List []*ListChatGroupResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 51
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListChatGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChatGroupResponseBodyData) GetList() []*ListChatGroupResponseBodyDataList {
	return s.List
}

func (s *ListChatGroupResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListChatGroupResponseBodyData) SetList(v []*ListChatGroupResponseBodyDataList) *ListChatGroupResponseBodyData {
	s.List = v
	return s
}

func (s *ListChatGroupResponseBodyData) SetTotal(v int64) *ListChatGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListChatGroupResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChatGroupResponseBodyDataList struct {
	// example:
	//
	// 8613800**
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// EA30d***
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// ACTIVE
	GroupStatus *string `json:"GroupStatus,omitempty" xml:"GroupStatus,omitempty"`
	// example:
	//
	// https://chat.whatsapp.com/***
	InviteLink *string `json:"InviteLink,omitempty" xml:"InviteLink,omitempty"`
	// example:
	//
	// https://aliyun.com/png.jpg
	ProfilePictureFile *string `json:"ProfilePictureFile,omitempty" xml:"ProfilePictureFile,omitempty"`
	// example:
	//
	// 示例值示例值
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s ListChatGroupResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListChatGroupResponseBodyDataList) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *ListChatGroupResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListChatGroupResponseBodyDataList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListChatGroupResponseBodyDataList) GetGroupStatus() *string {
	return s.GroupStatus
}

func (s *ListChatGroupResponseBodyDataList) GetInviteLink() *string {
	return s.InviteLink
}

func (s *ListChatGroupResponseBodyDataList) GetProfilePictureFile() *string {
	return s.ProfilePictureFile
}

func (s *ListChatGroupResponseBodyDataList) GetSubject() *string {
	return s.Subject
}

func (s *ListChatGroupResponseBodyDataList) SetBusinessNumber(v string) *ListChatGroupResponseBodyDataList {
	s.BusinessNumber = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetDescription(v string) *ListChatGroupResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetGroupId(v string) *ListChatGroupResponseBodyDataList {
	s.GroupId = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetGroupStatus(v string) *ListChatGroupResponseBodyDataList {
	s.GroupStatus = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetInviteLink(v string) *ListChatGroupResponseBodyDataList {
	s.InviteLink = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetProfilePictureFile(v string) *ListChatGroupResponseBodyDataList {
	s.ProfilePictureFile = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetSubject(v string) *ListChatGroupResponseBodyDataList {
	s.Subject = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

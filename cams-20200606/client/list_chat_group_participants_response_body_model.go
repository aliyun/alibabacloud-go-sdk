// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatGroupParticipantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListChatGroupParticipantsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListChatGroupParticipantsResponseBody
	GetCode() *string
	SetData(v *ListChatGroupParticipantsResponseBodyData) *ListChatGroupParticipantsResponseBody
	GetData() *ListChatGroupParticipantsResponseBodyData
	SetMessage(v string) *ListChatGroupParticipantsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChatGroupParticipantsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListChatGroupParticipantsResponseBody
	GetSuccess() *bool
}

type ListChatGroupParticipantsResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code. Valid values:
	//
	// - OK: The request was successful.
	//
	// - For other values, see the [error code list](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *ListChatGroupParticipantsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListChatGroupParticipantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupParticipantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatGroupParticipantsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListChatGroupParticipantsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatGroupParticipantsResponseBody) GetData() *ListChatGroupParticipantsResponseBodyData {
	return s.Data
}

func (s *ListChatGroupParticipantsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatGroupParticipantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatGroupParticipantsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChatGroupParticipantsResponseBody) SetAccessDeniedDetail(v string) *ListChatGroupParticipantsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListChatGroupParticipantsResponseBody) SetCode(v string) *ListChatGroupParticipantsResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatGroupParticipantsResponseBody) SetData(v *ListChatGroupParticipantsResponseBodyData) *ListChatGroupParticipantsResponseBody {
	s.Data = v
	return s
}

func (s *ListChatGroupParticipantsResponseBody) SetMessage(v string) *ListChatGroupParticipantsResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatGroupParticipantsResponseBody) SetRequestId(v string) *ListChatGroupParticipantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatGroupParticipantsResponseBody) SetSuccess(v bool) *ListChatGroupParticipantsResponseBody {
	s.Success = &v
	return s
}

func (s *ListChatGroupParticipantsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChatGroupParticipantsResponseBodyData struct {
	// The list of group members.
	List []*ListChatGroupParticipantsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The total number of participants.
	//
	// example:
	//
	// 8
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListChatGroupParticipantsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupParticipantsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChatGroupParticipantsResponseBodyData) GetList() []*ListChatGroupParticipantsResponseBodyDataList {
	return s.List
}

func (s *ListChatGroupParticipantsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListChatGroupParticipantsResponseBodyData) SetList(v []*ListChatGroupParticipantsResponseBodyDataList) *ListChatGroupParticipantsResponseBodyData {
	s.List = v
	return s
}

func (s *ListChatGroupParticipantsResponseBodyData) SetTotal(v int64) *ListChatGroupParticipantsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListChatGroupParticipantsResponseBodyData) Validate() error {
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

type ListChatGroupParticipantsResponseBodyDataList struct {
	// The phone number of the group member.
	//
	// example:
	//
	// 861382***
	ParticipantNumber *string `json:"ParticipantNumber,omitempty" xml:"ParticipantNumber,omitempty"`
}

func (s ListChatGroupParticipantsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupParticipantsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListChatGroupParticipantsResponseBodyDataList) GetParticipantNumber() *string {
	return s.ParticipantNumber
}

func (s *ListChatGroupParticipantsResponseBodyDataList) SetParticipantNumber(v string) *ListChatGroupParticipantsResponseBodyDataList {
	s.ParticipantNumber = &v
	return s
}

func (s *ListChatGroupParticipantsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

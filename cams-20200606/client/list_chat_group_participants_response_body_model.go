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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListChatGroupParticipantsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E9d9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type ListChatGroupParticipantsResponseBodyData struct {
	List []*ListChatGroupParticipantsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type ListChatGroupParticipantsResponseBodyDataList struct {
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

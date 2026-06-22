// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotListConversationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *KopilotListConversationsResponseBody
	GetCode() *int64
	SetData(v *KopilotListConversationsResponseBodyData) *KopilotListConversationsResponseBody
	GetData() *KopilotListConversationsResponseBodyData
	SetRequestId(v string) *KopilotListConversationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *KopilotListConversationsResponseBody
	GetSuccess() *bool
}

type KopilotListConversationsResponseBody struct {
	Code      *int64                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *KopilotListConversationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s KopilotListConversationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *KopilotListConversationsResponseBody) GetData() *KopilotListConversationsResponseBodyData {
	return s.Data
}

func (s *KopilotListConversationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KopilotListConversationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *KopilotListConversationsResponseBody) SetCode(v int64) *KopilotListConversationsResponseBody {
	s.Code = &v
	return s
}

func (s *KopilotListConversationsResponseBody) SetData(v *KopilotListConversationsResponseBodyData) *KopilotListConversationsResponseBody {
	s.Data = v
	return s
}

func (s *KopilotListConversationsResponseBody) SetRequestId(v string) *KopilotListConversationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *KopilotListConversationsResponseBody) SetSuccess(v bool) *KopilotListConversationsResponseBody {
	s.Success = &v
	return s
}

func (s *KopilotListConversationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotListConversationsResponseBodyData struct {
	ConversationIds []*string `json:"ConversationIds,omitempty" xml:"ConversationIds,omitempty" type:"Repeated"`
	Count           *int32    `json:"Count,omitempty" xml:"Count,omitempty"`
	Page            *int32    `json:"Page,omitempty" xml:"Page,omitempty"`
	Size            *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	Total           *int64    `json:"Total,omitempty" xml:"Total,omitempty"`
	TotalPages      *int32    `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	UserId          *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s KopilotListConversationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsResponseBodyData) GetConversationIds() []*string {
	return s.ConversationIds
}

func (s *KopilotListConversationsResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *KopilotListConversationsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *KopilotListConversationsResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *KopilotListConversationsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *KopilotListConversationsResponseBodyData) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *KopilotListConversationsResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *KopilotListConversationsResponseBodyData) SetConversationIds(v []*string) *KopilotListConversationsResponseBodyData {
	s.ConversationIds = v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetCount(v int32) *KopilotListConversationsResponseBodyData {
	s.Count = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetPage(v int32) *KopilotListConversationsResponseBodyData {
	s.Page = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetSize(v int32) *KopilotListConversationsResponseBodyData {
	s.Size = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetTotal(v int64) *KopilotListConversationsResponseBodyData {
	s.Total = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetTotalPages(v int32) *KopilotListConversationsResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) SetUserId(v string) *KopilotListConversationsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *KopilotListConversationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

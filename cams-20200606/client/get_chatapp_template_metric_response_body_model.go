// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappTemplateMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetChatappTemplateMetricResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetChatappTemplateMetricResponseBody
	GetCode() *string
	SetData(v []*GetChatappTemplateMetricResponseBodyData) *GetChatappTemplateMetricResponseBody
	GetData() []*GetChatappTemplateMetricResponseBodyData
	SetMessage(v string) *GetChatappTemplateMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetChatappTemplateMetricResponseBody
	GetRequestId() *string
}

type GetChatappTemplateMetricResponseBody struct {
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
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of returned data.
	Data []*GetChatappTemplateMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 608F9CCA-B5EB-3D72-8047-B25D6D75BDEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatappTemplateMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateMetricResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetChatappTemplateMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetChatappTemplateMetricResponseBody) GetData() []*GetChatappTemplateMetricResponseBodyData {
	return s.Data
}

func (s *GetChatappTemplateMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatappTemplateMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatappTemplateMetricResponseBody) SetAccessDeniedDetail(v string) *GetChatappTemplateMetricResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBody) SetCode(v string) *GetChatappTemplateMetricResponseBody {
	s.Code = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBody) SetData(v []*GetChatappTemplateMetricResponseBodyData) *GetChatappTemplateMetricResponseBody {
	s.Data = v
	return s
}

func (s *GetChatappTemplateMetricResponseBody) SetMessage(v string) *GetChatappTemplateMetricResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBody) SetRequestId(v string) *GetChatappTemplateMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBody) Validate() error {
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

type GetChatappTemplateMetricResponseBodyData struct {
	// The list of button click statistics.
	Cliented []*GetChatappTemplateMetricResponseBodyDataCliented `json:"Cliented,omitempty" xml:"Cliented,omitempty" type:"Repeated"`
	// The number of delivered messages.
	//
	// example:
	//
	// 6
	DeliveredCount *int32 `json:"DeliveredCount,omitempty" xml:"DeliveredCount,omitempty"`
	// The end time of the metric collection period. This value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1668138331485
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The template language.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of read messages.
	//
	// example:
	//
	// 3
	ReadCount *int32 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	// The number of sent messages.
	//
	// example:
	//
	// 10
	SentCount *int32 `json:"SentCount,omitempty" xml:"SentCount,omitempty"`
	// The start time of the metric collection period. This value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1673919240001
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The template code.
	//
	// example:
	//
	// 1100***************
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetChatappTemplateMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateMetricResponseBodyData) GetCliented() []*GetChatappTemplateMetricResponseBodyDataCliented {
	return s.Cliented
}

func (s *GetChatappTemplateMetricResponseBodyData) GetDeliveredCount() *int32 {
	return s.DeliveredCount
}

func (s *GetChatappTemplateMetricResponseBodyData) GetEnd() *int64 {
	return s.End
}

func (s *GetChatappTemplateMetricResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *GetChatappTemplateMetricResponseBodyData) GetReadCount() *int32 {
	return s.ReadCount
}

func (s *GetChatappTemplateMetricResponseBodyData) GetSentCount() *int32 {
	return s.SentCount
}

func (s *GetChatappTemplateMetricResponseBodyData) GetStart() *int64 {
	return s.Start
}

func (s *GetChatappTemplateMetricResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetChatappTemplateMetricResponseBodyData) SetCliented(v []*GetChatappTemplateMetricResponseBodyDataCliented) *GetChatappTemplateMetricResponseBodyData {
	s.Cliented = v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetDeliveredCount(v int32) *GetChatappTemplateMetricResponseBodyData {
	s.DeliveredCount = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetEnd(v int64) *GetChatappTemplateMetricResponseBodyData {
	s.End = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetLanguage(v string) *GetChatappTemplateMetricResponseBodyData {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetReadCount(v int32) *GetChatappTemplateMetricResponseBodyData {
	s.ReadCount = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetSentCount(v int32) *GetChatappTemplateMetricResponseBodyData {
	s.SentCount = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetStart(v int64) *GetChatappTemplateMetricResponseBodyData {
	s.Start = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) SetTemplateCode(v string) *GetChatappTemplateMetricResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyData) Validate() error {
	if s.Cliented != nil {
		for _, item := range s.Cliented {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetChatappTemplateMetricResponseBodyDataCliented struct {
	// The button text.
	//
	// example:
	//
	// reply
	ButtonContent *string `json:"ButtonContent,omitempty" xml:"ButtonContent,omitempty"`
	// The number of clicks.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The button type. Valid values:
	//
	// - phone_number_button: a phone call button.
	//
	// - url_button: a URL button.
	//
	// - quick_relpy_button: a quick reply button.
	//
	// example:
	//
	// quick_reply_button
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetChatappTemplateMetricResponseBodyDataCliented) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateMetricResponseBodyDataCliented) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateMetricResponseBodyDataCliented) GetButtonContent() *string {
	return s.ButtonContent
}

func (s *GetChatappTemplateMetricResponseBodyDataCliented) GetCount() *int32 {
	return s.Count
}

func (s *GetChatappTemplateMetricResponseBodyDataCliented) GetType() *string {
	return s.Type
}

func (s *GetChatappTemplateMetricResponseBodyDataCliented) SetButtonContent(v string) *GetChatappTemplateMetricResponseBodyDataCliented {
	s.ButtonContent = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyDataCliented) SetCount(v int32) *GetChatappTemplateMetricResponseBodyDataCliented {
	s.Count = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyDataCliented) SetType(v string) *GetChatappTemplateMetricResponseBodyDataCliented {
	s.Type = &v
	return s
}

func (s *GetChatappTemplateMetricResponseBodyDataCliented) Validate() error {
	return dara.Validate(s)
}

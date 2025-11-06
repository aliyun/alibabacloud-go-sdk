// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageByMessageIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryMessageByMessageIdResponseBody
	GetCode() *int32
	SetData(v *QueryMessageByMessageIdResponseBodyData) *QueryMessageByMessageIdResponseBody
	GetData() *QueryMessageByMessageIdResponseBodyData
	SetMessage(v string) *QueryMessageByMessageIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMessageByMessageIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMessageByMessageIdResponseBody
	GetSuccess() *bool
}

type QueryMessageByMessageIdResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryMessageByMessageIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMessageByMessageIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByMessageIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageByMessageIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryMessageByMessageIdResponseBody) GetData() *QueryMessageByMessageIdResponseBodyData {
	return s.Data
}

func (s *QueryMessageByMessageIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMessageByMessageIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMessageByMessageIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMessageByMessageIdResponseBody) SetCode(v int32) *QueryMessageByMessageIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBody) SetData(v *QueryMessageByMessageIdResponseBodyData) *QueryMessageByMessageIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryMessageByMessageIdResponseBody) SetMessage(v string) *QueryMessageByMessageIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBody) SetRequestId(v string) *QueryMessageByMessageIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBody) SetSuccess(v bool) *QueryMessageByMessageIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMessageByMessageIdResponseBodyData struct {
	CurrentPage *int32                                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId      *string                                          `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TotalCount  *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      []*QueryMessageByMessageIdResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Repeated"`
}

func (s QueryMessageByMessageIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByMessageIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMessageByMessageIdResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMessageByMessageIdResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMessageByMessageIdResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryMessageByMessageIdResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryMessageByMessageIdResponseBodyData) GetVoList() []*QueryMessageByMessageIdResponseBodyDataVoList {
	return s.VoList
}

func (s *QueryMessageByMessageIdResponseBodyData) SetCurrentPage(v int32) *QueryMessageByMessageIdResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyData) SetPageSize(v int32) *QueryMessageByMessageIdResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyData) SetTaskId(v string) *QueryMessageByMessageIdResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyData) SetTotalCount(v int32) *QueryMessageByMessageIdResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyData) SetVoList(v []*QueryMessageByMessageIdResponseBodyDataVoList) *QueryMessageByMessageIdResponseBodyData {
	s.VoList = v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyData) Validate() error {
	if s.VoList != nil {
		for _, item := range s.VoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMessageByMessageIdResponseBodyDataVoList struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Body            *string `json:"Body,omitempty" xml:"Body,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	ContentType     *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	CorrelationId   *string `json:"CorrelationId,omitempty" xml:"CorrelationId,omitempty"`
	DeliveryMode    *int32  `json:"DeliveryMode,omitempty" xml:"DeliveryMode,omitempty"`
	ExchangeName    *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Headers         *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	Immediate       *bool   `json:"Immediate,omitempty" xml:"Immediate,omitempty"`
	Mandatory       *bool   `json:"Mandatory,omitempty" xml:"Mandatory,omitempty"`
	MessageId       *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Priority        *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProcessToken    *string `json:"ProcessToken,omitempty" xml:"ProcessToken,omitempty"`
	ReconsumeTimes  *int32  `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	ReplyTo         *string `json:"ReplyTo,omitempty" xml:"ReplyTo,omitempty"`
	RoutingKey      *string `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty"`
	StoreTimestamp  *int64  `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Timestamp       *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryMessageByMessageIdResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByMessageIdResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetAppId() *string {
	return s.AppId
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetBody() *string {
	return s.Body
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetContentType() *string {
	return s.ContentType
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetCorrelationId() *string {
	return s.CorrelationId
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetDeliveryMode() *int32 {
	return s.DeliveryMode
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetExpiration() *string {
	return s.Expiration
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetHeaders() *string {
	return s.Headers
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetImmediate() *bool {
	return s.Immediate
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetMandatory() *bool {
	return s.Mandatory
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetPriority() *int32 {
	return s.Priority
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetProcessToken() *string {
	return s.ProcessToken
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetReconsumeTimes() *int32 {
	return s.ReconsumeTimes
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetReplyTo() *string {
	return s.ReplyTo
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetRoutingKey() *string {
	return s.RoutingKey
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetStoreTimestamp() *int64 {
	return s.StoreTimestamp
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetType() *string {
	return s.Type
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) GetUserId() *string {
	return s.UserId
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetAppId(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.AppId = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetBody(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.Body = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetClusterId(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.ClusterId = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetContentEncoding(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.ContentEncoding = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetContentType(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.ContentType = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetCorrelationId(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.CorrelationId = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetDeliveryMode(v int32) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.DeliveryMode = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetExchangeName(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.ExchangeName = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetExpiration(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.Expiration = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetHeaders(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.Headers = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetImmediate(v bool) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.Immediate = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetMandatory(v bool) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.Mandatory = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetMessageId(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.MessageId = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetPriority(v int32) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.Priority = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetProcessToken(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.ProcessToken = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetReconsumeTimes(v int32) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.ReconsumeTimes = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetReplyTo(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.ReplyTo = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetRoutingKey(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.RoutingKey = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetStoreTimestamp(v int64) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.StoreTimestamp = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetTimestamp(v int64) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.Timestamp = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetType(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.Type = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) SetUserId(v string) *QueryMessageByMessageIdResponseBodyDataVoList {
	s.UserId = &v
	return s
}

func (s *QueryMessageByMessageIdResponseBodyDataVoList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageByQueueNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryMessageByQueueNameResponseBody
	GetCode() *int32
	SetData(v *QueryMessageByQueueNameResponseBodyData) *QueryMessageByQueueNameResponseBody
	GetData() *QueryMessageByQueueNameResponseBodyData
	SetMessage(v string) *QueryMessageByQueueNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMessageByQueueNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMessageByQueueNameResponseBody
	GetSuccess() *bool
}

type QueryMessageByQueueNameResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryMessageByQueueNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMessageByQueueNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByQueueNameResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageByQueueNameResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryMessageByQueueNameResponseBody) GetData() *QueryMessageByQueueNameResponseBodyData {
	return s.Data
}

func (s *QueryMessageByQueueNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMessageByQueueNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMessageByQueueNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMessageByQueueNameResponseBody) SetCode(v int32) *QueryMessageByQueueNameResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBody) SetData(v *QueryMessageByQueueNameResponseBodyData) *QueryMessageByQueueNameResponseBody {
	s.Data = v
	return s
}

func (s *QueryMessageByQueueNameResponseBody) SetMessage(v string) *QueryMessageByQueueNameResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBody) SetRequestId(v string) *QueryMessageByQueueNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBody) SetSuccess(v bool) *QueryMessageByQueueNameResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMessageByQueueNameResponseBodyData struct {
	CurrentPage *int32                                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId      *string                                        `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TotalCount  *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      *QueryMessageByQueueNameResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s QueryMessageByQueueNameResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByQueueNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMessageByQueueNameResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMessageByQueueNameResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMessageByQueueNameResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryMessageByQueueNameResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryMessageByQueueNameResponseBodyData) GetVoList() *QueryMessageByQueueNameResponseBodyDataVoList {
	return s.VoList
}

func (s *QueryMessageByQueueNameResponseBodyData) SetCurrentPage(v int32) *QueryMessageByQueueNameResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyData) SetPageSize(v int32) *QueryMessageByQueueNameResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyData) SetTaskId(v string) *QueryMessageByQueueNameResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyData) SetTotalCount(v int64) *QueryMessageByQueueNameResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyData) SetVoList(v *QueryMessageByQueueNameResponseBodyDataVoList) *QueryMessageByQueueNameResponseBodyData {
	s.VoList = v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMessageByQueueNameResponseBodyDataVoList struct {
	AmqpMessageVO []*QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO `json:"AmqpMessageVO,omitempty" xml:"AmqpMessageVO,omitempty" type:"Repeated"`
}

func (s QueryMessageByQueueNameResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByQueueNameResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *QueryMessageByQueueNameResponseBodyDataVoList) GetAmqpMessageVO() []*QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	return s.AmqpMessageVO
}

func (s *QueryMessageByQueueNameResponseBodyDataVoList) SetAmqpMessageVO(v []*QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) *QueryMessageByQueueNameResponseBodyDataVoList {
	s.AmqpMessageVO = v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoList) Validate() error {
	if s.AmqpMessageVO != nil {
		for _, item := range s.AmqpMessageVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO struct {
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

func (s QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GoString() string {
	return s.String()
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetAppId() *string {
	return s.AppId
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetBody() *string {
	return s.Body
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetContentType() *string {
	return s.ContentType
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetCorrelationId() *string {
	return s.CorrelationId
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetDeliveryMode() *int32 {
	return s.DeliveryMode
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetExpiration() *string {
	return s.Expiration
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetHeaders() *string {
	return s.Headers
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetImmediate() *bool {
	return s.Immediate
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetMandatory() *bool {
	return s.Mandatory
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetPriority() *int32 {
	return s.Priority
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetProcessToken() *string {
	return s.ProcessToken
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetReconsumeTimes() *int32 {
	return s.ReconsumeTimes
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetReplyTo() *string {
	return s.ReplyTo
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetRoutingKey() *string {
	return s.RoutingKey
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetStoreTimestamp() *int64 {
	return s.StoreTimestamp
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetType() *string {
	return s.Type
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) GetUserId() *string {
	return s.UserId
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetAppId(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.AppId = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetBody(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.Body = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetClusterId(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.ClusterId = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetContentEncoding(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.ContentEncoding = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetContentType(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.ContentType = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetCorrelationId(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.CorrelationId = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetDeliveryMode(v int32) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.DeliveryMode = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetExchangeName(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.ExchangeName = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetExpiration(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.Expiration = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetHeaders(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.Headers = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetImmediate(v bool) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.Immediate = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetMandatory(v bool) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.Mandatory = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetMessageId(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.MessageId = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetPriority(v int32) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.Priority = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetProcessToken(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.ProcessToken = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetReconsumeTimes(v int32) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.ReconsumeTimes = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetReplyTo(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.ReplyTo = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetRoutingKey(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.RoutingKey = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetStoreTimestamp(v int64) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.StoreTimestamp = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetTimestamp(v int64) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.Timestamp = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetType(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.Type = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) SetUserId(v string) *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO {
	s.UserId = &v
	return s
}

func (s *QueryMessageByQueueNameResponseBodyDataVoListAmqpMessageVO) Validate() error {
	return dara.Validate(s)
}

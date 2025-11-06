// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody
	GetCode() *int32
	SetData(v *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody
	GetData() *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData
	SetMessage(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody
	GetSuccess() *bool
}

type GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody struct {
	Code      *int32                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) GetData() *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	return s.Data
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) SetCode(v int32) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody {
	s.Code = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) SetData(v *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody {
	s.Data = v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) SetMessage(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody {
	s.Message = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) SetRequestId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) SetSuccess(v bool) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody {
	s.Success = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData struct {
	Code          *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Delay         *string                `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Exchange      *string                `json:"Exchange,omitempty" xml:"Exchange,omitempty"`
	Expiration    *string                `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	MessageId     *string                `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	QueueMsgIdMap map[string]interface{} `json:"QueueMsgIdMap,omitempty" xml:"QueueMsgIdMap,omitempty"`
	RemoteAddress *string                `json:"RemoteAddress,omitempty" xml:"RemoteAddress,omitempty"`
	RoutingKey    *string                `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty"`
	SendErrorInfo *string                `json:"SendErrorInfo,omitempty" xml:"SendErrorInfo,omitempty"`
	TimeStamp     *string                `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	UserId        *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Vhost         *string                `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
	XDelay        *string                `json:"XDelay,omitempty" xml:"XDelay,omitempty"`
}

func (s GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetDelay() *string {
	return s.Delay
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetExchange() *string {
	return s.Exchange
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetExpiration() *string {
	return s.Expiration
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetMessageId() *string {
	return s.MessageId
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetQueueMsgIdMap() map[string]interface{} {
	return s.QueueMsgIdMap
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetRemoteAddress() *string {
	return s.RemoteAddress
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetRoutingKey() *string {
	return s.RoutingKey
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetSendErrorInfo() *string {
	return s.SendErrorInfo
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetVhost() *string {
	return s.Vhost
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) GetXDelay() *string {
	return s.XDelay
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetCode(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetDelay(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.Delay = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetExchange(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.Exchange = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetExpiration(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetMessageId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetQueueMsgIdMap(v map[string]interface{}) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.QueueMsgIdMap = v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetRemoteAddress(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.RemoteAddress = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetRoutingKey(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.RoutingKey = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetSendErrorInfo(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.SendErrorInfo = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetTimeStamp(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.TimeStamp = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetUserId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetVhost(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.Vhost = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) SetXDelay(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData {
	s.XDelay = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagResponseBodyData) Validate() error {
	return dara.Validate(s)
}

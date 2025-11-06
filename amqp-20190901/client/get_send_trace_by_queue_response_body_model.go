// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendTraceByQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetSendTraceByQueueResponseBody
	GetCode() *int32
	SetData(v *GetSendTraceByQueueResponseBodyData) *GetSendTraceByQueueResponseBody
	GetData() *GetSendTraceByQueueResponseBodyData
	SetMessage(v string) *GetSendTraceByQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSendTraceByQueueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSendTraceByQueueResponseBody
	GetSuccess() *bool
}

type GetSendTraceByQueueResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetSendTraceByQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSendTraceByQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByQueueResponseBody) GoString() string {
	return s.String()
}

func (s *GetSendTraceByQueueResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetSendTraceByQueueResponseBody) GetData() *GetSendTraceByQueueResponseBodyData {
	return s.Data
}

func (s *GetSendTraceByQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSendTraceByQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSendTraceByQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSendTraceByQueueResponseBody) SetCode(v int32) *GetSendTraceByQueueResponseBody {
	s.Code = &v
	return s
}

func (s *GetSendTraceByQueueResponseBody) SetData(v *GetSendTraceByQueueResponseBodyData) *GetSendTraceByQueueResponseBody {
	s.Data = v
	return s
}

func (s *GetSendTraceByQueueResponseBody) SetMessage(v string) *GetSendTraceByQueueResponseBody {
	s.Message = &v
	return s
}

func (s *GetSendTraceByQueueResponseBody) SetRequestId(v string) *GetSendTraceByQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSendTraceByQueueResponseBody) SetSuccess(v bool) *GetSendTraceByQueueResponseBody {
	s.Success = &v
	return s
}

func (s *GetSendTraceByQueueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSendTraceByQueueResponseBodyData struct {
	CurrentPage *int32                                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      []*GetSendTraceByQueueResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Repeated"`
}

func (s GetSendTraceByQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSendTraceByQueueResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSendTraceByQueueResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSendTraceByQueueResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetSendTraceByQueueResponseBodyData) GetVoList() []*GetSendTraceByQueueResponseBodyDataVoList {
	return s.VoList
}

func (s *GetSendTraceByQueueResponseBodyData) SetCurrentPage(v int32) *GetSendTraceByQueueResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyData) SetPageSize(v int32) *GetSendTraceByQueueResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyData) SetTotalCount(v int32) *GetSendTraceByQueueResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyData) SetVoList(v []*GetSendTraceByQueueResponseBodyDataVoList) *GetSendTraceByQueueResponseBodyData {
	s.VoList = v
	return s
}

func (s *GetSendTraceByQueueResponseBodyData) Validate() error {
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

type GetSendTraceByQueueResponseBodyDataVoList struct {
	Code                 *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Exchange             *string                `json:"Exchange,omitempty" xml:"Exchange,omitempty"`
	MessageBodyLength    *string                `json:"MessageBodyLength,omitempty" xml:"MessageBodyLength,omitempty"`
	MessageId            *string                `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	MessagePropertiesMap map[string]interface{} `json:"MessagePropertiesMap,omitempty" xml:"MessagePropertiesMap,omitempty"`
	QueueMsgIdMap        map[string]interface{} `json:"QueueMsgIdMap,omitempty" xml:"QueueMsgIdMap,omitempty"`
	RemoteAddress        *string                `json:"RemoteAddress,omitempty" xml:"RemoteAddress,omitempty"`
	RoutingKey           *string                `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty"`
	SendErrorInfo        *string                `json:"SendErrorInfo,omitempty" xml:"SendErrorInfo,omitempty"`
	TimeStamp            *string                `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	UserId               *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Vhost                *string                `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
}

func (s GetSendTraceByQueueResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByQueueResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetCode() *string {
	return s.Code
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetExchange() *string {
	return s.Exchange
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetMessageBodyLength() *string {
	return s.MessageBodyLength
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetMessageId() *string {
	return s.MessageId
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetMessagePropertiesMap() map[string]interface{} {
	return s.MessagePropertiesMap
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetQueueMsgIdMap() map[string]interface{} {
	return s.QueueMsgIdMap
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetRemoteAddress() *string {
	return s.RemoteAddress
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetRoutingKey() *string {
	return s.RoutingKey
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetSendErrorInfo() *string {
	return s.SendErrorInfo
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetUserId() *string {
	return s.UserId
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) GetVhost() *string {
	return s.Vhost
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetCode(v string) *GetSendTraceByQueueResponseBodyDataVoList {
	s.Code = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetExchange(v string) *GetSendTraceByQueueResponseBodyDataVoList {
	s.Exchange = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetMessageBodyLength(v string) *GetSendTraceByQueueResponseBodyDataVoList {
	s.MessageBodyLength = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetMessageId(v string) *GetSendTraceByQueueResponseBodyDataVoList {
	s.MessageId = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetMessagePropertiesMap(v map[string]interface{}) *GetSendTraceByQueueResponseBodyDataVoList {
	s.MessagePropertiesMap = v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetQueueMsgIdMap(v map[string]interface{}) *GetSendTraceByQueueResponseBodyDataVoList {
	s.QueueMsgIdMap = v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetRemoteAddress(v string) *GetSendTraceByQueueResponseBodyDataVoList {
	s.RemoteAddress = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetRoutingKey(v string) *GetSendTraceByQueueResponseBodyDataVoList {
	s.RoutingKey = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetSendErrorInfo(v string) *GetSendTraceByQueueResponseBodyDataVoList {
	s.SendErrorInfo = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetTimeStamp(v string) *GetSendTraceByQueueResponseBodyDataVoList {
	s.TimeStamp = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetUserId(v string) *GetSendTraceByQueueResponseBodyDataVoList {
	s.UserId = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) SetVhost(v string) *GetSendTraceByQueueResponseBodyDataVoList {
	s.Vhost = &v
	return s
}

func (s *GetSendTraceByQueueResponseBodyDataVoList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendTraceByMsgIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetSendTraceByMsgIdResponseBody
	GetCode() *int32
	SetData(v *GetSendTraceByMsgIdResponseBodyData) *GetSendTraceByMsgIdResponseBody
	GetData() *GetSendTraceByMsgIdResponseBodyData
	SetMessage(v string) *GetSendTraceByMsgIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSendTraceByMsgIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSendTraceByMsgIdResponseBody
	GetSuccess() *bool
}

type GetSendTraceByMsgIdResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetSendTraceByMsgIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSendTraceByMsgIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetSendTraceByMsgIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetSendTraceByMsgIdResponseBody) GetData() *GetSendTraceByMsgIdResponseBodyData {
	return s.Data
}

func (s *GetSendTraceByMsgIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSendTraceByMsgIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSendTraceByMsgIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSendTraceByMsgIdResponseBody) SetCode(v int32) *GetSendTraceByMsgIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBody) SetData(v *GetSendTraceByMsgIdResponseBodyData) *GetSendTraceByMsgIdResponseBody {
	s.Data = v
	return s
}

func (s *GetSendTraceByMsgIdResponseBody) SetMessage(v string) *GetSendTraceByMsgIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBody) SetRequestId(v string) *GetSendTraceByMsgIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBody) SetSuccess(v bool) *GetSendTraceByMsgIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSendTraceByMsgIdResponseBodyData struct {
	CurrentPage *int32                                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      []*GetSendTraceByMsgIdResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Repeated"`
}

func (s GetSendTraceByMsgIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByMsgIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSendTraceByMsgIdResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSendTraceByMsgIdResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSendTraceByMsgIdResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetSendTraceByMsgIdResponseBodyData) GetVoList() []*GetSendTraceByMsgIdResponseBodyDataVoList {
	return s.VoList
}

func (s *GetSendTraceByMsgIdResponseBodyData) SetCurrentPage(v int32) *GetSendTraceByMsgIdResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyData) SetPageSize(v int32) *GetSendTraceByMsgIdResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyData) SetTotalCount(v int32) *GetSendTraceByMsgIdResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyData) SetVoList(v []*GetSendTraceByMsgIdResponseBodyDataVoList) *GetSendTraceByMsgIdResponseBodyData {
	s.VoList = v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyData) Validate() error {
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

type GetSendTraceByMsgIdResponseBodyDataVoList struct {
	Code                 *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Exchange             *string                `json:"Exchange,omitempty" xml:"Exchange,omitempty"`
	InstanceId           *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MessageBodyLength    *string                `json:"MessageBodyLength,omitempty" xml:"MessageBodyLength,omitempty"`
	MessagePropertiesMap map[string]interface{} `json:"MessagePropertiesMap,omitempty" xml:"MessagePropertiesMap,omitempty"`
	QueueMsgIdMap        map[string]interface{} `json:"QueueMsgIdMap,omitempty" xml:"QueueMsgIdMap,omitempty"`
	RemoteAddress        *string                `json:"RemoteAddress,omitempty" xml:"RemoteAddress,omitempty"`
	RoutingKey           *string                `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty"`
	SendErrorInfo        *string                `json:"SendErrorInfo,omitempty" xml:"SendErrorInfo,omitempty"`
	TimeStamp            *string                `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	UserId               *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Vhost                *string                `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
}

func (s GetSendTraceByMsgIdResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByMsgIdResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetCode() *string {
	return s.Code
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetExchange() *string {
	return s.Exchange
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetMessageBodyLength() *string {
	return s.MessageBodyLength
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetMessagePropertiesMap() map[string]interface{} {
	return s.MessagePropertiesMap
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetQueueMsgIdMap() map[string]interface{} {
	return s.QueueMsgIdMap
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetRemoteAddress() *string {
	return s.RemoteAddress
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetRoutingKey() *string {
	return s.RoutingKey
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetSendErrorInfo() *string {
	return s.SendErrorInfo
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetUserId() *string {
	return s.UserId
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) GetVhost() *string {
	return s.Vhost
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetCode(v string) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.Code = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetExchange(v string) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.Exchange = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetInstanceId(v string) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.InstanceId = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetMessageBodyLength(v string) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.MessageBodyLength = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetMessagePropertiesMap(v map[string]interface{}) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.MessagePropertiesMap = v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetQueueMsgIdMap(v map[string]interface{}) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.QueueMsgIdMap = v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetRemoteAddress(v string) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.RemoteAddress = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetRoutingKey(v string) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.RoutingKey = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetSendErrorInfo(v string) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.SendErrorInfo = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetTimeStamp(v string) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.TimeStamp = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetUserId(v string) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.UserId = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) SetVhost(v string) *GetSendTraceByMsgIdResponseBodyDataVoList {
	s.Vhost = &v
	return s
}

func (s *GetSendTraceByMsgIdResponseBodyDataVoList) Validate() error {
	return dara.Validate(s)
}

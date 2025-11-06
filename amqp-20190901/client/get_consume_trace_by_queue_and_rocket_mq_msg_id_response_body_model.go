// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumeTraceByQueueAndRocketMqMsgIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody
	GetCode() *int32
	SetData(v []*GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody
	GetData() []*GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData
	SetMessage(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody
	GetSuccess() *bool
}

type GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody struct {
	Code      *int32                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) GetData() []*GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	return s.Data
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) SetCode(v int32) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) SetData(v []*GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) SetMessage(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) SetRequestId(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) SetSuccess(v bool) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) Validate() error {
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

type GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData struct {
	AutoAckTag        *string                `json:"AutoAckTag,omitempty" xml:"AutoAckTag,omitempty"`
	ClientAddress     *string                `json:"ClientAddress,omitempty" xml:"ClientAddress,omitempty"`
	Code              *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	ConsumeType       *string                `json:"ConsumeType,omitempty" xml:"ConsumeType,omitempty"`
	ConsumerTag       *string                `json:"ConsumerTag,omitempty" xml:"ConsumerTag,omitempty"`
	CurrentStatus     *string                `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	DeliveryErrorInfo *string                `json:"DeliveryErrorInfo,omitempty" xml:"DeliveryErrorInfo,omitempty"`
	DeliveryTag       *string                `json:"DeliveryTag,omitempty" xml:"DeliveryTag,omitempty"`
	DlqQueueMsgIdMap  map[string]interface{} `json:"DlqQueueMsgIdMap,omitempty" xml:"DlqQueueMsgIdMap,omitempty"`
	Reason            *string                `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ShowAckIcon       *bool                  `json:"ShowAckIcon,omitempty" xml:"ShowAckIcon,omitempty"`
	TimeStamp         *string                `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	UserId            *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetAutoAckTag() *string {
	return s.AutoAckTag
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetClientAddress() *string {
	return s.ClientAddress
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetConsumeType() *string {
	return s.ConsumeType
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetConsumerTag() *string {
	return s.ConsumerTag
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetDeliveryErrorInfo() *string {
	return s.DeliveryErrorInfo
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetDeliveryTag() *string {
	return s.DeliveryTag
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetDlqQueueMsgIdMap() map[string]interface{} {
	return s.DlqQueueMsgIdMap
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetShowAckIcon() *bool {
	return s.ShowAckIcon
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetAutoAckTag(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.AutoAckTag = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetClientAddress(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.ClientAddress = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetCode(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetConsumeType(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.ConsumeType = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetConsumerTag(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.ConsumerTag = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetCurrentStatus(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.CurrentStatus = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetDeliveryErrorInfo(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.DeliveryErrorInfo = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetDeliveryTag(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.DeliveryTag = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetDlqQueueMsgIdMap(v map[string]interface{}) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.DlqQueueMsgIdMap = v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetReason(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.Reason = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetShowAckIcon(v bool) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.ShowAckIcon = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetTimeStamp(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.TimeStamp = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) SetUserId(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}

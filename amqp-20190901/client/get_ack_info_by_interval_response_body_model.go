// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAckInfoByIntervalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAckInfoByIntervalResponseBody
	GetCode() *int32
	SetData(v *GetAckInfoByIntervalResponseBodyData) *GetAckInfoByIntervalResponseBody
	GetData() *GetAckInfoByIntervalResponseBodyData
	SetMessage(v string) *GetAckInfoByIntervalResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAckInfoByIntervalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAckInfoByIntervalResponseBody
	GetSuccess() *bool
}

type GetAckInfoByIntervalResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAckInfoByIntervalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAckInfoByIntervalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAckInfoByIntervalResponseBody) GoString() string {
	return s.String()
}

func (s *GetAckInfoByIntervalResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAckInfoByIntervalResponseBody) GetData() *GetAckInfoByIntervalResponseBodyData {
	return s.Data
}

func (s *GetAckInfoByIntervalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAckInfoByIntervalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAckInfoByIntervalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAckInfoByIntervalResponseBody) SetCode(v int32) *GetAckInfoByIntervalResponseBody {
	s.Code = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBody) SetData(v *GetAckInfoByIntervalResponseBodyData) *GetAckInfoByIntervalResponseBody {
	s.Data = v
	return s
}

func (s *GetAckInfoByIntervalResponseBody) SetMessage(v string) *GetAckInfoByIntervalResponseBody {
	s.Message = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBody) SetRequestId(v string) *GetAckInfoByIntervalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBody) SetSuccess(v bool) *GetAckInfoByIntervalResponseBody {
	s.Success = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAckInfoByIntervalResponseBodyData struct {
	CurrentPage *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      []*GetAckInfoByIntervalResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Repeated"`
}

func (s GetAckInfoByIntervalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAckInfoByIntervalResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAckInfoByIntervalResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAckInfoByIntervalResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAckInfoByIntervalResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAckInfoByIntervalResponseBodyData) GetVoList() []*GetAckInfoByIntervalResponseBodyDataVoList {
	return s.VoList
}

func (s *GetAckInfoByIntervalResponseBodyData) SetCurrentPage(v int32) *GetAckInfoByIntervalResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyData) SetPageSize(v int32) *GetAckInfoByIntervalResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyData) SetTotalCount(v int32) *GetAckInfoByIntervalResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyData) SetVoList(v []*GetAckInfoByIntervalResponseBodyDataVoList) *GetAckInfoByIntervalResponseBodyData {
	s.VoList = v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyData) Validate() error {
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

type GetAckInfoByIntervalResponseBodyDataVoList struct {
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ChannelId     *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ConnectionId  *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	DeliveryTag   *int64  `json:"DeliveryTag,omitempty" xml:"DeliveryTag,omitempty"`
	QueueName     *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RocketMqMsgId *string `json:"RocketMqMsgId,omitempty" xml:"RocketMqMsgId,omitempty"`
	Rt            *int32  `json:"Rt,omitempty" xml:"Rt,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s GetAckInfoByIntervalResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s GetAckInfoByIntervalResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) GetAction() *string {
	return s.Action
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) GetChannelId() *string {
	return s.ChannelId
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) GetDeliveryTag() *int64 {
	return s.DeliveryTag
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) GetQueueName() *string {
	return s.QueueName
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) GetRocketMqMsgId() *string {
	return s.RocketMqMsgId
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) GetRt() *int32 {
	return s.Rt
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) SetAction(v string) *GetAckInfoByIntervalResponseBodyDataVoList {
	s.Action = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) SetChannelId(v string) *GetAckInfoByIntervalResponseBodyDataVoList {
	s.ChannelId = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) SetConnectionId(v string) *GetAckInfoByIntervalResponseBodyDataVoList {
	s.ConnectionId = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) SetDeliveryTag(v int64) *GetAckInfoByIntervalResponseBodyDataVoList {
	s.DeliveryTag = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) SetQueueName(v string) *GetAckInfoByIntervalResponseBodyDataVoList {
	s.QueueName = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) SetRocketMqMsgId(v string) *GetAckInfoByIntervalResponseBodyDataVoList {
	s.RocketMqMsgId = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) SetRt(v int32) *GetAckInfoByIntervalResponseBodyDataVoList {
	s.Rt = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) SetTimeStamp(v string) *GetAckInfoByIntervalResponseBodyDataVoList {
	s.TimeStamp = &v
	return s
}

func (s *GetAckInfoByIntervalResponseBodyDataVoList) Validate() error {
	return dara.Validate(s)
}

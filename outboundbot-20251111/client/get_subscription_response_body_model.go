// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSubscriptionResponseBody
	GetCode() *string
	SetData(v *GetSubscriptionResponseBodyData) *GetSubscriptionResponseBody
	GetData() *GetSubscriptionResponseBodyData
	SetHttpStatusCode(v int32) *GetSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSubscriptionResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetSubscriptionResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSubscriptionResponseBody
	GetSuccess() *bool
}

type GetSubscriptionResponseBody struct {
	// 返回码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *GetSubscriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP状态码
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 错误信息
	//
	// example:
	//
	// Instance does not exist. Instance=ob-1234567890
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误信息中的变量值列表
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// 请求ID
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否调用成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSubscriptionResponseBody) GetData() *GetSubscriptionResponseBodyData {
	return s.Data
}

func (s *GetSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubscriptionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubscriptionResponseBody) SetCode(v string) *GetSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetData(v *GetSubscriptionResponseBodyData) *GetSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *GetSubscriptionResponseBody) SetHttpStatusCode(v int32) *GetSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetMessage(v string) *GetSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetParams(v []*string) *GetSubscriptionResponseBody {
	s.Params = v
	return s
}

func (s *GetSubscriptionResponseBody) SetRequestId(v string) *GetSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionResponseBody) SetSuccess(v bool) *GetSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubscriptionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSubscriptionResponseBodyData struct {
	// true 表示禁用，false 表示启用
	//
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// 接入点
	//
	// example:
	//
	// rmq-cn-h964u01wh12.cn-hangzhou.rmq.aliyuncs.com:8080
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// 事件列表
	EventList []*GetSubscriptionResponseBodyDataEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// 实例ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// MQ的实例ID
	//
	// example:
	//
	// rmq-cn-3g84vpf3712
	MqInstanceId *string `json:"MqInstanceId,omitempty" xml:"MqInstanceId,omitempty"`
	// 消息队列类型
	//
	// example:
	//
	// ROCKET_MQ_4
	MqType *string `json:"MqType,omitempty" xml:"MqType,omitempty"`
	// 田南+伽雷可斯
	//
	// example:
	//
	// pa44w0rd
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 生产者ID
	//
	// example:
	//
	// GID_123456
	ProducerId *string `json:"ProducerId,omitempty" xml:"ProducerId,omitempty"`
	// 主题
	//
	// example:
	//
	// OUTBOUND_BOT_TOPIC
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// 伽雷可斯
	//
	// example:
	//
	// admin
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetSubscriptionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyData) GetDisabled() *bool {
	return s.Disabled
}

func (s *GetSubscriptionResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetSubscriptionResponseBodyData) GetEventList() []*GetSubscriptionResponseBodyDataEventList {
	return s.EventList
}

func (s *GetSubscriptionResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSubscriptionResponseBodyData) GetMqInstanceId() *string {
	return s.MqInstanceId
}

func (s *GetSubscriptionResponseBodyData) GetMqType() *string {
	return s.MqType
}

func (s *GetSubscriptionResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *GetSubscriptionResponseBodyData) GetProducerId() *string {
	return s.ProducerId
}

func (s *GetSubscriptionResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *GetSubscriptionResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetSubscriptionResponseBodyData) SetDisabled(v bool) *GetSubscriptionResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetEndpoint(v string) *GetSubscriptionResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetEventList(v []*GetSubscriptionResponseBodyDataEventList) *GetSubscriptionResponseBodyData {
	s.EventList = v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetInstanceId(v string) *GetSubscriptionResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetMqInstanceId(v string) *GetSubscriptionResponseBodyData {
	s.MqInstanceId = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetMqType(v string) *GetSubscriptionResponseBodyData {
	s.MqType = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetPassword(v string) *GetSubscriptionResponseBodyData {
	s.Password = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetProducerId(v string) *GetSubscriptionResponseBodyData {
	s.ProducerId = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetTopic(v string) *GetSubscriptionResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) SetUserName(v string) *GetSubscriptionResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetSubscriptionResponseBodyData) Validate() error {
	if s.EventList != nil {
		for _, item := range s.EventList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSubscriptionResponseBodyDataEventList struct {
	// true 表示禁用，false 表示启用
	//
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// 显示名称
	//
	// example:
	//
	// 振铃
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 名称
	//
	// example:
	//
	// Ringing
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSubscriptionResponseBodyDataEventList) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionResponseBodyDataEventList) GoString() string {
	return s.String()
}

func (s *GetSubscriptionResponseBodyDataEventList) GetDisabled() *bool {
	return s.Disabled
}

func (s *GetSubscriptionResponseBodyDataEventList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetSubscriptionResponseBodyDataEventList) GetName() *string {
	return s.Name
}

func (s *GetSubscriptionResponseBodyDataEventList) SetDisabled(v bool) *GetSubscriptionResponseBodyDataEventList {
	s.Disabled = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataEventList) SetDisplayName(v string) *GetSubscriptionResponseBodyDataEventList {
	s.DisplayName = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataEventList) SetName(v string) *GetSubscriptionResponseBodyDataEventList {
	s.Name = &v
	return s
}

func (s *GetSubscriptionResponseBodyDataEventList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTopicResponseBody
	GetCode() *string
	SetData(v *GetTopicResponseBodyData) *GetTopicResponseBody
	GetData() *GetTopicResponseBodyData
	SetDynamicCode(v string) *GetTopicResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetTopicResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetTopicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTopicResponseBody
	GetSuccess() *bool
}

type GetTopicResponseBody struct {
	// Error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetTopicResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Dynamic error code.
	//
	// example:
	//
	// TopicName
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// Dynamic error message.
	//
	// example:
	//
	// topicName
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// The topic cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, each request\\"s ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the execution was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTopicResponseBody) GetData() *GetTopicResponseBodyData {
	return s.Data
}

func (s *GetTopicResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetTopicResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicResponseBody) SetCode(v string) *GetTopicResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicResponseBody) SetData(v *GetTopicResponseBodyData) *GetTopicResponseBody {
	s.Data = v
	return s
}

func (s *GetTopicResponseBody) SetDynamicCode(v string) *GetTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetTopicResponseBody) SetDynamicMessage(v string) *GetTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetTopicResponseBody) SetHttpStatusCode(v int32) *GetTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTopicResponseBody) SetMessage(v string) *GetTopicResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicResponseBody) SetRequestId(v string) *GetTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicResponseBody) SetSuccess(v bool) *GetTopicResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTopicResponseBodyData struct {
	// Creation time of the topic.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the instance to which the topic belongs.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The maximum TPS for message sending.
	//
	// example:
	//
	// 1000
	MaxSendTps *int64 `json:"maxSendTps,omitempty" xml:"maxSendTps,omitempty"`
	// The type of messages in the topic.
	//
	// Valid values:
	//
	// 	- TRANSACTION: transactional messages
	//
	// 	- FIFO: ordered messages
	//
	// 	- DELAY: scheduled or delayed messages
	//
	// 	- NORMAL: normal messages
	//
	// Valid values:
	//
	// 	- TRANSACTION: transactional messages
	//
	// 	- FIFO: ordered messages
	//
	// 	- DELAY: scheduled or delayed messages
	//
	// 	- NORMAL: normal messages
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The region ID to which the instance belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Remark information of the topic.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The topic status.
	//
	// Valid values:
	//
	// 	- RUNNING
	//
	// 	- CREATING
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Topic name.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// Last modification time of the topic.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetTopicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTopicResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTopicResponseBodyData) GetMaxSendTps() *int64 {
	return s.MaxSendTps
}

func (s *GetTopicResponseBodyData) GetMessageType() *string {
	return s.MessageType
}

func (s *GetTopicResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTopicResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *GetTopicResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTopicResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *GetTopicResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetTopicResponseBodyData) SetCreateTime(v string) *GetTopicResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTopicResponseBodyData) SetInstanceId(v string) *GetTopicResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetTopicResponseBodyData) SetMaxSendTps(v int64) *GetTopicResponseBodyData {
	s.MaxSendTps = &v
	return s
}

func (s *GetTopicResponseBodyData) SetMessageType(v string) *GetTopicResponseBodyData {
	s.MessageType = &v
	return s
}

func (s *GetTopicResponseBodyData) SetRegionId(v string) *GetTopicResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetTopicResponseBodyData) SetRemark(v string) *GetTopicResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetTopicResponseBodyData) SetStatus(v string) *GetTopicResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTopicResponseBodyData) SetTopicName(v string) *GetTopicResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetTopicResponseBodyData) SetUpdateTime(v string) *GetTopicResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetTopicResponseBodyData) Validate() error {
	return dara.Validate(s)
}

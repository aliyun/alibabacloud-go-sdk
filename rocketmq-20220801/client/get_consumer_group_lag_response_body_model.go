// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerGroupLagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConsumerGroupLagResponseBody
	GetCode() *string
	SetData(v *GetConsumerGroupLagResponseBodyData) *GetConsumerGroupLagResponseBody
	GetData() *GetConsumerGroupLagResponseBodyData
	SetDynamicCode(v string) *GetConsumerGroupLagResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetConsumerGroupLagResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetConsumerGroupLagResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetConsumerGroupLagResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConsumerGroupLagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConsumerGroupLagResponseBody
	GetSuccess() *bool
}

type GetConsumerGroupLagResponseBody struct {
	// Error code
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetConsumerGroupLagResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Dynamic error code
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5764C40-FB8C-53AE-B95D-96AB3D0E9375
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConsumerGroupLagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupLagResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupLagResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConsumerGroupLagResponseBody) GetData() *GetConsumerGroupLagResponseBodyData {
	return s.Data
}

func (s *GetConsumerGroupLagResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetConsumerGroupLagResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetConsumerGroupLagResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetConsumerGroupLagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsumerGroupLagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumerGroupLagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConsumerGroupLagResponseBody) SetCode(v string) *GetConsumerGroupLagResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetData(v *GetConsumerGroupLagResponseBodyData) *GetConsumerGroupLagResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetDynamicCode(v string) *GetConsumerGroupLagResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetDynamicMessage(v string) *GetConsumerGroupLagResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetHttpStatusCode(v int32) *GetConsumerGroupLagResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetMessage(v string) *GetConsumerGroupLagResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetRequestId(v string) *GetConsumerGroupLagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetSuccess(v bool) *GetConsumerGroupLagResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetConsumerGroupLagResponseBodyData struct {
	// Consumer Group ID
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// Instance ID
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Backlog for each topic
	TopicLagMap map[string]*DataTopicLagMapValue `json:"topicLagMap,omitempty" xml:"topicLagMap,omitempty"`
	// Total lag count
	TotalLag *GetConsumerGroupLagResponseBodyDataTotalLag `json:"totalLag,omitempty" xml:"totalLag,omitempty" type:"Struct"`
}

func (s GetConsumerGroupLagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupLagResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupLagResponseBodyData) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *GetConsumerGroupLagResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConsumerGroupLagResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConsumerGroupLagResponseBodyData) GetTopicLagMap() map[string]*DataTopicLagMapValue {
	return s.TopicLagMap
}

func (s *GetConsumerGroupLagResponseBodyData) GetTotalLag() *GetConsumerGroupLagResponseBodyDataTotalLag {
	return s.TotalLag
}

func (s *GetConsumerGroupLagResponseBodyData) SetConsumerGroupId(v string) *GetConsumerGroupLagResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyData) SetInstanceId(v string) *GetConsumerGroupLagResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyData) SetRegionId(v string) *GetConsumerGroupLagResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyData) SetTopicLagMap(v map[string]*DataTopicLagMapValue) *GetConsumerGroupLagResponseBodyData {
	s.TopicLagMap = v
	return s
}

func (s *GetConsumerGroupLagResponseBodyData) SetTotalLag(v *GetConsumerGroupLagResponseBodyDataTotalLag) *GetConsumerGroupLagResponseBodyData {
	s.TotalLag = v
	return s
}

func (s *GetConsumerGroupLagResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetConsumerGroupLagResponseBodyDataTotalLag struct {
	// Delivery delay time, in seconds
	//
	// example:
	//
	// 12
	DeliveryDuration *int64 `json:"deliveryDuration,omitempty" xml:"deliveryDuration,omitempty"`
	// The number of messages being consumed.
	//
	// example:
	//
	// 1
	InflightCount *int64 `json:"inflightCount,omitempty" xml:"inflightCount,omitempty"`
	// Last consumption time
	//
	// example:
	//
	// 1735629607846
	LastConsumeTimestamp *int64 `json:"lastConsumeTimestamp,omitempty" xml:"lastConsumeTimestamp,omitempty"`
	// Ready message count
	//
	// example:
	//
	// 1
	ReadyCount *int64 `json:"readyCount,omitempty" xml:"readyCount,omitempty"`
}

func (s GetConsumerGroupLagResponseBodyDataTotalLag) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupLagResponseBodyDataTotalLag) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) GetDeliveryDuration() *int64 {
	return s.DeliveryDuration
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) GetInflightCount() *int64 {
	return s.InflightCount
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) GetLastConsumeTimestamp() *int64 {
	return s.LastConsumeTimestamp
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) GetReadyCount() *int64 {
	return s.ReadyCount
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) SetDeliveryDuration(v int64) *GetConsumerGroupLagResponseBodyDataTotalLag {
	s.DeliveryDuration = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) SetInflightCount(v int64) *GetConsumerGroupLagResponseBodyDataTotalLag {
	s.InflightCount = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) SetLastConsumeTimestamp(v int64) *GetConsumerGroupLagResponseBodyDataTotalLag {
	s.LastConsumeTimestamp = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) SetReadyCount(v int64) *GetConsumerGroupLagResponseBodyDataTotalLag {
	s.ReadyCount = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) Validate() error {
	return dara.Validate(s)
}

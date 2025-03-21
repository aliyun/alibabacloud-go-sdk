// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListTagResourcesRequest struct {
	// The ID of the ApsaraMQ for RocketMQ instance to which the resource whose tags you want to query belongs.
	//
	// > This parameter is required when you query the tags of a topic or a group.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token that determines the start point of the next query.
	//
	// example:
	//
	// caeba0****be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of resource IDs.
	//
	// example:
	//
	// TopicA
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource whose tags you want to query. Valid values:
	//
	// 	- **INSTANCE**
	//
	// 	- **TOPIC**
	//
	// 	- **GROUP**
	//
	// This parameter is required.
	//
	// example:
	//
	// TOPIC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that you want to query. A maximum of 20 tags can be included in the list.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetInstanceId(v string) *ListTagResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of the tag that you want to detach from the resource.
	//
	// 	- If you include this parameter in a request, the value of this parameter cannot be an empty string.
	//
	// 	- The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// CartService
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that you want to query.
	//
	// 	- The value of this parameter can be an empty string.
	//
	// 	- The tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// ServiceA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The token that determines the start point of the next query.
	//
	// example:
	//
	// caeba0****be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 301D2CBE-66F8-403D-AEC0-82582478****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the resource and tags, including the resource ID, the resource type, tag keys, and tag values.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates the ID of the resource.
	//
	// example:
	//
	// TopicA
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource whose tags you want to query.
	//
	// 	- ALIYUN::MQ::INSTANCE: indicates that the resource is a ApsaraMQ for RocketMQ instance.
	//
	// 	- ALIYUN::MQ::TOPIC: indicates that the resource is a topic.
	//
	// 	- ALIYUN::MQ::GROUP: indicates that the resource is a group.
	//
	// example:
	//
	// ALIYUN::MQ::TOPIC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// CartService
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// ServiceA
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetInstanceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.InstanceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type OnsConsumerAccumulateRequest struct {
	// Specifies whether to query the details of each topic to which the consumer group subscribes. Valid values:
	//
	// 	- **true**: The details of each topic are queried. You can obtain the details from the **DetailInTopicList*	- response parameter.
	//
	// 	- **false**: The details of each topic are not queried. This is the default value. If you use this value, the value of the **DetailInTopicList*	- response parameter is empty.
	//
	// example:
	//
	// true
	Detail *bool `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the consumer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_consumer_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsConsumerAccumulateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerAccumulateRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateRequest) SetDetail(v bool) *OnsConsumerAccumulateRequest {
	s.Detail = &v
	return s
}

func (s *OnsConsumerAccumulateRequest) SetGroupId(v string) *OnsConsumerAccumulateRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerAccumulateRequest) SetInstanceId(v string) *OnsConsumerAccumulateRequest {
	s.InstanceId = &v
	return s
}

type OnsConsumerAccumulateResponseBody struct {
	// The message accumulation information about topics to which the specified consumer subscribes.
	Data *OnsConsumerAccumulateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// CE817BFF-B389-43CD-9419-95011AC9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsConsumerAccumulateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerAccumulateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponseBody) SetData(v *OnsConsumerAccumulateResponseBodyData) *OnsConsumerAccumulateResponseBody {
	s.Data = v
	return s
}

func (s *OnsConsumerAccumulateResponseBody) SetRequestId(v string) *OnsConsumerAccumulateResponseBody {
	s.RequestId = &v
	return s
}

type OnsConsumerAccumulateResponseBodyData struct {
	// The transactions per second (TPS) for message consumption performed by consumers in the group.
	//
	// example:
	//
	// 10
	ConsumeTps *float32 `json:"ConsumeTps,omitempty" xml:"ConsumeTps,omitempty"`
	// The consumption latency.
	//
	// example:
	//
	// 10000
	DelayTime *int64 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The information about each topic to which the consumer group subscribes. If the **Detail*	- parameter in the request is set to **false**, the value of this parameter is empty.
	DetailInTopicList *OnsConsumerAccumulateResponseBodyDataDetailInTopicList `json:"DetailInTopicList,omitempty" xml:"DetailInTopicList,omitempty" type:"Struct"`
	// The point in time when the latest message consumed by a consumer in the consumer group was produced.
	//
	// example:
	//
	// 1566231000000
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// Indicates whether the consumer group is online. The consumer group is online if one of the consumers in the group is online. Valid values:
	//
	// 	- **true**: The consumer group is online.
	//
	// 	- **false**: The consumer group is offline.
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The total number of accumulated messages in all topics to which the consumer group subscribes.
	//
	// example:
	//
	// 100
	TotalDiff *int64 `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s OnsConsumerAccumulateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerAccumulateResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponseBodyData) SetConsumeTps(v float32) *OnsConsumerAccumulateResponseBodyData {
	s.ConsumeTps = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetDelayTime(v int64) *OnsConsumerAccumulateResponseBodyData {
	s.DelayTime = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetDetailInTopicList(v *OnsConsumerAccumulateResponseBodyDataDetailInTopicList) *OnsConsumerAccumulateResponseBodyData {
	s.DetailInTopicList = v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetLastTimestamp(v int64) *OnsConsumerAccumulateResponseBodyData {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetOnline(v bool) *OnsConsumerAccumulateResponseBodyData {
	s.Online = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetTotalDiff(v int64) *OnsConsumerAccumulateResponseBodyData {
	s.TotalDiff = &v
	return s
}

type OnsConsumerAccumulateResponseBodyDataDetailInTopicList struct {
	DetailInTopicDo []*OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo `json:"DetailInTopicDo,omitempty" xml:"DetailInTopicDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerAccumulateResponseBodyDataDetailInTopicList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerAccumulateResponseBodyDataDetailInTopicList) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicList) SetDetailInTopicDo(v []*OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) *OnsConsumerAccumulateResponseBodyDataDetailInTopicList {
	s.DetailInTopicDo = v
	return s
}

type OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo struct {
	// The maximum latency of message consumption in the topic.
	//
	// example:
	//
	// 10000
	DelayTime *int64 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The point in time when the latest consumed message in the topic was produced.
	//
	// example:
	//
	// 1566231000000
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The topic name.
	//
	// example:
	//
	// test-mq-topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The number of accumulated messages in the topic.
	//
	// example:
	//
	// 100
	TotalDiff *int64 `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) SetDelayTime(v int64) *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.DelayTime = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) SetLastTimestamp(v int64) *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) SetTopic(v string) *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.Topic = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) SetTotalDiff(v int64) *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.TotalDiff = &v
	return s
}

type OnsConsumerAccumulateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsConsumerAccumulateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsConsumerAccumulateResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerAccumulateResponse) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponse) SetHeaders(v map[string]*string) *OnsConsumerAccumulateResponse {
	s.Headers = v
	return s
}

func (s *OnsConsumerAccumulateResponse) SetStatusCode(v int32) *OnsConsumerAccumulateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsConsumerAccumulateResponse) SetBody(v *OnsConsumerAccumulateResponseBody) *OnsConsumerAccumulateResponse {
	s.Body = v
	return s
}

type OnsConsumerGetConnectionRequest struct {
	// The ID of the consumer group whose client connection status you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_consumer_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsConsumerGetConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerGetConnectionRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionRequest) SetGroupId(v string) *OnsConsumerGetConnectionRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerGetConnectionRequest) SetInstanceId(v string) *OnsConsumerGetConnectionRequest {
	s.InstanceId = &v
	return s
}

type OnsConsumerGetConnectionResponseBody struct {
	// The data returned.
	Data *OnsConsumerGetConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// DE4140C7-F42D-473D-A5FF-B1E31692****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsConsumerGetConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBody) SetData(v *OnsConsumerGetConnectionResponseBodyData) *OnsConsumerGetConnectionResponseBody {
	s.Data = v
	return s
}

func (s *OnsConsumerGetConnectionResponseBody) SetRequestId(v string) *OnsConsumerGetConnectionResponseBody {
	s.RequestId = &v
	return s
}

type OnsConsumerGetConnectionResponseBodyData struct {
	// The client connection information of the consumer group.
	ConnectionList *OnsConsumerGetConnectionResponseBodyDataConnectionList `json:"ConnectionList,omitempty" xml:"ConnectionList,omitempty" type:"Struct"`
	MessageModel   *string                                                 `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
}

func (s OnsConsumerGetConnectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBodyData) SetConnectionList(v *OnsConsumerGetConnectionResponseBodyDataConnectionList) *OnsConsumerGetConnectionResponseBodyData {
	s.ConnectionList = v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyData) SetMessageModel(v string) *OnsConsumerGetConnectionResponseBodyData {
	s.MessageModel = &v
	return s
}

type OnsConsumerGetConnectionResponseBodyDataConnectionList struct {
	ConnectionDo []*OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo `json:"ConnectionDo,omitempty" xml:"ConnectionDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerGetConnectionResponseBodyDataConnectionList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBodyDataConnectionList) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionList) SetConnectionDo(v []*OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) *OnsConsumerGetConnectionResponseBodyDataConnectionList {
	s.ConnectionDo = v
	return s
}

type OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo struct {
	// The IP address and port number of the consumer client.
	//
	// example:
	//
	// 30.5.121.**
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	// The ID of the consumer client.
	//
	// example:
	//
	// 30.5.121.**@24813#-1999745829#-1737591554#453111174894656
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The programming language in which the consumer application was developed.
	//
	// example:
	//
	// JAVA
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The version of the consumer client.
	//
	// example:
	//
	// V4_3_6
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) SetClientAddr(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo {
	s.ClientAddr = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) SetClientId(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo {
	s.ClientId = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) SetLanguage(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo {
	s.Language = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo) SetVersion(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionListConnectionDo {
	s.Version = &v
	return s
}

type OnsConsumerGetConnectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsConsumerGetConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsConsumerGetConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerGetConnectionResponse) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponse) SetHeaders(v map[string]*string) *OnsConsumerGetConnectionResponse {
	s.Headers = v
	return s
}

func (s *OnsConsumerGetConnectionResponse) SetStatusCode(v int32) *OnsConsumerGetConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsConsumerGetConnectionResponse) SetBody(v *OnsConsumerGetConnectionResponseBody) *OnsConsumerGetConnectionResponse {
	s.Body = v
	return s
}

type OnsConsumerResetOffsetRequest struct {
	// The ID of the consumer group whose dead-letter message you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_consumer_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp to which you want to reset the consumer offset. This parameter takes effect only when the **Type*	- parameter is set to **1**. Unit: milliseconds.
	//
	// example:
	//
	// 1591153871000
	ResetTimestamp *int64 `json:"ResetTimestamp,omitempty" xml:"ResetTimestamp,omitempty"`
	// The name of the topic for which you want to reset the consumer offset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq-topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The method that you want to use to clear accumulated messages. Valid values:
	//
	// 	- **0:*	- All accumulated messages are cleared. Messages that are not consumed are ignored. Consumers in the specified consumer group consume only messages that are published to the topic after the specified point in time.
	//
	// If "reconsumeLater" is returned, the accumulated messages are not cleared because the system is retrying to resend the messages to consumers.
	//
	// 	- **1:*	- The messages that were published to the topic before the specified point in time are cleared. You must specify a point in time. Consumers in the specified consumer group consume only the messages that are published to the topic after the specified point in time.
	//
	// You can specify a point in time from the earliest point in time when a message was published to the topic to the most recent point in time when a message was published to the topic. Points in time that are not within the allowed time range are invalid.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OnsConsumerResetOffsetRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerResetOffsetRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerResetOffsetRequest) SetGroupId(v string) *OnsConsumerResetOffsetRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetInstanceId(v string) *OnsConsumerResetOffsetRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetResetTimestamp(v int64) *OnsConsumerResetOffsetRequest {
	s.ResetTimestamp = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetTopic(v string) *OnsConsumerResetOffsetRequest {
	s.Topic = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetType(v int32) *OnsConsumerResetOffsetRequest {
	s.Type = &v
	return s
}

type OnsConsumerResetOffsetResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// D52C68F8-EC5D-4294-BFFF-1A6A25AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsConsumerResetOffsetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerResetOffsetResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerResetOffsetResponseBody) SetRequestId(v string) *OnsConsumerResetOffsetResponseBody {
	s.RequestId = &v
	return s
}

type OnsConsumerResetOffsetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsConsumerResetOffsetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsConsumerResetOffsetResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerResetOffsetResponse) GoString() string {
	return s.String()
}

func (s *OnsConsumerResetOffsetResponse) SetHeaders(v map[string]*string) *OnsConsumerResetOffsetResponse {
	s.Headers = v
	return s
}

func (s *OnsConsumerResetOffsetResponse) SetStatusCode(v int32) *OnsConsumerResetOffsetResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsConsumerResetOffsetResponse) SetBody(v *OnsConsumerResetOffsetResponseBody) *OnsConsumerResetOffsetResponse {
	s.Body = v
	return s
}

type OnsConsumerStatusRequest struct {
	// Specifies whether to query the details of the consumer group. Valid values:
	//
	// 	- **true**: The details of the consumer group are queried. You can obtain details from the **ConsumerConnectionInfoList*	- and **DetailInTopicList*	- response parameters.
	//
	// 	- **false**: The details of the consumer group are not queried. The values of the **ConsumerConnectionInfoList*	- and **DetailInTopicList*	- response parameters are empty. This value is the default value of the Detail parameter.
	//
	// example:
	//
	// true
	Detail *bool `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the consumer group whose details you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to query the information about thread stack traces. Valid values:
	//
	// 	- **true**: The information about thread stack traces is queried. You can obtain the information from the **Jstack*	- response parameter.
	//
	// > If you want to obtain the information about thread stack traces, make sure that the **Detail*	- parameter in the request is set to **true**.
	//
	// 	- **false**: The information about thread stack traces is not queried. The value of the **Jstack*	- response parameter is empty. This value is the default value of the NeedJstack parameter.
	//
	// example:
	//
	// true
	NeedJstack *bool `json:"NeedJstack,omitempty" xml:"NeedJstack,omitempty"`
}

func (s OnsConsumerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusRequest) SetDetail(v bool) *OnsConsumerStatusRequest {
	s.Detail = &v
	return s
}

func (s *OnsConsumerStatusRequest) SetGroupId(v string) *OnsConsumerStatusRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerStatusRequest) SetInstanceId(v string) *OnsConsumerStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerStatusRequest) SetNeedJstack(v bool) *OnsConsumerStatusRequest {
	s.NeedJstack = &v
	return s
}

type OnsConsumerStatusResponseBody struct {
	// The data returned.
	Data *OnsConsumerStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 10EDC518-10E7-4B34-92FB-171235FA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsConsumerStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBody) SetData(v *OnsConsumerStatusResponseBodyData) *OnsConsumerStatusResponseBody {
	s.Data = v
	return s
}

func (s *OnsConsumerStatusResponseBody) SetRequestId(v string) *OnsConsumerStatusResponseBody {
	s.RequestId = &v
	return s
}

type OnsConsumerStatusResponseBodyData struct {
	// The information about online consumers in the consumer group.
	ConnectionSet *OnsConsumerStatusResponseBodyDataConnectionSet `json:"ConnectionSet,omitempty" xml:"ConnectionSet,omitempty" type:"Struct"`
	// The consumption mode. Valid values:
	//
	// 	- **CLUSTERING**: the clustering consumption mode
	//
	// 	- **BROADCASTING**: the broadcasting consumption mode
	//
	// For more information about consumption modes, see [Clustering consumption and broadcasting consumption](https://help.aliyun.com/document_detail/43163.html).
	//
	// example:
	//
	// CLUSTERING
	ConsumeModel *string `json:"ConsumeModel,omitempty" xml:"ConsumeModel,omitempty"`
	// The TPS for message consumption.
	//
	// example:
	//
	// 0
	ConsumeTps *float32 `json:"ConsumeTps,omitempty" xml:"ConsumeTps,omitempty"`
	// The details of online consumers in the consumer group, including the information about the thread stack traces and the consumption response time (RT). If you want to obtain the details of online consumers in the consumer group, make sure that the **Detail*	- parameter in the request is set to **true**. If the Detail parameter is not set to true, the value of this parameter is empty.
	ConsumerConnectionInfoList *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList `json:"ConsumerConnectionInfoList,omitempty" xml:"ConsumerConnectionInfoList,omitempty" type:"Struct"`
	// The maximum latency of message consumption in all topics to which the consumer group subscribes. Unit: milliseconds.
	//
	// example:
	//
	// 100857
	DelayTime *int64 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The information about message consumption by topic. If you want to obtain the information about the consumption status of each topic, make sure that the **Detail*	- parameter in the request is set to **true**. If the Detail parameter is not set to true, the value of this parameter is empty.
	DetailInTopicList *OnsConsumerStatusResponseBodyDataDetailInTopicList `json:"DetailInTopicList,omitempty" xml:"DetailInTopicList,omitempty" type:"Struct"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The most recent point in time when a message was consumed.
	//
	// The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1566883844954
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// Indicates whether the consumer group is online.
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// Indicates whether load balancing is performed as expected. Valid values:
	//
	// 	- **true**: Load balancing is performed as expected.
	//
	// 	- **false**: Load balancing is not performed as expected.
	//
	// example:
	//
	// true
	RebalanceOK *bool `json:"RebalanceOK,omitempty" xml:"RebalanceOK,omitempty"`
	// Indicates whether all consumers in the consumer group subscribe to the same topics and tags.
	//
	// example:
	//
	// true
	SubscriptionSame *bool `json:"SubscriptionSame,omitempty" xml:"SubscriptionSame,omitempty"`
	// The total number of accumulated messages.
	//
	// example:
	//
	// 197
	TotalDiff *int64 `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s OnsConsumerStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyData) SetConnectionSet(v *OnsConsumerStatusResponseBodyDataConnectionSet) *OnsConsumerStatusResponseBodyData {
	s.ConnectionSet = v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetConsumeModel(v string) *OnsConsumerStatusResponseBodyData {
	s.ConsumeModel = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetConsumeTps(v float32) *OnsConsumerStatusResponseBodyData {
	s.ConsumeTps = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetConsumerConnectionInfoList(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) *OnsConsumerStatusResponseBodyData {
	s.ConsumerConnectionInfoList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetDelayTime(v int64) *OnsConsumerStatusResponseBodyData {
	s.DelayTime = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetDetailInTopicList(v *OnsConsumerStatusResponseBodyDataDetailInTopicList) *OnsConsumerStatusResponseBodyData {
	s.DetailInTopicList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetInstanceId(v string) *OnsConsumerStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetLastTimestamp(v int64) *OnsConsumerStatusResponseBodyData {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetOnline(v bool) *OnsConsumerStatusResponseBodyData {
	s.Online = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetRebalanceOK(v bool) *OnsConsumerStatusResponseBodyData {
	s.RebalanceOK = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetSubscriptionSame(v bool) *OnsConsumerStatusResponseBodyData {
	s.SubscriptionSame = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetTotalDiff(v int64) *OnsConsumerStatusResponseBodyData {
	s.TotalDiff = &v
	return s
}

type OnsConsumerStatusResponseBodyDataConnectionSet struct {
	ConnectionDo []*OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo `json:"ConnectionDo,omitempty" xml:"ConnectionDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConnectionSet) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConnectionSet) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSet) SetConnectionDo(v []*OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) *OnsConsumerStatusResponseBodyDataConnectionSet {
	s.ConnectionDo = v
	return s
}

type OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo struct {
	// The IP address and port number of the consumer instance.
	//
	// example:
	//
	// 30.5.121.**
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	// The ID of the consumer instance.
	//
	// example:
	//
	// 30.5.121.**@25560#-1999745829#-1737591554#458773089270275
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The programming language in which the consumer is developed.
	//
	// example:
	//
	// JAVA
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The private or public IP address of the host.
	//
	// example:
	//
	// 42.120.74.**
	RemoteIP *string `json:"RemoteIP,omitempty" xml:"RemoteIP,omitempty"`
	// The version of the consumer client.
	//
	// example:
	//
	// V4_3_6_SNAPSHOT
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetClientAddr(v string) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.ClientAddr = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetClientId(v string) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.ClientId = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetLanguage(v string) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.Language = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetRemoteIP(v string) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.RemoteIP = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetVersion(v string) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.Version = &v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList struct {
	ConsumerConnectionInfoDo []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo `json:"ConsumerConnectionInfoDo,omitempty" xml:"ConsumerConnectionInfoDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetConsumerConnectionInfoDo(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.ConsumerConnectionInfoDo = v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo struct {
	// The ID of the consumer instance.
	//
	// example:
	//
	// ``30.5.**.**``@25560#-1999745829#-1737591554#458773089270275
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The connection information.
	//
	// example:
	//
	// **
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	// The consumption mode. Valid values:
	//
	// 	- **CLUSTERING**: the clustering consumption mode
	//
	// 	- **BROADCASTING**: the broadcasting consumption mode
	//
	// For more information about consumption modes, see [Clustering consumption and broadcasting consumption](https://help.aliyun.com/document_detail/43163.html).
	//
	// example:
	//
	// CLUSTERING
	ConsumeModel *string `json:"ConsumeModel,omitempty" xml:"ConsumeModel,omitempty"`
	// The mode in which the consumer consumes messages. Valid values:
	//
	// 	- **PUSH**: The ApsaraMQ for RocketMQ broker pushes messages to the consumer.
	//
	// 	- **PULL**: The consumer pulls messages from the ApsaraMQ for RocketMQ broker.
	//
	// example:
	//
	// PUSH
	ConsumeType *string `json:"ConsumeType,omitempty" xml:"ConsumeType,omitempty"`
	// The information about thread stack traces. If you want to obtain the information about thread stack traces, make sure that the **NeedJstack*	- parameter in the request is set to **true**. If the NeedJstack parameter is not set to true, the value of this parameter is empty.
	Jstack *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack `json:"Jstack,omitempty" xml:"Jstack,omitempty" type:"Struct"`
	// The programming language that the consumer supports.
	//
	// example:
	//
	// JAVA
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The most recent point in time when a message was consumed.
	//
	// The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1570701368114
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	// The real-time statistics.
	RunningDataList *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList `json:"RunningDataList,omitempty" xml:"RunningDataList,omitempty" type:"Struct"`
	// The earliest point in time when a message was consumed.
	//
	// The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1570701361528
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	// The information about subscriptions.
	SubscriptionSet *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet `json:"SubscriptionSet,omitempty" xml:"SubscriptionSet,omitempty" type:"Struct"`
	// The number of consumer threads.
	//
	// example:
	//
	// 20
	ThreadCount *int32 `json:"ThreadCount,omitempty" xml:"ThreadCount,omitempty"`
	// The version of the consumer client.
	//
	// example:
	//
	// V4_3_6
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetClientId(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.ClientId = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetConnection(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.Connection = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetConsumeModel(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.ConsumeModel = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetConsumeType(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.ConsumeType = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetJstack(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.Jstack = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetLanguage(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.Language = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetLastTimeStamp(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.LastTimeStamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetRunningDataList(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.RunningDataList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetStartTimeStamp(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.StartTimeStamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetSubscriptionSet(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.SubscriptionSet = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetThreadCount(v int32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.ThreadCount = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetVersion(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.Version = &v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack struct {
	ThreadTrackDo []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo `json:"ThreadTrackDo,omitempty" xml:"ThreadTrackDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack) SetThreadTrackDo(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack {
	s.ThreadTrackDo = v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo struct {
	// The name of the thread.
	//
	// example:
	//
	// ConsumeMessageThread_0
	Thread *string `json:"Thread,omitempty" xml:"Thread,omitempty"`
	// The details of thread stack traces.
	TrackList *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList `json:"TrackList,omitempty" xml:"TrackList,omitempty" type:"Struct"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) SetThread(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo {
	s.Thread = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) SetTrackList(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo {
	s.TrackList = v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList struct {
	Track []*string `json:"Track,omitempty" xml:"Track,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList) SetTrack(v []*string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList {
	s.Track = v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList struct {
	ConsumerRunningDataDo []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo `json:"ConsumerRunningDataDo,omitempty" xml:"ConsumerRunningDataDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList) SetConsumerRunningDataDo(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList {
	s.ConsumerRunningDataDo = v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo struct {
	// The number of messages that failed to be consumed each hour.
	//
	// example:
	//
	// 0
	FailedCountPerHour *int64 `json:"FailedCountPerHour,omitempty" xml:"FailedCountPerHour,omitempty"`
	// The TPS for failed message consumption.
	//
	// example:
	//
	// 0
	FailedTps *float32 `json:"FailedTps,omitempty" xml:"FailedTps,omitempty"`
	// The TPS for successful message consumption.
	//
	// example:
	//
	// 0
	OkTps *float32 `json:"OkTps,omitempty" xml:"OkTps,omitempty"`
	// The consumption RT. Unit: milliseconds.
	//
	// example:
	//
	// 0
	Rt *float32 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// The name of the topic to which the consumer subscribes.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetFailedCountPerHour(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.FailedCountPerHour = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetFailedTps(v float32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.FailedTps = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetOkTps(v float32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.OkTps = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetRt(v float32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.Rt = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetTopic(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.Topic = &v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet struct {
	SubscriptionData []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData `json:"SubscriptionData,omitempty" xml:"SubscriptionData,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet) SetSubscriptionData(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet {
	s.SubscriptionData = v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData struct {
	// The expression that is used to specify the tags of messages in the subscribed topic.
	//
	// example:
	//
	// *
	SubString *string `json:"SubString,omitempty" xml:"SubString,omitempty"`
	// The subscription version. The value is of the LONG type and is automatically incremented.
	//
	// example:
	//
	// 1570701364301
	SubVersion *int64 `json:"SubVersion,omitempty" xml:"SubVersion,omitempty"`
	// The information about the tags of the topic to which the consumer subscribes.
	TagsSet *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet `json:"TagsSet,omitempty" xml:"TagsSet,omitempty" type:"Struct"`
	// The name of the topic to which the consumer subscribes.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) SetSubString(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData {
	s.SubString = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) SetSubVersion(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData {
	s.SubVersion = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) SetTagsSet(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData {
	s.TagsSet = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) SetTopic(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData {
	s.Topic = &v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet struct {
	Tag []*string `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet) SetTag(v []*string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet {
	s.Tag = v
	return s
}

type OnsConsumerStatusResponseBodyDataDetailInTopicList struct {
	DetailInTopicDo []*OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo `json:"DetailInTopicDo,omitempty" xml:"DetailInTopicDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataDetailInTopicList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataDetailInTopicList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicList) SetDetailInTopicDo(v []*OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) *OnsConsumerStatusResponseBodyDataDetailInTopicList {
	s.DetailInTopicDo = v
	return s
}

type OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo struct {
	// The latency of message consumption in the topic. Unit: milliseconds.
	//
	// example:
	//
	// 0
	DelayTime *int64 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The most recent point in time when a message was consumed.
	//
	// The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1570701259403
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The topic name.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The number of accumulated messages in the topic.
	//
	// example:
	//
	// 0
	TotalDiff *int64 `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) SetDelayTime(v int64) *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.DelayTime = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) SetLastTimestamp(v int64) *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) SetTopic(v string) *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.Topic = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) SetTotalDiff(v int64) *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.TotalDiff = &v
	return s
}

type OnsConsumerStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsConsumerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsConsumerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponse) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponse) SetHeaders(v map[string]*string) *OnsConsumerStatusResponse {
	s.Headers = v
	return s
}

func (s *OnsConsumerStatusResponse) SetStatusCode(v int32) *OnsConsumerStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsConsumerStatusResponse) SetBody(v *OnsConsumerStatusResponseBody) *OnsConsumerStatusResponse {
	s.Body = v
	return s
}

type OnsConsumerTimeSpanRequest struct {
	// The ID of the consumer group whose reset time range you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The topic to which the consumer group subscribes.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsConsumerTimeSpanRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerTimeSpanRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerTimeSpanRequest) SetGroupId(v string) *OnsConsumerTimeSpanRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerTimeSpanRequest) SetInstanceId(v string) *OnsConsumerTimeSpanRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerTimeSpanRequest) SetTopic(v string) *OnsConsumerTimeSpanRequest {
	s.Topic = &v
	return s
}

type OnsConsumerTimeSpanResponseBody struct {
	// The data returned.
	Data *OnsConsumerTimeSpanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsConsumerTimeSpanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerTimeSpanResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerTimeSpanResponseBody) SetData(v *OnsConsumerTimeSpanResponseBodyData) *OnsConsumerTimeSpanResponseBody {
	s.Data = v
	return s
}

func (s *OnsConsumerTimeSpanResponseBody) SetRequestId(v string) *OnsConsumerTimeSpanResponseBody {
	s.RequestId = &v
	return s
}

type OnsConsumerTimeSpanResponseBodyData struct {
	// The most recent point in time when a message in the topic was consumed by the customer group.
	//
	// example:
	//
	// 1570761026400
	ConsumeTimeStamp *int64 `json:"ConsumeTimeStamp,omitempty" xml:"ConsumeTimeStamp,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The point in time when the earliest stored message was published to the topic.
	//
	// example:
	//
	// 1570761026804
	MaxTimeStamp *int64 `json:"MaxTimeStamp,omitempty" xml:"MaxTimeStamp,omitempty"`
	// The point in time when the most recently stored message was published to the topic.
	//
	// example:
	//
	// 1570701231122
	MinTimeStamp *int64 `json:"MinTimeStamp,omitempty" xml:"MinTimeStamp,omitempty"`
	// The name of the topic that you want to query.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsConsumerTimeSpanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerTimeSpanResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetConsumeTimeStamp(v int64) *OnsConsumerTimeSpanResponseBodyData {
	s.ConsumeTimeStamp = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetInstanceId(v string) *OnsConsumerTimeSpanResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetMaxTimeStamp(v int64) *OnsConsumerTimeSpanResponseBodyData {
	s.MaxTimeStamp = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetMinTimeStamp(v int64) *OnsConsumerTimeSpanResponseBodyData {
	s.MinTimeStamp = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetTopic(v string) *OnsConsumerTimeSpanResponseBodyData {
	s.Topic = &v
	return s
}

type OnsConsumerTimeSpanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsConsumerTimeSpanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsConsumerTimeSpanResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerTimeSpanResponse) GoString() string {
	return s.String()
}

func (s *OnsConsumerTimeSpanResponse) SetHeaders(v map[string]*string) *OnsConsumerTimeSpanResponse {
	s.Headers = v
	return s
}

func (s *OnsConsumerTimeSpanResponse) SetStatusCode(v int32) *OnsConsumerTimeSpanResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsConsumerTimeSpanResponse) SetBody(v *OnsConsumerTimeSpanResponseBody) *OnsConsumerTimeSpanResponse {
	s.Body = v
	return s
}

type OnsDLQMessageGetByIdRequest struct {
	// The ID of the consumer group whose dead-letter message you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the message you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the dead-letter message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0BC16699165C03B925DB8A404E2D****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s OnsDLQMessageGetByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageGetByIdRequest) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdRequest) SetGroupId(v string) *OnsDLQMessageGetByIdRequest {
	s.GroupId = &v
	return s
}

func (s *OnsDLQMessageGetByIdRequest) SetInstanceId(v string) *OnsDLQMessageGetByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessageGetByIdRequest) SetMsgId(v string) *OnsDLQMessageGetByIdRequest {
	s.MsgId = &v
	return s
}

type OnsDLQMessageGetByIdResponseBody struct {
	// The data returned.
	Data *OnsDLQMessageGetByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsDLQMessageGetByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBody) SetData(v *OnsDLQMessageGetByIdResponseBodyData) *OnsDLQMessageGetByIdResponseBody {
	s.Data = v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBody) SetRequestId(v string) *OnsDLQMessageGetByIdResponseBody {
	s.RequestId = &v
	return s
}

type OnsDLQMessageGetByIdResponseBodyData struct {
	// The cyclic redundancy check (CRC) value of the message body.
	//
	// example:
	//
	// 914112295
	BodyCRC *int32 `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	// The producer instance that generated the message.
	//
	// example:
	//
	// ``42.120.**.**``:64646
	BornHost *string `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	// The timestamp that indicates the point in time when the message was generated. Unit: milliseconds.
	//
	// example:
	//
	// 1570761026630
	BornTimestamp *int64 `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the dead-letter message.
	//
	// example:
	//
	// 0BC16699165C03B925DB8A404E2D****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The attributes of the message.
	PropertyList *OnsDLQMessageGetByIdResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	// The number of retries that were performed to send the message to consumers.
	//
	// example:
	//
	// 1
	ReconsumeTimes *int32 `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	// The ApsaraMQ for RocketMQ broker that stores the message.
	//
	// example:
	//
	// 11.220.***.***:10911
	StoreHost *string `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	// The size of the message. Unit: KB.
	//
	// example:
	//
	// 407
	StoreSize *int32 `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	// The timestamp that indicates the point in time when the ApsaraMQ for RocketMQ broker stored the message. Unit: milliseconds.
	//
	// example:
	//
	// 1570761026708
	StoreTimestamp *int64 `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsDLQMessageGetByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetBodyCRC(v int32) *OnsDLQMessageGetByIdResponseBodyData {
	s.BodyCRC = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetBornHost(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetBornTimestamp(v int64) *OnsDLQMessageGetByIdResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetInstanceId(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetMsgId(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetPropertyList(v *OnsDLQMessageGetByIdResponseBodyDataPropertyList) *OnsDLQMessageGetByIdResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetReconsumeTimes(v int32) *OnsDLQMessageGetByIdResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetStoreHost(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetStoreSize(v int32) *OnsDLQMessageGetByIdResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetStoreTimestamp(v int64) *OnsDLQMessageGetByIdResponseBodyData {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetTopic(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.Topic = &v
	return s
}

type OnsDLQMessageGetByIdResponseBodyDataPropertyList struct {
	MessageProperty []*OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty `json:"MessageProperty,omitempty" xml:"MessageProperty,omitempty" type:"Repeated"`
}

func (s OnsDLQMessageGetByIdResponseBodyDataPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyList) SetMessageProperty(v []*OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) *OnsDLQMessageGetByIdResponseBodyDataPropertyList {
	s.MessageProperty = v
	return s
}

type OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty struct {
	// The name of the attribute. Valid values:
	//
	// 	- **TRACE_ON**: indicates whether the message trace exists.
	//
	// 	- **KEYS**: indicates the key of the message.
	//
	// 	- **TAGS**: indicates the tag that is attached to the message.
	//
	// 	- **INSTANCE_ID**: indicates the ID of the instance that contains the message.
	//
	// For information about the terms that are used in ApsaraMQ for RocketMQ, see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// example:
	//
	// TAGS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// TagA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) SetName(v string) *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty {
	s.Name = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) SetValue(v string) *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty {
	s.Value = &v
	return s
}

type OnsDLQMessageGetByIdResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsDLQMessageGetByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsDLQMessageGetByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponse) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponse) SetHeaders(v map[string]*string) *OnsDLQMessageGetByIdResponse {
	s.Headers = v
	return s
}

func (s *OnsDLQMessageGetByIdResponse) SetStatusCode(v int32) *OnsDLQMessageGetByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponse) SetBody(v *OnsDLQMessageGetByIdResponseBody) *OnsDLQMessageGetByIdResponse {
	s.Body = v
	return s
}

type OnsDLQMessagePageQueryByGroupIdRequest struct {
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds. If you specify a valid value for the **TaskId*	- parameter in the request, this parameter does not take effect. The system uses the value of the BeginTime parameter that you specified in the request when you created the specified query task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570723200000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The number of the page to return. Pages start from page 1. Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds. If you specify a valid value for the **TaskId*	- parameter in the request, this parameter does not take effect. The system uses the value of the EndTime parameter that you specified in the request when you created the specified query task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570809600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the consumer group whose dead-letter messages you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the dead-letter messages you want to query belong.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of dead-letter messages to return on each page. Valid values: 5 to 50. Default value: 20. If you specify a valid value for the **TaskId*	- parameter in the request, this parameter does not take effect. The system uses the value of the PageSize parameter that you specified in the request when you created the specified query task.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the query task. The first time you call this operation to query dead-letter messages that are sent to a specified consumer group within a specified time range, this parameter is not required. This parameter is required in subsequent queries for dead-letter messages on a specified page. You can obtain the task ID from the returned result of the first query.
	//
	// example:
	//
	// 0BC1310300002A9F000021E4D7A48346
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetBeginTime(v int64) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetCurrentPage(v int32) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetEndTime(v int64) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.EndTime = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetGroupId(v string) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetInstanceId(v string) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetPageSize(v int32) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.PageSize = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetTaskId(v string) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.TaskId = &v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponseBody struct {
	// The information about dead-letter messages that are queried.
	MsgFoundDo *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo `json:"MsgFoundDo,omitempty" xml:"MsgFoundDo,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// B00CD3C8-D81E-4A41-85E2-38F19252****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBody) SetMsgFoundDo(v *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) *OnsDLQMessagePageQueryByGroupIdResponseBody {
	s.MsgFoundDo = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBody) SetRequestId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 400
	MaxPageCount *int64 `json:"MaxPageCount,omitempty" xml:"MaxPageCount,omitempty"`
	// The information about dead-letter messages that are returned on the current page. The information that is contained in this parameter is the same as the information that is returned by the [OnsDLQMessageGetById](https://help.aliyun.com/document_detail/112667.html) operation.
	MsgFoundList *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList `json:"MsgFoundList,omitempty" xml:"MsgFoundList,omitempty" type:"Struct"`
	// The ID of the query task. The first time you call this operation to query the dead-letter messages that are sent to a specified consumer group within a specified time range, this parameter is returned. You can use the task ID to query the details of dead-letter messages on other returned pages.
	//
	// example:
	//
	// 0BC1310300002A9F000021E4D7A48346
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetCurrentPage(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.CurrentPage = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetMaxPageCount(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.MaxPageCount = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetMsgFoundList(v *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.MsgFoundList = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetTaskId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.TaskId = &v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList struct {
	OnsRestMessageDo []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo `json:"OnsRestMessageDo,omitempty" xml:"OnsRestMessageDo,omitempty" type:"Repeated"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetOnsRestMessageDo(v []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.OnsRestMessageDo = v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo struct {
	// The cyclic redundancy check (CRC) value of the message body.
	//
	// example:
	//
	// 914112295
	BodyCRC *int32 `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	// The producer instance that generated the message.
	//
	// example:
	//
	// 42.120.***.***:59270
	BornHost *string `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	// The timestamp that indicates when the message was produced.
	//
	// example:
	//
	// 1570760999721
	BornTimestamp *int64 `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C05F2*****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The attributes of the message.
	PropertyList *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	// The number of retries that were performed to send the message to consumers.
	//
	// example:
	//
	// 1
	ReconsumeTimes *int32 `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	// The ApsaraMQ for RocketMQ broker that stores the message.
	//
	// example:
	//
	// 11.193.***.***:10911
	StoreHost *string `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	// The size of the message. Unit: KB.
	//
	// example:
	//
	// 406
	StoreSize *int32 `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	// The timestamp that indicates the point in time when the ApsaraMQ for RocketMQ broker stored the message.
	//
	// example:
	//
	// 1570760999811
	StoreTimestamp *int64 `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBodyCRC(v int32) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BodyCRC = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBornHost(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BornHost = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBornTimestamp(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BornTimestamp = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetInstanceId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetMsgId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.MsgId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetPropertyList(v *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.PropertyList = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetReconsumeTimes(v int32) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreHost(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreHost = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreSize(v int32) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreSize = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreTimestamp(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetTopic(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.Topic = &v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList struct {
	MessageProperty []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty `json:"MessageProperty,omitempty" xml:"MessageProperty,omitempty" type:"Repeated"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) SetMessageProperty(v []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList {
	s.MessageProperty = v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty struct {
	// The name of the attribute. Valid values:
	//
	// 	- **TRACE_ON**: indicates whether a trace of the message exists.
	//
	// 	- **KEYS**: indicates the key of the message.
	//
	// 	- **TAGS**: indicates the tag that is attached to the message.
	//
	// 	- **INSTANCE_ID**: indicates the ID of the instance that contains the message.
	//
	// For information about the terms that are used in ApsaraMQ for RocketMQ, see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// example:
	//
	// TAGS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// TagA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) SetName(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty {
	s.Name = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) SetValue(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty {
	s.Value = &v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsDLQMessagePageQueryByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponse) SetHeaders(v map[string]*string) *OnsDLQMessagePageQueryByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponse) SetStatusCode(v int32) *OnsDLQMessagePageQueryByGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponse) SetBody(v *OnsDLQMessagePageQueryByGroupIdResponseBody) *OnsDLQMessagePageQueryByGroupIdResponse {
	s.Body = v
	return s
}

type OnsDLQMessageResendByIdRequest struct {
	// The ID of the consumer group in which you want to query dead-letter messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance in which the dead-letter message you want to query resides.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the dead-letter message that you want to send to a consumer group for consumption.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0BC16699343051CD9F1D798E7734****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s OnsDLQMessageResendByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageResendByIdRequest) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageResendByIdRequest) SetGroupId(v string) *OnsDLQMessageResendByIdRequest {
	s.GroupId = &v
	return s
}

func (s *OnsDLQMessageResendByIdRequest) SetInstanceId(v string) *OnsDLQMessageResendByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessageResendByIdRequest) SetMsgId(v string) *OnsDLQMessageResendByIdRequest {
	s.MsgId = &v
	return s
}

type OnsDLQMessageResendByIdResponseBody struct {
	// The returned messages.
	Data *OnsDLQMessageResendByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// D94CC769-4DC3-4690-A868-9D0631B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsDLQMessageResendByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageResendByIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageResendByIdResponseBody) SetData(v *OnsDLQMessageResendByIdResponseBodyData) *OnsDLQMessageResendByIdResponseBody {
	s.Data = v
	return s
}

func (s *OnsDLQMessageResendByIdResponseBody) SetRequestId(v string) *OnsDLQMessageResendByIdResponseBody {
	s.RequestId = &v
	return s
}

type OnsDLQMessageResendByIdResponseBodyData struct {
	MsgId []*string `json:"MsgId,omitempty" xml:"MsgId,omitempty" type:"Repeated"`
}

func (s OnsDLQMessageResendByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageResendByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageResendByIdResponseBodyData) SetMsgId(v []*string) *OnsDLQMessageResendByIdResponseBodyData {
	s.MsgId = v
	return s
}

type OnsDLQMessageResendByIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsDLQMessageResendByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsDLQMessageResendByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageResendByIdResponse) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageResendByIdResponse) SetHeaders(v map[string]*string) *OnsDLQMessageResendByIdResponse {
	s.Headers = v
	return s
}

func (s *OnsDLQMessageResendByIdResponse) SetStatusCode(v int32) *OnsDLQMessageResendByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsDLQMessageResendByIdResponse) SetBody(v *OnsDLQMessageResendByIdResponseBody) *OnsDLQMessageResendByIdResponse {
	s.Body = v
	return s
}

type OnsGroupConsumerUpdateRequest struct {
	// The ID of the consumer group for which you want to configure read permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_groupId
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group you want to configure belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to authorize the consumer group to read messages. Valid values:
	//
	// 	- **true**: The consumer group can read messages.
	//
	// 	- **false**: The consumer group cannot read messages.
	//
	// Default value: **true**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ReadEnable *bool `json:"ReadEnable,omitempty" xml:"ReadEnable,omitempty"`
}

func (s OnsGroupConsumerUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupConsumerUpdateRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupConsumerUpdateRequest) SetGroupId(v string) *OnsGroupConsumerUpdateRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupConsumerUpdateRequest) SetInstanceId(v string) *OnsGroupConsumerUpdateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupConsumerUpdateRequest) SetReadEnable(v bool) *OnsGroupConsumerUpdateRequest {
	s.ReadEnable = &v
	return s
}

type OnsGroupConsumerUpdateResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsGroupConsumerUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupConsumerUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupConsumerUpdateResponseBody) SetRequestId(v string) *OnsGroupConsumerUpdateResponseBody {
	s.RequestId = &v
	return s
}

type OnsGroupConsumerUpdateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsGroupConsumerUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsGroupConsumerUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupConsumerUpdateResponse) GoString() string {
	return s.String()
}

func (s *OnsGroupConsumerUpdateResponse) SetHeaders(v map[string]*string) *OnsGroupConsumerUpdateResponse {
	s.Headers = v
	return s
}

func (s *OnsGroupConsumerUpdateResponse) SetStatusCode(v int32) *OnsGroupConsumerUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsGroupConsumerUpdateResponse) SetBody(v *OnsGroupConsumerUpdateResponseBody) *OnsGroupConsumerUpdateResponse {
	s.Body = v
	return s
}

type OnsGroupCreateRequest struct {
	// The ID of the consumer group that you want to create. The group ID must meet the following rules:
	//
	// 	- The group ID must be 2 to 64 characters in length and can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- If the ApsaraMQ for RocketMQ instance in which you want to create the consumer group uses a namespace, the group ID must be unique in the instance. The group ID cannot be the same as an existing group ID or a topic name in the instance. The group ID can be the same as a group ID or a topic name in another instance that uses a different namespace. For example, if Instance A and Instance B use different namespaces, a group ID in Instance A can be the same as a group ID or a topic name in Instance B.
	//
	// 	- If the instance does not use a namespace, the group ID must be globally unique across instances and regions. The group ID cannot be the same as an existing group ID or topic name in ApsaraMQ for RocketMQ instances that belong to your Alibaba Cloud account.
	//
	// >
	//
	// 	- After the consumer group is created, the group ID cannot be changed.
	//
	// 	- To check whether an instance uses a namespace, log on to the ApsaraMQ for RocketMQ console, go to the **Instance Details*	- page, and then view the value of the Namespace field in the **Basic Information*	- section.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_groupId
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The protocol over which clients in the consumer group communicate with the ApsaraMQ for RocketMQ broker. All clients in a consumer group communicate with the ApsaraMQ for RocketMQ broker over the same protocol. You must create different groups for TCP clients and HTTP clients. Valid values:
	//
	// 	- **tcp**: Clients in the consumer group consume messages over TCP. This is the default value.
	//
	// 	- **http**: Clients in the consumer group consume messages over HTTP.
	//
	// example:
	//
	// tcp
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The ID of the instance in which you want to create the consumer group.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The description of the consumer group.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s OnsGroupCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupCreateRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupCreateRequest) SetGroupId(v string) *OnsGroupCreateRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupCreateRequest) SetGroupType(v string) *OnsGroupCreateRequest {
	s.GroupType = &v
	return s
}

func (s *OnsGroupCreateRequest) SetInstanceId(v string) *OnsGroupCreateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupCreateRequest) SetRemark(v string) *OnsGroupCreateRequest {
	s.Remark = &v
	return s
}

type OnsGroupCreateResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsGroupCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupCreateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupCreateResponseBody) SetRequestId(v string) *OnsGroupCreateResponseBody {
	s.RequestId = &v
	return s
}

type OnsGroupCreateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsGroupCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsGroupCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupCreateResponse) GoString() string {
	return s.String()
}

func (s *OnsGroupCreateResponse) SetHeaders(v map[string]*string) *OnsGroupCreateResponse {
	s.Headers = v
	return s
}

func (s *OnsGroupCreateResponse) SetStatusCode(v int32) *OnsGroupCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsGroupCreateResponse) SetBody(v *OnsGroupCreateResponseBody) *OnsGroupCreateResponse {
	s.Body = v
	return s
}

type OnsGroupDeleteRequest struct {
	// The ID of the consumer group that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_groupId
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance to which the specified consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsGroupDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupDeleteRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupDeleteRequest) SetGroupId(v string) *OnsGroupDeleteRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupDeleteRequest) SetInstanceId(v string) *OnsGroupDeleteRequest {
	s.InstanceId = &v
	return s
}

type OnsGroupDeleteResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsGroupDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupDeleteResponseBody) SetRequestId(v string) *OnsGroupDeleteResponseBody {
	s.RequestId = &v
	return s
}

type OnsGroupDeleteResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsGroupDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsGroupDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupDeleteResponse) GoString() string {
	return s.String()
}

func (s *OnsGroupDeleteResponse) SetHeaders(v map[string]*string) *OnsGroupDeleteResponse {
	s.Headers = v
	return s
}

func (s *OnsGroupDeleteResponse) SetStatusCode(v int32) *OnsGroupDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsGroupDeleteResponse) SetBody(v *OnsGroupDeleteResponseBody) *OnsGroupDeleteResponse {
	s.Body = v
	return s
}

type OnsGroupListRequest struct {
	// This parameter is required only when you query specific consumer groups by using the fuzzy search method. If this parameter is not configured, the system queries all consumer groups that can be accessed by the current account.
	//
	// If you set this parameter to GID_ABC, the information about the consumer groups whose IDs contain GID_ABC is returned. For example, the information about the GID_test_GID_ABC_123 and GID_ABC_356 consumer groups is returned.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The protocol over which the queried consumer group publishes and subscribes to messages. All clients in a consumer group communicate with the ApsaraMQ for RocketMQ broker over the same protocol. You must create different consumer groups for TCP clients and HTTP clients. Valid values:
	//
	// 	- **tcp**: specifies that the consumer group publishes or subscribes to messages over TCP. This value is the default value.
	//
	// 	- **http**: specifies that the consumer group publishes or subscribes to messages over HTTP.
	//
	// example:
	//
	// tcp
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The ID of the instance to which the consumer group you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of tags that are attached to the consumer group. A maximum of 20 tags can be included in the list.
	Tag []*OnsGroupListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupListRequest) SetGroupId(v string) *OnsGroupListRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupListRequest) SetGroupType(v string) *OnsGroupListRequest {
	s.GroupType = &v
	return s
}

func (s *OnsGroupListRequest) SetInstanceId(v string) *OnsGroupListRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupListRequest) SetTag(v []*OnsGroupListRequestTag) *OnsGroupListRequest {
	s.Tag = v
	return s
}

type OnsGroupListRequestTag struct {
	// The key of the tag that is attached to the consumer group. This parameter is not required. If you configure this parameter, you must configure the **Key*	- parameter.***	- If you configure both the Key and Value parameters, the consumer groups are filtered based on the specified tag. If you do not configure these parameters, all consumer groups are queried.
	//
	// 	- The value of this parameter cannot be an empty string.
	//
	// 	- The tag value must be 1 to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// CartService
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is attached to the group. This parameter is not required. If you configure this parameter, you must configure the **Key*	- parameter.***	- If you configure both the Key and Value parameters, the consumer groups are filtered based on the specified tag. If you do not configure these parameters, all consumer groups are queried.
	//
	// 	- The value of this parameter can be an empty string.
	//
	// 	- The tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ServiceA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsGroupListRequestTag) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListRequestTag) GoString() string {
	return s.String()
}

func (s *OnsGroupListRequestTag) SetKey(v string) *OnsGroupListRequestTag {
	s.Key = &v
	return s
}

func (s *OnsGroupListRequestTag) SetValue(v string) *OnsGroupListRequestTag {
	s.Value = &v
	return s
}

type OnsGroupListResponseBody struct {
	// The returned list of subscriptions.
	Data *OnsGroupListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 16996623-AC4A-43AF-9248-FD9D2D75****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBody) SetData(v *OnsGroupListResponseBodyData) *OnsGroupListResponseBody {
	s.Data = v
	return s
}

func (s *OnsGroupListResponseBody) SetRequestId(v string) *OnsGroupListResponseBody {
	s.RequestId = &v
	return s
}

type OnsGroupListResponseBodyData struct {
	SubscribeInfoDo []*OnsGroupListResponseBodyDataSubscribeInfoDo `json:"SubscribeInfoDo,omitempty" xml:"SubscribeInfoDo,omitempty" type:"Repeated"`
}

func (s OnsGroupListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBodyData) SetSubscribeInfoDo(v []*OnsGroupListResponseBodyDataSubscribeInfoDo) *OnsGroupListResponseBodyData {
	s.SubscribeInfoDo = v
	return s
}

type OnsGroupListResponseBodyDataSubscribeInfoDo struct {
	// The point in time when the consumer group was created.
	//
	// example:
	//
	// 1568896605000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the consumer group.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The protocol over which the queried consumer group publishes and subscribes to messages. All clients in a consumer group communicate with the ApsaraMQ for RocketMQ broker over the same protocol. You must create different consumer groups for TCP clients and HTTP clients. Valid values:
	//
	// 	- **tcp**: indicates that the consumer group publishes and subscribes to messages over TCP.
	//
	// 	- **http**: indicates that the consumer group publishes and subscribes to messages over HTTP.
	//
	// example:
	//
	// tcp
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// Indicates whether the instance uses a namespace. Valid values:
	//
	// 	- **true**: The instance uses a separate namespace. The name of each resource must be unique in the instance. The names of resources in different instances can be the same.
	//
	// 	- **false**: The instance does not use a separate namespace. The name of each resource must be globally unique within the instance and across all instances.
	//
	// example:
	//
	// true
	IndependentNaming *bool `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Alibaba Cloud account ID of the user who created the consumer group.
	//
	// example:
	//
	// 138015630679****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The description of the consumer group.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The tags that are attached to the consumer group.
	Tags *OnsGroupListResponseBodyDataSubscribeInfoDoTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The most recent point in time when the consumer group was updated.
	//
	// example:
	//
	// 1570700979000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDo) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDo) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetCreateTime(v int64) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.CreateTime = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetGroupId(v string) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.GroupId = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetGroupType(v string) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.GroupType = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetIndependentNaming(v bool) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.IndependentNaming = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetInstanceId(v string) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetOwner(v string) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.Owner = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetRemark(v string) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.Remark = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetTags(v *OnsGroupListResponseBodyDataSubscribeInfoDoTags) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.Tags = v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDo) SetUpdateTime(v int64) *OnsGroupListResponseBodyDataSubscribeInfoDo {
	s.UpdateTime = &v
	return s
}

type OnsGroupListResponseBodyDataSubscribeInfoDoTags struct {
	Tag []*OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDoTags) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDoTags) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTags) SetTag(v []*OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) *OnsGroupListResponseBodyDataSubscribeInfoDoTags {
	s.Tag = v
	return s
}

type OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag struct {
	// The key of the tag that is attached to the consumer group.
	//
	// example:
	//
	// CartService
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is attached to the consumer group.
	//
	// example:
	//
	// ServiceA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) SetKey(v string) *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag {
	s.Key = &v
	return s
}

func (s *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag) SetValue(v string) *OnsGroupListResponseBodyDataSubscribeInfoDoTagsTag {
	s.Value = &v
	return s
}

type OnsGroupListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListResponse) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponse) SetHeaders(v map[string]*string) *OnsGroupListResponse {
	s.Headers = v
	return s
}

func (s *OnsGroupListResponse) SetStatusCode(v int32) *OnsGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsGroupListResponse) SetBody(v *OnsGroupListResponseBody) *OnsGroupListResponse {
	s.Body = v
	return s
}

type OnsGroupSubDetailRequest struct {
	// The ID of the consumer group that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group you want to query belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsGroupSubDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupSubDetailRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailRequest) SetGroupId(v string) *OnsGroupSubDetailRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupSubDetailRequest) SetInstanceId(v string) *OnsGroupSubDetailRequest {
	s.InstanceId = &v
	return s
}

type OnsGroupSubDetailResponseBody struct {
	// The data returned.
	Data *OnsGroupSubDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 3364E875-013B-442A-BC3C-C1A84DC6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsGroupSubDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupSubDetailResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBody) SetData(v *OnsGroupSubDetailResponseBodyData) *OnsGroupSubDetailResponseBody {
	s.Data = v
	return s
}

func (s *OnsGroupSubDetailResponseBody) SetRequestId(v string) *OnsGroupSubDetailResponseBody {
	s.RequestId = &v
	return s
}

type OnsGroupSubDetailResponseBodyData struct {
	// The ID of the consumer group that you want to query.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The consumption mode. Valid values:
	//
	// 	- **CLUSTERING**: the clustering consumption mode
	//
	// 	- **BROADCASTING**: the broadcasting consumption mode
	//
	// For more information about consumption modes, see [Clustering consumption and broadcasting consumption](https://help.aliyun.com/document_detail/43163.html).
	//
	// example:
	//
	// CLUSTERING
	MessageModel *string `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	// Indicates whether consumers in the group are online.
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The topics to which consumers in the consumer group subscribe. If all consumers in the specified group are offline, no topics are returned.
	SubscriptionDataList *OnsGroupSubDetailResponseBodyDataSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Struct"`
}

func (s OnsGroupSubDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupSubDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBodyData) SetGroupId(v string) *OnsGroupSubDetailResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *OnsGroupSubDetailResponseBodyData) SetMessageModel(v string) *OnsGroupSubDetailResponseBodyData {
	s.MessageModel = &v
	return s
}

func (s *OnsGroupSubDetailResponseBodyData) SetOnline(v bool) *OnsGroupSubDetailResponseBodyData {
	s.Online = &v
	return s
}

func (s *OnsGroupSubDetailResponseBodyData) SetSubscriptionDataList(v *OnsGroupSubDetailResponseBodyDataSubscriptionDataList) *OnsGroupSubDetailResponseBodyData {
	s.SubscriptionDataList = v
	return s
}

type OnsGroupSubDetailResponseBodyDataSubscriptionDataList struct {
	SubscriptionDataList []*OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Repeated"`
}

func (s OnsGroupSubDetailResponseBodyDataSubscriptionDataList) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupSubDetailResponseBodyDataSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataList) SetSubscriptionDataList(v []*OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) *OnsGroupSubDetailResponseBodyDataSubscriptionDataList {
	s.SubscriptionDataList = v
	return s
}

type OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList struct {
	// The expression based on which consumers in the consumer group subscribe to the topic.
	//
	// example:
	//
	// *
	SubString *string `json:"SubString,omitempty" xml:"SubString,omitempty"`
	// The name of the topic to which consumers in the consumer group subscribe.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetSubString(v string) *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.SubString = &v
	return s
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetTopic(v string) *OnsGroupSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.Topic = &v
	return s
}

type OnsGroupSubDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsGroupSubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsGroupSubDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupSubDetailResponse) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponse) SetHeaders(v map[string]*string) *OnsGroupSubDetailResponse {
	s.Headers = v
	return s
}

func (s *OnsGroupSubDetailResponse) SetStatusCode(v int32) *OnsGroupSubDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsGroupSubDetailResponse) SetBody(v *OnsGroupSubDetailResponseBody) *OnsGroupSubDetailResponse {
	s.Body = v
	return s
}

type OnsInstanceBaseInfoRequest struct {
	// The ID of the instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_138015630679****_BAAy1Hac
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsInstanceBaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoRequest) SetInstanceId(v string) *OnsInstanceBaseInfoRequest {
	s.InstanceId = &v
	return s
}

type OnsInstanceBaseInfoResponseBody struct {
	// The information about the instance.
	InstanceBaseInfo *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo `json:"InstanceBaseInfo,omitempty" xml:"InstanceBaseInfo,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 6CC46974-65E8-4C20-AB07-D20D102E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsInstanceBaseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoResponseBody) SetInstanceBaseInfo(v *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) *OnsInstanceBaseInfoResponseBody {
	s.InstanceBaseInfo = v
	return s
}

func (s *OnsInstanceBaseInfoResponseBody) SetRequestId(v string) *OnsInstanceBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

type OnsInstanceBaseInfoResponseBodyInstanceBaseInfo struct {
	// The time when the instance was created. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1570701259403
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The endpoints used to access ApsaraMQ for RocketMQ over different protocols.
	Endpoints *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// Indicates whether the instance uses a namespace. Valid values:
	//
	// 	- **true**: The instance uses a separate namespace. The name of each resource must be unique in the instance. The names of resources in different instances can be the same.
	//
	// 	- **false**: The instance does not use a separate namespace. The name of each resource must be globally unique within the instance and across all instances.
	//
	// example:
	//
	// true
	IndependentNaming *bool `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_138015630679****_BAAy1Hac
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// The name must be 3 to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **0**: The instance is being deployed. This value is valid only for Enterprise Platinum Edition instances.
	//
	// 	- **2**: The instance has overdue payments. This value is valid only for Standard Edition instances.
	//
	// 	- **5**: The instance is running. This value is valid for Standard Edition instances and Enterprise Platinum Edition instances.
	//
	// 	- **7**: The instance is being upgraded and is running. This value is valid only for Enterprise Platinum Edition instances.
	//
	// example:
	//
	// 5
	InstanceStatus *int32 `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance type. Valid values:
	//
	// 	- **1**: Standard Edition instances that use the pay-as-you-go billing method.
	//
	// 	- **2**: Enterprise Platinum Edition instances that use the subscription billing method.
	//
	// For information about the editions and specifications of ApsaraMQ for RocketMQ instances, see [Instance editions](https://help.aliyun.com/document_detail/185261.html).
	//
	// example:
	//
	// 2
	InstanceType *int32 `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum messaging transactions per second (TPS). Valid values: 5000, 10000, 20000, 50000, 100000, 200000, 300000, 500000, 800000, and 1000000.
	//
	// You can view the details about messaging TPS on the buy page of ApsaraMQ for RocketMQ.
	//
	// > This parameter is available only to the ApsaraMQ for RocketMQ Enterprise Platinum Edition instances.
	//
	// example:
	//
	// 10000
	MaxTps *int64 `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	// The time when the Enterprise Platinum Edition instance expires.
	//
	// example:
	//
	// 1603555200000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// ons-cn-m7r1r5f****
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The maximum number of topics that can be created on the instance. Valid values: 25, 50, 100, 300, and 500.
	//
	// > This parameter is available only to the ApsaraMQ for RocketMQ Enterprise Platinum Edition instances.
	//
	// example:
	//
	// 50
	TopicCapacity *int32 `json:"TopicCapacity,omitempty" xml:"TopicCapacity,omitempty"`
	// The commodity ID of the instance.
	//
	// example:
	//
	// ons-cn-m7r1r5f****
	SpInstanceId *string `json:"spInstanceId,omitempty" xml:"spInstanceId,omitempty"`
	// The commodity type of the instance.
	//
	// example:
	//
	// 1
	SpInstanceType *int32 `json:"spInstanceType,omitempty" xml:"spInstanceType,omitempty"`
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetCreateTime(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetEndpoints(v *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.Endpoints = v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetIndependentNaming(v bool) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.IndependentNaming = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceId(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceName(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceStatus(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceStatus = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceType(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceType = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetMaxTps(v int64) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.MaxTps = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetReleaseTime(v int64) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.ReleaseTime = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetRemark(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.Remark = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetTopicCapacity(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.TopicCapacity = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetSpInstanceId(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.SpInstanceId = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetSpInstanceType(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.SpInstanceType = &v
	return s
}

type OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints struct {
	// The private HTTP endpoint of the instance.
	//
	// example:
	//
	// http://138015630679****.mqrest.cn-chengdu-internal.aliyuncs.com
	HttpInternalEndpoint *string `json:"HttpInternalEndpoint,omitempty" xml:"HttpInternalEndpoint,omitempty"`
	// The public HTTP endpoint of the instance.
	//
	// example:
	//
	// http://138015630679****.mqrest.cn-chengdu.aliyuncs.com
	HttpInternetEndpoint *string `json:"HttpInternetEndpoint,omitempty" xml:"HttpInternetEndpoint,omitempty"`
	// The public HTTPS endpoint of the instance.
	//
	// example:
	//
	// https://138015630679****.mqrest.cn-chengdu.aliyuncs.com
	HttpInternetSecureEndpoint *string `json:"HttpInternetSecureEndpoint,omitempty" xml:"HttpInternetSecureEndpoint,omitempty"`
	// The private TCP endpoint of the instance.
	//
	// example:
	//
	// http://MQ_INST_138015630679****_BAAy1Hac.cn-chengdu.mq-internal.aliyuncs.com:8080
	TcpEndpoint *string `json:"TcpEndpoint,omitempty" xml:"TcpEndpoint,omitempty"`
	// The public TCP endpoint of the instance.
	//
	// 	- Only instances that are deployed in the China (Chengdu), China (Qingdao), or China (Shenzhen) region can be accessed by using public TCP endpoints.
	//
	// 	- Before you use a public TCP endpoint, make sure that your client SDK meets the following requirements:
	//
	//     	- TCP client SDK for Java: V2.0.0.Final or later For more information, see [Release notes for the SDK for Java](https://help.aliyun.com/document_detail/325569.html).
	//
	//     	- TCP client SDK for C++: V3.0.0 or later For more information, see [Release notes for the SDK for C++](https://help.aliyun.com/document_detail/325570.html).
	//
	// 	- You are charged for Internet traffic when you use a public TCP endpoint. For more information, see [Internet traffic fee](https://help.aliyun.com/document_detail/325571.html).
	//
	// example:
	//
	// http://MQ_INST_138015630679****_BAAy1Hac.mq.cn-chengdu.aliyuncs.com:80
	TcpInternetEndpoint *string `json:"TcpInternetEndpoint,omitempty" xml:"TcpInternetEndpoint,omitempty"`
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetHttpInternalEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.HttpInternalEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetHttpInternetEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.HttpInternetEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetHttpInternetSecureEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.HttpInternetSecureEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetTcpEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.TcpEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetTcpInternetEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.TcpInternetEndpoint = &v
	return s
}

type OnsInstanceBaseInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsInstanceBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsInstanceBaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoResponse) SetHeaders(v map[string]*string) *OnsInstanceBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *OnsInstanceBaseInfoResponse) SetStatusCode(v int32) *OnsInstanceBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsInstanceBaseInfoResponse) SetBody(v *OnsInstanceBaseInfoResponseBody) *OnsInstanceBaseInfoResponse {
	s.Body = v
	return s
}

type OnsInstanceCreateRequest struct {
	// The name of the instance. The name must meet the following rules:
	//
	// 	- The name of the instance must be unique in the region where the instance is deployed.
	//
	// 	- The name must be 3 to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// Test instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// Description
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s OnsInstanceCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceCreateRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceCreateRequest) SetInstanceName(v string) *OnsInstanceCreateRequest {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceCreateRequest) SetRemark(v string) *OnsInstanceCreateRequest {
	s.Remark = &v
	return s
}

type OnsInstanceCreateResponseBody struct {
	// The result returned.
	Data *OnsInstanceCreateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsInstanceCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceCreateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceCreateResponseBody) SetData(v *OnsInstanceCreateResponseBodyData) *OnsInstanceCreateResponseBody {
	s.Data = v
	return s
}

func (s *OnsInstanceCreateResponseBody) SetRequestId(v string) *OnsInstanceCreateResponseBody {
	s.RequestId = &v
	return s
}

type OnsInstanceCreateResponseBodyData struct {
	// The ID of the instance that you created.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The edition of the instance that you created. Valid value:
	//
	// 	- **1**: Standard Edition instances
	//
	// example:
	//
	// 1
	InstanceType *int32 `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s OnsInstanceCreateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceCreateResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsInstanceCreateResponseBodyData) SetInstanceId(v string) *OnsInstanceCreateResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceCreateResponseBodyData) SetInstanceType(v int32) *OnsInstanceCreateResponseBodyData {
	s.InstanceType = &v
	return s
}

type OnsInstanceCreateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsInstanceCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsInstanceCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceCreateResponse) GoString() string {
	return s.String()
}

func (s *OnsInstanceCreateResponse) SetHeaders(v map[string]*string) *OnsInstanceCreateResponse {
	s.Headers = v
	return s
}

func (s *OnsInstanceCreateResponse) SetStatusCode(v int32) *OnsInstanceCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsInstanceCreateResponse) SetBody(v *OnsInstanceCreateResponseBody) *OnsInstanceCreateResponse {
	s.Body = v
	return s
}

type OnsInstanceDeleteRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsInstanceDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceDeleteRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceDeleteRequest) SetInstanceId(v string) *OnsInstanceDeleteRequest {
	s.InstanceId = &v
	return s
}

type OnsInstanceDeleteResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsInstanceDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceDeleteResponseBody) SetRequestId(v string) *OnsInstanceDeleteResponseBody {
	s.RequestId = &v
	return s
}

type OnsInstanceDeleteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsInstanceDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsInstanceDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceDeleteResponse) GoString() string {
	return s.String()
}

func (s *OnsInstanceDeleteResponse) SetHeaders(v map[string]*string) *OnsInstanceDeleteResponse {
	s.Headers = v
	return s
}

func (s *OnsInstanceDeleteResponse) SetStatusCode(v int32) *OnsInstanceDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsInstanceDeleteResponse) SetBody(v *OnsInstanceDeleteResponseBody) *OnsInstanceDeleteResponse {
	s.Body = v
	return s
}

type OnsInstanceInServiceListRequest struct {
	// Specifies whether you want the resource information to be returned.
	//
	// example:
	//
	// true
	NeedResourceInfo *bool `json:"NeedResourceInfo,omitempty" xml:"NeedResourceInfo,omitempty"`
	// The tags that you want to attach to the instance. You can attach up to 20 tags to the instance.
	Tag []*OnsInstanceInServiceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsInstanceInServiceListRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListRequest) SetNeedResourceInfo(v bool) *OnsInstanceInServiceListRequest {
	s.NeedResourceInfo = &v
	return s
}

func (s *OnsInstanceInServiceListRequest) SetTag(v []*OnsInstanceInServiceListRequestTag) *OnsInstanceInServiceListRequest {
	s.Tag = v
	return s
}

type OnsInstanceInServiceListRequestTag struct {
	// The tag key. This parameter is not required. If you configure this parameter, you must also configure **Value**.***	- If you configure Key and Value in a request, this operation queries only the instances that use the specified tags. If you do not configure these parameters in a request, this operation queries all instances that you can access.
	//
	// 	- The value of this parameter cannot be an empty string.
	//
	// 	- The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// CartService
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. This parameter is not required. If you configure this parameter, you must also configure **Key**.***	- If you configure Key and Value in a request, this operation queries only the instances that use the specified tags. If you do not configure these parameters in a request, this operation queries all instances that you can access.
	//
	// 	- The value of this parameter can be an empty string.
	//
	// 	- The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// SericeA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsInstanceInServiceListRequestTag) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListRequestTag) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListRequestTag) SetKey(v string) *OnsInstanceInServiceListRequestTag {
	s.Key = &v
	return s
}

func (s *OnsInstanceInServiceListRequestTag) SetValue(v string) *OnsInstanceInServiceListRequestTag {
	s.Value = &v
	return s
}

type OnsInstanceInServiceListResponseBody struct {
	// The returned information about the queried instances.
	Data *OnsInstanceInServiceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 0598E46F-DB06-40E2-AD7B-C45923EE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsInstanceInServiceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBody) SetData(v *OnsInstanceInServiceListResponseBodyData) *OnsInstanceInServiceListResponseBody {
	s.Data = v
	return s
}

func (s *OnsInstanceInServiceListResponseBody) SetRequestId(v string) *OnsInstanceInServiceListResponseBody {
	s.RequestId = &v
	return s
}

type OnsInstanceInServiceListResponseBodyData struct {
	InstanceVO []*OnsInstanceInServiceListResponseBodyDataInstanceVO `json:"InstanceVO,omitempty" xml:"InstanceVO,omitempty" type:"Repeated"`
}

func (s OnsInstanceInServiceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBodyData) SetInstanceVO(v []*OnsInstanceInServiceListResponseBodyDataInstanceVO) *OnsInstanceInServiceListResponseBodyData {
	s.InstanceVO = v
	return s
}

type OnsInstanceInServiceListResponseBodyDataInstanceVO struct {
	// The time when the instance was created. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1640847284000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of consumer groups.
	//
	// example:
	//
	// 3
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// Indicates whether a namespace is used for the instance. Valid values:
	//
	// 	- **true**: A separate namespace is used for the instance. The identifier of each resource in the instance must be unique within the instance. However, the identifier of a resource in the instance can be the same as the identifier of a resource in another instance that uses a different namespace.
	//
	// 	- **false**: A separate namespace is not used for the instance. The name of each resource must be globally unique within the instance and across all instances.
	//
	// example:
	//
	// true
	IndependentNaming *bool `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// The name must be 3 to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// test1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance status. Valid values:
	//
	// 	- **0**: The instance is being deployed. This value is valid only for Enterprise Platinum Edition instances.
	//
	// 	- **2**: The instance has overdue payments. This value is valid only for Standard Edition instances.
	//
	// 	- **5**: The instance is running. This value is valid only for Standard Edition instances and Enterprise Platinum Edition instances.
	//
	// 	- **7**: The instance is being upgraded and is running. This value is valid only for Enterprise Platinum Edition instances.
	//
	// example:
	//
	// 5
	InstanceStatus *int32 `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance type. Valid values:
	//
	// 	- **1**: Standard Edition
	//
	// 	- **2**: Enterprise Platinum Edition
	//
	// For information about the instance editions and the differences between the editions, see [Instance editions](https://help.aliyun.com/document_detail/185261.html).
	//
	// example:
	//
	// 2
	InstanceType *int32 `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The time when the instance expires. If the instance is of Enterprise Platinum Edition, this parameter is returned.
	//
	// example:
	//
	// 1551024000000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The tags that are attached to the instance.
	Tags *OnsInstanceInServiceListResponseBodyDataInstanceVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The number of topics.
	//
	// example:
	//
	// 1
	TopicCount *int32 `json:"TopicCount,omitempty" xml:"TopicCount,omitempty"`
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVO) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVO) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetCreateTime(v int64) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.CreateTime = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetGroupCount(v int32) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.GroupCount = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetIndependentNaming(v bool) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.IndependentNaming = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetInstanceId(v string) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetInstanceName(v string) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetInstanceStatus(v int32) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.InstanceStatus = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetInstanceType(v int32) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.InstanceType = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetReleaseTime(v int64) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.ReleaseTime = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetTags(v *OnsInstanceInServiceListResponseBodyDataInstanceVOTags) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.Tags = v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVO) SetTopicCount(v int32) *OnsInstanceInServiceListResponseBodyDataInstanceVO {
	s.TopicCount = &v
	return s
}

type OnsInstanceInServiceListResponseBodyDataInstanceVOTags struct {
	Tag []*OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVOTags) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVOTags) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTags) SetTag(v []*OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) *OnsInstanceInServiceListResponseBodyDataInstanceVOTags {
	s.Tag = v
	return s
}

type OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// CartService
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// ServiceA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) SetKey(v string) *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag {
	s.Key = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag) SetValue(v string) *OnsInstanceInServiceListResponseBodyDataInstanceVOTagsTag {
	s.Value = &v
	return s
}

type OnsInstanceInServiceListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsInstanceInServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsInstanceInServiceListResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListResponse) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponse) SetHeaders(v map[string]*string) *OnsInstanceInServiceListResponse {
	s.Headers = v
	return s
}

func (s *OnsInstanceInServiceListResponse) SetStatusCode(v int32) *OnsInstanceInServiceListResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsInstanceInServiceListResponse) SetBody(v *OnsInstanceInServiceListResponseBody) *OnsInstanceInServiceListResponse {
	s.Body = v
	return s
}

type OnsInstanceUpdateRequest struct {
	// The ID of the instance whose name or description you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new name of the instance. The name must meet the following rules:
	//
	// 	- The name of the instance must be unique in the region where the instance is deployed.
	//
	// 	- The name must be 3 to 64 characters in length and can contain letters, digits, hyphens (-), underscores (_), and Chinese characters.
	//
	// 	- If you do not configure this parameter, the instance name remains unchanged.
	//
	// example:
	//
	// Updatedname
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The updated description of the instance. If you do not configure this parameter, the instance description remains unchanged.
	//
	// example:
	//
	// Updateddescription
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s OnsInstanceUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceUpdateRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceUpdateRequest) SetInstanceId(v string) *OnsInstanceUpdateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceUpdateRequest) SetInstanceName(v string) *OnsInstanceUpdateRequest {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceUpdateRequest) SetRemark(v string) *OnsInstanceUpdateRequest {
	s.Remark = &v
	return s
}

type OnsInstanceUpdateResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsInstanceUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceUpdateResponseBody) SetRequestId(v string) *OnsInstanceUpdateResponseBody {
	s.RequestId = &v
	return s
}

type OnsInstanceUpdateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsInstanceUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsInstanceUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceUpdateResponse) GoString() string {
	return s.String()
}

func (s *OnsInstanceUpdateResponse) SetHeaders(v map[string]*string) *OnsInstanceUpdateResponse {
	s.Headers = v
	return s
}

func (s *OnsInstanceUpdateResponse) SetStatusCode(v int32) *OnsInstanceUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsInstanceUpdateResponse) SetBody(v *OnsInstanceUpdateResponseBody) *OnsInstanceUpdateResponse {
	s.Body = v
	return s
}

type OnsMessageDetailRequest struct {
	// The ID of the ApsaraMQ for RocketMQ Instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_184681981******_BXig0x6A
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C0******
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The name of the topic in which the message you want to query is stored.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageDetailRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageDetailRequest) SetInstanceId(v string) *OnsMessageDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageDetailRequest) SetMsgId(v string) *OnsMessageDetailRequest {
	s.MsgId = &v
	return s
}

func (s *OnsMessageDetailRequest) SetTopic(v string) *OnsMessageDetailRequest {
	s.Topic = &v
	return s
}

type OnsMessageDetailResponseBody struct {
	// The data returned.
	Data *OnsMessageDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// EAE5BE23-37A1-4354-94D6-E44AE17E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessageDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageDetailResponseBody) SetData(v *OnsMessageDetailResponseBodyData) *OnsMessageDetailResponseBody {
	s.Data = v
	return s
}

func (s *OnsMessageDetailResponseBody) SetRequestId(v string) *OnsMessageDetailResponseBody {
	s.RequestId = &v
	return s
}

type OnsMessageDetailResponseBodyData struct {
	// The string that is obtained after the message body is encrypted by using the Base 64 algorithm.
	//
	// example:
	//
	// aGVsbG8=
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The cyclic redundancy check (CRC) value of the message body.
	//
	// example:
	//
	// 907060870
	BodyCRC *int32 `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	// The information about the message body.
	//
	// example:
	//
	// hello
	BodyStr *string `json:"BodyStr,omitempty" xml:"BodyStr,omitempty"`
	// The producer instance that generated the message.
	//
	// example:
	//
	// ``42.120.**.**``:64646
	BornHost *string `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	// The timestamp that indicates the point in time when the message was generated. Unit: milliseconds.
	//
	// example:
	//
	// 1570761026630
	BornTimestamp *int64 `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ Instance.
	//
	// example:
	//
	// MQ_INST_184681981******_BXig0x6A
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C05F2*****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The attributes of the message.
	PropertyList []*OnsMessageDetailResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// The number of retries that ApsaraMQ for RocketMQ performed to send the message to consumers.
	//
	// example:
	//
	// 0
	ReconsumeTimes *int32 `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	// The ApsaraMQ for RocketMQ broker that stores the message.
	//
	// example:
	//
	// 11.220.***.***:10911
	StoreHost *string `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	// The size of the message. Unit: KB.
	//
	// example:
	//
	// 2
	StoreSize *int32 `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	// The timestamp that indicates the point in time when the ApsaraMQ for RocketMQ broker stored the message. Unit: milliseconds.
	//
	// example:
	//
	// 1570761026708
	StoreTimestamp *int64 `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageDetailResponseBodyData) SetBody(v string) *OnsMessageDetailResponseBodyData {
	s.Body = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetBodyCRC(v int32) *OnsMessageDetailResponseBodyData {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetBodyStr(v string) *OnsMessageDetailResponseBodyData {
	s.BodyStr = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetBornHost(v string) *OnsMessageDetailResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetBornTimestamp(v int64) *OnsMessageDetailResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetInstanceId(v string) *OnsMessageDetailResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetMsgId(v string) *OnsMessageDetailResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetPropertyList(v []*OnsMessageDetailResponseBodyDataPropertyList) *OnsMessageDetailResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetReconsumeTimes(v int32) *OnsMessageDetailResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetStoreHost(v string) *OnsMessageDetailResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetStoreSize(v int32) *OnsMessageDetailResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetStoreTimestamp(v int64) *OnsMessageDetailResponseBodyData {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetTopic(v string) *OnsMessageDetailResponseBodyData {
	s.Topic = &v
	return s
}

type OnsMessageDetailResponseBodyDataPropertyList struct {
	// The name of the attribute. Valid values:
	//
	// 	- **TRACE_ON**: indicates whether the trace of the message exists.
	//
	// 	- **MSG_REGION**: The region ID of the instance to which the topic belongs.
	//
	// 	- **__MESSAGE_DECODED_TIME**: The time when the message was decoded.
	//
	// 	- **__BORNHOST**: The IP address of the producer client.
	//
	// 	- **UNIQ_KEY**: The ID of the message.
	//
	// For information about the terms that are used in ApsaraMQ for RocketMQ, see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// example:
	//
	// cn-hangzhou
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// MSG_REGION
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsMessageDetailResponseBodyDataPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageDetailResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessageDetailResponseBodyDataPropertyList) SetName(v string) *OnsMessageDetailResponseBodyDataPropertyList {
	s.Name = &v
	return s
}

func (s *OnsMessageDetailResponseBodyDataPropertyList) SetValue(v string) *OnsMessageDetailResponseBodyDataPropertyList {
	s.Value = &v
	return s
}

type OnsMessageDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessageDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageDetailResponse) GoString() string {
	return s.String()
}

func (s *OnsMessageDetailResponse) SetHeaders(v map[string]*string) *OnsMessageDetailResponse {
	s.Headers = v
	return s
}

func (s *OnsMessageDetailResponse) SetStatusCode(v int32) *OnsMessageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessageDetailResponse) SetBody(v *OnsMessageDetailResponseBody) *OnsMessageDetailResponse {
	s.Body = v
	return s
}

type OnsMessageGetByKeyRequest struct {
	// The ID of the instance to which the messages that you want to query belong.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The key of the messages that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// messageKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The topic that contains the messages that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageGetByKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyRequest) SetInstanceId(v string) *OnsMessageGetByKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageGetByKeyRequest) SetKey(v string) *OnsMessageGetByKeyRequest {
	s.Key = &v
	return s
}

func (s *OnsMessageGetByKeyRequest) SetTopic(v string) *OnsMessageGetByKeyRequest {
	s.Topic = &v
	return s
}

type OnsMessageGetByKeyResponseBody struct {
	// The list of returned results.
	Data *OnsMessageGetByKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessageGetByKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBody) SetData(v *OnsMessageGetByKeyResponseBodyData) *OnsMessageGetByKeyResponseBody {
	s.Data = v
	return s
}

func (s *OnsMessageGetByKeyResponseBody) SetRequestId(v string) *OnsMessageGetByKeyResponseBody {
	s.RequestId = &v
	return s
}

type OnsMessageGetByKeyResponseBodyData struct {
	OnsRestMessageDo []*OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo `json:"OnsRestMessageDo,omitempty" xml:"OnsRestMessageDo,omitempty" type:"Repeated"`
}

func (s OnsMessageGetByKeyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBodyData) SetOnsRestMessageDo(v []*OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) *OnsMessageGetByKeyResponseBodyData {
	s.OnsRestMessageDo = v
	return s
}

type OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo struct {
	// The cyclic redundancy check (CRC) value of the message body.
	//
	// example:
	//
	// 914112295
	BodyCRC *int32 `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	// The producer client that generated the message.
	//
	// example:
	//
	// 42.120.***.***:59270
	BornHost *string `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	// The timestamp that indicates when the message was produced.
	//
	// example:
	//
	// 1570760999721
	BornTimestamp *int64 `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C05F2*****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The attributes of the message.
	PropertyList *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	// The number of retries that were performed to send the message to consumers.
	//
	// example:
	//
	// 1
	ReconsumeTimes *int32 `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	// The ApsaraMQ for RocketMQ broker that stores the message.
	//
	// example:
	//
	// 11.193.***.***:10911
	StoreHost *string `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	// The size of the message.
	//
	// example:
	//
	// 406
	StoreSize *int32 `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	// The timestamp that indicates when the ApsaraMQ for RocketMQ broker stored the message.
	//
	// example:
	//
	// 1570760999811
	StoreTimestamp *int64 `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetBodyCRC(v int32) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetBornHost(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.BornHost = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetBornTimestamp(v int64) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetInstanceId(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetMsgId(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.MsgId = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetPropertyList(v *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.PropertyList = v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetReconsumeTimes(v int32) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetStoreHost(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.StoreHost = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetStoreSize(v int32) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.StoreSize = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetStoreTimestamp(v int64) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo) SetTopic(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDo {
	s.Topic = &v
	return s
}

type OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList struct {
	MessageProperty []*OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty `json:"MessageProperty,omitempty" xml:"MessageProperty,omitempty" type:"Repeated"`
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList) SetMessageProperty(v []*OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList {
	s.MessageProperty = v
	return s
}

type OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty struct {
	// The name of the attribute. Valid values:
	//
	// -  **TRACE_ON**: indicates whether the message trace exists.
	//
	// -   **KEYS**: indicates the key of the message.
	//
	// - **TAGS**: indicates the tag that is attached to the message.
	//
	// - **INSTANCE_ID**: indicates the ID of the instance that contains the message.
	//
	// For information about the terms that are used in ApsaraMQ for RocketMQ, see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// example:
	//
	// TAGS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// TagA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) SetName(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty {
	s.Name = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty) SetValue(v string) *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyListMessageProperty {
	s.Value = &v
	return s
}

type OnsMessageGetByKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessageGetByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessageGetByKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyResponse) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponse) SetHeaders(v map[string]*string) *OnsMessageGetByKeyResponse {
	s.Headers = v
	return s
}

func (s *OnsMessageGetByKeyResponse) SetStatusCode(v int32) *OnsMessageGetByKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessageGetByKeyResponse) SetBody(v *OnsMessageGetByKeyResponseBody) *OnsMessageGetByKeyResponse {
	s.Body = v
	return s
}

type OnsMessageGetByMsgIdRequest struct {
	// The ID of the instance to which the message you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C05F2*****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The topic that contains the message you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageGetByMsgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByMsgIdRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdRequest) SetInstanceId(v string) *OnsMessageGetByMsgIdRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageGetByMsgIdRequest) SetMsgId(v string) *OnsMessageGetByMsgIdRequest {
	s.MsgId = &v
	return s
}

func (s *OnsMessageGetByMsgIdRequest) SetTopic(v string) *OnsMessageGetByMsgIdRequest {
	s.Topic = &v
	return s
}

type OnsMessageGetByMsgIdResponseBody struct {
	// The data returned.
	Data *OnsMessageGetByMsgIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessageGetByMsgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBody) SetData(v *OnsMessageGetByMsgIdResponseBodyData) *OnsMessageGetByMsgIdResponseBody {
	s.Data = v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBody) SetRequestId(v string) *OnsMessageGetByMsgIdResponseBody {
	s.RequestId = &v
	return s
}

type OnsMessageGetByMsgIdResponseBodyData struct {
	// The cyclic redundancy check (CRC) value of the message body.
	//
	// example:
	//
	// 914112295
	BodyCRC *int32 `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	// The producer instance that generated the message.
	//
	// example:
	//
	// ``42.120.**.**``:64646
	BornHost *string `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	// The timestamp that indicates when the message was produced.
	//
	// example:
	//
	// 1570761026630
	BornTimestamp *int64 `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C0C8460BA2
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The attributes of the message.
	PropertyList *OnsMessageGetByMsgIdResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	// The number of retries that were performed to send the message to consumers.
	//
	// example:
	//
	// 1
	ReconsumeTimes *int32 `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	// The ApsaraMQ for RocketMQ broker that stores the message.
	//
	// example:
	//
	// 11.220.***.***:10911
	StoreHost *string `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	// The size of the message.
	//
	// example:
	//
	// 407
	StoreSize *int32 `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	// The timestamp that indicates when the ApsaraMQ for RocketMQ broker stored the message.
	//
	// example:
	//
	// 1570761026708
	StoreTimestamp *int64 `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageGetByMsgIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetBodyCRC(v int32) *OnsMessageGetByMsgIdResponseBodyData {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetBornHost(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetBornTimestamp(v int64) *OnsMessageGetByMsgIdResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetInstanceId(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetMsgId(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetPropertyList(v *OnsMessageGetByMsgIdResponseBodyDataPropertyList) *OnsMessageGetByMsgIdResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetReconsumeTimes(v int32) *OnsMessageGetByMsgIdResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetStoreHost(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetStoreSize(v int32) *OnsMessageGetByMsgIdResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetStoreTimestamp(v int64) *OnsMessageGetByMsgIdResponseBodyData {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetTopic(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.Topic = &v
	return s
}

type OnsMessageGetByMsgIdResponseBodyDataPropertyList struct {
	MessageProperty []*OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty `json:"MessageProperty,omitempty" xml:"MessageProperty,omitempty" type:"Repeated"`
}

func (s OnsMessageGetByMsgIdResponseBodyDataPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyList) SetMessageProperty(v []*OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) *OnsMessageGetByMsgIdResponseBodyDataPropertyList {
	s.MessageProperty = v
	return s
}

type OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty struct {
	// The name of the attribute. Valid values:
	//
	// 	- **TRACE_ON**: indicates whether a trace of the message exists.
	//
	// 	- **KEYS**: indicates the key of the message.
	//
	// 	- **TAGS**: indicates the tag that is attached to the message.
	//
	// 	- **INSTANCE_ID**: indicates the ID of the instance which contains the message.
	//
	// For information about the terms that are used in Message Queue for Apache RocketMQ, see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// example:
	//
	// TAGS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// TagA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) SetName(v string) *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty {
	s.Name = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) SetValue(v string) *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty {
	s.Value = &v
	return s
}

type OnsMessageGetByMsgIdResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessageGetByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessageGetByMsgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponse) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponse) SetHeaders(v map[string]*string) *OnsMessageGetByMsgIdResponse {
	s.Headers = v
	return s
}

func (s *OnsMessageGetByMsgIdResponse) SetStatusCode(v int32) *OnsMessageGetByMsgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponse) SetBody(v *OnsMessageGetByMsgIdResponseBody) *OnsMessageGetByMsgIdResponse {
	s.Body = v
	return s
}

type OnsMessagePageQueryByTopicRequest struct {
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds. If you specify a valid value for the **TaskId*	- parameter in the request, this parameter does not take effect. The system uses the value of the BeginTime parameter that you specified in the request when you created the specified query task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570723200000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The number of the page to return. Pages start from page 1. Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds. If you specify a valid value for the **TaskId*	- parameter in the request, this parameter does not take effect. The system uses the value of the EndTime parameter that you specified in the request when you created the specified query task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570809600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance to which the message you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. Valid values: 5 to 50. Default value: 20. If you specify a valid value for the **TaskId*	- parameter in the request, this parameter does not take effect. The system uses the value of the PageSize parameter that you specified in the request for creating the query task.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the query task. The first time you call this operation to query the information about messages in a specified topic within a specified time range, this parameter is not required. This parameter is required in subsequent queries for messages on a specified page. You can obtain the task ID from the returned result of the first query.
	//
	// example:
	//
	// 0BC1310300002A9F000021E4D7A48346
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the topic whose messages you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessagePageQueryByTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicRequest) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicRequest) SetBeginTime(v int64) *OnsMessagePageQueryByTopicRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetCurrentPage(v int32) *OnsMessagePageQueryByTopicRequest {
	s.CurrentPage = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetEndTime(v int64) *OnsMessagePageQueryByTopicRequest {
	s.EndTime = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetInstanceId(v string) *OnsMessagePageQueryByTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetPageSize(v int32) *OnsMessagePageQueryByTopicRequest {
	s.PageSize = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetTaskId(v string) *OnsMessagePageQueryByTopicRequest {
	s.TaskId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetTopic(v string) *OnsMessagePageQueryByTopicRequest {
	s.Topic = &v
	return s
}

type OnsMessagePageQueryByTopicResponseBody struct {
	// The information about the message that is queried.
	MsgFoundDo *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo `json:"MsgFoundDo,omitempty" xml:"MsgFoundDo,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 5DC2A47E-2B31-4722-96C8-FA59C9*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBody) SetMsgFoundDo(v *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) *OnsMessagePageQueryByTopicResponseBody {
	s.MsgFoundDo = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBody) SetRequestId(v string) *OnsMessagePageQueryByTopicResponseBody {
	s.RequestId = &v
	return s
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 400
	MaxPageCount *int64 `json:"MaxPageCount,omitempty" xml:"MaxPageCount,omitempty"`
	// The information about messages on the returned page. The information that is contained in this parameter is the same as the information that is returned by the [OnsMessageGetByMsgId](https://help.aliyun.com/document_detail/29607.html) operation.
	MsgFoundList *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList `json:"MsgFoundList,omitempty" xml:"MsgFoundList,omitempty" type:"Struct"`
	// The ID of the query task. The first time you call this operation to query the messages that are sent to a specified consumer group within a specified time range, this parameter is returned. You can use the task ID to query the details of messages on other returned pages.
	//
	// example:
	//
	// 0BC1310300002A9F000021E4D7A48346
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetCurrentPage(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.CurrentPage = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetMaxPageCount(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.MaxPageCount = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetMsgFoundList(v *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.MsgFoundList = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetTaskId(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.TaskId = &v
	return s
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList struct {
	OnsRestMessageDo []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo `json:"OnsRestMessageDo,omitempty" xml:"OnsRestMessageDo,omitempty" type:"Repeated"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetOnsRestMessageDo(v []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.OnsRestMessageDo = v
	return s
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo struct {
	// The cyclic redundancy check (CRC) value of the message body.
	//
	// example:
	//
	// 914112295
	BodyCRC *int32 `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	// The producer client that generated the message.
	//
	// example:
	//
	// 42.120.***.***:59270
	BornHost *string `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	// The time when the message was generated. The value is a UNIX timestamp that represents the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1570760999721
	BornTimestamp *int64 `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C05F2*****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The attributes of the message.
	PropertyList *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	// The number of retries that were performed to send the message to consumers.
	//
	// example:
	//
	// 1
	ReconsumeTimes *int32 `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	// The ApsaraMQ for RocketMQ broker that stores the message.
	//
	// example:
	//
	// 11.193.***.***:10911
	StoreHost *string `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	// The size of the message. Unit: KB.
	//
	// example:
	//
	// 406
	StoreSize *int32 `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	// The time when the ApsaraMQ for RocketMQ broker stored the message. The value is a UNIX timestamp that represents the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1570760999811
	StoreTimestamp *int64 `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBodyCRC(v int32) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBornHost(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BornHost = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetBornTimestamp(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetInstanceId(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.InstanceId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetMsgId(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.MsgId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetPropertyList(v *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.PropertyList = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetReconsumeTimes(v int32) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreHost(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreHost = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreSize(v int32) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreSize = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetStoreTimestamp(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo) SetTopic(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDo {
	s.Topic = &v
	return s
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList struct {
	MessageProperty []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty `json:"MessageProperty,omitempty" xml:"MessageProperty,omitempty" type:"Repeated"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList) SetMessageProperty(v []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList {
	s.MessageProperty = v
	return s
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty struct {
	// The name of the attribute. Valid values:
	//
	// 	- **TRACE_ON**: indicates whether a trace of the message exists.
	//
	// 	- **KEYS**: indicates the key of the message.
	//
	// 	- **TAGS**: indicates the tag of the message.
	//
	// 	- **INSTANCE_ID**: indicates the ID of the instance that contains the message.
	//
	// For information about the terms that are used in ApsaraMQ for RocketMQ see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// example:
	//
	// TAGS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// TagA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) SetName(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty {
	s.Name = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty) SetValue(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyListMessageProperty {
	s.Value = &v
	return s
}

type OnsMessagePageQueryByTopicResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessagePageQueryByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponse) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponse) SetHeaders(v map[string]*string) *OnsMessagePageQueryByTopicResponse {
	s.Headers = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponse) SetStatusCode(v int32) *OnsMessagePageQueryByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponse) SetBody(v *OnsMessagePageQueryByTopicResponseBody) *OnsMessagePageQueryByTopicResponse {
	s.Body = v
	return s
}

type OnsMessagePushRequest struct {
	// The ID of the consumer client. You can call the [OnsConsumerGetConnection](https://help.aliyun.com/document_detail/29598.html) operation to query client IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30.5.121.**@24813#-1999745829#-1737591554#453111174894656
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the consumer group. For information about what a consumer group is, see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance to which the specified consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0BC1669963053CF68F733BB70396****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The topic to which the message is pushed.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessagePushRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePushRequest) GoString() string {
	return s.String()
}

func (s *OnsMessagePushRequest) SetClientId(v string) *OnsMessagePushRequest {
	s.ClientId = &v
	return s
}

func (s *OnsMessagePushRequest) SetGroupId(v string) *OnsMessagePushRequest {
	s.GroupId = &v
	return s
}

func (s *OnsMessagePushRequest) SetInstanceId(v string) *OnsMessagePushRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessagePushRequest) SetMsgId(v string) *OnsMessagePushRequest {
	s.MsgId = &v
	return s
}

func (s *OnsMessagePushRequest) SetTopic(v string) *OnsMessagePushRequest {
	s.Topic = &v
	return s
}

type OnsMessagePushResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// B8EDC90D-F726-4B9E-8BEF-F0DD25EC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessagePushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePushResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessagePushResponseBody) SetRequestId(v string) *OnsMessagePushResponseBody {
	s.RequestId = &v
	return s
}

type OnsMessagePushResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessagePushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessagePushResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePushResponse) GoString() string {
	return s.String()
}

func (s *OnsMessagePushResponse) SetHeaders(v map[string]*string) *OnsMessagePushResponse {
	s.Headers = v
	return s
}

func (s *OnsMessagePushResponse) SetStatusCode(v int32) *OnsMessagePushResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessagePushResponse) SetBody(v *OnsMessagePushResponseBody) *OnsMessagePushResponse {
	s.Body = v
	return s
}

type OnsMessageTraceRequest struct {
	// The ID of the instance to which the message you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1E05791C117818B4AAC23B1BB0CE****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The topic to which the message belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageTraceRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceRequest) SetInstanceId(v string) *OnsMessageTraceRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageTraceRequest) SetMsgId(v string) *OnsMessageTraceRequest {
	s.MsgId = &v
	return s
}

func (s *OnsMessageTraceRequest) SetTopic(v string) *OnsMessageTraceRequest {
	s.Topic = &v
	return s
}

type OnsMessageTraceResponseBody struct {
	// The information about the message that is queried.
	Data *OnsMessageTraceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// EAE5BE23-37A1-4354-94D6-E44AE17E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessageTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageTraceResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceResponseBody) SetData(v *OnsMessageTraceResponseBodyData) *OnsMessageTraceResponseBody {
	s.Data = v
	return s
}

func (s *OnsMessageTraceResponseBody) SetRequestId(v string) *OnsMessageTraceResponseBody {
	s.RequestId = &v
	return s
}

type OnsMessageTraceResponseBodyData struct {
	MessageTrack []*OnsMessageTraceResponseBodyDataMessageTrack `json:"MessageTrack,omitempty" xml:"MessageTrack,omitempty" type:"Repeated"`
}

func (s OnsMessageTraceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageTraceResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceResponseBodyData) SetMessageTrack(v []*OnsMessageTraceResponseBodyDataMessageTrack) *OnsMessageTraceResponseBodyData {
	s.MessageTrack = v
	return s
}

type OnsMessageTraceResponseBodyDataMessageTrack struct {
	// The ID of the consumer group that subscribes to the topic.
	//
	// example:
	//
	// GID_test_group_id
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the instance to which the message you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the message. Valid values:
	//
	// 	- **CONSUMED**: The message is consumed.
	//
	// 	- **CONSUMED_BUT_FILTERED**: No consumer group subscribes to the message. The message is filtered out and not consumed.
	//
	// 	- **NOT_CONSUME_YET**: The message is not consumed.
	//
	// 	- **NOT_ONLINE**: The consumer group is offline.
	//
	// 	- **UNKNOWN**: The message is not consumed due to unknown reasons.
	//
	// example:
	//
	// CONSUMED
	TrackType *string `json:"TrackType,omitempty" xml:"TrackType,omitempty"`
}

func (s OnsMessageTraceResponseBodyDataMessageTrack) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageTraceResponseBodyDataMessageTrack) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceResponseBodyDataMessageTrack) SetConsumerGroup(v string) *OnsMessageTraceResponseBodyDataMessageTrack {
	s.ConsumerGroup = &v
	return s
}

func (s *OnsMessageTraceResponseBodyDataMessageTrack) SetInstanceId(v string) *OnsMessageTraceResponseBodyDataMessageTrack {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageTraceResponseBodyDataMessageTrack) SetTrackType(v string) *OnsMessageTraceResponseBodyDataMessageTrack {
	s.TrackType = &v
	return s
}

type OnsMessageTraceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessageTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessageTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageTraceResponse) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceResponse) SetHeaders(v map[string]*string) *OnsMessageTraceResponse {
	s.Headers = v
	return s
}

func (s *OnsMessageTraceResponse) SetStatusCode(v int32) *OnsMessageTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessageTraceResponse) SetBody(v *OnsMessageTraceResponseBody) *OnsMessageTraceResponse {
	s.Body = v
	return s
}

type OnsRegionListResponseBody struct {
	// The returned data.
	Data *OnsRegionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 72D14A84-45E5-4E01-A6DB-F63C4721****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsRegionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsRegionListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponseBody) SetData(v *OnsRegionListResponseBodyData) *OnsRegionListResponseBody {
	s.Data = v
	return s
}

func (s *OnsRegionListResponseBody) SetRequestId(v string) *OnsRegionListResponseBody {
	s.RequestId = &v
	return s
}

type OnsRegionListResponseBodyData struct {
	RegionDo []*OnsRegionListResponseBodyDataRegionDo `json:"RegionDo,omitempty" xml:"RegionDo,omitempty" type:"Repeated"`
}

func (s OnsRegionListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsRegionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponseBodyData) SetRegionDo(v []*OnsRegionListResponseBodyDataRegionDo) *OnsRegionListResponseBodyData {
	s.RegionDo = v
	return s
}

type OnsRegionListResponseBodyDataRegionDo struct {
	// The channel name.
	//
	// example:
	//
	// ALIYUN
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1411623866000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	OnsRegionId *string `json:"OnsRegionId,omitempty" xml:"OnsRegionId,omitempty"`
	// The region name.
	//
	// example:
	//
	// China (Hangzhou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The time when the instance was updated.
	//
	// example:
	//
	// 1411623866000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s OnsRegionListResponseBodyDataRegionDo) String() string {
	return tea.Prettify(s)
}

func (s OnsRegionListResponseBodyDataRegionDo) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetChannelName(v string) *OnsRegionListResponseBodyDataRegionDo {
	s.ChannelName = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetCreateTime(v int64) *OnsRegionListResponseBodyDataRegionDo {
	s.CreateTime = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetId(v int64) *OnsRegionListResponseBodyDataRegionDo {
	s.Id = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetOnsRegionId(v string) *OnsRegionListResponseBodyDataRegionDo {
	s.OnsRegionId = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetRegionName(v string) *OnsRegionListResponseBodyDataRegionDo {
	s.RegionName = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetUpdateTime(v int64) *OnsRegionListResponseBodyDataRegionDo {
	s.UpdateTime = &v
	return s
}

type OnsRegionListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsRegionListResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsRegionListResponse) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponse) SetHeaders(v map[string]*string) *OnsRegionListResponse {
	s.Headers = v
	return s
}

func (s *OnsRegionListResponse) SetStatusCode(v int32) *OnsRegionListResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsRegionListResponse) SetBody(v *OnsRegionListResponseBody) *OnsRegionListResponse {
	s.Body = v
	return s
}

type OnsTopicCreateRequest struct {
	// The ID of the instance in which you want to create the topic.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of messages that you want to publish to the topic. Valid values:
	//
	// 	- **0**: normal messages
	//
	// 	- **1**: partitionally ordered messages
	//
	// 	- **2**: globally ordered messages
	//
	// 	- **4**: transactional messages
	//
	// 	- **5**: scheduled or delayed messages
	//
	// For more information about message types, see [Message types](https://help.aliyun.com/document_detail/155952.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	MessageType *int32 `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The description of the topic that you want to create.
	//
	// example:
	//
	// Test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of the topic that you want to create. The name must meet the following rules:
	//
	// 	- The name must be 3 to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The topic name cannot start with CID or GID because CID and GID are reserved prefixes for group IDs.
	//
	// 	- If the ApsaraMQ for RocketMQ instance in which you want to create the topic uses a namespace, the topic name must be unique in the instance. The topic name cannot be the same as an existing topic name or a group ID in the instance. The topic name can be the same as a topic name or a group ID in another instance that uses a different namespace. For example, if Instance A and Instance B use different namespaces, a topic name in Instance A can be the same as a topic name or a group ID in Instance B.
	//
	// 	- If the instance in which you want to create the topic does not use a namespace, the topic name must be globally unique across instances and regions. The topic name cannot be the same as an existing topic name or group ID in all ApsaraMQ for RocketMQ instances that belong to your Alibaba Cloud account.
	//
	// > To check whether an instance uses a namespace, log on to the ApsaraMQ for RocketMQ console, go to the **Instance Details*	- page, and then view the value of the Namespace field in the **Basic Information*	- section.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicCreateRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicCreateRequest) SetInstanceId(v string) *OnsTopicCreateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicCreateRequest) SetMessageType(v int32) *OnsTopicCreateRequest {
	s.MessageType = &v
	return s
}

func (s *OnsTopicCreateRequest) SetRemark(v string) *OnsTopicCreateRequest {
	s.Remark = &v
	return s
}

func (s *OnsTopicCreateRequest) SetTopic(v string) *OnsTopicCreateRequest {
	s.Topic = &v
	return s
}

type OnsTopicCreateResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// B6949B58-223E-4B75-B4FE-7797C15E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicCreateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicCreateResponseBody) SetRequestId(v string) *OnsTopicCreateResponseBody {
	s.RequestId = &v
	return s
}

type OnsTopicCreateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicCreateResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicCreateResponse) SetHeaders(v map[string]*string) *OnsTopicCreateResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicCreateResponse) SetStatusCode(v int32) *OnsTopicCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicCreateResponse) SetBody(v *OnsTopicCreateResponseBody) *OnsTopicCreateResponse {
	s.Body = v
	return s
}

type OnsTopicDeleteRequest struct {
	// The ID of the instance to which the topic you want to delete belongs.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the topic that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicDeleteRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicDeleteRequest) SetInstanceId(v string) *OnsTopicDeleteRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicDeleteRequest) SetTopic(v string) *OnsTopicDeleteRequest {
	s.Topic = &v
	return s
}

type OnsTopicDeleteResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 4189D4A6-231A-4028-8D89-F66A76C1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicDeleteResponseBody) SetRequestId(v string) *OnsTopicDeleteResponseBody {
	s.RequestId = &v
	return s
}

type OnsTopicDeleteResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicDeleteResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicDeleteResponse) SetHeaders(v map[string]*string) *OnsTopicDeleteResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicDeleteResponse) SetStatusCode(v int32) *OnsTopicDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicDeleteResponse) SetBody(v *OnsTopicDeleteResponseBody) *OnsTopicDeleteResponse {
	s.Body = v
	return s
}

type OnsTopicListRequest struct {
	// The ID of the instance that contains the topics you want to query.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of tags that are attached to the topic. A maximum of 20 tags can be included in the list.
	Tag []*OnsTopicListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic that you want to query. This parameter is required if you want to query a specific topic. If you do not include this parameter in a request, all topics that you can access are queried.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The user ID of the topic owner. Set this parameter to an Alibaba Cloud account ID.
	//
	// example:
	//
	// 138015630679****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s OnsTopicListRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicListRequest) SetInstanceId(v string) *OnsTopicListRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicListRequest) SetTag(v []*OnsTopicListRequestTag) *OnsTopicListRequest {
	s.Tag = v
	return s
}

func (s *OnsTopicListRequest) SetTopic(v string) *OnsTopicListRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicListRequest) SetUserId(v string) *OnsTopicListRequest {
	s.UserId = &v
	return s
}

type OnsTopicListRequestTag struct {
	// The key of the tag that is attached to the topics you want to query. This parameter is not required. If you configure this parameter, you must also configure the **Value*	- parameter.***	- If you include the Key and Value parameters in a request, this operation queries only the topics that use the specified tag. If you do not include these parameters in a request, this operation queries all topics that you can access.
	//
	// 	- The value of this parameter cannot be an empty string.
	//
	// 	- A tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// CartService
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is attached to the topics you want to query. This parameter is not required. If you configure this parameter, you must also configure the **Key*	- parameter.***	- If you include the Key and Value parameters in a request, this operation queries only the topics that use the specified tag. If you do not include these parameters in a request, this operation queries all topics that you can access.
	//
	// 	- The value of this parameter can be an empty string.
	//
	// 	- A tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ServiceA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsTopicListRequestTag) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListRequestTag) GoString() string {
	return s.String()
}

func (s *OnsTopicListRequestTag) SetKey(v string) *OnsTopicListRequestTag {
	s.Key = &v
	return s
}

func (s *OnsTopicListRequestTag) SetValue(v string) *OnsTopicListRequestTag {
	s.Value = &v
	return s
}

type OnsTopicListResponseBody struct {
	// The information about the topics.
	Data *OnsTopicListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 4A978869-7681-4529-B470-107E1379****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBody) SetData(v *OnsTopicListResponseBodyData) *OnsTopicListResponseBody {
	s.Data = v
	return s
}

func (s *OnsTopicListResponseBody) SetRequestId(v string) *OnsTopicListResponseBody {
	s.RequestId = &v
	return s
}

type OnsTopicListResponseBodyData struct {
	PublishInfoDo []*OnsTopicListResponseBodyDataPublishInfoDo `json:"PublishInfoDo,omitempty" xml:"PublishInfoDo,omitempty" type:"Repeated"`
}

func (s OnsTopicListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBodyData) SetPublishInfoDo(v []*OnsTopicListResponseBodyDataPublishInfoDo) *OnsTopicListResponseBodyData {
	s.PublishInfoDo = v
	return s
}

type OnsTopicListResponseBodyDataPublishInfoDo struct {
	// The time when the topic was created.
	//
	// example:
	//
	// 1570700947000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the instance that contains the topic uses a namespace. Valid values:
	//
	// 	- **true**: The instance uses a separate namespace. The name of each resource must be unique in the instance. The names of resources in different instances can be the same.
	//
	// 	- **false**: The instance does not use a separate namespace. The name of each resource must be globally unique within an instance and across all instances.
	//
	// example:
	//
	// true
	IndependentNaming *bool `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	// The ID of the instance that contains the topic.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The message type. Valid values:
	//
	// 	- **0**: normal messages
	//
	// 	- **1**: partitionally ordered messages
	//
	// 	- **2**: globally ordered messages
	//
	// 	- **4**: transactional messages
	//
	// 	- **5**: scheduled or delayed messages
	//
	// example:
	//
	// 0
	MessageType *int32 `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The user ID of the topic owner. The value of this parameter is an Alibaba account ID.
	//
	// example:
	//
	// 138015630679****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// Indicates the relationship between the current account and the topic. Valid values:
	//
	// 	- **1**: The current account is the owner of the topic.
	//
	// 	- **2**: The current account can publish messages to the topic.
	//
	// 	- **4**: The current account can subscribe to the topic.
	//
	// 	- **6**: The current account can publish messages to and subscribe to the topic.
	//
	// example:
	//
	// 6
	Relation *int32 `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The relationship between the current account and the topic. The value of this parameter indicates whether the current account is the owner of the topic, and whether the current account can subscribe or publish messages to the topic. the topic.
	//
	// example:
	//
	// Publish and subscribe
	RelationName *string `json:"RelationName,omitempty" xml:"RelationName,omitempty"`
	// The description of the topic.
	//
	// example:
	//
	// Test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the topic that is asynchronously created. Valid values:
	//
	// 	- **0**: The topic is being created.
	//
	// 	- **1**: The topic is being used.
	//
	// example:
	//
	// 0
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The tags that are attached to the topic.
	Tags *OnsTopicListResponseBodyDataPublishInfoDoTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The topic name.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicListResponseBodyDataPublishInfoDo) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListResponseBodyDataPublishInfoDo) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetCreateTime(v int64) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.CreateTime = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetIndependentNaming(v bool) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.IndependentNaming = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetInstanceId(v string) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetMessageType(v int32) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.MessageType = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetOwner(v string) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.Owner = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetRelation(v int32) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.Relation = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetRelationName(v string) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.RelationName = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetRemark(v string) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.Remark = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetServiceStatus(v int32) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.ServiceStatus = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetTags(v *OnsTopicListResponseBodyDataPublishInfoDoTags) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.Tags = v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDo) SetTopic(v string) *OnsTopicListResponseBodyDataPublishInfoDo {
	s.Topic = &v
	return s
}

type OnsTopicListResponseBodyDataPublishInfoDoTags struct {
	Tag []*OnsTopicListResponseBodyDataPublishInfoDoTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsTopicListResponseBodyDataPublishInfoDoTags) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListResponseBodyDataPublishInfoDoTags) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTags) SetTag(v []*OnsTopicListResponseBodyDataPublishInfoDoTagsTag) *OnsTopicListResponseBodyDataPublishInfoDoTags {
	s.Tag = v
	return s
}

type OnsTopicListResponseBodyDataPublishInfoDoTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// CartService
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// SrviceA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsTopicListResponseBodyDataPublishInfoDoTagsTag) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListResponseBodyDataPublishInfoDoTagsTag) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTagsTag) SetKey(v string) *OnsTopicListResponseBodyDataPublishInfoDoTagsTag {
	s.Key = &v
	return s
}

func (s *OnsTopicListResponseBodyDataPublishInfoDoTagsTag) SetValue(v string) *OnsTopicListResponseBodyDataPublishInfoDoTagsTag {
	s.Value = &v
	return s
}

type OnsTopicListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicListResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponse) SetHeaders(v map[string]*string) *OnsTopicListResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicListResponse) SetStatusCode(v int32) *OnsTopicListResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicListResponse) SetBody(v *OnsTopicListResponseBody) *OnsTopicListResponse {
	s.Body = v
	return s
}

type OnsTopicStatusRequest struct {
	// The ID of the instance that contains the topic you want to query.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the topic that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicStatusRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusRequest) SetInstanceId(v string) *OnsTopicStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicStatusRequest) SetTopic(v string) *OnsTopicStatusRequest {
	s.Topic = &v
	return s
}

type OnsTopicStatusResponseBody struct {
	// The data structure of the queried topic.
	Data *OnsTopicStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 427EE49D-D762-41FB-8F3D-9BAC96C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicStatusResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusResponseBody) SetData(v *OnsTopicStatusResponseBodyData) *OnsTopicStatusResponseBody {
	s.Data = v
	return s
}

func (s *OnsTopicStatusResponseBody) SetRequestId(v string) *OnsTopicStatusResponseBody {
	s.RequestId = &v
	return s
}

type OnsTopicStatusResponseBodyData struct {
	// The point in time when the latest message is ready in the topic. If no message exists in the topic, the return value of this parameter is 0.
	//
	// The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// For information about the definition of ready messages and ready time, see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// example:
	//
	// 1570864984364
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	// Indicates the operations that you can perform on the topic. Valid values:
	//
	// 	- **2**: You can publish messages to the topic.
	//
	// 	- **4**: You can subscribe to the topic.
	//
	// 	- **6**: You can publish messages to and subscribe to the topic.
	//
	// example:
	//
	// 6
	Perm *int32 `json:"Perm,omitempty" xml:"Perm,omitempty"`
	// The total number of messages in all partitions of the topic.
	//
	// example:
	//
	// 2310
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s OnsTopicStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusResponseBodyData) SetLastTimeStamp(v int64) *OnsTopicStatusResponseBodyData {
	s.LastTimeStamp = &v
	return s
}

func (s *OnsTopicStatusResponseBodyData) SetPerm(v int32) *OnsTopicStatusResponseBodyData {
	s.Perm = &v
	return s
}

func (s *OnsTopicStatusResponseBodyData) SetTotalCount(v int64) *OnsTopicStatusResponseBodyData {
	s.TotalCount = &v
	return s
}

type OnsTopicStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicStatusResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusResponse) SetHeaders(v map[string]*string) *OnsTopicStatusResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicStatusResponse) SetStatusCode(v int32) *OnsTopicStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicStatusResponse) SetBody(v *OnsTopicStatusResponseBody) *OnsTopicStatusResponse {
	s.Body = v
	return s
}

type OnsTopicSubDetailRequest struct {
	// The ID of the instance that contains the topic you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the topic that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicSubDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicSubDetailRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailRequest) SetInstanceId(v string) *OnsTopicSubDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicSubDetailRequest) SetTopic(v string) *OnsTopicSubDetailRequest {
	s.Topic = &v
	return s
}

type OnsTopicSubDetailResponseBody struct {
	// The data returned.
	Data *OnsTopicSubDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 87B6207F-2908-42B5-A134-84956DCA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicSubDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicSubDetailResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBody) SetData(v *OnsTopicSubDetailResponseBodyData) *OnsTopicSubDetailResponseBody {
	s.Data = v
	return s
}

func (s *OnsTopicSubDetailResponseBody) SetRequestId(v string) *OnsTopicSubDetailResponseBody {
	s.RequestId = &v
	return s
}

type OnsTopicSubDetailResponseBodyData struct {
	// The information about the online consumer groups that subscribe to the topic.
	SubscriptionDataList *OnsTopicSubDetailResponseBodyDataSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Struct"`
	// The topic name.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicSubDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicSubDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBodyData) SetSubscriptionDataList(v *OnsTopicSubDetailResponseBodyDataSubscriptionDataList) *OnsTopicSubDetailResponseBodyData {
	s.SubscriptionDataList = v
	return s
}

func (s *OnsTopicSubDetailResponseBodyData) SetTopic(v string) *OnsTopicSubDetailResponseBodyData {
	s.Topic = &v
	return s
}

type OnsTopicSubDetailResponseBodyDataSubscriptionDataList struct {
	SubscriptionDataList []*OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Repeated"`
}

func (s OnsTopicSubDetailResponseBodyDataSubscriptionDataList) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicSubDetailResponseBodyDataSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataList) SetSubscriptionDataList(v []*OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) *OnsTopicSubDetailResponseBodyDataSubscriptionDataList {
	s.SubscriptionDataList = v
	return s
}

type OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList struct {
	// The ID of the consumer group that subscribes to the topic.
	//
	// example:
	//
	// GID_test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The consumption mode. Valid values:
	//
	// 	- **CLUSTERING**: the clustering consumption mode
	//
	// 	- **BROADCASTING**: the broadcasting consumption mode
	//
	// For more information about consumption modes, see [Clustering consumption and broadcasting consumption](https://help.aliyun.com/document_detail/43163.html).
	//
	// example:
	//
	// CLUSTERING
	MessageModel *string `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	Online       *string `json:"Online,omitempty" xml:"Online,omitempty"`
	// The expression based on which consumers in the consumer group subscribe to the topic.
	//
	// example:
	//
	// *
	SubString *string `json:"SubString,omitempty" xml:"SubString,omitempty"`
}

func (s OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetGroupId(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.GroupId = &v
	return s
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetMessageModel(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.MessageModel = &v
	return s
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetOnline(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.Online = &v
	return s
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetSubString(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.SubString = &v
	return s
}

type OnsTopicSubDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicSubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicSubDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicSubDetailResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponse) SetHeaders(v map[string]*string) *OnsTopicSubDetailResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicSubDetailResponse) SetStatusCode(v int32) *OnsTopicSubDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicSubDetailResponse) SetBody(v *OnsTopicSubDetailResponseBody) *OnsTopicSubDetailResponse {
	s.Body = v
	return s
}

type OnsTopicUpdateRequest struct {
	// The ID of the instance to which the topic belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The read/write mode that you want to configure for the topic. Valid values:
	//
	// 	- **6**: Both read and write operations are allowed.
	//
	// 	- **4**: Write operations are forbidden.
	//
	// 	- **2**: Read operations are forbidden.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	Perm *int32 `json:"Perm,omitempty" xml:"Perm,omitempty"`
	// The name of the topic that you want to manage.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicUpdateRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicUpdateRequest) SetInstanceId(v string) *OnsTopicUpdateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicUpdateRequest) SetPerm(v int32) *OnsTopicUpdateRequest {
	s.Perm = &v
	return s
}

func (s *OnsTopicUpdateRequest) SetTopic(v string) *OnsTopicUpdateRequest {
	s.Topic = &v
	return s
}

type OnsTopicUpdateResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 81979ADA-4A78-4F64-9DEC-5700446D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicUpdateResponseBody) SetRequestId(v string) *OnsTopicUpdateResponseBody {
	s.RequestId = &v
	return s
}

type OnsTopicUpdateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTopicUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTopicUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicUpdateResponse) GoString() string {
	return s.String()
}

func (s *OnsTopicUpdateResponse) SetHeaders(v map[string]*string) *OnsTopicUpdateResponse {
	s.Headers = v
	return s
}

func (s *OnsTopicUpdateResponse) SetStatusCode(v int32) *OnsTopicUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTopicUpdateResponse) SetBody(v *OnsTopicUpdateResponseBody) *OnsTopicUpdateResponse {
	s.Body = v
	return s
}

type OnsTraceGetResultRequest struct {
	// The ID of the instance to which the message you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the task that was created to query the trace of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 272967562652883649157096685****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTraceGetResultRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultRequest) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultRequest) SetInstanceId(v string) *OnsTraceGetResultRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTraceGetResultRequest) SetQueryId(v string) *OnsTraceGetResultRequest {
	s.QueryId = &v
	return s
}

func (s *OnsTraceGetResultRequest) SetTopic(v string) *OnsTraceGetResultRequest {
	s.Topic = &v
	return s
}

type OnsTraceGetResultResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 84EE24D2-851F-40D6-B99E-4D6AB909****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the message trace.
	TraceData *OnsTraceGetResultResponseBodyTraceData `json:"TraceData,omitempty" xml:"TraceData,omitempty" type:"Struct"`
}

func (s OnsTraceGetResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBody) SetRequestId(v string) *OnsTraceGetResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTraceGetResultResponseBody) SetTraceData(v *OnsTraceGetResultResponseBodyTraceData) *OnsTraceGetResultResponseBody {
	s.TraceData = v
	return s
}

type OnsTraceGetResultResponseBodyTraceData struct {
	// The point in time when the task was created.
	//
	// example:
	//
	// 1570966857000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message that is queried.
	//
	// example:
	//
	// 1E05791C117818B4AAC23B1BB0CE****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The key of the message that is queried.
	//
	// example:
	//
	// ORDERID_100
	MsgKey *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 272967562652883649157096685****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **finish**: The task is complete.
	//
	// 	- **working**: The task is in progress.
	//
	// 	- **removed**: The task is deleted.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The topic in which the message is stored.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The details of the message trace.
	TraceList *OnsTraceGetResultResponseBodyTraceDataTraceList `json:"TraceList,omitempty" xml:"TraceList,omitempty" type:"Struct"`
	// The most recent point in time when the task was updated.
	//
	// example:
	//
	// 1570966877000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user who created the task.
	//
	// example:
	//
	// 27296756265288****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceData) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceData) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetCreateTime(v int64) *OnsTraceGetResultResponseBodyTraceData {
	s.CreateTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetInstanceId(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.InstanceId = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetMsgId(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.MsgId = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetMsgKey(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.MsgKey = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetQueryId(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.QueryId = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetStatus(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.Status = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetTopic(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.Topic = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetTraceList(v *OnsTraceGetResultResponseBodyTraceDataTraceList) *OnsTraceGetResultResponseBodyTraceData {
	s.TraceList = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetUpdateTime(v int64) *OnsTraceGetResultResponseBodyTraceData {
	s.UpdateTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetUserId(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.UserId = &v
	return s
}

type OnsTraceGetResultResponseBodyTraceDataTraceList struct {
	TraceMapDo []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo `json:"TraceMapDo,omitempty" xml:"TraceMapDo,omitempty" type:"Repeated"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceList) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceList) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetTraceMapDo(v []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.TraceMapDo = v
	return s
}

type OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo struct {
	// The address of the producer that generated the message.
	//
	// example:
	//
	// ``30.5.**.**``
	BornHost *string `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	// The period of time that the system took to send the message. Unit: milliseconds.
	//
	// example:
	//
	// 24
	CostTime *int32 `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 0BC1F01800002A9F000000531246****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The key of the message.
	//
	// example:
	//
	// ORDERID_100
	MsgKey *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	// The ID of the group that contains the producer.
	//
	// example:
	//
	// GID_test
	PubGroupName *string `json:"PubGroupName,omitempty" xml:"PubGroupName,omitempty"`
	// The point in time when the message was sent.
	//
	// example:
	//
	// 1570850870478
	PubTime *int64 `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// Indicates whether the message is sent. Valid values:
	//
	// 	- **SEND_SUCCESS**: The message is sent.
	//
	// 	- **SEND_FAILED**: The message failed to be sent.
	//
	// 	- **SEND_ROLLBACK:*	- The message is a transactional message and is rolled back.
	//
	// 	- **SEND_UNKNOWN:*	- The message is a transactional message and is not committed.
	//
	// 	- **SEND_DELAY:*	- The message is a scheduled or delayed message and is waiting to be consumed at the specified point in time.
	//
	// example:
	//
	// SEND_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The consumption traces of the message.
	SubList *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList `json:"SubList,omitempty" xml:"SubList,omitempty" type:"Struct"`
	// The tag of the message.
	//
	// example:
	//
	// TagA
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetBornHost(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.BornHost = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetCostTime(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.CostTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetMsgId(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.MsgId = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetMsgKey(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.MsgKey = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetPubGroupName(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.PubGroupName = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetPubTime(v int64) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.PubTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetStatus(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.Status = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetSubList(v *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.SubList = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetTag(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.Tag = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo) SetTopic(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDo {
	s.Topic = &v
	return s
}

type OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList struct {
	SubMapDo []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo `json:"SubMapDo,omitempty" xml:"SubMapDo,omitempty" type:"Repeated"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList) SetSubMapDo(v []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList {
	s.SubMapDo = v
	return s
}

type OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo struct {
	// The information about message consumption by consumers in the group.
	ClientList *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList `json:"ClientList,omitempty" xml:"ClientList,omitempty" type:"Struct"`
	// The number of consumption failures.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The ID of the consumer group.
	//
	// example:
	//
	// GID_test
	SubGroupName *string `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	// The number of successful consumptions.
	//
	// example:
	//
	// 1
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) SetClientList(v *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo {
	s.ClientList = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) SetFailCount(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo {
	s.FailCount = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) SetSubGroupName(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo {
	s.SubGroupName = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo) SetSuccessCount(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDo {
	s.SuccessCount = &v
	return s
}

type OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList struct {
	SubClientInfoDo []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo `json:"SubClientInfoDo,omitempty" xml:"SubClientInfoDo,omitempty" type:"Repeated"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList) SetSubClientInfoDo(v []*OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList {
	s.SubClientInfoDo = v
	return s
}

type OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo struct {
	// The address of the consumer.
	//
	// example:
	//
	// ``30.5.**.**``
	ClientHost *string `json:"ClientHost,omitempty" xml:"ClientHost,omitempty"`
	// The period of time that the system took to consume the message. Unit: milliseconds.
	//
	// example:
	//
	// 43
	CostTime *int32 `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	// The number of attempts that the ApsaraMQ for RocketMQ broker tried to send the message to the consumer.
	//
	// example:
	//
	// 1
	ReconsumeTimes *int32 `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	// Indicates whether the message is consumed. Valid values:
	//
	// 	- **CONSUME_FAILED**: The message failed to be consumed.
	//
	// 	- **CONSUME_SUCCESS**: The message is consumed.
	//
	// 	- **CONSUME_NOT_RETURN:*	- No responses are returned.
	//
	// 	- **SEND_UNKNOWN:*	- The message is a transactional message and is not committed.
	//
	// 	- **SEND_DELAY:*	- The message is a scheduled or delayed message and is waiting to be consumed at the specified point in time.
	//
	// example:
	//
	// CONSUME_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the group that contains the consumer.
	//
	// example:
	//
	// GID_test
	SubGroupName *string `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	// The earliest point in time when the message was consumed.
	//
	// example:
	//
	// 1570851590511
	SubTime *int64 `json:"SubTime,omitempty" xml:"SubTime,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetClientHost(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.ClientHost = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetCostTime(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.CostTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetReconsumeTimes(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetStatus(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.Status = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetSubGroupName(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.SubGroupName = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo) SetSubTime(v int64) *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientListSubClientInfoDo {
	s.SubTime = &v
	return s
}

type OnsTraceGetResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTraceGetResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTraceGetResultResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponse) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponse) SetHeaders(v map[string]*string) *OnsTraceGetResultResponse {
	s.Headers = v
	return s
}

func (s *OnsTraceGetResultResponse) SetStatusCode(v int32) *OnsTraceGetResultResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTraceGetResultResponse) SetBody(v *OnsTraceGetResultResponseBody) *OnsTraceGetResultResponse {
	s.Body = v
	return s
}

type OnsTraceQueryByMsgIdRequest struct {
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570852800000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570968000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance that contains the specified topic.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1E05791C117818B4AAC23B1BB0CE****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The topic that contains the message you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTraceQueryByMsgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceQueryByMsgIdRequest) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgIdRequest) SetBeginTime(v int64) *OnsTraceQueryByMsgIdRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsTraceQueryByMsgIdRequest) SetEndTime(v int64) *OnsTraceQueryByMsgIdRequest {
	s.EndTime = &v
	return s
}

func (s *OnsTraceQueryByMsgIdRequest) SetInstanceId(v string) *OnsTraceQueryByMsgIdRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTraceQueryByMsgIdRequest) SetMsgId(v string) *OnsTraceQueryByMsgIdRequest {
	s.MsgId = &v
	return s
}

func (s *OnsTraceQueryByMsgIdRequest) SetTopic(v string) *OnsTraceQueryByMsgIdRequest {
	s.Topic = &v
	return s
}

type OnsTraceQueryByMsgIdResponseBody struct {
	// The ID of the query task. You can call the [OnsTraceGetResult](https://help.aliyun.com/document_detail/59832.html) operation to query the details of the message trace based on the task ID.
	//
	// example:
	//
	// 272967562652883649157096685****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// B93332A3-160D-404F-880F-1F8736D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTraceQueryByMsgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceQueryByMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgIdResponseBody) SetQueryId(v string) *OnsTraceQueryByMsgIdResponseBody {
	s.QueryId = &v
	return s
}

func (s *OnsTraceQueryByMsgIdResponseBody) SetRequestId(v string) *OnsTraceQueryByMsgIdResponseBody {
	s.RequestId = &v
	return s
}

type OnsTraceQueryByMsgIdResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTraceQueryByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTraceQueryByMsgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceQueryByMsgIdResponse) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgIdResponse) SetHeaders(v map[string]*string) *OnsTraceQueryByMsgIdResponse {
	s.Headers = v
	return s
}

func (s *OnsTraceQueryByMsgIdResponse) SetStatusCode(v int32) *OnsTraceQueryByMsgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTraceQueryByMsgIdResponse) SetBody(v *OnsTraceQueryByMsgIdResponseBody) *OnsTraceQueryByMsgIdResponse {
	s.Body = v
	return s
}

type OnsTraceQueryByMsgKeyRequest struct {
	// The start of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570852800000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570968000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance that contains the specified topic.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The key of the message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ORDERID_100
	MsgKey *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	// The topic that contains the message you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTraceQueryByMsgKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceQueryByMsgKeyRequest) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgKeyRequest) SetBeginTime(v int64) *OnsTraceQueryByMsgKeyRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyRequest) SetEndTime(v int64) *OnsTraceQueryByMsgKeyRequest {
	s.EndTime = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyRequest) SetInstanceId(v string) *OnsTraceQueryByMsgKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyRequest) SetMsgKey(v string) *OnsTraceQueryByMsgKeyRequest {
	s.MsgKey = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyRequest) SetTopic(v string) *OnsTraceQueryByMsgKeyRequest {
	s.Topic = &v
	return s
}

type OnsTraceQueryByMsgKeyResponseBody struct {
	// The ID of the query task. You can call the [OnsTraceGetResult](https://help.aliyun.com/document_detail/59832.html) operation to query the details of the message trace based on the task ID.
	//
	// example:
	//
	// 272967562652883649157096685****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// F8654231-122A-4DBD-801F-38E35538****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTraceQueryByMsgKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceQueryByMsgKeyResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgKeyResponseBody) SetQueryId(v string) *OnsTraceQueryByMsgKeyResponseBody {
	s.QueryId = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyResponseBody) SetRequestId(v string) *OnsTraceQueryByMsgKeyResponseBody {
	s.RequestId = &v
	return s
}

type OnsTraceQueryByMsgKeyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTraceQueryByMsgKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTraceQueryByMsgKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceQueryByMsgKeyResponse) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgKeyResponse) SetHeaders(v map[string]*string) *OnsTraceQueryByMsgKeyResponse {
	s.Headers = v
	return s
}

func (s *OnsTraceQueryByMsgKeyResponse) SetStatusCode(v int32) *OnsTraceQueryByMsgKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyResponse) SetBody(v *OnsTraceQueryByMsgKeyResponseBody) *OnsTraceQueryByMsgKeyResponse {
	s.Body = v
	return s
}

type OnsTrendGroupOutputTpsRequest struct {
	// The timestamp that indicates the beginning of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570852800000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The timestamp that indicates the end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570868400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the consumer group that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The sampling period. Unit: minutes. Valid values: 1, 5, and 10.
	//
	// example:
	//
	// 10
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The name of the topic that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The type of information that you want to query. Valid values:
	//
	// 	- **0**: the number of messages that are consumed during each sampling period.
	//
	// 	- **1**: the TPS for message consumption during each sampling period.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OnsTrendGroupOutputTpsRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendGroupOutputTpsRequest) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsRequest) SetBeginTime(v int64) *OnsTrendGroupOutputTpsRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetEndTime(v int64) *OnsTrendGroupOutputTpsRequest {
	s.EndTime = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetGroupId(v string) *OnsTrendGroupOutputTpsRequest {
	s.GroupId = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetInstanceId(v string) *OnsTrendGroupOutputTpsRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetPeriod(v int64) *OnsTrendGroupOutputTpsRequest {
	s.Period = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetTopic(v string) *OnsTrendGroupOutputTpsRequest {
	s.Topic = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetType(v int32) *OnsTrendGroupOutputTpsRequest {
	s.Type = &v
	return s
}

type OnsTrendGroupOutputTpsResponseBody struct {
	// The data returned.
	Data *OnsTrendGroupOutputTpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// CE57AEDC-8FD2-43ED-8E3B-1F878077****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTrendGroupOutputTpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBody) SetData(v *OnsTrendGroupOutputTpsResponseBodyData) *OnsTrendGroupOutputTpsResponseBody {
	s.Data = v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBody) SetRequestId(v string) *OnsTrendGroupOutputTpsResponseBody {
	s.RequestId = &v
	return s
}

type OnsTrendGroupOutputTpsResponseBodyData struct {
	// The data set returned based on sampling period.
	Records *OnsTrendGroupOutputTpsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// The name of the table.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx%test@MQ_INST_111111111111_DOxxxxxx%GID_test trend chart of delivered messages
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The unit of the timestamp.
	//
	// example:
	//
	// time
	XUnit *string `json:"XUnit,omitempty" xml:"XUnit,omitempty"`
	// The total number of messages.
	//
	// example:
	//
	// msg
	YUnit *string `json:"YUnit,omitempty" xml:"YUnit,omitempty"`
}

func (s OnsTrendGroupOutputTpsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) SetRecords(v *OnsTrendGroupOutputTpsResponseBodyDataRecords) *OnsTrendGroupOutputTpsResponseBodyData {
	s.Records = v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) SetTitle(v string) *OnsTrendGroupOutputTpsResponseBodyData {
	s.Title = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) SetXUnit(v string) *OnsTrendGroupOutputTpsResponseBodyData {
	s.XUnit = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) SetYUnit(v string) *OnsTrendGroupOutputTpsResponseBodyData {
	s.YUnit = &v
	return s
}

type OnsTrendGroupOutputTpsResponseBodyDataRecords struct {
	StatsDataDo []*OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo `json:"StatsDataDo,omitempty" xml:"StatsDataDo,omitempty" type:"Repeated"`
}

func (s OnsTrendGroupOutputTpsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecords) SetStatsDataDo(v []*OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) *OnsTrendGroupOutputTpsResponseBodyDataRecords {
	s.StatsDataDo = v
	return s
}

type OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo struct {
	// The X axis. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1570867800000
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y axis. This parameter indicates the TPS for message consumption or the number of messages that are consumed.
	//
	// example:
	//
	// 0
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) SetX(v int64) *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo {
	s.X = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo) SetY(v float32) *OnsTrendGroupOutputTpsResponseBodyDataRecordsStatsDataDo {
	s.Y = &v
	return s
}

type OnsTrendGroupOutputTpsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTrendGroupOutputTpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTrendGroupOutputTpsResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponse) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponse) SetHeaders(v map[string]*string) *OnsTrendGroupOutputTpsResponse {
	s.Headers = v
	return s
}

func (s *OnsTrendGroupOutputTpsResponse) SetStatusCode(v int32) *OnsTrendGroupOutputTpsResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponse) SetBody(v *OnsTrendGroupOutputTpsResponseBody) *OnsTrendGroupOutputTpsResponse {
	s.Body = v
	return s
}

type OnsTrendTopicInputTpsRequest struct {
	// The timestamp that indicates the beginning of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570852800000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The timestamp that indicates the end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570868400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance to which the topic you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The sampling period. Unit: minutes. Valid values: 1, 5, and 10.
	//
	// example:
	//
	// 10
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The name of the topic that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The type of information that you want to query. Valid values:
	//
	// 	- **0**: the number of messages that are published to the topic during each sampling period.
	//
	// 	- **1**: the TPS for message publishing in the topic during each sampling period.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OnsTrendTopicInputTpsRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendTopicInputTpsRequest) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsRequest) SetBeginTime(v int64) *OnsTrendTopicInputTpsRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetEndTime(v int64) *OnsTrendTopicInputTpsRequest {
	s.EndTime = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetInstanceId(v string) *OnsTrendTopicInputTpsRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetPeriod(v int64) *OnsTrendTopicInputTpsRequest {
	s.Period = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetTopic(v string) *OnsTrendTopicInputTpsRequest {
	s.Topic = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetType(v int32) *OnsTrendTopicInputTpsRequest {
	s.Type = &v
	return s
}

type OnsTrendTopicInputTpsResponseBody struct {
	// The data returned.
	Data *OnsTrendTopicInputTpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// E213AD8A-0730-4B3D-A35A-340DA47D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTrendTopicInputTpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBody) SetData(v *OnsTrendTopicInputTpsResponseBodyData) *OnsTrendTopicInputTpsResponseBody {
	s.Data = v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBody) SetRequestId(v string) *OnsTrendTopicInputTpsResponseBody {
	s.RequestId = &v
	return s
}

type OnsTrendTopicInputTpsResponseBodyData struct {
	// The data set returned based on sampling period.
	Records *OnsTrendTopicInputTpsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// The name of the table.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx%test trend of received messages
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The unit of the timestamp.
	//
	// example:
	//
	// time
	XUnit *string `json:"XUnit,omitempty" xml:"XUnit,omitempty"`
	// The unit of the Y axis.
	//
	// example:
	//
	// msg
	YUnit *string `json:"YUnit,omitempty" xml:"YUnit,omitempty"`
}

func (s OnsTrendTopicInputTpsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBodyData) SetRecords(v *OnsTrendTopicInputTpsResponseBodyDataRecords) *OnsTrendTopicInputTpsResponseBodyData {
	s.Records = v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyData) SetTitle(v string) *OnsTrendTopicInputTpsResponseBodyData {
	s.Title = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyData) SetXUnit(v string) *OnsTrendTopicInputTpsResponseBodyData {
	s.XUnit = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyData) SetYUnit(v string) *OnsTrendTopicInputTpsResponseBodyData {
	s.YUnit = &v
	return s
}

type OnsTrendTopicInputTpsResponseBodyDataRecords struct {
	StatsDataDo []*OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo `json:"StatsDataDo,omitempty" xml:"StatsDataDo,omitempty" type:"Repeated"`
}

func (s OnsTrendTopicInputTpsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecords) SetStatsDataDo(v []*OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) *OnsTrendTopicInputTpsResponseBodyDataRecords {
	s.StatsDataDo = v
	return s
}

type OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo struct {
	// The X axis. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1570852800000
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y axis. This parameter indicates the TPS for message publishing or the number of messages that are published.
	//
	// example:
	//
	// 0
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) SetX(v int64) *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo {
	s.X = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo) SetY(v float32) *OnsTrendTopicInputTpsResponseBodyDataRecordsStatsDataDo {
	s.Y = &v
	return s
}

type OnsTrendTopicInputTpsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTrendTopicInputTpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTrendTopicInputTpsResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponse) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponse) SetHeaders(v map[string]*string) *OnsTrendTopicInputTpsResponse {
	s.Headers = v
	return s
}

func (s *OnsTrendTopicInputTpsResponse) SetStatusCode(v int32) *OnsTrendTopicInputTpsResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponse) SetBody(v *OnsTrendTopicInputTpsResponseBody) *OnsTrendTopicInputTpsResponse {
	s.Body = v
	return s
}

type OpenOnsServiceResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 2068689****0272
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 8C5B4603-8977-4513-AB60-9C3E2F88****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenOnsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenOnsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenOnsServiceResponseBody) SetOrderId(v string) *OpenOnsServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenOnsServiceResponseBody) SetRequestId(v string) *OpenOnsServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenOnsServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenOnsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenOnsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenOnsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenOnsServiceResponse) SetHeaders(v map[string]*string) *OpenOnsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenOnsServiceResponse) SetStatusCode(v int32) *OpenOnsServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenOnsServiceResponse) SetBody(v *OpenOnsServiceResponseBody) *OpenOnsServiceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The ID of the ApsaraMQ for RocketMQ instance that contains the resource to which you want to attach tags.
	//
	// > This parameter is required when you attach tags to a topic or a group.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The resource IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// TopicA
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource to which you want to attach tags. Valid values:
	//
	// 	- **INSTANCE**
	//
	// 	- **TOPIC**
	//
	// 	- **GROUP**
	//
	// This parameter is required.
	//
	// example:
	//
	// TOPIC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that you want to attach to the resource.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetInstanceId(v string) *TagResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The tag key. If you configure this parameter, you must also configure the **Value*	- parameter.****
	//
	// 	- The value of this parameter cannot be an empty string.
	//
	// 	- A tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. A tag key cannot start with `acs:` or `aliyun`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that you want to attach to the resource. If you configure this parameter, you must also configure the **Key*	- parameter.****
	//
	// 	- The value of this parameter can be an empty string.
	//
	// 	- A tag value must be 1 to 128 characters in length and cannot contain `http://` or `https://`. A tag value cannot start with `acs:` or `aliyun`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 301D2CBE-66F8-403D-AEC0-82582478****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags that are attached to the specified resource. This parameter takes effect only if the **TagKey*	- parameter is empty. Default value: **false**.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required when you detach tags from a topic or a group.
	//
	// example:
	//
	// MQ_INST_188077086902****_BX4jvZZG
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The resource IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// TopicA
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource from which you want to detach tags. Valid values:
	//
	// 	- **INSTANCE**
	//
	// 	- **TOPIC**
	//
	// 	- **GROUP**
	//
	// This parameter is required.
	//
	// example:
	//
	// TOPIC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys of the resource.
	//
	// example:
	//
	// CartService
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetInstanceId(v string) *UntagResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 19780F2E-7841-4E0F-A5D9-C64A0530****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("ons.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("ons.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("ons.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("ons.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("ons.aliyuncs.com"),
		"cn-edge-1":                   tea.String("ons.aliyuncs.com"),
		"cn-fujian":                   tea.String("ons.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("ons.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("ons.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("ons.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("ons.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("ons.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("ons.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("ons.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("ons.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("ons.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("ons.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("ons.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("ons.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("ons.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("ons.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("ons.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("ons.aliyuncs.com"),
		"cn-wuhan":                    tea.String("ons.aliyuncs.com"),
		"cn-yushanfang":               tea.String("ons.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("ons.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("ons.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("ons.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("ons.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("ons.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ons"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are attached to a specified resource.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// When you call the **ListTagResources*	- operation, specify at least one of the following parameters in the request: **Key*	- and **ResourceId**. You can specify a resource ID to query all tags that are attached to the specified resource. You can also specify a tag key to query the tag value and resource to which the tag is attached.
//
//   - If you include the **Key*	- parameter in a request, you can obtain the tag value and the ID of the resource to which the tag is attached.
//
//   - If you include the **ResourceId*	- parameter in a request, you can obtain the keys and values of all tags that are attached to the specified resource.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the tags that are attached to a specified resource.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// When you call the **ListTagResources*	- operation, specify at least one of the following parameters in the request: **Key*	- and **ResourceId**. You can specify a resource ID to query all tags that are attached to the specified resource. You can also specify a tag key to query the tag value and resource to which the tag is attached.
//
//   - If you include the **Key*	- parameter in a request, you can obtain the tag value and the ID of the resource to which the tag is attached.
//
//   - If you include the **ResourceId*	- parameter in a request, you can obtain the keys and values of all tags that are attached to the specified resource.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about message accumulation in topics to which a specified consumer group subscribes. The returned information includes the number of accumulated messages and the consumption latency.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation in scenarios in which you want to know the message consumption progress of a specified consumer group in production environments. You can obtain the information about message consumption and consumption latency based on the returned information. This operation returns the total number of accumulated messages in all topics to which the specified consumer group subscribes and the number of accumulated messages in each topic.
//
// @param request - OnsConsumerAccumulateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsConsumerAccumulateResponse
func (client *Client) OnsConsumerAccumulateWithOptions(request *OnsConsumerAccumulateRequest, runtime *util.RuntimeOptions) (_result *OnsConsumerAccumulateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		query["Detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsConsumerAccumulate"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsConsumerAccumulateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsConsumerAccumulateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about message accumulation in topics to which a specified consumer group subscribes. The returned information includes the number of accumulated messages and the consumption latency.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation in scenarios in which you want to know the message consumption progress of a specified consumer group in production environments. You can obtain the information about message consumption and consumption latency based on the returned information. This operation returns the total number of accumulated messages in all topics to which the specified consumer group subscribes and the number of accumulated messages in each topic.
//
// @param request - OnsConsumerAccumulateRequest
//
// @return OnsConsumerAccumulateResponse
func (client *Client) OnsConsumerAccumulate(request *OnsConsumerAccumulateRequest) (_result *OnsConsumerAccumulateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsConsumerAccumulateResponse{}
	_body, _err := client.OnsConsumerAccumulateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the client connection status of a specified consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// When messages are accumulated in a topic, you can call this operation to check whether a consumer is online.
//
// @param request - OnsConsumerGetConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsConsumerGetConnectionResponse
func (client *Client) OnsConsumerGetConnectionWithOptions(request *OnsConsumerGetConnectionRequest, runtime *util.RuntimeOptions) (_result *OnsConsumerGetConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsConsumerGetConnection"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsConsumerGetConnectionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsConsumerGetConnectionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the client connection status of a specified consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// When messages are accumulated in a topic, you can call this operation to check whether a consumer is online.
//
// @param request - OnsConsumerGetConnectionRequest
//
// @return OnsConsumerGetConnectionResponse
func (client *Client) OnsConsumerGetConnection(request *OnsConsumerGetConnectionRequest) (_result *OnsConsumerGetConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsConsumerGetConnectionResponse{}
	_body, _err := client.OnsConsumerGetConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets a consumer offset to a specified timestamp for a consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to clear accumulated messages or reset a consumer offset to a specified timestamp. You can use one of the following methods to clear accumulated messages:
//
//   - Clear all accumulated messages in a specified topic.
//
//   - Clear the messages that were published to the specified topic before a specified point in time.
//
// @param request - OnsConsumerResetOffsetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsConsumerResetOffsetResponse
func (client *Client) OnsConsumerResetOffsetWithOptions(request *OnsConsumerResetOffsetRequest, runtime *util.RuntimeOptions) (_result *OnsConsumerResetOffsetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResetTimestamp)) {
		query["ResetTimestamp"] = request.ResetTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsConsumerResetOffset"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsConsumerResetOffsetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsConsumerResetOffsetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Resets a consumer offset to a specified timestamp for a consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to clear accumulated messages or reset a consumer offset to a specified timestamp. You can use one of the following methods to clear accumulated messages:
//
//   - Clear all accumulated messages in a specified topic.
//
//   - Clear the messages that were published to the specified topic before a specified point in time.
//
// @param request - OnsConsumerResetOffsetRequest
//
// @return OnsConsumerResetOffsetResponse
func (client *Client) OnsConsumerResetOffset(request *OnsConsumerResetOffsetRequest) (_result *OnsConsumerResetOffsetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsConsumerResetOffsetResponse{}
	_body, _err := client.OnsConsumerResetOffsetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about the status of a specified consumer group. This operation returns the transactions per second (TPS) for message consumption, load balancing status, consumer connection status, and whether all consumers in the consumer group subscribe to the same topics and tags.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - You can call this operation in scenarios in which consumers are online and messages are accumulated. You can troubleshoot errors based on the information that is returned by this operation. You can check whether all consumers in the consumer group subscribe to the same topics and tags, and whether load balancing is performed as expected. You can also obtain the information about thread stack traces of online consumers.
//
//   - This operation uses multiple backend operations to query and aggregate data. The system requires a long period of time to process a request. We recommend that you do not frequently call this operation.
//
// @param request - OnsConsumerStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsConsumerStatusResponse
func (client *Client) OnsConsumerStatusWithOptions(request *OnsConsumerStatusRequest, runtime *util.RuntimeOptions) (_result *OnsConsumerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		query["Detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NeedJstack)) {
		query["NeedJstack"] = request.NeedJstack
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsConsumerStatus"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsConsumerStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsConsumerStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the detailed information about the status of a specified consumer group. This operation returns the transactions per second (TPS) for message consumption, load balancing status, consumer connection status, and whether all consumers in the consumer group subscribe to the same topics and tags.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - You can call this operation in scenarios in which consumers are online and messages are accumulated. You can troubleshoot errors based on the information that is returned by this operation. You can check whether all consumers in the consumer group subscribe to the same topics and tags, and whether load balancing is performed as expected. You can also obtain the information about thread stack traces of online consumers.
//
//   - This operation uses multiple backend operations to query and aggregate data. The system requires a long period of time to process a request. We recommend that you do not frequently call this operation.
//
// @param request - OnsConsumerStatusRequest
//
// @return OnsConsumerStatusResponse
func (client *Client) OnsConsumerStatus(request *OnsConsumerStatusRequest) (_result *OnsConsumerStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsConsumerStatusResponse{}
	_body, _err := client.OnsConsumerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the time range within which you can specify a point in time to reset the consumer offset for a specified topic. The time range is from the point in time when the earliest stored message was published to the topic to the point in time when the most recently stored message was published to the topic.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to query the point in time when the earliest stored message was published to a specified topic and the point in time when the most recently stored message was published to the specified topic. You can also call this operation to query the most recent point in time when a message in the topic was consumed. This operation is usually used with the \\*\\*OnsConsumerAccumulate\\*\\	- operation to display the overview of the consumption progress.
//
// @param request - OnsConsumerTimeSpanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsConsumerTimeSpanResponse
func (client *Client) OnsConsumerTimeSpanWithOptions(request *OnsConsumerTimeSpanRequest, runtime *util.RuntimeOptions) (_result *OnsConsumerTimeSpanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsConsumerTimeSpan"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsConsumerTimeSpanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsConsumerTimeSpanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the time range within which you can specify a point in time to reset the consumer offset for a specified topic. The time range is from the point in time when the earliest stored message was published to the topic to the point in time when the most recently stored message was published to the topic.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to query the point in time when the earliest stored message was published to a specified topic and the point in time when the most recently stored message was published to the specified topic. You can also call this operation to query the most recent point in time when a message in the topic was consumed. This operation is usually used with the \\*\\*OnsConsumerAccumulate\\*\\	- operation to display the overview of the consumption progress.
//
// @param request - OnsConsumerTimeSpanRequest
//
// @return OnsConsumerTimeSpanResponse
func (client *Client) OnsConsumerTimeSpan(request *OnsConsumerTimeSpanRequest) (_result *OnsConsumerTimeSpanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsConsumerTimeSpanResponse{}
	_body, _err := client.OnsConsumerTimeSpanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a dead-letter message based on the message ID. The queried information about the dead-letter message includes the point in time when the message is stored, the message body, and attributes such as the message tag and the message key.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// This operation uses the exact match method to query a dead-letter message based on the message ID. You can obtain the message ID that is required to query the information about a dead-letter message from the SendResult parameter that is returned after the message is sent. You can also obtain the message ID by calling the OnsDLQMessagePageQueryByGroupId operation to query multiple messages at a time. The queried information about the dead-letter message includes the point in time when the message is stored, the message body, and attributes such as the message tag and the message key.
//
// @param request - OnsDLQMessageGetByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsDLQMessageGetByIdResponse
func (client *Client) OnsDLQMessageGetByIdWithOptions(request *OnsDLQMessageGetByIdRequest, runtime *util.RuntimeOptions) (_result *OnsDLQMessageGetByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		query["MsgId"] = request.MsgId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsDLQMessageGetById"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsDLQMessageGetByIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsDLQMessageGetByIdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a dead-letter message based on the message ID. The queried information about the dead-letter message includes the point in time when the message is stored, the message body, and attributes such as the message tag and the message key.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// This operation uses the exact match method to query a dead-letter message based on the message ID. You can obtain the message ID that is required to query the information about a dead-letter message from the SendResult parameter that is returned after the message is sent. You can also obtain the message ID by calling the OnsDLQMessagePageQueryByGroupId operation to query multiple messages at a time. The queried information about the dead-letter message includes the point in time when the message is stored, the message body, and attributes such as the message tag and the message key.
//
// @param request - OnsDLQMessageGetByIdRequest
//
// @return OnsDLQMessageGetByIdResponse
func (client *Client) OnsDLQMessageGetById(request *OnsDLQMessageGetByIdRequest) (_result *OnsDLQMessageGetByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsDLQMessageGetByIdResponse{}
	_body, _err := client.OnsDLQMessageGetByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all dead-letter messages in a group within a period of time by page.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - If you do not know the ID of the dead-letter message that you want to query, you can call this operation to query all dead-letter messages that are sent to a specified consumer group within a specified time range. The results are returned by page.
//
//   - We recommend that you specify a short time range to query dead-letter messages in this method. If you specify a long time range, a large number of dead-letter messages are returned. In this case, you cannot find the dead-letter message that you want to query in an efficient manner. You can perform the following steps to query dead-letter messages:
//
//     1.  Perform a paged query by specifying the group ID, start time, end time, and number of entries to return on each page. If matched messages are found, the information about the dead-letter messages on the first page, total number of pages, and task ID are returned by default.
//
//     2.  Specify the task ID and a page number to call this operation again to query the dead-letter messages on the specified page. In this query, the BeginTime, EndTime, and PageSize parameters do not take effect. By default, the system uses the values of these parameters that you specified in the request when you created the specified query task.
//
// @param request - OnsDLQMessagePageQueryByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsDLQMessagePageQueryByGroupIdResponse
func (client *Client) OnsDLQMessagePageQueryByGroupIdWithOptions(request *OnsDLQMessagePageQueryByGroupIdRequest, runtime *util.RuntimeOptions) (_result *OnsDLQMessagePageQueryByGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsDLQMessagePageQueryByGroupId"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsDLQMessagePageQueryByGroupIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsDLQMessagePageQueryByGroupIdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries all dead-letter messages in a group within a period of time by page.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - If you do not know the ID of the dead-letter message that you want to query, you can call this operation to query all dead-letter messages that are sent to a specified consumer group within a specified time range. The results are returned by page.
//
//   - We recommend that you specify a short time range to query dead-letter messages in this method. If you specify a long time range, a large number of dead-letter messages are returned. In this case, you cannot find the dead-letter message that you want to query in an efficient manner. You can perform the following steps to query dead-letter messages:
//
//     1.  Perform a paged query by specifying the group ID, start time, end time, and number of entries to return on each page. If matched messages are found, the information about the dead-letter messages on the first page, total number of pages, and task ID are returned by default.
//
//     2.  Specify the task ID and a page number to call this operation again to query the dead-letter messages on the specified page. In this query, the BeginTime, EndTime, and PageSize parameters do not take effect. By default, the system uses the values of these parameters that you specified in the request when you created the specified query task.
//
// @param request - OnsDLQMessagePageQueryByGroupIdRequest
//
// @return OnsDLQMessagePageQueryByGroupIdResponse
func (client *Client) OnsDLQMessagePageQueryByGroupId(request *OnsDLQMessagePageQueryByGroupIdRequest) (_result *OnsDLQMessagePageQueryByGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsDLQMessagePageQueryByGroupIdResponse{}
	_body, _err := client.OnsDLQMessagePageQueryByGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resends a dead-letter message based on a specified message ID so that the dead-letter message can be consumed by consumers.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - A dead-letter message is a message that still fails to be consumed after the number of consumption retries reaches the upper limit. If the message still cannot be consumed after you re-send it, a message with the same message ID is added to the corresponding dead-letter queue. You can query the message ID on the Dead-letter Queues page in the ApsaraMQ for RocketMQ console or by calling API operations. You can obtain the number of consumption failures for a message based on the number of dead-letter messages with the same message ID in the dead-letter queue.
//
//   - A dead-letter message is a message that fails to be consumed after the number of consumption retries reaches the upper limit. Generally, dead-letter messages are produced because of incorrect consumption logic. We recommend that you troubleshoot the consumption failures and then call this operation to send the message to the consumer group for consumption again.
//
//   - ApsaraMQ for RocketMQ does not manage the status of dead-letter messages based on the consumption status of the dead-letter messages. After you call this operation to send a dead-letter message to a consumer group and the message is consumed, ApsaraMQ for RocketMQ does not remove the dead-letter message from the dead-letter queue. You must manage dead-letter messages and determine whether to send a dead-letter message to a consumer group for consumption. This way, you do not resend or reconsume the messages that are consumed.
//
// @param request - OnsDLQMessageResendByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsDLQMessageResendByIdResponse
func (client *Client) OnsDLQMessageResendByIdWithOptions(request *OnsDLQMessageResendByIdRequest, runtime *util.RuntimeOptions) (_result *OnsDLQMessageResendByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		query["MsgId"] = request.MsgId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsDLQMessageResendById"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsDLQMessageResendByIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsDLQMessageResendByIdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Resends a dead-letter message based on a specified message ID so that the dead-letter message can be consumed by consumers.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - A dead-letter message is a message that still fails to be consumed after the number of consumption retries reaches the upper limit. If the message still cannot be consumed after you re-send it, a message with the same message ID is added to the corresponding dead-letter queue. You can query the message ID on the Dead-letter Queues page in the ApsaraMQ for RocketMQ console or by calling API operations. You can obtain the number of consumption failures for a message based on the number of dead-letter messages with the same message ID in the dead-letter queue.
//
//   - A dead-letter message is a message that fails to be consumed after the number of consumption retries reaches the upper limit. Generally, dead-letter messages are produced because of incorrect consumption logic. We recommend that you troubleshoot the consumption failures and then call this operation to send the message to the consumer group for consumption again.
//
//   - ApsaraMQ for RocketMQ does not manage the status of dead-letter messages based on the consumption status of the dead-letter messages. After you call this operation to send a dead-letter message to a consumer group and the message is consumed, ApsaraMQ for RocketMQ does not remove the dead-letter message from the dead-letter queue. You must manage dead-letter messages and determine whether to send a dead-letter message to a consumer group for consumption. This way, you do not resend or reconsume the messages that are consumed.
//
// @param request - OnsDLQMessageResendByIdRequest
//
// @return OnsDLQMessageResendByIdResponse
func (client *Client) OnsDLQMessageResendById(request *OnsDLQMessageResendByIdRequest) (_result *OnsDLQMessageResendByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsDLQMessageResendByIdResponse{}
	_body, _err := client.OnsDLQMessageResendByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures read permissions on messages for a consumer group that is identified by a group ID.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to configure the permissions for a consumer group to read messages based on a specified region of ApsaraMQ for RocketMQ and a specified group ID. You can call this operation in scenarios in which you want to forbid consumers in a specific group from reading messages.
//
// @param request - OnsGroupConsumerUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsGroupConsumerUpdateResponse
func (client *Client) OnsGroupConsumerUpdateWithOptions(request *OnsGroupConsumerUpdateRequest, runtime *util.RuntimeOptions) (_result *OnsGroupConsumerUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ReadEnable)) {
		query["ReadEnable"] = request.ReadEnable
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsGroupConsumerUpdate"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsGroupConsumerUpdateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsGroupConsumerUpdateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Configures read permissions on messages for a consumer group that is identified by a group ID.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to configure the permissions for a consumer group to read messages based on a specified region of ApsaraMQ for RocketMQ and a specified group ID. You can call this operation in scenarios in which you want to forbid consumers in a specific group from reading messages.
//
// @param request - OnsGroupConsumerUpdateRequest
//
// @return OnsGroupConsumerUpdateResponse
func (client *Client) OnsGroupConsumerUpdate(request *OnsGroupConsumerUpdateRequest) (_result *OnsGroupConsumerUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsGroupConsumerUpdateResponse{}
	_body, _err := client.OnsGroupConsumerUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// When you release a new application or implement new business logic, you need new consumer groups. You can call this operation to create a consumer group.
//
// @param request - OnsGroupCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsGroupCreateResponse
func (client *Client) OnsGroupCreateWithOptions(request *OnsGroupCreateRequest, runtime *util.RuntimeOptions) (_result *OnsGroupCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsGroupCreate"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsGroupCreateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsGroupCreateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// When you release a new application or implement new business logic, you need new consumer groups. You can call this operation to create a consumer group.
//
// @param request - OnsGroupCreateRequest
//
// @return OnsGroupCreateResponse
func (client *Client) OnsGroupCreate(request *OnsGroupCreateRequest) (_result *OnsGroupCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsGroupCreateResponse{}
	_body, _err := client.OnsGroupCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a consumer group.
//
// Description:
//
// >
//
//   - API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - After you delete a group, the consumers in the group immediately stop receiving messages. Exercise caution when you call this operation.
//
// You can call this operation to delete a group when you need to reclaim the resources of the group. For example, after an application is brought offline, you can delete the groups that are used for the application. After you delete a group, the backend of ApsaraMQ for RocketMQ reclaims the resources of the group. The system requires a long period of time to reclaim the resources. We recommend that you do not create a group that uses the same name as a deleted group immediately after you delete the group. If the system fails to delete the specified group, troubleshoot the issue based on the error code.
//
// @param request - OnsGroupDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsGroupDeleteResponse
func (client *Client) OnsGroupDeleteWithOptions(request *OnsGroupDeleteRequest, runtime *util.RuntimeOptions) (_result *OnsGroupDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsGroupDelete"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsGroupDeleteResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsGroupDeleteResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a consumer group.
//
// Description:
//
// >
//
//   - API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - After you delete a group, the consumers in the group immediately stop receiving messages. Exercise caution when you call this operation.
//
// You can call this operation to delete a group when you need to reclaim the resources of the group. For example, after an application is brought offline, you can delete the groups that are used for the application. After you delete a group, the backend of ApsaraMQ for RocketMQ reclaims the resources of the group. The system requires a long period of time to reclaim the resources. We recommend that you do not create a group that uses the same name as a deleted group immediately after you delete the group. If the system fails to delete the specified group, troubleshoot the issue based on the error code.
//
// @param request - OnsGroupDeleteRequest
//
// @return OnsGroupDeleteResponse
func (client *Client) OnsGroupDelete(request *OnsGroupDeleteRequest) (_result *OnsGroupDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsGroupDeleteResponse{}
	_body, _err := client.OnsGroupDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more group IDs.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - OnsGroupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsGroupListResponse
func (client *Client) OnsGroupListWithOptions(request *OnsGroupListRequest, runtime *util.RuntimeOptions) (_result *OnsGroupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsGroupList"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsGroupListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsGroupListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries one or more group IDs.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - OnsGroupListRequest
//
// @return OnsGroupListResponse
func (client *Client) OnsGroupList(request *OnsGroupListRequest) (_result *OnsGroupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsGroupListResponse{}
	_body, _err := client.OnsGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all topics to which a specified consumer group subscribes. If no consumer instance in the consumer group is online, no data is returned.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - OnsGroupSubDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsGroupSubDetailResponse
func (client *Client) OnsGroupSubDetailWithOptions(request *OnsGroupSubDetailRequest, runtime *util.RuntimeOptions) (_result *OnsGroupSubDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsGroupSubDetail"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsGroupSubDetailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsGroupSubDetailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries all topics to which a specified consumer group subscribes. If no consumer instance in the consumer group is online, no data is returned.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - OnsGroupSubDetailRequest
//
// @return OnsGroupSubDetailResponse
func (client *Client) OnsGroupSubDetail(request *OnsGroupSubDetailRequest) (_result *OnsGroupSubDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsGroupSubDetailResponse{}
	_body, _err := client.OnsGroupSubDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information of a ApsaraMQ for RocketMQ instance and the endpoint that a client uses to connect to the ApsaraMQ for RocketMQ instance to send and receive messages.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// To send and receive messages, a client must be connected to a ApsaraMQ for RocketMQ instance by using an endpoint. You can call this operation to query the endpoints of the instance.
//
// @param request - OnsInstanceBaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsInstanceBaseInfoResponse
func (client *Client) OnsInstanceBaseInfoWithOptions(request *OnsInstanceBaseInfoRequest, runtime *util.RuntimeOptions) (_result *OnsInstanceBaseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsInstanceBaseInfo"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsInstanceBaseInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsInstanceBaseInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the basic information of a ApsaraMQ for RocketMQ instance and the endpoint that a client uses to connect to the ApsaraMQ for RocketMQ instance to send and receive messages.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// To send and receive messages, a client must be connected to a ApsaraMQ for RocketMQ instance by using an endpoint. You can call this operation to query the endpoints of the instance.
//
// @param request - OnsInstanceBaseInfoRequest
//
// @return OnsInstanceBaseInfoResponse
func (client *Client) OnsInstanceBaseInfo(request *OnsInstanceBaseInfoRequest) (_result *OnsInstanceBaseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsInstanceBaseInfoResponse{}
	_body, _err := client.OnsInstanceBaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// An instance is a virtual machine (VM) that can be used to store information about the topics and groups of ApsaraMQ for RocketMQ. You can call this operation when you need to create service resources for the business that you want to launch. Before you call this operation, take note of the following limits:
//
//   - A maximum of eight ApsaraMQ for RocketMQ instances can be deployed in each region.
//
//   - This operation can be called to create only a Standard Edition instance. You can use the ApsaraMQ for RocketMQ console to create Standard Edition instances and Enterprise Platinum Edition instances. For information about how to create ApsaraMQ for RocketMQ instances, see [Manage instances](https://help.aliyun.com/document_detail/200153.html).
//
// @param request - OnsInstanceCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsInstanceCreateResponse
func (client *Client) OnsInstanceCreateWithOptions(request *OnsInstanceCreateRequest, runtime *util.RuntimeOptions) (_result *OnsInstanceCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsInstanceCreate"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsInstanceCreateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsInstanceCreateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// An instance is a virtual machine (VM) that can be used to store information about the topics and groups of ApsaraMQ for RocketMQ. You can call this operation when you need to create service resources for the business that you want to launch. Before you call this operation, take note of the following limits:
//
//   - A maximum of eight ApsaraMQ for RocketMQ instances can be deployed in each region.
//
//   - This operation can be called to create only a Standard Edition instance. You can use the ApsaraMQ for RocketMQ console to create Standard Edition instances and Enterprise Platinum Edition instances. For information about how to create ApsaraMQ for RocketMQ instances, see [Manage instances](https://help.aliyun.com/document_detail/200153.html).
//
// @param request - OnsInstanceCreateRequest
//
// @return OnsInstanceCreateResponse
func (client *Client) OnsInstanceCreate(request *OnsInstanceCreateRequest) (_result *OnsInstanceCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsInstanceCreateResponse{}
	_body, _err := client.OnsInstanceCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - You can call this operation when you need to reclaim resources. For example, after you unpublish an application, you can reclaim the resources that were used for the application. An instance can be deleted only when the instance does not contain topics and groups.
//
//   - After an instance is deleted, the instance cannot be restored. Exercise caution when you call this operation.
//
// @param request - OnsInstanceDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsInstanceDeleteResponse
func (client *Client) OnsInstanceDeleteWithOptions(request *OnsInstanceDeleteRequest, runtime *util.RuntimeOptions) (_result *OnsInstanceDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsInstanceDelete"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsInstanceDeleteResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsInstanceDeleteResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - You can call this operation when you need to reclaim resources. For example, after you unpublish an application, you can reclaim the resources that were used for the application. An instance can be deleted only when the instance does not contain topics and groups.
//
//   - After an instance is deleted, the instance cannot be restored. Exercise caution when you call this operation.
//
// @param request - OnsInstanceDeleteRequest
//
// @return OnsInstanceDeleteResponse
func (client *Client) OnsInstanceDelete(request *OnsInstanceDeleteRequest) (_result *OnsInstanceDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsInstanceDeleteResponse{}
	_body, _err := client.OnsInstanceDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all Message Queue for Apache RocketMQ instances in a specified region within the current account.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - OnsInstanceInServiceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsInstanceInServiceListResponse
func (client *Client) OnsInstanceInServiceListWithOptions(request *OnsInstanceInServiceListRequest, runtime *util.RuntimeOptions) (_result *OnsInstanceInServiceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedResourceInfo)) {
		query["NeedResourceInfo"] = request.NeedResourceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsInstanceInServiceList"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsInstanceInServiceListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsInstanceInServiceListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries all Message Queue for Apache RocketMQ instances in a specified region within the current account.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - OnsInstanceInServiceListRequest
//
// @return OnsInstanceInServiceListResponse
func (client *Client) OnsInstanceInServiceList(request *OnsInstanceInServiceListRequest) (_result *OnsInstanceInServiceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsInstanceInServiceListResponse{}
	_body, _err := client.OnsInstanceInServiceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name and description of a ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// A maximum of eight ApsaraMQ for RocketMQ instances can be deployed in each region.
//
// @param request - OnsInstanceUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsInstanceUpdateResponse
func (client *Client) OnsInstanceUpdateWithOptions(request *OnsInstanceUpdateRequest, runtime *util.RuntimeOptions) (_result *OnsInstanceUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsInstanceUpdate"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsInstanceUpdateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsInstanceUpdateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates the name and description of a ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// A maximum of eight ApsaraMQ for RocketMQ instances can be deployed in each region.
//
// @param request - OnsInstanceUpdateRequest
//
// @return OnsInstanceUpdateResponse
func (client *Client) OnsInstanceUpdate(request *OnsInstanceUpdateRequest) (_result *OnsInstanceUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsInstanceUpdateResponse{}
	_body, _err := client.OnsInstanceUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a message.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - OnsMessageDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsMessageDetailResponse
func (client *Client) OnsMessageDetailWithOptions(request *OnsMessageDetailRequest, runtime *util.RuntimeOptions) (_result *OnsMessageDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsMessageDetail"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsMessageDetailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsMessageDetailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of a message.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - OnsMessageDetailRequest
//
// @return OnsMessageDetailResponse
func (client *Client) OnsMessageDetail(request *OnsMessageDetailRequest) (_result *OnsMessageDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMessageDetailResponse{}
	_body, _err := client.OnsMessageDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries messages by using a specified topic and message key.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - This operation uses the fuzzy match method to query messages based on a specified message key. The same message key may be used by multiple messages. Therefore, the returned result may contain information about multiple messages.
//
//   - This operation can be used in scenarios in which you cannot obtain the IDs of the messages that you want to query. You can perform the following steps to query the information about messages:
//
//     1.  Call this operation to query message IDs.
//
//     2.  Call the **OnsMessageGetByMsgId*	- operation that uses the exact match method to query the details of a specified message. For more information about the **OnsMessageGetByMsgId*	- operation, see [OnsMessageGetByMsgId](https://help.aliyun.com/document_detail/29607.html).
//
// @param request - OnsMessageGetByKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsMessageGetByKeyResponse
func (client *Client) OnsMessageGetByKeyWithOptions(request *OnsMessageGetByKeyRequest, runtime *util.RuntimeOptions) (_result *OnsMessageGetByKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsMessageGetByKey"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsMessageGetByKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsMessageGetByKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries messages by using a specified topic and message key.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - This operation uses the fuzzy match method to query messages based on a specified message key. The same message key may be used by multiple messages. Therefore, the returned result may contain information about multiple messages.
//
//   - This operation can be used in scenarios in which you cannot obtain the IDs of the messages that you want to query. You can perform the following steps to query the information about messages:
//
//     1.  Call this operation to query message IDs.
//
//     2.  Call the **OnsMessageGetByMsgId*	- operation that uses the exact match method to query the details of a specified message. For more information about the **OnsMessageGetByMsgId*	- operation, see [OnsMessageGetByMsgId](https://help.aliyun.com/document_detail/29607.html).
//
// @param request - OnsMessageGetByKeyRequest
//
// @return OnsMessageGetByKeyResponse
func (client *Client) OnsMessageGetByKey(request *OnsMessageGetByKeyRequest) (_result *OnsMessageGetByKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMessageGetByKeyResponse{}
	_body, _err := client.OnsMessageGetByKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a message by specifying the message ID and determines whether the message has been consumed.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - If a message is not consumed as expected, you can call this operation to query the information about the message for troubleshooting.
//
//   - This operation uses the exact match method to query a message based on the message ID. You can obtain the message ID from the SendResult parameter that is returned after the message is sent. You must store the returned information after each message is sent. The queried information about a message includes the point in time when the message was sent, the broker on which the message is stored, and the attributes of the message such as the message key and tag.
//
// @param request - OnsMessageGetByMsgIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsMessageGetByMsgIdResponse
func (client *Client) OnsMessageGetByMsgIdWithOptions(request *OnsMessageGetByMsgIdRequest, runtime *util.RuntimeOptions) (_result *OnsMessageGetByMsgIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		query["MsgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsMessageGetByMsgId"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsMessageGetByMsgIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsMessageGetByMsgIdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about a message by specifying the message ID and determines whether the message has been consumed.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - If a message is not consumed as expected, you can call this operation to query the information about the message for troubleshooting.
//
//   - This operation uses the exact match method to query a message based on the message ID. You can obtain the message ID from the SendResult parameter that is returned after the message is sent. You must store the returned information after each message is sent. The queried information about a message includes the point in time when the message was sent, the broker on which the message is stored, and the attributes of the message such as the message key and tag.
//
// @param request - OnsMessageGetByMsgIdRequest
//
// @return OnsMessageGetByMsgIdResponse
func (client *Client) OnsMessageGetByMsgId(request *OnsMessageGetByMsgIdRequest) (_result *OnsMessageGetByMsgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMessageGetByMsgIdResponse{}
	_body, _err := client.OnsMessageGetByMsgIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all messages that are stored in a specified topic within a specified time range by page.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - If you do not know the ID or key of a message that you want to query, you can call this operation to query all messages that are stored in the topic within a specified time range. The results are displayed by page.
//
//   - We recommend that you specify a short time range to query messages. If you specify a long time range, a large number of messages are returned. In this case, you cannot find the message that you want to query in an efficient manner. You can perform the following steps to query messages:
//
//     1.  Perform a paged query by specifying the topic, start time, end time, and number of entries to return on each page. If the topic contains messages, the information about the messages on the first page, total number of pages, and task ID are returned by default.
//
//     2.  Specify the task ID and a page number to call this operation again to query the messages on the specified page. The BeginTime, EndTime, and PageSize parameters do not take effect. By default, the system uses the values of these parameters that you specified in the request when you created the specified query task.
//
// @param request - OnsMessagePageQueryByTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsMessagePageQueryByTopicResponse
func (client *Client) OnsMessagePageQueryByTopicWithOptions(request *OnsMessagePageQueryByTopicRequest, runtime *util.RuntimeOptions) (_result *OnsMessagePageQueryByTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsMessagePageQueryByTopic"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsMessagePageQueryByTopicResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsMessagePageQueryByTopicResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries all messages that are stored in a specified topic within a specified time range by page.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - If you do not know the ID or key of a message that you want to query, you can call this operation to query all messages that are stored in the topic within a specified time range. The results are displayed by page.
//
//   - We recommend that you specify a short time range to query messages. If you specify a long time range, a large number of messages are returned. In this case, you cannot find the message that you want to query in an efficient manner. You can perform the following steps to query messages:
//
//     1.  Perform a paged query by specifying the topic, start time, end time, and number of entries to return on each page. If the topic contains messages, the information about the messages on the first page, total number of pages, and task ID are returned by default.
//
//     2.  Specify the task ID and a page number to call this operation again to query the messages on the specified page. The BeginTime, EndTime, and PageSize parameters do not take effect. By default, the system uses the values of these parameters that you specified in the request when you created the specified query task.
//
// @param request - OnsMessagePageQueryByTopicRequest
//
// @return OnsMessagePageQueryByTopicResponse
func (client *Client) OnsMessagePageQueryByTopic(request *OnsMessagePageQueryByTopicRequest) (_result *OnsMessagePageQueryByTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMessagePageQueryByTopicResponse{}
	_body, _err := client.OnsMessagePageQueryByTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pushes a message to a specified consumer.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// This operation can be used to check whether messages in a specified topic can be consumed by consumers in a specified consumer group. This operation obtains the body of the message that is specified by the MsgId parameter, re-encapsulates the message body to produce a new message, and then pushes the new message to a specified consumer. The content of the message that is sent to the consumer is the same as the content of the original message. They are not the same message because they use different message IDs.
//
// @param request - OnsMessagePushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsMessagePushResponse
func (client *Client) OnsMessagePushWithOptions(request *OnsMessagePushRequest, runtime *util.RuntimeOptions) (_result *OnsMessagePushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		query["MsgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsMessagePush"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsMessagePushResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsMessagePushResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Pushes a message to a specified consumer.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// This operation can be used to check whether messages in a specified topic can be consumed by consumers in a specified consumer group. This operation obtains the body of the message that is specified by the MsgId parameter, re-encapsulates the message body to produce a new message, and then pushes the new message to a specified consumer. The content of the message that is sent to the consumer is the same as the content of the original message. They are not the same message because they use different message IDs.
//
// @param request - OnsMessagePushRequest
//
// @return OnsMessagePushResponse
func (client *Client) OnsMessagePush(request *OnsMessagePushRequest) (_result *OnsMessagePushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMessagePushResponse{}
	_body, _err := client.OnsMessagePushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the consumption status of a message by using the message ID.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - You can call this operation to check whether a specified message is consumed. If the message is not consumed, you can troubleshoot the issue based on the returned information.
//
//   - This operation queries information based on the built-in offset mechanism of ApsaraMQ for RocketMQ. In most cases, the results are correct. If you have reset the consumer offset or cleared accumulated messages, the results may not be correct.
//
// @param request - OnsMessageTraceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsMessageTraceResponse
func (client *Client) OnsMessageTraceWithOptions(request *OnsMessageTraceRequest, runtime *util.RuntimeOptions) (_result *OnsMessageTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		query["MsgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsMessageTrace"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsMessageTraceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsMessageTraceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the consumption status of a message by using the message ID.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - You can call this operation to check whether a specified message is consumed. If the message is not consumed, you can troubleshoot the issue based on the returned information.
//
//   - This operation queries information based on the built-in offset mechanism of ApsaraMQ for RocketMQ. In most cases, the results are correct. If you have reset the consumer offset or cleared accumulated messages, the results may not be correct.
//
// @param request - OnsMessageTraceRequest
//
// @return OnsMessageTraceResponse
func (client *Client) OnsMessageTrace(request *OnsMessageTraceRequest) (_result *OnsMessageTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMessageTraceResponse{}
	_body, _err := client.OnsMessageTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries regions where ApsaraMQ for RocketMQ is available.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// When you use an SDK to access and manage a ApsaraMQ for RocketMQ instance, you must sequentially specify the information about two regions. You can query the information about the second region by calling the OnsRegionList operation. You must apply for a public endpoint in the following scenarios:
//
//   - Connect your application to ApsaraMQ for RocketMQ: Select the nearest API gateway endpoint based on the region where your application is deployed, and enter the corresponding **region ID**. The **regionId*	- is used to access Alibaba Cloud API Gateway because ApsaraMQ for RocketMQ instances provide API services by using the OpenAPI Explorer platform, which is also called POP.
//
//   - Access a region to manage its resources: Specify a region where you want to manage ApsaraMQ for RocketMQ resources and enter the region ID. You can call the **OnsRegionList*	- operation to query a region ID.
//
// @param request - OnsRegionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsRegionListResponse
func (client *Client) OnsRegionListWithOptions(runtime *util.RuntimeOptions) (_result *OnsRegionListResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("OnsRegionList"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsRegionListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsRegionListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries regions where ApsaraMQ for RocketMQ is available.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// When you use an SDK to access and manage a ApsaraMQ for RocketMQ instance, you must sequentially specify the information about two regions. You can query the information about the second region by calling the OnsRegionList operation. You must apply for a public endpoint in the following scenarios:
//
//   - Connect your application to ApsaraMQ for RocketMQ: Select the nearest API gateway endpoint based on the region where your application is deployed, and enter the corresponding **region ID**. The **regionId*	- is used to access Alibaba Cloud API Gateway because ApsaraMQ for RocketMQ instances provide API services by using the OpenAPI Explorer platform, which is also called POP.
//
//   - Access a region to manage its resources: Specify a region where you want to manage ApsaraMQ for RocketMQ resources and enter the region ID. You can call the **OnsRegionList*	- operation to query a region ID.
//
// @return OnsRegionListResponse
func (client *Client) OnsRegionList() (_result *OnsRegionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsRegionListResponse{}
	_body, _err := client.OnsRegionListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a topic.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// When you want to release a new application or expand your business, you can call this operation to create a topic based on your business requirements.
//
// @param request - OnsTopicCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTopicCreateResponse
func (client *Client) OnsTopicCreateWithOptions(request *OnsTopicCreateRequest, runtime *util.RuntimeOptions) (_result *OnsTopicCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTopicCreate"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTopicCreateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTopicCreateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a topic.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// When you want to release a new application or expand your business, you can call this operation to create a topic based on your business requirements.
//
// @param request - OnsTopicCreateRequest
//
// @return OnsTopicCreateResponse
func (client *Client) OnsTopicCreate(request *OnsTopicCreateRequest) (_result *OnsTopicCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTopicCreateResponse{}
	_body, _err := client.OnsTopicCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a topic.
//
// Description:
//
// >  API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur. - After you delete the topic, the publishing and subscription relationships that are constructed based on the topic are cleared. Exercise caution when you call this operation.
//
// You can call this operation to delete a topic when you need to reclaim the resources from the topic. For example, after an application is brought offline, you can delete the topics that are used for the application. After you delete a topic, the backend of ApsaraMQ for RocketMQ reclaims the resources from the topic. The system requires a long period of time to reclaim the resources. After you delete a topic, we recommend that you do not create a topic that uses the same name as the deleted topic within a short period of time. If the system fails to delete the specified topic, troubleshoot the issue based on the error code.
//
// @param request - OnsTopicDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTopicDeleteResponse
func (client *Client) OnsTopicDeleteWithOptions(request *OnsTopicDeleteRequest, runtime *util.RuntimeOptions) (_result *OnsTopicDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTopicDelete"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTopicDeleteResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTopicDeleteResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a topic.
//
// Description:
//
// >  API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur. - After you delete the topic, the publishing and subscription relationships that are constructed based on the topic are cleared. Exercise caution when you call this operation.
//
// You can call this operation to delete a topic when you need to reclaim the resources from the topic. For example, after an application is brought offline, you can delete the topics that are used for the application. After you delete a topic, the backend of ApsaraMQ for RocketMQ reclaims the resources from the topic. The system requires a long period of time to reclaim the resources. After you delete a topic, we recommend that you do not create a topic that uses the same name as the deleted topic within a short period of time. If the system fails to delete the specified topic, troubleshoot the issue based on the error code.
//
// @param request - OnsTopicDeleteRequest
//
// @return OnsTopicDeleteResponse
func (client *Client) OnsTopicDelete(request *OnsTopicDeleteRequest) (_result *OnsTopicDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTopicDeleteResponse{}
	_body, _err := client.OnsTopicDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about topics that belong to the current account.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// This operation returns the basic information about topics and does not return the details of topics.
//
// @param request - OnsTopicListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTopicListResponse
func (client *Client) OnsTopicListWithOptions(request *OnsTopicListRequest, runtime *util.RuntimeOptions) (_result *OnsTopicListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTopicList"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTopicListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTopicListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about topics that belong to the current account.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// This operation returns the basic information about topics and does not return the details of topics.
//
// @param request - OnsTopicListRequest
//
// @return OnsTopicListResponse
func (client *Client) OnsTopicList(request *OnsTopicListRequest) (_result *OnsTopicListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTopicListResponse{}
	_body, _err := client.OnsTopicListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total number of messages in a topic and the status of the topic.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can determine the resource usage of a topic based on the information that is returned by this operation. The returned information includes the total number of messages in the topic and the most recent point in time when a message was published to the topic.
//
// @param request - OnsTopicStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTopicStatusResponse
func (client *Client) OnsTopicStatusWithOptions(request *OnsTopicStatusRequest, runtime *util.RuntimeOptions) (_result *OnsTopicStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTopicStatus"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTopicStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTopicStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the total number of messages in a topic and the status of the topic.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can determine the resource usage of a topic based on the information that is returned by this operation. The returned information includes the total number of messages in the topic and the most recent point in time when a message was published to the topic.
//
// @param request - OnsTopicStatusRequest
//
// @return OnsTopicStatusResponse
func (client *Client) OnsTopicStatus(request *OnsTopicStatusRequest) (_result *OnsTopicStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTopicStatusResponse{}
	_body, _err := client.OnsTopicStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the online consumer groups that subscribe to a specified topic.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to query the online consumer groups that subscribe to a specified topic. If all consumers in a group are offline, the information about the group is not returned.
//
// @param request - OnsTopicSubDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTopicSubDetailResponse
func (client *Client) OnsTopicSubDetailWithOptions(request *OnsTopicSubDetailRequest, runtime *util.RuntimeOptions) (_result *OnsTopicSubDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTopicSubDetail"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTopicSubDetailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTopicSubDetailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the online consumer groups that subscribe to a specified topic.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to query the online consumer groups that subscribe to a specified topic. If all consumers in a group are offline, the information about the group is not returned.
//
// @param request - OnsTopicSubDetailRequest
//
// @return OnsTopicSubDetailResponse
func (client *Client) OnsTopicSubDetail(request *OnsTopicSubDetailRequest) (_result *OnsTopicSubDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTopicSubDetailResponse{}
	_body, _err := client.OnsTopicSubDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI OnsTopicUpdate is deprecated
//
// Summary:
//
// Configures the read/write mode for a specified topic.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to forbid read or write operations on a specific topic.
//
// @param request - OnsTopicUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTopicUpdateResponse
// Deprecated
func (client *Client) OnsTopicUpdateWithOptions(request *OnsTopicUpdateRequest, runtime *util.RuntimeOptions) (_result *OnsTopicUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Perm)) {
		query["Perm"] = request.Perm
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTopicUpdate"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTopicUpdateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTopicUpdateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Deprecated: OpenAPI OnsTopicUpdate is deprecated
//
// Summary:
//
// Configures the read/write mode for a specified topic.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to forbid read or write operations on a specific topic.
//
// @param request - OnsTopicUpdateRequest
//
// @return OnsTopicUpdateResponse
// Deprecated
func (client *Client) OnsTopicUpdate(request *OnsTopicUpdateRequest) (_result *OnsTopicUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTopicUpdateResponse{}
	_body, _err := client.OnsTopicUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The tracing results are queried by specifying the ID of a trace query task.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - Before you call this operation to query the details of the trace of a message, you must create a task to query the trace of the message based on the message ID or message key and obtain the task ID. Then, you can call this operation to query the details of the message trace based on the task ID. You can call the [OnsTraceQueryByMsgId](https://help.aliyun.com/document_detail/445322.html) operation or the [OnsTraceQueryByMsgKey](https://help.aliyun.com/document_detail/445324.html) operation to create a task to query the trace of the message and obtain the task ID from the **QueryId*	- response parameter.
//
//   - A trace query task is time-consuming. If you call this operation to query the details immediately after you create a trace query task, the results may be empty. In this case, we recommend that you try again later.
//
// @param request - OnsTraceGetResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTraceGetResultResponse
func (client *Client) OnsTraceGetResultWithOptions(request *OnsTraceGetResultRequest, runtime *util.RuntimeOptions) (_result *OnsTraceGetResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTraceGetResult"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTraceGetResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTraceGetResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// The tracing results are queried by specifying the ID of a trace query task.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - Before you call this operation to query the details of the trace of a message, you must create a task to query the trace of the message based on the message ID or message key and obtain the task ID. Then, you can call this operation to query the details of the message trace based on the task ID. You can call the [OnsTraceQueryByMsgId](https://help.aliyun.com/document_detail/445322.html) operation or the [OnsTraceQueryByMsgKey](https://help.aliyun.com/document_detail/445324.html) operation to create a task to query the trace of the message and obtain the task ID from the **QueryId*	- response parameter.
//
//   - A trace query task is time-consuming. If you call this operation to query the details immediately after you create a trace query task, the results may be empty. In this case, we recommend that you try again later.
//
// @param request - OnsTraceGetResultRequest
//
// @return OnsTraceGetResultResponse
func (client *Client) OnsTraceGetResult(request *OnsTraceGetResultRequest) (_result *OnsTraceGetResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTraceGetResultResponse{}
	_body, _err := client.OnsTraceGetResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task to query the trace of a message based on the message ID and the name of the topic in which the message is stored. The task ID is returned.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// If you want to query the trace of a message based on the message ID, you can call this operation to create a query task. After you obtain the task ID, you can call the [OnsTraceGetResult](https://help.aliyun.com/document_detail/59832.html) operation to query the details of the message trace based on the task ID.
//
// @param request - OnsTraceQueryByMsgIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTraceQueryByMsgIdResponse
func (client *Client) OnsTraceQueryByMsgIdWithOptions(request *OnsTraceQueryByMsgIdRequest, runtime *util.RuntimeOptions) (_result *OnsTraceQueryByMsgIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		query["MsgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTraceQueryByMsgId"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTraceQueryByMsgIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTraceQueryByMsgIdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a task to query the trace of a message based on the message ID and the name of the topic in which the message is stored. The task ID is returned.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// If you want to query the trace of a message based on the message ID, you can call this operation to create a query task. After you obtain the task ID, you can call the [OnsTraceGetResult](https://help.aliyun.com/document_detail/59832.html) operation to query the details of the message trace based on the task ID.
//
// @param request - OnsTraceQueryByMsgIdRequest
//
// @return OnsTraceQueryByMsgIdResponse
func (client *Client) OnsTraceQueryByMsgId(request *OnsTraceQueryByMsgIdRequest) (_result *OnsTraceQueryByMsgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTraceQueryByMsgIdResponse{}
	_body, _err := client.OnsTraceQueryByMsgIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a trace query task based on the topic name and message key and obtains the ID of the query task.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// If you want to query the trace of a message based on the message key that you obtained, you can call this operation to create a query task. After you obtain the task ID, you can call the OnsTraceGetResult operation to query the details of the message trace based on the task ID.
//
// @param request - OnsTraceQueryByMsgKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTraceQueryByMsgKeyResponse
func (client *Client) OnsTraceQueryByMsgKeyWithOptions(request *OnsTraceQueryByMsgKeyRequest, runtime *util.RuntimeOptions) (_result *OnsTraceQueryByMsgKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgKey)) {
		query["MsgKey"] = request.MsgKey
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTraceQueryByMsgKey"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTraceQueryByMsgKeyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTraceQueryByMsgKeyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a trace query task based on the topic name and message key and obtains the ID of the query task.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// If you want to query the trace of a message based on the message key that you obtained, you can call this operation to create a query task. After you obtain the task ID, you can call the OnsTraceGetResult operation to query the details of the message trace based on the task ID.
//
// @param request - OnsTraceQueryByMsgKeyRequest
//
// @return OnsTraceQueryByMsgKeyResponse
func (client *Client) OnsTraceQueryByMsgKey(request *OnsTraceQueryByMsgKeyRequest) (_result *OnsTraceQueryByMsgKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTraceQueryByMsgKeyResponse{}
	_body, _err := client.OnsTraceQueryByMsgKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics about messages that are consumed by a consumer group within a specific period of time.
//
// Description:
//
// >  API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to query the following statistics that are collected in a production environment:
//
//   - The number of messages that are consumed during each sampling period
//
//   - The transactions per second (TPS) for message consumption during each sampling period
//
// If your application consumes a small number of messages and does not consume messages at specific intervals, we recommend that you query the number of messages that are consumed during each sampling period because the statistics of TPS may not show a clear change trend.
//
// @param request - OnsTrendGroupOutputTpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTrendGroupOutputTpsResponse
func (client *Client) OnsTrendGroupOutputTpsWithOptions(request *OnsTrendGroupOutputTpsRequest, runtime *util.RuntimeOptions) (_result *OnsTrendGroupOutputTpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTrendGroupOutputTps"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTrendGroupOutputTpsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTrendGroupOutputTpsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the statistics about messages that are consumed by a consumer group within a specific period of time.
//
// Description:
//
// >  API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to query the following statistics that are collected in a production environment:
//
//   - The number of messages that are consumed during each sampling period
//
//   - The transactions per second (TPS) for message consumption during each sampling period
//
// If your application consumes a small number of messages and does not consume messages at specific intervals, we recommend that you query the number of messages that are consumed during each sampling period because the statistics of TPS may not show a clear change trend.
//
// @param request - OnsTrendGroupOutputTpsRequest
//
// @return OnsTrendGroupOutputTpsResponse
func (client *Client) OnsTrendGroupOutputTps(request *OnsTrendGroupOutputTpsRequest) (_result *OnsTrendGroupOutputTpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTrendGroupOutputTpsResponse{}
	_body, _err := client.OnsTrendGroupOutputTpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics about messages that are published to a topic within a specific period of time.
//
// Description:
//
// >  API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to query the statistics of messages that are published to a specific topic in a production environment. You can query the number of messages that are published to the topic or the transactions per second (TPS) for message publishing within a specified time range based on your business requirements.
//
// If your application publishes a small number of messages and does not publish messages at specific intervals, we recommend that you query the number of messages that are published to the topic during each sampling period because the statistics of TPS may not show a clear change trend.
//
// @param request - OnsTrendTopicInputTpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnsTrendTopicInputTpsResponse
func (client *Client) OnsTrendTopicInputTpsWithOptions(request *OnsTrendTopicInputTpsRequest, runtime *util.RuntimeOptions) (_result *OnsTrendTopicInputTpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OnsTrendTopicInputTps"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OnsTrendTopicInputTpsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OnsTrendTopicInputTpsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the statistics about messages that are published to a topic within a specific period of time.
//
// Description:
//
// >  API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to query the statistics of messages that are published to a specific topic in a production environment. You can query the number of messages that are published to the topic or the transactions per second (TPS) for message publishing within a specified time range based on your business requirements.
//
// If your application publishes a small number of messages and does not publish messages at specific intervals, we recommend that you query the number of messages that are published to the topic during each sampling period because the statistics of TPS may not show a clear change trend.
//
// @param request - OnsTrendTopicInputTpsRequest
//
// @return OnsTrendTopicInputTpsResponse
func (client *Client) OnsTrendTopicInputTps(request *OnsTrendTopicInputTpsRequest) (_result *OnsTrendTopicInputTpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsTrendTopicInputTpsResponse{}
	_body, _err := client.OnsTrendTopicInputTpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates ApsaraMQ for RocketMQ.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation the first time you use ApsaraMQ for RocketMQ. You can use ApsaraMQ for RocketMQ only after the service is activated.
//
// The ApsaraMQ for RocketMQ service can be activated only in the China (Hangzhou) region. Service activation is not billed.
//
// @param request - OpenOnsServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenOnsServiceResponse
func (client *Client) OpenOnsServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenOnsServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("OpenOnsService"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OpenOnsServiceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OpenOnsServiceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Activates ApsaraMQ for RocketMQ.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation the first time you use ApsaraMQ for RocketMQ. You can use ApsaraMQ for RocketMQ only after the service is activated.
//
// The ApsaraMQ for RocketMQ service can be activated only in the China (Hangzhou) region. Service activation is not billed.
//
// @return OpenOnsServiceResponse
func (client *Client) OpenOnsService() (_result *OpenOnsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenOnsServiceResponse{}
	_body, _err := client.OpenOnsServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches tags to resources.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to attach tags to a source. You can use tags to classify resources in ApsaraMQ for RocketMQ. This can help you aggregate and search resources in an efficient manner.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Attaches tags to resources.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// You can call this operation to attach tags to a source. You can use tags to classify resources in ApsaraMQ for RocketMQ. This can help you aggregate and search resources in an efficient manner.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches and removes tags from a specific resource.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UntagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UntagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Detaches and removes tags from a specific resource.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

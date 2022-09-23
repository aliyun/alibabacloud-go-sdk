// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListTagResourcesRequest struct {
	InstanceId   *string                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Detail     *bool   `json:"Detail,omitempty" xml:"Detail,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	Data      *OnsConsumerAccumulateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ConsumeTps        *float32                                                `json:"ConsumeTps,omitempty" xml:"ConsumeTps,omitempty"`
	DelayTime         *int64                                                  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DetailInTopicList *OnsConsumerAccumulateResponseBodyDataDetailInTopicList `json:"DetailInTopicList,omitempty" xml:"DetailInTopicList,omitempty" type:"Struct"`
	LastTimestamp     *int64                                                  `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Online            *bool                                                   `json:"Online,omitempty" xml:"Online,omitempty"`
	TotalDiff         *int64                                                  `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
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
	DelayTime     *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	LastTimestamp *int64  `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TotalDiff     *int64  `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsConsumerAccumulateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	Data      *OnsConsumerGetConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ConnectionList *OnsConsumerGetConnectionResponseBodyDataConnectionList `json:"ConnectionList,omitempty" xml:"ConnectionList,omitempty" type:"Struct"`
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
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsConsumerGetConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResetTimestamp *int64  `json:"ResetTimestamp,omitempty" xml:"ResetTimestamp,omitempty"`
	Topic          *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsConsumerResetOffsetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Detail     *bool   `json:"Detail,omitempty" xml:"Detail,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NeedJstack *bool   `json:"NeedJstack,omitempty" xml:"NeedJstack,omitempty"`
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
	Data      *OnsConsumerStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ConnectionSet              *OnsConsumerStatusResponseBodyDataConnectionSet              `json:"ConnectionSet,omitempty" xml:"ConnectionSet,omitempty" type:"Struct"`
	ConsumeModel               *string                                                      `json:"ConsumeModel,omitempty" xml:"ConsumeModel,omitempty"`
	ConsumeTps                 *float32                                                     `json:"ConsumeTps,omitempty" xml:"ConsumeTps,omitempty"`
	ConsumerConnectionInfoList *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList `json:"ConsumerConnectionInfoList,omitempty" xml:"ConsumerConnectionInfoList,omitempty" type:"Struct"`
	DelayTime                  *int64                                                       `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DetailInTopicList          *OnsConsumerStatusResponseBodyDataDetailInTopicList          `json:"DetailInTopicList,omitempty" xml:"DetailInTopicList,omitempty" type:"Struct"`
	InstanceId                 *string                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LastTimestamp              *int64                                                       `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Online                     *bool                                                        `json:"Online,omitempty" xml:"Online,omitempty"`
	RebalanceOK                *bool                                                        `json:"RebalanceOK,omitempty" xml:"RebalanceOK,omitempty"`
	SubscriptionSame           *bool                                                        `json:"SubscriptionSame,omitempty" xml:"SubscriptionSame,omitempty"`
	TotalDiff                  *int64                                                       `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
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
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	RemoteIP   *string `json:"RemoteIP,omitempty" xml:"RemoteIP,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
	ClientId        *string                                                                                             `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Connection      *string                                                                                             `json:"Connection,omitempty" xml:"Connection,omitempty"`
	ConsumeModel    *string                                                                                             `json:"ConsumeModel,omitempty" xml:"ConsumeModel,omitempty"`
	ConsumeType     *string                                                                                             `json:"ConsumeType,omitempty" xml:"ConsumeType,omitempty"`
	Jstack          *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack          `json:"Jstack,omitempty" xml:"Jstack,omitempty" type:"Struct"`
	Language        *string                                                                                             `json:"Language,omitempty" xml:"Language,omitempty"`
	LastTimeStamp   *int64                                                                                              `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	RunningDataList *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList `json:"RunningDataList,omitempty" xml:"RunningDataList,omitempty" type:"Struct"`
	StartTimeStamp  *int64                                                                                              `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	SubscriptionSet *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet `json:"SubscriptionSet,omitempty" xml:"SubscriptionSet,omitempty" type:"Struct"`
	ThreadCount     *int32                                                                                              `json:"ThreadCount,omitempty" xml:"ThreadCount,omitempty"`
	Version         *string                                                                                             `json:"Version,omitempty" xml:"Version,omitempty"`
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
	Thread    *string                                                                                                          `json:"Thread,omitempty" xml:"Thread,omitempty"`
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
	FailedCountPerHour *int64   `json:"FailedCountPerHour,omitempty" xml:"FailedCountPerHour,omitempty"`
	FailedTps          *float32 `json:"FailedTps,omitempty" xml:"FailedTps,omitempty"`
	GroupId            *string  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OkTps              *float32 `json:"OkTps,omitempty" xml:"OkTps,omitempty"`
	Rt                 *float32 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Topic              *string  `json:"Topic,omitempty" xml:"Topic,omitempty"`
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

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetGroupId(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.GroupId = &v
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
	SubString  *string                                                                                                                    `json:"SubString,omitempty" xml:"SubString,omitempty"`
	SubVersion *int64                                                                                                                     `json:"SubVersion,omitempty" xml:"SubVersion,omitempty"`
	TagsSet    *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet `json:"TagsSet,omitempty" xml:"TagsSet,omitempty" type:"Struct"`
	Topic      *string                                                                                                                    `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	DelayTime     *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	LastTimestamp *int64  `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TotalDiff     *int64  `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsConsumerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Data      *OnsConsumerTimeSpanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ConsumeTimeStamp *int64  `json:"ConsumeTimeStamp,omitempty" xml:"ConsumeTimeStamp,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxTimeStamp     *int64  `json:"MaxTimeStamp,omitempty" xml:"MaxTimeStamp,omitempty"`
	MinTimeStamp     *int64  `json:"MinTimeStamp,omitempty" xml:"MinTimeStamp,omitempty"`
	Topic            *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsConsumerTimeSpanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
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
	Data      *OnsDLQMessageGetByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	BodyCRC        *int32                                            `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                           `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                            `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                           `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   *OnsDLQMessageGetByIdResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	ReconsumeTimes *int32                                            `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                           `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int32                                            `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                            `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                           `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsDLQMessageGetByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BeginTime   *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	MsgFoundDo *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo `json:"MsgFoundDo,omitempty" xml:"MsgFoundDo,omitempty" type:"Struct"`
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CurrentPage  *int64                                                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	MaxPageCount *int64                                                             `json:"MaxPageCount,omitempty" xml:"MaxPageCount,omitempty"`
	MsgFoundList *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList `json:"MsgFoundList,omitempty" xml:"MsgFoundList,omitempty" type:"Struct"`
	TaskId       *string                                                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	BodyCRC        *int32                                                                                         `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                                                                        `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                                                         `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                                                        `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	ReconsumeTimes *int32                                                                                         `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                                                                        `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int32                                                                                         `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                                                         `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                                                                        `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsDLQMessagePageQueryByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
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
	Data      *OnsDLQMessageResendByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsDLQMessageResendByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ReadEnable *bool   `json:"ReadEnable,omitempty" xml:"ReadEnable,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsGroupConsumerUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupType  *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsGroupCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsGroupDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupId    *string                   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupType  *string                   `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	InstanceId *string                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Tag        []*OnsGroupListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Data      *OnsGroupListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime        *int64                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GroupId           *string                                          `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupType         *string                                          `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	IndependentNaming *bool                                            `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	InstanceId        *string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Owner             *string                                          `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Remark            *string                                          `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Tags              *OnsGroupListResponseBodyDataSubscribeInfoDoTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UpdateTime        *int64                                           `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	Data      *OnsGroupSubDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	GroupId              *string                                                `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MessageModel         *string                                                `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	Online               *bool                                                  `json:"Online,omitempty" xml:"Online,omitempty"`
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
	SubString *string `json:"SubString,omitempty" xml:"SubString,omitempty"`
	Topic     *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsGroupSubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceBaseInfo *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo `json:"InstanceBaseInfo,omitempty" xml:"InstanceBaseInfo,omitempty" type:"Struct"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime        *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Endpoints         *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	IndependentNaming *bool                                                     `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	InstanceId        *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus    *int32                                                    `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType      *int32                                                    `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaxTps            *int64                                                    `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	ReleaseTime       *int64                                                    `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	Remark            *string                                                   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TopicCapacity     *int32                                                    `json:"TopicCapacity,omitempty" xml:"TopicCapacity,omitempty"`
	SpInstanceId      *string                                                   `json:"spInstanceId,omitempty" xml:"spInstanceId,omitempty"`
	SpInstanceType    *int32                                                    `json:"spInstanceType,omitempty" xml:"spInstanceType,omitempty"`
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
	HttpInternalEndpoint       *string `json:"HttpInternalEndpoint,omitempty" xml:"HttpInternalEndpoint,omitempty"`
	HttpInternetEndpoint       *string `json:"HttpInternetEndpoint,omitempty" xml:"HttpInternetEndpoint,omitempty"`
	HttpInternetSecureEndpoint *string `json:"HttpInternetSecureEndpoint,omitempty" xml:"HttpInternetSecureEndpoint,omitempty"`
	TcpEndpoint                *string `json:"TcpEndpoint,omitempty" xml:"TcpEndpoint,omitempty"`
	TcpInternetEndpoint        *string `json:"TcpInternetEndpoint,omitempty" xml:"TcpInternetEndpoint,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsInstanceBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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
	Data      *OnsInstanceCreateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *int32  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsInstanceCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsInstanceDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Tag []*OnsInstanceInServiceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsInstanceInServiceListRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListRequest) SetTag(v []*OnsInstanceInServiceListRequestTag) *OnsInstanceInServiceListRequest {
	s.Tag = v
	return s
}

type OnsInstanceInServiceListRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Data      *OnsInstanceInServiceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime        *int64                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IndependentNaming *bool                                                   `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	InstanceId        *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string                                                 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus    *int32                                                  `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType      *int32                                                  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	ReleaseTime       *int64                                                  `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	Tags              *OnsInstanceInServiceListResponseBodyDataInstanceVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsInstanceInServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsInstanceUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type OnsMessageGetByKeyRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Data      *OnsMessageGetByKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	BodyCRC        *int32                                                          `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                                         `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                          `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                         `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   *OnsMessageGetByKeyResponseBodyDataOnsRestMessageDoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	ReconsumeTimes *int32                                                          `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                                         `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int32                                                          `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                          `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                                         `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsMessageGetByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Data      *OnsMessageGetByMsgIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	BodyCRC        *int32                                            `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                           `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                            `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                           `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   *OnsMessageGetByMsgIdResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	ReconsumeTimes *int32                                            `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                           `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int32                                            `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                            `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                           `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsMessageGetByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BeginTime   *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Topic       *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	MsgFoundDo *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo `json:"MsgFoundDo,omitempty" xml:"MsgFoundDo,omitempty" type:"Struct"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CurrentPage  *int64                                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	MaxPageCount *int64                                                        `json:"MaxPageCount,omitempty" xml:"MaxPageCount,omitempty"`
	MsgFoundList *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList `json:"MsgFoundList,omitempty" xml:"MsgFoundList,omitempty" type:"Struct"`
	TaskId       *string                                                       `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	BodyCRC        *int32                                                                                    `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                                                                   `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	BornTimestamp  *int64                                                                                    `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	InstanceId     *string                                                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                                                   `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList   *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListOnsRestMessageDoPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	ReconsumeTimes *int32                                                                                    `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreHost      *string                                                                                   `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	StoreSize      *int32                                                                                    `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	StoreTimestamp *int64                                                                                    `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	Topic          *string                                                                                   `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsMessagePageQueryByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsMessagePushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Data      *OnsMessageTraceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TrackType     *string `json:"TrackType,omitempty" xml:"TrackType,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsMessageTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Data      *OnsRegionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	OnsRegionId *string `json:"OnsRegionId,omitempty" xml:"OnsRegionId,omitempty"`
	RegionName  *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s OnsRegionListResponseBodyDataRegionDo) String() string {
	return tea.Prettify(s)
}

func (s OnsRegionListResponseBodyDataRegionDo) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetOnsRegionId(v string) *OnsRegionListResponseBodyDataRegionDo {
	s.OnsRegionId = &v
	return s
}

func (s *OnsRegionListResponseBodyDataRegionDo) SetRegionName(v string) *OnsRegionListResponseBodyDataRegionDo {
	s.RegionName = &v
	return s
}

type OnsRegionListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MessageType *int32  `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Topic       *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTopicCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTopicDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Tag        []*OnsTopicListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Topic      *string                   `json:"Topic,omitempty" xml:"Topic,omitempty"`
	UserId     *string                   `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Data      *OnsTopicListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime        *int64                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IndependentNaming *bool                                          `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	InstanceId        *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MessageType       *int32                                         `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Owner             *string                                        `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Relation          *int32                                         `json:"Relation,omitempty" xml:"Relation,omitempty"`
	RelationName      *string                                        `json:"RelationName,omitempty" xml:"RelationName,omitempty"`
	Remark            *string                                        `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ServiceStatus     *int32                                         `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	Tags              *OnsTopicListResponseBodyDataPublishInfoDoTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Topic             *string                                        `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTopicListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Data      *OnsTopicStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	Perm          *int32 `json:"Perm,omitempty" xml:"Perm,omitempty"`
	TotalCount    *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTopicStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Data      *OnsTopicSubDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	SubscriptionDataList *OnsTopicSubDetailResponseBodyDataSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Struct"`
	Topic                *string                                                `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MessageModel *string `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	SubString    *string `json:"SubString,omitempty" xml:"SubString,omitempty"`
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

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList) SetSubString(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataListSubscriptionDataList {
	s.SubString = &v
	return s
}

type OnsTopicSubDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTopicSubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Perm       *int32  `json:"Perm,omitempty" xml:"Perm,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTopicUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s OnsTraceGetResultRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultRequest) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultRequest) SetQueryId(v string) *OnsTraceGetResultRequest {
	s.QueryId = &v
	return s
}

type OnsTraceGetResultResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime *int64                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId *string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string                                          `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	MsgKey     *string                                          `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	QueryId    *string                                          `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	Status     *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Topic      *string                                          `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TraceList  *OnsTraceGetResultResponseBodyTraceDataTraceList `json:"TraceList,omitempty" xml:"TraceList,omitempty" type:"Struct"`
	UpdateTime *int64                                           `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string                                          `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	BornHost     *string                                                           `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	CostTime     *int32                                                            `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	MsgId        *string                                                           `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	MsgKey       *string                                                           `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	PubGroupName *string                                                           `json:"PubGroupName,omitempty" xml:"PubGroupName,omitempty"`
	PubTime      *int64                                                            `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	Status       *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	SubList      *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubList `json:"SubList,omitempty" xml:"SubList,omitempty" type:"Struct"`
	Tag          *string                                                           `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Topic        *string                                                           `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	ClientList   *OnsTraceGetResultResponseBodyTraceDataTraceListTraceMapDoSubListSubMapDoClientList `json:"ClientList,omitempty" xml:"ClientList,omitempty" type:"Struct"`
	FailCount    *int32                                                                              `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	SubGroupName *string                                                                             `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	SuccessCount *int32                                                                              `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
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
	ClientHost     *string `json:"ClientHost,omitempty" xml:"ClientHost,omitempty"`
	CostTime       *int32  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	ReconsumeTimes *int32  `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubGroupName   *string `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	SubTime        *int64  `json:"SubTime,omitempty" xml:"SubTime,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTraceGetResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	QueryId   *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTraceQueryByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgKey     *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	QueryId   *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTraceQueryByMsgKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Period     *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Data      *OnsTrendGroupOutputTpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Records *OnsTrendGroupOutputTpsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	Title   *string                                        `json:"Title,omitempty" xml:"Title,omitempty"`
	XUnit   *string                                        `json:"XUnit,omitempty" xml:"XUnit,omitempty"`
	YUnit   *string                                        `json:"YUnit,omitempty" xml:"YUnit,omitempty"`
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
	X *int64   `json:"X,omitempty" xml:"X,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTrendGroupOutputTpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Period     *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Data      *OnsTrendTopicInputTpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Records *OnsTrendTopicInputTpsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	Title   *string                                       `json:"Title,omitempty" xml:"Title,omitempty"`
	XUnit   *string                                       `json:"XUnit,omitempty" xml:"XUnit,omitempty"`
	YUnit   *string                                       `json:"YUnit,omitempty" xml:"YUnit,omitempty"`
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
	X *int64   `json:"X,omitempty" xml:"X,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OnsTrendTopicInputTpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenOnsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId   *string                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsConsumerAccumulateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsConsumerGetConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsConsumerResetOffsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsConsumerStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsConsumerTimeSpanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsDLQMessageGetByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsDLQMessagePageQueryByGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsDLQMessageResendByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsGroupConsumerUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsGroupCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsGroupDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsGroupListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsGroupSubDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsInstanceBaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsInstanceCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsInstanceDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) OnsInstanceInServiceListWithOptions(request *OnsInstanceInServiceListRequest, runtime *util.RuntimeOptions) (_result *OnsInstanceInServiceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	_result = &OnsInstanceInServiceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsInstanceUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsMessageGetByKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsMessageGetByMsgIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsMessagePageQueryByTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsMessagePushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsMessageTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsRegionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsTopicCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsTopicDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsTopicListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsTopicStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsTopicSubDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsTopicUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) OnsTraceGetResultWithOptions(request *OnsTraceGetResultRequest, runtime *util.RuntimeOptions) (_result *OnsTraceGetResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
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
	_result = &OnsTraceGetResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsTraceQueryByMsgIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsTraceQueryByMsgKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsTrendGroupOutputTpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OnsTrendTopicInputTpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &OpenOnsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

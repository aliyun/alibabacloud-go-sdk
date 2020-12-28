// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListTagResourcesRequest struct {
	InstanceId   *string                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
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
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetInstanceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.InstanceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type OnsConsumerAccumulateRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Detail     *bool   `json:"Detail,omitempty" xml:"Detail,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsConsumerAccumulateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerAccumulateRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateRequest) SetGroupId(v string) *OnsConsumerAccumulateRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerAccumulateRequest) SetDetail(v bool) *OnsConsumerAccumulateRequest {
	s.Detail = &v
	return s
}

func (s *OnsConsumerAccumulateRequest) SetInstanceId(v string) *OnsConsumerAccumulateRequest {
	s.InstanceId = &v
	return s
}

type OnsConsumerAccumulateResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsConsumerAccumulateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsConsumerAccumulateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerAccumulateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponseBody) SetRequestId(v string) *OnsConsumerAccumulateResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBody) SetData(v *OnsConsumerAccumulateResponseBodyData) *OnsConsumerAccumulateResponseBody {
	s.Data = v
	return s
}

type OnsConsumerAccumulateResponseBodyData struct {
	ConsumeTps        *float32                                                  `json:"ConsumeTps,omitempty" xml:"ConsumeTps,omitempty"`
	DelayTime         *int64                                                    `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	LastTimestamp     *int64                                                    `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	TotalDiff         *int64                                                    `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
	Online            *bool                                                     `json:"Online,omitempty" xml:"Online,omitempty"`
	DetailInTopicList []*OnsConsumerAccumulateResponseBodyDataDetailInTopicList `json:"DetailInTopicList,omitempty" xml:"DetailInTopicList,omitempty" type:"Repeated"`
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

func (s *OnsConsumerAccumulateResponseBodyData) SetLastTimestamp(v int64) *OnsConsumerAccumulateResponseBodyData {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetTotalDiff(v int64) *OnsConsumerAccumulateResponseBodyData {
	s.TotalDiff = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetOnline(v bool) *OnsConsumerAccumulateResponseBodyData {
	s.Online = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetDetailInTopicList(v []*OnsConsumerAccumulateResponseBodyDataDetailInTopicList) *OnsConsumerAccumulateResponseBodyData {
	s.DetailInTopicList = v
	return s
}

type OnsConsumerAccumulateResponseBodyDataDetailInTopicList struct {
	DelayTime     *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	TotalDiff     *int64  `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
	LastTimestamp *int64  `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsConsumerAccumulateResponseBodyDataDetailInTopicList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerAccumulateResponseBodyDataDetailInTopicList) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicList) SetDelayTime(v int64) *OnsConsumerAccumulateResponseBodyDataDetailInTopicList {
	s.DelayTime = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicList) SetTotalDiff(v int64) *OnsConsumerAccumulateResponseBodyDataDetailInTopicList {
	s.TotalDiff = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicList) SetLastTimestamp(v int64) *OnsConsumerAccumulateResponseBodyDataDetailInTopicList {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicList) SetTopic(v string) *OnsConsumerAccumulateResponseBodyDataDetailInTopicList {
	s.Topic = &v
	return s
}

type OnsConsumerAccumulateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsConsumerAccumulateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsConsumerGetConnectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsConsumerGetConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBody) SetRequestId(v string) *OnsConsumerGetConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBody) SetData(v *OnsConsumerGetConnectionResponseBodyData) *OnsConsumerGetConnectionResponseBody {
	s.Data = v
	return s
}

type OnsConsumerGetConnectionResponseBodyData struct {
	ConnectionList []*OnsConsumerGetConnectionResponseBodyDataConnectionList `json:"ConnectionList,omitempty" xml:"ConnectionList,omitempty" type:"Repeated"`
}

func (s OnsConsumerGetConnectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBodyData) SetConnectionList(v []*OnsConsumerGetConnectionResponseBodyDataConnectionList) *OnsConsumerGetConnectionResponseBodyData {
	s.ConnectionList = v
	return s
}

type OnsConsumerGetConnectionResponseBodyDataConnectionList struct {
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s OnsConsumerGetConnectionResponseBodyDataConnectionList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerGetConnectionResponseBodyDataConnectionList) GoString() string {
	return s.String()
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionList) SetVersion(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionList {
	s.Version = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionList) SetClientAddr(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionList {
	s.ClientAddr = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionList) SetLanguage(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionList {
	s.Language = &v
	return s
}

func (s *OnsConsumerGetConnectionResponseBodyDataConnectionList) SetClientId(v string) *OnsConsumerGetConnectionResponseBodyDataConnectionList {
	s.ClientId = &v
	return s
}

type OnsConsumerGetConnectionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsConsumerGetConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsConsumerGetConnectionResponse) SetBody(v *OnsConsumerGetConnectionResponseBody) *OnsConsumerGetConnectionResponse {
	s.Body = v
	return s
}

type OnsConsumerResetOffsetRequest struct {
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Topic          *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	ResetTimestamp *int64  `json:"ResetTimestamp,omitempty" xml:"ResetTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *OnsConsumerResetOffsetRequest) SetTopic(v string) *OnsConsumerResetOffsetRequest {
	s.Topic = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetType(v int32) *OnsConsumerResetOffsetRequest {
	s.Type = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetResetTimestamp(v int64) *OnsConsumerResetOffsetRequest {
	s.ResetTimestamp = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetInstanceId(v string) *OnsConsumerResetOffsetRequest {
	s.InstanceId = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsConsumerResetOffsetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsConsumerResetOffsetResponse) SetBody(v *OnsConsumerResetOffsetResponseBody) *OnsConsumerResetOffsetResponse {
	s.Body = v
	return s
}

type OnsConsumerStatusRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Detail     *bool   `json:"Detail,omitempty" xml:"Detail,omitempty"`
	NeedJstack *bool   `json:"NeedJstack,omitempty" xml:"NeedJstack,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsConsumerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusRequest) SetGroupId(v string) *OnsConsumerStatusRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerStatusRequest) SetDetail(v bool) *OnsConsumerStatusRequest {
	s.Detail = &v
	return s
}

func (s *OnsConsumerStatusRequest) SetNeedJstack(v bool) *OnsConsumerStatusRequest {
	s.NeedJstack = &v
	return s
}

func (s *OnsConsumerStatusRequest) SetInstanceId(v string) *OnsConsumerStatusRequest {
	s.InstanceId = &v
	return s
}

type OnsConsumerStatusResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsConsumerStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsConsumerStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBody) SetRequestId(v string) *OnsConsumerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsConsumerStatusResponseBody) SetData(v *OnsConsumerStatusResponseBodyData) *OnsConsumerStatusResponseBody {
	s.Data = v
	return s
}

type OnsConsumerStatusResponseBodyData struct {
	ConsumeTps                 *float32                                                       `json:"ConsumeTps,omitempty" xml:"ConsumeTps,omitempty"`
	ConsumeModel               *string                                                        `json:"ConsumeModel,omitempty" xml:"ConsumeModel,omitempty"`
	ConnectionSet              []*OnsConsumerStatusResponseBodyDataConnectionSet              `json:"ConnectionSet,omitempty" xml:"ConnectionSet,omitempty" type:"Repeated"`
	TotalDiff                  *int64                                                         `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
	ConsumerConnectionInfoList []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList `json:"ConsumerConnectionInfoList,omitempty" xml:"ConsumerConnectionInfoList,omitempty" type:"Repeated"`
	InstanceId                 *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DetailInTopicList          []*OnsConsumerStatusResponseBodyDataDetailInTopicList          `json:"DetailInTopicList,omitempty" xml:"DetailInTopicList,omitempty" type:"Repeated"`
	SubscriptionSame           *bool                                                          `json:"SubscriptionSame,omitempty" xml:"SubscriptionSame,omitempty"`
	DelayTime                  *int64                                                         `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	LastTimestamp              *int64                                                         `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Online                     *bool                                                          `json:"Online,omitempty" xml:"Online,omitempty"`
	RebalanceOK                *bool                                                          `json:"RebalanceOK,omitempty" xml:"RebalanceOK,omitempty"`
}

func (s OnsConsumerStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyData) SetConsumeTps(v float32) *OnsConsumerStatusResponseBodyData {
	s.ConsumeTps = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetConsumeModel(v string) *OnsConsumerStatusResponseBodyData {
	s.ConsumeModel = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetConnectionSet(v []*OnsConsumerStatusResponseBodyDataConnectionSet) *OnsConsumerStatusResponseBodyData {
	s.ConnectionSet = v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetTotalDiff(v int64) *OnsConsumerStatusResponseBodyData {
	s.TotalDiff = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetConsumerConnectionInfoList(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) *OnsConsumerStatusResponseBodyData {
	s.ConsumerConnectionInfoList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetInstanceId(v string) *OnsConsumerStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetDetailInTopicList(v []*OnsConsumerStatusResponseBodyDataDetailInTopicList) *OnsConsumerStatusResponseBodyData {
	s.DetailInTopicList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetSubscriptionSame(v bool) *OnsConsumerStatusResponseBodyData {
	s.SubscriptionSame = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetDelayTime(v int64) *OnsConsumerStatusResponseBodyData {
	s.DelayTime = &v
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

type OnsConsumerStatusResponseBodyDataConnectionSet struct {
	RemoteIP   *string `json:"RemoteIP,omitempty" xml:"RemoteIP,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConnectionSet) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConnectionSet) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSet) SetRemoteIP(v string) *OnsConsumerStatusResponseBodyDataConnectionSet {
	s.RemoteIP = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSet) SetVersion(v string) *OnsConsumerStatusResponseBodyDataConnectionSet {
	s.Version = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSet) SetClientAddr(v string) *OnsConsumerStatusResponseBodyDataConnectionSet {
	s.ClientAddr = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSet) SetLanguage(v string) *OnsConsumerStatusResponseBodyDataConnectionSet {
	s.Language = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSet) SetClientId(v string) *OnsConsumerStatusResponseBodyDataConnectionSet {
	s.ClientId = &v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList struct {
	ConsumeModel    *string                                                                       `json:"ConsumeModel,omitempty" xml:"ConsumeModel,omitempty"`
	RunningDataList []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList `json:"RunningDataList,omitempty" xml:"RunningDataList,omitempty" type:"Repeated"`
	SubscriptionSet []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet `json:"SubscriptionSet,omitempty" xml:"SubscriptionSet,omitempty" type:"Repeated"`
	Jstack          []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack          `json:"Jstack,omitempty" xml:"Jstack,omitempty" type:"Repeated"`
	LastTimeStamp   *int64                                                                        `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	StartTimeStamp  *int64                                                                        `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	Language        *string                                                                       `json:"Language,omitempty" xml:"Language,omitempty"`
	ClientId        *string                                                                       `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Connection      *string                                                                       `json:"Connection,omitempty" xml:"Connection,omitempty"`
	Version         *string                                                                       `json:"Version,omitempty" xml:"Version,omitempty"`
	ConsumeType     *string                                                                       `json:"ConsumeType,omitempty" xml:"ConsumeType,omitempty"`
	ThreadCount     *int32                                                                        `json:"ThreadCount,omitempty" xml:"ThreadCount,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetConsumeModel(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.ConsumeModel = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetRunningDataList(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.RunningDataList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetSubscriptionSet(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.SubscriptionSet = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetJstack(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.Jstack = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetLastTimeStamp(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.LastTimeStamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetStartTimeStamp(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.StartTimeStamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetLanguage(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.Language = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetClientId(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.ClientId = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetConnection(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.Connection = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetVersion(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.Version = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetConsumeType(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.ConsumeType = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetThreadCount(v int32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.ThreadCount = &v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList struct {
	GroupId            *string  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Rt                 *float32 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Topic              *string  `json:"Topic,omitempty" xml:"Topic,omitempty"`
	FailedCountPerHour *int64   `json:"FailedCountPerHour,omitempty" xml:"FailedCountPerHour,omitempty"`
	OkTps              *float32 `json:"OkTps,omitempty" xml:"OkTps,omitempty"`
	FailedTps          *float32 `json:"FailedTps,omitempty" xml:"FailedTps,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetGroupId(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetRt(v float32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.Rt = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetTopic(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.Topic = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetFailedCountPerHour(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.FailedCountPerHour = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetOkTps(v float32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.OkTps = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList) SetFailedTps(v float32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListRunningDataList {
	s.FailedTps = &v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet struct {
	SubString  *string   `json:"SubString,omitempty" xml:"SubString,omitempty"`
	SubVersion *int64    `json:"SubVersion,omitempty" xml:"SubVersion,omitempty"`
	Topic      *string   `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TagsSet    []*string `json:"TagsSet,omitempty" xml:"TagsSet,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) SetSubString(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet {
	s.SubString = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) SetSubVersion(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet {
	s.SubVersion = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) SetTopic(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet {
	s.Topic = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet) SetTagsSet(v []*string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListSubscriptionSet {
	s.TagsSet = v
	return s
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack struct {
	TrackList []*string `json:"TrackList,omitempty" xml:"TrackList,omitempty" type:"Repeated"`
	Thread    *string   `json:"Thread,omitempty" xml:"Thread,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack) SetTrackList(v []*string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack {
	s.TrackList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack) SetThread(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListJstack {
	s.Thread = &v
	return s
}

type OnsConsumerStatusResponseBodyDataDetailInTopicList struct {
	DelayTime     *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	TotalDiff     *int64  `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
	LastTimestamp *int64  `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataDetailInTopicList) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataDetailInTopicList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicList) SetDelayTime(v int64) *OnsConsumerStatusResponseBodyDataDetailInTopicList {
	s.DelayTime = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicList) SetTotalDiff(v int64) *OnsConsumerStatusResponseBodyDataDetailInTopicList {
	s.TotalDiff = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicList) SetLastTimestamp(v int64) *OnsConsumerStatusResponseBodyDataDetailInTopicList {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicList) SetTopic(v string) *OnsConsumerStatusResponseBodyDataDetailInTopicList {
	s.Topic = &v
	return s
}

type OnsConsumerStatusResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsConsumerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsConsumerStatusResponse) SetBody(v *OnsConsumerStatusResponseBody) *OnsConsumerStatusResponse {
	s.Body = v
	return s
}

type OnsConsumerTimeSpanRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *OnsConsumerTimeSpanRequest) SetTopic(v string) *OnsConsumerTimeSpanRequest {
	s.Topic = &v
	return s
}

func (s *OnsConsumerTimeSpanRequest) SetInstanceId(v string) *OnsConsumerTimeSpanRequest {
	s.InstanceId = &v
	return s
}

type OnsConsumerTimeSpanResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsConsumerTimeSpanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsConsumerTimeSpanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerTimeSpanResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerTimeSpanResponseBody) SetRequestId(v string) *OnsConsumerTimeSpanResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBody) SetData(v *OnsConsumerTimeSpanResponseBodyData) *OnsConsumerTimeSpanResponseBody {
	s.Data = v
	return s
}

type OnsConsumerTimeSpanResponseBodyData struct {
	MaxTimeStamp     *int64  `json:"MaxTimeStamp,omitempty" xml:"MaxTimeStamp,omitempty"`
	ConsumeTimeStamp *int64  `json:"ConsumeTimeStamp,omitempty" xml:"ConsumeTimeStamp,omitempty"`
	Topic            *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	MinTimeStamp     *int64  `json:"MinTimeStamp,omitempty" xml:"MinTimeStamp,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsConsumerTimeSpanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsConsumerTimeSpanResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetMaxTimeStamp(v int64) *OnsConsumerTimeSpanResponseBodyData {
	s.MaxTimeStamp = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetConsumeTimeStamp(v int64) *OnsConsumerTimeSpanResponseBodyData {
	s.ConsumeTimeStamp = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetTopic(v string) *OnsConsumerTimeSpanResponseBodyData {
	s.Topic = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetMinTimeStamp(v int64) *OnsConsumerTimeSpanResponseBodyData {
	s.MinTimeStamp = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetInstanceId(v string) *OnsConsumerTimeSpanResponseBodyData {
	s.InstanceId = &v
	return s
}

type OnsConsumerTimeSpanResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsConsumerTimeSpanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsConsumerTimeSpanResponse) SetBody(v *OnsConsumerTimeSpanResponseBody) *OnsConsumerTimeSpanResponse {
	s.Body = v
	return s
}

type OnsDLQMessageGetByIdRequest struct {
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsDLQMessageGetByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageGetByIdRequest) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdRequest) SetMsgId(v string) *OnsDLQMessageGetByIdRequest {
	s.MsgId = &v
	return s
}

func (s *OnsDLQMessageGetByIdRequest) SetGroupId(v string) *OnsDLQMessageGetByIdRequest {
	s.GroupId = &v
	return s
}

func (s *OnsDLQMessageGetByIdRequest) SetInstanceId(v string) *OnsDLQMessageGetByIdRequest {
	s.InstanceId = &v
	return s
}

type OnsDLQMessageGetByIdResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsDLQMessageGetByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsDLQMessageGetByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBody) SetRequestId(v string) *OnsDLQMessageGetByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBody) SetData(v *OnsDLQMessageGetByIdResponseBodyData) *OnsDLQMessageGetByIdResponseBody {
	s.Data = v
	return s
}

type OnsDLQMessageGetByIdResponseBodyData struct {
	StoreSize      *int32                                              `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	ReconsumeTimes *int32                                              `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreTimestamp *int64                                              `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	InstanceId     *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                             `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	StoreHost      *string                                             `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	Topic          *string                                             `json:"Topic,omitempty" xml:"Topic,omitempty"`
	PropertyList   []*OnsDLQMessageGetByIdResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	BornTimestamp  *int64                                              `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	BodyCRC        *int32                                              `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                             `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
}

func (s OnsDLQMessageGetByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetStoreSize(v int32) *OnsDLQMessageGetByIdResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetReconsumeTimes(v int32) *OnsDLQMessageGetByIdResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetStoreTimestamp(v int64) *OnsDLQMessageGetByIdResponseBodyData {
	s.StoreTimestamp = &v
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

func (s *OnsDLQMessageGetByIdResponseBodyData) SetStoreHost(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetTopic(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.Topic = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetPropertyList(v []*OnsDLQMessageGetByIdResponseBodyDataPropertyList) *OnsDLQMessageGetByIdResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetBornTimestamp(v int64) *OnsDLQMessageGetByIdResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetBodyCRC(v int32) *OnsDLQMessageGetByIdResponseBodyData {
	s.BodyCRC = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetBornHost(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.BornHost = &v
	return s
}

type OnsDLQMessageGetByIdResponseBodyDataPropertyList struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s OnsDLQMessageGetByIdResponseBodyDataPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyList) SetValue(v string) *OnsDLQMessageGetByIdResponseBodyDataPropertyList {
	s.Value = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyList) SetName(v string) *OnsDLQMessageGetByIdResponseBodyDataPropertyList {
	s.Name = &v
	return s
}

type OnsDLQMessageGetByIdResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsDLQMessageGetByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsDLQMessageGetByIdResponse) SetBody(v *OnsDLQMessageGetByIdResponseBody) *OnsDLQMessageGetByIdResponse {
	s.Body = v
	return s
}

type OnsDLQMessagePageQueryByGroupIdRequest struct {
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	BeginTime   *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetGroupId(v string) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetBeginTime(v int64) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetEndTime(v int64) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.EndTime = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetTaskId(v string) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.TaskId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetCurrentPage(v int32) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetPageSize(v int32) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.PageSize = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetInstanceId(v string) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.InstanceId = &v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponseBody struct {
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MsgFoundDo *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo `json:"MsgFoundDo,omitempty" xml:"MsgFoundDo,omitempty" type:"Struct"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBody) SetRequestId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBody) SetMsgFoundDo(v *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) *OnsDLQMessagePageQueryByGroupIdResponseBody {
	s.MsgFoundDo = v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo struct {
	CurrentPage  *int64                                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	MsgFoundList []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList `json:"MsgFoundList,omitempty" xml:"MsgFoundList,omitempty" type:"Repeated"`
	MaxPageCount *int64                                                               `json:"MaxPageCount,omitempty" xml:"MaxPageCount,omitempty"`
	TaskId       *string                                                              `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetMsgFoundList(v []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.MsgFoundList = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetMaxPageCount(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.MaxPageCount = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo) SetTaskId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDo {
	s.TaskId = &v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList struct {
	StoreSize      *int32                                                                           `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	ReconsumeTimes *int32                                                                           `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreTimestamp *int64                                                                           `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	InstanceId     *string                                                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                                          `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	StoreHost      *string                                                                          `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	Topic          *string                                                                          `json:"Topic,omitempty" xml:"Topic,omitempty"`
	PropertyList   []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	BornTimestamp  *int64                                                                           `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	BodyCRC        *int32                                                                           `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                                                          `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetStoreSize(v int32) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.StoreSize = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetReconsumeTimes(v int32) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetStoreTimestamp(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetInstanceId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetMsgId(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.MsgId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetStoreHost(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.StoreHost = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetTopic(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.Topic = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetPropertyList(v []*OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListPropertyList) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.PropertyList = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetBornTimestamp(v int64) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.BornTimestamp = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetBodyCRC(v int32) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.BodyCRC = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList) SetBornHost(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundList {
	s.BornHost = &v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListPropertyList struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListPropertyList) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListPropertyList) SetValue(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListPropertyList {
	s.Value = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListPropertyList) SetName(v string) *OnsDLQMessagePageQueryByGroupIdResponseBodyMsgFoundDoMsgFoundListPropertyList {
	s.Name = &v
	return s
}

type OnsDLQMessagePageQueryByGroupIdResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsDLQMessagePageQueryByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsDLQMessagePageQueryByGroupIdResponse) SetBody(v *OnsDLQMessagePageQueryByGroupIdResponseBody) *OnsDLQMessagePageQueryByGroupIdResponse {
	s.Body = v
	return s
}

type OnsDLQMessageResendByIdRequest struct {
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsDLQMessageResendByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageResendByIdRequest) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageResendByIdRequest) SetMsgId(v string) *OnsDLQMessageResendByIdRequest {
	s.MsgId = &v
	return s
}

func (s *OnsDLQMessageResendByIdRequest) SetGroupId(v string) *OnsDLQMessageResendByIdRequest {
	s.GroupId = &v
	return s
}

func (s *OnsDLQMessageResendByIdRequest) SetInstanceId(v string) *OnsDLQMessageResendByIdRequest {
	s.InstanceId = &v
	return s
}

type OnsDLQMessageResendByIdResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s OnsDLQMessageResendByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsDLQMessageResendByIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageResendByIdResponseBody) SetRequestId(v string) *OnsDLQMessageResendByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsDLQMessageResendByIdResponseBody) SetData(v []*string) *OnsDLQMessageResendByIdResponseBody {
	s.Data = v
	return s
}

type OnsDLQMessageResendByIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsDLQMessageResendByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsDLQMessageResendByIdResponse) SetBody(v *OnsDLQMessageResendByIdResponseBody) *OnsDLQMessageResendByIdResponse {
	s.Body = v
	return s
}

type OnsGroupConsumerUpdateRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ReadEnable *bool   `json:"ReadEnable,omitempty" xml:"ReadEnable,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *OnsGroupConsumerUpdateRequest) SetReadEnable(v bool) *OnsGroupConsumerUpdateRequest {
	s.ReadEnable = &v
	return s
}

func (s *OnsGroupConsumerUpdateRequest) SetInstanceId(v string) *OnsGroupConsumerUpdateRequest {
	s.InstanceId = &v
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsGroupConsumerUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsGroupConsumerUpdateResponse) SetBody(v *OnsGroupConsumerUpdateResponseBody) *OnsGroupConsumerUpdateResponse {
	s.Body = v
	return s
}

type OnsGroupCreateRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupType  *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
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

func (s *OnsGroupCreateRequest) SetRemark(v string) *OnsGroupCreateRequest {
	s.Remark = &v
	return s
}

func (s *OnsGroupCreateRequest) SetInstanceId(v string) *OnsGroupCreateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupCreateRequest) SetGroupType(v string) *OnsGroupCreateRequest {
	s.GroupType = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsGroupCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsGroupDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsGroupDeleteResponse) SetBody(v *OnsGroupDeleteResponseBody) *OnsGroupDeleteResponse {
	s.Body = v
	return s
}

type OnsGroupListRequest struct {
	InstanceId *string                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupId    *string                   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupType  *string                   `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	Tag        []*OnsGroupListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupListRequest) SetInstanceId(v string) *OnsGroupListRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupListRequest) SetGroupId(v string) *OnsGroupListRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupListRequest) SetGroupType(v string) *OnsGroupListRequest {
	s.GroupType = &v
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
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*OnsGroupListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HelpUrl   *string                         `json:"HelpUrl,omitempty" xml:"HelpUrl,omitempty"`
}

func (s OnsGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBody) SetRequestId(v string) *OnsGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsGroupListResponseBody) SetData(v []*OnsGroupListResponseBodyData) *OnsGroupListResponseBody {
	s.Data = v
	return s
}

func (s *OnsGroupListResponseBody) SetHelpUrl(v string) *OnsGroupListResponseBody {
	s.HelpUrl = &v
	return s
}

type OnsGroupListResponseBodyData struct {
	Owner             *string                             `json:"Owner,omitempty" xml:"Owner,omitempty"`
	UpdateTime        *int64                              `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	IndependentNaming *bool                               `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	GroupId           *string                             `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Remark            *string                             `json:"Remark,omitempty" xml:"Remark,omitempty"`
	CreateTime        *int64                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Tags              []*OnsGroupListResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	InstanceId        *string                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupType         *string                             `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
}

func (s OnsGroupListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBodyData) SetOwner(v string) *OnsGroupListResponseBodyData {
	s.Owner = &v
	return s
}

func (s *OnsGroupListResponseBodyData) SetUpdateTime(v int64) *OnsGroupListResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *OnsGroupListResponseBodyData) SetIndependentNaming(v bool) *OnsGroupListResponseBodyData {
	s.IndependentNaming = &v
	return s
}

func (s *OnsGroupListResponseBodyData) SetGroupId(v string) *OnsGroupListResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *OnsGroupListResponseBodyData) SetRemark(v string) *OnsGroupListResponseBodyData {
	s.Remark = &v
	return s
}

func (s *OnsGroupListResponseBodyData) SetCreateTime(v int64) *OnsGroupListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *OnsGroupListResponseBodyData) SetTags(v []*OnsGroupListResponseBodyDataTags) *OnsGroupListResponseBodyData {
	s.Tags = v
	return s
}

func (s *OnsGroupListResponseBodyData) SetInstanceId(v string) *OnsGroupListResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupListResponseBodyData) SetGroupType(v string) *OnsGroupListResponseBodyData {
	s.GroupType = &v
	return s
}

type OnsGroupListResponseBodyDataTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsGroupListResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupListResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *OnsGroupListResponseBodyDataTags) SetKey(v string) *OnsGroupListResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *OnsGroupListResponseBodyDataTags) SetValue(v string) *OnsGroupListResponseBodyDataTags {
	s.Value = &v
	return s
}

type OnsGroupListResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsGroupListResponse) SetBody(v *OnsGroupListResponseBody) *OnsGroupListResponse {
	s.Body = v
	return s
}

type OnsGroupSubDetailRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s OnsGroupSubDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupSubDetailRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailRequest) SetInstanceId(v string) *OnsGroupSubDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupSubDetailRequest) SetGroupId(v string) *OnsGroupSubDetailRequest {
	s.GroupId = &v
	return s
}

type OnsGroupSubDetailResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsGroupSubDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsGroupSubDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupSubDetailResponseBody) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBody) SetRequestId(v string) *OnsGroupSubDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsGroupSubDetailResponseBody) SetData(v *OnsGroupSubDetailResponseBodyData) *OnsGroupSubDetailResponseBody {
	s.Data = v
	return s
}

type OnsGroupSubDetailResponseBodyData struct {
	SubscriptionDataList []*OnsGroupSubDetailResponseBodyDataSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Repeated"`
	GroupId              *string                                                  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MessageModel         *string                                                  `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	Online               *bool                                                    `json:"Online,omitempty" xml:"Online,omitempty"`
}

func (s OnsGroupSubDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupSubDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBodyData) SetSubscriptionDataList(v []*OnsGroupSubDetailResponseBodyDataSubscriptionDataList) *OnsGroupSubDetailResponseBodyData {
	s.SubscriptionDataList = v
	return s
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

type OnsGroupSubDetailResponseBodyDataSubscriptionDataList struct {
	SubString *string `json:"SubString,omitempty" xml:"SubString,omitempty"`
	Topic     *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsGroupSubDetailResponseBodyDataSubscriptionDataList) String() string {
	return tea.Prettify(s)
}

func (s OnsGroupSubDetailResponseBodyDataSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataList) SetSubString(v string) *OnsGroupSubDetailResponseBodyDataSubscriptionDataList {
	s.SubString = &v
	return s
}

func (s *OnsGroupSubDetailResponseBodyDataSubscriptionDataList) SetTopic(v string) *OnsGroupSubDetailResponseBodyDataSubscriptionDataList {
	s.Topic = &v
	return s
}

type OnsGroupSubDetailResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsGroupSubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Endpoints         *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	IndependentNaming *bool                                                     `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	MaxTps            *int64                                                    `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	Remark            *string                                                   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	InstanceName      *string                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ReleaseTime       *int64                                                    `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	TopicCapacity     *int32                                                    `json:"TopicCapacity,omitempty" xml:"TopicCapacity,omitempty"`
	InstanceStatus    *int32                                                    `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceId        *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType      *int32                                                    `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetEndpoints(v *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.Endpoints = v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetIndependentNaming(v bool) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.IndependentNaming = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetMaxTps(v int64) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.MaxTps = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetRemark(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.Remark = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceName(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetReleaseTime(v int64) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.ReleaseTime = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetTopicCapacity(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.TopicCapacity = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceStatus(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceStatus = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceId(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceType(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceType = &v
	return s
}

type OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints struct {
	TcpEndpoint                *string `json:"TcpEndpoint,omitempty" xml:"TcpEndpoint,omitempty"`
	HttpInternetEndpoint       *string `json:"HttpInternetEndpoint,omitempty" xml:"HttpInternetEndpoint,omitempty"`
	HttpInternalEndpoint       *string `json:"HttpInternalEndpoint,omitempty" xml:"HttpInternalEndpoint,omitempty"`
	HttpInternetSecureEndpoint *string `json:"HttpInternetSecureEndpoint,omitempty" xml:"HttpInternetSecureEndpoint,omitempty"`
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetTcpEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.TcpEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetHttpInternetEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.HttpInternetEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetHttpInternalEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.HttpInternalEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetHttpInternetSecureEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.HttpInternetSecureEndpoint = &v
	return s
}

type OnsInstanceBaseInfoResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsInstanceBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsInstanceBaseInfoResponse) SetBody(v *OnsInstanceBaseInfoResponseBody) *OnsInstanceBaseInfoResponse {
	s.Body = v
	return s
}

type OnsInstanceCreateRequest struct {
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s OnsInstanceCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceCreateRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceCreateRequest) SetRemark(v string) *OnsInstanceCreateRequest {
	s.Remark = &v
	return s
}

func (s *OnsInstanceCreateRequest) SetInstanceName(v string) *OnsInstanceCreateRequest {
	s.InstanceName = &v
	return s
}

type OnsInstanceCreateResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsInstanceCreateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsInstanceCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceCreateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceCreateResponseBody) SetRequestId(v string) *OnsInstanceCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsInstanceCreateResponseBody) SetData(v *OnsInstanceCreateResponseBodyData) *OnsInstanceCreateResponseBody {
	s.Data = v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsInstanceCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsInstanceDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*OnsInstanceInServiceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s OnsInstanceInServiceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBody) SetRequestId(v string) *OnsInstanceInServiceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBody) SetData(v []*OnsInstanceInServiceListResponseBodyData) *OnsInstanceInServiceListResponseBody {
	s.Data = v
	return s
}

type OnsInstanceInServiceListResponseBodyData struct {
	IndependentNaming *bool                                           `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	InstanceName      *string                                         `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ReleaseTime       *int64                                          `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	InstanceStatus    *int32                                          `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Tags              []*OnsInstanceInServiceListResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	InstanceId        *string                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType      *int32                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s OnsInstanceInServiceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBodyData) SetIndependentNaming(v bool) *OnsInstanceInServiceListResponseBodyData {
	s.IndependentNaming = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyData) SetInstanceName(v string) *OnsInstanceInServiceListResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyData) SetReleaseTime(v int64) *OnsInstanceInServiceListResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyData) SetInstanceStatus(v int32) *OnsInstanceInServiceListResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyData) SetTags(v []*OnsInstanceInServiceListResponseBodyDataTags) *OnsInstanceInServiceListResponseBodyData {
	s.Tags = v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyData) SetInstanceId(v string) *OnsInstanceInServiceListResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyData) SetInstanceType(v int32) *OnsInstanceInServiceListResponseBodyData {
	s.InstanceType = &v
	return s
}

type OnsInstanceInServiceListResponseBodyDataTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsInstanceInServiceListResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceInServiceListResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListResponseBodyDataTags) SetKey(v string) *OnsInstanceInServiceListResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *OnsInstanceInServiceListResponseBodyDataTags) SetValue(v string) *OnsInstanceInServiceListResponseBodyDataTags {
	s.Value = &v
	return s
}

type OnsInstanceInServiceListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsInstanceInServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsInstanceInServiceListResponse) SetBody(v *OnsInstanceInServiceListResponseBody) *OnsInstanceInServiceListResponse {
	s.Body = v
	return s
}

type OnsInstanceUpdateRequest struct {
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsInstanceUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsInstanceUpdateRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceUpdateRequest) SetRemark(v string) *OnsInstanceUpdateRequest {
	s.Remark = &v
	return s
}

func (s *OnsInstanceUpdateRequest) SetInstanceName(v string) *OnsInstanceUpdateRequest {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceUpdateRequest) SetInstanceId(v string) *OnsInstanceUpdateRequest {
	s.InstanceId = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsInstanceUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsInstanceUpdateResponse) SetBody(v *OnsInstanceUpdateResponseBody) *OnsInstanceUpdateResponse {
	s.Body = v
	return s
}

type OnsMessageGetByKeyRequest struct {
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMessageGetByKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyRequest) SetTopic(v string) *OnsMessageGetByKeyRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessageGetByKeyRequest) SetKey(v string) *OnsMessageGetByKeyRequest {
	s.Key = &v
	return s
}

func (s *OnsMessageGetByKeyRequest) SetInstanceId(v string) *OnsMessageGetByKeyRequest {
	s.InstanceId = &v
	return s
}

type OnsMessageGetByKeyResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*OnsMessageGetByKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s OnsMessageGetByKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBody) SetRequestId(v string) *OnsMessageGetByKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBody) SetData(v []*OnsMessageGetByKeyResponseBodyData) *OnsMessageGetByKeyResponseBody {
	s.Data = v
	return s
}

type OnsMessageGetByKeyResponseBodyData struct {
	StoreSize      *int32                                            `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	ReconsumeTimes *int32                                            `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreTimestamp *int64                                            `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	InstanceId     *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                           `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	StoreHost      *string                                           `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	Topic          *string                                           `json:"Topic,omitempty" xml:"Topic,omitempty"`
	PropertyList   []*OnsMessageGetByKeyResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	BornTimestamp  *int64                                            `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	BodyCRC        *int32                                            `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                           `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
}

func (s OnsMessageGetByKeyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBodyData) SetStoreSize(v int32) *OnsMessageGetByKeyResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) SetReconsumeTimes(v int32) *OnsMessageGetByKeyResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) SetStoreTimestamp(v int64) *OnsMessageGetByKeyResponseBodyData {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) SetInstanceId(v string) *OnsMessageGetByKeyResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) SetMsgId(v string) *OnsMessageGetByKeyResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) SetStoreHost(v string) *OnsMessageGetByKeyResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) SetTopic(v string) *OnsMessageGetByKeyResponseBodyData {
	s.Topic = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) SetPropertyList(v []*OnsMessageGetByKeyResponseBodyDataPropertyList) *OnsMessageGetByKeyResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) SetBornTimestamp(v int64) *OnsMessageGetByKeyResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) SetBodyCRC(v int32) *OnsMessageGetByKeyResponseBodyData {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyData) SetBornHost(v string) *OnsMessageGetByKeyResponseBodyData {
	s.BornHost = &v
	return s
}

type OnsMessageGetByKeyResponseBodyDataPropertyList struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s OnsMessageGetByKeyResponseBodyDataPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByKeyResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponseBodyDataPropertyList) SetValue(v string) *OnsMessageGetByKeyResponseBodyDataPropertyList {
	s.Value = &v
	return s
}

func (s *OnsMessageGetByKeyResponseBodyDataPropertyList) SetName(v string) *OnsMessageGetByKeyResponseBodyDataPropertyList {
	s.Name = &v
	return s
}

type OnsMessageGetByKeyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMessageGetByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsMessageGetByKeyResponse) SetBody(v *OnsMessageGetByKeyResponseBody) *OnsMessageGetByKeyResponse {
	s.Body = v
	return s
}

type OnsMessageGetByMsgIdRequest struct {
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMessageGetByMsgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByMsgIdRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdRequest) SetMsgId(v string) *OnsMessageGetByMsgIdRequest {
	s.MsgId = &v
	return s
}

func (s *OnsMessageGetByMsgIdRequest) SetTopic(v string) *OnsMessageGetByMsgIdRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessageGetByMsgIdRequest) SetInstanceId(v string) *OnsMessageGetByMsgIdRequest {
	s.InstanceId = &v
	return s
}

type OnsMessageGetByMsgIdResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsMessageGetByMsgIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsMessageGetByMsgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBody) SetRequestId(v string) *OnsMessageGetByMsgIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBody) SetData(v *OnsMessageGetByMsgIdResponseBodyData) *OnsMessageGetByMsgIdResponseBody {
	s.Data = v
	return s
}

type OnsMessageGetByMsgIdResponseBodyData struct {
	StoreSize      *int32                                              `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	ReconsumeTimes *int32                                              `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreTimestamp *int64                                              `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	InstanceId     *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                             `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	StoreHost      *string                                             `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	Topic          *string                                             `json:"Topic,omitempty" xml:"Topic,omitempty"`
	PropertyList   []*OnsMessageGetByMsgIdResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	BornTimestamp  *int64                                              `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	BodyCRC        *int32                                              `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                             `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
}

func (s OnsMessageGetByMsgIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetStoreSize(v int32) *OnsMessageGetByMsgIdResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetReconsumeTimes(v int32) *OnsMessageGetByMsgIdResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetStoreTimestamp(v int64) *OnsMessageGetByMsgIdResponseBodyData {
	s.StoreTimestamp = &v
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

func (s *OnsMessageGetByMsgIdResponseBodyData) SetStoreHost(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetTopic(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.Topic = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetPropertyList(v []*OnsMessageGetByMsgIdResponseBodyDataPropertyList) *OnsMessageGetByMsgIdResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetBornTimestamp(v int64) *OnsMessageGetByMsgIdResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetBodyCRC(v int32) *OnsMessageGetByMsgIdResponseBodyData {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetBornHost(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.BornHost = &v
	return s
}

type OnsMessageGetByMsgIdResponseBodyDataPropertyList struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s OnsMessageGetByMsgIdResponseBodyDataPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyList) SetValue(v string) *OnsMessageGetByMsgIdResponseBodyDataPropertyList {
	s.Value = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyList) SetName(v string) *OnsMessageGetByMsgIdResponseBodyDataPropertyList {
	s.Name = &v
	return s
}

type OnsMessageGetByMsgIdResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMessageGetByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsMessageGetByMsgIdResponse) SetBody(v *OnsMessageGetByMsgIdResponseBody) *OnsMessageGetByMsgIdResponse {
	s.Body = v
	return s
}

type OnsMessagePageQueryByTopicRequest struct {
	Topic       *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	BeginTime   *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMessagePageQueryByTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicRequest) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicRequest) SetTopic(v string) *OnsMessagePageQueryByTopicRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetBeginTime(v int64) *OnsMessagePageQueryByTopicRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetEndTime(v int64) *OnsMessagePageQueryByTopicRequest {
	s.EndTime = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetTaskId(v string) *OnsMessagePageQueryByTopicRequest {
	s.TaskId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetCurrentPage(v int32) *OnsMessagePageQueryByTopicRequest {
	s.CurrentPage = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetPageSize(v int32) *OnsMessagePageQueryByTopicRequest {
	s.PageSize = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetInstanceId(v string) *OnsMessagePageQueryByTopicRequest {
	s.InstanceId = &v
	return s
}

type OnsMessagePageQueryByTopicResponseBody struct {
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MsgFoundDo *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo `json:"MsgFoundDo,omitempty" xml:"MsgFoundDo,omitempty" type:"Struct"`
}

func (s OnsMessagePageQueryByTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBody) SetRequestId(v string) *OnsMessagePageQueryByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBody) SetMsgFoundDo(v *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) *OnsMessagePageQueryByTopicResponseBody {
	s.MsgFoundDo = v
	return s
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDo struct {
	CurrentPage  *int64                                                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	MsgFoundList []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList `json:"MsgFoundList,omitempty" xml:"MsgFoundList,omitempty" type:"Repeated"`
	MaxPageCount *int64                                                          `json:"MaxPageCount,omitempty" xml:"MaxPageCount,omitempty"`
	TaskId       *string                                                         `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetMsgFoundList(v []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.MsgFoundList = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetMaxPageCount(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.MaxPageCount = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo) SetTaskId(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDo {
	s.TaskId = &v
	return s
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList struct {
	StoreSize      *int32                                                                      `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	ReconsumeTimes *int32                                                                      `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	StoreTimestamp *int64                                                                      `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	InstanceId     *string                                                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId          *string                                                                     `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	StoreHost      *string                                                                     `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	Topic          *string                                                                     `json:"Topic,omitempty" xml:"Topic,omitempty"`
	PropertyList   []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	BornTimestamp  *int64                                                                      `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	BodyCRC        *int32                                                                      `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	BornHost       *string                                                                     `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetStoreSize(v int32) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.StoreSize = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetReconsumeTimes(v int32) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetStoreTimestamp(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetInstanceId(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.InstanceId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetMsgId(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.MsgId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetStoreHost(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.StoreHost = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetTopic(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.Topic = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetPropertyList(v []*OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListPropertyList) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.PropertyList = v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetBornTimestamp(v int64) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetBodyCRC(v int32) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList) SetBornHost(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundList {
	s.BornHost = &v
	return s
}

type OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListPropertyList struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListPropertyList) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListPropertyList) SetValue(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListPropertyList {
	s.Value = &v
	return s
}

func (s *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListPropertyList) SetName(v string) *OnsMessagePageQueryByTopicResponseBodyMsgFoundDoMsgFoundListPropertyList {
	s.Name = &v
	return s
}

type OnsMessagePageQueryByTopicResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMessagePageQueryByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsMessagePageQueryByTopicResponse) SetBody(v *OnsMessagePageQueryByTopicResponseBody) *OnsMessagePageQueryByTopicResponse {
	s.Body = v
	return s
}

type OnsMessagePushRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMessagePushRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessagePushRequest) GoString() string {
	return s.String()
}

func (s *OnsMessagePushRequest) SetGroupId(v string) *OnsMessagePushRequest {
	s.GroupId = &v
	return s
}

func (s *OnsMessagePushRequest) SetClientId(v string) *OnsMessagePushRequest {
	s.ClientId = &v
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

func (s *OnsMessagePushRequest) SetInstanceId(v string) *OnsMessagePushRequest {
	s.InstanceId = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMessagePushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsMessagePushResponse) SetBody(v *OnsMessagePushResponseBody) *OnsMessagePushResponse {
	s.Body = v
	return s
}

type OnsMessageSendRequest struct {
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Tag        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMessageSendRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageSendRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageSendRequest) SetTopic(v string) *OnsMessageSendRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessageSendRequest) SetTag(v string) *OnsMessageSendRequest {
	s.Tag = &v
	return s
}

func (s *OnsMessageSendRequest) SetKey(v string) *OnsMessageSendRequest {
	s.Key = &v
	return s
}

func (s *OnsMessageSendRequest) SetMessage(v string) *OnsMessageSendRequest {
	s.Message = &v
	return s
}

func (s *OnsMessageSendRequest) SetInstanceId(v string) *OnsMessageSendRequest {
	s.InstanceId = &v
	return s
}

type OnsMessageSendResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s OnsMessageSendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageSendResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageSendResponseBody) SetRequestId(v string) *OnsMessageSendResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessageSendResponseBody) SetData(v string) *OnsMessageSendResponseBody {
	s.Data = &v
	return s
}

type OnsMessageSendResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMessageSendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsMessageSendResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageSendResponse) GoString() string {
	return s.String()
}

func (s *OnsMessageSendResponse) SetHeaders(v map[string]*string) *OnsMessageSendResponse {
	s.Headers = v
	return s
}

func (s *OnsMessageSendResponse) SetBody(v *OnsMessageSendResponseBody) *OnsMessageSendResponse {
	s.Body = v
	return s
}

type OnsMessageTraceRequest struct {
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMessageTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageTraceRequest) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceRequest) SetTopic(v string) *OnsMessageTraceRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessageTraceRequest) SetMsgId(v string) *OnsMessageTraceRequest {
	s.MsgId = &v
	return s
}

func (s *OnsMessageTraceRequest) SetInstanceId(v string) *OnsMessageTraceRequest {
	s.InstanceId = &v
	return s
}

type OnsMessageTraceResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*OnsMessageTraceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s OnsMessageTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageTraceResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceResponseBody) SetRequestId(v string) *OnsMessageTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessageTraceResponseBody) SetData(v []*OnsMessageTraceResponseBodyData) *OnsMessageTraceResponseBody {
	s.Data = v
	return s
}

type OnsMessageTraceResponseBodyData struct {
	TrackType     *string `json:"TrackType,omitempty" xml:"TrackType,omitempty"`
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMessageTraceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsMessageTraceResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageTraceResponseBodyData) SetTrackType(v string) *OnsMessageTraceResponseBodyData {
	s.TrackType = &v
	return s
}

func (s *OnsMessageTraceResponseBodyData) SetConsumerGroup(v string) *OnsMessageTraceResponseBodyData {
	s.ConsumerGroup = &v
	return s
}

func (s *OnsMessageTraceResponseBodyData) SetInstanceId(v string) *OnsMessageTraceResponseBodyData {
	s.InstanceId = &v
	return s
}

type OnsMessageTraceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMessageTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsMessageTraceResponse) SetBody(v *OnsMessageTraceResponseBody) *OnsMessageTraceResponse {
	s.Body = v
	return s
}

type OnsMqttGroupIdCreateRequest struct {
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMqttGroupIdCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttGroupIdCreateRequest) GoString() string {
	return s.String()
}

func (s *OnsMqttGroupIdCreateRequest) SetTopic(v string) *OnsMqttGroupIdCreateRequest {
	s.Topic = &v
	return s
}

func (s *OnsMqttGroupIdCreateRequest) SetGroupId(v string) *OnsMqttGroupIdCreateRequest {
	s.GroupId = &v
	return s
}

func (s *OnsMqttGroupIdCreateRequest) SetInstanceId(v string) *OnsMqttGroupIdCreateRequest {
	s.InstanceId = &v
	return s
}

type OnsMqttGroupIdCreateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMqttGroupIdCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttGroupIdCreateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMqttGroupIdCreateResponseBody) SetRequestId(v string) *OnsMqttGroupIdCreateResponseBody {
	s.RequestId = &v
	return s
}

type OnsMqttGroupIdCreateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMqttGroupIdCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsMqttGroupIdCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttGroupIdCreateResponse) GoString() string {
	return s.String()
}

func (s *OnsMqttGroupIdCreateResponse) SetHeaders(v map[string]*string) *OnsMqttGroupIdCreateResponse {
	s.Headers = v
	return s
}

func (s *OnsMqttGroupIdCreateResponse) SetBody(v *OnsMqttGroupIdCreateResponseBody) *OnsMqttGroupIdCreateResponse {
	s.Body = v
	return s
}

type OnsMqttGroupIdDeleteRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMqttGroupIdDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttGroupIdDeleteRequest) GoString() string {
	return s.String()
}

func (s *OnsMqttGroupIdDeleteRequest) SetGroupId(v string) *OnsMqttGroupIdDeleteRequest {
	s.GroupId = &v
	return s
}

func (s *OnsMqttGroupIdDeleteRequest) SetInstanceId(v string) *OnsMqttGroupIdDeleteRequest {
	s.InstanceId = &v
	return s
}

type OnsMqttGroupIdDeleteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMqttGroupIdDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttGroupIdDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMqttGroupIdDeleteResponseBody) SetRequestId(v string) *OnsMqttGroupIdDeleteResponseBody {
	s.RequestId = &v
	return s
}

type OnsMqttGroupIdDeleteResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMqttGroupIdDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsMqttGroupIdDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttGroupIdDeleteResponse) GoString() string {
	return s.String()
}

func (s *OnsMqttGroupIdDeleteResponse) SetHeaders(v map[string]*string) *OnsMqttGroupIdDeleteResponse {
	s.Headers = v
	return s
}

func (s *OnsMqttGroupIdDeleteResponse) SetBody(v *OnsMqttGroupIdDeleteResponseBody) *OnsMqttGroupIdDeleteResponse {
	s.Body = v
	return s
}

type OnsMqttGroupIdListRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMqttGroupIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttGroupIdListRequest) GoString() string {
	return s.String()
}

func (s *OnsMqttGroupIdListRequest) SetInstanceId(v string) *OnsMqttGroupIdListRequest {
	s.InstanceId = &v
	return s
}

type OnsMqttGroupIdListResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*OnsMqttGroupIdListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s OnsMqttGroupIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttGroupIdListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMqttGroupIdListResponseBody) SetRequestId(v string) *OnsMqttGroupIdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMqttGroupIdListResponseBody) SetData(v []*OnsMqttGroupIdListResponseBodyData) *OnsMqttGroupIdListResponseBody {
	s.Data = v
	return s
}

type OnsMqttGroupIdListResponseBodyData struct {
	UpdateTime        *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	IndependentNaming *bool   `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMqttGroupIdListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttGroupIdListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMqttGroupIdListResponseBodyData) SetUpdateTime(v int64) *OnsMqttGroupIdListResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *OnsMqttGroupIdListResponseBodyData) SetIndependentNaming(v bool) *OnsMqttGroupIdListResponseBodyData {
	s.IndependentNaming = &v
	return s
}

func (s *OnsMqttGroupIdListResponseBodyData) SetGroupId(v string) *OnsMqttGroupIdListResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *OnsMqttGroupIdListResponseBodyData) SetCreateTime(v int64) *OnsMqttGroupIdListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *OnsMqttGroupIdListResponseBodyData) SetInstanceId(v string) *OnsMqttGroupIdListResponseBodyData {
	s.InstanceId = &v
	return s
}

type OnsMqttGroupIdListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMqttGroupIdListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsMqttGroupIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttGroupIdListResponse) GoString() string {
	return s.String()
}

func (s *OnsMqttGroupIdListResponse) SetHeaders(v map[string]*string) *OnsMqttGroupIdListResponse {
	s.Headers = v
	return s
}

func (s *OnsMqttGroupIdListResponse) SetBody(v *OnsMqttGroupIdListResponseBody) *OnsMqttGroupIdListResponse {
	s.Body = v
	return s
}

type OnsMqttQueryClientByClientIdRequest struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMqttQueryClientByClientIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByClientIdRequest) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByClientIdRequest) SetClientId(v string) *OnsMqttQueryClientByClientIdRequest {
	s.ClientId = &v
	return s
}

func (s *OnsMqttQueryClientByClientIdRequest) SetInstanceId(v string) *OnsMqttQueryClientByClientIdRequest {
	s.InstanceId = &v
	return s
}

type OnsMqttQueryClientByClientIdResponseBody struct {
	RequestId        *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MqttClientInfoDo *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo `json:"MqttClientInfoDo,omitempty" xml:"MqttClientInfoDo,omitempty" type:"Struct"`
}

func (s OnsMqttQueryClientByClientIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByClientIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByClientIdResponseBody) SetRequestId(v string) *OnsMqttQueryClientByClientIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMqttQueryClientByClientIdResponseBody) SetMqttClientInfoDo(v *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo) *OnsMqttQueryClientByClientIdResponseBody {
	s.MqttClientInfoDo = v
	return s
}

type OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo struct {
	Online          *bool                                                                      `json:"Online,omitempty" xml:"Online,omitempty"`
	LastTouch       *int64                                                                     `json:"LastTouch,omitempty" xml:"LastTouch,omitempty"`
	SocketChannel   *string                                                                    `json:"SocketChannel,omitempty" xml:"SocketChannel,omitempty"`
	ClientId        *string                                                                    `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	SubScriptonData []*OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData `json:"SubScriptonData,omitempty" xml:"SubScriptonData,omitempty" type:"Repeated"`
}

func (s OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo) SetOnline(v bool) *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo {
	s.Online = &v
	return s
}

func (s *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo) SetLastTouch(v int64) *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo {
	s.LastTouch = &v
	return s
}

func (s *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo) SetSocketChannel(v string) *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo {
	s.SocketChannel = &v
	return s
}

func (s *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo) SetClientId(v string) *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo {
	s.ClientId = &v
	return s
}

func (s *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo) SetSubScriptonData(v []*OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData) *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDo {
	s.SubScriptonData = v
	return s
}

type OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData struct {
	SubTopic    *string `json:"SubTopic,omitempty" xml:"SubTopic,omitempty"`
	ParentTopic *string `json:"ParentTopic,omitempty" xml:"ParentTopic,omitempty"`
	Qos         *int32  `json:"Qos,omitempty" xml:"Qos,omitempty"`
}

func (s OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData) SetSubTopic(v string) *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData {
	s.SubTopic = &v
	return s
}

func (s *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData) SetParentTopic(v string) *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData {
	s.ParentTopic = &v
	return s
}

func (s *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData) SetQos(v int32) *OnsMqttQueryClientByClientIdResponseBodyMqttClientInfoDoSubScriptonData {
	s.Qos = &v
	return s
}

type OnsMqttQueryClientByClientIdResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMqttQueryClientByClientIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsMqttQueryClientByClientIdResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByClientIdResponse) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByClientIdResponse) SetHeaders(v map[string]*string) *OnsMqttQueryClientByClientIdResponse {
	s.Headers = v
	return s
}

func (s *OnsMqttQueryClientByClientIdResponse) SetBody(v *OnsMqttQueryClientByClientIdResponseBody) *OnsMqttQueryClientByClientIdResponse {
	s.Body = v
	return s
}

type OnsMqttQueryClientByGroupIdRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMqttQueryClientByGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByGroupIdRequest) SetGroupId(v string) *OnsMqttQueryClientByGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *OnsMqttQueryClientByGroupIdRequest) SetInstanceId(v string) *OnsMqttQueryClientByGroupIdRequest {
	s.InstanceId = &v
	return s
}

type OnsMqttQueryClientByGroupIdResponseBody struct {
	RequestId       *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MqttClientSetDo *OnsMqttQueryClientByGroupIdResponseBodyMqttClientSetDo `json:"MqttClientSetDo,omitempty" xml:"MqttClientSetDo,omitempty" type:"Struct"`
}

func (s OnsMqttQueryClientByGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByGroupIdResponseBody) SetRequestId(v string) *OnsMqttQueryClientByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMqttQueryClientByGroupIdResponseBody) SetMqttClientSetDo(v *OnsMqttQueryClientByGroupIdResponseBodyMqttClientSetDo) *OnsMqttQueryClientByGroupIdResponseBody {
	s.MqttClientSetDo = v
	return s
}

type OnsMqttQueryClientByGroupIdResponseBodyMqttClientSetDo struct {
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
}

func (s OnsMqttQueryClientByGroupIdResponseBodyMqttClientSetDo) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByGroupIdResponseBodyMqttClientSetDo) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByGroupIdResponseBodyMqttClientSetDo) SetOnlineCount(v int64) *OnsMqttQueryClientByGroupIdResponseBodyMqttClientSetDo {
	s.OnlineCount = &v
	return s
}

type OnsMqttQueryClientByGroupIdResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMqttQueryClientByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsMqttQueryClientByGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByGroupIdResponse) SetHeaders(v map[string]*string) *OnsMqttQueryClientByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *OnsMqttQueryClientByGroupIdResponse) SetBody(v *OnsMqttQueryClientByGroupIdResponseBody) *OnsMqttQueryClientByGroupIdResponse {
	s.Body = v
	return s
}

type OnsMqttQueryClientByTopicRequest struct {
	ParentTopic *string `json:"ParentTopic,omitempty" xml:"ParentTopic,omitempty"`
	SubTopic    *string `json:"SubTopic,omitempty" xml:"SubTopic,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMqttQueryClientByTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByTopicRequest) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByTopicRequest) SetParentTopic(v string) *OnsMqttQueryClientByTopicRequest {
	s.ParentTopic = &v
	return s
}

func (s *OnsMqttQueryClientByTopicRequest) SetSubTopic(v string) *OnsMqttQueryClientByTopicRequest {
	s.SubTopic = &v
	return s
}

func (s *OnsMqttQueryClientByTopicRequest) SetInstanceId(v string) *OnsMqttQueryClientByTopicRequest {
	s.InstanceId = &v
	return s
}

type OnsMqttQueryClientByTopicResponseBody struct {
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MqttClientSetDo *OnsMqttQueryClientByTopicResponseBodyMqttClientSetDo `json:"MqttClientSetDo,omitempty" xml:"MqttClientSetDo,omitempty" type:"Struct"`
}

func (s OnsMqttQueryClientByTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByTopicResponseBody) SetRequestId(v string) *OnsMqttQueryClientByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMqttQueryClientByTopicResponseBody) SetMqttClientSetDo(v *OnsMqttQueryClientByTopicResponseBodyMqttClientSetDo) *OnsMqttQueryClientByTopicResponseBody {
	s.MqttClientSetDo = v
	return s
}

type OnsMqttQueryClientByTopicResponseBodyMqttClientSetDo struct {
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
}

func (s OnsMqttQueryClientByTopicResponseBodyMqttClientSetDo) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByTopicResponseBodyMqttClientSetDo) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByTopicResponseBodyMqttClientSetDo) SetOnlineCount(v int64) *OnsMqttQueryClientByTopicResponseBodyMqttClientSetDo {
	s.OnlineCount = &v
	return s
}

type OnsMqttQueryClientByTopicResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMqttQueryClientByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsMqttQueryClientByTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryClientByTopicResponse) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryClientByTopicResponse) SetHeaders(v map[string]*string) *OnsMqttQueryClientByTopicResponse {
	s.Headers = v
	return s
}

func (s *OnsMqttQueryClientByTopicResponse) SetBody(v *OnsMqttQueryClientByTopicResponseBody) *OnsMqttQueryClientByTopicResponse {
	s.Body = v
	return s
}

type OnsMqttQueryHistoryOnlineRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMqttQueryHistoryOnlineRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryHistoryOnlineRequest) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryHistoryOnlineRequest) SetGroupId(v string) *OnsMqttQueryHistoryOnlineRequest {
	s.GroupId = &v
	return s
}

func (s *OnsMqttQueryHistoryOnlineRequest) SetBeginTime(v int64) *OnsMqttQueryHistoryOnlineRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsMqttQueryHistoryOnlineRequest) SetEndTime(v int64) *OnsMqttQueryHistoryOnlineRequest {
	s.EndTime = &v
	return s
}

func (s *OnsMqttQueryHistoryOnlineRequest) SetInstanceId(v string) *OnsMqttQueryHistoryOnlineRequest {
	s.InstanceId = &v
	return s
}

type OnsMqttQueryHistoryOnlineResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsMqttQueryHistoryOnlineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsMqttQueryHistoryOnlineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryHistoryOnlineResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryHistoryOnlineResponseBody) SetRequestId(v string) *OnsMqttQueryHistoryOnlineResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMqttQueryHistoryOnlineResponseBody) SetData(v *OnsMqttQueryHistoryOnlineResponseBodyData) *OnsMqttQueryHistoryOnlineResponseBody {
	s.Data = v
	return s
}

type OnsMqttQueryHistoryOnlineResponseBodyData struct {
	Records []*OnsMqttQueryHistoryOnlineResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	XUnit   *string                                             `json:"XUnit,omitempty" xml:"XUnit,omitempty"`
	YUnit   *string                                             `json:"YUnit,omitempty" xml:"YUnit,omitempty"`
	Title   *string                                             `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s OnsMqttQueryHistoryOnlineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryHistoryOnlineResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryHistoryOnlineResponseBodyData) SetRecords(v []*OnsMqttQueryHistoryOnlineResponseBodyDataRecords) *OnsMqttQueryHistoryOnlineResponseBodyData {
	s.Records = v
	return s
}

func (s *OnsMqttQueryHistoryOnlineResponseBodyData) SetXUnit(v string) *OnsMqttQueryHistoryOnlineResponseBodyData {
	s.XUnit = &v
	return s
}

func (s *OnsMqttQueryHistoryOnlineResponseBodyData) SetYUnit(v string) *OnsMqttQueryHistoryOnlineResponseBodyData {
	s.YUnit = &v
	return s
}

func (s *OnsMqttQueryHistoryOnlineResponseBodyData) SetTitle(v string) *OnsMqttQueryHistoryOnlineResponseBodyData {
	s.Title = &v
	return s
}

type OnsMqttQueryHistoryOnlineResponseBodyDataRecords struct {
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *int64   `json:"X,omitempty" xml:"X,omitempty"`
}

func (s OnsMqttQueryHistoryOnlineResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryHistoryOnlineResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryHistoryOnlineResponseBodyDataRecords) SetY(v float32) *OnsMqttQueryHistoryOnlineResponseBodyDataRecords {
	s.Y = &v
	return s
}

func (s *OnsMqttQueryHistoryOnlineResponseBodyDataRecords) SetX(v int64) *OnsMqttQueryHistoryOnlineResponseBodyDataRecords {
	s.X = &v
	return s
}

type OnsMqttQueryHistoryOnlineResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMqttQueryHistoryOnlineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsMqttQueryHistoryOnlineResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryHistoryOnlineResponse) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryHistoryOnlineResponse) SetHeaders(v map[string]*string) *OnsMqttQueryHistoryOnlineResponse {
	s.Headers = v
	return s
}

func (s *OnsMqttQueryHistoryOnlineResponse) SetBody(v *OnsMqttQueryHistoryOnlineResponseBody) *OnsMqttQueryHistoryOnlineResponse {
	s.Body = v
	return s
}

type OnsMqttQueryMsgTransTrendRequest struct {
	TpsType     *string `json:"TpsType,omitempty" xml:"TpsType,omitempty"`
	TransType   *string `json:"TransType,omitempty" xml:"TransType,omitempty"`
	ParentTopic *string `json:"ParentTopic,omitempty" xml:"ParentTopic,omitempty"`
	SubTopic    *string `json:"SubTopic,omitempty" xml:"SubTopic,omitempty"`
	MsgType     *string `json:"MsgType,omitempty" xml:"MsgType,omitempty"`
	Qos         *int32  `json:"Qos,omitempty" xml:"Qos,omitempty"`
	BeginTime   *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsMqttQueryMsgTransTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryMsgTransTrendRequest) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryMsgTransTrendRequest) SetTpsType(v string) *OnsMqttQueryMsgTransTrendRequest {
	s.TpsType = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendRequest) SetTransType(v string) *OnsMqttQueryMsgTransTrendRequest {
	s.TransType = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendRequest) SetParentTopic(v string) *OnsMqttQueryMsgTransTrendRequest {
	s.ParentTopic = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendRequest) SetSubTopic(v string) *OnsMqttQueryMsgTransTrendRequest {
	s.SubTopic = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendRequest) SetMsgType(v string) *OnsMqttQueryMsgTransTrendRequest {
	s.MsgType = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendRequest) SetQos(v int32) *OnsMqttQueryMsgTransTrendRequest {
	s.Qos = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendRequest) SetBeginTime(v int64) *OnsMqttQueryMsgTransTrendRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendRequest) SetEndTime(v int64) *OnsMqttQueryMsgTransTrendRequest {
	s.EndTime = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendRequest) SetInstanceId(v string) *OnsMqttQueryMsgTransTrendRequest {
	s.InstanceId = &v
	return s
}

type OnsMqttQueryMsgTransTrendResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsMqttQueryMsgTransTrendResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsMqttQueryMsgTransTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryMsgTransTrendResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryMsgTransTrendResponseBody) SetRequestId(v string) *OnsMqttQueryMsgTransTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendResponseBody) SetData(v *OnsMqttQueryMsgTransTrendResponseBodyData) *OnsMqttQueryMsgTransTrendResponseBody {
	s.Data = v
	return s
}

type OnsMqttQueryMsgTransTrendResponseBodyData struct {
	Records []*OnsMqttQueryMsgTransTrendResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	XUnit   *string                                             `json:"XUnit,omitempty" xml:"XUnit,omitempty"`
	YUnit   *string                                             `json:"YUnit,omitempty" xml:"YUnit,omitempty"`
	Title   *string                                             `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s OnsMqttQueryMsgTransTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryMsgTransTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryMsgTransTrendResponseBodyData) SetRecords(v []*OnsMqttQueryMsgTransTrendResponseBodyDataRecords) *OnsMqttQueryMsgTransTrendResponseBodyData {
	s.Records = v
	return s
}

func (s *OnsMqttQueryMsgTransTrendResponseBodyData) SetXUnit(v string) *OnsMqttQueryMsgTransTrendResponseBodyData {
	s.XUnit = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendResponseBodyData) SetYUnit(v string) *OnsMqttQueryMsgTransTrendResponseBodyData {
	s.YUnit = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendResponseBodyData) SetTitle(v string) *OnsMqttQueryMsgTransTrendResponseBodyData {
	s.Title = &v
	return s
}

type OnsMqttQueryMsgTransTrendResponseBodyDataRecords struct {
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *int64   `json:"X,omitempty" xml:"X,omitempty"`
}

func (s OnsMqttQueryMsgTransTrendResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryMsgTransTrendResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryMsgTransTrendResponseBodyDataRecords) SetY(v float32) *OnsMqttQueryMsgTransTrendResponseBodyDataRecords {
	s.Y = &v
	return s
}

func (s *OnsMqttQueryMsgTransTrendResponseBodyDataRecords) SetX(v int64) *OnsMqttQueryMsgTransTrendResponseBodyDataRecords {
	s.X = &v
	return s
}

type OnsMqttQueryMsgTransTrendResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsMqttQueryMsgTransTrendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsMqttQueryMsgTransTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsMqttQueryMsgTransTrendResponse) GoString() string {
	return s.String()
}

func (s *OnsMqttQueryMsgTransTrendResponse) SetHeaders(v map[string]*string) *OnsMqttQueryMsgTransTrendResponse {
	s.Headers = v
	return s
}

func (s *OnsMqttQueryMsgTransTrendResponse) SetBody(v *OnsMqttQueryMsgTransTrendResponseBody) *OnsMqttQueryMsgTransTrendResponse {
	s.Body = v
	return s
}

type OnsRegionListResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*OnsRegionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s OnsRegionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsRegionListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponseBody) SetRequestId(v string) *OnsRegionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsRegionListResponseBody) SetData(v []*OnsRegionListResponseBodyData) *OnsRegionListResponseBody {
	s.Data = v
	return s
}

type OnsRegionListResponseBodyData struct {
	RegionName  *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	OnsRegionId *string `json:"OnsRegionId,omitempty" xml:"OnsRegionId,omitempty"`
}

func (s OnsRegionListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsRegionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponseBodyData) SetRegionName(v string) *OnsRegionListResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *OnsRegionListResponseBodyData) SetOnsRegionId(v string) *OnsRegionListResponseBodyData {
	s.OnsRegionId = &v
	return s
}

type OnsRegionListResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsRegionListResponse) SetBody(v *OnsRegionListResponseBody) *OnsRegionListResponse {
	s.Body = v
	return s
}

type OnsTopicCreateRequest struct {
	Topic       *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	MessageType *int32  `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsTopicCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicCreateRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicCreateRequest) SetTopic(v string) *OnsTopicCreateRequest {
	s.Topic = &v
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

func (s *OnsTopicCreateRequest) SetInstanceId(v string) *OnsTopicCreateRequest {
	s.InstanceId = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTopicCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsTopicCreateResponse) SetBody(v *OnsTopicCreateResponseBody) *OnsTopicCreateResponse {
	s.Body = v
	return s
}

type OnsTopicDeleteRequest struct {
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsTopicDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicDeleteRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicDeleteRequest) SetTopic(v string) *OnsTopicDeleteRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicDeleteRequest) SetInstanceId(v string) *OnsTopicDeleteRequest {
	s.InstanceId = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTopicDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsTopicDeleteResponse) SetBody(v *OnsTopicDeleteResponseBody) *OnsTopicDeleteResponse {
	s.Body = v
	return s
}

type OnsTopicListRequest struct {
	Topic      *string                   `json:"Topic,omitempty" xml:"Topic,omitempty"`
	InstanceId *string                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Tag        []*OnsTopicListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsTopicListRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicListRequest) SetTopic(v string) *OnsTopicListRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicListRequest) SetInstanceId(v string) *OnsTopicListRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicListRequest) SetTag(v []*OnsTopicListRequestTag) *OnsTopicListRequest {
	s.Tag = v
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
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*OnsTopicListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HelpUrl   *string                         `json:"HelpUrl,omitempty" xml:"HelpUrl,omitempty"`
}

func (s OnsTopicListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBody) SetRequestId(v string) *OnsTopicListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTopicListResponseBody) SetData(v []*OnsTopicListResponseBodyData) *OnsTopicListResponseBody {
	s.Data = v
	return s
}

func (s *OnsTopicListResponseBody) SetHelpUrl(v string) *OnsTopicListResponseBody {
	s.HelpUrl = &v
	return s
}

type OnsTopicListResponseBodyData struct {
	MessageType       *int32                              `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	RelationName      *string                             `json:"RelationName,omitempty" xml:"RelationName,omitempty"`
	Owner             *string                             `json:"Owner,omitempty" xml:"Owner,omitempty"`
	IndependentNaming *bool                               `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	Remark            *string                             `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Relation          *int32                              `json:"Relation,omitempty" xml:"Relation,omitempty"`
	CreateTime        *int64                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Topic             *string                             `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Tags              []*OnsTopicListResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	InstanceId        *string                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsTopicListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBodyData) SetMessageType(v int32) *OnsTopicListResponseBodyData {
	s.MessageType = &v
	return s
}

func (s *OnsTopicListResponseBodyData) SetRelationName(v string) *OnsTopicListResponseBodyData {
	s.RelationName = &v
	return s
}

func (s *OnsTopicListResponseBodyData) SetOwner(v string) *OnsTopicListResponseBodyData {
	s.Owner = &v
	return s
}

func (s *OnsTopicListResponseBodyData) SetIndependentNaming(v bool) *OnsTopicListResponseBodyData {
	s.IndependentNaming = &v
	return s
}

func (s *OnsTopicListResponseBodyData) SetRemark(v string) *OnsTopicListResponseBodyData {
	s.Remark = &v
	return s
}

func (s *OnsTopicListResponseBodyData) SetRelation(v int32) *OnsTopicListResponseBodyData {
	s.Relation = &v
	return s
}

func (s *OnsTopicListResponseBodyData) SetCreateTime(v int64) *OnsTopicListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *OnsTopicListResponseBodyData) SetTopic(v string) *OnsTopicListResponseBodyData {
	s.Topic = &v
	return s
}

func (s *OnsTopicListResponseBodyData) SetTags(v []*OnsTopicListResponseBodyDataTags) *OnsTopicListResponseBodyData {
	s.Tags = v
	return s
}

func (s *OnsTopicListResponseBodyData) SetInstanceId(v string) *OnsTopicListResponseBodyData {
	s.InstanceId = &v
	return s
}

type OnsTopicListResponseBodyDataTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsTopicListResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicListResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *OnsTopicListResponseBodyDataTags) SetKey(v string) *OnsTopicListResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *OnsTopicListResponseBodyDataTags) SetValue(v string) *OnsTopicListResponseBodyDataTags {
	s.Value = &v
	return s
}

type OnsTopicListResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTopicListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsTopicListResponse) SetBody(v *OnsTopicListResponseBody) *OnsTopicListResponse {
	s.Body = v
	return s
}

type OnsTopicStatusRequest struct {
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsTopicStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicStatusRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusRequest) SetTopic(v string) *OnsTopicStatusRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicStatusRequest) SetInstanceId(v string) *OnsTopicStatusRequest {
	s.InstanceId = &v
	return s
}

type OnsTopicStatusResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsTopicStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsTopicStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicStatusResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusResponseBody) SetRequestId(v string) *OnsTopicStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTopicStatusResponseBody) SetData(v *OnsTopicStatusResponseBodyData) *OnsTopicStatusResponseBody {
	s.Data = v
	return s
}

type OnsTopicStatusResponseBodyData struct {
	Perm          *int32 `json:"Perm,omitempty" xml:"Perm,omitempty"`
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	TotalCount    *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s OnsTopicStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusResponseBodyData) SetPerm(v int32) *OnsTopicStatusResponseBodyData {
	s.Perm = &v
	return s
}

func (s *OnsTopicStatusResponseBodyData) SetLastTimeStamp(v int64) *OnsTopicStatusResponseBodyData {
	s.LastTimeStamp = &v
	return s
}

func (s *OnsTopicStatusResponseBodyData) SetTotalCount(v int64) *OnsTopicStatusResponseBodyData {
	s.TotalCount = &v
	return s
}

type OnsTopicStatusResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTopicStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsTopicSubDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsTopicSubDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicSubDetailResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBody) SetRequestId(v string) *OnsTopicSubDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTopicSubDetailResponseBody) SetData(v *OnsTopicSubDetailResponseBodyData) *OnsTopicSubDetailResponseBody {
	s.Data = v
	return s
}

type OnsTopicSubDetailResponseBodyData struct {
	SubscriptionDataList []*OnsTopicSubDetailResponseBodyDataSubscriptionDataList `json:"SubscriptionDataList,omitempty" xml:"SubscriptionDataList,omitempty" type:"Repeated"`
	Topic                *string                                                  `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicSubDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicSubDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBodyData) SetSubscriptionDataList(v []*OnsTopicSubDetailResponseBodyDataSubscriptionDataList) *OnsTopicSubDetailResponseBodyData {
	s.SubscriptionDataList = v
	return s
}

func (s *OnsTopicSubDetailResponseBodyData) SetTopic(v string) *OnsTopicSubDetailResponseBodyData {
	s.Topic = &v
	return s
}

type OnsTopicSubDetailResponseBodyDataSubscriptionDataList struct {
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MessageModel *string `json:"MessageModel,omitempty" xml:"MessageModel,omitempty"`
	SubString    *string `json:"SubString,omitempty" xml:"SubString,omitempty"`
}

func (s OnsTopicSubDetailResponseBodyDataSubscriptionDataList) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicSubDetailResponseBodyDataSubscriptionDataList) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataList) SetGroupId(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataList {
	s.GroupId = &v
	return s
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataList) SetMessageModel(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataList {
	s.MessageModel = &v
	return s
}

func (s *OnsTopicSubDetailResponseBodyDataSubscriptionDataList) SetSubString(v string) *OnsTopicSubDetailResponseBodyDataSubscriptionDataList {
	s.SubString = &v
	return s
}

type OnsTopicSubDetailResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTopicSubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsTopicSubDetailResponse) SetBody(v *OnsTopicSubDetailResponseBody) *OnsTopicSubDetailResponse {
	s.Body = v
	return s
}

type OnsTopicUpdateRequest struct {
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Perm       *int32  `json:"Perm,omitempty" xml:"Perm,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsTopicUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTopicUpdateRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicUpdateRequest) SetTopic(v string) *OnsTopicUpdateRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicUpdateRequest) SetPerm(v int32) *OnsTopicUpdateRequest {
	s.Perm = &v
	return s
}

func (s *OnsTopicUpdateRequest) SetInstanceId(v string) *OnsTopicUpdateRequest {
	s.InstanceId = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTopicUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Status     *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	MsgKey     *string                                            `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	UpdateTime *int64                                             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime *int64                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Topic      *string                                            `json:"Topic,omitempty" xml:"Topic,omitempty"`
	UserId     *string                                            `json:"UserId,omitempty" xml:"UserId,omitempty"`
	InstanceId *string                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId      *string                                            `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	TraceList  []*OnsTraceGetResultResponseBodyTraceDataTraceList `json:"TraceList,omitempty" xml:"TraceList,omitempty" type:"Repeated"`
	QueryId    *string                                            `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceData) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceData) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetStatus(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.Status = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetMsgKey(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.MsgKey = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetUpdateTime(v int64) *OnsTraceGetResultResponseBodyTraceData {
	s.UpdateTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetCreateTime(v int64) *OnsTraceGetResultResponseBodyTraceData {
	s.CreateTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetTopic(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.Topic = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetUserId(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.UserId = &v
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

func (s *OnsTraceGetResultResponseBodyTraceData) SetTraceList(v []*OnsTraceGetResultResponseBodyTraceDataTraceList) *OnsTraceGetResultResponseBodyTraceData {
	s.TraceList = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceData) SetQueryId(v string) *OnsTraceGetResultResponseBodyTraceData {
	s.QueryId = &v
	return s
}

type OnsTraceGetResultResponseBodyTraceDataTraceList struct {
	Status       *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	MsgKey       *string                                                   `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	PubTime      *int64                                                    `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	SubList      []*OnsTraceGetResultResponseBodyTraceDataTraceListSubList `json:"SubList,omitempty" xml:"SubList,omitempty" type:"Repeated"`
	Topic        *string                                                   `json:"Topic,omitempty" xml:"Topic,omitempty"`
	CostTime     *int32                                                    `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Tag          *string                                                   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	MsgId        *string                                                   `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PubGroupName *string                                                   `json:"PubGroupName,omitempty" xml:"PubGroupName,omitempty"`
	BornHost     *string                                                   `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceList) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceList) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetStatus(v string) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.Status = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetMsgKey(v string) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.MsgKey = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetPubTime(v int64) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.PubTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetSubList(v []*OnsTraceGetResultResponseBodyTraceDataTraceListSubList) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.SubList = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetTopic(v string) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.Topic = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetCostTime(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.CostTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetTag(v string) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.Tag = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetMsgId(v string) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.MsgId = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetPubGroupName(v string) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.PubGroupName = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceList) SetBornHost(v string) *OnsTraceGetResultResponseBodyTraceDataTraceList {
	s.BornHost = &v
	return s
}

type OnsTraceGetResultResponseBodyTraceDataTraceListSubList struct {
	ClientList   []*OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList `json:"ClientList,omitempty" xml:"ClientList,omitempty" type:"Repeated"`
	FailCount    *int32                                                              `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	SubGroupName *string                                                             `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	SuccessCount *int32                                                              `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListSubList) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListSubList) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListSubList) SetClientList(v []*OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList) *OnsTraceGetResultResponseBodyTraceDataTraceListSubList {
	s.ClientList = v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListSubList) SetFailCount(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListSubList {
	s.FailCount = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListSubList) SetSubGroupName(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListSubList {
	s.SubGroupName = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListSubList) SetSuccessCount(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListSubList {
	s.SuccessCount = &v
	return s
}

type OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubTime        *int64  `json:"SubTime,omitempty" xml:"SubTime,omitempty"`
	ReconsumeTimes *int32  `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	SubGroupName   *string `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	ClientHost     *string `json:"ClientHost,omitempty" xml:"ClientHost,omitempty"`
	CostTime       *int32  `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList) SetStatus(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList {
	s.Status = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList) SetSubTime(v int64) *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList {
	s.SubTime = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList) SetReconsumeTimes(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList) SetSubGroupName(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList {
	s.SubGroupName = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList) SetClientHost(v string) *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList {
	s.ClientHost = &v
	return s
}

func (s *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList) SetCostTime(v int32) *OnsTraceGetResultResponseBodyTraceDataTraceListSubListClientList {
	s.CostTime = &v
	return s
}

type OnsTraceGetResultResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTraceGetResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsTraceGetResultResponse) SetBody(v *OnsTraceGetResultResponseBody) *OnsTraceGetResultResponse {
	s.Body = v
	return s
}

type OnsTraceQueryByMsgIdRequest struct {
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsTraceQueryByMsgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceQueryByMsgIdRequest) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgIdRequest) SetTopic(v string) *OnsTraceQueryByMsgIdRequest {
	s.Topic = &v
	return s
}

func (s *OnsTraceQueryByMsgIdRequest) SetMsgId(v string) *OnsTraceQueryByMsgIdRequest {
	s.MsgId = &v
	return s
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

type OnsTraceQueryByMsgIdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	QueryId   *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s OnsTraceQueryByMsgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceQueryByMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgIdResponseBody) SetRequestId(v string) *OnsTraceQueryByMsgIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTraceQueryByMsgIdResponseBody) SetQueryId(v string) *OnsTraceQueryByMsgIdResponseBody {
	s.QueryId = &v
	return s
}

type OnsTraceQueryByMsgIdResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTraceQueryByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsTraceQueryByMsgIdResponse) SetBody(v *OnsTraceQueryByMsgIdResponseBody) *OnsTraceQueryByMsgIdResponse {
	s.Body = v
	return s
}

type OnsTraceQueryByMsgKeyRequest struct {
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	MsgKey     *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsTraceQueryByMsgKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceQueryByMsgKeyRequest) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgKeyRequest) SetTopic(v string) *OnsTraceQueryByMsgKeyRequest {
	s.Topic = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyRequest) SetMsgKey(v string) *OnsTraceQueryByMsgKeyRequest {
	s.MsgKey = &v
	return s
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

type OnsTraceQueryByMsgKeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	QueryId   *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s OnsTraceQueryByMsgKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTraceQueryByMsgKeyResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgKeyResponseBody) SetRequestId(v string) *OnsTraceQueryByMsgKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyResponseBody) SetQueryId(v string) *OnsTraceQueryByMsgKeyResponseBody {
	s.QueryId = &v
	return s
}

type OnsTraceQueryByMsgKeyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTraceQueryByMsgKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsTraceQueryByMsgKeyResponse) SetBody(v *OnsTraceQueryByMsgKeyResponseBody) *OnsTraceQueryByMsgKeyResponse {
	s.Body = v
	return s
}

type OnsTrendGroupOutputTpsRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Period     *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s OnsTrendGroupOutputTpsRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendGroupOutputTpsRequest) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsRequest) SetGroupId(v string) *OnsTrendGroupOutputTpsRequest {
	s.GroupId = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetTopic(v string) *OnsTrendGroupOutputTpsRequest {
	s.Topic = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetBeginTime(v int64) *OnsTrendGroupOutputTpsRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetEndTime(v int64) *OnsTrendGroupOutputTpsRequest {
	s.EndTime = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetType(v int32) *OnsTrendGroupOutputTpsRequest {
	s.Type = &v
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

type OnsTrendGroupOutputTpsResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsTrendGroupOutputTpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsTrendGroupOutputTpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBody) SetRequestId(v string) *OnsTrendGroupOutputTpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBody) SetData(v *OnsTrendGroupOutputTpsResponseBodyData) *OnsTrendGroupOutputTpsResponseBody {
	s.Data = v
	return s
}

type OnsTrendGroupOutputTpsResponseBodyData struct {
	Records []*OnsTrendGroupOutputTpsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	XUnit   *string                                          `json:"XUnit,omitempty" xml:"XUnit,omitempty"`
	YUnit   *string                                          `json:"YUnit,omitempty" xml:"YUnit,omitempty"`
	Title   *string                                          `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s OnsTrendGroupOutputTpsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBodyData) SetRecords(v []*OnsTrendGroupOutputTpsResponseBodyDataRecords) *OnsTrendGroupOutputTpsResponseBodyData {
	s.Records = v
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

func (s *OnsTrendGroupOutputTpsResponseBodyData) SetTitle(v string) *OnsTrendGroupOutputTpsResponseBodyData {
	s.Title = &v
	return s
}

type OnsTrendGroupOutputTpsResponseBodyDataRecords struct {
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *int64   `json:"X,omitempty" xml:"X,omitempty"`
}

func (s OnsTrendGroupOutputTpsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecords) SetY(v float32) *OnsTrendGroupOutputTpsResponseBodyDataRecords {
	s.Y = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponseBodyDataRecords) SetX(v int64) *OnsTrendGroupOutputTpsResponseBodyDataRecords {
	s.X = &v
	return s
}

type OnsTrendGroupOutputTpsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTrendGroupOutputTpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsTrendGroupOutputTpsResponse) SetBody(v *OnsTrendGroupOutputTpsResponseBody) *OnsTrendGroupOutputTpsResponse {
	s.Body = v
	return s
}

type OnsTrendTopicInputTpsRequest struct {
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Period     *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s OnsTrendTopicInputTpsRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendTopicInputTpsRequest) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsRequest) SetTopic(v string) *OnsTrendTopicInputTpsRequest {
	s.Topic = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetBeginTime(v int64) *OnsTrendTopicInputTpsRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetEndTime(v int64) *OnsTrendTopicInputTpsRequest {
	s.EndTime = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetType(v int32) *OnsTrendTopicInputTpsRequest {
	s.Type = &v
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

type OnsTrendTopicInputTpsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *OnsTrendTopicInputTpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s OnsTrendTopicInputTpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBody) SetRequestId(v string) *OnsTrendTopicInputTpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBody) SetData(v *OnsTrendTopicInputTpsResponseBodyData) *OnsTrendTopicInputTpsResponseBody {
	s.Data = v
	return s
}

type OnsTrendTopicInputTpsResponseBodyData struct {
	Records []*OnsTrendTopicInputTpsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	XUnit   *string                                         `json:"XUnit,omitempty" xml:"XUnit,omitempty"`
	YUnit   *string                                         `json:"YUnit,omitempty" xml:"YUnit,omitempty"`
	Title   *string                                         `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s OnsTrendTopicInputTpsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBodyData) SetRecords(v []*OnsTrendTopicInputTpsResponseBodyDataRecords) *OnsTrendTopicInputTpsResponseBodyData {
	s.Records = v
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

func (s *OnsTrendTopicInputTpsResponseBodyData) SetTitle(v string) *OnsTrendTopicInputTpsResponseBodyData {
	s.Title = &v
	return s
}

type OnsTrendTopicInputTpsResponseBodyDataRecords struct {
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *int64   `json:"X,omitempty" xml:"X,omitempty"`
}

func (s OnsTrendTopicInputTpsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecords) SetY(v float32) *OnsTrendTopicInputTpsResponseBodyDataRecords {
	s.Y = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponseBodyDataRecords) SetX(v int64) *OnsTrendTopicInputTpsResponseBodyDataRecords {
	s.X = &v
	return s
}

type OnsTrendTopicInputTpsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsTrendTopicInputTpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OnsTrendTopicInputTpsResponse) SetBody(v *OnsTrendTopicInputTpsResponseBody) *OnsTrendTopicInputTpsResponse {
	s.Body = v
	return s
}

type OnsWarnCreateRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Threshold  *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Contacts   *string `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	DelayTime  *string `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	BlockTime  *string `json:"BlockTime,omitempty" xml:"BlockTime,omitempty"`
	AlertTime  *string `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	Level      *string `json:"Level,omitempty" xml:"Level,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsWarnCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsWarnCreateRequest) GoString() string {
	return s.String()
}

func (s *OnsWarnCreateRequest) SetGroupId(v string) *OnsWarnCreateRequest {
	s.GroupId = &v
	return s
}

func (s *OnsWarnCreateRequest) SetTopic(v string) *OnsWarnCreateRequest {
	s.Topic = &v
	return s
}

func (s *OnsWarnCreateRequest) SetThreshold(v string) *OnsWarnCreateRequest {
	s.Threshold = &v
	return s
}

func (s *OnsWarnCreateRequest) SetContacts(v string) *OnsWarnCreateRequest {
	s.Contacts = &v
	return s
}

func (s *OnsWarnCreateRequest) SetDelayTime(v string) *OnsWarnCreateRequest {
	s.DelayTime = &v
	return s
}

func (s *OnsWarnCreateRequest) SetBlockTime(v string) *OnsWarnCreateRequest {
	s.BlockTime = &v
	return s
}

func (s *OnsWarnCreateRequest) SetAlertTime(v string) *OnsWarnCreateRequest {
	s.AlertTime = &v
	return s
}

func (s *OnsWarnCreateRequest) SetLevel(v string) *OnsWarnCreateRequest {
	s.Level = &v
	return s
}

func (s *OnsWarnCreateRequest) SetInstanceId(v string) *OnsWarnCreateRequest {
	s.InstanceId = &v
	return s
}

type OnsWarnCreateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsWarnCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsWarnCreateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsWarnCreateResponseBody) SetRequestId(v string) *OnsWarnCreateResponseBody {
	s.RequestId = &v
	return s
}

type OnsWarnCreateResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsWarnCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsWarnCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsWarnCreateResponse) GoString() string {
	return s.String()
}

func (s *OnsWarnCreateResponse) SetHeaders(v map[string]*string) *OnsWarnCreateResponse {
	s.Headers = v
	return s
}

func (s *OnsWarnCreateResponse) SetBody(v *OnsWarnCreateResponseBody) *OnsWarnCreateResponse {
	s.Body = v
	return s
}

type OnsWarnDeleteRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnsWarnDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s OnsWarnDeleteRequest) GoString() string {
	return s.String()
}

func (s *OnsWarnDeleteRequest) SetGroupId(v string) *OnsWarnDeleteRequest {
	s.GroupId = &v
	return s
}

func (s *OnsWarnDeleteRequest) SetTopic(v string) *OnsWarnDeleteRequest {
	s.Topic = &v
	return s
}

func (s *OnsWarnDeleteRequest) SetInstanceId(v string) *OnsWarnDeleteRequest {
	s.InstanceId = &v
	return s
}

type OnsWarnDeleteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsWarnDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OnsWarnDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *OnsWarnDeleteResponseBody) SetRequestId(v string) *OnsWarnDeleteResponseBody {
	s.RequestId = &v
	return s
}

type OnsWarnDeleteResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OnsWarnDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OnsWarnDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s OnsWarnDeleteResponse) GoString() string {
	return s.String()
}

func (s *OnsWarnDeleteResponse) SetHeaders(v map[string]*string) *OnsWarnDeleteResponse {
	s.Headers = v
	return s
}

func (s *OnsWarnDeleteResponse) SetBody(v *OnsWarnDeleteResponseBody) *OnsWarnDeleteResponse {
	s.Body = v
	return s
}

type OpenOnsServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s OpenOnsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenOnsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenOnsServiceResponseBody) SetRequestId(v string) *OpenOnsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenOnsServiceResponseBody) SetOrderId(v string) *OpenOnsServiceResponseBody {
	s.OrderId = &v
	return s
}

type OpenOnsServiceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenOnsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *OpenOnsServiceResponse) SetBody(v *OpenOnsServiceResponseBody) *OpenOnsServiceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	InstanceId   *string                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
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

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetInstanceId(v string) *UntagResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsConsumerAccumulateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsConsumerAccumulate"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsConsumerGetConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsConsumerGetConnection"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsConsumerResetOffsetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsConsumerResetOffset"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsConsumerStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsConsumerStatus"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsConsumerTimeSpanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsConsumerTimeSpan"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsDLQMessageGetByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsDLQMessageGetById"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsDLQMessagePageQueryByGroupIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsDLQMessagePageQueryByGroupId"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsDLQMessageResendByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsDLQMessageResendById"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsGroupConsumerUpdateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsGroupConsumerUpdate"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsGroupCreateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsGroupCreate"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsGroupDeleteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsGroupDelete"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsGroupListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsGroupList"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsGroupSubDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsGroupSubDetail"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsInstanceBaseInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsInstanceBaseInfo"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsInstanceCreateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsInstanceCreate"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsInstanceDeleteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsInstanceDelete"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsInstanceInServiceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsInstanceInServiceList"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsInstanceUpdateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsInstanceUpdate"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMessageGetByKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMessageGetByKey"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMessageGetByMsgIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMessageGetByMsgId"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMessagePageQueryByTopicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMessagePageQueryByTopic"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMessagePushResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMessagePush"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) OnsMessageSendWithOptions(request *OnsMessageSendRequest, runtime *util.RuntimeOptions) (_result *OnsMessageSendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMessageSendResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMessageSend"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsMessageSend(request *OnsMessageSendRequest) (_result *OnsMessageSendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMessageSendResponse{}
	_body, _err := client.OnsMessageSendWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMessageTraceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMessageTrace"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) OnsMqttGroupIdCreateWithOptions(request *OnsMqttGroupIdCreateRequest, runtime *util.RuntimeOptions) (_result *OnsMqttGroupIdCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMqttGroupIdCreateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMqttGroupIdCreate"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsMqttGroupIdCreate(request *OnsMqttGroupIdCreateRequest) (_result *OnsMqttGroupIdCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMqttGroupIdCreateResponse{}
	_body, _err := client.OnsMqttGroupIdCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OnsMqttGroupIdDeleteWithOptions(request *OnsMqttGroupIdDeleteRequest, runtime *util.RuntimeOptions) (_result *OnsMqttGroupIdDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMqttGroupIdDeleteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMqttGroupIdDelete"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsMqttGroupIdDelete(request *OnsMqttGroupIdDeleteRequest) (_result *OnsMqttGroupIdDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMqttGroupIdDeleteResponse{}
	_body, _err := client.OnsMqttGroupIdDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OnsMqttGroupIdListWithOptions(request *OnsMqttGroupIdListRequest, runtime *util.RuntimeOptions) (_result *OnsMqttGroupIdListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMqttGroupIdListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMqttGroupIdList"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsMqttGroupIdList(request *OnsMqttGroupIdListRequest) (_result *OnsMqttGroupIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMqttGroupIdListResponse{}
	_body, _err := client.OnsMqttGroupIdListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OnsMqttQueryClientByClientIdWithOptions(request *OnsMqttQueryClientByClientIdRequest, runtime *util.RuntimeOptions) (_result *OnsMqttQueryClientByClientIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMqttQueryClientByClientIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMqttQueryClientByClientId"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsMqttQueryClientByClientId(request *OnsMqttQueryClientByClientIdRequest) (_result *OnsMqttQueryClientByClientIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMqttQueryClientByClientIdResponse{}
	_body, _err := client.OnsMqttQueryClientByClientIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OnsMqttQueryClientByGroupIdWithOptions(request *OnsMqttQueryClientByGroupIdRequest, runtime *util.RuntimeOptions) (_result *OnsMqttQueryClientByGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMqttQueryClientByGroupIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMqttQueryClientByGroupId"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsMqttQueryClientByGroupId(request *OnsMqttQueryClientByGroupIdRequest) (_result *OnsMqttQueryClientByGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMqttQueryClientByGroupIdResponse{}
	_body, _err := client.OnsMqttQueryClientByGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OnsMqttQueryClientByTopicWithOptions(request *OnsMqttQueryClientByTopicRequest, runtime *util.RuntimeOptions) (_result *OnsMqttQueryClientByTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMqttQueryClientByTopicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMqttQueryClientByTopic"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsMqttQueryClientByTopic(request *OnsMqttQueryClientByTopicRequest) (_result *OnsMqttQueryClientByTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMqttQueryClientByTopicResponse{}
	_body, _err := client.OnsMqttQueryClientByTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OnsMqttQueryHistoryOnlineWithOptions(request *OnsMqttQueryHistoryOnlineRequest, runtime *util.RuntimeOptions) (_result *OnsMqttQueryHistoryOnlineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMqttQueryHistoryOnlineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMqttQueryHistoryOnline"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsMqttQueryHistoryOnline(request *OnsMqttQueryHistoryOnlineRequest) (_result *OnsMqttQueryHistoryOnlineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMqttQueryHistoryOnlineResponse{}
	_body, _err := client.OnsMqttQueryHistoryOnlineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OnsMqttQueryMsgTransTrendWithOptions(request *OnsMqttQueryMsgTransTrendRequest, runtime *util.RuntimeOptions) (_result *OnsMqttQueryMsgTransTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsMqttQueryMsgTransTrendResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsMqttQueryMsgTransTrend"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsMqttQueryMsgTransTrend(request *OnsMqttQueryMsgTransTrendRequest) (_result *OnsMqttQueryMsgTransTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsMqttQueryMsgTransTrendResponse{}
	_body, _err := client.OnsMqttQueryMsgTransTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OnsRegionListWithOptions(runtime *util.RuntimeOptions) (_result *OnsRegionListResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &OnsRegionListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsRegionList"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTopicCreateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTopicCreate"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTopicDeleteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTopicDelete"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTopicListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTopicList"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTopicStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTopicStatus"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTopicSubDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTopicSubDetail"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTopicUpdateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTopicUpdate"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTraceGetResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTraceGetResult"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTraceQueryByMsgIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTraceQueryByMsgId"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTraceQueryByMsgKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTraceQueryByMsgKey"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTrendGroupOutputTpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTrendGroupOutputTps"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsTrendTopicInputTpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsTrendTopicInputTps"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) OnsWarnCreateWithOptions(request *OnsWarnCreateRequest, runtime *util.RuntimeOptions) (_result *OnsWarnCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsWarnCreateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsWarnCreate"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsWarnCreate(request *OnsWarnCreateRequest) (_result *OnsWarnCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsWarnCreateResponse{}
	_body, _err := client.OnsWarnCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OnsWarnDeleteWithOptions(request *OnsWarnDeleteRequest, runtime *util.RuntimeOptions) (_result *OnsWarnDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OnsWarnDeleteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OnsWarnDelete"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OnsWarnDelete(request *OnsWarnDeleteRequest) (_result *OnsWarnDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OnsWarnDeleteResponse{}
	_body, _err := client.OnsWarnDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenOnsServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenOnsServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &OpenOnsServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenOnsService"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2019-02-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

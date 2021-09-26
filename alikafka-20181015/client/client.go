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

type CreateConsumerGroupRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) SetInstanceId(v string) *CreateConsumerGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetConsumerId(v string) *CreateConsumerGroupRequest {
	s.ConsumerId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRegionId(v string) *CreateConsumerGroupRequest {
	s.RegionId = &v
	return s
}

type CreateConsumerGroupResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBody) SetCode(v int32) *CreateConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetMessage(v string) *CreateConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetRequestId(v string) *CreateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetSuccess(v bool) *CreateConsumerGroupResponseBody {
	s.Success = &v
	return s
}

type CreateConsumerGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponse) SetHeaders(v map[string]*string) *CreateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerGroupResponse) SetBody(v *CreateConsumerGroupResponseBody) *CreateConsumerGroupResponse {
	s.Body = v
	return s
}

type CreateTopicRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) SetInstanceId(v string) *CreateTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTopicRequest) SetTopic(v string) *CreateTopicRequest {
	s.Topic = &v
	return s
}

func (s *CreateTopicRequest) SetRemark(v string) *CreateTopicRequest {
	s.Remark = &v
	return s
}

func (s *CreateTopicRequest) SetRegionId(v string) *CreateTopicRequest {
	s.RegionId = &v
	return s
}

type CreateTopicResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) SetCode(v int32) *CreateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBody) SetMessage(v string) *CreateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBody) SetRequestId(v string) *CreateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTopicResponseBody) SetSuccess(v bool) *CreateTopicResponseBody {
	s.Success = &v
	return s
}

type CreateTopicResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponse) GoString() string {
	return s.String()
}

func (s *CreateTopicResponse) SetHeaders(v map[string]*string) *CreateTopicResponse {
	s.Headers = v
	return s
}

func (s *CreateTopicResponse) SetBody(v *CreateTopicResponseBody) *CreateTopicResponse {
	s.Body = v
	return s
}

type DeleteConsumerGroupRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupRequest) SetInstanceId(v string) *DeleteConsumerGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetConsumerId(v string) *DeleteConsumerGroupRequest {
	s.ConsumerId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetRegionId(v string) *DeleteConsumerGroupRequest {
	s.RegionId = &v
	return s
}

type DeleteConsumerGroupResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponseBody) SetCode(v int32) *DeleteConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetMessage(v string) *DeleteConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetRequestId(v string) *DeleteConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetSuccess(v bool) *DeleteConsumerGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteConsumerGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponse) SetHeaders(v map[string]*string) *DeleteConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerGroupResponse) SetBody(v *DeleteConsumerGroupResponseBody) *DeleteConsumerGroupResponse {
	s.Body = v
	return s
}

type DeleteTopicRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteTopicRequest) SetInstanceId(v string) *DeleteTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTopicRequest) SetTopic(v string) *DeleteTopicRequest {
	s.Topic = &v
	return s
}

func (s *DeleteTopicRequest) SetRegionId(v string) *DeleteTopicRequest {
	s.RegionId = &v
	return s
}

type DeleteTopicResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) SetCode(v int32) *DeleteTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTopicResponseBody) SetMessage(v string) *DeleteTopicResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTopicResponseBody) SetRequestId(v string) *DeleteTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTopicResponseBody) SetSuccess(v bool) *DeleteTopicResponseBody {
	s.Success = &v
	return s
}

type DeleteTopicResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponse) SetHeaders(v map[string]*string) *DeleteTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteTopicResponse) SetBody(v *DeleteTopicResponseBody) *DeleteTopicResponse {
	s.Body = v
	return s
}

type GetConsumerListRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetConsumerListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerListRequest) SetInstanceId(v string) *GetConsumerListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerListRequest) SetRegionId(v string) *GetConsumerListRequest {
	s.RegionId = &v
	return s
}

type GetConsumerListResponseBody struct {
	Code         *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	ConsumerList *GetConsumerListResponseBodyConsumerList `json:"ConsumerList,omitempty" xml:"ConsumerList,omitempty" type:"Struct"`
}

func (s GetConsumerListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBody) SetCode(v int32) *GetConsumerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerListResponseBody) SetMessage(v string) *GetConsumerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerListResponseBody) SetRequestId(v string) *GetConsumerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerListResponseBody) SetSuccess(v bool) *GetConsumerListResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsumerListResponseBody) SetConsumerList(v *GetConsumerListResponseBodyConsumerList) *GetConsumerListResponseBody {
	s.ConsumerList = v
	return s
}

type GetConsumerListResponseBodyConsumerList struct {
	ConsumerVO []*GetConsumerListResponseBodyConsumerListConsumerVO `json:"ConsumerVO,omitempty" xml:"ConsumerVO,omitempty" type:"Repeated"`
}

func (s GetConsumerListResponseBodyConsumerList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerList) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBodyConsumerList) SetConsumerVO(v []*GetConsumerListResponseBodyConsumerListConsumerVO) *GetConsumerListResponseBodyConsumerList {
	s.ConsumerVO = v
	return s
}

type GetConsumerListResponseBodyConsumerListConsumerVO struct {
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetConsumerListResponseBodyConsumerListConsumerVO) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerListConsumerVO) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetConsumerId(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.ConsumerId = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetInstanceId(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetRegionId(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.RegionId = &v
	return s
}

type GetConsumerListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConsumerListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConsumerListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponse) SetHeaders(v map[string]*string) *GetConsumerListResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerListResponse) SetBody(v *GetConsumerListResponseBody) *GetConsumerListResponse {
	s.Body = v
	return s
}

type GetConsumerProgressRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetConsumerProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressRequest) SetInstanceId(v string) *GetConsumerProgressRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerProgressRequest) SetConsumerId(v string) *GetConsumerProgressRequest {
	s.ConsumerId = &v
	return s
}

func (s *GetConsumerProgressRequest) SetRegionId(v string) *GetConsumerProgressRequest {
	s.RegionId = &v
	return s
}

type GetConsumerProgressResponseBody struct {
	Code             *int32                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	ConsumerProgress *GetConsumerProgressResponseBodyConsumerProgress `json:"ConsumerProgress,omitempty" xml:"ConsumerProgress,omitempty" type:"Struct"`
}

func (s GetConsumerProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBody) SetCode(v int32) *GetConsumerProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerProgressResponseBody) SetMessage(v string) *GetConsumerProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerProgressResponseBody) SetRequestId(v string) *GetConsumerProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerProgressResponseBody) SetSuccess(v bool) *GetConsumerProgressResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsumerProgressResponseBody) SetConsumerProgress(v *GetConsumerProgressResponseBodyConsumerProgress) *GetConsumerProgressResponseBody {
	s.ConsumerProgress = v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgress struct {
	LastTimestamp *int64                                                    `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	TotalDiff     *int64                                                    `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
	TopicList     *GetConsumerProgressResponseBodyConsumerProgressTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Struct"`
}

func (s GetConsumerProgressResponseBodyConsumerProgress) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgress) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetLastTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgress {
	s.LastTimestamp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetTotalDiff(v int64) *GetConsumerProgressResponseBodyConsumerProgress {
	s.TotalDiff = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetTopicList(v *GetConsumerProgressResponseBodyConsumerProgressTopicList) *GetConsumerProgressResponseBodyConsumerProgress {
	s.TopicList = v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgressTopicList struct {
	TopicList []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Repeated"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicList) SetTopicList(v []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) *GetConsumerProgressResponseBodyConsumerProgressTopicList {
	s.TopicList = v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList struct {
	TotalDiff     *int64                                                                       `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
	LastTimestamp *int64                                                                       `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Topic         *string                                                                      `json:"Topic,omitempty" xml:"Topic,omitempty"`
	OffsetList    *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList `json:"OffsetList,omitempty" xml:"OffsetList,omitempty" type:"Struct"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetTotalDiff(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.TotalDiff = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetLastTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.LastTimestamp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetTopic(v string) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.Topic = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetOffsetList(v *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.OffsetList = v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList struct {
	OffsetList []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList `json:"OffsetList,omitempty" xml:"OffsetList,omitempty" type:"Repeated"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) SetOffsetList(v []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList {
	s.OffsetList = v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList struct {
	Partition      *int32 `json:"Partition,omitempty" xml:"Partition,omitempty"`
	BrokerOffset   *int64 `json:"BrokerOffset,omitempty" xml:"BrokerOffset,omitempty"`
	ConsumerOffset *int64 `json:"ConsumerOffset,omitempty" xml:"ConsumerOffset,omitempty"`
	LastTimestamp  *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetPartition(v int32) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.Partition = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetBrokerOffset(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.BrokerOffset = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetConsumerOffset(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.ConsumerOffset = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetLastTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.LastTimestamp = &v
	return s
}

type GetConsumerProgressResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConsumerProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConsumerProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponse) SetHeaders(v map[string]*string) *GetConsumerProgressResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerProgressResponse) SetBody(v *GetConsumerProgressResponseBody) *GetConsumerProgressResponse {
	s.Body = v
	return s
}

type GetInstanceListRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceListRequest) SetRegionId(v string) *GetInstanceListRequest {
	s.RegionId = &v
	return s
}

type GetInstanceListResponseBody struct {
	Code         *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	InstanceList *GetInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
}

func (s GetInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBody) SetCode(v int32) *GetInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceListResponseBody) SetMessage(v string) *GetInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceListResponseBody) SetRequestId(v string) *GetInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceListResponseBody) SetSuccess(v bool) *GetInstanceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceListResponseBody) SetInstanceList(v *GetInstanceListResponseBodyInstanceList) *GetInstanceListResponseBody {
	s.InstanceList = v
	return s
}

type GetInstanceListResponseBodyInstanceList struct {
	InstanceVO []*GetInstanceListResponseBodyInstanceListInstanceVO `json:"InstanceVO,omitempty" xml:"InstanceVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceList) SetInstanceVO(v []*GetInstanceListResponseBodyInstanceListInstanceVO) *GetInstanceListResponseBodyInstanceList {
	s.InstanceVO = v
	return s
}

type GetInstanceListResponseBodyInstanceListInstanceVO struct {
	VpcId                    *string                                                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId                *string                                                                    `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ExpiredTime              *int64                                                                     `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	DeployType               *int32                                                                     `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	CreateTime               *int64                                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SslEndPoint              *string                                                                    `json:"SslEndPoint,omitempty" xml:"SslEndPoint,omitempty"`
	InstanceId               *string                                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name                     *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceStatus            *int32                                                                     `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	EndPoint                 *string                                                                    `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	RegionId                 *string                                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UpgradeServiceDetailInfo *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo `json:"UpgradeServiceDetailInfo,omitempty" xml:"UpgradeServiceDetailInfo,omitempty" type:"Struct"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVO) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVSwitchId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetExpiredTime(v int64) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ExpiredTime = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDeployType(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DeployType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetCreateTime(v int64) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSslEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SslEndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetInstanceId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetName(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.Name = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetServiceStatus(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ServiceStatus = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.EndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetRegionId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.RegionId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUpgradeServiceDetailInfo(v *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UpgradeServiceDetailInfo = v
	return s
}

type GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo struct {
	UpgradeServiceDetailInfoVO []*GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfoUpgradeServiceDetailInfoVO `json:"UpgradeServiceDetailInfoVO,omitempty" xml:"UpgradeServiceDetailInfoVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) SetUpgradeServiceDetailInfoVO(v []*GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfoUpgradeServiceDetailInfoVO) *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo {
	s.UpgradeServiceDetailInfoVO = v
	return s
}

type GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfoUpgradeServiceDetailInfoVO struct {
	Current2OpenSourceVersion *string `json:"Current2OpenSourceVersion,omitempty" xml:"Current2OpenSourceVersion,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfoUpgradeServiceDetailInfoVO) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfoUpgradeServiceDetailInfoVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfoUpgradeServiceDetailInfoVO) SetCurrent2OpenSourceVersion(v string) *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfoUpgradeServiceDetailInfoVO {
	s.Current2OpenSourceVersion = &v
	return s
}

type GetInstanceListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponse) SetHeaders(v map[string]*string) *GetInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceListResponse) SetBody(v *GetInstanceListResponseBody) *GetInstanceListResponse {
	s.Body = v
	return s
}

type GetTopicListRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTopicListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListRequest) GoString() string {
	return s.String()
}

func (s *GetTopicListRequest) SetInstanceId(v string) *GetTopicListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTopicListRequest) SetCurrentPage(v string) *GetTopicListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetTopicListRequest) SetPageSize(v string) *GetTopicListRequest {
	s.PageSize = &v
	return s
}

func (s *GetTopicListRequest) SetRegionId(v string) *GetTopicListRequest {
	s.RegionId = &v
	return s
}

type GetTopicListResponseBody struct {
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Code        *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize    *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32                             `json:"Total,omitempty" xml:"Total,omitempty"`
	TopicList   *GetTopicListResponseBodyTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Struct"`
}

func (s GetTopicListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBody) SetCurrentPage(v int32) *GetTopicListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetTopicListResponseBody) SetRequestId(v string) *GetTopicListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicListResponseBody) SetSuccess(v bool) *GetTopicListResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicListResponseBody) SetCode(v int32) *GetTopicListResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicListResponseBody) SetMessage(v string) *GetTopicListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicListResponseBody) SetPageSize(v int32) *GetTopicListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTopicListResponseBody) SetTotal(v int32) *GetTopicListResponseBody {
	s.Total = &v
	return s
}

func (s *GetTopicListResponseBody) SetTopicList(v *GetTopicListResponseBodyTopicList) *GetTopicListResponseBody {
	s.TopicList = v
	return s
}

type GetTopicListResponseBodyTopicList struct {
	TopicVO []*GetTopicListResponseBodyTopicListTopicVO `json:"TopicVO,omitempty" xml:"TopicVO,omitempty" type:"Repeated"`
}

func (s GetTopicListResponseBodyTopicList) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponseBodyTopicList) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBodyTopicList) SetTopicVO(v []*GetTopicListResponseBodyTopicListTopicVO) *GetTopicListResponseBodyTopicList {
	s.TopicVO = v
	return s
}

type GetTopicListResponseBodyTopicListTopicVO struct {
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTopicListResponseBodyTopicListTopicVO) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponseBodyTopicListTopicVO) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetStatus(v int32) *GetTopicListResponseBodyTopicListTopicVO {
	s.Status = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetRemark(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.Remark = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetCreateTime(v int64) *GetTopicListResponseBodyTopicListTopicVO {
	s.CreateTime = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetTopic(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.Topic = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetStatusName(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.StatusName = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetInstanceId(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.InstanceId = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetRegionId(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.RegionId = &v
	return s
}

type GetTopicListResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTopicListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTopicListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponse) GoString() string {
	return s.String()
}

func (s *GetTopicListResponse) SetHeaders(v map[string]*string) *GetTopicListResponse {
	s.Headers = v
	return s
}

func (s *GetTopicListResponse) SetBody(v *GetTopicListResponseBody) *GetTopicListResponse {
	s.Body = v
	return s
}

type GetTopicStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTopicStatusRequest) SetInstanceId(v string) *GetTopicStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTopicStatusRequest) SetTopic(v string) *GetTopicStatusRequest {
	s.Topic = &v
	return s
}

type GetTopicStatusResponseBody struct {
	Code        *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TopicStatus *GetTopicStatusResponseBodyTopicStatus `json:"TopicStatus,omitempty" xml:"TopicStatus,omitempty" type:"Struct"`
}

func (s GetTopicStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBody) SetCode(v int32) *GetTopicStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetMessage(v string) *GetTopicStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetRequestId(v string) *GetTopicStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetSuccess(v bool) *GetTopicStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetTopicStatus(v *GetTopicStatusResponseBodyTopicStatus) *GetTopicStatusResponseBody {
	s.TopicStatus = v
	return s
}

type GetTopicStatusResponseBodyTopicStatus struct {
	TotalCount    *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LastTimeStamp *int64                                            `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	OffsetTable   *GetTopicStatusResponseBodyTopicStatusOffsetTable `json:"OffsetTable,omitempty" xml:"OffsetTable,omitempty" type:"Struct"`
}

func (s GetTopicStatusResponseBodyTopicStatus) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusResponseBodyTopicStatus) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBodyTopicStatus) SetTotalCount(v int64) *GetTopicStatusResponseBodyTopicStatus {
	s.TotalCount = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatus) SetLastTimeStamp(v int64) *GetTopicStatusResponseBodyTopicStatus {
	s.LastTimeStamp = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatus) SetOffsetTable(v *GetTopicStatusResponseBodyTopicStatusOffsetTable) *GetTopicStatusResponseBodyTopicStatus {
	s.OffsetTable = v
	return s
}

type GetTopicStatusResponseBodyTopicStatusOffsetTable struct {
	OffsetTable []*GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable `json:"OffsetTable,omitempty" xml:"OffsetTable,omitempty" type:"Repeated"`
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTable) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTable) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTable) SetOffsetTable(v []*GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) *GetTopicStatusResponseBodyTopicStatusOffsetTable {
	s.OffsetTable = v
	return s
}

type GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable struct {
	Partition           *int32  `json:"Partition,omitempty" xml:"Partition,omitempty"`
	MinOffset           *int64  `json:"MinOffset,omitempty" xml:"MinOffset,omitempty"`
	LastUpdateTimestamp *int64  `json:"LastUpdateTimestamp,omitempty" xml:"LastUpdateTimestamp,omitempty"`
	MaxOffset           *int64  `json:"MaxOffset,omitempty" xml:"MaxOffset,omitempty"`
	Topic               *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetPartition(v int32) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.Partition = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetMinOffset(v int64) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.MinOffset = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetLastUpdateTimestamp(v int64) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.LastUpdateTimestamp = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetMaxOffset(v int64) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.MaxOffset = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetTopic(v string) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.Topic = &v
	return s
}

type GetTopicStatusResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTopicStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTopicStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponse) SetHeaders(v map[string]*string) *GetTopicStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTopicStatusResponse) SetBody(v *GetTopicStatusResponseBody) *GetTopicStatusResponse {
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
		"ap-northeast-2-pop":          tea.String("alikafka.aliyuncs.com"),
		"ap-southeast-2":              tea.String("alikafka.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("alikafka.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("alikafka.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("alikafka.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("alikafka.aliyuncs.com"),
		"cn-edge-1":                   tea.String("alikafka.aliyuncs.com"),
		"cn-fujian":                   tea.String("alikafka.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("alikafka.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("alikafka.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("alikafka.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("alikafka.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("alikafka.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("alikafka.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("alikafka.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("alikafka.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("alikafka.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("alikafka.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("alikafka.aliyuncs.com"),
		"cn-wuhan":                    tea.String("alikafka.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("alikafka.aliyuncs.com"),
		"cn-yushanfang":               tea.String("alikafka.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("alikafka.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("alikafka.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("alikafka.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("alikafka.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("alikafka.aliyuncs.com"),
		"me-east-1":                   tea.String("alikafka.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("alikafka.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("alikafka"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateConsumerGroupWithOptions(request *CreateConsumerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateConsumerGroup"), tea.String("2018-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CreateConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTopicWithOptions(request *CreateTopicRequest, runtime *util.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTopicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTopic"), tea.String("2018-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTopic(request *CreateTopicRequest) (_result *CreateTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTopicResponse{}
	_body, _err := client.CreateTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConsumerGroupWithOptions(request *DeleteConsumerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteConsumerGroup"), tea.String("2018-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (_result *DeleteConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DeleteConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTopicWithOptions(request *DeleteTopicRequest, runtime *util.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTopicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTopic"), tea.String("2018-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTopic(request *DeleteTopicRequest) (_result *DeleteTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTopicResponse{}
	_body, _err := client.DeleteTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConsumerListWithOptions(request *GetConsumerListRequest, runtime *util.RuntimeOptions) (_result *GetConsumerListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetConsumerListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConsumerList"), tea.String("2018-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConsumerList(request *GetConsumerListRequest) (_result *GetConsumerListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConsumerListResponse{}
	_body, _err := client.GetConsumerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConsumerProgressWithOptions(request *GetConsumerProgressRequest, runtime *util.RuntimeOptions) (_result *GetConsumerProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetConsumerProgressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConsumerProgress"), tea.String("2018-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConsumerProgress(request *GetConsumerProgressRequest) (_result *GetConsumerProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConsumerProgressResponse{}
	_body, _err := client.GetConsumerProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceListWithOptions(request *GetInstanceListRequest, runtime *util.RuntimeOptions) (_result *GetInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetInstanceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetInstanceList"), tea.String("2018-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceList(request *GetInstanceListRequest) (_result *GetInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceListResponse{}
	_body, _err := client.GetInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTopicListWithOptions(request *GetTopicListRequest, runtime *util.RuntimeOptions) (_result *GetTopicListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTopicListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTopicList"), tea.String("2018-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTopicList(request *GetTopicListRequest) (_result *GetTopicListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTopicListResponse{}
	_body, _err := client.GetTopicListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTopicStatusWithOptions(request *GetTopicStatusRequest, runtime *util.RuntimeOptions) (_result *GetTopicStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTopicStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTopicStatus"), tea.String("2018-10-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTopicStatus(request *GetTopicStatusRequest) (_result *GetTopicStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTopicStatusResponse{}
	_body, _err := client.GetTopicStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

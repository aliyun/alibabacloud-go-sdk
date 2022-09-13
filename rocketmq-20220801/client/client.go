// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateConsumerGroupRequest struct {
	ConsumeRetryPolicy *CreateConsumerGroupRequestConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	DeliveryOrderType  *string                                       `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	Remark             *string                                       `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) SetConsumeRetryPolicy(v *CreateConsumerGroupRequestConsumeRetryPolicy) *CreateConsumerGroupRequest {
	s.ConsumeRetryPolicy = v
	return s
}

func (s *CreateConsumerGroupRequest) SetDeliveryOrderType(v string) *CreateConsumerGroupRequest {
	s.DeliveryOrderType = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRemark(v string) *CreateConsumerGroupRequest {
	s.Remark = &v
	return s
}

type CreateConsumerGroupRequestConsumeRetryPolicy struct {
	MaxRetryTimes *int32  `json:"maxRetryTimes,omitempty" xml:"maxRetryTimes,omitempty"`
	RetryPolicy   *string `json:"retryPolicy,omitempty" xml:"retryPolicy,omitempty"`
}

func (s CreateConsumerGroupRequestConsumeRetryPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupRequestConsumeRetryPolicy) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) SetMaxRetryTimes(v int32) *CreateConsumerGroupRequestConsumeRetryPolicy {
	s.MaxRetryTimes = &v
	return s
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) SetRetryPolicy(v string) *CreateConsumerGroupRequestConsumeRetryPolicy {
	s.RetryPolicy = &v
	return s
}

type CreateConsumerGroupResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *bool   `json:"data,omitempty" xml:"data,omitempty"`
	DynamicCode    *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBody) SetCode(v string) *CreateConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetData(v bool) *CreateConsumerGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetDynamicCode(v string) *CreateConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetDynamicMessage(v string) *CreateConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetHttpStatusCode(v int32) *CreateConsumerGroupResponseBody {
	s.HttpStatusCode = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateConsumerGroupResponse) SetStatusCode(v int32) *CreateConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConsumerGroupResponse) SetBody(v *CreateConsumerGroupResponseBody) *CreateConsumerGroupResponse {
	s.Body = v
	return s
}

type CreateTopicRequest struct {
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	Remark      *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) SetMessageType(v string) *CreateTopicRequest {
	s.MessageType = &v
	return s
}

func (s *CreateTopicRequest) SetRemark(v string) *CreateTopicRequest {
	s.Remark = &v
	return s
}

type CreateTopicResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *bool   `json:"data,omitempty" xml:"data,omitempty"`
	DynamicCode    *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) SetCode(v string) *CreateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBody) SetData(v bool) *CreateTopicResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTopicResponseBody) SetDynamicCode(v string) *CreateTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateTopicResponseBody) SetDynamicMessage(v string) *CreateTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateTopicResponseBody) SetHttpStatusCode(v int32) *CreateTopicResponseBody {
	s.HttpStatusCode = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateTopicResponse) SetStatusCode(v int32) *CreateTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTopicResponse) SetBody(v *CreateTopicResponseBody) *CreateTopicResponse {
	s.Body = v
	return s
}

type DeleteConsumerGroupResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *bool   `json:"data,omitempty" xml:"data,omitempty"`
	DynamicCode    *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponseBody) SetCode(v string) *DeleteConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetData(v bool) *DeleteConsumerGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetDynamicCode(v string) *DeleteConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetDynamicMessage(v string) *DeleteConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetHttpStatusCode(v int32) *DeleteConsumerGroupResponseBody {
	s.HttpStatusCode = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteConsumerGroupResponse) SetStatusCode(v int32) *DeleteConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConsumerGroupResponse) SetBody(v *DeleteConsumerGroupResponseBody) *DeleteConsumerGroupResponse {
	s.Body = v
	return s
}

type DeleteInstanceResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *bool   `json:"data,omitempty" xml:"data,omitempty"`
	DynamicCode    *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetCode(v string) *DeleteInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetData(v bool) *DeleteInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetDynamicCode(v string) *DeleteInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetDynamicMessage(v string) *DeleteInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetMessage(v string) *DeleteInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetSuccess(v bool) *DeleteInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteTopicResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *bool   `json:"data,omitempty" xml:"data,omitempty"`
	DynamicCode    *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) SetCode(v string) *DeleteTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTopicResponseBody) SetData(v bool) *DeleteTopicResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTopicResponseBody) SetDynamicCode(v string) *DeleteTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteTopicResponseBody) SetDynamicMessage(v string) *DeleteTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteTopicResponseBody) SetHttpStatusCode(v int32) *DeleteTopicResponseBody {
	s.HttpStatusCode = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteTopicResponse) SetStatusCode(v int32) *DeleteTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTopicResponse) SetBody(v *DeleteTopicResponseBody) *DeleteTopicResponse {
	s.Body = v
	return s
}

type GetConsumerGroupResponseBody struct {
	Code           *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data           *GetConsumerGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	DynamicCode    *string                           `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string                           `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32                            `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                           `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBody) SetCode(v string) *GetConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetData(v *GetConsumerGroupResponseBodyData) *GetConsumerGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerGroupResponseBody) SetDynamicCode(v string) *GetConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetDynamicMessage(v string) *GetConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetHttpStatusCode(v int32) *GetConsumerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetMessage(v string) *GetConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetRequestId(v string) *GetConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetSuccess(v bool) *GetConsumerGroupResponseBody {
	s.Success = &v
	return s
}

type GetConsumerGroupResponseBodyData struct {
	ConsumeRetryPolicy *GetConsumerGroupResponseBodyDataConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	ConsumerGroupId    *string                                             `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	CreateTime         *string                                             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DeliveryOrderType  *string                                             `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	InstanceId         *string                                             `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	RegionId           *string                                             `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Remark             *string                                             `json:"remark,omitempty" xml:"remark,omitempty"`
	Status             *string                                             `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTime         *string                                             `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetConsumerGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBodyData) SetConsumeRetryPolicy(v *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) *GetConsumerGroupResponseBodyData {
	s.ConsumeRetryPolicy = v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetConsumerGroupId(v string) *GetConsumerGroupResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetCreateTime(v string) *GetConsumerGroupResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetDeliveryOrderType(v string) *GetConsumerGroupResponseBodyData {
	s.DeliveryOrderType = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetInstanceId(v string) *GetConsumerGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetRegionId(v string) *GetConsumerGroupResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetRemark(v string) *GetConsumerGroupResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetStatus(v string) *GetConsumerGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetUpdateTime(v string) *GetConsumerGroupResponseBodyData {
	s.UpdateTime = &v
	return s
}

type GetConsumerGroupResponseBodyDataConsumeRetryPolicy struct {
	MaxRetryTimes *int32  `json:"maxRetryTimes,omitempty" xml:"maxRetryTimes,omitempty"`
	RetryPolicy   *string `json:"retryPolicy,omitempty" xml:"retryPolicy,omitempty"`
}

func (s GetConsumerGroupResponseBodyDataConsumeRetryPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupResponseBodyDataConsumeRetryPolicy) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) SetMaxRetryTimes(v int32) *GetConsumerGroupResponseBodyDataConsumeRetryPolicy {
	s.MaxRetryTimes = &v
	return s
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) SetRetryPolicy(v string) *GetConsumerGroupResponseBodyDataConsumeRetryPolicy {
	s.RetryPolicy = &v
	return s
}

type GetConsumerGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponse) SetHeaders(v map[string]*string) *GetConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerGroupResponse) SetStatusCode(v int32) *GetConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerGroupResponse) SetBody(v *GetConsumerGroupResponseBody) *GetConsumerGroupResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	Code           *string                      `json:"code,omitempty" xml:"code,omitempty"`
	Data           *GetInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	DynamicCode    *string                      `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string                      `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32                       `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                      `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceResponseBody) SetDynamicCode(v string) *GetInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetDynamicMessage(v string) *GetInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v int32) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

type GetInstanceResponseBodyData struct {
	AccountInfo    *GetInstanceResponseBodyDataAccountInfo      `json:"accountInfo,omitempty" xml:"accountInfo,omitempty" type:"Struct"`
	Bid            *string                                      `json:"bid,omitempty" xml:"bid,omitempty"`
	CommodityCode  *string                                      `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CreateTime     *string                                      `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ExpireTime     *string                                      `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	ExtConfig      *GetInstanceResponseBodyDataExtConfig        `json:"extConfig,omitempty" xml:"extConfig,omitempty" type:"Struct"`
	GroupCount     *int64                                       `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	InstanceId     *string                                      `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceName   *string                                      `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	InstanceQuotas []*GetInstanceResponseBodyDataInstanceQuotas `json:"instanceQuotas,omitempty" xml:"instanceQuotas,omitempty" type:"Repeated"`
	NetworkInfo    *GetInstanceResponseBodyDataNetworkInfo      `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	PaymentType    *string                                      `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	RegionId       *string                                      `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ReleaseTime    *string                                      `json:"releaseTime,omitempty" xml:"releaseTime,omitempty"`
	Remark         *string                                      `json:"remark,omitempty" xml:"remark,omitempty"`
	SeriesCode     *string                                      `json:"seriesCode,omitempty" xml:"seriesCode,omitempty"`
	ServiceCode    *string                                      `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	StartTime      *string                                      `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status         *string                                      `json:"status,omitempty" xml:"status,omitempty"`
	SubSeriesCode  *string                                      `json:"subSeriesCode,omitempty" xml:"subSeriesCode,omitempty"`
	TopicCount     *int64                                       `json:"topicCount,omitempty" xml:"topicCount,omitempty"`
	UpdateTime     *string                                      `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId         *string                                      `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyData) SetAccountInfo(v *GetInstanceResponseBodyDataAccountInfo) *GetInstanceResponseBodyData {
	s.AccountInfo = v
	return s
}

func (s *GetInstanceResponseBodyData) SetBid(v string) *GetInstanceResponseBodyData {
	s.Bid = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetCommodityCode(v string) *GetInstanceResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetCreateTime(v string) *GetInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetExpireTime(v string) *GetInstanceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetExtConfig(v *GetInstanceResponseBodyDataExtConfig) *GetInstanceResponseBodyData {
	s.ExtConfig = v
	return s
}

func (s *GetInstanceResponseBodyData) SetGroupCount(v int64) *GetInstanceResponseBodyData {
	s.GroupCount = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceId(v string) *GetInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceName(v string) *GetInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceQuotas(v []*GetInstanceResponseBodyDataInstanceQuotas) *GetInstanceResponseBodyData {
	s.InstanceQuotas = v
	return s
}

func (s *GetInstanceResponseBodyData) SetNetworkInfo(v *GetInstanceResponseBodyDataNetworkInfo) *GetInstanceResponseBodyData {
	s.NetworkInfo = v
	return s
}

func (s *GetInstanceResponseBodyData) SetPaymentType(v string) *GetInstanceResponseBodyData {
	s.PaymentType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetRegionId(v string) *GetInstanceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetReleaseTime(v string) *GetInstanceResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetRemark(v string) *GetInstanceResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSeriesCode(v string) *GetInstanceResponseBodyData {
	s.SeriesCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetServiceCode(v string) *GetInstanceResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetStartTime(v string) *GetInstanceResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetStatus(v string) *GetInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSubSeriesCode(v string) *GetInstanceResponseBodyData {
	s.SubSeriesCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetTopicCount(v int64) *GetInstanceResponseBodyData {
	s.TopicCount = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetUpdateTime(v string) *GetInstanceResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetUserId(v string) *GetInstanceResponseBodyData {
	s.UserId = &v
	return s
}

type GetInstanceResponseBodyDataAccountInfo struct {
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetInstanceResponseBodyDataAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataAccountInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataAccountInfo) SetUsername(v string) *GetInstanceResponseBodyDataAccountInfo {
	s.Username = &v
	return s
}

type GetInstanceResponseBodyDataExtConfig struct {
	AclType              *string  `json:"aclType,omitempty" xml:"aclType,omitempty"`
	AutoScaling          *bool    `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	FlowOutBandwidth     *int32   `json:"flowOutBandwidth,omitempty" xml:"flowOutBandwidth,omitempty"`
	FlowOutType          *string  `json:"flowOutType,omitempty" xml:"flowOutType,omitempty"`
	InternetSpec         *string  `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	MessageRetentionTime *int32   `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	MsgProcessSpec       *string  `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	SendReceiveRatio     *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	SupportAutoScaling   *bool    `json:"supportAutoScaling,omitempty" xml:"supportAutoScaling,omitempty"`
}

func (s GetInstanceResponseBodyDataExtConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataExtConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataExtConfig) SetAclType(v string) *GetInstanceResponseBodyDataExtConfig {
	s.AclType = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetAutoScaling(v bool) *GetInstanceResponseBodyDataExtConfig {
	s.AutoScaling = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetFlowOutBandwidth(v int32) *GetInstanceResponseBodyDataExtConfig {
	s.FlowOutBandwidth = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetFlowOutType(v string) *GetInstanceResponseBodyDataExtConfig {
	s.FlowOutType = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetInternetSpec(v string) *GetInstanceResponseBodyDataExtConfig {
	s.InternetSpec = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetMessageRetentionTime(v int32) *GetInstanceResponseBodyDataExtConfig {
	s.MessageRetentionTime = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetMsgProcessSpec(v string) *GetInstanceResponseBodyDataExtConfig {
	s.MsgProcessSpec = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetSendReceiveRatio(v float32) *GetInstanceResponseBodyDataExtConfig {
	s.SendReceiveRatio = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetSupportAutoScaling(v bool) *GetInstanceResponseBodyDataExtConfig {
	s.SupportAutoScaling = &v
	return s
}

type GetInstanceResponseBodyDataInstanceQuotas struct {
	FreeCount  *float64 `json:"freeCount,omitempty" xml:"freeCount,omitempty"`
	QuotaName  *string  `json:"quotaName,omitempty" xml:"quotaName,omitempty"`
	TotalCount *float64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	UsedCount  *float64 `json:"usedCount,omitempty" xml:"usedCount,omitempty"`
}

func (s GetInstanceResponseBodyDataInstanceQuotas) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataInstanceQuotas) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetFreeCount(v float64) *GetInstanceResponseBodyDataInstanceQuotas {
	s.FreeCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetQuotaName(v string) *GetInstanceResponseBodyDataInstanceQuotas {
	s.QuotaName = &v
	return s
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetTotalCount(v float64) *GetInstanceResponseBodyDataInstanceQuotas {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetUsedCount(v float64) *GetInstanceResponseBodyDataInstanceQuotas {
	s.UsedCount = &v
	return s
}

type GetInstanceResponseBodyDataNetworkInfo struct {
	Endpoints []*GetInstanceResponseBodyDataNetworkInfoEndpoints `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	VpcInfo   *GetInstanceResponseBodyDataNetworkInfoVpcInfo     `json:"vpcInfo,omitempty" xml:"vpcInfo,omitempty" type:"Struct"`
}

func (s GetInstanceResponseBodyDataNetworkInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfo) SetEndpoints(v []*GetInstanceResponseBodyDataNetworkInfoEndpoints) *GetInstanceResponseBodyDataNetworkInfo {
	s.Endpoints = v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfo) SetVpcInfo(v *GetInstanceResponseBodyDataNetworkInfoVpcInfo) *GetInstanceResponseBodyDataNetworkInfo {
	s.VpcInfo = v
	return s
}

type GetInstanceResponseBodyDataNetworkInfoEndpoints struct {
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	EndpointUrl  *string `json:"endpointUrl,omitempty" xml:"endpointUrl,omitempty"`
	IpWhitelist  *string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty"`
}

func (s GetInstanceResponseBodyDataNetworkInfoEndpoints) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfoEndpoints) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) SetEndpointType(v string) *GetInstanceResponseBodyDataNetworkInfoEndpoints {
	s.EndpointType = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) SetEndpointUrl(v string) *GetInstanceResponseBodyDataNetworkInfoEndpoints {
	s.EndpointUrl = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) SetIpWhitelist(v string) *GetInstanceResponseBodyDataNetworkInfoEndpoints {
	s.IpWhitelist = &v
	return s
}

type GetInstanceResponseBodyDataNetworkInfoVpcInfo struct {
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	VpcId     *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetInstanceResponseBodyDataNetworkInfoVpcInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfoVpcInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) SetVSwitchId(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) SetVpcId(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	s.VpcId = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetTopicResponseBody struct {
	Code           *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data           *GetTopicResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	DynamicCode    *string                   `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string                   `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32                    `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                   `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicResponseBody) GoString() string {
	return s.String()
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

type GetTopicResponseBodyData struct {
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	RegionId    *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Remark      *string `json:"remark,omitempty" xml:"remark,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	TopicName   *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	UpdateTime  *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetTopicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBodyData) SetCreateTime(v string) *GetTopicResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTopicResponseBodyData) SetInstanceId(v string) *GetTopicResponseBodyData {
	s.InstanceId = &v
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

type GetTopicResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicResponse) GoString() string {
	return s.String()
}

func (s *GetTopicResponse) SetHeaders(v map[string]*string) *GetTopicResponse {
	s.Headers = v
	return s
}

func (s *GetTopicResponse) SetStatusCode(v int32) *GetTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicResponse) SetBody(v *GetTopicResponseBody) *GetTopicResponse {
	s.Body = v
	return s
}

type ListConsumerGroupsRequest struct {
	Filter     *string `json:"filter,omitempty" xml:"filter,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListConsumerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsRequest) SetFilter(v string) *ListConsumerGroupsRequest {
	s.Filter = &v
	return s
}

func (s *ListConsumerGroupsRequest) SetPageNumber(v int32) *ListConsumerGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupsRequest) SetPageSize(v int32) *ListConsumerGroupsRequest {
	s.PageSize = &v
	return s
}

type ListConsumerGroupsResponseBody struct {
	Code           *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data           *ListConsumerGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	DynamicCode    *string                             `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string                             `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32                              `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                             `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListConsumerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBody) SetCode(v string) *ListConsumerGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetData(v *ListConsumerGroupsResponseBodyData) *ListConsumerGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetDynamicCode(v string) *ListConsumerGroupsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetDynamicMessage(v string) *ListConsumerGroupsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetHttpStatusCode(v int32) *ListConsumerGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetMessage(v string) *ListConsumerGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetRequestId(v string) *ListConsumerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetSuccess(v bool) *ListConsumerGroupsResponseBody {
	s.Success = &v
	return s
}

type ListConsumerGroupsResponseBodyData struct {
	List       []*ListConsumerGroupsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNumber *int64                                    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int64                                    `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListConsumerGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBodyData) SetList(v []*ListConsumerGroupsResponseBodyDataList) *ListConsumerGroupsResponseBodyData {
	s.List = v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetPageNumber(v int64) *ListConsumerGroupsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetPageSize(v int64) *ListConsumerGroupsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetTotalCount(v int64) *ListConsumerGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListConsumerGroupsResponseBodyDataList struct {
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	CreateTime      *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	RegionId        *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Remark          *string `json:"remark,omitempty" xml:"remark,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTime      *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListConsumerGroupsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBodyDataList) SetConsumerGroupId(v string) *ListConsumerGroupsResponseBodyDataList {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetCreateTime(v string) *ListConsumerGroupsResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetInstanceId(v string) *ListConsumerGroupsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetRegionId(v string) *ListConsumerGroupsResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetRemark(v string) *ListConsumerGroupsResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetStatus(v string) *ListConsumerGroupsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetUpdateTime(v string) *ListConsumerGroupsResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

type ListConsumerGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConsumerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConsumerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponse) SetHeaders(v map[string]*string) *ListConsumerGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerGroupsResponse) SetStatusCode(v int32) *ListConsumerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerGroupsResponse) SetBody(v *ListConsumerGroupsResponseBody) *ListConsumerGroupsResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	Filter     *string `json:"filter,omitempty" xml:"filter,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetFilter(v string) *ListInstancesRequest {
	s.Filter = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

type ListInstancesResponseBody struct {
	Code           *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data           *ListInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	DynamicCode    *string                        `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string                        `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32                         `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                        `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetDynamicCode(v string) *ListInstancesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetDynamicMessage(v string) *ListInstancesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

type ListInstancesResponseBodyData struct {
	List       []*ListInstancesResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNumber *int64                               `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                               `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int64                               `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) SetList(v []*ListInstancesResponseBodyDataList) *ListInstancesResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageNumber(v int64) *ListInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageSize(v int64) *ListInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetTotalCount(v int64) *ListInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyDataList struct {
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CreateTime    *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ExpireTime    *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	GroupCount    *int64  `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	InstanceId    *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceName  *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	PaymentType   *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	RegionId      *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ReleaseTime   *string `json:"releaseTime,omitempty" xml:"releaseTime,omitempty"`
	Remark        *string `json:"remark,omitempty" xml:"remark,omitempty"`
	SeriesCode    *string `json:"seriesCode,omitempty" xml:"seriesCode,omitempty"`
	ServiceCode   *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	StartTime     *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
	SubSeriesCode *string `json:"subSeriesCode,omitempty" xml:"subSeriesCode,omitempty"`
	TopicCount    *int64  `json:"topicCount,omitempty" xml:"topicCount,omitempty"`
	UpdateTime    *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListInstancesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataList) SetCommodityCode(v string) *ListInstancesResponseBodyDataList {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetCreateTime(v string) *ListInstancesResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetExpireTime(v string) *ListInstancesResponseBodyDataList {
	s.ExpireTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetGroupCount(v int64) *ListInstancesResponseBodyDataList {
	s.GroupCount = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetInstanceId(v string) *ListInstancesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetInstanceName(v string) *ListInstancesResponseBodyDataList {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetPaymentType(v string) *ListInstancesResponseBodyDataList {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetRegionId(v string) *ListInstancesResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetReleaseTime(v string) *ListInstancesResponseBodyDataList {
	s.ReleaseTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetRemark(v string) *ListInstancesResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetSeriesCode(v string) *ListInstancesResponseBodyDataList {
	s.SeriesCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetServiceCode(v string) *ListInstancesResponseBodyDataList {
	s.ServiceCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetStartTime(v string) *ListInstancesResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetStatus(v string) *ListInstancesResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetSubSeriesCode(v string) *ListInstancesResponseBodyDataList {
	s.SubSeriesCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetTopicCount(v int64) *ListInstancesResponseBodyDataList {
	s.TopicCount = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetUpdateTime(v string) *ListInstancesResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetUserId(v string) *ListInstancesResponseBodyDataList {
	s.UserId = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListTopicsRequest struct {
	Filter     *string `json:"filter,omitempty" xml:"filter,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListTopicsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsRequest) GoString() string {
	return s.String()
}

func (s *ListTopicsRequest) SetFilter(v string) *ListTopicsRequest {
	s.Filter = &v
	return s
}

func (s *ListTopicsRequest) SetPageNumber(v int32) *ListTopicsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTopicsRequest) SetPageSize(v int32) *ListTopicsRequest {
	s.PageSize = &v
	return s
}

type ListTopicsResponseBody struct {
	Code           *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data           *ListTopicsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	DynamicCode    *string                     `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string                     `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32                      `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                     `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTopicsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBody) SetCode(v string) *ListTopicsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicsResponseBody) SetData(v *ListTopicsResponseBodyData) *ListTopicsResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicsResponseBody) SetDynamicCode(v string) *ListTopicsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListTopicsResponseBody) SetDynamicMessage(v string) *ListTopicsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListTopicsResponseBody) SetHttpStatusCode(v int32) *ListTopicsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicsResponseBody) SetMessage(v string) *ListTopicsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicsResponseBody) SetRequestId(v string) *ListTopicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicsResponseBody) SetSuccess(v bool) *ListTopicsResponseBody {
	s.Success = &v
	return s
}

type ListTopicsResponseBodyData struct {
	List       []*ListTopicsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNumber *int64                            `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                            `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int64                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTopicsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBodyData) SetList(v []*ListTopicsResponseBodyDataList) *ListTopicsResponseBodyData {
	s.List = v
	return s
}

func (s *ListTopicsResponseBodyData) SetPageNumber(v int64) *ListTopicsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTopicsResponseBodyData) SetPageSize(v int64) *ListTopicsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTopicsResponseBodyData) SetTotalCount(v int64) *ListTopicsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListTopicsResponseBodyDataList struct {
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	RegionId    *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Remark      *string `json:"remark,omitempty" xml:"remark,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	TopicName   *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	UpdateTime  *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListTopicsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBodyDataList) SetCreateTime(v string) *ListTopicsResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetInstanceId(v string) *ListTopicsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetMessageType(v string) *ListTopicsResponseBodyDataList {
	s.MessageType = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetRegionId(v string) *ListTopicsResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetRemark(v string) *ListTopicsResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetStatus(v string) *ListTopicsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetTopicName(v string) *ListTopicsResponseBodyDataList {
	s.TopicName = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetUpdateTime(v string) *ListTopicsResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

type ListTopicsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTopicsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTopicsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsResponse) GoString() string {
	return s.String()
}

func (s *ListTopicsResponse) SetHeaders(v map[string]*string) *ListTopicsResponse {
	s.Headers = v
	return s
}

func (s *ListTopicsResponse) SetStatusCode(v int32) *ListTopicsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicsResponse) SetBody(v *ListTopicsResponseBody) *ListTopicsResponse {
	s.Body = v
	return s
}

type UpdateConsumerGroupRequest struct {
	ConsumeRetryPolicy *UpdateConsumerGroupRequestConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	DeliveryOrderType  *string                                       `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	Remark             *string                                       `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupRequest) SetConsumeRetryPolicy(v *UpdateConsumerGroupRequestConsumeRetryPolicy) *UpdateConsumerGroupRequest {
	s.ConsumeRetryPolicy = v
	return s
}

func (s *UpdateConsumerGroupRequest) SetDeliveryOrderType(v string) *UpdateConsumerGroupRequest {
	s.DeliveryOrderType = &v
	return s
}

func (s *UpdateConsumerGroupRequest) SetRemark(v string) *UpdateConsumerGroupRequest {
	s.Remark = &v
	return s
}

type UpdateConsumerGroupRequestConsumeRetryPolicy struct {
	MaxRetryTimes *int32  `json:"maxRetryTimes,omitempty" xml:"maxRetryTimes,omitempty"`
	RetryPolicy   *string `json:"retryPolicy,omitempty" xml:"retryPolicy,omitempty"`
}

func (s UpdateConsumerGroupRequestConsumeRetryPolicy) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupRequestConsumeRetryPolicy) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) SetMaxRetryTimes(v int32) *UpdateConsumerGroupRequestConsumeRetryPolicy {
	s.MaxRetryTimes = &v
	return s
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) SetRetryPolicy(v string) *UpdateConsumerGroupRequestConsumeRetryPolicy {
	s.RetryPolicy = &v
	return s
}

type UpdateConsumerGroupResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *bool   `json:"data,omitempty" xml:"data,omitempty"`
	DynamicCode    *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupResponseBody) SetCode(v string) *UpdateConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetData(v bool) *UpdateConsumerGroupResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetDynamicCode(v string) *UpdateConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetDynamicMessage(v string) *UpdateConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetHttpStatusCode(v int32) *UpdateConsumerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetMessage(v string) *UpdateConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetRequestId(v string) *UpdateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetSuccess(v bool) *UpdateConsumerGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateConsumerGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupResponse) SetHeaders(v map[string]*string) *UpdateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateConsumerGroupResponse) SetStatusCode(v int32) *UpdateConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConsumerGroupResponse) SetBody(v *UpdateConsumerGroupResponseBody) *UpdateConsumerGroupResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	ExtConfig    *UpdateInstanceRequestExtConfig   `json:"extConfig,omitempty" xml:"extConfig,omitempty" type:"Struct"`
	InstanceName *string                           `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	NetworkInfo  *UpdateInstanceRequestNetworkInfo `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	Remark       *string                           `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetExtConfig(v *UpdateInstanceRequestExtConfig) *UpdateInstanceRequest {
	s.ExtConfig = v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetNetworkInfo(v *UpdateInstanceRequestNetworkInfo) *UpdateInstanceRequest {
	s.NetworkInfo = v
	return s
}

func (s *UpdateInstanceRequest) SetRemark(v string) *UpdateInstanceRequest {
	s.Remark = &v
	return s
}

type UpdateInstanceRequestExtConfig struct {
	AutoScaling          *bool    `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	MessageRetentionTime *int32   `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	SendReceiveRatio     *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
}

func (s UpdateInstanceRequestExtConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestExtConfig) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestExtConfig) SetAutoScaling(v bool) *UpdateInstanceRequestExtConfig {
	s.AutoScaling = &v
	return s
}

func (s *UpdateInstanceRequestExtConfig) SetMessageRetentionTime(v int32) *UpdateInstanceRequestExtConfig {
	s.MessageRetentionTime = &v
	return s
}

func (s *UpdateInstanceRequestExtConfig) SetSendReceiveRatio(v float32) *UpdateInstanceRequestExtConfig {
	s.SendReceiveRatio = &v
	return s
}

type UpdateInstanceRequestNetworkInfo struct {
	Endpoints []*UpdateInstanceRequestNetworkInfoEndpoints `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
}

func (s UpdateInstanceRequestNetworkInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestNetworkInfo) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestNetworkInfo) SetEndpoints(v []*UpdateInstanceRequestNetworkInfoEndpoints) *UpdateInstanceRequestNetworkInfo {
	s.Endpoints = v
	return s
}

type UpdateInstanceRequestNetworkInfoEndpoints struct {
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	IpWhitelist  *string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty"`
}

func (s UpdateInstanceRequestNetworkInfoEndpoints) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestNetworkInfoEndpoints) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestNetworkInfoEndpoints) SetEndpointType(v string) *UpdateInstanceRequestNetworkInfoEndpoints {
	s.EndpointType = &v
	return s
}

func (s *UpdateInstanceRequestNetworkInfoEndpoints) SetIpWhitelist(v string) *UpdateInstanceRequestNetworkInfoEndpoints {
	s.IpWhitelist = &v
	return s
}

type UpdateInstanceResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *bool   `json:"data,omitempty" xml:"data,omitempty"`
	DynamicCode    *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetCode(v string) *UpdateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetData(v bool) *UpdateInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetDynamicCode(v string) *UpdateInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetDynamicMessage(v string) *UpdateInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetMessage(v string) *UpdateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetSuccess(v bool) *UpdateInstanceResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
	s.Body = v
	return s
}

type UpdateTopicRequest struct {
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTopicRequest) GoString() string {
	return s.String()
}

func (s *UpdateTopicRequest) SetRemark(v string) *UpdateTopicRequest {
	s.Remark = &v
	return s
}

type UpdateTopicResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *bool   `json:"data,omitempty" xml:"data,omitempty"`
	DynamicCode    *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTopicResponseBody) SetCode(v string) *UpdateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTopicResponseBody) SetData(v bool) *UpdateTopicResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateTopicResponseBody) SetDynamicCode(v string) *UpdateTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateTopicResponseBody) SetDynamicMessage(v string) *UpdateTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateTopicResponseBody) SetHttpStatusCode(v int32) *UpdateTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTopicResponseBody) SetMessage(v string) *UpdateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTopicResponseBody) SetRequestId(v string) *UpdateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTopicResponseBody) SetSuccess(v bool) *UpdateTopicResponseBody {
	s.Success = &v
	return s
}

type UpdateTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTopicResponse) GoString() string {
	return s.String()
}

func (s *UpdateTopicResponse) SetHeaders(v map[string]*string) *UpdateTopicResponse {
	s.Headers = v
	return s
}

func (s *UpdateTopicResponse) SetStatusCode(v int32) *UpdateTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTopicResponse) SetBody(v *UpdateTopicResponseBody) *UpdateTopicResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("rocketmq"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateConsumerGroup(instanceId *string, consumerGroupId *string, request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CreateConsumerGroupWithOptions(instanceId, consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, request *CreateConsumerGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ConsumeRetryPolicy))) {
		body["consumeRetryPolicy"] = request.ConsumeRetryPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryOrderType)) {
		body["deliveryOrderType"] = request.DeliveryOrderType
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConsumerGroup"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTopic(instanceId *string, topicName *string, request *CreateTopicRequest) (_result *CreateTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTopicResponse{}
	_body, _err := client.CreateTopicWithOptions(instanceId, topicName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTopicWithOptions(instanceId *string, topicName *string, request *CreateTopicRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTopic"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConsumerGroup(instanceId *string, consumerGroupId *string) (_result *DeleteConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DeleteConsumerGroupWithOptions(instanceId, consumerGroupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConsumerGroup"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(instanceId *string) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTopic(instanceId *string, topicName *string) (_result *DeleteTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTopicResponse{}
	_body, _err := client.DeleteTopicWithOptions(instanceId, topicName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTopicWithOptions(instanceId *string, topicName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTopic"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConsumerGroup(instanceId *string, consumerGroupId *string) (_result *GetConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConsumerGroupResponse{}
	_body, _err := client.GetConsumerGroupWithOptions(instanceId, consumerGroupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetConsumerGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetConsumerGroup"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(instanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTopic(instanceId *string, topicName *string) (_result *GetTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTopicResponse{}
	_body, _err := client.GetTopicWithOptions(instanceId, topicName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTopicWithOptions(instanceId *string, topicName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTopicResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTopic"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConsumerGroups(instanceId *string, request *ListConsumerGroupsRequest) (_result *ListConsumerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerGroupsResponse{}
	_body, _err := client.ListConsumerGroupsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConsumerGroupsWithOptions(instanceId *string, request *ListConsumerGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConsumerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConsumerGroups"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConsumerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTopics(instanceId *string, request *ListTopicsRequest) (_result *ListTopicsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTopicsResponse{}
	_body, _err := client.ListTopicsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTopicsWithOptions(instanceId *string, request *ListTopicsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTopicsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTopics"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTopicsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConsumerGroup(instanceId *string, consumerGroupId *string, request *UpdateConsumerGroupRequest) (_result *UpdateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConsumerGroupResponse{}
	_body, _err := client.UpdateConsumerGroupWithOptions(instanceId, consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, request *UpdateConsumerGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ConsumeRetryPolicy))) {
		body["consumeRetryPolicy"] = request.ConsumeRetryPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryOrderType)) {
		body["deliveryOrderType"] = request.DeliveryOrderType
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConsumerGroup"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstance(instanceId *string, request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceWithOptions(instanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ExtConfig))) {
		body["extConfig"] = request.ExtConfig
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.NetworkInfo))) {
		body["networkInfo"] = request.NetworkInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTopic(instanceId *string, topicName *string, request *UpdateTopicRequest) (_result *UpdateTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTopicResponse{}
	_body, _err := client.UpdateTopicWithOptions(instanceId, topicName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTopicWithOptions(instanceId *string, topicName *string, request *UpdateTopicRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTopic"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

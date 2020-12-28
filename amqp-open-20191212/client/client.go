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

type CreateBindingRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost     *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	Argument        *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
}

func (s CreateBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateBindingRequest) SetInstanceId(v string) *CreateBindingRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBindingRequest) SetVirtualHost(v string) *CreateBindingRequest {
	s.VirtualHost = &v
	return s
}

func (s *CreateBindingRequest) SetSourceExchange(v string) *CreateBindingRequest {
	s.SourceExchange = &v
	return s
}

func (s *CreateBindingRequest) SetDestinationName(v string) *CreateBindingRequest {
	s.DestinationName = &v
	return s
}

func (s *CreateBindingRequest) SetBindingKey(v string) *CreateBindingRequest {
	s.BindingKey = &v
	return s
}

func (s *CreateBindingRequest) SetBindingType(v string) *CreateBindingRequest {
	s.BindingType = &v
	return s
}

func (s *CreateBindingRequest) SetArgument(v string) *CreateBindingRequest {
	s.Argument = &v
	return s
}

type CreateBindingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBindingResponseBody) SetRequestId(v string) *CreateBindingResponseBody {
	s.RequestId = &v
	return s
}

type CreateBindingResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingResponse) GoString() string {
	return s.String()
}

func (s *CreateBindingResponse) SetHeaders(v map[string]*string) *CreateBindingResponse {
	s.Headers = v
	return s
}

func (s *CreateBindingResponse) SetBody(v *CreateBindingResponseBody) *CreateBindingResponse {
	s.Body = v
	return s
}

type CreateExchangeRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost       *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	ExchangeName      *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	ExchangeType      *string `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
	AutoDeleteState   *bool   `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	Internal          *bool   `json:"Internal,omitempty" xml:"Internal,omitempty"`
	AlternateExchange *string `json:"AlternateExchange,omitempty" xml:"AlternateExchange,omitempty"`
}

func (s CreateExchangeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExchangeRequest) GoString() string {
	return s.String()
}

func (s *CreateExchangeRequest) SetInstanceId(v string) *CreateExchangeRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExchangeRequest) SetVirtualHost(v string) *CreateExchangeRequest {
	s.VirtualHost = &v
	return s
}

func (s *CreateExchangeRequest) SetExchangeName(v string) *CreateExchangeRequest {
	s.ExchangeName = &v
	return s
}

func (s *CreateExchangeRequest) SetExchangeType(v string) *CreateExchangeRequest {
	s.ExchangeType = &v
	return s
}

func (s *CreateExchangeRequest) SetAutoDeleteState(v bool) *CreateExchangeRequest {
	s.AutoDeleteState = &v
	return s
}

func (s *CreateExchangeRequest) SetInternal(v bool) *CreateExchangeRequest {
	s.Internal = &v
	return s
}

func (s *CreateExchangeRequest) SetAlternateExchange(v string) *CreateExchangeRequest {
	s.AlternateExchange = &v
	return s
}

type CreateExchangeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExchangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExchangeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExchangeResponseBody) SetRequestId(v string) *CreateExchangeResponseBody {
	s.RequestId = &v
	return s
}

type CreateExchangeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateExchangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateExchangeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExchangeResponse) GoString() string {
	return s.String()
}

func (s *CreateExchangeResponse) SetHeaders(v map[string]*string) *CreateExchangeResponse {
	s.Headers = v
	return s
}

func (s *CreateExchangeResponse) SetBody(v *CreateExchangeResponseBody) *CreateExchangeResponse {
	s.Body = v
	return s
}

type CreateQueueRequest struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost          *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	QueueName            *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	AutoDeleteState      *bool   `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	ExclusiveState       *bool   `json:"ExclusiveState,omitempty" xml:"ExclusiveState,omitempty"`
	MessageTTL           *int64  `json:"MessageTTL,omitempty" xml:"MessageTTL,omitempty"`
	AutoExpireState      *int64  `json:"AutoExpireState,omitempty" xml:"AutoExpireState,omitempty"`
	MaxLength            *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	DeadLetterExchange   *string `json:"DeadLetterExchange,omitempty" xml:"DeadLetterExchange,omitempty"`
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" xml:"DeadLetterRoutingKey,omitempty"`
	MaximumPriority      *int32  `json:"MaximumPriority,omitempty" xml:"MaximumPriority,omitempty"`
}

func (s CreateQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueRequest) SetInstanceId(v string) *CreateQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateQueueRequest) SetVirtualHost(v string) *CreateQueueRequest {
	s.VirtualHost = &v
	return s
}

func (s *CreateQueueRequest) SetQueueName(v string) *CreateQueueRequest {
	s.QueueName = &v
	return s
}

func (s *CreateQueueRequest) SetAutoDeleteState(v bool) *CreateQueueRequest {
	s.AutoDeleteState = &v
	return s
}

func (s *CreateQueueRequest) SetExclusiveState(v bool) *CreateQueueRequest {
	s.ExclusiveState = &v
	return s
}

func (s *CreateQueueRequest) SetMessageTTL(v int64) *CreateQueueRequest {
	s.MessageTTL = &v
	return s
}

func (s *CreateQueueRequest) SetAutoExpireState(v int64) *CreateQueueRequest {
	s.AutoExpireState = &v
	return s
}

func (s *CreateQueueRequest) SetMaxLength(v int64) *CreateQueueRequest {
	s.MaxLength = &v
	return s
}

func (s *CreateQueueRequest) SetDeadLetterExchange(v string) *CreateQueueRequest {
	s.DeadLetterExchange = &v
	return s
}

func (s *CreateQueueRequest) SetDeadLetterRoutingKey(v string) *CreateQueueRequest {
	s.DeadLetterRoutingKey = &v
	return s
}

func (s *CreateQueueRequest) SetMaximumPriority(v int32) *CreateQueueRequest {
	s.MaximumPriority = &v
	return s
}

type CreateQueueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBody) SetRequestId(v string) *CreateQueueResponseBody {
	s.RequestId = &v
	return s
}

type CreateQueueResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateQueueResponse) SetHeaders(v map[string]*string) *CreateQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateQueueResponse) SetBody(v *CreateQueueResponseBody) *CreateQueueResponse {
	s.Body = v
	return s
}

type CreateVirtualHostRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s CreateVirtualHostRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualHostRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualHostRequest) SetInstanceId(v string) *CreateVirtualHostRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVirtualHostRequest) SetVirtualHost(v string) *CreateVirtualHostRequest {
	s.VirtualHost = &v
	return s
}

type CreateVirtualHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVirtualHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualHostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualHostResponseBody) SetRequestId(v string) *CreateVirtualHostResponseBody {
	s.RequestId = &v
	return s
}

type CreateVirtualHostResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVirtualHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVirtualHostResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualHostResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualHostResponse) SetHeaders(v map[string]*string) *CreateVirtualHostResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualHostResponse) SetBody(v *CreateVirtualHostResponseBody) *CreateVirtualHostResponse {
	s.Body = v
	return s
}

type DeleteBindingRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost     *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
}

func (s DeleteBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingRequest) GoString() string {
	return s.String()
}

func (s *DeleteBindingRequest) SetInstanceId(v string) *DeleteBindingRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBindingRequest) SetVirtualHost(v string) *DeleteBindingRequest {
	s.VirtualHost = &v
	return s
}

func (s *DeleteBindingRequest) SetSourceExchange(v string) *DeleteBindingRequest {
	s.SourceExchange = &v
	return s
}

func (s *DeleteBindingRequest) SetDestinationName(v string) *DeleteBindingRequest {
	s.DestinationName = &v
	return s
}

func (s *DeleteBindingRequest) SetBindingType(v string) *DeleteBindingRequest {
	s.BindingType = &v
	return s
}

func (s *DeleteBindingRequest) SetBindingKey(v string) *DeleteBindingRequest {
	s.BindingKey = &v
	return s
}

type DeleteBindingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBindingResponseBody) SetRequestId(v string) *DeleteBindingResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBindingResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingResponse) GoString() string {
	return s.String()
}

func (s *DeleteBindingResponse) SetHeaders(v map[string]*string) *DeleteBindingResponse {
	s.Headers = v
	return s
}

func (s *DeleteBindingResponse) SetBody(v *DeleteBindingResponseBody) *DeleteBindingResponse {
	s.Body = v
	return s
}

type DeleteExchangeRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost  *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
}

func (s DeleteExchangeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExchangeRequest) GoString() string {
	return s.String()
}

func (s *DeleteExchangeRequest) SetInstanceId(v string) *DeleteExchangeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteExchangeRequest) SetVirtualHost(v string) *DeleteExchangeRequest {
	s.VirtualHost = &v
	return s
}

func (s *DeleteExchangeRequest) SetExchangeName(v string) *DeleteExchangeRequest {
	s.ExchangeName = &v
	return s
}

type DeleteExchangeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExchangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExchangeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExchangeResponseBody) SetRequestId(v string) *DeleteExchangeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExchangeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteExchangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteExchangeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExchangeResponse) GoString() string {
	return s.String()
}

func (s *DeleteExchangeResponse) SetHeaders(v map[string]*string) *DeleteExchangeResponse {
	s.Headers = v
	return s
}

func (s *DeleteExchangeResponse) SetBody(v *DeleteExchangeResponseBody) *DeleteExchangeResponse {
	s.Body = v
	return s
}

type DeleteQueueRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName   *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s DeleteQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueRequest) SetInstanceId(v string) *DeleteQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteQueueRequest) SetQueueName(v string) *DeleteQueueRequest {
	s.QueueName = &v
	return s
}

func (s *DeleteQueueRequest) SetVirtualHost(v string) *DeleteQueueRequest {
	s.VirtualHost = &v
	return s
}

type DeleteQueueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBody) SetRequestId(v string) *DeleteQueueResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQueueResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponse) SetHeaders(v map[string]*string) *DeleteQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueueResponse) SetBody(v *DeleteQueueResponseBody) *DeleteQueueResponse {
	s.Body = v
	return s
}

type DeleteVirtualHostRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s DeleteVirtualHostRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualHostRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualHostRequest) SetInstanceId(v string) *DeleteVirtualHostRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVirtualHostRequest) SetVirtualHost(v string) *DeleteVirtualHostRequest {
	s.VirtualHost = &v
	return s
}

type DeleteVirtualHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualHostResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualHostResponseBody) SetRequestId(v string) *DeleteVirtualHostResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVirtualHostResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVirtualHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVirtualHostResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualHostResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualHostResponse) SetHeaders(v map[string]*string) *DeleteVirtualHostResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualHostResponse) SetBody(v *DeleteVirtualHostResponseBody) *DeleteVirtualHostResponse {
	s.Body = v
	return s
}

type GetMetadataAmountRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMetadataAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMetadataAmountRequest) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountRequest) SetInstanceId(v string) *GetMetadataAmountRequest {
	s.InstanceId = &v
	return s
}

type GetMetadataAmountResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetMetadataAmountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetMetadataAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMetadataAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponseBody) SetRequestId(v string) *GetMetadataAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetadataAmountResponseBody) SetData(v *GetMetadataAmountResponseBodyData) *GetMetadataAmountResponseBody {
	s.Data = v
	return s
}

type GetMetadataAmountResponseBodyData struct {
	MaxVirtualHosts     *int32 `json:"MaxVirtualHosts,omitempty" xml:"MaxVirtualHosts,omitempty"`
	CurrentVirtualHosts *int32 `json:"CurrentVirtualHosts,omitempty" xml:"CurrentVirtualHosts,omitempty"`
	MaxQueues           *int32 `json:"MaxQueues,omitempty" xml:"MaxQueues,omitempty"`
	CurrentExchanges    *int32 `json:"CurrentExchanges,omitempty" xml:"CurrentExchanges,omitempty"`
	MaxExchanges        *int32 `json:"MaxExchanges,omitempty" xml:"MaxExchanges,omitempty"`
	CurrentQueues       *int32 `json:"CurrentQueues,omitempty" xml:"CurrentQueues,omitempty"`
}

func (s GetMetadataAmountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMetadataAmountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponseBodyData) SetMaxVirtualHosts(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxVirtualHosts = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentVirtualHosts(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentVirtualHosts = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxQueues(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxQueues = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentExchanges(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentExchanges = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxExchanges(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxExchanges = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentQueues(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentQueues = &v
	return s
}

type GetMetadataAmountResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMetadataAmountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMetadataAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMetadataAmountResponse) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponse) SetHeaders(v map[string]*string) *GetMetadataAmountResponse {
	s.Headers = v
	return s
}

func (s *GetMetadataAmountResponse) SetBody(v *GetMetadataAmountResponseBody) *GetMetadataAmountResponse {
	s.Body = v
	return s
}

type ListBindingsRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListBindingsRequest) SetInstanceId(v string) *ListBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBindingsRequest) SetVirtualHost(v string) *ListBindingsRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListBindingsRequest) SetNextToken(v string) *ListBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBindingsRequest) SetMaxResults(v int32) *ListBindingsRequest {
	s.MaxResults = &v
	return s
}

type ListBindingsResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBody) SetRequestId(v string) *ListBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBindingsResponseBody) SetData(v *ListBindingsResponseBodyData) *ListBindingsResponseBody {
	s.Data = v
	return s
}

type ListBindingsResponseBodyData struct {
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Bindings   []*ListBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
}

func (s ListBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBodyData) SetNextToken(v string) *ListBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListBindingsResponseBodyData) SetMaxResults(v int32) *ListBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListBindingsResponseBodyData) SetBindings(v []*ListBindingsResponseBodyDataBindings) *ListBindingsResponseBodyData {
	s.Bindings = v
	return s
}

type ListBindingsResponseBodyDataBindings struct {
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	Argument        *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
}

func (s ListBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetBindingType(v string) *ListBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetArgument(v string) *ListBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

type ListBindingsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListBindingsResponse) SetHeaders(v map[string]*string) *ListBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListBindingsResponse) SetBody(v *ListBindingsResponseBody) *ListBindingsResponse {
	s.Body = v
	return s
}

type ListDownStreamBindingsRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost  *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListDownStreamBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsRequest) SetInstanceId(v string) *ListDownStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetVirtualHost(v string) *ListDownStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetExchangeName(v string) *ListDownStreamBindingsRequest {
	s.ExchangeName = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetNextToken(v string) *ListDownStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetMaxResults(v int32) *ListDownStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

type ListDownStreamBindingsResponseBody struct {
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListDownStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDownStreamBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBody) SetMessage(v string) *ListDownStreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetRequestId(v string) *ListDownStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetData(v *ListDownStreamBindingsResponseBodyData) *ListDownStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetCode(v int32) *ListDownStreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetSuccess(v bool) *ListDownStreamBindingsResponseBody {
	s.Success = &v
	return s
}

type ListDownStreamBindingsResponseBodyData struct {
	NextToken  *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Bindings   []*ListDownStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
}

func (s ListDownStreamBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBodyData) SetNextToken(v string) *ListDownStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyData) SetMaxResults(v int32) *ListDownStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyData) SetBindings(v []*ListDownStreamBindingsResponseBodyDataBindings) *ListDownStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

type ListDownStreamBindingsResponseBodyDataBindings struct {
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	Argument        *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
}

func (s ListDownStreamBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetBindingType(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

type ListDownStreamBindingsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDownStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDownStreamBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponse) SetHeaders(v map[string]*string) *ListDownStreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListDownStreamBindingsResponse) SetBody(v *ListDownStreamBindingsResponseBody) *ListDownStreamBindingsResponse {
	s.Body = v
	return s
}

type ListExchangesRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListExchangesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesRequest) GoString() string {
	return s.String()
}

func (s *ListExchangesRequest) SetInstanceId(v string) *ListExchangesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExchangesRequest) SetVirtualHost(v string) *ListExchangesRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListExchangesRequest) SetNextToken(v string) *ListExchangesRequest {
	s.NextToken = &v
	return s
}

func (s *ListExchangesRequest) SetMaxResults(v int32) *ListExchangesRequest {
	s.MaxResults = &v
	return s
}

type ListExchangesResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListExchangesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListExchangesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBody) SetRequestId(v string) *ListExchangesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExchangesResponseBody) SetData(v *ListExchangesResponseBodyData) *ListExchangesResponseBody {
	s.Data = v
	return s
}

type ListExchangesResponseBodyData struct {
	NextToken  *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Exchanges  []*ListExchangesResponseBodyDataExchanges `json:"Exchanges,omitempty" xml:"Exchanges,omitempty" type:"Repeated"`
}

func (s ListExchangesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBodyData) SetNextToken(v string) *ListExchangesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListExchangesResponseBodyData) SetMaxResults(v int32) *ListExchangesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListExchangesResponseBodyData) SetExchanges(v []*ListExchangesResponseBodyDataExchanges) *ListExchangesResponseBodyData {
	s.Exchanges = v
	return s
}

type ListExchangesResponseBodyDataExchanges struct {
	AutoDeleteState *bool                  `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	CreateTime      *int64                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Attributes      map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	VHostName       *string                `json:"VHostName,omitempty" xml:"VHostName,omitempty"`
	Name            *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	ExchangeType    *string                `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
}

func (s ListExchangesResponseBodyDataExchanges) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponseBodyDataExchanges) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBodyDataExchanges) SetAutoDeleteState(v bool) *ListExchangesResponseBodyDataExchanges {
	s.AutoDeleteState = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetCreateTime(v int64) *ListExchangesResponseBodyDataExchanges {
	s.CreateTime = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetAttributes(v map[string]interface{}) *ListExchangesResponseBodyDataExchanges {
	s.Attributes = v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetVHostName(v string) *ListExchangesResponseBodyDataExchanges {
	s.VHostName = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetName(v string) *ListExchangesResponseBodyDataExchanges {
	s.Name = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetExchangeType(v string) *ListExchangesResponseBodyDataExchanges {
	s.ExchangeType = &v
	return s
}

type ListExchangesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExchangesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExchangesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponse) GoString() string {
	return s.String()
}

func (s *ListExchangesResponse) SetHeaders(v map[string]*string) *ListExchangesResponse {
	s.Headers = v
	return s
}

func (s *ListExchangesResponse) SetBody(v *ListExchangesResponseBody) *ListExchangesResponse {
	s.Body = v
	return s
}

type ListExchangeUpStreamBindingsRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost  *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListExchangeUpStreamBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsRequest) SetInstanceId(v string) *ListExchangeUpStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetVirtualHost(v string) *ListExchangeUpStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetExchangeName(v string) *ListExchangeUpStreamBindingsRequest {
	s.ExchangeName = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetNextToken(v string) *ListExchangeUpStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetMaxResults(v int32) *ListExchangeUpStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

type ListExchangeUpStreamBindingsResponseBody struct {
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListExchangeUpStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetMessage(v string) *ListExchangeUpStreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetRequestId(v string) *ListExchangeUpStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetData(v *ListExchangeUpStreamBindingsResponseBodyData) *ListExchangeUpStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetCode(v int32) *ListExchangeUpStreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetSuccess(v bool) *ListExchangeUpStreamBindingsResponseBody {
	s.Success = &v
	return s
}

type ListExchangeUpStreamBindingsResponseBodyData struct {
	NextToken  *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Bindings   []*ListExchangeUpStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
}

func (s ListExchangeUpStreamBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetNextToken(v string) *ListExchangeUpStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetMaxResults(v int32) *ListExchangeUpStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetBindings(v []*ListExchangeUpStreamBindingsResponseBodyDataBindings) *ListExchangeUpStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

type ListExchangeUpStreamBindingsResponseBodyDataBindings struct {
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	Argument        *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetBindingType(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

type ListExchangeUpStreamBindingsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExchangeUpStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExchangeUpStreamBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponse) SetHeaders(v map[string]*string) *ListExchangeUpStreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponse) SetBody(v *ListExchangeUpStreamBindingsResponseBody) *ListExchangeUpStreamBindingsResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetNextToken(v string) *ListInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesRequest) SetMaxResults(v int32) *ListInstancesRequest {
	s.MaxResults = &v
	return s
}

type ListInstancesResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

type ListInstancesResponseBodyData struct {
	NextToken  *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Instances  []*ListInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) SetNextToken(v string) *ListInstancesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetMaxResults(v int32) *ListInstancesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetInstances(v []*ListInstancesResponseBodyDataInstances) *ListInstancesResponseBodyData {
	s.Instances = v
	return s
}

type ListInstancesResponseBodyDataInstances struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportEIP        *bool   `json:"SupportEIP,omitempty" xml:"SupportEIP,omitempty"`
	AutoRenewInstance *bool   `json:"AutoRenewInstance,omitempty" xml:"AutoRenewInstance,omitempty"`
	ExpireTime        *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	OrderCreateTime   *int64  `json:"OrderCreateTime,omitempty" xml:"OrderCreateTime,omitempty"`
	InstanceName      *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PrivateEndpoint   *string `json:"PrivateEndpoint,omitempty" xml:"PrivateEndpoint,omitempty"`
	OrderType         *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType      *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PublicEndpoint    *string `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
}

func (s ListInstancesResponseBodyDataInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataInstances) SetStatus(v string) *ListInstancesResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetSupportEIP(v bool) *ListInstancesResponseBodyDataInstances {
	s.SupportEIP = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetAutoRenewInstance(v bool) *ListInstancesResponseBodyDataInstances {
	s.AutoRenewInstance = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetExpireTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.ExpireTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetOrderCreateTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.OrderCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceName(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetPrivateEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.PrivateEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetOrderType(v string) *ListInstancesResponseBodyDataInstances {
	s.OrderType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceId(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceType(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetPublicEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.PublicEndpoint = &v
	return s
}

type ListInstancesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListQueueConsumersRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	Queue       *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	QueryCount  *int32  `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
}

func (s ListQueueConsumersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersRequest) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersRequest) SetInstanceId(v string) *ListQueueConsumersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueueConsumersRequest) SetVirtualHost(v string) *ListQueueConsumersRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListQueueConsumersRequest) SetQueue(v string) *ListQueueConsumersRequest {
	s.Queue = &v
	return s
}

func (s *ListQueueConsumersRequest) SetNextToken(v string) *ListQueueConsumersRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueueConsumersRequest) SetQueryCount(v int32) *ListQueueConsumersRequest {
	s.QueryCount = &v
	return s
}

type ListQueueConsumersResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListQueueConsumersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListQueueConsumersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBody) SetRequestId(v string) *ListQueueConsumersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueueConsumersResponseBody) SetData(v *ListQueueConsumersResponseBodyData) *ListQueueConsumersResponseBody {
	s.Data = v
	return s
}

type ListQueueConsumersResponseBodyData struct {
	NextToken  *string                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Consumers  []*ListQueueConsumersResponseBodyDataConsumers `json:"Consumers,omitempty" xml:"Consumers,omitempty" type:"Repeated"`
	MaxResults *int32                                         `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListQueueConsumersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBodyData) SetNextToken(v string) *ListQueueConsumersResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQueueConsumersResponseBodyData) SetConsumers(v []*ListQueueConsumersResponseBodyDataConsumers) *ListQueueConsumersResponseBodyData {
	s.Consumers = v
	return s
}

func (s *ListQueueConsumersResponseBodyData) SetMaxResults(v int32) *ListQueueConsumersResponseBodyData {
	s.MaxResults = &v
	return s
}

type ListQueueConsumersResponseBodyDataConsumers struct {
	ConsumerTag *string `json:"ConsumerTag,omitempty" xml:"ConsumerTag,omitempty"`
}

func (s ListQueueConsumersResponseBodyDataConsumers) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersResponseBodyDataConsumers) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBodyDataConsumers) SetConsumerTag(v string) *ListQueueConsumersResponseBodyDataConsumers {
	s.ConsumerTag = &v
	return s
}

type ListQueueConsumersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQueueConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueueConsumersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersResponse) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponse) SetHeaders(v map[string]*string) *ListQueueConsumersResponse {
	s.Headers = v
	return s
}

func (s *ListQueueConsumersResponse) SetBody(v *ListQueueConsumersResponseBody) *ListQueueConsumersResponse {
	s.Body = v
	return s
}

type ListQueuesRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListQueuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListQueuesRequest) SetInstanceId(v string) *ListQueuesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueuesRequest) SetVirtualHost(v string) *ListQueuesRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListQueuesRequest) SetNextToken(v string) *ListQueuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueuesRequest) SetMaxResults(v int32) *ListQueuesRequest {
	s.MaxResults = &v
	return s
}

type ListQueuesResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListQueuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListQueuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBody) SetRequestId(v string) *ListQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueuesResponseBody) SetData(v *ListQueuesResponseBodyData) *ListQueuesResponseBody {
	s.Data = v
	return s
}

type ListQueuesResponseBodyData struct {
	NextToken  *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Queues     []*ListQueuesResponseBodyDataQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	MaxResults *int32                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListQueuesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyData) SetNextToken(v string) *ListQueuesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQueuesResponseBodyData) SetQueues(v []*ListQueuesResponseBodyDataQueues) *ListQueuesResponseBodyData {
	s.Queues = v
	return s
}

func (s *ListQueuesResponseBodyData) SetMaxResults(v int32) *ListQueuesResponseBodyData {
	s.MaxResults = &v
	return s
}

type ListQueuesResponseBodyDataQueues struct {
	ExclusiveState  *bool                  `json:"ExclusiveState,omitempty" xml:"ExclusiveState,omitempty"`
	AutoDeleteState *bool                  `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	CreateTime      *int64                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Attributes      map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	VHostName       *string                `json:"VHostName,omitempty" xml:"VHostName,omitempty"`
	Name            *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	LastConsumeTime *int64                 `json:"LastConsumeTime,omitempty" xml:"LastConsumeTime,omitempty"`
}

func (s ListQueuesResponseBodyDataQueues) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyDataQueues) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyDataQueues) SetExclusiveState(v bool) *ListQueuesResponseBodyDataQueues {
	s.ExclusiveState = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetAutoDeleteState(v bool) *ListQueuesResponseBodyDataQueues {
	s.AutoDeleteState = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetCreateTime(v int64) *ListQueuesResponseBodyDataQueues {
	s.CreateTime = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetAttributes(v map[string]interface{}) *ListQueuesResponseBodyDataQueues {
	s.Attributes = v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetVHostName(v string) *ListQueuesResponseBodyDataQueues {
	s.VHostName = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetName(v string) *ListQueuesResponseBodyDataQueues {
	s.Name = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetOwnerId(v string) *ListQueuesResponseBodyDataQueues {
	s.OwnerId = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetLastConsumeTime(v int64) *ListQueuesResponseBodyDataQueues {
	s.LastConsumeTime = &v
	return s
}

type ListQueuesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponse) GoString() string {
	return s.String()
}

func (s *ListQueuesResponse) SetHeaders(v map[string]*string) *ListQueuesResponse {
	s.Headers = v
	return s
}

func (s *ListQueuesResponse) SetBody(v *ListQueuesResponseBody) *ListQueuesResponse {
	s.Body = v
	return s
}

type ListQueueUpStreamBindingsRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
	QueueName   *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListQueueUpStreamBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsRequest) SetInstanceId(v string) *ListQueueUpStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetVirtualHost(v string) *ListQueueUpStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetQueueName(v string) *ListQueueUpStreamBindingsRequest {
	s.QueueName = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetNextToken(v string) *ListQueueUpStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetMaxResults(v int32) *ListQueueUpStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

type ListQueueUpStreamBindingsResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListQueueUpStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListQueueUpStreamBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBody) SetRequestId(v string) *ListQueueUpStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBody) SetData(v *ListQueueUpStreamBindingsResponseBodyData) *ListQueueUpStreamBindingsResponseBody {
	s.Data = v
	return s
}

type ListQueueUpStreamBindingsResponseBodyData struct {
	NextToken  *string                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *string                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Bindings   []*ListQueueUpStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
}

func (s ListQueueUpStreamBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetNextToken(v string) *ListQueueUpStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetMaxResults(v string) *ListQueueUpStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetBindings(v []*ListQueueUpStreamBindingsResponseBodyDataBindings) *ListQueueUpStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

type ListQueueUpStreamBindingsResponseBodyDataBindings struct {
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	Argument        *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
}

func (s ListQueueUpStreamBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetBindingKey(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.BindingKey = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetBindingType(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.BindingType = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

type ListQueueUpStreamBindingsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQueueUpStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueueUpStreamBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponse) SetHeaders(v map[string]*string) *ListQueueUpStreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListQueueUpStreamBindingsResponse) SetBody(v *ListQueueUpStreamBindingsResponseBody) *ListQueueUpStreamBindingsResponse {
	s.Body = v
	return s
}

type ListVirtualHostsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListVirtualHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsRequest) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsRequest) SetInstanceId(v string) *ListVirtualHostsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVirtualHostsRequest) SetNextToken(v string) *ListVirtualHostsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVirtualHostsRequest) SetMaxResults(v int32) *ListVirtualHostsRequest {
	s.MaxResults = &v
	return s
}

type ListVirtualHostsResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListVirtualHostsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListVirtualHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBody) SetRequestId(v string) *ListVirtualHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirtualHostsResponseBody) SetData(v *ListVirtualHostsResponseBodyData) *ListVirtualHostsResponseBody {
	s.Data = v
	return s
}

type ListVirtualHostsResponseBodyData struct {
	NextToken    *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *int32                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	VirtualHosts []*ListVirtualHostsResponseBodyDataVirtualHosts `json:"VirtualHosts,omitempty" xml:"VirtualHosts,omitempty" type:"Repeated"`
}

func (s ListVirtualHostsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBodyData) SetNextToken(v string) *ListVirtualHostsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListVirtualHostsResponseBodyData) SetMaxResults(v int32) *ListVirtualHostsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualHostsResponseBodyData) SetVirtualHosts(v []*ListVirtualHostsResponseBodyDataVirtualHosts) *ListVirtualHostsResponseBodyData {
	s.VirtualHosts = v
	return s
}

type ListVirtualHostsResponseBodyDataVirtualHosts struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListVirtualHostsResponseBodyDataVirtualHosts) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsResponseBodyDataVirtualHosts) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBodyDataVirtualHosts) SetName(v string) *ListVirtualHostsResponseBodyDataVirtualHosts {
	s.Name = &v
	return s
}

type ListVirtualHostsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVirtualHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVirtualHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponse) SetHeaders(v map[string]*string) *ListVirtualHostsResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualHostsResponse) SetBody(v *ListVirtualHostsResponseBody) *ListVirtualHostsResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("amqp-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateBindingWithOptions(request *CreateBindingRequest, runtime *util.RuntimeOptions) (_result *CreateBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBindingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBinding"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBinding(request *CreateBindingRequest) (_result *CreateBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBindingResponse{}
	_body, _err := client.CreateBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateExchangeWithOptions(request *CreateExchangeRequest, runtime *util.RuntimeOptions) (_result *CreateExchangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateExchangeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateExchange"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateExchange(request *CreateExchangeRequest) (_result *CreateExchangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateExchangeResponse{}
	_body, _err := client.CreateExchangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateQueueWithOptions(request *CreateQueueRequest, runtime *util.RuntimeOptions) (_result *CreateQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateQueue"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateQueue(request *CreateQueueRequest) (_result *CreateQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQueueResponse{}
	_body, _err := client.CreateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVirtualHostWithOptions(request *CreateVirtualHostRequest, runtime *util.RuntimeOptions) (_result *CreateVirtualHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVirtualHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVirtualHost"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVirtualHost(request *CreateVirtualHostRequest) (_result *CreateVirtualHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVirtualHostResponse{}
	_body, _err := client.CreateVirtualHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBindingWithOptions(request *DeleteBindingRequest, runtime *util.RuntimeOptions) (_result *DeleteBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBindingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBinding"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBinding(request *DeleteBindingRequest) (_result *DeleteBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBindingResponse{}
	_body, _err := client.DeleteBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteExchangeWithOptions(request *DeleteExchangeRequest, runtime *util.RuntimeOptions) (_result *DeleteExchangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteExchangeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteExchange"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteExchange(request *DeleteExchangeRequest) (_result *DeleteExchangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExchangeResponse{}
	_body, _err := client.DeleteExchangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteQueueWithOptions(request *DeleteQueueRequest, runtime *util.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteQueue"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteQueue(request *DeleteQueueRequest) (_result *DeleteQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQueueResponse{}
	_body, _err := client.DeleteQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVirtualHostWithOptions(request *DeleteVirtualHostRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVirtualHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVirtualHost"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVirtualHost(request *DeleteVirtualHostRequest) (_result *DeleteVirtualHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualHostResponse{}
	_body, _err := client.DeleteVirtualHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMetadataAmountWithOptions(request *GetMetadataAmountRequest, runtime *util.RuntimeOptions) (_result *GetMetadataAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetMetadataAmountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMetadataAmount"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMetadataAmount(request *GetMetadataAmountRequest) (_result *GetMetadataAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMetadataAmountResponse{}
	_body, _err := client.GetMetadataAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBindingsWithOptions(request *ListBindingsRequest, runtime *util.RuntimeOptions) (_result *ListBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListBindingsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBindings"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBindings(request *ListBindingsRequest) (_result *ListBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBindingsResponse{}
	_body, _err := client.ListBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDownStreamBindingsWithOptions(request *ListDownStreamBindingsRequest, runtime *util.RuntimeOptions) (_result *ListDownStreamBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDownStreamBindingsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDownStreamBindings"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDownStreamBindings(request *ListDownStreamBindingsRequest) (_result *ListDownStreamBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDownStreamBindingsResponse{}
	_body, _err := client.ListDownStreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExchangesWithOptions(request *ListExchangesRequest, runtime *util.RuntimeOptions) (_result *ListExchangesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListExchangesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListExchanges"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExchanges(request *ListExchangesRequest) (_result *ListExchangesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExchangesResponse{}
	_body, _err := client.ListExchangesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExchangeUpStreamBindingsWithOptions(request *ListExchangeUpStreamBindingsRequest, runtime *util.RuntimeOptions) (_result *ListExchangeUpStreamBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListExchangeUpStreamBindingsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListExchangeUpStreamBindings"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExchangeUpStreamBindings(request *ListExchangeUpStreamBindingsRequest) (_result *ListExchangeUpStreamBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExchangeUpStreamBindingsResponse{}
	_body, _err := client.ListExchangeUpStreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstances"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueueConsumersWithOptions(request *ListQueueConsumersRequest, runtime *util.RuntimeOptions) (_result *ListQueueConsumersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListQueueConsumersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListQueueConsumers"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueueConsumers(request *ListQueueConsumersRequest) (_result *ListQueueConsumersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueueConsumersResponse{}
	_body, _err := client.ListQueueConsumersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueuesWithOptions(request *ListQueuesRequest, runtime *util.RuntimeOptions) (_result *ListQueuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListQueuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListQueues"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueues(request *ListQueuesRequest) (_result *ListQueuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueuesResponse{}
	_body, _err := client.ListQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueueUpStreamBindingsWithOptions(request *ListQueueUpStreamBindingsRequest, runtime *util.RuntimeOptions) (_result *ListQueueUpStreamBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListQueueUpStreamBindingsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListQueueUpStreamBindings"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueueUpStreamBindings(request *ListQueueUpStreamBindingsRequest) (_result *ListQueueUpStreamBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueueUpStreamBindingsResponse{}
	_body, _err := client.ListQueueUpStreamBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVirtualHostsWithOptions(request *ListVirtualHostsRequest, runtime *util.RuntimeOptions) (_result *ListVirtualHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListVirtualHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVirtualHosts"), tea.String("2019-12-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVirtualHosts(request *ListVirtualHostsRequest) (_result *ListVirtualHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVirtualHostsResponse{}
	_body, _err := client.ListVirtualHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

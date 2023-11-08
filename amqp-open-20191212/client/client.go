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

type DataValue struct {
	MasterUid       *int64  `json:"masterUid,omitempty" xml:"masterUid,omitempty"`
	CInstanceId     *string `json:"cInstanceId,omitempty" xml:"cInstanceId,omitempty"`
	AccessKey       *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	UserName        *string `json:"userName,omitempty" xml:"userName,omitempty"`
	Password        *string `json:"password,omitempty" xml:"password,omitempty"`
	Deleted         *int64  `json:"deleted,omitempty" xml:"deleted,omitempty"`
	CreateTimestamp *int64  `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
}

func (s DataValue) String() string {
	return tea.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) SetMasterUid(v int64) *DataValue {
	s.MasterUid = &v
	return s
}

func (s *DataValue) SetCInstanceId(v string) *DataValue {
	s.CInstanceId = &v
	return s
}

func (s *DataValue) SetAccessKey(v string) *DataValue {
	s.AccessKey = &v
	return s
}

func (s *DataValue) SetUserName(v string) *DataValue {
	s.UserName = &v
	return s
}

func (s *DataValue) SetPassword(v string) *DataValue {
	s.Password = &v
	return s
}

func (s *DataValue) SetDeleted(v int64) *DataValue {
	s.Deleted = &v
	return s
}

func (s *DataValue) SetCreateTimestamp(v int64) *DataValue {
	s.CreateTimestamp = &v
	return s
}

type CreateAccountRequest struct {
	AccountAccessKey *string `json:"accountAccessKey,omitempty" xml:"accountAccessKey,omitempty"`
	CreateTimestamp  *int64  `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	InstanceId       *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	SecretSign       *string `json:"secretSign,omitempty" xml:"secretSign,omitempty"`
	Signature        *string `json:"signature,omitempty" xml:"signature,omitempty"`
	UserName         *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetAccountAccessKey(v string) *CreateAccountRequest {
	s.AccountAccessKey = &v
	return s
}

func (s *CreateAccountRequest) SetCreateTimestamp(v int64) *CreateAccountRequest {
	s.CreateTimestamp = &v
	return s
}

func (s *CreateAccountRequest) SetInstanceId(v string) *CreateAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetSecretSign(v string) *CreateAccountRequest {
	s.SecretSign = &v
	return s
}

func (s *CreateAccountRequest) SetSignature(v string) *CreateAccountRequest {
	s.Signature = &v
	return s
}

func (s *CreateAccountRequest) SetUserName(v string) *CreateAccountRequest {
	s.UserName = &v
	return s
}

type CreateAccountResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetCode(v int32) *CreateAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAccountResponseBody) SetData(v *CreateAccountResponseBodyData) *CreateAccountResponseBody {
	s.Data = v
	return s
}

func (s *CreateAccountResponseBody) SetMessage(v string) *CreateAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) SetSuccess(v bool) *CreateAccountResponseBody {
	s.Success = &v
	return s
}

type CreateAccountResponseBodyData struct {
	// AccessKey ID。
	AccessKey       *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	CreateTimeStamp *int64  `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MasterUId       *int64  `json:"MasterUId,omitempty" xml:"MasterUId,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBodyData) SetAccessKey(v string) *CreateAccountResponseBodyData {
	s.AccessKey = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetCreateTimeStamp(v int64) *CreateAccountResponseBodyData {
	s.CreateTimeStamp = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetInstanceId(v string) *CreateAccountResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetMasterUId(v int64) *CreateAccountResponseBodyData {
	s.MasterUId = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetPassword(v string) *CreateAccountResponseBodyData {
	s.Password = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetUserName(v string) *CreateAccountResponseBodyData {
	s.UserName = &v
	return s
}

type CreateAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) SetHeaders(v map[string]*string) *CreateAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountResponse) SetStatusCode(v int32) *CreateAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateBindingRequest struct {
	Argument        *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	VirtualHost     *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s CreateBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateBindingRequest) SetArgument(v string) *CreateBindingRequest {
	s.Argument = &v
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

func (s *CreateBindingRequest) SetDestinationName(v string) *CreateBindingRequest {
	s.DestinationName = &v
	return s
}

func (s *CreateBindingRequest) SetInstanceId(v string) *CreateBindingRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBindingRequest) SetSourceExchange(v string) *CreateBindingRequest {
	s.SourceExchange = &v
	return s
}

func (s *CreateBindingRequest) SetVirtualHost(v string) *CreateBindingRequest {
	s.VirtualHost = &v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateBindingResponse) SetStatusCode(v int32) *CreateBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBindingResponse) SetBody(v *CreateBindingResponseBody) *CreateBindingResponse {
	s.Body = v
	return s
}

type CreateExchangeRequest struct {
	AlternateExchange *string `json:"AlternateExchange,omitempty" xml:"AlternateExchange,omitempty"`
	AutoDeleteState   *bool   `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	ExchangeName      *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	ExchangeType      *string `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Internal          *bool   `json:"Internal,omitempty" xml:"Internal,omitempty"`
	VirtualHost       *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s CreateExchangeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExchangeRequest) GoString() string {
	return s.String()
}

func (s *CreateExchangeRequest) SetAlternateExchange(v string) *CreateExchangeRequest {
	s.AlternateExchange = &v
	return s
}

func (s *CreateExchangeRequest) SetAutoDeleteState(v bool) *CreateExchangeRequest {
	s.AutoDeleteState = &v
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

func (s *CreateExchangeRequest) SetInstanceId(v string) *CreateExchangeRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExchangeRequest) SetInternal(v bool) *CreateExchangeRequest {
	s.Internal = &v
	return s
}

func (s *CreateExchangeRequest) SetVirtualHost(v string) *CreateExchangeRequest {
	s.VirtualHost = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateExchangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateExchangeResponse) SetStatusCode(v int32) *CreateExchangeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExchangeResponse) SetBody(v *CreateExchangeResponseBody) *CreateExchangeResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	AutoRenew          *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod    *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaxConnections     *int32  `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MaxEipTps          *int64  `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	MaxPrivateTps      *int64  `json:"MaxPrivateTps,omitempty" xml:"MaxPrivateTps,omitempty"`
	PaymentType        *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodCycle        *string `json:"PeriodCycle,omitempty" xml:"PeriodCycle,omitempty"`
	QueueCapacity      *int32  `json:"QueueCapacity,omitempty" xml:"QueueCapacity,omitempty"`
	StorageSize        *int32  `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	SupportEip         *bool   `json:"SupportEip,omitempty" xml:"SupportEip,omitempty"`
	SupportTracing     *bool   `json:"SupportTracing,omitempty" xml:"SupportTracing,omitempty"`
	TracingStorageTime *int32  `json:"TracingStorageTime,omitempty" xml:"TracingStorageTime,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v int32) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxConnections(v int32) *CreateInstanceRequest {
	s.MaxConnections = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxEipTps(v int64) *CreateInstanceRequest {
	s.MaxEipTps = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxPrivateTps(v int64) *CreateInstanceRequest {
	s.MaxPrivateTps = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodCycle(v string) *CreateInstanceRequest {
	s.PeriodCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetQueueCapacity(v int32) *CreateInstanceRequest {
	s.QueueCapacity = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSize(v int32) *CreateInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateInstanceRequest) SetSupportEip(v bool) *CreateInstanceRequest {
	s.SupportEip = &v
	return s
}

func (s *CreateInstanceRequest) SetSupportTracing(v bool) *CreateInstanceRequest {
	s.SupportTracing = &v
	return s
}

func (s *CreateInstanceRequest) SetTracingStorageTime(v int32) *CreateInstanceRequest {
	s.TracingStorageTime = &v
	return s
}

type CreateInstanceResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetCode(v int32) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetData(v interface{}) *CreateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateQueueRequest struct {
	AutoDeleteState      *bool   `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	AutoExpireState      *int64  `json:"AutoExpireState,omitempty" xml:"AutoExpireState,omitempty"`
	DeadLetterExchange   *string `json:"DeadLetterExchange,omitempty" xml:"DeadLetterExchange,omitempty"`
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" xml:"DeadLetterRoutingKey,omitempty"`
	ExclusiveState       *bool   `json:"ExclusiveState,omitempty" xml:"ExclusiveState,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxLength            *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaximumPriority      *int32  `json:"MaximumPriority,omitempty" xml:"MaximumPriority,omitempty"`
	MessageTTL           *int64  `json:"MessageTTL,omitempty" xml:"MessageTTL,omitempty"`
	QueueName            *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	VirtualHost          *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s CreateQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueRequest) SetAutoDeleteState(v bool) *CreateQueueRequest {
	s.AutoDeleteState = &v
	return s
}

func (s *CreateQueueRequest) SetAutoExpireState(v int64) *CreateQueueRequest {
	s.AutoExpireState = &v
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

func (s *CreateQueueRequest) SetExclusiveState(v bool) *CreateQueueRequest {
	s.ExclusiveState = &v
	return s
}

func (s *CreateQueueRequest) SetInstanceId(v string) *CreateQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateQueueRequest) SetMaxLength(v int64) *CreateQueueRequest {
	s.MaxLength = &v
	return s
}

func (s *CreateQueueRequest) SetMaximumPriority(v int32) *CreateQueueRequest {
	s.MaximumPriority = &v
	return s
}

func (s *CreateQueueRequest) SetMessageTTL(v int64) *CreateQueueRequest {
	s.MessageTTL = &v
	return s
}

func (s *CreateQueueRequest) SetQueueName(v string) *CreateQueueRequest {
	s.QueueName = &v
	return s
}

func (s *CreateQueueRequest) SetVirtualHost(v string) *CreateQueueRequest {
	s.VirtualHost = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateQueueResponse) SetStatusCode(v int32) *CreateQueueResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVirtualHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateVirtualHostResponse) SetStatusCode(v int32) *CreateVirtualHostResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualHostResponse) SetBody(v *CreateVirtualHostResponseBody) *CreateVirtualHostResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetCreateTimestamp(v int64) *DeleteAccountRequest {
	s.CreateTimestamp = &v
	return s
}

func (s *DeleteAccountRequest) SetUserName(v string) *DeleteAccountRequest {
	s.UserName = &v
	return s
}

type DeleteAccountResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetCode(v int32) *DeleteAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAccountResponseBody) SetData(v bool) *DeleteAccountResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAccountResponseBody) SetMessage(v string) *DeleteAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountResponseBody) SetSuccess(v bool) *DeleteAccountResponseBody {
	s.Success = &v
	return s
}

type DeleteAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetHeaders(v map[string]*string) *DeleteAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountResponse) SetStatusCode(v int32) *DeleteAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type DeleteBindingRequest struct {
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
	VirtualHost     *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s DeleteBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBindingRequest) GoString() string {
	return s.String()
}

func (s *DeleteBindingRequest) SetBindingKey(v string) *DeleteBindingRequest {
	s.BindingKey = &v
	return s
}

func (s *DeleteBindingRequest) SetBindingType(v string) *DeleteBindingRequest {
	s.BindingType = &v
	return s
}

func (s *DeleteBindingRequest) SetDestinationName(v string) *DeleteBindingRequest {
	s.DestinationName = &v
	return s
}

func (s *DeleteBindingRequest) SetInstanceId(v string) *DeleteBindingRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBindingRequest) SetSourceExchange(v string) *DeleteBindingRequest {
	s.SourceExchange = &v
	return s
}

func (s *DeleteBindingRequest) SetVirtualHost(v string) *DeleteBindingRequest {
	s.VirtualHost = &v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteBindingResponse) SetStatusCode(v int32) *DeleteBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBindingResponse) SetBody(v *DeleteBindingResponseBody) *DeleteBindingResponse {
	s.Body = v
	return s
}

type DeleteExchangeRequest struct {
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VirtualHost  *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s DeleteExchangeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExchangeRequest) GoString() string {
	return s.String()
}

func (s *DeleteExchangeRequest) SetExchangeName(v string) *DeleteExchangeRequest {
	s.ExchangeName = &v
	return s
}

func (s *DeleteExchangeRequest) SetInstanceId(v string) *DeleteExchangeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteExchangeRequest) SetVirtualHost(v string) *DeleteExchangeRequest {
	s.VirtualHost = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteExchangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteExchangeResponse) SetStatusCode(v int32) *DeleteExchangeResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteQueueResponse) SetStatusCode(v int32) *DeleteQueueResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVirtualHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVirtualHostResponse) SetStatusCode(v int32) *DeleteVirtualHostResponse {
	s.StatusCode = &v
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
	Data      *GetMetadataAmountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMetadataAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMetadataAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponseBody) SetData(v *GetMetadataAmountResponseBodyData) *GetMetadataAmountResponseBody {
	s.Data = v
	return s
}

func (s *GetMetadataAmountResponseBody) SetRequestId(v string) *GetMetadataAmountResponseBody {
	s.RequestId = &v
	return s
}

type GetMetadataAmountResponseBodyData struct {
	CurrentExchanges    *int32 `json:"CurrentExchanges,omitempty" xml:"CurrentExchanges,omitempty"`
	CurrentQueues       *int32 `json:"CurrentQueues,omitempty" xml:"CurrentQueues,omitempty"`
	CurrentVirtualHosts *int32 `json:"CurrentVirtualHosts,omitempty" xml:"CurrentVirtualHosts,omitempty"`
	MaxExchanges        *int32 `json:"MaxExchanges,omitempty" xml:"MaxExchanges,omitempty"`
	MaxQueues           *int32 `json:"MaxQueues,omitempty" xml:"MaxQueues,omitempty"`
	MaxVirtualHosts     *int32 `json:"MaxVirtualHosts,omitempty" xml:"MaxVirtualHosts,omitempty"`
}

func (s GetMetadataAmountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMetadataAmountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentExchanges(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentExchanges = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentQueues(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentQueues = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentVirtualHosts(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentVirtualHosts = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxExchanges(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxExchanges = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxQueues(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxQueues = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxVirtualHosts(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxVirtualHosts = &v
	return s
}

type GetMetadataAmountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMetadataAmountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetMetadataAmountResponse) SetStatusCode(v int32) *GetMetadataAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetadataAmountResponse) SetBody(v *GetMetadataAmountResponseBody) *GetMetadataAmountResponse {
	s.Body = v
	return s
}

type ListAccountsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsRequest) SetInstanceId(v string) *ListAccountsRequest {
	s.InstanceId = &v
	return s
}

type ListAccountsResponseBody struct {
	Code      *int32                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string][]*DataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsResponseBody) SetCode(v int32) *ListAccountsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAccountsResponseBody) SetData(v map[string][]*DataValue) *ListAccountsResponseBody {
	s.Data = v
	return s
}

func (s *ListAccountsResponseBody) SetMessage(v string) *ListAccountsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAccountsResponseBody) SetRequestId(v string) *ListAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsResponseBody) SetSuccess(v bool) *ListAccountsResponseBody {
	s.Success = &v
	return s
}

type ListAccountsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsResponse) SetHeaders(v map[string]*string) *ListAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsResponse) SetStatusCode(v int32) *ListAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsResponse) SetBody(v *ListAccountsResponseBody) *ListAccountsResponse {
	s.Body = v
	return s
}

type ListBindingsRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
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

func (s *ListBindingsRequest) SetMaxResults(v int32) *ListBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBindingsRequest) SetNextToken(v string) *ListBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBindingsRequest) SetVirtualHost(v string) *ListBindingsRequest {
	s.VirtualHost = &v
	return s
}

type ListBindingsResponseBody struct {
	Data      *ListBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBody) SetData(v *ListBindingsResponseBodyData) *ListBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListBindingsResponseBody) SetRequestId(v string) *ListBindingsResponseBody {
	s.RequestId = &v
	return s
}

type ListBindingsResponseBodyData struct {
	Bindings   []*ListBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	MaxResults *int32                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBodyData) SetBindings(v []*ListBindingsResponseBodyDataBindings) *ListBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListBindingsResponseBodyData) SetMaxResults(v int32) *ListBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListBindingsResponseBodyData) SetNextToken(v string) *ListBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

type ListBindingsResponseBodyDataBindings struct {
	Argument        *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
}

func (s ListBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBodyDataBindings) SetArgument(v string) *ListBindingsResponseBodyDataBindings {
	s.Argument = &v
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

func (s *ListBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

type ListBindingsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListBindingsResponse) SetStatusCode(v int32) *ListBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindingsResponse) SetBody(v *ListBindingsResponseBody) *ListBindingsResponse {
	s.Body = v
	return s
}

type ListDownStreamBindingsRequest struct {
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	VirtualHost  *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListDownStreamBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsRequest) SetExchangeName(v string) *ListDownStreamBindingsRequest {
	s.ExchangeName = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetInstanceId(v string) *ListDownStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetMaxResults(v int32) *ListDownStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetNextToken(v string) *ListDownStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDownStreamBindingsRequest) SetVirtualHost(v string) *ListDownStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

type ListDownStreamBindingsResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDownStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDownStreamBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBody) SetCode(v int32) *ListDownStreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetData(v *ListDownStreamBindingsResponseBodyData) *ListDownStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetMessage(v string) *ListDownStreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetRequestId(v string) *ListDownStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDownStreamBindingsResponseBody) SetSuccess(v bool) *ListDownStreamBindingsResponseBody {
	s.Success = &v
	return s
}

type ListDownStreamBindingsResponseBodyData struct {
	Bindings   []*ListDownStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	MaxResults *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDownStreamBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBodyData) SetBindings(v []*ListDownStreamBindingsResponseBodyDataBindings) *ListDownStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListDownStreamBindingsResponseBodyData) SetMaxResults(v int32) *ListDownStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyData) SetNextToken(v string) *ListDownStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

type ListDownStreamBindingsResponseBodyDataBindings struct {
	Argument        *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
}

func (s ListDownStreamBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListDownStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
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

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListDownStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListDownStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

type ListDownStreamBindingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDownStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDownStreamBindingsResponse) SetStatusCode(v int32) *ListDownStreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDownStreamBindingsResponse) SetBody(v *ListDownStreamBindingsResponseBody) *ListDownStreamBindingsResponse {
	s.Body = v
	return s
}

type ListExchangeUpStreamBindingsRequest struct {
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	VirtualHost  *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s ListExchangeUpStreamBindingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsRequest) SetExchangeName(v string) *ListExchangeUpStreamBindingsRequest {
	s.ExchangeName = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetInstanceId(v string) *ListExchangeUpStreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetMaxResults(v int32) *ListExchangeUpStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetNextToken(v string) *ListExchangeUpStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExchangeUpStreamBindingsRequest) SetVirtualHost(v string) *ListExchangeUpStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

type ListExchangeUpStreamBindingsResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListExchangeUpStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetCode(v int32) *ListExchangeUpStreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetData(v *ListExchangeUpStreamBindingsResponseBodyData) *ListExchangeUpStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetMessage(v string) *ListExchangeUpStreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetRequestId(v string) *ListExchangeUpStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBody) SetSuccess(v bool) *ListExchangeUpStreamBindingsResponseBody {
	s.Success = &v
	return s
}

type ListExchangeUpStreamBindingsResponseBodyData struct {
	Bindings   []*ListExchangeUpStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	MaxResults *int32                                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetBindings(v []*ListExchangeUpStreamBindingsResponseBodyDataBindings) *ListExchangeUpStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetMaxResults(v int32) *ListExchangeUpStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyData) SetNextToken(v string) *ListExchangeUpStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

type ListExchangeUpStreamBindingsResponseBodyDataBindings struct {
	Argument        *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
}

func (s ListExchangeUpStreamBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListExchangeUpStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
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

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListExchangeUpStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

type ListExchangeUpStreamBindingsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListExchangeUpStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListExchangeUpStreamBindingsResponse) SetStatusCode(v int32) *ListExchangeUpStreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExchangeUpStreamBindingsResponse) SetBody(v *ListExchangeUpStreamBindingsResponseBody) *ListExchangeUpStreamBindingsResponse {
	s.Body = v
	return s
}

type ListExchangesRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
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

func (s *ListExchangesRequest) SetMaxResults(v int32) *ListExchangesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExchangesRequest) SetNextToken(v string) *ListExchangesRequest {
	s.NextToken = &v
	return s
}

func (s *ListExchangesRequest) SetVirtualHost(v string) *ListExchangesRequest {
	s.VirtualHost = &v
	return s
}

type ListExchangesResponseBody struct {
	Data      *ListExchangesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExchangesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBody) SetData(v *ListExchangesResponseBodyData) *ListExchangesResponseBody {
	s.Data = v
	return s
}

func (s *ListExchangesResponseBody) SetRequestId(v string) *ListExchangesResponseBody {
	s.RequestId = &v
	return s
}

type ListExchangesResponseBodyData struct {
	// Exchange。
	Exchanges  []*ListExchangesResponseBodyDataExchanges `json:"Exchanges,omitempty" xml:"Exchanges,omitempty" type:"Repeated"`
	MaxResults *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListExchangesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBodyData) SetExchanges(v []*ListExchangesResponseBodyDataExchanges) *ListExchangesResponseBodyData {
	s.Exchanges = v
	return s
}

func (s *ListExchangesResponseBodyData) SetMaxResults(v int32) *ListExchangesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListExchangesResponseBodyData) SetNextToken(v string) *ListExchangesResponseBodyData {
	s.NextToken = &v
	return s
}

type ListExchangesResponseBodyDataExchanges struct {
	Attributes      map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	AutoDeleteState *bool                  `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	CreateTime      *int64                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExchangeType    *string                `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
	Name            *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	VHostName       *string                `json:"VHostName,omitempty" xml:"VHostName,omitempty"`
}

func (s ListExchangesResponseBodyDataExchanges) String() string {
	return tea.Prettify(s)
}

func (s ListExchangesResponseBodyDataExchanges) GoString() string {
	return s.String()
}

func (s *ListExchangesResponseBodyDataExchanges) SetAttributes(v map[string]interface{}) *ListExchangesResponseBodyDataExchanges {
	s.Attributes = v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetAutoDeleteState(v bool) *ListExchangesResponseBodyDataExchanges {
	s.AutoDeleteState = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetCreateTime(v int64) *ListExchangesResponseBodyDataExchanges {
	s.CreateTime = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetExchangeType(v string) *ListExchangesResponseBodyDataExchanges {
	s.ExchangeType = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetName(v string) *ListExchangesResponseBodyDataExchanges {
	s.Name = &v
	return s
}

func (s *ListExchangesResponseBodyDataExchanges) SetVHostName(v string) *ListExchangesResponseBodyDataExchanges {
	s.VHostName = &v
	return s
}

type ListExchangesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListExchangesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListExchangesResponse) SetStatusCode(v int32) *ListExchangesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExchangesResponse) SetBody(v *ListExchangesResponseBody) *ListExchangesResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetMaxResults(v int32) *ListInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesRequest) SetNextToken(v string) *ListInstancesRequest {
	s.NextToken = &v
	return s
}

type ListInstancesResponseBody struct {
	Data      *ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancesResponseBodyData struct {
	Instances  []*ListInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	MaxResults *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) SetInstances(v []*ListInstancesResponseBodyDataInstances) *ListInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBodyData) SetMaxResults(v int32) *ListInstancesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetNextToken(v string) *ListInstancesResponseBodyData {
	s.NextToken = &v
	return s
}

type ListInstancesResponseBodyDataInstances struct {
	AutoRenewInstance *bool                                         `json:"AutoRenewInstance,omitempty" xml:"AutoRenewInstance,omitempty"`
	ClassicEndpoint   *string                                       `json:"ClassicEndpoint,omitempty" xml:"ClassicEndpoint,omitempty"`
	ExpireTime        *int64                                        `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId        *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string                                       `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType      *string                                       `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaxEipTps         *int32                                        `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	MaxQueue          *int32                                        `json:"MaxQueue,omitempty" xml:"MaxQueue,omitempty"`
	MaxTps            *int32                                        `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	MaxVhost          *int32                                        `json:"MaxVhost,omitempty" xml:"MaxVhost,omitempty"`
	OrderCreateTime   *int64                                        `json:"OrderCreateTime,omitempty" xml:"OrderCreateTime,omitempty"`
	OrderType         *string                                       `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PrivateEndpoint   *string                                       `json:"PrivateEndpoint,omitempty" xml:"PrivateEndpoint,omitempty"`
	PublicEndpoint    *string                                       `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
	Status            *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageSize       *int32                                        `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	SupportEIP        *bool                                         `json:"SupportEIP,omitempty" xml:"SupportEIP,omitempty"`
	Tags              []*ListInstancesResponseBodyDataInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyDataInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataInstances) SetAutoRenewInstance(v bool) *ListInstancesResponseBodyDataInstances {
	s.AutoRenewInstance = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetClassicEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.ClassicEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetExpireTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.ExpireTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceId(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceName(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceType(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxEipTps(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxEipTps = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxQueue(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxQueue = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxTps(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxTps = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxVhost(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxVhost = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetOrderCreateTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.OrderCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetOrderType(v string) *ListInstancesResponseBodyDataInstances {
	s.OrderType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetPrivateEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.PrivateEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetPublicEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.PublicEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetStatus(v string) *ListInstancesResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetStorageSize(v int32) *ListInstancesResponseBodyDataInstances {
	s.StorageSize = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetSupportEIP(v bool) *ListInstancesResponseBodyDataInstances {
	s.SupportEIP = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetTags(v []*ListInstancesResponseBodyDataInstancesTags) *ListInstancesResponseBodyDataInstances {
	s.Tags = v
	return s
}

type ListInstancesResponseBodyDataInstancesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyDataInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataInstancesTags) SetKey(v string) *ListInstancesResponseBodyDataInstancesTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstancesTags) SetValue(v string) *ListInstancesResponseBodyDataInstancesTags {
	s.Value = &v
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

type ListQueueConsumersRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	QueryCount  *int32  `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	Queue       *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
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

func (s *ListQueueConsumersRequest) SetNextToken(v string) *ListQueueConsumersRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueueConsumersRequest) SetQueryCount(v int32) *ListQueueConsumersRequest {
	s.QueryCount = &v
	return s
}

func (s *ListQueueConsumersRequest) SetQueue(v string) *ListQueueConsumersRequest {
	s.Queue = &v
	return s
}

func (s *ListQueueConsumersRequest) SetVirtualHost(v string) *ListQueueConsumersRequest {
	s.VirtualHost = &v
	return s
}

type ListQueueConsumersResponseBody struct {
	Data      *ListQueueConsumersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueueConsumersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBody) SetData(v *ListQueueConsumersResponseBodyData) *ListQueueConsumersResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueConsumersResponseBody) SetRequestId(v string) *ListQueueConsumersResponseBody {
	s.RequestId = &v
	return s
}

type ListQueueConsumersResponseBodyData struct {
	Consumers  []*ListQueueConsumersResponseBodyDataConsumers `json:"Consumers,omitempty" xml:"Consumers,omitempty" type:"Repeated"`
	MaxResults *int32                                         `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListQueueConsumersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueueConsumersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponseBodyData) SetConsumers(v []*ListQueueConsumersResponseBodyDataConsumers) *ListQueueConsumersResponseBodyData {
	s.Consumers = v
	return s
}

func (s *ListQueueConsumersResponseBodyData) SetMaxResults(v int32) *ListQueueConsumersResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListQueueConsumersResponseBodyData) SetNextToken(v string) *ListQueueConsumersResponseBodyData {
	s.NextToken = &v
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQueueConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListQueueConsumersResponse) SetStatusCode(v int32) *ListQueueConsumersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueConsumersResponse) SetBody(v *ListQueueConsumersResponseBody) *ListQueueConsumersResponse {
	s.Body = v
	return s
}

type ListQueueUpStreamBindingsRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	QueueName   *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
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

func (s *ListQueueUpStreamBindingsRequest) SetMaxResults(v int32) *ListQueueUpStreamBindingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetNextToken(v string) *ListQueueUpStreamBindingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetQueueName(v string) *ListQueueUpStreamBindingsRequest {
	s.QueueName = &v
	return s
}

func (s *ListQueueUpStreamBindingsRequest) SetVirtualHost(v string) *ListQueueUpStreamBindingsRequest {
	s.VirtualHost = &v
	return s
}

type ListQueueUpStreamBindingsResponseBody struct {
	Data      *ListQueueUpStreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueueUpStreamBindingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBody) SetData(v *ListQueueUpStreamBindingsResponseBodyData) *ListQueueUpStreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBody) SetRequestId(v string) *ListQueueUpStreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

type ListQueueUpStreamBindingsResponseBodyData struct {
	Bindings   []*ListQueueUpStreamBindingsResponseBodyDataBindings `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	MaxResults *string                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListQueueUpStreamBindingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetBindings(v []*ListQueueUpStreamBindingsResponseBodyDataBindings) *ListQueueUpStreamBindingsResponseBodyData {
	s.Bindings = v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetMaxResults(v string) *ListQueueUpStreamBindingsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyData) SetNextToken(v string) *ListQueueUpStreamBindingsResponseBodyData {
	s.NextToken = &v
	return s
}

type ListQueueUpStreamBindingsResponseBodyDataBindings struct {
	Argument        *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	BindingKey      *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType     *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty" xml:"DestinationName,omitempty"`
	SourceExchange  *string `json:"SourceExchange,omitempty" xml:"SourceExchange,omitempty"`
}

func (s ListQueueUpStreamBindingsResponseBodyDataBindings) String() string {
	return tea.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponseBodyDataBindings) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetArgument(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.Argument = &v
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

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetDestinationName(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.DestinationName = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponseBodyDataBindings) SetSourceExchange(v string) *ListQueueUpStreamBindingsResponseBodyDataBindings {
	s.SourceExchange = &v
	return s
}

type ListQueueUpStreamBindingsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQueueUpStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListQueueUpStreamBindingsResponse) SetStatusCode(v int32) *ListQueueUpStreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponse) SetBody(v *ListQueueUpStreamBindingsResponseBody) *ListQueueUpStreamBindingsResponse {
	s.Body = v
	return s
}

type ListQueuesRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
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

func (s *ListQueuesRequest) SetMaxResults(v int32) *ListQueuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQueuesRequest) SetNextToken(v string) *ListQueuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueuesRequest) SetVirtualHost(v string) *ListQueuesRequest {
	s.VirtualHost = &v
	return s
}

type ListQueuesResponseBody struct {
	Data      *ListQueuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBody) SetData(v *ListQueuesResponseBodyData) *ListQueuesResponseBody {
	s.Data = v
	return s
}

func (s *ListQueuesResponseBody) SetRequestId(v string) *ListQueuesResponseBody {
	s.RequestId = &v
	return s
}

type ListQueuesResponseBodyData struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Queue。
	Queues []*ListQueuesResponseBodyDataQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyData) SetMaxResults(v int32) *ListQueuesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListQueuesResponseBodyData) SetNextToken(v string) *ListQueuesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQueuesResponseBodyData) SetQueues(v []*ListQueuesResponseBodyDataQueues) *ListQueuesResponseBodyData {
	s.Queues = v
	return s
}

type ListQueuesResponseBodyDataQueues struct {
	Attributes      map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	AutoDeleteState *bool                  `json:"AutoDeleteState,omitempty" xml:"AutoDeleteState,omitempty"`
	CreateTime      *int64                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExclusiveState  *bool                  `json:"ExclusiveState,omitempty" xml:"ExclusiveState,omitempty"`
	LastConsumeTime *int64                 `json:"LastConsumeTime,omitempty" xml:"LastConsumeTime,omitempty"`
	Name            *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	VHostName       *string                `json:"VHostName,omitempty" xml:"VHostName,omitempty"`
}

func (s ListQueuesResponseBodyDataQueues) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyDataQueues) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyDataQueues) SetAttributes(v map[string]interface{}) *ListQueuesResponseBodyDataQueues {
	s.Attributes = v
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

func (s *ListQueuesResponseBodyDataQueues) SetExclusiveState(v bool) *ListQueuesResponseBodyDataQueues {
	s.ExclusiveState = &v
	return s
}

func (s *ListQueuesResponseBodyDataQueues) SetLastConsumeTime(v int64) *ListQueuesResponseBodyDataQueues {
	s.LastConsumeTime = &v
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

func (s *ListQueuesResponseBodyDataQueues) SetVHostName(v string) *ListQueuesResponseBodyDataQueues {
	s.VHostName = &v
	return s
}

type ListQueuesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListQueuesResponse) SetStatusCode(v int32) *ListQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueuesResponse) SetBody(v *ListQueuesResponseBody) *ListQueuesResponse {
	s.Body = v
	return s
}

type ListVirtualHostsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

func (s *ListVirtualHostsRequest) SetMaxResults(v int32) *ListVirtualHostsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualHostsRequest) SetNextToken(v string) *ListVirtualHostsRequest {
	s.NextToken = &v
	return s
}

type ListVirtualHostsResponseBody struct {
	Data      *ListVirtualHostsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVirtualHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBody) SetData(v *ListVirtualHostsResponseBodyData) *ListVirtualHostsResponseBody {
	s.Data = v
	return s
}

func (s *ListVirtualHostsResponseBody) SetRequestId(v string) *ListVirtualHostsResponseBody {
	s.RequestId = &v
	return s
}

type ListVirtualHostsResponseBodyData struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Vhost。
	VirtualHosts []*ListVirtualHostsResponseBodyDataVirtualHosts `json:"VirtualHosts,omitempty" xml:"VirtualHosts,omitempty" type:"Repeated"`
}

func (s ListVirtualHostsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualHostsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBodyData) SetMaxResults(v int32) *ListVirtualHostsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualHostsResponseBodyData) SetNextToken(v string) *ListVirtualHostsResponseBodyData {
	s.NextToken = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVirtualHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListVirtualHostsResponse) SetStatusCode(v int32) *ListVirtualHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualHostsResponse) SetBody(v *ListVirtualHostsResponseBody) *ListVirtualHostsResponse {
	s.Body = v
	return s
}

type UpdateInstanceNameRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s UpdateInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameRequest) SetInstanceId(v string) *UpdateInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceNameRequest) SetInstanceName(v string) *UpdateInstanceNameRequest {
	s.InstanceName = &v
	return s
}

type UpdateInstanceNameResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponseBody) SetCode(v int32) *UpdateInstanceNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetData(v string) *UpdateInstanceNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetMessage(v string) *UpdateInstanceNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetRequestId(v string) *UpdateInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetSuccess(v bool) *UpdateInstanceNameResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponse) SetHeaders(v map[string]*string) *UpdateInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNameResponse) SetStatusCode(v int32) *UpdateInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponse) SetBody(v *UpdateInstanceNameResponseBody) *UpdateInstanceNameResponse {
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

func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountAccessKey)) {
		query["accountAccessKey"] = request.AccountAccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimestamp)) {
		query["createTimestamp"] = request.CreateTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretSign)) {
		query["secretSign"] = request.SecretSign
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["userName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccount"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Argument)) {
		body["Argument"] = request.Argument
	}

	if !tea.BoolValue(util.IsUnset(request.BindingKey)) {
		body["BindingKey"] = request.BindingKey
	}

	if !tea.BoolValue(util.IsUnset(request.BindingType)) {
		body["BindingType"] = request.BindingType
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationName)) {
		body["DestinationName"] = request.DestinationName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceExchange)) {
		body["SourceExchange"] = request.SourceExchange
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBinding"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlternateExchange)) {
		body["AlternateExchange"] = request.AlternateExchange
	}

	if !tea.BoolValue(util.IsUnset(request.AutoDeleteState)) {
		body["AutoDeleteState"] = request.AutoDeleteState
	}

	if !tea.BoolValue(util.IsUnset(request.ExchangeName)) {
		body["ExchangeName"] = request.ExchangeName
	}

	if !tea.BoolValue(util.IsUnset(request.ExchangeType)) {
		body["ExchangeType"] = request.ExchangeType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Internal)) {
		body["Internal"] = request.Internal
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExchange"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExchangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConnections)) {
		query["MaxConnections"] = request.MaxConnections
	}

	if !tea.BoolValue(util.IsUnset(request.MaxEipTps)) {
		query["MaxEipTps"] = request.MaxEipTps
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPrivateTps)) {
		query["MaxPrivateTps"] = request.MaxPrivateTps
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		query["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodCycle)) {
		query["PeriodCycle"] = request.PeriodCycle
	}

	if !tea.BoolValue(util.IsUnset(request.QueueCapacity)) {
		query["QueueCapacity"] = request.QueueCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		query["StorageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SupportEip)) {
		query["SupportEip"] = request.SupportEip
	}

	if !tea.BoolValue(util.IsUnset(request.SupportTracing)) {
		query["SupportTracing"] = request.SupportTracing
	}

	if !tea.BoolValue(util.IsUnset(request.TracingStorageTime)) {
		query["TracingStorageTime"] = request.TracingStorageTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoDeleteState)) {
		body["AutoDeleteState"] = request.AutoDeleteState
	}

	if !tea.BoolValue(util.IsUnset(request.AutoExpireState)) {
		body["AutoExpireState"] = request.AutoExpireState
	}

	if !tea.BoolValue(util.IsUnset(request.DeadLetterExchange)) {
		body["DeadLetterExchange"] = request.DeadLetterExchange
	}

	if !tea.BoolValue(util.IsUnset(request.DeadLetterRoutingKey)) {
		body["DeadLetterRoutingKey"] = request.DeadLetterRoutingKey
	}

	if !tea.BoolValue(util.IsUnset(request.ExclusiveState)) {
		body["ExclusiveState"] = request.ExclusiveState
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxLength)) {
		body["MaxLength"] = request.MaxLength
	}

	if !tea.BoolValue(util.IsUnset(request.MaximumPriority)) {
		body["MaximumPriority"] = request.MaximumPriority
	}

	if !tea.BoolValue(util.IsUnset(request.MessageTTL)) {
		body["MessageTTL"] = request.MessageTTL
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		body["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQueue"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVirtualHost"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVirtualHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimestamp)) {
		query["CreateTimestamp"] = request.CreateTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccount"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindingKey)) {
		body["BindingKey"] = request.BindingKey
	}

	if !tea.BoolValue(util.IsUnset(request.BindingType)) {
		body["BindingType"] = request.BindingType
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationName)) {
		body["DestinationName"] = request.DestinationName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceExchange)) {
		body["SourceExchange"] = request.SourceExchange
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBinding"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExchangeName)) {
		body["ExchangeName"] = request.ExchangeName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExchange"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExchangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		body["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQueue"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["VirtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVirtualHost"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVirtualHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMetadataAmount"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMetadataAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListAccountsWithOptions(request *ListAccountsRequest, runtime *util.RuntimeOptions) (_result *ListAccountsResponse, _err error) {
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
		Action:      tea.String("ListAccounts"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccounts(request *ListAccountsRequest) (_result *ListAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountsResponse{}
	_body, _err := client.ListAccountsWithOptions(request, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBindings"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDownStreamBindings"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDownStreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListExchangeUpStreamBindingsWithOptions(request *ListExchangeUpStreamBindingsRequest, runtime *util.RuntimeOptions) (_result *ListExchangeUpStreamBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExchangeUpStreamBindings"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExchangeUpStreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListExchangesWithOptions(request *ListExchangesRequest, runtime *util.RuntimeOptions) (_result *ListExchangesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExchanges"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExchangesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueueConsumers"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueueConsumersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListQueueUpStreamBindingsWithOptions(request *ListQueueUpStreamBindingsRequest, runtime *util.RuntimeOptions) (_result *ListQueueUpStreamBindingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueueUpStreamBindings"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueueUpStreamBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListQueuesWithOptions(request *ListQueuesRequest, runtime *util.RuntimeOptions) (_result *ListQueuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueues"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListVirtualHostsWithOptions(request *ListVirtualHostsRequest, runtime *util.RuntimeOptions) (_result *ListVirtualHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVirtualHosts"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVirtualHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateInstanceNameWithOptions(request *UpdateInstanceNameRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceNameResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceName"),
		Version:     tea.String("2019-12-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceName(request *UpdateInstanceNameRequest) (_result *UpdateInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.UpdateInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

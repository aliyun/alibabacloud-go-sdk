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

type EndpointStatus struct {
	Code    *string               `json:"Code,omitempty" xml:"Code,omitempty"`
	Detail  *EndpointStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Message *string               `json:"Message,omitempty" xml:"Message,omitempty"`
	Phase   *string               `json:"Phase,omitempty" xml:"Phase,omitempty"`
}

func (s EndpointStatus) String() string {
	return tea.Prettify(s)
}

func (s EndpointStatus) GoString() string {
	return s.String()
}

func (s *EndpointStatus) SetCode(v string) *EndpointStatus {
	s.Code = &v
	return s
}

func (s *EndpointStatus) SetDetail(v *EndpointStatusDetail) *EndpointStatus {
	s.Detail = v
	return s
}

func (s *EndpointStatus) SetMessage(v string) *EndpointStatus {
	s.Message = &v
	return s
}

func (s *EndpointStatus) SetPhase(v string) *EndpointStatus {
	s.Phase = &v
	return s
}

type EndpointStatusDetail struct {
	IpPortMapping map[string]*IpPort `json:"IpPortMapping,omitempty" xml:"IpPortMapping,omitempty"`
}

func (s EndpointStatusDetail) String() string {
	return tea.Prettify(s)
}

func (s EndpointStatusDetail) GoString() string {
	return s.String()
}

func (s *EndpointStatusDetail) SetIpPortMapping(v map[string]*IpPort) *EndpointStatusDetail {
	s.IpPortMapping = v
	return s
}

type InstanceLifeCycle struct {
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s InstanceLifeCycle) String() string {
	return tea.Prettify(s)
}

func (s InstanceLifeCycle) GoString() string {
	return s.String()
}

func (s *InstanceLifeCycle) SetConfig(v string) *InstanceLifeCycle {
	s.Config = &v
	return s
}

func (s *InstanceLifeCycle) SetType(v string) *InstanceLifeCycle {
	s.Type = &v
	return s
}

type InstanceStatus struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Phase        *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	SlotNum      *int32  `json:"SlotNum,omitempty" xml:"SlotNum,omitempty"`
	UsedCapacity *string `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s InstanceStatus) String() string {
	return tea.Prettify(s)
}

func (s InstanceStatus) GoString() string {
	return s.String()
}

func (s *InstanceStatus) SetCode(v string) *InstanceStatus {
	s.Code = &v
	return s
}

func (s *InstanceStatus) SetMessage(v string) *InstanceStatus {
	s.Message = &v
	return s
}

func (s *InstanceStatus) SetPhase(v string) *InstanceStatus {
	s.Phase = &v
	return s
}

func (s *InstanceStatus) SetSlotNum(v int32) *InstanceStatus {
	s.SlotNum = &v
	return s
}

func (s *InstanceStatus) SetUsedCapacity(v string) *InstanceStatus {
	s.UsedCapacity = &v
	return s
}

type IpPort struct {
	Ip   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s IpPort) String() string {
	return tea.Prettify(s)
}

func (s IpPort) GoString() string {
	return s.String()
}

func (s *IpPort) SetIp(v string) *IpPort {
	s.Ip = &v
	return s
}

func (s *IpPort) SetPort(v string) *IpPort {
	s.Port = &v
	return s
}

type Metric struct {
	Timestamp *string  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Metric) String() string {
	return tea.Prettify(s)
}

func (s Metric) GoString() string {
	return s.String()
}

func (s *Metric) SetTimestamp(v string) *Metric {
	s.Timestamp = &v
	return s
}

func (s *Metric) SetValue(v float64) *Metric {
	s.Value = &v
	return s
}

type SlotLifeCycle struct {
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SlotLifeCycle) String() string {
	return tea.Prettify(s)
}

func (s SlotLifeCycle) GoString() string {
	return s.String()
}

func (s *SlotLifeCycle) SetConfig(v string) *SlotLifeCycle {
	s.Config = &v
	return s
}

func (s *SlotLifeCycle) SetType(v string) *SlotLifeCycle {
	s.Type = &v
	return s
}

type SlotStatus struct {
	Code    *string           `json:"Code,omitempty" xml:"Code,omitempty"`
	Detail  *SlotStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Message *string           `json:"Message,omitempty" xml:"Message,omitempty"`
	Phase   *string           `json:"Phase,omitempty" xml:"Phase,omitempty"`
}

func (s SlotStatus) String() string {
	return tea.Prettify(s)
}

func (s SlotStatus) GoString() string {
	return s.String()
}

func (s *SlotStatus) SetCode(v string) *SlotStatus {
	s.Code = &v
	return s
}

func (s *SlotStatus) SetDetail(v *SlotStatusDetail) *SlotStatus {
	s.Detail = v
	return s
}

func (s *SlotStatus) SetMessage(v string) *SlotStatus {
	s.Message = &v
	return s
}

func (s *SlotStatus) SetPhase(v string) *SlotStatus {
	s.Phase = &v
	return s
}

type SlotStatusDetail struct {
	LoadedFileNum   *int64  `json:"LoadedFileNum,omitempty" xml:"LoadedFileNum,omitempty"`
	LoadedFileSize  *string `json:"LoadedFileSize,omitempty" xml:"LoadedFileSize,omitempty"`
	LoadingTimeCost *int64  `json:"LoadingTimeCost,omitempty" xml:"LoadingTimeCost,omitempty"`
}

func (s SlotStatusDetail) String() string {
	return tea.Prettify(s)
}

func (s SlotStatusDetail) GoString() string {
	return s.String()
}

func (s *SlotStatusDetail) SetLoadedFileNum(v int64) *SlotStatusDetail {
	s.LoadedFileNum = &v
	return s
}

func (s *SlotStatusDetail) SetLoadedFileSize(v string) *SlotStatusDetail {
	s.LoadedFileSize = &v
	return s
}

func (s *SlotStatusDetail) SetLoadingTimeCost(v int64) *SlotStatusDetail {
	s.LoadingTimeCost = &v
	return s
}

type BindEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *BindEndpointResponseBody) SetRequestId(v string) *BindEndpointResponseBody {
	s.RequestId = &v
	return s
}

type BindEndpointResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s BindEndpointResponse) GoString() string {
	return s.String()
}

func (s *BindEndpointResponse) SetHeaders(v map[string]*string) *BindEndpointResponse {
	s.Headers = v
	return s
}

func (s *BindEndpointResponse) SetStatusCode(v int32) *BindEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *BindEndpointResponse) SetBody(v *BindEndpointResponseBody) *BindEndpointResponse {
	s.Body = v
	return s
}

type CreateEndpointRequest struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointRequest) SetName(v string) *CreateEndpointRequest {
	s.Name = &v
	return s
}

func (s *CreateEndpointRequest) SetType(v string) *CreateEndpointRequest {
	s.Type = &v
	return s
}

func (s *CreateEndpointRequest) SetVpcId(v string) *CreateEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateEndpointRequest) SetVswitchId(v string) *CreateEndpointRequest {
	s.VswitchId = &v
	return s
}

type CreateEndpointResponseBody struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponseBody) SetEndpointId(v string) *CreateEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *CreateEndpointResponseBody) SetRequestId(v string) *CreateEndpointResponseBody {
	s.RequestId = &v
	return s
}

type CreateEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponse) SetHeaders(v map[string]*string) *CreateEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateEndpointResponse) SetStatusCode(v int32) *CreateEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEndpointResponse) SetBody(v *CreateEndpointResponseBody) *CreateEndpointResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	Capacity    *string                      `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Description *string                      `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxSlot     *string                      `json:"MaxSlot,omitempty" xml:"MaxSlot,omitempty"`
	Name        *string                      `json:"Name,omitempty" xml:"Name,omitempty"`
	PaymentType *string                      `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	Tags        []*CreateInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type        *string                      `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetCapacity(v string) *CreateInstanceRequest {
	s.Capacity = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxSlot(v string) *CreateInstanceRequest {
	s.MaxSlot = &v
	return s
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreateInstanceRequest) SetType(v string) *CreateInstanceRequest {
	s.Type = &v
	return s
}

type CreateInstanceRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTags) SetKey(v string) *CreateInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTags) SetValue(v string) *CreateInstanceRequestTags {
	s.Value = &v
	return s
}

type CreateInstanceResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
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

type CreateSlotRequest struct {
	Capacity    *string                       `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Description *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointIds *string                       `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty"`
	Endpoints   []*CreateSlotRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	InstanceId  *string                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LifeCycle   *SlotLifeCycle                `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	Name        *string                       `json:"Name,omitempty" xml:"Name,omitempty"`
	StorageType *string                       `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	StorageUri  *string                       `json:"StorageUri,omitempty" xml:"StorageUri,omitempty"`
	Tags        []*CreateSlotRequestTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateSlotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSlotRequest) GoString() string {
	return s.String()
}

func (s *CreateSlotRequest) SetCapacity(v string) *CreateSlotRequest {
	s.Capacity = &v
	return s
}

func (s *CreateSlotRequest) SetDescription(v string) *CreateSlotRequest {
	s.Description = &v
	return s
}

func (s *CreateSlotRequest) SetEndpointIds(v string) *CreateSlotRequest {
	s.EndpointIds = &v
	return s
}

func (s *CreateSlotRequest) SetEndpoints(v []*CreateSlotRequestEndpoints) *CreateSlotRequest {
	s.Endpoints = v
	return s
}

func (s *CreateSlotRequest) SetInstanceId(v string) *CreateSlotRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSlotRequest) SetLifeCycle(v *SlotLifeCycle) *CreateSlotRequest {
	s.LifeCycle = v
	return s
}

func (s *CreateSlotRequest) SetName(v string) *CreateSlotRequest {
	s.Name = &v
	return s
}

func (s *CreateSlotRequest) SetStorageType(v string) *CreateSlotRequest {
	s.StorageType = &v
	return s
}

func (s *CreateSlotRequest) SetStorageUri(v string) *CreateSlotRequest {
	s.StorageUri = &v
	return s
}

func (s *CreateSlotRequest) SetTags(v []*CreateSlotRequestTags) *CreateSlotRequest {
	s.Tags = v
	return s
}

type CreateSlotRequestEndpoints struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateSlotRequestEndpoints) String() string {
	return tea.Prettify(s)
}

func (s CreateSlotRequestEndpoints) GoString() string {
	return s.String()
}

func (s *CreateSlotRequestEndpoints) SetName(v string) *CreateSlotRequestEndpoints {
	s.Name = &v
	return s
}

func (s *CreateSlotRequestEndpoints) SetType(v string) *CreateSlotRequestEndpoints {
	s.Type = &v
	return s
}

func (s *CreateSlotRequestEndpoints) SetVpcId(v string) *CreateSlotRequestEndpoints {
	s.VpcId = &v
	return s
}

func (s *CreateSlotRequestEndpoints) SetVswitchId(v string) *CreateSlotRequestEndpoints {
	s.VswitchId = &v
	return s
}

type CreateSlotRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSlotRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateSlotRequestTags) GoString() string {
	return s.String()
}

func (s *CreateSlotRequestTags) SetKey(v string) *CreateSlotRequestTags {
	s.Key = &v
	return s
}

func (s *CreateSlotRequestTags) SetValue(v string) *CreateSlotRequestTags {
	s.Value = &v
	return s
}

type CreateSlotResponseBody struct {
	EndpointIds *string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlotId      *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
}

func (s CreateSlotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSlotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlotResponseBody) SetEndpointIds(v string) *CreateSlotResponseBody {
	s.EndpointIds = &v
	return s
}

func (s *CreateSlotResponseBody) SetRequestId(v string) *CreateSlotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSlotResponseBody) SetSlotId(v string) *CreateSlotResponseBody {
	s.SlotId = &v
	return s
}

type CreateSlotResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSlotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSlotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSlotResponse) GoString() string {
	return s.String()
}

func (s *CreateSlotResponse) SetHeaders(v map[string]*string) *CreateSlotResponse {
	s.Headers = v
	return s
}

func (s *CreateSlotResponse) SetStatusCode(v int32) *CreateSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSlotResponse) SetBody(v *CreateSlotResponseBody) *CreateSlotResponse {
	s.Body = v
	return s
}

type CreateTagRequest struct {
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTagRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) SetKey(v string) *CreateTagRequest {
	s.Key = &v
	return s
}

func (s *CreateTagRequest) SetResourceId(v string) *CreateTagRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateTagRequest) SetResourceType(v string) *CreateTagRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateTagRequest) SetValue(v string) *CreateTagRequest {
	s.Value = &v
	return s
}

type CreateTagResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}

type CreateTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTagResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponse) GoString() string {
	return s.String()
}

func (s *CreateTagResponse) SetHeaders(v map[string]*string) *CreateTagResponse {
	s.Headers = v
	return s
}

func (s *CreateTagResponse) SetStatusCode(v int32) *CreateTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTagResponse) SetBody(v *CreateTagResponseBody) *CreateTagResponse {
	s.Body = v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
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

type DeleteSlotRequest struct {
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s DeleteSlotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSlotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSlotRequest) SetForce(v bool) *DeleteSlotRequest {
	s.Force = &v
	return s
}

type DeleteSlotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSlotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSlotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSlotResponseBody) SetRequestId(v string) *DeleteSlotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSlotResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSlotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSlotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSlotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSlotResponse) SetHeaders(v map[string]*string) *DeleteSlotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSlotResponse) SetStatusCode(v int32) *DeleteSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSlotResponse) SetBody(v *DeleteSlotResponseBody) *DeleteSlotResponse {
	s.Body = v
	return s
}

type DeleteTagRequest struct {
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DeleteTagRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagRequest) SetKey(v string) *DeleteTagRequest {
	s.Key = &v
	return s
}

func (s *DeleteTagRequest) SetResourceId(v string) *DeleteTagRequest {
	s.ResourceId = &v
	return s
}

func (s *DeleteTagRequest) SetResourceType(v string) *DeleteTagRequest {
	s.ResourceType = &v
	return s
}

type DeleteTagResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagResponseBody) SetRequestId(v string) *DeleteTagResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagResponse) SetHeaders(v map[string]*string) *DeleteTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagResponse) SetStatusCode(v int32) *DeleteTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTagResponse) SetBody(v *DeleteTagResponseBody) *DeleteTagResponse {
	s.Body = v
	return s
}

type DescribeComponentRequest struct {
	RenderTemplate *bool                  `json:"RenderTemplate,omitempty" xml:"RenderTemplate,omitempty"`
	Values         map[string]interface{} `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentRequest) SetRenderTemplate(v bool) *DescribeComponentRequest {
	s.RenderTemplate = &v
	return s
}

func (s *DescribeComponentRequest) SetValues(v map[string]interface{}) *DescribeComponentRequest {
	s.Values = v
	return s
}

type DescribeComponentShrinkRequest struct {
	RenderTemplate *bool   `json:"RenderTemplate,omitempty" xml:"RenderTemplate,omitempty"`
	ValuesShrink   *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeComponentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentShrinkRequest) SetRenderTemplate(v bool) *DescribeComponentShrinkRequest {
	s.RenderTemplate = &v
	return s
}

func (s *DescribeComponentShrinkRequest) SetValuesShrink(v string) *DescribeComponentShrinkRequest {
	s.ValuesShrink = &v
	return s
}

type DescribeComponentResponseBody struct {
	GmtCreateTime   *string                                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RenderedContent *string                                `json:"RenderedContent,omitempty" xml:"RenderedContent,omitempty"`
	RequestId       *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template        *DescribeComponentResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	UserId          *string                                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid            *string                                `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Version         *string                                `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentResponseBody) SetGmtCreateTime(v string) *DescribeComponentResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribeComponentResponseBody) SetGmtModifiedTime(v string) *DescribeComponentResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *DescribeComponentResponseBody) SetName(v string) *DescribeComponentResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeComponentResponseBody) SetOwnerId(v string) *DescribeComponentResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeComponentResponseBody) SetRenderedContent(v string) *DescribeComponentResponseBody {
	s.RenderedContent = &v
	return s
}

func (s *DescribeComponentResponseBody) SetRequestId(v string) *DescribeComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentResponseBody) SetTemplate(v *DescribeComponentResponseBodyTemplate) *DescribeComponentResponseBody {
	s.Template = v
	return s
}

func (s *DescribeComponentResponseBody) SetUserId(v string) *DescribeComponentResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeComponentResponseBody) SetUuid(v string) *DescribeComponentResponseBody {
	s.Uuid = &v
	return s
}

func (s *DescribeComponentResponseBody) SetVersion(v string) *DescribeComponentResponseBody {
	s.Version = &v
	return s
}

type DescribeComponentResponseBodyTemplate struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri  *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeComponentResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *DescribeComponentResponseBodyTemplate) SetType(v string) *DescribeComponentResponseBodyTemplate {
	s.Type = &v
	return s
}

func (s *DescribeComponentResponseBodyTemplate) SetUri(v string) *DescribeComponentResponseBodyTemplate {
	s.Uri = &v
	return s
}

type DescribeComponentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeComponentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentResponse) SetHeaders(v map[string]*string) *DescribeComponentResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentResponse) SetStatusCode(v int32) *DescribeComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentResponse) SetBody(v *DescribeComponentResponseBody) *DescribeComponentResponse {
	s.Body = v
	return s
}

type DescribeEndpointResponseBody struct {
	GmtCreateTime   *string         `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string         `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string         `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RequestId       *string         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status          *EndpointStatus `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *string         `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId          *string         `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid            *string         `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	VpcId           *string         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId       *string         `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointResponseBody) SetGmtCreateTime(v string) *DescribeEndpointResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetGmtModifiedTime(v string) *DescribeEndpointResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetName(v string) *DescribeEndpointResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetOwnerId(v string) *DescribeEndpointResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetRequestId(v string) *DescribeEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetStatus(v *EndpointStatus) *DescribeEndpointResponseBody {
	s.Status = v
	return s
}

func (s *DescribeEndpointResponseBody) SetType(v string) *DescribeEndpointResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetUserId(v string) *DescribeEndpointResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetUuid(v string) *DescribeEndpointResponseBody {
	s.Uuid = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetVpcId(v string) *DescribeEndpointResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeEndpointResponseBody) SetVswitchId(v string) *DescribeEndpointResponseBody {
	s.VswitchId = &v
	return s
}

type DescribeEndpointResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointResponse) SetHeaders(v map[string]*string) *DescribeEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointResponse) SetStatusCode(v int32) *DescribeEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndpointResponse) SetBody(v *DescribeEndpointResponseBody) *DescribeEndpointResponse {
	s.Body = v
	return s
}

type DescribeInstanceResponseBody struct {
	Capacity        *string                             `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Description     *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string                             `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                             `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	MaxSlot         *int32                              `json:"MaxSlot,omitempty" xml:"MaxSlot,omitempty"`
	Name            *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PaymentType     *string                             `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	RequestId       *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status          *InstanceStatus                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            []*DescribeInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type            *string                             `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId          *string                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid            *string                             `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetCapacity(v string) *DescribeInstanceResponseBody {
	s.Capacity = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetDescription(v string) *DescribeInstanceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetGmtCreateTime(v string) *DescribeInstanceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetGmtModifiedTime(v string) *DescribeInstanceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMaxSlot(v int32) *DescribeInstanceResponseBody {
	s.MaxSlot = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetName(v string) *DescribeInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetOwnerId(v string) *DescribeInstanceResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPaymentType(v string) *DescribeInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v *InstanceStatus) *DescribeInstanceResponseBody {
	s.Status = v
	return s
}

func (s *DescribeInstanceResponseBody) SetTags(v []*DescribeInstanceResponseBodyTags) *DescribeInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeInstanceResponseBody) SetType(v string) *DescribeInstanceResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetUserId(v string) *DescribeInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetUuid(v string) *DescribeInstanceResponseBody {
	s.Uuid = &v
	return s
}

type DescribeInstanceResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyTags) SetKey(v string) *DescribeInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeInstanceResponseBodyTags) SetValue(v string) *DescribeInstanceResponseBodyTags {
	s.Value = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeSlotResponseBody struct {
	Capacity        *string                         `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Description     *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string                         `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                         `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InstanceId      *string                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LifeCycle       *SlotLifeCycle                  `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	Name            *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RequestId       *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status          *SlotStatus                     `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType     *string                         `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	StorageUri      *string                         `json:"StorageUri,omitempty" xml:"StorageUri,omitempty"`
	Tags            []*DescribeSlotResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UserId          *string                         `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid            *string                         `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSlotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlotResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlotResponseBody) SetCapacity(v string) *DescribeSlotResponseBody {
	s.Capacity = &v
	return s
}

func (s *DescribeSlotResponseBody) SetDescription(v string) *DescribeSlotResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeSlotResponseBody) SetGmtCreateTime(v string) *DescribeSlotResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribeSlotResponseBody) SetGmtModifiedTime(v string) *DescribeSlotResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *DescribeSlotResponseBody) SetInstanceId(v string) *DescribeSlotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlotResponseBody) SetLifeCycle(v *SlotLifeCycle) *DescribeSlotResponseBody {
	s.LifeCycle = v
	return s
}

func (s *DescribeSlotResponseBody) SetName(v string) *DescribeSlotResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeSlotResponseBody) SetOwnerId(v string) *DescribeSlotResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlotResponseBody) SetRequestId(v string) *DescribeSlotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlotResponseBody) SetStatus(v *SlotStatus) *DescribeSlotResponseBody {
	s.Status = v
	return s
}

func (s *DescribeSlotResponseBody) SetStorageType(v string) *DescribeSlotResponseBody {
	s.StorageType = &v
	return s
}

func (s *DescribeSlotResponseBody) SetStorageUri(v string) *DescribeSlotResponseBody {
	s.StorageUri = &v
	return s
}

func (s *DescribeSlotResponseBody) SetTags(v []*DescribeSlotResponseBodyTags) *DescribeSlotResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeSlotResponseBody) SetUserId(v string) *DescribeSlotResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeSlotResponseBody) SetUuid(v string) *DescribeSlotResponseBody {
	s.Uuid = &v
	return s
}

type DescribeSlotResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSlotResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlotResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeSlotResponseBodyTags) SetKey(v string) *DescribeSlotResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeSlotResponseBodyTags) SetValue(v string) *DescribeSlotResponseBodyTags {
	s.Value = &v
	return s
}

type DescribeSlotResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSlotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlotResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlotResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlotResponse) SetHeaders(v map[string]*string) *DescribeSlotResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlotResponse) SetStatusCode(v int32) *DescribeSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlotResponse) SetBody(v *DescribeSlotResponseBody) *DescribeSlotResponse {
	s.Body = v
	return s
}

type ListComponentsRequest struct {
	ComponentIds *string `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy       *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentsRequest) SetComponentIds(v string) *ListComponentsRequest {
	s.ComponentIds = &v
	return s
}

func (s *ListComponentsRequest) SetName(v string) *ListComponentsRequest {
	s.Name = &v
	return s
}

func (s *ListComponentsRequest) SetOrder(v string) *ListComponentsRequest {
	s.Order = &v
	return s
}

func (s *ListComponentsRequest) SetPageNumber(v int32) *ListComponentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComponentsRequest) SetPageSize(v int32) *ListComponentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListComponentsRequest) SetSortBy(v string) *ListComponentsRequest {
	s.SortBy = &v
	return s
}

func (s *ListComponentsRequest) SetVersion(v string) *ListComponentsRequest {
	s.Version = &v
	return s
}

type ListComponentsResponseBody struct {
	Components []*ListComponentsResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListComponentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBody) SetComponents(v []*ListComponentsResponseBodyComponents) *ListComponentsResponseBody {
	s.Components = v
	return s
}

func (s *ListComponentsResponseBody) SetRequestId(v string) *ListComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComponentsResponseBody) SetTotalCount(v int32) *ListComponentsResponseBody {
	s.TotalCount = &v
	return s
}

type ListComponentsResponseBodyComponents struct {
	GmtCreateTime   *string                                       `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                                       `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Template        *ListComponentsResponseBodyComponentsTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	UserId          *string                                       `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid            *string                                       `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Version         *string                                       `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListComponentsResponseBodyComponents) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyComponents) SetGmtCreateTime(v string) *ListComponentsResponseBodyComponents {
	s.GmtCreateTime = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetGmtModifiedTime(v string) *ListComponentsResponseBodyComponents {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetName(v string) *ListComponentsResponseBodyComponents {
	s.Name = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetOwnerId(v string) *ListComponentsResponseBodyComponents {
	s.OwnerId = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetTemplate(v *ListComponentsResponseBodyComponentsTemplate) *ListComponentsResponseBodyComponents {
	s.Template = v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetUserId(v string) *ListComponentsResponseBodyComponents {
	s.UserId = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetUuid(v string) *ListComponentsResponseBodyComponents {
	s.Uuid = &v
	return s
}

func (s *ListComponentsResponseBodyComponents) SetVersion(v string) *ListComponentsResponseBodyComponents {
	s.Version = &v
	return s
}

type ListComponentsResponseBodyComponentsTemplate struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri  *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ListComponentsResponseBodyComponentsTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyComponentsTemplate) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyComponentsTemplate) SetType(v string) *ListComponentsResponseBodyComponentsTemplate {
	s.Type = &v
	return s
}

func (s *ListComponentsResponseBodyComponentsTemplate) SetUri(v string) *ListComponentsResponseBodyComponentsTemplate {
	s.Uri = &v
	return s
}

type ListComponentsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponse) GoString() string {
	return s.String()
}

func (s *ListComponentsResponse) SetHeaders(v map[string]*string) *ListComponentsResponse {
	s.Headers = v
	return s
}

func (s *ListComponentsResponse) SetStatusCode(v int32) *ListComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComponentsResponse) SetBody(v *ListComponentsResponseBody) *ListComponentsResponse {
	s.Body = v
	return s
}

type ListEndpointsRequest struct {
	EndpointIds *string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SlotIds     *string `json:"SlotIds,omitempty" xml:"SlotIds,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListEndpointsRequest) SetEndpointIds(v string) *ListEndpointsRequest {
	s.EndpointIds = &v
	return s
}

func (s *ListEndpointsRequest) SetName(v string) *ListEndpointsRequest {
	s.Name = &v
	return s
}

func (s *ListEndpointsRequest) SetOrder(v string) *ListEndpointsRequest {
	s.Order = &v
	return s
}

func (s *ListEndpointsRequest) SetPageNumber(v int32) *ListEndpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEndpointsRequest) SetPageSize(v int32) *ListEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEndpointsRequest) SetSlotIds(v string) *ListEndpointsRequest {
	s.SlotIds = &v
	return s
}

func (s *ListEndpointsRequest) SetSortBy(v string) *ListEndpointsRequest {
	s.SortBy = &v
	return s
}

func (s *ListEndpointsRequest) SetType(v string) *ListEndpointsRequest {
	s.Type = &v
	return s
}

type ListEndpointsResponseBody struct {
	Endpoints  []*ListEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBody) SetEndpoints(v []*ListEndpointsResponseBodyEndpoints) *ListEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListEndpointsResponseBody) SetRequestId(v string) *ListEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEndpointsResponseBody) SetTotalCount(v int32) *ListEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEndpointsResponseBodyEndpoints struct {
	GmtCreateTime   *string         `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string         `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string         `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status          *EndpointStatus `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *string         `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId          *string         `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid            *string         `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	VpcId           *string         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId       *string         `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListEndpointsResponseBodyEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBodyEndpoints) SetGmtCreateTime(v string) *ListEndpointsResponseBodyEndpoints {
	s.GmtCreateTime = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetGmtModifiedTime(v string) *ListEndpointsResponseBodyEndpoints {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetName(v string) *ListEndpointsResponseBodyEndpoints {
	s.Name = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetOwnerId(v string) *ListEndpointsResponseBodyEndpoints {
	s.OwnerId = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetStatus(v *EndpointStatus) *ListEndpointsResponseBodyEndpoints {
	s.Status = v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetType(v string) *ListEndpointsResponseBodyEndpoints {
	s.Type = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetUserId(v string) *ListEndpointsResponseBodyEndpoints {
	s.UserId = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetUuid(v string) *ListEndpointsResponseBodyEndpoints {
	s.Uuid = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetVpcId(v string) *ListEndpointsResponseBodyEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListEndpointsResponseBodyEndpoints) SetVswitchId(v string) *ListEndpointsResponseBodyEndpoints {
	s.VswitchId = &v
	return s
}

type ListEndpointsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponse) SetHeaders(v map[string]*string) *ListEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListEndpointsResponse) SetStatusCode(v int32) *ListEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEndpointsResponse) SetBody(v *ListEndpointsResponseBody) *ListEndpointsResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	Phase       *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetInstanceIds(v string) *ListInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstancesRequest) SetName(v string) *ListInstancesRequest {
	s.Name = &v
	return s
}

func (s *ListInstancesRequest) SetOrder(v string) *ListInstancesRequest {
	s.Order = &v
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

func (s *ListInstancesRequest) SetPaymentType(v string) *ListInstancesRequest {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesRequest) SetPhase(v string) *ListInstancesRequest {
	s.Phase = &v
	return s
}

func (s *ListInstancesRequest) SetSortBy(v string) *ListInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListInstancesRequest) SetType(v string) *ListInstancesRequest {
	s.Type = &v
	return s
}

type ListInstancesResponseBody struct {
	Instances  []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	Capacity        *string                                   `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Description     *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string                                   `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                                   `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	MaxSlot         *int32                                    `json:"MaxSlot,omitempty" xml:"MaxSlot,omitempty"`
	Name            *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PaymentType     *string                                   `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	Status          *InstanceStatus                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            []*ListInstancesResponseBodyInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type            *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId          *string                                   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid            *string                                   `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetCapacity(v string) *ListInstancesResponseBodyInstances {
	s.Capacity = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDescription(v string) *ListInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetMaxSlot(v int32) *ListInstancesResponseBodyInstances {
	s.MaxSlot = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetName(v string) *ListInstancesResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetOwnerId(v string) *ListInstancesResponseBodyInstances {
	s.OwnerId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPaymentType(v string) *ListInstancesResponseBodyInstances {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v *InstanceStatus) *ListInstancesResponseBodyInstances {
	s.Status = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetTags(v []*ListInstancesResponseBodyInstancesTags) *ListInstancesResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetType(v string) *ListInstancesResponseBodyInstances {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUserId(v string) *ListInstancesResponseBodyInstances {
	s.UserId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUuid(v string) *ListInstancesResponseBodyInstances {
	s.Uuid = &v
	return s
}

type ListInstancesResponseBodyInstancesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesTags) SetKey(v string) *ListInstancesResponseBodyInstancesTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesTags) SetValue(v string) *ListInstancesResponseBodyInstancesTags {
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

type ListSlotsRequest struct {
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Phase       *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	SlotIds     *string `json:"SlotIds,omitempty" xml:"SlotIds,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ListSlotsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSlotsRequest) GoString() string {
	return s.String()
}

func (s *ListSlotsRequest) SetInstanceIds(v string) *ListSlotsRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListSlotsRequest) SetName(v string) *ListSlotsRequest {
	s.Name = &v
	return s
}

func (s *ListSlotsRequest) SetOrder(v string) *ListSlotsRequest {
	s.Order = &v
	return s
}

func (s *ListSlotsRequest) SetPageNumber(v int32) *ListSlotsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSlotsRequest) SetPageSize(v int32) *ListSlotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSlotsRequest) SetPhase(v string) *ListSlotsRequest {
	s.Phase = &v
	return s
}

func (s *ListSlotsRequest) SetSlotIds(v string) *ListSlotsRequest {
	s.SlotIds = &v
	return s
}

func (s *ListSlotsRequest) SetSortBy(v string) *ListSlotsRequest {
	s.SortBy = &v
	return s
}

func (s *ListSlotsRequest) SetStorageType(v string) *ListSlotsRequest {
	s.StorageType = &v
	return s
}

type ListSlotsResponseBody struct {
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Slots      []*ListSlotsResponseBodySlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSlotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSlotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlotsResponseBody) SetRequestId(v string) *ListSlotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlotsResponseBody) SetSlots(v []*ListSlotsResponseBodySlots) *ListSlotsResponseBody {
	s.Slots = v
	return s
}

func (s *ListSlotsResponseBody) SetTotalCount(v int32) *ListSlotsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSlotsResponseBodySlots struct {
	Capacity        *string                                `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Description     *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Endpoints       []*ListSlotsResponseBodySlotsEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	GmtCreateTime   *string                                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InstanceId      *string                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LifeCycle       *SlotLifeCycle                         `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	Name            *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status          *SlotStatus                            `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType     *string                                `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	StorageUri      *string                                `json:"StorageUri,omitempty" xml:"StorageUri,omitempty"`
	Tags            []*ListSlotsResponseBodySlotsTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UserId          *string                                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid            *string                                `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListSlotsResponseBodySlots) String() string {
	return tea.Prettify(s)
}

func (s ListSlotsResponseBodySlots) GoString() string {
	return s.String()
}

func (s *ListSlotsResponseBodySlots) SetCapacity(v string) *ListSlotsResponseBodySlots {
	s.Capacity = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetDescription(v string) *ListSlotsResponseBodySlots {
	s.Description = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetEndpoints(v []*ListSlotsResponseBodySlotsEndpoints) *ListSlotsResponseBodySlots {
	s.Endpoints = v
	return s
}

func (s *ListSlotsResponseBodySlots) SetGmtCreateTime(v string) *ListSlotsResponseBodySlots {
	s.GmtCreateTime = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetGmtModifiedTime(v string) *ListSlotsResponseBodySlots {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetInstanceId(v string) *ListSlotsResponseBodySlots {
	s.InstanceId = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetLifeCycle(v *SlotLifeCycle) *ListSlotsResponseBodySlots {
	s.LifeCycle = v
	return s
}

func (s *ListSlotsResponseBodySlots) SetName(v string) *ListSlotsResponseBodySlots {
	s.Name = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetOwnerId(v string) *ListSlotsResponseBodySlots {
	s.OwnerId = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetStatus(v *SlotStatus) *ListSlotsResponseBodySlots {
	s.Status = v
	return s
}

func (s *ListSlotsResponseBodySlots) SetStorageType(v string) *ListSlotsResponseBodySlots {
	s.StorageType = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetStorageUri(v string) *ListSlotsResponseBodySlots {
	s.StorageUri = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetTags(v []*ListSlotsResponseBodySlotsTags) *ListSlotsResponseBodySlots {
	s.Tags = v
	return s
}

func (s *ListSlotsResponseBodySlots) SetUserId(v string) *ListSlotsResponseBodySlots {
	s.UserId = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetUuid(v string) *ListSlotsResponseBodySlots {
	s.Uuid = &v
	return s
}

type ListSlotsResponseBodySlotsEndpoints struct {
	GmtCreateTime   *string         `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string         `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Name            *string         `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status          *EndpointStatus `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *string         `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId          *string         `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid            *string         `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	VpcId           *string         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId       *string         `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListSlotsResponseBodySlotsEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ListSlotsResponseBodySlotsEndpoints) GoString() string {
	return s.String()
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetGmtCreateTime(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.GmtCreateTime = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetGmtModifiedTime(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetName(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.Name = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetOwnerId(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.OwnerId = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetStatus(v *EndpointStatus) *ListSlotsResponseBodySlotsEndpoints {
	s.Status = v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetType(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.Type = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetUserId(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.UserId = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetUuid(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.Uuid = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetVpcId(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetVswitchId(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.VswitchId = &v
	return s
}

type ListSlotsResponseBodySlotsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSlotsResponseBodySlotsTags) String() string {
	return tea.Prettify(s)
}

func (s ListSlotsResponseBodySlotsTags) GoString() string {
	return s.String()
}

func (s *ListSlotsResponseBodySlotsTags) SetKey(v string) *ListSlotsResponseBodySlotsTags {
	s.Key = &v
	return s
}

func (s *ListSlotsResponseBodySlotsTags) SetValue(v string) *ListSlotsResponseBodySlotsTags {
	s.Value = &v
	return s
}

type ListSlotsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSlotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSlotsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSlotsResponse) GoString() string {
	return s.String()
}

func (s *ListSlotsResponse) SetHeaders(v map[string]*string) *ListSlotsResponse {
	s.Headers = v
	return s
}

func (s *ListSlotsResponse) SetStatusCode(v int32) *ListSlotsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSlotsResponse) SetBody(v *ListSlotsResponseBody) *ListSlotsResponse {
	s.Body = v
	return s
}

type ListTagsRequest struct {
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SortBy       *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) SetKey(v string) *ListTagsRequest {
	s.Key = &v
	return s
}

func (s *ListTagsRequest) SetOrder(v string) *ListTagsRequest {
	s.Order = &v
	return s
}

func (s *ListTagsRequest) SetPageNumber(v int32) *ListTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagsRequest) SetPageSize(v int32) *ListTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagsRequest) SetResourceId(v string) *ListTagsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListTagsRequest) SetResourceType(v string) *ListTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagsRequest) SetSortBy(v string) *ListTagsRequest {
	s.SortBy = &v
	return s
}

func (s *ListTagsRequest) SetValue(v string) *ListTagsRequest {
	s.Value = &v
	return s
}

type ListTagsResponseBody struct {
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags       []*ListTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetTags(v []*ListTagsResponseBodyTags) *ListTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListTagsResponseBody) SetTotalCount(v int32) *ListTagsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagsResponseBodyTags struct {
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Key             *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Value           *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTags) SetGmtCreateTime(v string) *ListTagsResponseBodyTags {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetGmtModifiedTime(v string) *ListTagsResponseBodyTags {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetKey(v string) *ListTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetOwnerId(v string) *ListTagsResponseBodyTags {
	s.OwnerId = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetResourceId(v string) *ListTagsResponseBodyTags {
	s.ResourceId = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetResourceType(v string) *ListTagsResponseBodyTags {
	s.ResourceType = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetUserId(v string) *ListTagsResponseBodyTags {
	s.UserId = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetValue(v string) *ListTagsResponseBodyTags {
	s.Value = &v
	return s
}

type ListTagsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponse) GoString() string {
	return s.String()
}

func (s *ListTagsResponse) SetHeaders(v map[string]*string) *ListTagsResponse {
	s.Headers = v
	return s
}

func (s *ListTagsResponse) SetStatusCode(v int32) *ListTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsResponse) SetBody(v *ListTagsResponseBody) *ListTagsResponse {
	s.Body = v
	return s
}

type QueryInstanceMetricsRequest struct {
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EndTime    *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricType *string                `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	StartTime  *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeStep   *string                `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
}

func (s QueryInstanceMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceMetricsRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceMetricsRequest) SetDimensions(v map[string]interface{}) *QueryInstanceMetricsRequest {
	s.Dimensions = v
	return s
}

func (s *QueryInstanceMetricsRequest) SetEndTime(v string) *QueryInstanceMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryInstanceMetricsRequest) SetMetricType(v string) *QueryInstanceMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *QueryInstanceMetricsRequest) SetStartTime(v string) *QueryInstanceMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryInstanceMetricsRequest) SetTimeStep(v string) *QueryInstanceMetricsRequest {
	s.TimeStep = &v
	return s
}

type QueryInstanceMetricsShrinkRequest struct {
	DimensionsShrink *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricType       *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeStep         *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
}

func (s QueryInstanceMetricsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceMetricsShrinkRequest) SetDimensionsShrink(v string) *QueryInstanceMetricsShrinkRequest {
	s.DimensionsShrink = &v
	return s
}

func (s *QueryInstanceMetricsShrinkRequest) SetEndTime(v string) *QueryInstanceMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *QueryInstanceMetricsShrinkRequest) SetMetricType(v string) *QueryInstanceMetricsShrinkRequest {
	s.MetricType = &v
	return s
}

func (s *QueryInstanceMetricsShrinkRequest) SetStartTime(v string) *QueryInstanceMetricsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *QueryInstanceMetricsShrinkRequest) SetTimeStep(v string) *QueryInstanceMetricsShrinkRequest {
	s.TimeStep = &v
	return s
}

type QueryInstanceMetricsResponseBody struct {
	Metrics   []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	Period    *string   `json:"Period,omitempty" xml:"Period,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryInstanceMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceMetricsResponseBody) SetMetrics(v []*Metric) *QueryInstanceMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *QueryInstanceMetricsResponseBody) SetPeriod(v string) *QueryInstanceMetricsResponseBody {
	s.Period = &v
	return s
}

func (s *QueryInstanceMetricsResponseBody) SetRequestId(v string) *QueryInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

type QueryInstanceMetricsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryInstanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInstanceMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceMetricsResponse) SetHeaders(v map[string]*string) *QueryInstanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceMetricsResponse) SetStatusCode(v int32) *QueryInstanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceMetricsResponse) SetBody(v *QueryInstanceMetricsResponseBody) *QueryInstanceMetricsResponse {
	s.Body = v
	return s
}

type QuerySlotMetricsRequest struct {
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EndTime    *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricType *string                `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	StartTime  *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeStep   *string                `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
}

func (s QuerySlotMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySlotMetricsRequest) GoString() string {
	return s.String()
}

func (s *QuerySlotMetricsRequest) SetDimensions(v map[string]interface{}) *QuerySlotMetricsRequest {
	s.Dimensions = v
	return s
}

func (s *QuerySlotMetricsRequest) SetEndTime(v string) *QuerySlotMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySlotMetricsRequest) SetMetricType(v string) *QuerySlotMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *QuerySlotMetricsRequest) SetStartTime(v string) *QuerySlotMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySlotMetricsRequest) SetTimeStep(v string) *QuerySlotMetricsRequest {
	s.TimeStep = &v
	return s
}

type QuerySlotMetricsShrinkRequest struct {
	DimensionsShrink *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricType       *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeStep         *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
}

func (s QuerySlotMetricsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySlotMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySlotMetricsShrinkRequest) SetDimensionsShrink(v string) *QuerySlotMetricsShrinkRequest {
	s.DimensionsShrink = &v
	return s
}

func (s *QuerySlotMetricsShrinkRequest) SetEndTime(v string) *QuerySlotMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySlotMetricsShrinkRequest) SetMetricType(v string) *QuerySlotMetricsShrinkRequest {
	s.MetricType = &v
	return s
}

func (s *QuerySlotMetricsShrinkRequest) SetStartTime(v string) *QuerySlotMetricsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySlotMetricsShrinkRequest) SetTimeStep(v string) *QuerySlotMetricsShrinkRequest {
	s.TimeStep = &v
	return s
}

type QuerySlotMetricsResponseBody struct {
	Metrics   []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	Period    *string   `json:"Period,omitempty" xml:"Period,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySlotMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySlotMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySlotMetricsResponseBody) SetMetrics(v []*Metric) *QuerySlotMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *QuerySlotMetricsResponseBody) SetPeriod(v string) *QuerySlotMetricsResponseBody {
	s.Period = &v
	return s
}

func (s *QuerySlotMetricsResponseBody) SetRequestId(v string) *QuerySlotMetricsResponseBody {
	s.RequestId = &v
	return s
}

type QuerySlotMetricsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySlotMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySlotMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySlotMetricsResponse) GoString() string {
	return s.String()
}

func (s *QuerySlotMetricsResponse) SetHeaders(v map[string]*string) *QuerySlotMetricsResponse {
	s.Headers = v
	return s
}

func (s *QuerySlotMetricsResponse) SetStatusCode(v int32) *QuerySlotMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySlotMetricsResponse) SetBody(v *QuerySlotMetricsResponseBody) *QuerySlotMetricsResponse {
	s.Body = v
	return s
}

type QueryStatisticRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Fields    *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryStatisticRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticRequest) GoString() string {
	return s.String()
}

func (s *QueryStatisticRequest) SetEndTime(v string) *QueryStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *QueryStatisticRequest) SetFields(v string) *QueryStatisticRequest {
	s.Fields = &v
	return s
}

func (s *QueryStatisticRequest) SetStartTime(v string) *QueryStatisticRequest {
	s.StartTime = &v
	return s
}

type QueryStatisticResponseBody struct {
	InstanceCapacityEachType map[string]interface{} `json:"InstanceCapacityEachType,omitempty" xml:"InstanceCapacityEachType,omitempty"`
	InstanceNumEachType      map[string]interface{} `json:"InstanceNumEachType,omitempty" xml:"InstanceNumEachType,omitempty"`
	RequestId                *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlotNumEachType          map[string]interface{} `json:"SlotNumEachType,omitempty" xml:"SlotNumEachType,omitempty"`
}

func (s QueryStatisticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStatisticResponseBody) SetInstanceCapacityEachType(v map[string]interface{}) *QueryStatisticResponseBody {
	s.InstanceCapacityEachType = v
	return s
}

func (s *QueryStatisticResponseBody) SetInstanceNumEachType(v map[string]interface{}) *QueryStatisticResponseBody {
	s.InstanceNumEachType = v
	return s
}

func (s *QueryStatisticResponseBody) SetRequestId(v string) *QueryStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStatisticResponseBody) SetSlotNumEachType(v map[string]interface{}) *QueryStatisticResponseBody {
	s.SlotNumEachType = v
	return s
}

type QueryStatisticResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStatisticResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticResponse) GoString() string {
	return s.String()
}

func (s *QueryStatisticResponse) SetHeaders(v map[string]*string) *QueryStatisticResponse {
	s.Headers = v
	return s
}

func (s *QueryStatisticResponse) SetStatusCode(v int32) *QueryStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStatisticResponse) SetBody(v *QueryStatisticResponseBody) *QueryStatisticResponse {
	s.Body = v
	return s
}

type StopSlotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopSlotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopSlotResponseBody) GoString() string {
	return s.String()
}

func (s *StopSlotResponseBody) SetRequestId(v string) *StopSlotResponseBody {
	s.RequestId = &v
	return s
}

type StopSlotResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopSlotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopSlotResponse) String() string {
	return tea.Prettify(s)
}

func (s StopSlotResponse) GoString() string {
	return s.String()
}

func (s *StopSlotResponse) SetHeaders(v map[string]*string) *StopSlotResponse {
	s.Headers = v
	return s
}

func (s *StopSlotResponse) SetStatusCode(v int32) *StopSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSlotResponse) SetBody(v *StopSlotResponseBody) *StopSlotResponse {
	s.Body = v
	return s
}

type UnbindEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindEndpointResponseBody) SetRequestId(v string) *UnbindEndpointResponseBody {
	s.RequestId = &v
	return s
}

type UnbindEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindEndpointResponse) GoString() string {
	return s.String()
}

func (s *UnbindEndpointResponse) SetHeaders(v map[string]*string) *UnbindEndpointResponse {
	s.Headers = v
	return s
}

func (s *UnbindEndpointResponse) SetStatusCode(v int32) *UnbindEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindEndpointResponse) SetBody(v *UnbindEndpointResponseBody) *UnbindEndpointResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxSlot     *string `json:"MaxSlot,omitempty" xml:"MaxSlot,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetDescription(v string) *UpdateInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstanceRequest) SetMaxSlot(v string) *UpdateInstanceRequest {
	s.MaxSlot = &v
	return s
}

func (s *UpdateInstanceRequest) SetName(v string) *UpdateInstanceRequest {
	s.Name = &v
	return s
}

type UpdateInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
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

type UpdateSlotRequest struct {
	Capacity    *string                  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Description *string                  `json:"Description,omitempty" xml:"Description,omitempty"`
	LifeCycle   *SlotLifeCycle           `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	Name        *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	StorageType *string                  `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	StorageUri  *string                  `json:"StorageUri,omitempty" xml:"StorageUri,omitempty"`
	Tags        []*UpdateSlotRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateSlotRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlotRequest) GoString() string {
	return s.String()
}

func (s *UpdateSlotRequest) SetCapacity(v string) *UpdateSlotRequest {
	s.Capacity = &v
	return s
}

func (s *UpdateSlotRequest) SetDescription(v string) *UpdateSlotRequest {
	s.Description = &v
	return s
}

func (s *UpdateSlotRequest) SetLifeCycle(v *SlotLifeCycle) *UpdateSlotRequest {
	s.LifeCycle = v
	return s
}

func (s *UpdateSlotRequest) SetName(v string) *UpdateSlotRequest {
	s.Name = &v
	return s
}

func (s *UpdateSlotRequest) SetStorageType(v string) *UpdateSlotRequest {
	s.StorageType = &v
	return s
}

func (s *UpdateSlotRequest) SetStorageUri(v string) *UpdateSlotRequest {
	s.StorageUri = &v
	return s
}

func (s *UpdateSlotRequest) SetTags(v []*UpdateSlotRequestTags) *UpdateSlotRequest {
	s.Tags = v
	return s
}

type UpdateSlotRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateSlotRequestTags) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlotRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateSlotRequestTags) SetKey(v string) *UpdateSlotRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateSlotRequestTags) SetValue(v string) *UpdateSlotRequestTags {
	s.Value = &v
	return s
}

type UpdateSlotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSlotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlotResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSlotResponseBody) SetRequestId(v string) *UpdateSlotResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSlotResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSlotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSlotResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlotResponse) GoString() string {
	return s.String()
}

func (s *UpdateSlotResponse) SetHeaders(v map[string]*string) *UpdateSlotResponse {
	s.Headers = v
	return s
}

func (s *UpdateSlotResponse) SetStatusCode(v int32) *UpdateSlotResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSlotResponse) SetBody(v *UpdateSlotResponseBody) *UpdateSlotResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("paielasticdatasetaccelerator"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BindEndpointWithOptions(EndpointId *string, SlotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BindEndpointResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("BindEndpoint"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/endpoints/" + tea.StringValue(openapiutil.GetEncodeParam(EndpointId)) + "/slots/" + tea.StringValue(openapiutil.GetEncodeParam(SlotId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BindEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindEndpoint(EndpointId *string, SlotId *string) (_result *BindEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindEndpointResponse{}
	_body, _err := client.BindEndpointWithOptions(EndpointId, SlotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEndpointWithOptions(request *CreateEndpointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchId)) {
		body["VswitchId"] = request.VswitchId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEndpoint"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/endpoints"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEndpoint(request *CreateEndpointRequest) (_result *CreateEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEndpointResponse{}
	_body, _err := client.CreateEndpointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		body["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSlot)) {
		body["MaxSlot"] = request.MaxSlot
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		body["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSlotWithOptions(request *CreateSlotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSlotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		body["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointIds)) {
		body["EndpointIds"] = request.EndpointIds
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoints)) {
		body["Endpoints"] = request.Endpoints
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LifeCycle)) {
		body["LifeCycle"] = request.LifeCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		body["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.StorageUri)) {
		body["StorageUri"] = request.StorageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSlot"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/slots"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSlotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSlot(request *CreateSlotRequest) (_result *CreateSlotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSlotResponse{}
	_body, _err := client.CreateSlotWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTagWithOptions(request *CreateTagRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		body["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTag"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tags"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTag(request *CreateTagRequest) (_result *CreateTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTagResponse{}
	_body, _err := client.CreateTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId))),
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

func (client *Client) DeleteInstance(InstanceId *string) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSlotWithOptions(SlotId *string, request *DeleteSlotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSlotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSlot"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/slots/" + tea.StringValue(openapiutil.GetEncodeParam(SlotId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSlotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSlot(SlotId *string, request *DeleteSlotRequest) (_result *DeleteSlotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSlotResponse{}
	_body, _err := client.DeleteSlotWithOptions(SlotId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTagWithOptions(request *DeleteTagRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTag"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tags"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTag(request *DeleteTagRequest) (_result *DeleteTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTagResponse{}
	_body, _err := client.DeleteTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComponentWithOptions(ComponentId *string, tmpReq *DescribeComponentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeComponentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeComponentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Values)) {
		request.ValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Values, tea.String("Values"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RenderTemplate)) {
		query["RenderTemplate"] = request.RenderTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.ValuesShrink)) {
		query["Values"] = request.ValuesShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponent"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/components/" + tea.StringValue(openapiutil.GetEncodeParam(ComponentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComponentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComponent(ComponentId *string, request *DescribeComponentRequest) (_result *DescribeComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeComponentResponse{}
	_body, _err := client.DescribeComponentWithOptions(ComponentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEndpointWithOptions(EndpointId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeEndpointResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEndpoint"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/endpoints/" + tea.StringValue(openapiutil.GetEncodeParam(EndpointId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEndpoint(EndpointId *string) (_result *DescribeEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeEndpointResponse{}
	_body, _err := client.DescribeEndpointWithOptions(EndpointId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(InstanceId *string) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlotWithOptions(SlotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSlotResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlot"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/slots/" + tea.StringValue(openapiutil.GetEncodeParam(SlotId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlot(SlotId *string) (_result *DescribeSlotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSlotResponse{}
	_body, _err := client.DescribeSlotWithOptions(SlotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListComponentsWithOptions(request *ListComponentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComponentIds)) {
		query["ComponentIds"] = request.ComponentIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListComponents"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/components"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListComponentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListComponents(request *ListComponentsRequest) (_result *ListComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListComponentsResponse{}
	_body, _err := client.ListComponentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEndpointsWithOptions(request *ListEndpointsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointIds)) {
		query["EndpointIds"] = request.EndpointIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SlotIds)) {
		query["SlotIds"] = request.SlotIds
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEndpoints"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/endpoints"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEndpoints(request *ListEndpointsRequest) (_result *ListEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEndpointsResponse{}
	_body, _err := client.ListEndpointsWithOptions(request, headers, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		query["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		query["Phase"] = request.Phase
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances"),
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

func (client *Client) ListSlotsWithOptions(request *ListSlotsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSlotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Phase)) {
		query["Phase"] = request.Phase
	}

	if !tea.BoolValue(util.IsUnset(request.SlotIds)) {
		query["SlotIds"] = request.SlotIds
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSlots"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/slots"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSlotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSlots(request *ListSlotsRequest) (_result *ListSlotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSlotsResponse{}
	_body, _err := client.ListSlotsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagsWithOptions(request *ListTagsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTags"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInstanceMetricsWithOptions(InstanceId *string, tmpReq *QueryInstanceMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryInstanceMetricsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryInstanceMetricsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Dimensions)) {
		request.DimensionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dimensions, tea.String("Dimensions"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DimensionsShrink)) {
		query["Dimensions"] = request.DimensionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStep)) {
		query["TimeStep"] = request.TimeStep
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryInstanceMetrics"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId)) + "/metrics/action/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInstanceMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInstanceMetrics(InstanceId *string, request *QueryInstanceMetricsRequest) (_result *QueryInstanceMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryInstanceMetricsResponse{}
	_body, _err := client.QueryInstanceMetricsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySlotMetricsWithOptions(SlotId *string, tmpReq *QuerySlotMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QuerySlotMetricsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QuerySlotMetricsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Dimensions)) {
		request.DimensionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dimensions, tea.String("Dimensions"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DimensionsShrink)) {
		query["Dimensions"] = request.DimensionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStep)) {
		query["TimeStep"] = request.TimeStep
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySlotMetrics"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/slots/" + tea.StringValue(openapiutil.GetEncodeParam(SlotId)) + "/metrics/action/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySlotMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySlotMetrics(SlotId *string, request *QuerySlotMetricsRequest) (_result *QuerySlotMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySlotMetricsResponse{}
	_body, _err := client.QuerySlotMetricsWithOptions(SlotId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStatisticWithOptions(request *QueryStatisticRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryStatisticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		query["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryStatistic"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/action/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStatistic(request *QueryStatisticRequest) (_result *QueryStatisticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryStatisticResponse{}
	_body, _err := client.QueryStatisticWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopSlotWithOptions(SlotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopSlotResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopSlot"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/slots/" + tea.StringValue(openapiutil.GetEncodeParam(SlotId)) + "/action/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopSlotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopSlot(SlotId *string) (_result *StopSlotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopSlotResponse{}
	_body, _err := client.StopSlotWithOptions(SlotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindEndpointWithOptions(EndpointId *string, SlotId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnbindEndpointResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindEndpoint"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/endpoints/" + tea.StringValue(openapiutil.GetEncodeParam(EndpointId)) + "/slots/" + tea.StringValue(openapiutil.GetEncodeParam(SlotId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindEndpoint(EndpointId *string, SlotId *string) (_result *UnbindEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnbindEndpointResponse{}
	_body, _err := client.UnbindEndpointWithOptions(EndpointId, SlotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceWithOptions(InstanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSlot)) {
		body["MaxSlot"] = request.MaxSlot
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId))),
		Method:      tea.String("PUT"),
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

func (client *Client) UpdateInstance(InstanceId *string, request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSlotWithOptions(SlotId *string, request *UpdateSlotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSlotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		body["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.LifeCycle)) {
		body["LifeCycle"] = request.LifeCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		body["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.StorageUri)) {
		body["StorageUri"] = request.StorageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSlot"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/slots/" + tea.StringValue(openapiutil.GetEncodeParam(SlotId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSlotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSlot(SlotId *string, request *UpdateSlotRequest) (_result *UpdateSlotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSlotResponse{}
	_body, _err := client.UpdateSlotWithOptions(SlotId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

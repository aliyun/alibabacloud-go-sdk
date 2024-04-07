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

type ConvertInstanceRequest struct {
	ConvertPostpayInstanceRequest *ConvertInstanceRequestConvertPostpayInstanceRequest `json:"ConvertPostpayInstanceRequest,omitempty" xml:"ConvertPostpayInstanceRequest,omitempty" type:"Struct"`
}

func (s ConvertInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequest) SetConvertPostpayInstanceRequest(v *ConvertInstanceRequestConvertPostpayInstanceRequest) *ConvertInstanceRequest {
	s.ConvertPostpayInstanceRequest = v
	return s
}

type ConvertInstanceRequestConvertPostpayInstanceRequest struct {
	Duration               *int32                                                                       `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId             *string                                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsAutoRenew            *bool                                                                        `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	NamespaceResourceSpecs []*ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty" type:"Repeated"`
	PricingCycle           *string                                                                      `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Region                 *string                                                                      `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetDuration(v int32) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.Duration = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetInstanceId(v string) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetIsAutoRenew(v bool) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetNamespaceResourceSpecs(v []*ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.NamespaceResourceSpecs = v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetPricingCycle(v string) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetRegion(v string) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.Region = &v
	return s
}

type ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs struct {
	Namespace    *string                                                                                `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ResourceSpec *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) SetNamespace(v string) *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs {
	s.Namespace = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) SetResourceSpec(v *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs {
	s.ResourceSpec = v
	return s
}

type ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) SetCpu(v int32) *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) SetMemoryGB(v int32) *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.MemoryGB = &v
	return s
}

type ConvertInstanceResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConvertInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResponseBody) SetOrderId(v int64) *ConvertInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ConvertInstanceResponseBody) SetRequestId(v string) *ConvertInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertInstanceResponseBody) SetSuccess(v bool) *ConvertInstanceResponseBody {
	s.Success = &v
	return s
}

type ConvertInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConvertInstanceResponse) SetHeaders(v map[string]*string) *ConvertInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConvertInstanceResponse) SetStatusCode(v int32) *ConvertInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertInstanceResponse) SetBody(v *ConvertInstanceResponseBody) *ConvertInstanceResponse {
	s.Body = v
	return s
}

type ConvertPrepayInstanceRequest struct {
	ConvertPrepayInstanceRequest *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest `json:"ConvertPrepayInstanceRequest,omitempty" xml:"ConvertPrepayInstanceRequest,omitempty" type:"Struct"`
}

func (s ConvertPrepayInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertPrepayInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertPrepayInstanceRequest) SetConvertPrepayInstanceRequest(v *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) *ConvertPrepayInstanceRequest {
	s.ConvertPrepayInstanceRequest = v
	return s
}

type ConvertPrepayInstanceRequestConvertPrepayInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) SetInstanceId(v string) *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest) SetRegion(v string) *ConvertPrepayInstanceRequestConvertPrepayInstanceRequest {
	s.Region = &v
	return s
}

type ConvertPrepayInstanceResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConvertPrepayInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertPrepayInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertPrepayInstanceResponseBody) SetOrderId(v int64) *ConvertPrepayInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ConvertPrepayInstanceResponseBody) SetRequestId(v string) *ConvertPrepayInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertPrepayInstanceResponseBody) SetSuccess(v bool) *ConvertPrepayInstanceResponseBody {
	s.Success = &v
	return s
}

type ConvertPrepayInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertPrepayInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertPrepayInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertPrepayInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConvertPrepayInstanceResponse) SetHeaders(v map[string]*string) *ConvertPrepayInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConvertPrepayInstanceResponse) SetStatusCode(v int32) *ConvertPrepayInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertPrepayInstanceResponse) SetBody(v *ConvertPrepayInstanceResponseBody) *ConvertPrepayInstanceResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	CreateInstanceRequest *CreateInstanceRequestCreateInstanceRequest `json:"CreateInstanceRequest,omitempty" xml:"CreateInstanceRequest,omitempty" type:"Struct"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetCreateInstanceRequest(v *CreateInstanceRequestCreateInstanceRequest) *CreateInstanceRequest {
	s.CreateInstanceRequest = v
	return s
}

type CreateInstanceRequestCreateInstanceRequest struct {
	ArchitectureType *string                                                   `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	AutoRenew        *bool                                                     `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType       *string                                                   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Duration         *int32                                                    `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extra            *string                                                   `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Ha               *bool                                                     `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaResourceSpec   *CreateInstanceRequestCreateInstanceRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	HaVSwitchIds     []*string                                                 `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	HaZoneId         *string                                                   `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	InstanceName     *string                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MonitorType      *string                                                   `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	PricingCycle     *string                                                   `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	PromotionCode    *string                                                   `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	Region           *string                                                   `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId  *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceSpec     *CreateInstanceRequestCreateInstanceRequestResourceSpec   `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage          *CreateInstanceRequestCreateInstanceRequestStorage        `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	UsePromotionCode *bool                                                     `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	VSwitchIds       []*string                                                 `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// VPC ID。
	VpcId  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequestCreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetArchitectureType(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.ArchitectureType = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequestCreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetDuration(v int32) *CreateInstanceRequestCreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetExtra(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.Extra = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetHa(v bool) *CreateInstanceRequestCreateInstanceRequest {
	s.Ha = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetHaResourceSpec(v *CreateInstanceRequestCreateInstanceRequestHaResourceSpec) *CreateInstanceRequestCreateInstanceRequest {
	s.HaResourceSpec = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetHaVSwitchIds(v []*string) *CreateInstanceRequestCreateInstanceRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetHaZoneId(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.HaZoneId = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetMonitorType(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.MonitorType = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetPricingCycle(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetPromotionCode(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetRegion(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.Region = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetResourceSpec(v *CreateInstanceRequestCreateInstanceRequestResourceSpec) *CreateInstanceRequestCreateInstanceRequest {
	s.ResourceSpec = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetStorage(v *CreateInstanceRequestCreateInstanceRequestStorage) *CreateInstanceRequestCreateInstanceRequest {
	s.Storage = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetUsePromotionCode(v bool) *CreateInstanceRequestCreateInstanceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetVSwitchIds(v []*string) *CreateInstanceRequestCreateInstanceRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateInstanceRequestCreateInstanceRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateInstanceRequestCreateInstanceRequestHaResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCreateInstanceRequestHaResourceSpec) SetCpu(v int32) *CreateInstanceRequestCreateInstanceRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequestHaResourceSpec) SetMemoryGB(v int32) *CreateInstanceRequestCreateInstanceRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

type CreateInstanceRequestCreateInstanceRequestResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateInstanceRequestCreateInstanceRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCreateInstanceRequestResourceSpec) SetCpu(v int32) *CreateInstanceRequestCreateInstanceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequestResourceSpec) SetMemoryGB(v int32) *CreateInstanceRequestCreateInstanceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type CreateInstanceRequestCreateInstanceRequestStorage struct {
	Oss *CreateInstanceRequestCreateInstanceRequestStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s CreateInstanceRequestCreateInstanceRequestStorage) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequestStorage) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCreateInstanceRequestStorage) SetOss(v *CreateInstanceRequestCreateInstanceRequestStorageOss) *CreateInstanceRequestCreateInstanceRequestStorage {
	s.Oss = v
	return s
}

type CreateInstanceRequestCreateInstanceRequestStorageOss struct {
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s CreateInstanceRequestCreateInstanceRequestStorageOss) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequestStorageOss) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCreateInstanceRequestStorageOss) SetBucket(v string) *CreateInstanceRequestCreateInstanceRequestStorageOss {
	s.Bucket = &v
	return s
}

type CreateInstanceResponseBody struct {
	OrderInfo *CreateInstanceResponseBodyOrderInfo `json:"OrderInfo,omitempty" xml:"OrderInfo,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetOrderInfo(v *CreateInstanceResponseBodyOrderInfo) *CreateInstanceResponseBody {
	s.OrderInfo = v
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

type CreateInstanceResponseBodyOrderInfo struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateInstanceResponseBodyOrderInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyOrderInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyOrderInfo) SetInstanceId(v string) *CreateInstanceResponseBodyOrderInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyOrderInfo) SetOrderId(v int64) *CreateInstanceResponseBodyOrderInfo {
	s.OrderId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateNamespaceRequest struct {
	CreateNamespaceRequest *CreateNamespaceRequestCreateNamespaceRequest `json:"CreateNamespaceRequest,omitempty" xml:"CreateNamespaceRequest,omitempty" type:"Struct"`
}

func (s CreateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) SetCreateNamespaceRequest(v *CreateNamespaceRequestCreateNamespaceRequest) *CreateNamespaceRequest {
	s.CreateNamespaceRequest = v
	return s
}

type CreateNamespaceRequestCreateNamespaceRequest struct {
	Ha           *bool                                                     `json:"Ha,omitempty" xml:"Ha,omitempty"`
	InstanceId   *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Namespace    *string                                                   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Region       *string                                                   `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec *CreateNamespaceRequestCreateNamespaceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s CreateNamespaceRequestCreateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequestCreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) SetHa(v bool) *CreateNamespaceRequestCreateNamespaceRequest {
	s.Ha = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) SetInstanceId(v string) *CreateNamespaceRequestCreateNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) SetNamespace(v string) *CreateNamespaceRequestCreateNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) SetRegion(v string) *CreateNamespaceRequestCreateNamespaceRequest {
	s.Region = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) SetResourceSpec(v *CreateNamespaceRequestCreateNamespaceRequestResourceSpec) *CreateNamespaceRequestCreateNamespaceRequest {
	s.ResourceSpec = v
	return s
}

type CreateNamespaceRequestCreateNamespaceRequestResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateNamespaceRequestCreateNamespaceRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequestCreateNamespaceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequestCreateNamespaceRequestResourceSpec) SetCpu(v int32) *CreateNamespaceRequestCreateNamespaceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequestResourceSpec) SetMemoryGB(v int32) *CreateNamespaceRequestCreateNamespaceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type CreateNamespaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetSuccess(v bool) *CreateNamespaceResponseBody {
	s.Success = &v
	return s
}

type CreateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponse) SetHeaders(v map[string]*string) *CreateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateNamespaceResponse) SetStatusCode(v int32) *CreateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNamespaceResponse) SetBody(v *CreateNamespaceResponseBody) *CreateNamespaceResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	DeleteInstanceRequest *DeleteInstanceRequestDeleteInstanceRequest `json:"DeleteInstanceRequest,omitempty" xml:"DeleteInstanceRequest,omitempty" type:"Struct"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetDeleteInstanceRequest(v *DeleteInstanceRequestDeleteInstanceRequest) *DeleteInstanceRequest {
	s.DeleteInstanceRequest = v
	return s
}

type DeleteInstanceRequestDeleteInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteInstanceRequestDeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequestDeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequestDeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequestDeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequestDeleteInstanceRequest) SetRegion(v string) *DeleteInstanceRequestDeleteInstanceRequest {
	s.Region = &v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *DeleteInstanceResponseBody) SetSuccess(v bool) *DeleteInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteNamespaceRequest struct {
	DeleteNamespaceRequest *DeleteNamespaceRequestDeleteNamespaceRequest `json:"DeleteNamespaceRequest,omitempty" xml:"DeleteNamespaceRequest,omitempty" type:"Struct"`
}

func (s DeleteNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) SetDeleteNamespaceRequest(v *DeleteNamespaceRequestDeleteNamespaceRequest) *DeleteNamespaceRequest {
	s.DeleteNamespaceRequest = v
	return s
}

type DeleteNamespaceRequestDeleteNamespaceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteNamespaceRequestDeleteNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceRequestDeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequestDeleteNamespaceRequest) SetInstanceId(v string) *DeleteNamespaceRequestDeleteNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNamespaceRequestDeleteNamespaceRequest) SetNamespace(v string) *DeleteNamespaceRequestDeleteNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteNamespaceRequestDeleteNamespaceRequest) SetRegion(v string) *DeleteNamespaceRequestDeleteNamespaceRequest {
	s.Region = &v
	return s
}

type DeleteNamespaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponseBody) SetRequestId(v string) *DeleteNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetSuccess(v bool) *DeleteNamespaceResponseBody {
	s.Success = &v
	return s
}

type DeleteNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponse) SetHeaders(v map[string]*string) *DeleteNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNamespaceResponse) SetStatusCode(v int32) *DeleteNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNamespaceResponse) SetBody(v *DeleteNamespaceResponseBody) *DeleteNamespaceResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	DescribeInstancesRequest *DescribeInstancesRequestDescribeInstancesRequest `json:"DescribeInstancesRequest,omitempty" xml:"DescribeInstancesRequest,omitempty" type:"Struct"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetDescribeInstancesRequest(v *DescribeInstancesRequestDescribeInstancesRequest) *DescribeInstancesRequest {
	s.DescribeInstancesRequest = v
	return s
}

type DescribeInstancesRequestDescribeInstancesRequest struct {
	ArchitectureType *string                                                 `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	ChargeType       *string                                                 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceId       *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageIndex        *int32                                                  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize         *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Region           *string                                                 `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId  *string                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags             []*DescribeInstancesRequestDescribeInstancesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequestDescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestDescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetArchitectureType(v string) *DescribeInstancesRequestDescribeInstancesRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetChargeType(v string) *DescribeInstancesRequestDescribeInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequestDescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetPageIndex(v int32) *DescribeInstancesRequestDescribeInstancesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequestDescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetRegion(v string) *DescribeInstancesRequestDescribeInstancesRequest {
	s.Region = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequestDescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequest) SetTags(v []*DescribeInstancesRequestDescribeInstancesRequestTags) *DescribeInstancesRequestDescribeInstancesRequest {
	s.Tags = v
	return s
}

type DescribeInstancesRequestDescribeInstancesRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestDescribeInstancesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestDescribeInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestDescribeInstancesRequestTags) SetKey(v string) *DescribeInstancesRequestDescribeInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestDescribeInstancesRequestTags) SetValue(v string) *DescribeInstancesRequestDescribeInstancesRequestTags {
	s.Value = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Instances  []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	PageIndex  *int32                                    `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                    `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageIndex(v int32) *DescribeInstancesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetSuccess(v bool) *DescribeInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalPage(v int32) *DescribeInstancesResponseBody {
	s.TotalPage = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	ArchitectureType    *string                                               `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	AskClusterId        *string                                               `json:"AskClusterId,omitempty" xml:"AskClusterId,omitempty"`
	ChargeType          *string                                               `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterStatus       *string                                               `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	Ha                  *bool                                                 `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaResourceSpec      *DescribeInstancesResponseBodyInstancesHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	HaVSwitchIds        []*string                                             `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	HaZoneId            *string                                               `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	HostAliases         []*DescribeInstancesResponseBodyInstancesHostAliases  `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	InstanceId          *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName        *string                                               `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MonitorType         *string                                               `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	OrderState          *string                                               `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	Region              *string                                               `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCreateTime  *int64                                                `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	ResourceExpiredTime *int64                                                `json:"ResourceExpiredTime,omitempty" xml:"ResourceExpiredTime,omitempty"`
	ResourceGroupId     *string                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId          *string                                               `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceSpec        *DescribeInstancesResponseBodyInstancesResourceSpec   `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage             *DescribeInstancesResponseBodyInstancesStorage        `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	Tags                []*DescribeInstancesResponseBodyInstancesTags         `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Uid                 *string                                               `json:"Uid,omitempty" xml:"Uid,omitempty"`
	VSwitchIds          []*string                                             `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VpcId               *string                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId              *string                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetArchitectureType(v string) *DescribeInstancesResponseBodyInstances {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetAskClusterId(v string) *DescribeInstancesResponseBodyInstances {
	s.AskClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetChargeType(v string) *DescribeInstancesResponseBodyInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.ClusterStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHa(v bool) *DescribeInstancesResponseBodyInstances {
	s.Ha = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHaResourceSpec(v *DescribeInstancesResponseBodyInstancesHaResourceSpec) *DescribeInstancesResponseBodyInstances {
	s.HaResourceSpec = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHaVSwitchIds(v []*string) *DescribeInstancesResponseBodyInstances {
	s.HaVSwitchIds = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHaZoneId(v string) *DescribeInstancesResponseBodyInstances {
	s.HaZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHostAliases(v []*DescribeInstancesResponseBodyInstancesHostAliases) *DescribeInstancesResponseBodyInstances {
	s.HostAliases = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceName(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetMonitorType(v string) *DescribeInstancesResponseBodyInstances {
	s.MonitorType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetOrderState(v string) *DescribeInstancesResponseBodyInstances {
	s.OrderState = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRegion(v string) *DescribeInstancesResponseBodyInstances {
	s.Region = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceCreateTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ResourceCreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceExpiredTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ResourceExpiredTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceSpec(v *DescribeInstancesResponseBodyInstancesResourceSpec) *DescribeInstancesResponseBodyInstances {
	s.ResourceSpec = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStorage(v *DescribeInstancesResponseBodyInstancesStorage) *DescribeInstancesResponseBodyInstances {
	s.Storage = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetTags(v []*DescribeInstancesResponseBodyInstancesTags) *DescribeInstancesResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUid(v string) *DescribeInstancesResponseBodyInstances {
	s.Uid = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVSwitchIds(v []*string) *DescribeInstancesResponseBodyInstances {
	s.VSwitchIds = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetZoneId(v string) *DescribeInstancesResponseBodyInstances {
	s.ZoneId = &v
	return s
}

type DescribeInstancesResponseBodyInstancesHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesHaResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesHaResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) SetCpu(v int32) *DescribeInstancesResponseBodyInstancesHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) SetMemoryGB(v int32) *DescribeInstancesResponseBodyInstancesHaResourceSpec {
	s.MemoryGB = &v
	return s
}

type DescribeInstancesResponseBodyInstancesHostAliases struct {
	HostNames []*string `json:"HostNames,omitempty" xml:"HostNames,omitempty" type:"Repeated"`
	Ip        *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesHostAliases) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesHostAliases) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) SetHostNames(v []*string) *DescribeInstancesResponseBodyInstancesHostAliases {
	s.HostNames = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) SetIp(v string) *DescribeInstancesResponseBodyInstancesHostAliases {
	s.Ip = &v
	return s
}

type DescribeInstancesResponseBodyInstancesResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) SetCpu(v int32) *DescribeInstancesResponseBodyInstancesResourceSpec {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) SetMemoryGB(v int32) *DescribeInstancesResponseBodyInstancesResourceSpec {
	s.MemoryGB = &v
	return s
}

type DescribeInstancesResponseBodyInstancesStorage struct {
	Oss *DescribeInstancesResponseBodyInstancesStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s DescribeInstancesResponseBodyInstancesStorage) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesStorage) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesStorage) SetOss(v *DescribeInstancesResponseBodyInstancesStorageOss) *DescribeInstancesResponseBodyInstancesStorage {
	s.Oss = v
	return s
}

type DescribeInstancesResponseBodyInstancesStorageOss struct {
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesStorageOss) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesStorageOss) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesStorageOss) SetBucket(v string) *DescribeInstancesResponseBodyInstancesStorageOss {
	s.Bucket = &v
	return s
}

type DescribeInstancesResponseBodyInstancesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesTags) SetKey(v string) *DescribeInstancesResponseBodyInstancesTags {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesTags) SetValue(v string) *DescribeInstancesResponseBodyInstancesTags {
	s.Value = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetStatusCode(v int32) *DescribeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeNamespacesRequest struct {
	DescribeNamespacesRequest *DescribeNamespacesRequestDescribeNamespacesRequest `json:"DescribeNamespacesRequest,omitempty" xml:"DescribeNamespacesRequest,omitempty" type:"Struct"`
}

func (s DescribeNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequest) SetDescribeNamespacesRequest(v *DescribeNamespacesRequestDescribeNamespacesRequest) *DescribeNamespacesRequest {
	s.DescribeNamespacesRequest = v
	return s
}

type DescribeNamespacesRequestDescribeNamespacesRequest struct {
	Ha         *bool                                                     `json:"Ha,omitempty" xml:"Ha,omitempty"`
	InstanceId *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Namespace  *string                                                   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PageIndex  *int32                                                    `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Region     *string                                                   `json:"Region,omitempty" xml:"Region,omitempty"`
	Tags       []*DescribeNamespacesRequestDescribeNamespacesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeNamespacesRequestDescribeNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesRequestDescribeNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetHa(v bool) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.Ha = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetInstanceId(v string) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetNamespace(v string) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetPageIndex(v int32) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetPageSize(v int32) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetRegion(v string) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.Region = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetTags(v []*DescribeNamespacesRequestDescribeNamespacesRequestTags) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.Tags = v
	return s
}

type DescribeNamespacesRequestDescribeNamespacesRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNamespacesRequestDescribeNamespacesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesRequestDescribeNamespacesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequestTags) SetKey(v string) *DescribeNamespacesRequestDescribeNamespacesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequestTags) SetValue(v string) *DescribeNamespacesRequestDescribeNamespacesRequestTags {
	s.Value = &v
	return s
}

type DescribeNamespacesResponseBody struct {
	Namespaces []*DescribeNamespacesResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	PageIndex  *int32                                      `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                      `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBody) SetNamespaces(v []*DescribeNamespacesResponseBodyNamespaces) *DescribeNamespacesResponseBody {
	s.Namespaces = v
	return s
}

func (s *DescribeNamespacesResponseBody) SetPageIndex(v int32) *DescribeNamespacesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetPageSize(v int32) *DescribeNamespacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetRequestId(v string) *DescribeNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetSuccess(v bool) *DescribeNamespacesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetTotalCount(v int64) *DescribeNamespacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetTotalPage(v int32) *DescribeNamespacesResponseBody {
	s.TotalPage = &v
	return s
}

type DescribeNamespacesResponseBodyNamespaces struct {
	GmtCreate    *int64                                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *int64                                                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Ha           *bool                                                 `json:"Ha,omitempty" xml:"Ha,omitempty"`
	Namespace    *string                                               `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ResourceSpec *DescribeNamespacesResponseBodyNamespacesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	ResourceUsed *DescribeNamespacesResponseBodyNamespacesResourceUsed `json:"ResourceUsed,omitempty" xml:"ResourceUsed,omitempty" type:"Struct"`
	Status       *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags         []*DescribeNamespacesResponseBodyNamespacesTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeNamespacesResponseBodyNamespaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetGmtCreate(v int64) *DescribeNamespacesResponseBodyNamespaces {
	s.GmtCreate = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetGmtModified(v int64) *DescribeNamespacesResponseBodyNamespaces {
	s.GmtModified = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetHa(v bool) *DescribeNamespacesResponseBodyNamespaces {
	s.Ha = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetNamespace(v string) *DescribeNamespacesResponseBodyNamespaces {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetResourceSpec(v *DescribeNamespacesResponseBodyNamespacesResourceSpec) *DescribeNamespacesResponseBodyNamespaces {
	s.ResourceSpec = v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetResourceUsed(v *DescribeNamespacesResponseBodyNamespacesResourceUsed) *DescribeNamespacesResponseBodyNamespaces {
	s.ResourceUsed = v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetStatus(v string) *DescribeNamespacesResponseBodyNamespaces {
	s.Status = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetTags(v []*DescribeNamespacesResponseBodyNamespacesTags) *DescribeNamespacesResponseBodyNamespaces {
	s.Tags = v
	return s
}

type DescribeNamespacesResponseBodyNamespacesResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeNamespacesResponseBodyNamespacesResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponseBodyNamespacesResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceSpec) SetCpu(v int32) *DescribeNamespacesResponseBodyNamespacesResourceSpec {
	s.Cpu = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceSpec) SetMemoryGB(v int32) *DescribeNamespacesResponseBodyNamespacesResourceSpec {
	s.MemoryGB = &v
	return s
}

type DescribeNamespacesResponseBodyNamespacesResourceUsed struct {
	Cpu      *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Cu       *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	MemoryGB *float32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeNamespacesResponseBodyNamespacesResourceUsed) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponseBodyNamespacesResourceUsed) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) SetCpu(v float32) *DescribeNamespacesResponseBodyNamespacesResourceUsed {
	s.Cpu = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) SetCu(v float32) *DescribeNamespacesResponseBodyNamespacesResourceUsed {
	s.Cu = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) SetMemoryGB(v float32) *DescribeNamespacesResponseBodyNamespacesResourceUsed {
	s.MemoryGB = &v
	return s
}

type DescribeNamespacesResponseBodyNamespacesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNamespacesResponseBodyNamespacesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponseBodyNamespacesTags) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyNamespacesTags) SetKey(v string) *DescribeNamespacesResponseBodyNamespacesTags {
	s.Key = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesTags) SetValue(v string) *DescribeNamespacesResponseBodyNamespacesTags {
	s.Value = &v
	return s
}

type DescribeNamespacesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponse) SetHeaders(v map[string]*string) *DescribeNamespacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespacesResponse) SetStatusCode(v int32) *DescribeNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespacesResponse) SetBody(v *DescribeNamespacesResponseBody) *DescribeNamespacesResponse {
	s.Body = v
	return s
}

type DescribeSupportedRegionsResponseBody struct {
	Regions   []*DescribeSupportedRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSupportedRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportedRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportedRegionsResponseBody) SetRegions(v []*DescribeSupportedRegionsResponseBodyRegions) *DescribeSupportedRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetRequestId(v string) *DescribeSupportedRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetSuccess(v bool) *DescribeSupportedRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeSupportedRegionsResponseBodyRegions struct {
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeSupportedRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportedRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeSupportedRegionsResponseBodyRegions) SetRegion(v string) *DescribeSupportedRegionsResponseBodyRegions {
	s.Region = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBodyRegions) SetRegionName(v string) *DescribeSupportedRegionsResponseBodyRegions {
	s.RegionName = &v
	return s
}

type DescribeSupportedRegionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupportedRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupportedRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportedRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportedRegionsResponse) SetHeaders(v map[string]*string) *DescribeSupportedRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportedRegionsResponse) SetStatusCode(v int32) *DescribeSupportedRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportedRegionsResponse) SetBody(v *DescribeSupportedRegionsResponseBody) *DescribeSupportedRegionsResponse {
	s.Body = v
	return s
}

type DescribeSupportedZonesRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeSupportedZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportedZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupportedZonesRequest) SetArchitectureType(v string) *DescribeSupportedZonesRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeSupportedZonesRequest) SetRegion(v string) *DescribeSupportedZonesRequest {
	s.Region = &v
	return s
}

type DescribeSupportedZonesResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	ZoneIds   []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s DescribeSupportedZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportedZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportedZonesResponseBody) SetRequestId(v string) *DescribeSupportedZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetSuccess(v bool) *DescribeSupportedZonesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetZoneIds(v []*string) *DescribeSupportedZonesResponseBody {
	s.ZoneIds = v
	return s
}

type DescribeSupportedZonesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupportedZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupportedZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportedZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportedZonesResponse) SetHeaders(v map[string]*string) *DescribeSupportedZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportedZonesResponse) SetStatusCode(v int32) *DescribeSupportedZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportedZonesResponse) SetBody(v *DescribeSupportedZonesResponseBody) *DescribeSupportedZonesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
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
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// requestID。
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TagReponseId *string                                     `json:"TagReponseId,omitempty" xml:"TagReponseId,omitempty"`
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

func (s *ListTagResourcesResponseBody) SetSuccess(v bool) *ListTagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagReponseId(v string) *ListTagResourcesResponseBody {
	s.TagReponseId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
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

type ModifyPrepayInstanceSpecRequest struct {
	ModifyPrepayInstanceSpecRequest *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest `json:"ModifyPrepayInstanceSpecRequest,omitempty" xml:"ModifyPrepayInstanceSpecRequest,omitempty" type:"Struct"`
}

func (s ModifyPrepayInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequest) SetModifyPrepayInstanceSpecRequest(v *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) *ModifyPrepayInstanceSpecRequest {
	s.ModifyPrepayInstanceSpecRequest = v
	return s
}

type ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest struct {
	Ha             *bool                                                                         `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaResourceSpec *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	HaVSwitchIds   []*string                                                                     `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	HaZoneId       *string                                                                       `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	InstanceId     *string                                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region         *string                                                                       `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec   *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec   `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetHa(v bool) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.Ha = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetHaResourceSpec(v *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.HaResourceSpec = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetHaVSwitchIds(v []*string) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetHaZoneId(v string) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.HaZoneId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetInstanceId(v string) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetRegion(v string) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetResourceSpec(v *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.ResourceSpec = v
	return s
}

type ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) SetCpu(v int32) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) SetMemoryGB(v int32) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

type ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) SetCpu(v int32) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) SetMemoryGB(v int32) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type ModifyPrepayInstanceSpecResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPrepayInstanceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecResponseBody) SetOrderId(v int64) *ModifyPrepayInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecResponseBody) SetRequestId(v string) *ModifyPrepayInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecResponseBody) SetSuccess(v bool) *ModifyPrepayInstanceSpecResponseBody {
	s.Success = &v
	return s
}

type ModifyPrepayInstanceSpecResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPrepayInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPrepayInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyPrepayInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyPrepayInstanceSpecResponse) SetStatusCode(v int32) *ModifyPrepayInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPrepayInstanceSpecResponse) SetBody(v *ModifyPrepayInstanceSpecResponseBody) *ModifyPrepayInstanceSpecResponse {
	s.Body = v
	return s
}

type ModifyPrepayNamespaceSpecRequest struct {
	ModifyPrepayNamespaceSpecRequest *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest `json:"ModifyPrepayNamespaceSpecRequest,omitempty" xml:"ModifyPrepayNamespaceSpecRequest,omitempty" type:"Struct"`
}

func (s ModifyPrepayNamespaceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecRequest) SetModifyPrepayNamespaceSpecRequest(v *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) *ModifyPrepayNamespaceSpecRequest {
	s.ModifyPrepayNamespaceSpecRequest = v
	return s
}

type ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest struct {
	InstanceId   *string                                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Namespace    *string                                                                       `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Region       *string                                                                       `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) SetInstanceId(v string) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) SetNamespace(v string) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) SetRegion(v string) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) SetResourceSpec(v *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest {
	s.ResourceSpec = v
	return s
}

type ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) SetCpu(v int32) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) SetMemoryGB(v int32) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type ModifyPrepayNamespaceSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPrepayNamespaceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecResponseBody) SetRequestId(v string) *ModifyPrepayNamespaceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecResponseBody) SetSuccess(v bool) *ModifyPrepayNamespaceSpecResponseBody {
	s.Success = &v
	return s
}

type ModifyPrepayNamespaceSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPrepayNamespaceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPrepayNamespaceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecResponse) SetHeaders(v map[string]*string) *ModifyPrepayNamespaceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyPrepayNamespaceSpecResponse) SetStatusCode(v int32) *ModifyPrepayNamespaceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecResponse) SetBody(v *ModifyPrepayNamespaceSpecResponseBody) *ModifyPrepayNamespaceSpecResponse {
	s.Body = v
	return s
}

type QueryConvertInstancePriceRequest struct {
	ConvertPostpayInstanceRequest *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest `json:"ConvertPostpayInstanceRequest,omitempty" xml:"ConvertPostpayInstanceRequest,omitempty" type:"Struct"`
}

func (s QueryConvertInstancePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequest) SetConvertPostpayInstanceRequest(v *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) *QueryConvertInstancePriceRequest {
	s.ConvertPostpayInstanceRequest = v
	return s
}

type QueryConvertInstancePriceRequestConvertPostpayInstanceRequest struct {
	Duration               *int32                                                                                 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId             *string                                                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsAutoRenew            *bool                                                                                  `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	NamespaceResourceSpecs []*QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty" type:"Repeated"`
	PricingCycle           *string                                                                                `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Region                 *string                                                                                `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetDuration(v int32) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.Duration = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetInstanceId(v string) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetIsAutoRenew(v bool) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetNamespaceResourceSpecs(v []*QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.NamespaceResourceSpecs = v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetPricingCycle(v string) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetRegion(v string) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.Region = &v
	return s
}

type QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs struct {
	Namespace    *string                                                                                          `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ResourceSpec *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) SetNamespace(v string) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs {
	s.Namespace = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) SetResourceSpec(v *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs {
	s.ResourceSpec = v
	return s
}

type QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) SetCpu(v int32) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) SetMemoryGB(v int32) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.MemoryGB = &v
	return s
}

type QueryConvertInstancePriceResponseBody struct {
	PriceInfo *QueryConvertInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryConvertInstancePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponseBody) SetPriceInfo(v *QueryConvertInstancePriceResponseBodyPriceInfo) *QueryConvertInstancePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *QueryConvertInstancePriceResponseBody) SetRequestId(v string) *QueryConvertInstancePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBody) SetSuccess(v bool) *QueryConvertInstancePriceResponseBody {
	s.Success = &v
	return s
}

type QueryConvertInstancePriceResponseBodyPriceInfo struct {
	Code               *string                                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Currency           *string                                                             `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo     *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo       `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	DiscountAmount     *float32                                                            `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	IsContractActivity *bool                                                               `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	Message            *string                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	OriginalAmount     *float32                                                            `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryConvertInstancePriceResponseBodyPriceInfoRules              `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                                             `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                                             `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	TradeAmount        *float32                                                            `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryConvertInstancePriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetCode(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.Code = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetCurrency(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetDepreciateInfo(v *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.DepreciateInfo = v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetIsContractActivity(v bool) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.IsContractActivity = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v []*QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetOriginalAmount(v float32) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetRules(v []*QueryConvertInstancePriceResponseBodyPriceInfoRules) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetStandDiscountPrice(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetStandPrice(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.StandPrice = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

type QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo struct {
	CheapRate           *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	CheapStandAmount    *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	IsShow              *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	MonthPrice          *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapRate(v string) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapStandAmount(v string) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetIsShow(v bool) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetMonthPrice(v string) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetOriginalStandAmount(v string) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo) SetStartTime(v string) *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.StartTime = &v
	return s
}

type QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
	PromotionDesc     *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionName     *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	Selected          *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionDesc(v string) *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionName(v string) *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionOptionNo(v string) *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) SetSelected(v bool) *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.Selected = &v
	return s
}

type QueryConvertInstancePriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoRules) SetDescription(v string) *QueryConvertInstancePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *QueryConvertInstancePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type QueryConvertInstancePriceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConvertInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConvertInstancePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceResponse) SetHeaders(v map[string]*string) *QueryConvertInstancePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryConvertInstancePriceResponse) SetStatusCode(v int32) *QueryConvertInstancePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConvertInstancePriceResponse) SetBody(v *QueryConvertInstancePriceResponseBody) *QueryConvertInstancePriceResponse {
	s.Body = v
	return s
}

type QueryConvertPrepayInstancePriceRequest struct {
	ConvertPrepayInstanceRequest *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest `json:"ConvertPrepayInstanceRequest,omitempty" xml:"ConvertPrepayInstanceRequest,omitempty" type:"Struct"`
}

func (s QueryConvertPrepayInstancePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceRequest) SetConvertPrepayInstanceRequest(v *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) *QueryConvertPrepayInstancePriceRequest {
	s.ConvertPrepayInstanceRequest = v
	return s
}

type QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) SetInstanceId(v string) *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest) SetRegion(v string) *QueryConvertPrepayInstancePriceRequestConvertPrepayInstanceRequest {
	s.Region = &v
	return s
}

type QueryConvertPrepayInstancePriceResponseBody struct {
	PriceInfo *QueryConvertPrepayInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponseBody) SetPriceInfo(v *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) *QueryConvertPrepayInstancePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBody) SetRequestId(v string) *QueryConvertPrepayInstancePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBody) SetSuccess(v bool) *QueryConvertPrepayInstancePriceResponseBody {
	s.Success = &v
	return s
}

type QueryConvertPrepayInstancePriceResponseBodyPriceInfo struct {
	Code               *string                                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Currency           *string                                                                   `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo     *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo       `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	DiscountAmount     *float32                                                                  `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	IsContractActivity *bool                                                                     `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	Message            *string                                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	OriginalAmount     *float32                                                                  `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules              `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                                                   `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                                                   `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	TradeAmount        *float32                                                                  `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetCode(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.Code = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetCurrency(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetDepreciateInfo(v *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.DepreciateInfo = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetIsContractActivity(v bool) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.IsContractActivity = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetOriginalAmount(v float32) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetRules(v []*QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetStandDiscountPrice(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetStandPrice(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.StandPrice = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryConvertPrepayInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

type QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo struct {
	CheapRate           *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	CheapStandAmount    *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	IsShow              *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	MonthPrice          *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapRate(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapStandAmount(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetIsShow(v bool) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetMonthPrice(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetOriginalStandAmount(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo) SetStartTime(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.StartTime = &v
	return s
}

type QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
	PromotionDesc     *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionName     *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	Selected          *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionDesc(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionName(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionOptionNo(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions) SetSelected(v bool) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.Selected = &v
	return s
}

type QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) SetDescription(v string) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *QueryConvertPrepayInstancePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type QueryConvertPrepayInstancePriceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConvertPrepayInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConvertPrepayInstancePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertPrepayInstancePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryConvertPrepayInstancePriceResponse) SetHeaders(v map[string]*string) *QueryConvertPrepayInstancePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponse) SetStatusCode(v int32) *QueryConvertPrepayInstancePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConvertPrepayInstancePriceResponse) SetBody(v *QueryConvertPrepayInstancePriceResponseBody) *QueryConvertPrepayInstancePriceResponse {
	s.Body = v
	return s
}

type QueryCreateInstancePriceRequest struct {
	CreateInstanceRequest *QueryCreateInstancePriceRequestCreateInstanceRequest `json:"CreateInstanceRequest,omitempty" xml:"CreateInstanceRequest,omitempty" type:"Struct"`
}

func (s QueryCreateInstancePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequest) SetCreateInstanceRequest(v *QueryCreateInstancePriceRequestCreateInstanceRequest) *QueryCreateInstancePriceRequest {
	s.CreateInstanceRequest = v
	return s
}

type QueryCreateInstancePriceRequestCreateInstanceRequest struct {
	ArchitectureType *string                                                             `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	AutoRenew        *bool                                                               `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType       *string                                                             `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Duration         *int32                                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extra            *string                                                             `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Ha               *bool                                                               `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaResourceSpec   *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	InstanceName     *string                                                             `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PricingCycle     *string                                                             `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	PromotionCode    *string                                                             `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	Region           *string                                                             `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec     *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec   `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage          *QueryCreateInstancePriceRequestCreateInstanceRequestStorage        `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	UsePromotionCode *bool                                                               `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	VSwitchIds       []*string                                                           `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// VPC ID。
	VpcId  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetArchitectureType(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.ArchitectureType = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetAutoRenew(v bool) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetChargeType(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetDuration(v int32) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetExtra(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.Extra = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetHa(v bool) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.Ha = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetHaResourceSpec(v *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.HaResourceSpec = v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetInstanceName(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetPricingCycle(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetPromotionCode(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetRegion(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.Region = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetResourceSpec(v *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.ResourceSpec = v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetStorage(v *QueryCreateInstancePriceRequestCreateInstanceRequestStorage) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.Storage = v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetUsePromotionCode(v bool) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetVSwitchIds(v []*string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.VSwitchIds = v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetVpcId(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetZoneId(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.ZoneId = &v
	return s
}

type QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) SetCpu(v int32) *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) SetMemoryGB(v int32) *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

type QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) SetCpu(v int32) *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) SetMemoryGB(v int32) *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type QueryCreateInstancePriceRequestCreateInstanceRequestStorage struct {
	Oss *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestStorage) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestStorage) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestStorage) SetOss(v *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss) *QueryCreateInstancePriceRequestCreateInstanceRequestStorage {
	s.Oss = v
	return s
}

type QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss struct {
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss) SetBucket(v string) *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss {
	s.Bucket = &v
	return s
}

type QueryCreateInstancePriceResponseBody struct {
	PriceInfo *QueryCreateInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCreateInstancePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBody) SetPriceInfo(v *QueryCreateInstancePriceResponseBodyPriceInfo) *QueryCreateInstancePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *QueryCreateInstancePriceResponseBody) SetRequestId(v string) *QueryCreateInstancePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBody) SetSuccess(v bool) *QueryCreateInstancePriceResponseBody {
	s.Success = &v
	return s
}

type QueryCreateInstancePriceResponseBodyPriceInfo struct {
	Code               *string                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Currency           *string                                                            `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo     *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo       `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	DiscountAmount     *float32                                                           `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	IsContractActivity *bool                                                              `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	Message            *string                                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	OriginalAmount     *float32                                                           `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryCreateInstancePriceResponseBodyPriceInfoRules              `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                                            `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                                            `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	TradeAmount        *float32                                                           `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryCreateInstancePriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetCode(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.Code = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetCurrency(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetDepreciateInfo(v *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.DepreciateInfo = v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetIsContractActivity(v bool) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.IsContractActivity = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v []*QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetOriginalAmount(v float32) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetRules(v []*QueryCreateInstancePriceResponseBodyPriceInfoRules) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetStandDiscountPrice(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetStandPrice(v string) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.StandPrice = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryCreateInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

type QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo struct {
	CheapRate           *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	CheapStandAmount    *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	IsShow              *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	MonthPrice          *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapRate(v string) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapStandAmount(v string) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetIsShow(v bool) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetMonthPrice(v string) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetOriginalStandAmount(v string) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo) SetStartTime(v string) *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.StartTime = &v
	return s
}

type QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
	PromotionDesc     *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionName     *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	Selected          *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionDesc(v string) *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionName(v string) *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionOptionNo(v string) *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions) SetSelected(v bool) *QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.Selected = &v
	return s
}

type QueryCreateInstancePriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoRules) SetDescription(v string) *QueryCreateInstancePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *QueryCreateInstancePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *QueryCreateInstancePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type QueryCreateInstancePriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCreateInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCreateInstancePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceResponse) SetHeaders(v map[string]*string) *QueryCreateInstancePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryCreateInstancePriceResponse) SetStatusCode(v int32) *QueryCreateInstancePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCreateInstancePriceResponse) SetBody(v *QueryCreateInstancePriceResponseBody) *QueryCreateInstancePriceResponse {
	s.Body = v
	return s
}

type QueryModifyInstancePriceRequest struct {
	ModifyPrepayInstanceSpecRequest *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest `json:"ModifyPrepayInstanceSpecRequest,omitempty" xml:"ModifyPrepayInstanceSpecRequest,omitempty" type:"Struct"`
}

func (s QueryModifyInstancePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequest) SetModifyPrepayInstanceSpecRequest(v *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) *QueryModifyInstancePriceRequest {
	s.ModifyPrepayInstanceSpecRequest = v
	return s
}

type QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest struct {
	Ha             *bool                                                                         `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaResourceSpec *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	HaVSwitchIds   []*string                                                                     `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	HaZoneId       *string                                                                       `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	InstanceId     *string                                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region         *string                                                                       `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec   *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec   `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetHa(v bool) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.Ha = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetHaResourceSpec(v *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.HaResourceSpec = v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetHaVSwitchIds(v []*string) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetHaZoneId(v string) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.HaZoneId = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetInstanceId(v string) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetRegion(v string) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.Region = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetResourceSpec(v *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.ResourceSpec = v
	return s
}

type QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) SetCpu(v int32) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) SetMemoryGB(v int32) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

type QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) SetCpu(v int32) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) SetMemoryGB(v int32) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type QueryModifyInstancePriceResponseBody struct {
	PriceInfo *QueryModifyInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryModifyInstancePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponseBody) SetPriceInfo(v *QueryModifyInstancePriceResponseBodyPriceInfo) *QueryModifyInstancePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *QueryModifyInstancePriceResponseBody) SetRequestId(v string) *QueryModifyInstancePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBody) SetSuccess(v bool) *QueryModifyInstancePriceResponseBody {
	s.Success = &v
	return s
}

type QueryModifyInstancePriceResponseBodyPriceInfo struct {
	Code               *string                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Currency           *string                                                            `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo     *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo       `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	DiscountAmount     *float32                                                           `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	IsContractActivity *bool                                                              `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	Message            *string                                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	OriginalAmount     *float32                                                           `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryModifyInstancePriceResponseBodyPriceInfoRules              `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                                            `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                                            `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	TradeAmount        *float32                                                           `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryModifyInstancePriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetCode(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.Code = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetCurrency(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetDepreciateInfo(v *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.DepreciateInfo = v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetIsContractActivity(v bool) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.IsContractActivity = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v []*QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetOriginalAmount(v float32) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetRules(v []*QueryModifyInstancePriceResponseBodyPriceInfoRules) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetStandDiscountPrice(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetStandPrice(v string) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.StandPrice = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryModifyInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

type QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo struct {
	CheapRate           *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	CheapStandAmount    *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	IsShow              *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	MonthPrice          *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapRate(v string) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapStandAmount(v string) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetIsShow(v bool) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetMonthPrice(v string) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetOriginalStandAmount(v string) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo) SetStartTime(v string) *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.StartTime = &v
	return s
}

type QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
	PromotionDesc     *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionName     *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	Selected          *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionDesc(v string) *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionName(v string) *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionOptionNo(v string) *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions) SetSelected(v bool) *QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.Selected = &v
	return s
}

type QueryModifyInstancePriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoRules) SetDescription(v string) *QueryModifyInstancePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *QueryModifyInstancePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *QueryModifyInstancePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type QueryModifyInstancePriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryModifyInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryModifyInstancePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceResponse) SetHeaders(v map[string]*string) *QueryModifyInstancePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryModifyInstancePriceResponse) SetStatusCode(v int32) *QueryModifyInstancePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryModifyInstancePriceResponse) SetBody(v *QueryModifyInstancePriceResponseBody) *QueryModifyInstancePriceResponse {
	s.Body = v
	return s
}

type QueryRenewInstancePriceRequest struct {
	RenewInstanceRequest *QueryRenewInstancePriceRequestRenewInstanceRequest `json:"RenewInstanceRequest,omitempty" xml:"RenewInstanceRequest,omitempty" type:"Struct"`
}

func (s QueryRenewInstancePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceRequest) SetRenewInstanceRequest(v *QueryRenewInstancePriceRequestRenewInstanceRequest) *QueryRenewInstancePriceRequest {
	s.RenewInstanceRequest = v
	return s
}

type QueryRenewInstancePriceRequestRenewInstanceRequest struct {
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryRenewInstancePriceRequestRenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewInstancePriceRequestRenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) SetDuration(v int32) *QueryRenewInstancePriceRequestRenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) SetInstanceId(v string) *QueryRenewInstancePriceRequestRenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) SetPricingCycle(v string) *QueryRenewInstancePriceRequestRenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryRenewInstancePriceRequestRenewInstanceRequest) SetRegion(v string) *QueryRenewInstancePriceRequestRenewInstanceRequest {
	s.Region = &v
	return s
}

type QueryRenewInstancePriceResponseBody struct {
	PriceInfo *QueryRenewInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRenewInstancePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewInstancePriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponseBody) SetPriceInfo(v *QueryRenewInstancePriceResponseBodyPriceInfo) *QueryRenewInstancePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *QueryRenewInstancePriceResponseBody) SetRequestId(v string) *QueryRenewInstancePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBody) SetSuccess(v bool) *QueryRenewInstancePriceResponseBody {
	s.Success = &v
	return s
}

type QueryRenewInstancePriceResponseBodyPriceInfo struct {
	Code               *string                                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Currency           *string                                                           `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo     *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo       `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	DiscountAmount     *float32                                                          `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	IsContractActivity *bool                                                             `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	Message            *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	OriginalAmount     *float32                                                          `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryRenewInstancePriceResponseBodyPriceInfoRules              `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                                           `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                                           `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	TradeAmount        *float32                                                          `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s QueryRenewInstancePriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewInstancePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetCode(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.Code = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetCurrency(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetDepreciateInfo(v *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.DepreciateInfo = v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetIsContractActivity(v bool) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.IsContractActivity = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v []*QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetOriginalAmount(v float32) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetRules(v []*QueryRenewInstancePriceResponseBodyPriceInfoRules) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetStandDiscountPrice(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.StandDiscountPrice = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetStandPrice(v string) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.StandPrice = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryRenewInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

type QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo struct {
	CheapRate           *string `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	CheapStandAmount    *string `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	IsShow              *bool   `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	MonthPrice          *string `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	OriginalStandAmount *string `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapRate(v string) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetCheapStandAmount(v string) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetIsShow(v bool) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetMonthPrice(v string) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetOriginalStandAmount(v string) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo) SetStartTime(v string) *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo {
	s.StartTime = &v
	return s
}

type QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
	PromotionDesc     *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionName     *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	Selected          *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionDesc(v string) *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionName(v string) *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) SetPromotionOptionNo(v string) *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions) SetSelected(v bool) *QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions {
	s.Selected = &v
	return s
}

type QueryRenewInstancePriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewInstancePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoRules) SetDescription(v string) *QueryRenewInstancePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *QueryRenewInstancePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *QueryRenewInstancePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type QueryRenewInstancePriceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRenewInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRenewInstancePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewInstancePriceResponse) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceResponse) SetHeaders(v map[string]*string) *QueryRenewInstancePriceResponse {
	s.Headers = v
	return s
}

func (s *QueryRenewInstancePriceResponse) SetStatusCode(v int32) *QueryRenewInstancePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRenewInstancePriceResponse) SetBody(v *QueryRenewInstancePriceResponseBody) *QueryRenewInstancePriceResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	RenewInstanceRequest *RenewInstanceRequestRenewInstanceRequest `json:"RenewInstanceRequest,omitempty" xml:"RenewInstanceRequest,omitempty" type:"Struct"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetRenewInstanceRequest(v *RenewInstanceRequestRenewInstanceRequest) *RenewInstanceRequest {
	s.RenewInstanceRequest = v
	return s
}

type RenewInstanceRequestRenewInstanceRequest struct {
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s RenewInstanceRequestRenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequestRenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequestRenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequestRenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewInstanceRequestRenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequestRenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequestRenewInstanceRequest) SetPricingCycle(v string) *RenewInstanceRequestRenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewInstanceRequestRenewInstanceRequest) SetRegion(v string) *RenewInstanceRequestRenewInstanceRequest {
	s.Region = &v
	return s
}

type RenewInstanceResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetOrderId(v int64) *RenewInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetSuccess(v bool) *RenewInstanceResponseBody {
	s.Success = &v
	return s
}

type RenewInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetStatusCode(v int32) *RenewInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
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
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TagResponseId *string `json:"TagResponseId,omitempty" xml:"TagResponseId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetCode(v string) *TagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBody) SetMessage(v string) *TagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *TagResourcesResponseBody) SetTagResponseId(v string) *TagResourcesResponseBody {
	s.TagResponseId = &v
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
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
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
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TagResponseId *string `json:"TagResponseId,omitempty" xml:"TagResponseId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetCode(v string) *UntagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBody) SetMessage(v string) *UntagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) SetTagResponseId(v string) *UntagResourcesResponseBody {
	s.TagResponseId = &v
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("foasconsole"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ConvertInstanceWithOptions(request *ConvertInstanceRequest, runtime *util.RuntimeOptions) (_result *ConvertInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConvertPostpayInstanceRequest)) {
		bodyFlat["ConvertPostpayInstanceRequest"] = request.ConvertPostpayInstanceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConvertInstance"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertInstance(request *ConvertInstanceRequest) (_result *ConvertInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.ConvertInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConvertPrepayInstanceWithOptions(request *ConvertPrepayInstanceRequest, runtime *util.RuntimeOptions) (_result *ConvertPrepayInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConvertPrepayInstanceRequest)) {
		bodyFlat["ConvertPrepayInstanceRequest"] = request.ConvertPrepayInstanceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConvertPrepayInstance"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConvertPrepayInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertPrepayInstance(request *ConvertPrepayInstanceRequest) (_result *ConvertPrepayInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertPrepayInstanceResponse{}
	_body, _err := client.ConvertPrepayInstanceWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateInstanceRequest)) {
		bodyFlat["CreateInstanceRequest"] = request.CreateInstanceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2019-06-01"),
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

func (client *Client) CreateNamespaceWithOptions(request *CreateNamespaceRequest, runtime *util.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateNamespaceRequest)) {
		bodyFlat["CreateNamespaceRequest"] = request.CreateNamespaceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNamespace"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteInstanceRequest)) {
		bodyFlat["DeleteInstanceRequest"] = request.DeleteInstanceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *util.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteNamespaceRequest)) {
		bodyFlat["DeleteNamespaceRequest"] = request.DeleteNamespaceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNamespace"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNamespacesWithOptions(request *DescribeNamespacesRequest, runtime *util.RuntimeOptions) (_result *DescribeNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespaces"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (_result *DescribeNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNamespacesResponse{}
	_body, _err := client.DescribeNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSupportedRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeSupportedRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeSupportedRegions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSupportedRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSupportedRegions() (_result *DescribeSupportedRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSupportedRegionsResponse{}
	_body, _err := client.DescribeSupportedRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSupportedZonesWithOptions(request *DescribeSupportedZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeSupportedZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSupportedZones"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSupportedZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSupportedZones(request *DescribeSupportedZonesRequest) (_result *DescribeSupportedZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSupportedZonesResponse{}
	_body, _err := client.DescribeSupportedZonesWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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
		Version:     tea.String("2019-06-01"),
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

/**
 * @deprecated : ModifyPrepayInstanceSpec is deprecated, please use foasconsole::2019-06-01::ModifyInstanceSpec instead.
 *
 * @param request ModifyPrepayInstanceSpecRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyPrepayInstanceSpecResponse
 */
// Deprecated
func (client *Client) ModifyPrepayInstanceSpecWithOptions(request *ModifyPrepayInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyPrepayInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModifyPrepayInstanceSpecRequest)) {
		bodyFlat["ModifyPrepayInstanceSpecRequest"] = request.ModifyPrepayInstanceSpecRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPrepayInstanceSpec"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPrepayInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : ModifyPrepayInstanceSpec is deprecated, please use foasconsole::2019-06-01::ModifyInstanceSpec instead.
 *
 * @param request ModifyPrepayInstanceSpecRequest
 * @return ModifyPrepayInstanceSpecResponse
 */
// Deprecated
func (client *Client) ModifyPrepayInstanceSpec(request *ModifyPrepayInstanceSpecRequest) (_result *ModifyPrepayInstanceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPrepayInstanceSpecResponse{}
	_body, _err := client.ModifyPrepayInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated : ModifyPrepayNamespaceSpec is deprecated, please use foasconsole::2019-06-01::ModifyNamespaceSpec instead.
 *
 * @param request ModifyPrepayNamespaceSpecRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyPrepayNamespaceSpecResponse
 */
// Deprecated
func (client *Client) ModifyPrepayNamespaceSpecWithOptions(request *ModifyPrepayNamespaceSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyPrepayNamespaceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModifyPrepayNamespaceSpecRequest)) {
		bodyFlat["ModifyPrepayNamespaceSpecRequest"] = request.ModifyPrepayNamespaceSpecRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPrepayNamespaceSpec"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPrepayNamespaceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : ModifyPrepayNamespaceSpec is deprecated, please use foasconsole::2019-06-01::ModifyNamespaceSpec instead.
 *
 * @param request ModifyPrepayNamespaceSpecRequest
 * @return ModifyPrepayNamespaceSpecResponse
 */
// Deprecated
func (client *Client) ModifyPrepayNamespaceSpec(request *ModifyPrepayNamespaceSpecRequest) (_result *ModifyPrepayNamespaceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPrepayNamespaceSpecResponse{}
	_body, _err := client.ModifyPrepayNamespaceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryConvertInstancePriceWithOptions(request *QueryConvertInstancePriceRequest, runtime *util.RuntimeOptions) (_result *QueryConvertInstancePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConvertPostpayInstanceRequest)) {
		bodyFlat["ConvertPostpayInstanceRequest"] = request.ConvertPostpayInstanceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryConvertInstancePrice"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryConvertInstancePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryConvertInstancePrice(request *QueryConvertInstancePriceRequest) (_result *QueryConvertInstancePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryConvertInstancePriceResponse{}
	_body, _err := client.QueryConvertInstancePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryConvertPrepayInstancePriceWithOptions(request *QueryConvertPrepayInstancePriceRequest, runtime *util.RuntimeOptions) (_result *QueryConvertPrepayInstancePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConvertPrepayInstanceRequest)) {
		bodyFlat["ConvertPrepayInstanceRequest"] = request.ConvertPrepayInstanceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryConvertPrepayInstancePrice"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryConvertPrepayInstancePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryConvertPrepayInstancePrice(request *QueryConvertPrepayInstancePriceRequest) (_result *QueryConvertPrepayInstancePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryConvertPrepayInstancePriceResponse{}
	_body, _err := client.QueryConvertPrepayInstancePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCreateInstancePriceWithOptions(request *QueryCreateInstancePriceRequest, runtime *util.RuntimeOptions) (_result *QueryCreateInstancePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateInstanceRequest)) {
		bodyFlat["CreateInstanceRequest"] = request.CreateInstanceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCreateInstancePrice"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCreateInstancePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCreateInstancePrice(request *QueryCreateInstancePriceRequest) (_result *QueryCreateInstancePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCreateInstancePriceResponse{}
	_body, _err := client.QueryCreateInstancePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryModifyInstancePriceWithOptions(request *QueryModifyInstancePriceRequest, runtime *util.RuntimeOptions) (_result *QueryModifyInstancePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModifyPrepayInstanceSpecRequest)) {
		bodyFlat["ModifyPrepayInstanceSpecRequest"] = request.ModifyPrepayInstanceSpecRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryModifyInstancePrice"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryModifyInstancePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryModifyInstancePrice(request *QueryModifyInstancePriceRequest) (_result *QueryModifyInstancePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryModifyInstancePriceResponse{}
	_body, _err := client.QueryModifyInstancePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRenewInstancePriceWithOptions(request *QueryRenewInstancePriceRequest, runtime *util.RuntimeOptions) (_result *QueryRenewInstancePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RenewInstanceRequest)) {
		bodyFlat["RenewInstanceRequest"] = request.RenewInstanceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRenewInstancePrice"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRenewInstancePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRenewInstancePrice(request *QueryRenewInstancePriceRequest) (_result *QueryRenewInstancePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRenewInstancePriceResponse{}
	_body, _err := client.QueryRenewInstancePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RenewInstanceRequest)) {
		bodyFlat["RenewInstanceRequest"] = request.RenewInstanceRequest
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewInstance"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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
		Version:     tea.String("2019-06-01"),
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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
		Version:     tea.String("2019-06-01"),
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

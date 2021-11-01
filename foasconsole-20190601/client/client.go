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
	// 订购周期数量
	Duration   *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 是否自动续费
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// 项目空间资源规格。
	NamespaceResourceSpecs []*ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty" type:"Repeated"`
	// 订购周期
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
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
	// namespace名称，
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 资源规格。
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
	// cpu数量。
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// 内存大小。
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
	// 订单id
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConvertInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConvertInstanceResponse) SetBody(v *ConvertInstanceResponseBody) *ConvertInstanceResponse {
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
	AutoRenew    *bool                                                   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType   *string                                                 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Duration     *int32                                                  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceName *string                                                 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PricingCycle *string                                                 `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Region       *string                                                 `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec *CreateInstanceRequestCreateInstanceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage      *CreateInstanceRequestCreateInstanceRequestStorage      `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	VSwitchIds   []*string                                               `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VpcId        *string                                                 `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId       *string                                                 `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequestCreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequest) GoString() string {
	return s.String()
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

func (s *CreateInstanceRequestCreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetPricingCycle(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetRegion(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.Region = &v
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
	// 订单信息
	OrderInfo *CreateInstanceResponseBodyOrderInfo `json:"OrderInfo,omitempty" xml:"OrderInfo,omitempty" type:"Struct"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 订单id
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 付款类型
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageIndex  *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeInstancesRequestDescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestDescribeInstancesRequest) GoString() string {
	return s.String()
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
	ChargeType         *string                                             `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterStatus      *string                                             `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	InstanceId         *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName       *string                                             `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OrderState         *string                                             `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	Region             *string                                             `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCreateTime *int64                                              `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	ResourceId         *string                                             `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceSpec       *DescribeInstancesResponseBodyInstancesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage            *DescribeInstancesResponseBodyInstancesStorage      `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	Uid                *string                                             `json:"Uid,omitempty" xml:"Uid,omitempty"`
	VSwitchIds         []*string                                           `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VpcId              *string                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId             *string                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetChargeType(v string) *DescribeInstancesResponseBodyInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.ClusterStatus = &v
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

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 当前页数
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// 每页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// regionId
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeNamespacesRequestDescribeNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesRequestDescribeNamespacesRequest) GoString() string {
	return s.String()
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

type DescribeNamespacesResponseBody struct {
	Namespaces []*DescribeNamespacesResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	PageIndex  *int32                                      `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success    *bool  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	Namespace    *string                                               `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ResourceSpec *DescribeNamespacesResponseBodyNamespacesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	ResourceUsed *DescribeNamespacesResponseBodyNamespacesResourceUsed `json:"ResourceUsed,omitempty" xml:"ResourceUsed,omitempty" type:"Struct"`
	Status       *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) SetMemoryGB(v float32) *DescribeNamespacesResponseBodyNamespacesResourceUsed {
	s.MemoryGB = &v
	return s
}

type DescribeNamespacesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeNamespacesResponse) SetBody(v *DescribeNamespacesResponseBody) *DescribeNamespacesResponse {
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
	InstanceId   *string                                                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region       *string                                                                     `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
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
	// 订单id
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPrepayInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPrepayNamespaceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 订购周期数量
	Duration   *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 是否自动续费
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// 项目空间资源规格。
	NamespaceResourceSpecs []*QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty" type:"Repeated"`
	// 订购周期
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
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
	// namespace名称，
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 资源规格。
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
	// cpu数量。
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// 内存大小。
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
	// 价格信息，包括价格和优惠规则。
	PriceInfo *QueryConvertInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 货币单位。
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// 折扣
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 可选择的优惠券
	OptionalPromotions *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Struct"`
	// 原价
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// 活动规则。
	Rules []*QueryConvertInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// 最终价，为原价减去折扣。
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
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

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetDiscountAmount(v float32) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetMessage(v string) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.Message = &v
	return s
}

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetOptionalPromotions(v *QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions) *QueryConvertInstancePriceResponseBodyPriceInfo {
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

func (s *QueryConvertInstancePriceResponseBodyPriceInfo) SetTradeAmount(v float32) *QueryConvertInstancePriceResponseBodyPriceInfo {
	s.TradeAmount = &v
	return s
}

type QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions struct {
	// 优惠券描述
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// 优惠券名称
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// 优惠券编号
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
	// 活动规则描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 活动ID。
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryConvertInstancePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryConvertInstancePriceResponse) SetBody(v *QueryConvertInstancePriceResponseBody) *QueryConvertInstancePriceResponse {
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
	// 订购周期数量
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 订购周期
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// 地域id
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
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
	// orderId
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConvertInstance"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateNamespace"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteInstance"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNamespace"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &DescribeNamespacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNamespaces"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyPrepayInstanceSpecWithOptions(request *ModifyPrepayInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyPrepayInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPrepayInstanceSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPrepayInstanceSpec"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ModifyPrepayNamespaceSpecWithOptions(request *ModifyPrepayNamespaceSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyPrepayNamespaceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPrepayNamespaceSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPrepayNamespaceSpec"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryConvertInstancePriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryConvertInstancePrice"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewInstance"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConvertInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// This parameter is required.
	NamespaceResourceSpecs []*ConvertInstanceRequestNamespaceResourceSpecs `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ConvertInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequest) SetDuration(v int32) *ConvertInstanceRequest {
	s.Duration = &v
	return s
}

func (s *ConvertInstanceRequest) SetInstanceId(v string) *ConvertInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertInstanceRequest) SetIsAutoRenew(v bool) *ConvertInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *ConvertInstanceRequest) SetNamespaceResourceSpecs(v []*ConvertInstanceRequestNamespaceResourceSpecs) *ConvertInstanceRequest {
	s.NamespaceResourceSpecs = v
	return s
}

func (s *ConvertInstanceRequest) SetPricingCycle(v string) *ConvertInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConvertInstanceRequest) SetRegion(v string) *ConvertInstanceRequest {
	s.Region = &v
	return s
}

type ConvertInstanceRequestNamespaceResourceSpecs struct {
	// This parameter is required.
	//
	// example:
	//
	// ns-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	ResourceSpec *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ConvertInstanceRequestNamespaceResourceSpecs) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceRequestNamespaceResourceSpecs) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequestNamespaceResourceSpecs) SetNamespace(v string) *ConvertInstanceRequestNamespaceResourceSpecs {
	s.Namespace = &v
	return s
}

func (s *ConvertInstanceRequestNamespaceResourceSpecs) SetResourceSpec(v *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) *ConvertInstanceRequestNamespaceResourceSpecs {
	s.ResourceSpec = v
	return s
}

type ConvertInstanceRequestNamespaceResourceSpecsResourceSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) SetCpu(v int32) *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) SetMemoryGB(v int32) *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.MemoryGB = &v
	return s
}

type ConvertInstanceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// This parameter is required.
	NamespaceResourceSpecsShrink *string `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ConvertInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceShrinkRequest) SetDuration(v int32) *ConvertInstanceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetInstanceId(v string) *ConvertInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetIsAutoRenew(v bool) *ConvertInstanceShrinkRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetNamespaceResourceSpecsShrink(v string) *ConvertInstanceShrinkRequest {
	s.NamespaceResourceSpecsShrink = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetPricingCycle(v string) *ConvertInstanceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetRegion(v string) *ConvertInstanceShrinkRequest {
	s.Region = &v
	return s
}

type ConvertInstanceResponseBody struct {
	// example:
	//
	// 211473228320700
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
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

type CreateInstanceRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 1
	Duration *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extra    *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *CreateInstanceRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	// if can be null:
	// true
	HaVSwitchIds []*string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	// if can be null:
	// true
	HaZoneId *string `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtc-e2e-test-pre
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MonitorType  *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 500043499350689
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region          *string                            `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceSpec    *CreateInstanceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	// This parameter is required.
	Storage          *CreateInstanceRequestStorage `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	Tag              []*CreateInstanceRequestTag   `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UsePromotionCode *bool                         `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	// This parameter is required.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-2ze9xoh8qyt1rnxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetArchitectureType(v string) *CreateInstanceRequest {
	s.ArchitectureType = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetDuration(v int32) *CreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequest) SetExtra(v string) *CreateInstanceRequest {
	s.Extra = &v
	return s
}

func (s *CreateInstanceRequest) SetHa(v bool) *CreateInstanceRequest {
	s.Ha = &v
	return s
}

func (s *CreateInstanceRequest) SetHaResourceSpec(v *CreateInstanceRequestHaResourceSpec) *CreateInstanceRequest {
	s.HaResourceSpec = v
	return s
}

func (s *CreateInstanceRequest) SetHaVSwitchIds(v []*string) *CreateInstanceRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *CreateInstanceRequest) SetHaZoneId(v string) *CreateInstanceRequest {
	s.HaZoneId = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetMonitorType(v string) *CreateInstanceRequest {
	s.MonitorType = &v
	return s
}

func (s *CreateInstanceRequest) SetPricingCycle(v string) *CreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetPromotionCode(v string) *CreateInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateInstanceRequest) SetRegion(v string) *CreateInstanceRequest {
	s.Region = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceSpec(v *CreateInstanceRequestResourceSpec) *CreateInstanceRequest {
	s.ResourceSpec = v
	return s
}

func (s *CreateInstanceRequest) SetStorage(v *CreateInstanceRequestStorage) *CreateInstanceRequest {
	s.Storage = v
	return s
}

func (s *CreateInstanceRequest) SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateInstanceRequest) SetUsePromotionCode(v bool) *CreateInstanceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchIds(v []*string) *CreateInstanceRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateInstanceRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateInstanceRequestHaResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestHaResourceSpec) SetCpu(v int32) *CreateInstanceRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateInstanceRequestHaResourceSpec) SetMemoryGB(v int32) *CreateInstanceRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

type CreateInstanceRequestResourceSpec struct {
	// example:
	//
	// 30
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 120
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateInstanceRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestResourceSpec) SetCpu(v int32) *CreateInstanceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateInstanceRequestResourceSpec) SetMemoryGB(v int32) *CreateInstanceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type CreateInstanceRequestStorage struct {
	FullyManaged *bool                            `json:"FullyManaged,omitempty" xml:"FullyManaged,omitempty"`
	Oss          *CreateInstanceRequestStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s CreateInstanceRequestStorage) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestStorage) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestStorage) SetFullyManaged(v bool) *CreateInstanceRequestStorage {
	s.FullyManaged = &v
	return s
}

func (s *CreateInstanceRequestStorage) SetOss(v *CreateInstanceRequestStorageOss) *CreateInstanceRequestStorage {
	s.Oss = v
	return s
}

type CreateInstanceRequestStorageOss struct {
	// example:
	//
	// oss-flink-cn-shanghai-260343971602724445
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s CreateInstanceRequestStorageOss) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestStorageOss) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestStorageOss) SetBucket(v string) *CreateInstanceRequestStorageOss {
	s.Bucket = &v
	return s
}

type CreateInstanceRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTag) SetKey(v string) *CreateInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTag) SetValue(v string) *CreateInstanceRequestTag {
	s.Value = &v
	return s
}

type CreateInstanceShrinkRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 1
	Duration *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extra    *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpecShrink *string `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty"`
	// if can be null:
	// true
	HaVSwitchIdsShrink *string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty"`
	// if can be null:
	// true
	HaZoneId *string `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtc-e2e-test-pre
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MonitorType  *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 500043499350689
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId    *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
	// This parameter is required.
	StorageShrink    *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	TagShrink        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	UsePromotionCode *bool   `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	// This parameter is required.
	VSwitchIdsShrink *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-2ze9xoh8qyt1rnxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceShrinkRequest) SetArchitectureType(v string) *CreateInstanceShrinkRequest {
	s.ArchitectureType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetAutoRenew(v bool) *CreateInstanceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetChargeType(v string) *CreateInstanceShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetDuration(v int32) *CreateInstanceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetExtra(v string) *CreateInstanceShrinkRequest {
	s.Extra = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetHa(v bool) *CreateInstanceShrinkRequest {
	s.Ha = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetHaResourceSpecShrink(v string) *CreateInstanceShrinkRequest {
	s.HaResourceSpecShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetHaVSwitchIdsShrink(v string) *CreateInstanceShrinkRequest {
	s.HaVSwitchIdsShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetHaZoneId(v string) *CreateInstanceShrinkRequest {
	s.HaZoneId = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetInstanceName(v string) *CreateInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetMonitorType(v string) *CreateInstanceShrinkRequest {
	s.MonitorType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetPricingCycle(v string) *CreateInstanceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetPromotionCode(v string) *CreateInstanceShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetRegion(v string) *CreateInstanceShrinkRequest {
	s.Region = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetResourceGroupId(v string) *CreateInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetResourceSpecShrink(v string) *CreateInstanceShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetStorageShrink(v string) *CreateInstanceShrinkRequest {
	s.StorageShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetTagShrink(v string) *CreateInstanceShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetUsePromotionCode(v bool) *CreateInstanceShrinkRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetVSwitchIdsShrink(v string) *CreateInstanceShrinkRequest {
	s.VSwitchIdsShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetVpcId(v string) *CreateInstanceShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetZoneId(v string) *CreateInstanceShrinkRequest {
	s.ZoneId = &v
	return s
}

type CreateInstanceResponseBody struct {
	OrderInfo *CreateInstanceResponseBodyOrderInfo `json:"OrderInfo,omitempty" xml:"OrderInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
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
	// example:
	//
	// f-cn-zvp2q0zik06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 210406354694567
	OrderId           *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	StorageInstanceId *string `json:"StorageInstanceId,omitempty" xml:"StorageInstanceId,omitempty"`
	StorageOrderId    *int64  `json:"StorageOrderId,omitempty" xml:"StorageOrderId,omitempty"`
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

func (s *CreateInstanceResponseBodyOrderInfo) SetStorageInstanceId(v string) *CreateInstanceResponseBodyOrderInfo {
	s.StorageInstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyOrderInfo) SetStorageOrderId(v int64) *CreateInstanceResponseBodyOrderInfo {
	s.StorageOrderId = &v
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
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// di-593440390152545
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region       *string                             `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec *CreateNamespaceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s CreateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) SetHa(v bool) *CreateNamespaceRequest {
	s.Ha = &v
	return s
}

func (s *CreateNamespaceRequest) SetInstanceId(v string) *CreateNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespace(v string) *CreateNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *CreateNamespaceRequest) SetRegion(v string) *CreateNamespaceRequest {
	s.Region = &v
	return s
}

func (s *CreateNamespaceRequest) SetResourceSpec(v *CreateNamespaceRequestResourceSpec) *CreateNamespaceRequest {
	s.ResourceSpec = v
	return s
}

type CreateNamespaceRequestResourceSpec struct {
	// example:
	//
	// 30
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 120
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateNamespaceRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequestResourceSpec) SetCpu(v int32) *CreateNamespaceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateNamespaceRequestResourceSpec) SetMemoryGB(v int32) *CreateNamespaceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type CreateNamespaceShrinkRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// di-593440390152545
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
}

func (s CreateNamespaceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceShrinkRequest) SetHa(v bool) *CreateNamespaceShrinkRequest {
	s.Ha = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetInstanceId(v string) *CreateNamespaceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetNamespace(v string) *CreateNamespaceShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetRegion(v string) *CreateNamespaceShrinkRequest {
	s.Region = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetResourceSpecShrink(v string) *CreateNamespaceShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

type CreateNamespaceResponseBody struct {
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
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
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetRegion(v string) *DeleteInstanceRequest {
	s.Region = &v
	return s
}

type DeleteInstanceResponseBody struct {
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
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
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// di-593439443804417
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) SetInstanceId(v string) *DeleteNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNamespaceRequest) SetNamespace(v string) *DeleteNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteNamespaceRequest) SetRegion(v string) *DeleteNamespaceRequest {
	s.Region = &v
	return s
}

type DeleteNamespaceResponseBody struct {
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
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
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region          *string                         `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*DescribeInstancesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetArchitectureType(v string) *DescribeInstancesRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesRequest) SetChargeType(v string) *DescribeInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageIndex(v int32) *DescribeInstancesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegion(v string) *DescribeInstancesRequest {
	s.Region = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetTags(v []*DescribeInstancesRequestTags) *DescribeInstancesRequest {
	s.Tags = v
	return s
}

type DescribeInstancesRequestTags struct {
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// ys
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTags) SetKey(v string) *DescribeInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTags) SetValue(v string) *DescribeInstancesRequestTags {
	s.Value = &v
	return s
}

type DescribeInstancesShrinkRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TagsShrink      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesShrinkRequest) SetArchitectureType(v string) *DescribeInstancesShrinkRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetChargeType(v string) *DescribeInstancesShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetInstanceId(v string) *DescribeInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetPageIndex(v int32) *DescribeInstancesShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetPageSize(v int32) *DescribeInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetRegion(v string) *DescribeInstancesShrinkRequest {
	s.Region = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetResourceGroupId(v string) *DescribeInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetTagsShrink(v string) *DescribeInstancesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// C8DF2A5B-6FBA-5651-A3D4-960F36640E6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	AskClusterId     *string `json:"AskClusterId,omitempty" xml:"AskClusterId,omitempty"`
	// example:
	//
	// PRE
	ChargeType   *string                                             `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterState *DescribeInstancesResponseBodyInstancesClusterState `json:"ClusterState,omitempty" xml:"ClusterState,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	ClusterStatus        *string                                                       `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	ClusterUsedResources []*DescribeInstancesResponseBodyInstancesClusterUsedResources `json:"ClusterUsedResources,omitempty" xml:"ClusterUsedResources,omitempty" type:"Repeated"`
	ClusterUsedStorage   *DescribeInstancesResponseBodyInstancesClusterUsedStorage     `json:"ClusterUsedStorage,omitempty" xml:"ClusterUsedStorage,omitempty" type:"Struct"`
	Ha                   *bool                                                         `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaResourceSpec       *DescribeInstancesResponseBodyInstancesHaResourceSpec         `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	HaVSwitchIds         []*string                                                     `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	HaVSwitchInfo        []*DescribeInstancesResponseBodyInstancesHaVSwitchInfo        `json:"HaVSwitchInfo,omitempty" xml:"HaVSwitchInfo,omitempty" type:"Repeated"`
	HaZoneId             *string                                                       `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// This parameter is required.
	HostAliases []*DescribeInstancesResponseBodyInstancesHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// example:
	//
	// f-cn-zvp2q0zik06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// vvp1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MonitorType  *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// example:
	//
	// NORMAL
	OrderState *string                                        `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	OssInfo    *DescribeInstancesResponseBodyInstancesOssInfo `json:"OssInfo,omitempty" xml:"OssInfo,omitempty" type:"Struct"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1629879567394
	ResourceCreateTime *int64 `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	// example:
	//
	// 1637337600000
	ResourceExpiredTime *int64  `json:"ResourceExpiredTime,omitempty" xml:"ResourceExpiredTime,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// b3690a1655da47
	ResourceId   *string                                             `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceSpec *DescribeInstancesResponseBodyInstancesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage      *DescribeInstancesResponseBodyInstancesStorage      `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	Tags         []*DescribeInstancesResponseBodyInstancesTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 1838996687368452
	Uid         *string                                              `json:"Uid,omitempty" xml:"Uid,omitempty"`
	VSwitchIds  []*string                                            `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VSwitchInfo []*DescribeInstancesResponseBodyInstancesVSwitchInfo `json:"VSwitchInfo,omitempty" xml:"VSwitchInfo,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-2ze9*******nxfmfcdi
	VpcId   *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcInfo *DescribeInstancesResponseBodyInstancesVpcInfo `json:"VpcInfo,omitempty" xml:"VpcInfo,omitempty" type:"Struct"`
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DescribeInstancesResponseBodyInstances) SetClusterState(v *DescribeInstancesResponseBodyInstancesClusterState) *DescribeInstancesResponseBodyInstances {
	s.ClusterState = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.ClusterStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterUsedResources(v []*DescribeInstancesResponseBodyInstancesClusterUsedResources) *DescribeInstancesResponseBodyInstances {
	s.ClusterUsedResources = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterUsedStorage(v *DescribeInstancesResponseBodyInstancesClusterUsedStorage) *DescribeInstancesResponseBodyInstances {
	s.ClusterUsedStorage = v
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

func (s *DescribeInstancesResponseBodyInstances) SetHaVSwitchInfo(v []*DescribeInstancesResponseBodyInstancesHaVSwitchInfo) *DescribeInstancesResponseBodyInstances {
	s.HaVSwitchInfo = v
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

func (s *DescribeInstancesResponseBodyInstances) SetOssInfo(v *DescribeInstancesResponseBodyInstancesOssInfo) *DescribeInstancesResponseBodyInstances {
	s.OssInfo = v
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

func (s *DescribeInstancesResponseBodyInstances) SetVSwitchInfo(v []*DescribeInstancesResponseBodyInstancesVSwitchInfo) *DescribeInstancesResponseBodyInstances {
	s.VSwitchInfo = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcInfo(v *DescribeInstancesResponseBodyInstancesVpcInfo) *DescribeInstancesResponseBodyInstances {
	s.VpcInfo = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetZoneId(v string) *DescribeInstancesResponseBodyInstances {
	s.ZoneId = &v
	return s
}

type DescribeInstancesResponseBodyInstancesClusterState struct {
	ClusterId     *string                                                         `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterStage  *DescribeInstancesResponseBodyInstancesClusterStateClusterStage `json:"ClusterStage,omitempty" xml:"ClusterStage,omitempty" type:"Struct"`
	CreateTimeout *bool                                                           `json:"CreateTimeout,omitempty" xml:"CreateTimeout,omitempty"`
	Status        *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus     *string                                                         `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	Url           *string                                                         `json:"Url,omitempty" xml:"Url,omitempty"`
	UserSlbDto    *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto   `json:"UserSlbDto,omitempty" xml:"UserSlbDto,omitempty" type:"Struct"`
	VpcCidr       *string                                                         `json:"VpcCidr,omitempty" xml:"VpcCidr,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesClusterState) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterState) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesClusterState {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetClusterStage(v *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) *DescribeInstancesResponseBodyInstancesClusterState {
	s.ClusterStage = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetCreateTimeout(v bool) *DescribeInstancesResponseBodyInstancesClusterState {
	s.CreateTimeout = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetStatus(v string) *DescribeInstancesResponseBodyInstancesClusterState {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetSubStatus(v string) *DescribeInstancesResponseBodyInstancesClusterState {
	s.SubStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetUrl(v string) *DescribeInstancesResponseBodyInstancesClusterState {
	s.Url = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetUserSlbDto(v *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) *DescribeInstancesResponseBodyInstancesClusterState {
	s.UserSlbDto = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetVpcCidr(v string) *DescribeInstancesResponseBodyInstancesClusterState {
	s.VpcCidr = &v
	return s
}

type DescribeInstancesResponseBodyInstancesClusterStateClusterStage struct {
	ClusterId            *string                                                                               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CurrentStage         *int32                                                                                `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	Message              *string                                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Status               *string                                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalStageWithWeight []*DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight `json:"TotalStageWithWeight,omitempty" xml:"TotalStageWithWeight,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesClusterStateClusterStage) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterStateClusterStage) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) SetCurrentStage(v int32) *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	s.CurrentStage = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) SetMessage(v string) *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	s.Message = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) SetStatus(v string) *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) SetTotalStageWithWeight(v []*DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	s.TotalStageWithWeight = v
	return s
}

type DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight struct {
	StepIndex *int32  `json:"StepIndex,omitempty" xml:"StepIndex,omitempty"`
	StepName  *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	Weight    *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) SetStepIndex(v int32) *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight {
	s.StepIndex = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) SetStepName(v string) *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight {
	s.StepName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) SetWeight(v int32) *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight {
	s.Weight = &v
	return s
}

type DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto struct {
	ExistSlb         *bool                                                                           `json:"ExistSlb,omitempty" xml:"ExistSlb,omitempty"`
	SlbId            *string                                                                         `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SlbIp            *string                                                                         `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	SlbStatus        *string                                                                         `json:"SlbStatus,omitempty" xml:"SlbStatus,omitempty"`
	UserSlbListeners []*DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners `json:"UserSlbListeners,omitempty" xml:"UserSlbListeners,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) SetExistSlb(v bool) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	s.ExistSlb = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) SetSlbId(v string) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	s.SlbId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) SetSlbIp(v string) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	s.SlbIp = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) SetSlbStatus(v string) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	s.SlbStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) SetUserSlbListeners(v []*DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	s.UserSlbListeners = v
	return s
}

type DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners struct {
	ListenersStatus *string `json:"ListenersStatus,omitempty" xml:"ListenersStatus,omitempty"`
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) SetListenersStatus(v string) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners {
	s.ListenersStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) SetPort(v string) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners {
	s.Port = &v
	return s
}

type DescribeInstancesResponseBodyInstancesClusterUsedResources struct {
	ClusterId      *string  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Ha             *bool    `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaUsedCpu      *float32 `json:"HaUsedCpu,omitempty" xml:"HaUsedCpu,omitempty"`
	HaUsedMemory   *float32 `json:"HaUsedMemory,omitempty" xml:"HaUsedMemory,omitempty"`
	HaUsedResource *float32 `json:"HaUsedResource,omitempty" xml:"HaUsedResource,omitempty"`
	UsedCpu        *float32 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
	UsedMemory     *float32 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
	UsedResource   *float32 `json:"UsedResource,omitempty" xml:"UsedResource,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesClusterUsedResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterUsedResources) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetHa(v bool) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.Ha = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetHaUsedCpu(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.HaUsedCpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetHaUsedMemory(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.HaUsedMemory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetHaUsedResource(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.HaUsedResource = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetUsedCpu(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.UsedCpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetUsedMemory(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.UsedMemory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetUsedResource(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.UsedResource = &v
	return s
}

type DescribeInstancesResponseBodyInstancesClusterUsedStorage struct {
	ClusterId   *string  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	UsedStorage *float32 `json:"UsedStorage,omitempty" xml:"UsedStorage,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesClusterUsedStorage) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterUsedStorage) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedStorage) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesClusterUsedStorage {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedStorage) SetUsedStorage(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedStorage {
	s.UsedStorage = &v
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

type DescribeInstancesResponseBodyInstancesHaVSwitchInfo struct {
	AvailableIpAddressCount *int64  `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchCidr             *string `json:"VSwitchCidr,omitempty" xml:"VSwitchCidr,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName             *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	VpcId                   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesHaVSwitchInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesHaVSwitchInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetAvailableIpAddressCount(v int64) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetDescription(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetVSwitchCidr(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.VSwitchCidr = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetVSwitchId(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetVSwitchName(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.VSwitchName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.ZoneId = &v
	return s
}

type DescribeInstancesResponseBodyInstancesHostAliases struct {
	// This parameter is required.
	HostNames []*string `json:"HostNames,omitempty" xml:"HostNames,omitempty" type:"Repeated"`
	// This parameter is required.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
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

type DescribeInstancesResponseBodyInstancesOssInfo struct {
	AccessId               *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	AccessKey              *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	Bucket                 *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	BucketVersioningStatus *string `json:"BucketVersioningStatus,omitempty" xml:"BucketVersioningStatus,omitempty"`
	Endpoint               *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesOssInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesOssInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) SetAccessId(v string) *DescribeInstancesResponseBodyInstancesOssInfo {
	s.AccessId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) SetAccessKey(v string) *DescribeInstancesResponseBodyInstancesOssInfo {
	s.AccessKey = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) SetBucket(v string) *DescribeInstancesResponseBodyInstancesOssInfo {
	s.Bucket = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) SetBucketVersioningStatus(v string) *DescribeInstancesResponseBodyInstancesOssInfo {
	s.BucketVersioningStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) SetEndpoint(v string) *DescribeInstancesResponseBodyInstancesOssInfo {
	s.Endpoint = &v
	return s
}

type DescribeInstancesResponseBodyInstancesResourceSpec struct {
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 40
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
	FullyManaged *bool                                             `json:"FullyManaged,omitempty" xml:"FullyManaged,omitempty"`
	OrderState   *string                                           `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	Oss          *DescribeInstancesResponseBodyInstancesStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s DescribeInstancesResponseBodyInstancesStorage) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesStorage) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesStorage) SetFullyManaged(v bool) *DescribeInstancesResponseBodyInstancesStorage {
	s.FullyManaged = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesStorage) SetOrderState(v string) *DescribeInstancesResponseBodyInstancesStorage {
	s.OrderState = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesStorage) SetOss(v *DescribeInstancesResponseBodyInstancesStorageOss) *DescribeInstancesResponseBodyInstancesStorage {
	s.Oss = v
	return s
}

type DescribeInstancesResponseBodyInstancesStorageOss struct {
	// example:
	//
	// oss_flink
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
	// example:
	//
	// flink
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
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

type DescribeInstancesResponseBodyInstancesVSwitchInfo struct {
	AvailableIpAddressCount *string `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchCidr             *string `json:"VSwitchCidr,omitempty" xml:"VSwitchCidr,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName             *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	VpcId                   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesVSwitchInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesVSwitchInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetAvailableIpAddressCount(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetDescription(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetVSwitchCidr(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.VSwitchCidr = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetVSwitchId(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetVSwitchName(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.VSwitchName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.ZoneId = &v
	return s
}

type DescribeInstancesResponseBodyInstancesVpcInfo struct {
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName     *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesVpcInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesVpcInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetCidrBlock(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.CidrBlock = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetDescription(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetStatus(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetVpcName(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.VpcName = &v
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
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// di-590843445844225
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string                          `json:"Region,omitempty" xml:"Region,omitempty"`
	Tags   []*DescribeNamespacesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequest) SetHa(v bool) *DescribeNamespacesRequest {
	s.Ha = &v
	return s
}

func (s *DescribeNamespacesRequest) SetInstanceId(v string) *DescribeNamespacesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNamespacesRequest) SetNamespace(v string) *DescribeNamespacesRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespacesRequest) SetPageIndex(v int32) *DescribeNamespacesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeNamespacesRequest) SetPageSize(v int32) *DescribeNamespacesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesRequest) SetRegion(v string) *DescribeNamespacesRequest {
	s.Region = &v
	return s
}

func (s *DescribeNamespacesRequest) SetTags(v []*DescribeNamespacesRequestTags) *DescribeNamespacesRequest {
	s.Tags = v
	return s
}

type DescribeNamespacesRequestTags struct {
	// example:
	//
	// FLink
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNamespacesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequestTags) SetKey(v string) *DescribeNamespacesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeNamespacesRequestTags) SetValue(v string) *DescribeNamespacesRequestTags {
	s.Value = &v
	return s
}

type DescribeNamespacesShrinkRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// di-590843445844225
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeNamespacesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesShrinkRequest) SetHa(v bool) *DescribeNamespacesShrinkRequest {
	s.Ha = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetInstanceId(v string) *DescribeNamespacesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetNamespace(v string) *DescribeNamespacesShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetPageIndex(v int32) *DescribeNamespacesShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetPageSize(v int32) *DescribeNamespacesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetRegion(v string) *DescribeNamespacesShrinkRequest {
	s.Region = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetTagsShrink(v string) *DescribeNamespacesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type DescribeNamespacesResponseBody struct {
	Namespaces []*DescribeNamespacesResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	// example:
	//
	// 1629879567394
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1629879567394
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Ha          *bool  `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// example:
	//
	// ns-1
	Namespace    *string                                               `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ResourceSpec *DescribeNamespacesResponseBodyNamespacesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	ResourceUsed *DescribeNamespacesResponseBodyNamespacesResourceUsed `json:"ResourceUsed,omitempty" xml:"ResourceUsed,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESS
	Status *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*DescribeNamespacesResponseBodyNamespacesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 40
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
	// example:
	//
	// 2
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Cu  *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// 4
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
	// example:
	//
	// flink
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
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
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Regions  []*DescribeSupportedRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// B21DC47E-8928-199A-9F32-36D45E4693B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeSupportedRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportedRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportedRegionsResponseBody) SetPageIndex(v int32) *DescribeSupportedRegionsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetPageSize(v int32) *DescribeSupportedRegionsResponseBody {
	s.PageSize = &v
	return s
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

func (s *DescribeSupportedRegionsResponseBody) SetTotalCount(v int64) *DescribeSupportedRegionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetTotalPage(v int32) *DescribeSupportedRegionsResponseBody {
	s.TotalPage = &v
	return s
}

type DescribeSupportedRegionsResponseBodyRegions struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Extra       *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 华北2 (北京)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeSupportedRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportedRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeSupportedRegionsResponseBodyRegions) SetDescription(v string) *DescribeSupportedRegionsResponseBodyRegions {
	s.Description = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBodyRegions) SetExtra(v string) *DescribeSupportedRegionsResponseBodyRegions {
	s.Extra = &v
	return s
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
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
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
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 23A9C718-DDAB-1696-B025-18FBC830F7C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32    `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	ZoneIds   []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s DescribeSupportedZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupportedZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportedZonesResponseBody) SetPageIndex(v int32) *DescribeSupportedZonesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetPageSize(v int32) *DescribeSupportedZonesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetRequestId(v string) *DescribeSupportedZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetSuccess(v bool) *DescribeSupportedZonesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetTotalCount(v int64) *DescribeSupportedZonesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSupportedZonesResponseBody) SetTotalPage(v int32) *DescribeSupportedZonesResponseBody {
	s.TotalPage = &v
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
	// example:
	//
	// 27AE00
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vvpinstance
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
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// tag
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
	// example:
	//
	// 27AE00
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0E5D17CE-BD83-5DC9-8CD2-3C40C2F7A135
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 87AE00
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
	// example:
	//
	// f-cn-tyts
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// vvpinstance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// tag
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *ModifyPrepayInstanceSpecRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	// if can be null:
	// true
	HaVSwitchIds []*string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	// if can be null:
	// true
	HaZoneId *string `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec *ModifyPrepayInstanceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ModifyPrepayInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequest) SetHa(v bool) *ModifyPrepayInstanceSpecRequest {
	s.Ha = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetHaResourceSpec(v *ModifyPrepayInstanceSpecRequestHaResourceSpec) *ModifyPrepayInstanceSpecRequest {
	s.HaResourceSpec = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetHaVSwitchIds(v []*string) *ModifyPrepayInstanceSpecRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetHaZoneId(v string) *ModifyPrepayInstanceSpecRequest {
	s.HaZoneId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetInstanceId(v string) *ModifyPrepayInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetRegion(v string) *ModifyPrepayInstanceSpecRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetResourceSpec(v *ModifyPrepayInstanceSpecRequestResourceSpec) *ModifyPrepayInstanceSpecRequest {
	s.ResourceSpec = v
	return s
}

type ModifyPrepayInstanceSpecRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyPrepayInstanceSpecRequestHaResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestHaResourceSpec) SetCpu(v int32) *ModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestHaResourceSpec) SetMemoryGB(v int32) *ModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

type ModifyPrepayInstanceSpecRequestResourceSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyPrepayInstanceSpecRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestResourceSpec) SetCpu(v int32) *ModifyPrepayInstanceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestResourceSpec) SetMemoryGB(v int32) *ModifyPrepayInstanceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type ModifyPrepayInstanceSpecShrinkRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpecShrink *string `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty"`
	// if can be null:
	// true
	HaVSwitchIdsShrink *string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty"`
	// if can be null:
	// true
	HaZoneId *string `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
}

func (s ModifyPrepayInstanceSpecShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetHa(v bool) *ModifyPrepayInstanceSpecShrinkRequest {
	s.Ha = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetHaResourceSpecShrink(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.HaResourceSpecShrink = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetHaVSwitchIdsShrink(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.HaVSwitchIdsShrink = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetHaZoneId(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.HaZoneId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetInstanceId(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetRegion(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetResourceSpecShrink(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

type ModifyPrepayInstanceSpecResponseBody struct {
	// example:
	//
	// 210406354690749
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
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
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// di-593440219799842
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec *ModifyPrepayNamespaceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ModifyPrepayNamespaceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecRequest) SetInstanceId(v string) *ModifyPrepayNamespaceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequest) SetNamespace(v string) *ModifyPrepayNamespaceSpecRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequest) SetRegion(v string) *ModifyPrepayNamespaceSpecRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequest) SetResourceSpec(v *ModifyPrepayNamespaceSpecRequestResourceSpec) *ModifyPrepayNamespaceSpecRequest {
	s.ResourceSpec = v
	return s
}

type ModifyPrepayNamespaceSpecRequestResourceSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyPrepayNamespaceSpecRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecRequestResourceSpec) SetCpu(v int32) *ModifyPrepayNamespaceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestResourceSpec) SetMemoryGB(v int32) *ModifyPrepayNamespaceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type ModifyPrepayNamespaceSpecShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// di-593440219799842
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
}

func (s ModifyPrepayNamespaceSpecShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) SetInstanceId(v string) *ModifyPrepayNamespaceSpecShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) SetNamespace(v string) *ModifyPrepayNamespaceSpecShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) SetRegion(v string) *ModifyPrepayNamespaceSpecShrinkRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecShrinkRequest) SetResourceSpecShrink(v string) *ModifyPrepayNamespaceSpecShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

type ModifyPrepayNamespaceSpecResponseBody struct {
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// This parameter is required.
	NamespaceResourceSpecs []*QueryConvertInstancePriceRequestNamespaceResourceSpecs `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryConvertInstancePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequest) SetDuration(v int32) *QueryConvertInstancePriceRequest {
	s.Duration = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetInstanceId(v string) *QueryConvertInstancePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetIsAutoRenew(v bool) *QueryConvertInstancePriceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetNamespaceResourceSpecs(v []*QueryConvertInstancePriceRequestNamespaceResourceSpecs) *QueryConvertInstancePriceRequest {
	s.NamespaceResourceSpecs = v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetPricingCycle(v string) *QueryConvertInstancePriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetRegion(v string) *QueryConvertInstancePriceRequest {
	s.Region = &v
	return s
}

type QueryConvertInstancePriceRequestNamespaceResourceSpecs struct {
	// This parameter is required.
	//
	// example:
	//
	// lm-test-default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	ResourceSpec *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s QueryConvertInstancePriceRequestNamespaceResourceSpecs) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceRequestNamespaceResourceSpecs) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecs) SetNamespace(v string) *QueryConvertInstancePriceRequestNamespaceResourceSpecs {
	s.Namespace = &v
	return s
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecs) SetResourceSpec(v *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) *QueryConvertInstancePriceRequestNamespaceResourceSpecs {
	s.ResourceSpec = v
	return s
}

type QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// 6
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 24
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) SetCpu(v int32) *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) SetMemoryGB(v int32) *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec {
	s.MemoryGB = &v
	return s
}

type QueryConvertInstancePriceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// This parameter is required.
	NamespaceResourceSpecsShrink *string `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryConvertInstancePriceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConvertInstancePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceShrinkRequest) SetDuration(v int32) *QueryConvertInstancePriceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetInstanceId(v string) *QueryConvertInstancePriceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetIsAutoRenew(v bool) *QueryConvertInstancePriceShrinkRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetNamespaceResourceSpecsShrink(v string) *QueryConvertInstancePriceShrinkRequest {
	s.NamespaceResourceSpecsShrink = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetPricingCycle(v string) *QueryConvertInstancePriceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetRegion(v string) *QueryConvertInstancePriceShrinkRequest {
	s.Region = &v
	return s
}

type QueryConvertInstancePriceResponseBody struct {
	PriceInfo *QueryConvertInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
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
	// example:
	//
	// ORDER.INST_HAS_UNPAID_ORDER
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// CNY
	Currency       *string                                                       `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryConvertInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 655.2
	DiscountAmount     *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	IsContractActivity *bool    `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// example:
	//
	// 存在未支付订单，请先支付或取消原有订单
	Message            *string                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryConvertInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 4368
	OriginalAmount     *float32                                               `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryConvertInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                                `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                                `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 3712.8
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
	// example:
	//
	// ￥1,391.5 优惠券 (有效期至 03/23/2022)
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// ￥1,391.5 优惠券
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// 500011220010099
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
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
	// example:
	//
	// 买满1年，立享官网价格8.5折优惠。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 587
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

type QueryCreateInstancePriceRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 1
	Duration *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extra    *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Ha       *bool   `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *QueryCreateInstancePriceRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	// example:
	//
	// rtc-e2e-test-post
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 500041860100636
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region           *string                                      `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec     *QueryCreateInstancePriceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage          *QueryCreateInstancePriceRequestStorage      `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	UsePromotionCode *bool                                        `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	VSwitchIds       []*string                                    `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-2ze9xoh8qyt1rnxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryCreateInstancePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequest) SetArchitectureType(v string) *QueryCreateInstancePriceRequest {
	s.ArchitectureType = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetAutoRenew(v bool) *QueryCreateInstancePriceRequest {
	s.AutoRenew = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetChargeType(v string) *QueryCreateInstancePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetDuration(v int32) *QueryCreateInstancePriceRequest {
	s.Duration = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetExtra(v string) *QueryCreateInstancePriceRequest {
	s.Extra = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetHa(v bool) *QueryCreateInstancePriceRequest {
	s.Ha = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetHaResourceSpec(v *QueryCreateInstancePriceRequestHaResourceSpec) *QueryCreateInstancePriceRequest {
	s.HaResourceSpec = v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetInstanceName(v string) *QueryCreateInstancePriceRequest {
	s.InstanceName = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetPricingCycle(v string) *QueryCreateInstancePriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetPromotionCode(v string) *QueryCreateInstancePriceRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetRegion(v string) *QueryCreateInstancePriceRequest {
	s.Region = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetResourceSpec(v *QueryCreateInstancePriceRequestResourceSpec) *QueryCreateInstancePriceRequest {
	s.ResourceSpec = v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetStorage(v *QueryCreateInstancePriceRequestStorage) *QueryCreateInstancePriceRequest {
	s.Storage = v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetUsePromotionCode(v bool) *QueryCreateInstancePriceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetVSwitchIds(v []*string) *QueryCreateInstancePriceRequest {
	s.VSwitchIds = v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetVpcId(v string) *QueryCreateInstancePriceRequest {
	s.VpcId = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetZoneId(v string) *QueryCreateInstancePriceRequest {
	s.ZoneId = &v
	return s
}

type QueryCreateInstancePriceRequestHaResourceSpec struct {
	// if can be null:
	// false
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// if can be null:
	// false
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryCreateInstancePriceRequestHaResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestHaResourceSpec) SetCpu(v int32) *QueryCreateInstancePriceRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryCreateInstancePriceRequestHaResourceSpec) SetMemoryGB(v int32) *QueryCreateInstancePriceRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

type QueryCreateInstancePriceRequestResourceSpec struct {
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 16
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryCreateInstancePriceRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestResourceSpec) SetCpu(v int32) *QueryCreateInstancePriceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryCreateInstancePriceRequestResourceSpec) SetMemoryGB(v int32) *QueryCreateInstancePriceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type QueryCreateInstancePriceRequestStorage struct {
	Oss *QueryCreateInstancePriceRequestStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s QueryCreateInstancePriceRequestStorage) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequestStorage) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestStorage) SetOss(v *QueryCreateInstancePriceRequestStorageOss) *QueryCreateInstancePriceRequestStorage {
	s.Oss = v
	return s
}

type QueryCreateInstancePriceRequestStorageOss struct {
	// example:
	//
	// quicktracing
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s QueryCreateInstancePriceRequestStorageOss) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceRequestStorageOss) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestStorageOss) SetBucket(v string) *QueryCreateInstancePriceRequestStorageOss {
	s.Bucket = &v
	return s
}

type QueryCreateInstancePriceShrinkRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 1
	Duration *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extra    *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Ha       *bool   `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpecShrink *string `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty"`
	// example:
	//
	// rtc-e2e-test-post
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 500041860100636
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
	StorageShrink      *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	UsePromotionCode   *bool   `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	VSwitchIdsShrink   *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// vpc-2ze9xoh8qyt1rnxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryCreateInstancePriceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreateInstancePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceShrinkRequest) SetArchitectureType(v string) *QueryCreateInstancePriceShrinkRequest {
	s.ArchitectureType = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetAutoRenew(v bool) *QueryCreateInstancePriceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetChargeType(v string) *QueryCreateInstancePriceShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetDuration(v int32) *QueryCreateInstancePriceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetExtra(v string) *QueryCreateInstancePriceShrinkRequest {
	s.Extra = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetHa(v bool) *QueryCreateInstancePriceShrinkRequest {
	s.Ha = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetHaResourceSpecShrink(v string) *QueryCreateInstancePriceShrinkRequest {
	s.HaResourceSpecShrink = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetInstanceName(v string) *QueryCreateInstancePriceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetPricingCycle(v string) *QueryCreateInstancePriceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetPromotionCode(v string) *QueryCreateInstancePriceShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetRegion(v string) *QueryCreateInstancePriceShrinkRequest {
	s.Region = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetResourceSpecShrink(v string) *QueryCreateInstancePriceShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetStorageShrink(v string) *QueryCreateInstancePriceShrinkRequest {
	s.StorageShrink = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetUsePromotionCode(v bool) *QueryCreateInstancePriceShrinkRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetVSwitchIdsShrink(v string) *QueryCreateInstancePriceShrinkRequest {
	s.VSwitchIdsShrink = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetVpcId(v string) *QueryCreateInstancePriceShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetZoneId(v string) *QueryCreateInstancePriceShrinkRequest {
	s.ZoneId = &v
	return s
}

type QueryCreateInstancePriceResponseBody struct {
	PriceInfo *QueryCreateInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// ORDER.INST_HAS_UNPAID_ORDER
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// CNY
	Currency       *string                                                      `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryCreateInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 655.2
	DiscountAmount     *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	IsContractActivity *bool    `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// example:
	//
	// 存在未支付订单，请先支付或取消原有订单
	Message            *string                                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryCreateInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 4368
	OriginalAmount     *float32                                              `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryCreateInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                               `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                               `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 3712.8
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
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
	// example:
	//
	// ￥1,391.5 优惠券 (有效期至 03/23/2022)
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// ￥1,391.5 优惠券
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// 500011220010099
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
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
	// example:
	//
	// 买满1年，立享官网价格8.5折优惠。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 587
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *QueryModifyInstancePriceRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	// if can be null:
	// true
	HaVSwitchIds []*string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	// if can be null:
	// true
	HaZoneId *string `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec *QueryModifyInstancePriceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s QueryModifyInstancePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequest) SetHa(v bool) *QueryModifyInstancePriceRequest {
	s.Ha = &v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetHaResourceSpec(v *QueryModifyInstancePriceRequestHaResourceSpec) *QueryModifyInstancePriceRequest {
	s.HaResourceSpec = v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetHaVSwitchIds(v []*string) *QueryModifyInstancePriceRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetHaZoneId(v string) *QueryModifyInstancePriceRequest {
	s.HaZoneId = &v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetInstanceId(v string) *QueryModifyInstancePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetRegion(v string) *QueryModifyInstancePriceRequest {
	s.Region = &v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetResourceSpec(v *QueryModifyInstancePriceRequestResourceSpec) *QueryModifyInstancePriceRequest {
	s.ResourceSpec = v
	return s
}

type QueryModifyInstancePriceRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryModifyInstancePriceRequestHaResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequestHaResourceSpec) SetCpu(v int32) *QueryModifyInstancePriceRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryModifyInstancePriceRequestHaResourceSpec) SetMemoryGB(v int32) *QueryModifyInstancePriceRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

type QueryModifyInstancePriceRequestResourceSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryModifyInstancePriceRequestResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequestResourceSpec) SetCpu(v int32) *QueryModifyInstancePriceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryModifyInstancePriceRequestResourceSpec) SetMemoryGB(v int32) *QueryModifyInstancePriceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

type QueryModifyInstancePriceShrinkRequest struct {
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpecShrink *string `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty"`
	// if can be null:
	// true
	HaVSwitchIdsShrink *string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty"`
	// if can be null:
	// true
	HaZoneId *string `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
}

func (s QueryModifyInstancePriceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryModifyInstancePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceShrinkRequest) SetHa(v bool) *QueryModifyInstancePriceShrinkRequest {
	s.Ha = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetHaResourceSpecShrink(v string) *QueryModifyInstancePriceShrinkRequest {
	s.HaResourceSpecShrink = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetHaVSwitchIdsShrink(v string) *QueryModifyInstancePriceShrinkRequest {
	s.HaVSwitchIdsShrink = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetHaZoneId(v string) *QueryModifyInstancePriceShrinkRequest {
	s.HaZoneId = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetInstanceId(v string) *QueryModifyInstancePriceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetRegion(v string) *QueryModifyInstancePriceShrinkRequest {
	s.Region = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetResourceSpecShrink(v string) *QueryModifyInstancePriceShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

type QueryModifyInstancePriceResponseBody struct {
	PriceInfo *QueryModifyInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// ORDER.INST_HAS_UNPAID_ORDER
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// CNY
	Currency       *string                                                      `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryModifyInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 655.2
	DiscountAmount     *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	IsContractActivity *bool    `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// example:
	//
	// 存在未支付订单，请先支付或取消原有订单
	Message            *string                                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryModifyInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 4368
	OriginalAmount     *float32                                              `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryModifyInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                               `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                               `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 3712.8
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
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
	// example:
	//
	// ￥1,391.5 优惠券 (有效期至 03/23/2022)
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// ￥1,391.5 优惠券
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// 500011220010099
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
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
	// example:
	//
	// 买满1年，立享官网价格8.5折优惠。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 587
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sc_flinkserverless_public_cn-7e22ae5sess
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryRenewInstancePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryRenewInstancePriceRequest) SetDuration(v int32) *QueryRenewInstancePriceRequest {
	s.Duration = &v
	return s
}

func (s *QueryRenewInstancePriceRequest) SetInstanceId(v string) *QueryRenewInstancePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRenewInstancePriceRequest) SetPricingCycle(v string) *QueryRenewInstancePriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryRenewInstancePriceRequest) SetRegion(v string) *QueryRenewInstancePriceRequest {
	s.Region = &v
	return s
}

type QueryRenewInstancePriceResponseBody struct {
	PriceInfo *QueryRenewInstancePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// ORDER.INST_HAS_UNPAID_ORDER
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// CNY
	Currency       *string                                                     `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DepreciateInfo *QueryRenewInstancePriceResponseBodyPriceInfoDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// example:
	//
	// 655.2
	DiscountAmount     *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	IsContractActivity *bool    `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// example:
	//
	// 存在未支付订单，请先支付或取消原有订单
	Message            *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	OptionalPromotions []*QueryRenewInstancePriceResponseBodyPriceInfoOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	// example:
	//
	// 4368
	OriginalAmount     *float32                                             `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	Rules              []*QueryRenewInstancePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	StandDiscountPrice *string                                              `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	StandPrice         *string                                              `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// example:
	//
	// 3712.8
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
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
	// example:
	//
	// ￥1,391.5 优惠券 (有效期至 03/23/2022)
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// ￥1,391.5 优惠券
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// 500011220010099
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
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
	// example:
	//
	// 买满1年，立享官网价格8.5折优惠。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 587
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sc_flinkserverless_public_cn-7e22ae5sess
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetPricingCycle(v string) *RenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewInstanceRequest) SetRegion(v string) *RenewInstanceRequest {
	s.Region = &v
	return s
}

type RenewInstanceResponseBody struct {
	// orderId
	//
	// example:
	//
	// 210406354690749
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
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
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vvpinstance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// tag
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
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 154FT
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
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vvpinstance
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
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F59597FC-CD05-536D-B75B-6F45B8CC8539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 154FT
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

// Summary:
//
// 按量付费转包年包月
//
// @param tmpReq - ConvertInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertInstanceResponse
func (client *Client) ConvertInstanceWithOptions(tmpReq *ConvertInstanceRequest, runtime *util.RuntimeOptions) (_result *ConvertInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ConvertInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NamespaceResourceSpecs)) {
		request.NamespaceResourceSpecsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NamespaceResourceSpecs, tea.String("NamespaceResourceSpecs"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoRenew)) {
		body["IsAutoRenew"] = request.IsAutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceResourceSpecsShrink)) {
		body["NamespaceResourceSpecs"] = request.NamespaceResourceSpecsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConvertInstance"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 按量付费转包年包月
//
// @param request - ConvertInstanceRequest
//
// @return ConvertInstanceResponse
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

// Summary:
//
// 创建实例
//
// @param tmpReq - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(tmpReq *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HaResourceSpec)) {
		request.HaResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaResourceSpec, tea.String("HaResourceSpec"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.HaVSwitchIds)) {
		request.HaVSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaVSwitchIds, tea.String("HaVSwitchIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceSpec)) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, tea.String("ResourceSpec"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Storage)) {
		request.StorageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Storage, tea.String("Storage"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VSwitchIds)) {
		request.VSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIds, tea.String("VSwitchIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArchitectureType)) {
		body["ArchitectureType"] = request.ArchitectureType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Ha)) {
		body["Ha"] = request.Ha
	}

	if !tea.BoolValue(util.IsUnset(request.HaResourceSpecShrink)) {
		body["HaResourceSpec"] = request.HaResourceSpecShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HaVSwitchIdsShrink)) {
		body["HaVSwitchIds"] = request.HaVSwitchIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HaZoneId)) {
		body["HaZoneId"] = request.HaZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorType)) {
		body["MonitorType"] = request.MonitorType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionCode)) {
		body["PromotionCode"] = request.PromotionCode
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSpecShrink)) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StorageShrink)) {
		body["Storage"] = request.StorageShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		body["Tag"] = request.TagShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UsePromotionCode)) {
		body["UsePromotionCode"] = request.UsePromotionCode
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIdsShrink)) {
		body["VSwitchIds"] = request.VSwitchIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 创建实例
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
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

// Summary:
//
// 创建命名空间
//
// @param tmpReq - CreateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithOptions(tmpReq *CreateNamespaceRequest, runtime *util.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateNamespaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceSpec)) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, tea.String("ResourceSpec"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ha)) {
		body["Ha"] = request.Ha
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSpecShrink)) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNamespace"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 创建命名空间
//
// @param request - CreateNamespaceRequest
//
// @return CreateNamespaceResponse
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

// Summary:
//
// 释放按量付费的实例
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 释放按量付费的实例
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
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

// Summary:
//
// 删除namespace
//
// @param request - DeleteNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *util.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNamespace"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 删除namespace
//
// @param request - DeleteNamespaceRequest
//
// @return DeleteNamespaceResponse
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

// Summary:
//
// instance列表
//
// @param tmpReq - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(tmpReq *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// instance列表
//
// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
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

// Summary:
//
// namespace列表
//
// @param tmpReq - DescribeNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNamespacesResponse
func (client *Client) DescribeNamespacesWithOptions(tmpReq *DescribeNamespacesRequest, runtime *util.RuntimeOptions) (_result *DescribeNamespacesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeNamespacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespaces"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// namespace列表
//
// @param request - DescribeNamespacesRequest
//
// @return DescribeNamespacesResponse
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

// Summary:
//
// 获取支持的region列表
//
// @param request - DescribeSupportedRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupportedRegionsResponse
func (client *Client) DescribeSupportedRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeSupportedRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeSupportedRegions"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 获取支持的region列表
//
// @return DescribeSupportedRegionsResponse
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

// Summary:
//
// 获取支持的zoneId列表
//
// @param request - DescribeSupportedZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupportedZonesResponse
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
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 获取支持的zoneId列表
//
// @param request - DescribeSupportedZonesRequest
//
// @return DescribeSupportedZonesResponse
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

// Summary:
//
// 列举flinkasi标签
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
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 列举flinkasi标签
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

// Deprecated: OpenAPI ModifyPrepayInstanceSpec is deprecated, please use foasconsole::2021-10-28::ModifyInstanceSpec instead.
//
// Summary:
//
// 扩容/缩容
//
// @param tmpReq - ModifyPrepayInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPrepayInstanceSpecResponse
// Deprecated
func (client *Client) ModifyPrepayInstanceSpecWithOptions(tmpReq *ModifyPrepayInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyPrepayInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyPrepayInstanceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HaResourceSpec)) {
		request.HaResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaResourceSpec, tea.String("HaResourceSpec"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.HaVSwitchIds)) {
		request.HaVSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaVSwitchIds, tea.String("HaVSwitchIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceSpec)) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, tea.String("ResourceSpec"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ha)) {
		body["Ha"] = request.Ha
	}

	if !tea.BoolValue(util.IsUnset(request.HaResourceSpecShrink)) {
		body["HaResourceSpec"] = request.HaResourceSpecShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HaVSwitchIdsShrink)) {
		body["HaVSwitchIds"] = request.HaVSwitchIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HaZoneId)) {
		body["HaZoneId"] = request.HaZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSpecShrink)) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPrepayInstanceSpec"),
		Version:     tea.String("2021-10-28"),
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

// Deprecated: OpenAPI ModifyPrepayInstanceSpec is deprecated, please use foasconsole::2021-10-28::ModifyInstanceSpec instead.
//
// Summary:
//
// 扩容/缩容
//
// @param request - ModifyPrepayInstanceSpecRequest
//
// @return ModifyPrepayInstanceSpecResponse
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

// Deprecated: OpenAPI ModifyPrepayNamespaceSpec is deprecated, please use foasconsole::2021-10-28::ModifyNamespaceSpec instead.
//
// Summary:
//
// 修改namespace资源分配
//
// @param tmpReq - ModifyPrepayNamespaceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPrepayNamespaceSpecResponse
// Deprecated
func (client *Client) ModifyPrepayNamespaceSpecWithOptions(tmpReq *ModifyPrepayNamespaceSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyPrepayNamespaceSpecResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyPrepayNamespaceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceSpec)) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, tea.String("ResourceSpec"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSpecShrink)) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPrepayNamespaceSpec"),
		Version:     tea.String("2021-10-28"),
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

// Deprecated: OpenAPI ModifyPrepayNamespaceSpec is deprecated, please use foasconsole::2021-10-28::ModifyNamespaceSpec instead.
//
// Summary:
//
// 修改namespace资源分配
//
// @param request - ModifyPrepayNamespaceSpecRequest
//
// @return ModifyPrepayNamespaceSpecResponse
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

// Summary:
//
// 按量付费转包年包月询价
//
// @param tmpReq - QueryConvertInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConvertInstancePriceResponse
func (client *Client) QueryConvertInstancePriceWithOptions(tmpReq *QueryConvertInstancePriceRequest, runtime *util.RuntimeOptions) (_result *QueryConvertInstancePriceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryConvertInstancePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NamespaceResourceSpecs)) {
		request.NamespaceResourceSpecsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NamespaceResourceSpecs, tea.String("NamespaceResourceSpecs"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoRenew)) {
		body["IsAutoRenew"] = request.IsAutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceResourceSpecsShrink)) {
		body["NamespaceResourceSpecs"] = request.NamespaceResourceSpecsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryConvertInstancePrice"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 按量付费转包年包月询价
//
// @param request - QueryConvertInstancePriceRequest
//
// @return QueryConvertInstancePriceResponse
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

// Summary:
//
// 获取创建实例的价格
//
// @param tmpReq - QueryCreateInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCreateInstancePriceResponse
func (client *Client) QueryCreateInstancePriceWithOptions(tmpReq *QueryCreateInstancePriceRequest, runtime *util.RuntimeOptions) (_result *QueryCreateInstancePriceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryCreateInstancePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HaResourceSpec)) {
		request.HaResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaResourceSpec, tea.String("HaResourceSpec"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceSpec)) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, tea.String("ResourceSpec"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Storage)) {
		request.StorageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Storage, tea.String("Storage"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VSwitchIds)) {
		request.VSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIds, tea.String("VSwitchIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArchitectureType)) {
		body["ArchitectureType"] = request.ArchitectureType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Ha)) {
		body["Ha"] = request.Ha
	}

	if !tea.BoolValue(util.IsUnset(request.HaResourceSpecShrink)) {
		body["HaResourceSpec"] = request.HaResourceSpecShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionCode)) {
		body["PromotionCode"] = request.PromotionCode
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSpecShrink)) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StorageShrink)) {
		body["Storage"] = request.StorageShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UsePromotionCode)) {
		body["UsePromotionCode"] = request.UsePromotionCode
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIdsShrink)) {
		body["VSwitchIds"] = request.VSwitchIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCreateInstancePrice"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 获取创建实例的价格
//
// @param request - QueryCreateInstancePriceRequest
//
// @return QueryCreateInstancePriceResponse
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

// Summary:
//
// 查询付费类型为包年包月的实例修改资源规格的价格
//
// @param tmpReq - QueryModifyInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifyInstancePriceResponse
func (client *Client) QueryModifyInstancePriceWithOptions(tmpReq *QueryModifyInstancePriceRequest, runtime *util.RuntimeOptions) (_result *QueryModifyInstancePriceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryModifyInstancePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HaResourceSpec)) {
		request.HaResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaResourceSpec, tea.String("HaResourceSpec"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.HaVSwitchIds)) {
		request.HaVSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaVSwitchIds, tea.String("HaVSwitchIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceSpec)) {
		request.ResourceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSpec, tea.String("ResourceSpec"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ha)) {
		body["Ha"] = request.Ha
	}

	if !tea.BoolValue(util.IsUnset(request.HaResourceSpecShrink)) {
		body["HaResourceSpec"] = request.HaResourceSpecShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HaVSwitchIdsShrink)) {
		body["HaVSwitchIds"] = request.HaVSwitchIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HaZoneId)) {
		body["HaZoneId"] = request.HaZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceSpecShrink)) {
		body["ResourceSpec"] = request.ResourceSpecShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryModifyInstancePrice"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 查询付费类型为包年包月的实例修改资源规格的价格
//
// @param request - QueryModifyInstancePriceRequest
//
// @return QueryModifyInstancePriceResponse
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

// Summary:
//
// 查询付费类型为包年包月的实例续费价格
//
// @param request - QueryRenewInstancePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRenewInstancePriceResponse
func (client *Client) QueryRenewInstancePriceWithOptions(request *QueryRenewInstancePriceRequest, runtime *util.RuntimeOptions) (_result *QueryRenewInstancePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRenewInstancePrice"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 查询付费类型为包年包月的实例续费价格
//
// @param request - QueryRenewInstancePriceRequest
//
// @return QueryRenewInstancePriceResponse
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

// Summary:
//
// 续费
//
// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewInstance"),
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 续费
//
// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
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

// Summary:
//
// 打标签接口
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
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// 打标签接口
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
// flinkasi去标签
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
		Version:     tea.String("2021-10-28"),
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

// Summary:
//
// flinkasi去标签
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

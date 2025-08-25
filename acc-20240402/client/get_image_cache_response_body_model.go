// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetImageCacheResponseBody
	GetCreateTime() *string
	SetEvents(v []*GetImageCacheResponseBodyEvents) *GetImageCacheResponseBody
	GetEvents() []*GetImageCacheResponseBodyEvents
	SetImageCacheId(v string) *GetImageCacheResponseBody
	GetImageCacheId() *string
	SetImageCacheName(v string) *GetImageCacheResponseBody
	GetImageCacheName() *string
	SetImages(v []*string) *GetImageCacheResponseBody
	GetImages() []*string
	SetNetworkConfig(v *GetImageCacheResponseBodyNetworkConfig) *GetImageCacheResponseBody
	GetNetworkConfig() *GetImageCacheResponseBodyNetworkConfig
	SetPaymentType(v string) *GetImageCacheResponseBody
	GetPaymentType() *string
	SetReadyTime(v string) *GetImageCacheResponseBody
	GetReadyTime() *string
	SetRegionId(v string) *GetImageCacheResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetImageCacheResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetImageCacheResponseBody
	GetResourceGroupId() *string
	SetSize(v int32) *GetImageCacheResponseBody
	GetSize() *int32
	SetStatus(v string) *GetImageCacheResponseBody
	GetStatus() *string
	SetTags(v []*GetImageCacheResponseBodyTags) *GetImageCacheResponseBody
	GetTags() []*GetImageCacheResponseBodyTags
}

type GetImageCacheResponseBody struct {
	// example:
	//
	// 2025-**-**T07:55:25Z
	CreateTime *string                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Events     []*GetImageCacheResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// example:
	//
	// imc-bp1dj*****
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// example:
	//
	// my-imc
	ImageCacheName *string                                 `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	Images         []*string                               `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	NetworkConfig  *GetImageCacheResponseBodyNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// 2025-**-**T07:58:25Z
	ReadyTime *string `json:"ReadyTime,omitempty" xml:"ReadyTime,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 0E234675-3465-4CC3-9D0F-9A864BC*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-aekzh43v*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 8
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// Ready
	Status *string                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*GetImageCacheResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetImageCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageCacheResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetImageCacheResponseBody) GetEvents() []*GetImageCacheResponseBodyEvents {
	return s.Events
}

func (s *GetImageCacheResponseBody) GetImageCacheId() *string {
	return s.ImageCacheId
}

func (s *GetImageCacheResponseBody) GetImageCacheName() *string {
	return s.ImageCacheName
}

func (s *GetImageCacheResponseBody) GetImages() []*string {
	return s.Images
}

func (s *GetImageCacheResponseBody) GetNetworkConfig() *GetImageCacheResponseBodyNetworkConfig {
	return s.NetworkConfig
}

func (s *GetImageCacheResponseBody) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetImageCacheResponseBody) GetReadyTime() *string {
	return s.ReadyTime
}

func (s *GetImageCacheResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetImageCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageCacheResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetImageCacheResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *GetImageCacheResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetImageCacheResponseBody) GetTags() []*GetImageCacheResponseBodyTags {
	return s.Tags
}

func (s *GetImageCacheResponseBody) SetCreateTime(v string) *GetImageCacheResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetImageCacheResponseBody) SetEvents(v []*GetImageCacheResponseBodyEvents) *GetImageCacheResponseBody {
	s.Events = v
	return s
}

func (s *GetImageCacheResponseBody) SetImageCacheId(v string) *GetImageCacheResponseBody {
	s.ImageCacheId = &v
	return s
}

func (s *GetImageCacheResponseBody) SetImageCacheName(v string) *GetImageCacheResponseBody {
	s.ImageCacheName = &v
	return s
}

func (s *GetImageCacheResponseBody) SetImages(v []*string) *GetImageCacheResponseBody {
	s.Images = v
	return s
}

func (s *GetImageCacheResponseBody) SetNetworkConfig(v *GetImageCacheResponseBodyNetworkConfig) *GetImageCacheResponseBody {
	s.NetworkConfig = v
	return s
}

func (s *GetImageCacheResponseBody) SetPaymentType(v string) *GetImageCacheResponseBody {
	s.PaymentType = &v
	return s
}

func (s *GetImageCacheResponseBody) SetReadyTime(v string) *GetImageCacheResponseBody {
	s.ReadyTime = &v
	return s
}

func (s *GetImageCacheResponseBody) SetRegionId(v string) *GetImageCacheResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetImageCacheResponseBody) SetRequestId(v string) *GetImageCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageCacheResponseBody) SetResourceGroupId(v string) *GetImageCacheResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetImageCacheResponseBody) SetSize(v int32) *GetImageCacheResponseBody {
	s.Size = &v
	return s
}

func (s *GetImageCacheResponseBody) SetStatus(v string) *GetImageCacheResponseBody {
	s.Status = &v
	return s
}

func (s *GetImageCacheResponseBody) SetTags(v []*GetImageCacheResponseBodyTags) *GetImageCacheResponseBody {
	s.Tags = v
	return s
}

func (s *GetImageCacheResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImageCacheResponseBodyEvents struct {
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 2025-**-**T02:24:48Z
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	// example:
	//
	// 2025-**-**T02:24:48Z
	LastTimestamp *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// example:
	//
	// Image cache [my-imc] has been created successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// imagetest.1661f31f851*****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ImageCacheCreated
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetImageCacheResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *GetImageCacheResponseBodyEvents) GetCount() *int32 {
	return s.Count
}

func (s *GetImageCacheResponseBodyEvents) GetFirstTimestamp() *string {
	return s.FirstTimestamp
}

func (s *GetImageCacheResponseBodyEvents) GetLastTimestamp() *string {
	return s.LastTimestamp
}

func (s *GetImageCacheResponseBodyEvents) GetMessage() *string {
	return s.Message
}

func (s *GetImageCacheResponseBodyEvents) GetName() *string {
	return s.Name
}

func (s *GetImageCacheResponseBodyEvents) GetReason() *string {
	return s.Reason
}

func (s *GetImageCacheResponseBodyEvents) GetType() *string {
	return s.Type
}

func (s *GetImageCacheResponseBodyEvents) SetCount(v int32) *GetImageCacheResponseBodyEvents {
	s.Count = &v
	return s
}

func (s *GetImageCacheResponseBodyEvents) SetFirstTimestamp(v string) *GetImageCacheResponseBodyEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *GetImageCacheResponseBodyEvents) SetLastTimestamp(v string) *GetImageCacheResponseBodyEvents {
	s.LastTimestamp = &v
	return s
}

func (s *GetImageCacheResponseBodyEvents) SetMessage(v string) *GetImageCacheResponseBodyEvents {
	s.Message = &v
	return s
}

func (s *GetImageCacheResponseBodyEvents) SetName(v string) *GetImageCacheResponseBodyEvents {
	s.Name = &v
	return s
}

func (s *GetImageCacheResponseBodyEvents) SetReason(v string) *GetImageCacheResponseBodyEvents {
	s.Reason = &v
	return s
}

func (s *GetImageCacheResponseBodyEvents) SetType(v string) *GetImageCacheResponseBodyEvents {
	s.Type = &v
	return s
}

func (s *GetImageCacheResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}

type GetImageCacheResponseBodyNetworkConfig struct {
	EipInstance *GetImageCacheResponseBodyNetworkConfigEipInstance `json:"EipInstance,omitempty" xml:"EipInstance,omitempty" type:"Struct"`
	// example:
	//
	// sg-0jlgektkddwa42n*****
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchIds      []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s GetImageCacheResponseBodyNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheResponseBodyNetworkConfig) GoString() string {
	return s.String()
}

func (s *GetImageCacheResponseBodyNetworkConfig) GetEipInstance() *GetImageCacheResponseBodyNetworkConfigEipInstance {
	return s.EipInstance
}

func (s *GetImageCacheResponseBodyNetworkConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetImageCacheResponseBodyNetworkConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetImageCacheResponseBodyNetworkConfig) SetEipInstance(v *GetImageCacheResponseBodyNetworkConfigEipInstance) *GetImageCacheResponseBodyNetworkConfig {
	s.EipInstance = v
	return s
}

func (s *GetImageCacheResponseBodyNetworkConfig) SetSecurityGroupId(v string) *GetImageCacheResponseBodyNetworkConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *GetImageCacheResponseBodyNetworkConfig) SetVSwitchIds(v []*string) *GetImageCacheResponseBodyNetworkConfig {
	s.VSwitchIds = v
	return s
}

func (s *GetImageCacheResponseBodyNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type GetImageCacheResponseBodyNetworkConfigEipInstance struct {
	// example:
	//
	// true
	AutoCreate *bool `json:"AutoCreate,omitempty" xml:"AutoCreate,omitempty"`
	// example:
	//
	// 100
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// example:
	//
	// eip-0jl0bx3fnpnjc9i4*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetImageCacheResponseBodyNetworkConfigEipInstance) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheResponseBodyNetworkConfigEipInstance) GoString() string {
	return s.String()
}

func (s *GetImageCacheResponseBodyNetworkConfigEipInstance) GetAutoCreate() *bool {
	return s.AutoCreate
}

func (s *GetImageCacheResponseBodyNetworkConfigEipInstance) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *GetImageCacheResponseBodyNetworkConfigEipInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetImageCacheResponseBodyNetworkConfigEipInstance) SetAutoCreate(v bool) *GetImageCacheResponseBodyNetworkConfigEipInstance {
	s.AutoCreate = &v
	return s
}

func (s *GetImageCacheResponseBodyNetworkConfigEipInstance) SetBandwidth(v int32) *GetImageCacheResponseBodyNetworkConfigEipInstance {
	s.Bandwidth = &v
	return s
}

func (s *GetImageCacheResponseBodyNetworkConfigEipInstance) SetInstanceId(v string) *GetImageCacheResponseBodyNetworkConfigEipInstance {
	s.InstanceId = &v
	return s
}

func (s *GetImageCacheResponseBodyNetworkConfigEipInstance) Validate() error {
	return dara.Validate(s)
}

type GetImageCacheResponseBodyTags struct {
	// example:
	//
	// imc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetImageCacheResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetImageCacheResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetImageCacheResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetImageCacheResponseBodyTags) SetKey(v string) *GetImageCacheResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetImageCacheResponseBodyTags) SetValue(v string) *GetImageCacheResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetImageCacheResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

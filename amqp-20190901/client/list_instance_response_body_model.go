// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListInstanceResponseBody
	GetCode() *int32
	SetData(v *ListInstanceResponseBodyData) *ListInstanceResponseBody
	GetData() *ListInstanceResponseBodyData
	SetMessage(v string) *ListInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceResponseBody
	GetSuccess() *bool
}

type ListInstanceResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListInstanceResponseBody) GetData() *ListInstanceResponseBodyData {
	return s.Data
}

func (s *ListInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceResponseBody) SetCode(v int32) *ListInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceResponseBody) SetData(v *ListInstanceResponseBodyData) *ListInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceResponseBody) SetMessage(v string) *ListInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetSuccess(v bool) *ListInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceResponseBodyData struct {
	Instances []*ListInstanceResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s ListInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyData) GetInstances() []*ListInstanceResponseBodyDataInstances {
	return s.Instances
}

func (s *ListInstanceResponseBodyData) SetInstances(v []*ListInstanceResponseBodyDataInstances) *ListInstanceResponseBodyData {
	s.Instances = v
	return s
}

func (s *ListInstanceResponseBodyData) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceResponseBodyDataInstances struct {
	AutoRenew       *bool                                      `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClassicEndpoint *string                                    `json:"ClassicEndpoint,omitempty" xml:"ClassicEndpoint,omitempty"`
	Expire          *int64                                     `json:"Expire,omitempty" xml:"Expire,omitempty"`
	InstanceId      *string                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string                                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType    *string                                    `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaxEIPTPS       *int32                                     `json:"MaxEIPTPS,omitempty" xml:"MaxEIPTPS,omitempty"`
	MaxQueue        *int32                                     `json:"MaxQueue,omitempty" xml:"MaxQueue,omitempty"`
	MaxTPS          *int32                                     `json:"MaxTPS,omitempty" xml:"MaxTPS,omitempty"`
	MaxVhost        *int32                                     `json:"MaxVhost,omitempty" xml:"MaxVhost,omitempty"`
	OrderCreate     *int64                                     `json:"OrderCreate,omitempty" xml:"OrderCreate,omitempty"`
	OrderType       *string                                    `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PrivateEndpoint *string                                    `json:"PrivateEndpoint,omitempty" xml:"PrivateEndpoint,omitempty"`
	PublicEndpoint  *string                                    `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
	Status          *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageSize     *int32                                     `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	SupportEIP      *bool                                      `json:"SupportEIP,omitempty" xml:"SupportEIP,omitempty"`
	Tags            *ListInstanceResponseBodyDataInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s ListInstanceResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyDataInstances) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ListInstanceResponseBodyDataInstances) GetClassicEndpoint() *string {
	return s.ClassicEndpoint
}

func (s *ListInstanceResponseBodyDataInstances) GetExpire() *int64 {
	return s.Expire
}

func (s *ListInstanceResponseBodyDataInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceResponseBodyDataInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstanceResponseBodyDataInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListInstanceResponseBodyDataInstances) GetMaxEIPTPS() *int32 {
	return s.MaxEIPTPS
}

func (s *ListInstanceResponseBodyDataInstances) GetMaxQueue() *int32 {
	return s.MaxQueue
}

func (s *ListInstanceResponseBodyDataInstances) GetMaxTPS() *int32 {
	return s.MaxTPS
}

func (s *ListInstanceResponseBodyDataInstances) GetMaxVhost() *int32 {
	return s.MaxVhost
}

func (s *ListInstanceResponseBodyDataInstances) GetOrderCreate() *int64 {
	return s.OrderCreate
}

func (s *ListInstanceResponseBodyDataInstances) GetOrderType() *string {
	return s.OrderType
}

func (s *ListInstanceResponseBodyDataInstances) GetPrivateEndpoint() *string {
	return s.PrivateEndpoint
}

func (s *ListInstanceResponseBodyDataInstances) GetPublicEndpoint() *string {
	return s.PublicEndpoint
}

func (s *ListInstanceResponseBodyDataInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceResponseBodyDataInstances) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *ListInstanceResponseBodyDataInstances) GetSupportEIP() *bool {
	return s.SupportEIP
}

func (s *ListInstanceResponseBodyDataInstances) GetTags() *ListInstanceResponseBodyDataInstancesTags {
	return s.Tags
}

func (s *ListInstanceResponseBodyDataInstances) SetAutoRenew(v bool) *ListInstanceResponseBodyDataInstances {
	s.AutoRenew = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetClassicEndpoint(v string) *ListInstanceResponseBodyDataInstances {
	s.ClassicEndpoint = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetExpire(v int64) *ListInstanceResponseBodyDataInstances {
	s.Expire = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetInstanceId(v string) *ListInstanceResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetInstanceName(v string) *ListInstanceResponseBodyDataInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetInstanceType(v string) *ListInstanceResponseBodyDataInstances {
	s.InstanceType = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetMaxEIPTPS(v int32) *ListInstanceResponseBodyDataInstances {
	s.MaxEIPTPS = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetMaxQueue(v int32) *ListInstanceResponseBodyDataInstances {
	s.MaxQueue = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetMaxTPS(v int32) *ListInstanceResponseBodyDataInstances {
	s.MaxTPS = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetMaxVhost(v int32) *ListInstanceResponseBodyDataInstances {
	s.MaxVhost = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetOrderCreate(v int64) *ListInstanceResponseBodyDataInstances {
	s.OrderCreate = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetOrderType(v string) *ListInstanceResponseBodyDataInstances {
	s.OrderType = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetPrivateEndpoint(v string) *ListInstanceResponseBodyDataInstances {
	s.PrivateEndpoint = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetPublicEndpoint(v string) *ListInstanceResponseBodyDataInstances {
	s.PublicEndpoint = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetStatus(v string) *ListInstanceResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetStorageSize(v int32) *ListInstanceResponseBodyDataInstances {
	s.StorageSize = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetSupportEIP(v bool) *ListInstanceResponseBodyDataInstances {
	s.SupportEIP = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) SetTags(v *ListInstanceResponseBodyDataInstancesTags) *ListInstanceResponseBodyDataInstances {
	s.Tags = v
	return s
}

func (s *ListInstanceResponseBodyDataInstances) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceResponseBodyDataInstancesTags struct {
	Tags []*ListInstanceResponseBodyDataInstancesTagsTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListInstanceResponseBodyDataInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyDataInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyDataInstancesTags) GetTags() []*ListInstanceResponseBodyDataInstancesTagsTags {
	return s.Tags
}

func (s *ListInstanceResponseBodyDataInstancesTags) SetTags(v []*ListInstanceResponseBodyDataInstancesTagsTags) *ListInstanceResponseBodyDataInstancesTags {
	s.Tags = v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesTags) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceResponseBodyDataInstancesTagsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstanceResponseBodyDataInstancesTagsTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResponseBodyDataInstancesTagsTags) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyDataInstancesTagsTags) GetKey() *string {
	return s.Key
}

func (s *ListInstanceResponseBodyDataInstancesTagsTags) GetValue() *string {
	return s.Value
}

func (s *ListInstanceResponseBodyDataInstancesTagsTags) SetKey(v string) *ListInstanceResponseBodyDataInstancesTagsTags {
	s.Key = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesTagsTags) SetValue(v string) *ListInstanceResponseBodyDataInstancesTagsTags {
	s.Value = &v
	return s
}

func (s *ListInstanceResponseBodyDataInstancesTagsTags) Validate() error {
	return dara.Validate(s)
}

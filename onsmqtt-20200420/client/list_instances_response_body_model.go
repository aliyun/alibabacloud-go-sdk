// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesResponseBody
	GetCode() *string
	SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody
	GetInstances() []*ListInstancesResponseBodyInstances
	SetMessage(v string) *ListInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListInstancesResponseBody
	GetSuccess() *string
}

type ListInstancesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of instances.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 294F5B26-FADF-41D0-BE68-14ABE4EB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesResponseBody) GetInstances() []*ListInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
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

func (s *ListInstancesResponseBody) SetSuccess(v string) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
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

type ListInstancesResponseBodyInstances struct {
	// The time when the instance was created.
	//
	// example:
	//
	// 1739345768000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the instance.
	//
	// example:
	//
	// 1741795200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// mqtt-cn-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// mqtt-cn-xxxxxxx
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance status.
	//
	// - 0: Preparing
	//
	// - 5: In Service
	//
	// - 7: Upgrading
	//
	// - 15: Frozen
	//
	// - 10: Released
	//
	// - 1: Not Activated
	//
	// - 2: Overdue
	//
	// - 3: Expired
	//
	// example:
	//
	// 5
	InstanceStatus *int32 `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance type:
	//
	// - 0: Basic Edition
	//
	// - 1: Platinum Edition
	//
	// - 2: Pay-as-you-go
	//
	// - 3: Subscription Lightweight
	//
	// - 4: Professional Edition
	//
	// - 5: Serverless Edition
	//
	// example:
	//
	// 4
	InstanceType *int32 `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The Milvus version.
	//
	// example:
	//
	// V3.3.0
	KernelVersion *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	// The list of tags.
	MqttTags []*ListInstancesResponseBodyInstancesMqttTags `json:"MqttTags,omitempty" xml:"MqttTags,omitempty" type:"Repeated"`
	// The order ID.
	//
	// example:
	//
	// 245949616640095
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance type of the Professional Edition instance.
	//
	// example:
	//
	// mqtt.p2.1xlarge
	Specific *string `json:"Specific,omitempty" xml:"Specific,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListInstancesResponseBodyInstances) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesResponseBodyInstances) GetInstanceStatus() *int32 {
	return s.InstanceStatus
}

func (s *ListInstancesResponseBodyInstances) GetInstanceType() *int32 {
	return s.InstanceType
}

func (s *ListInstancesResponseBodyInstances) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *ListInstancesResponseBodyInstances) GetMqttTags() []*ListInstancesResponseBodyInstancesMqttTags {
	return s.MqttTags
}

func (s *ListInstancesResponseBodyInstances) GetOrderId() *string {
	return s.OrderId
}

func (s *ListInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesResponseBodyInstances) GetSpecific() *string {
	return s.Specific
}

func (s *ListInstancesResponseBodyInstances) SetCreateTime(v int64) *ListInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetExpireTime(v int64) *ListInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceName(v string) *ListInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceStatus(v int32) *ListInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceType(v int32) *ListInstancesResponseBodyInstances {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetKernelVersion(v string) *ListInstancesResponseBodyInstances {
	s.KernelVersion = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetMqttTags(v []*ListInstancesResponseBodyInstancesMqttTags) *ListInstancesResponseBodyInstances {
	s.MqttTags = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetOrderId(v string) *ListInstancesResponseBodyInstances {
	s.OrderId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRegionId(v string) *ListInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetSpecific(v string) *ListInstancesResponseBodyInstances {
	s.Specific = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
	if s.MqttTags != nil {
		for _, item := range s.MqttTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstancesMqttTags struct {
	// The tag key.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListInstancesResponseBodyInstancesMqttTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesMqttTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesMqttTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListInstancesResponseBodyInstancesMqttTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListInstancesResponseBodyInstancesMqttTags) SetTagKey(v string) *ListInstancesResponseBodyInstancesMqttTags {
	s.TagKey = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesMqttTags) SetTagValue(v string) *ListInstancesResponseBodyInstancesMqttTags {
	s.TagValue = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesMqttTags) Validate() error {
	return dara.Validate(s)
}

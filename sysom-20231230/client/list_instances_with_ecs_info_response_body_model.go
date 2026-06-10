// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesWithEcsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesWithEcsInfoResponseBody
	GetCode() *string
	SetData(v []*ListInstancesWithEcsInfoResponseBodyData) *ListInstancesWithEcsInfoResponseBody
	GetData() []*ListInstancesWithEcsInfoResponseBodyData
	SetMessage(v string) *ListInstancesWithEcsInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstancesWithEcsInfoResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListInstancesWithEcsInfoResponseBody
	GetTotal() *int64
}

type ListInstancesWithEcsInfoResponseBody struct {
	// Status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data.
	Data []*ListInstancesWithEcsInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Error message. An empty value indicates that the data has been read completely.
	//
	// example:
	//
	// result: code=1 msg=(Request failed, status_code != 200)
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request RequestId
	//
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Total number of records
	//
	// example:
	//
	// 319
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInstancesWithEcsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesWithEcsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesWithEcsInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesWithEcsInfoResponseBody) GetData() []*ListInstancesWithEcsInfoResponseBodyData {
	return s.Data
}

func (s *ListInstancesWithEcsInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesWithEcsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesWithEcsInfoResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListInstancesWithEcsInfoResponseBody) SetCode(v string) *ListInstancesWithEcsInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBody) SetData(v []*ListInstancesWithEcsInfoResponseBodyData) *ListInstancesWithEcsInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBody) SetMessage(v string) *ListInstancesWithEcsInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBody) SetRequestId(v string) *ListInstancesWithEcsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBody) SetTotal(v int64) *ListInstancesWithEcsInfoResponseBody {
	s.Total = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesWithEcsInfoResponseBodyData struct {
	// Cluster ID
	//
	// example:
	//
	// cbf7a37bc905d4682a3338b3744810269
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// instance ID
	//
	// example:
	//
	// i-bp118piqcio9tiwgh84b
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// Instance Name.
	//
	// example:
	//
	// allowed-repos-r2tzl
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	// tags of instances
	InstanceTag []*ListInstancesWithEcsInfoResponseBodyDataInstanceTag `json:"instance_tag,omitempty" xml:"instance_tag,omitempty" type:"Repeated"`
	// Milvus version
	//
	// example:
	//
	// 5.10.134-14.an8.x86_64
	KernelVersion *string `json:"kernel_version,omitempty" xml:"kernel_version,omitempty"`
	// ECS instance architecture
	//
	// example:
	//
	// x86
	OsArch *string `json:"os_arch,omitempty" xml:"os_arch,omitempty"`
	// Instance health score
	//
	// example:
	//
	// 100
	OsHealthScore *string `json:"os_health_score,omitempty" xml:"os_health_score,omitempty"`
	// The operating system name of the instance
	//
	// example:
	//
	// Alibaba Cloud Linux  3.2104 LTS 64bit
	OsName *string `json:"os_name,omitempty" xml:"os_name,omitempty"`
	// Instance private IP
	//
	// example:
	//
	// 1.1.1.1
	PrivateIp *string `json:"private_ip,omitempty" xml:"private_ip,omitempty"`
	// Instance Internet IP
	//
	// example:
	//
	// 1.1.1.1
	PublicIp *string `json:"public_ip,omitempty" xml:"public_ip,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-xxxxxx
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// Resource group name
	//
	// example:
	//
	// default resource group
	ResourceGroupName *string `json:"resource_group_name,omitempty" xml:"resource_group_name,omitempty"`
	// The running status of the instance. Valid values:
	//
	// - **Running**: The instance is running.
	//
	// - **Offline**: The instance is offline.
	//
	// > An instance in the Offline state indicates that the heartbeat from the edge zone to the SysOM Server has been lost. This does not mean that the corresponding ECS instance is not running.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstancesWithEcsInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesWithEcsInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetInstanceTag() []*ListInstancesWithEcsInfoResponseBodyDataInstanceTag {
	return s.InstanceTag
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetOsArch() *string {
	return s.OsArch
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetOsHealthScore() *string {
	return s.OsHealthScore
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetOsName() *string {
	return s.OsName
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListInstancesWithEcsInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetClusterId(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetInstanceId(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetInstanceName(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetInstanceTag(v []*ListInstancesWithEcsInfoResponseBodyDataInstanceTag) *ListInstancesWithEcsInfoResponseBodyData {
	s.InstanceTag = v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetKernelVersion(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.KernelVersion = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetOsArch(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.OsArch = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetOsHealthScore(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.OsHealthScore = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetOsName(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.OsName = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetPrivateIp(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.PrivateIp = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetPublicIp(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.PublicIp = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetResourceGroupId(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetResourceGroupName(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.ResourceGroupName = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) SetStatus(v string) *ListInstancesWithEcsInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyData) Validate() error {
	if s.InstanceTag != nil {
		for _, item := range s.InstanceTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesWithEcsInfoResponseBodyDataInstanceTag struct {
	// Name of the tag.
	//
	// example:
	//
	// test_tag_key
	TagKey *string `json:"tag_key,omitempty" xml:"tag_key,omitempty"`
	// Tag value.
	//
	// example:
	//
	// test_tag_value
	TagValue *string `json:"tag_value,omitempty" xml:"tag_value,omitempty"`
}

func (s ListInstancesWithEcsInfoResponseBodyDataInstanceTag) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesWithEcsInfoResponseBodyDataInstanceTag) GoString() string {
	return s.String()
}

func (s *ListInstancesWithEcsInfoResponseBodyDataInstanceTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListInstancesWithEcsInfoResponseBodyDataInstanceTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListInstancesWithEcsInfoResponseBodyDataInstanceTag) SetTagKey(v string) *ListInstancesWithEcsInfoResponseBodyDataInstanceTag {
	s.TagKey = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyDataInstanceTag) SetTagValue(v string) *ListInstancesWithEcsInfoResponseBodyDataInstanceTag {
	s.TagValue = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponseBodyDataInstanceTag) Validate() error {
	return dara.Validate(s)
}

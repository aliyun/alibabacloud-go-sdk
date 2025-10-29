// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstashResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListLogstashResponseBodyHeaders) *ListLogstashResponseBody
	GetHeaders() *ListLogstashResponseBodyHeaders
	SetRequestId(v string) *ListLogstashResponseBody
	GetRequestId() *string
	SetResult(v []*ListLogstashResponseBodyResult) *ListLogstashResponseBody
	GetResult() []*ListLogstashResponseBodyResult
}

type ListLogstashResponseBody struct {
	// The billing method of the instance. Supported: prepaid (subscription) and postpaid (pay-as-you-go).
	Headers *ListLogstashResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// Detailed information about the matching instances.
	//
	// example:
	//
	// AC442F2F-5068-4434-AA21-E78947A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the VPC.
	Result []*ListLogstashResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListLogstashResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBody) GetHeaders() *ListLogstashResponseBodyHeaders {
	return s.Headers
}

func (s *ListLogstashResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogstashResponseBody) GetResult() []*ListLogstashResponseBodyResult {
	return s.Result
}

func (s *ListLogstashResponseBody) SetHeaders(v *ListLogstashResponseBodyHeaders) *ListLogstashResponseBody {
	s.Headers = v
	return s
}

func (s *ListLogstashResponseBody) SetRequestId(v string) *ListLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogstashResponseBody) SetResult(v []*ListLogstashResponseBodyResult) *ListLogstashResponseBody {
	s.Result = v
	return s
}

func (s *ListLogstashResponseBody) Validate() error {
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLogstashResponseBodyHeaders struct {
	// The number of data nodes.
	//
	// example:
	//
	// 10
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListLogstashResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListLogstashResponseBodyHeaders) SetXTotalCount(v int32) *ListLogstashResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListLogstashResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListLogstashResponseBodyResult struct {
	// The configuration information of the data node.
	Tags []*ListLogstashResponseBodyResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// example:
	//
	// 2018-07-13T03:58:07.253Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// ls-cn-abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The tag value of the cloud disk.
	//
	// example:
	//
	// ls-cn-n6w1o5jq****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The network type. Currently, only Virtual Private Cloud (VPC) is supported.
	NetworkConfig *ListLogstashResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	// The state of the instance. Valid values: Normal, Active, Inactive, and Invalid.
	//
	// example:
	//
	// 2
	NodeAmount *int32 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// Specifies whether to use disk encryption. Valid values:
	//
	// 	- true: Enables the concurrent query feature for queries other than aggregate queries.
	//
	// 	- false: Disables the concurrent query feature for queries other than aggregate queries.
	NodeSpec *ListLogstashResponseBodyResultNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// The time when the instance was created.
	//
	// example:
	//
	// postpaid
	PaymentType     *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The version of the instance. Currently, only 6.7.0_with_X-Pack and 7.4.0_with_X-Pack are supported.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the instance. Valid values:
	//
	// example:
	//
	// 2018-07-18T10:10:04.484Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The tag key of the cloud disk.
	//
	// example:
	//
	// 6.7.0_with_X-Pack
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListLogstashResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBodyResult) GetTags() []*ListLogstashResponseBodyResultTags {
	return s.Tags
}

func (s *ListLogstashResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListLogstashResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListLogstashResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLogstashResponseBodyResult) GetNetworkConfig() *ListLogstashResponseBodyResultNetworkConfig {
	return s.NetworkConfig
}

func (s *ListLogstashResponseBodyResult) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *ListLogstashResponseBodyResult) GetNodeSpec() *ListLogstashResponseBodyResultNodeSpec {
	return s.NodeSpec
}

func (s *ListLogstashResponseBodyResult) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListLogstashResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLogstashResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListLogstashResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListLogstashResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *ListLogstashResponseBodyResult) SetTags(v []*ListLogstashResponseBodyResultTags) *ListLogstashResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListLogstashResponseBodyResult) SetCreatedAt(v string) *ListLogstashResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetDescription(v string) *ListLogstashResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetInstanceId(v string) *ListLogstashResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetNetworkConfig(v *ListLogstashResponseBodyResultNetworkConfig) *ListLogstashResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *ListLogstashResponseBodyResult) SetNodeAmount(v int32) *ListLogstashResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetNodeSpec(v *ListLogstashResponseBodyResultNodeSpec) *ListLogstashResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *ListLogstashResponseBodyResult) SetPaymentType(v string) *ListLogstashResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetResourceGroupId(v string) *ListLogstashResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetStatus(v string) *ListLogstashResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetUpdatedAt(v string) *ListLogstashResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListLogstashResponseBodyResult) SetVersion(v string) *ListLogstashResponseBodyResult {
	s.Version = &v
	return s
}

func (s *ListLogstashResponseBodyResult) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkConfig != nil {
		if err := s.NetworkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NodeSpec != nil {
		if err := s.NodeSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLogstashResponseBodyResultTags struct {
	// The disk size of the node.
	//
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The instance type of the ECS instance.
	//
	// example:
	//
	// dev
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListLogstashResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBodyResultTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListLogstashResponseBodyResultTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListLogstashResponseBodyResultTags) SetTagKey(v string) *ListLogstashResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *ListLogstashResponseBodyResultTags) SetTagValue(v string) *ListLogstashResponseBodyResultTags {
	s.TagValue = &v
	return s
}

func (s *ListLogstashResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}

type ListLogstashResponseBodyResultNetworkConfig struct {
	// example:
	//
	// vpc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vpc-abc
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-*
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// example:
	//
	// vsw-def
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s ListLogstashResponseBodyResultNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBodyResultNetworkConfig) GetType() *string {
	return s.Type
}

func (s *ListLogstashResponseBodyResultNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *ListLogstashResponseBodyResultNetworkConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *ListLogstashResponseBodyResultNetworkConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListLogstashResponseBodyResultNetworkConfig) SetType(v string) *ListLogstashResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *ListLogstashResponseBodyResultNetworkConfig) SetVpcId(v string) *ListLogstashResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *ListLogstashResponseBodyResultNetworkConfig) SetVsArea(v string) *ListLogstashResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *ListLogstashResponseBodyResultNetworkConfig) SetVswitchId(v string) *ListLogstashResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *ListLogstashResponseBodyResultNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type ListLogstashResponseBodyResultNodeSpec struct {
	// The network configurations.
	//
	// example:
	//
	// 50
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// false
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// The zone where the cluster resides.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The type of the disk.
	//
	// example:
	//
	// logstash.n4.small
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListLogstashResponseBodyResultNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *ListLogstashResponseBodyResultNodeSpec) GetDisk() *int32 {
	return s.Disk
}

func (s *ListLogstashResponseBodyResultNodeSpec) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *ListLogstashResponseBodyResultNodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *ListLogstashResponseBodyResultNodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *ListLogstashResponseBodyResultNodeSpec) SetDisk(v int32) *ListLogstashResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *ListLogstashResponseBodyResultNodeSpec) SetDiskEncryption(v bool) *ListLogstashResponseBodyResultNodeSpec {
	s.DiskEncryption = &v
	return s
}

func (s *ListLogstashResponseBodyResultNodeSpec) SetDiskType(v string) *ListLogstashResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *ListLogstashResponseBodyResultNodeSpec) SetSpec(v string) *ListLogstashResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

func (s *ListLogstashResponseBodyResultNodeSpec) Validate() error {
	return dara.Validate(s)
}

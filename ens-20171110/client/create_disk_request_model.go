// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateDiskRequest
	GetCategory() *string
	SetDiskName(v string) *CreateDiskRequest
	GetDiskName() *string
	SetEncrypted(v bool) *CreateDiskRequest
	GetEncrypted() *bool
	SetEnsRegionId(v string) *CreateDiskRequest
	GetEnsRegionId() *string
	SetInstanceBillingCycle(v string) *CreateDiskRequest
	GetInstanceBillingCycle() *string
	SetInstanceChargeType(v string) *CreateDiskRequest
	GetInstanceChargeType() *string
	SetKMSKeyId(v string) *CreateDiskRequest
	GetKMSKeyId() *string
	SetSize(v string) *CreateDiskRequest
	GetSize() *string
	SetSnapshotId(v string) *CreateDiskRequest
	GetSnapshotId() *string
	SetTag(v []*CreateDiskRequestTag) *CreateDiskRequest
	GetTag() []*CreateDiskRequestTag
}

type CreateDiskRequest struct {
	// The category of the disk. Valid values:
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: all-flash disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The name of the disk.
	//
	// example:
	//
	// yourDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the new system disk. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default): no
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the edge node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu-telecom
	EnsRegionId          *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The billing method of the instance. Set the value to **PostPaid**.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used by the cloud disk.
	//
	// >  If you set the **Encrypted*	- parameter to **true**, the default service key is used when the **KMSKeyId*	- parameter is empty.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fxxxxx
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The size of the disk. Unit: GiB.
	//
	// example:
	//
	// 20
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot that you want to use to create the disk.
	//
	// The following limits apply to the **SnapshotId*	- and **Size*	- parameters:
	//
	// 	- If the size of the snapshot specified by **SnapshotId*	- is greater than the specified **Size*	- value, the size of the created disk is equal to the specified snapshot size.
	//
	// 	- If the size of the snapshot specified by **SnapshotId*	- is smaller than the specified **Size*	- value, the size of the created disk is equal to the specified **Size*	- value.
	//
	// example:
	//
	// s-897654321****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The tags of the instance. You can specify at most 20 tags in each call.
	Tag []*CreateDiskRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateDiskRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateDiskRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateDiskRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateDiskRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *CreateDiskRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateDiskRequest) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateDiskRequest) GetSize() *string {
	return s.Size
}

func (s *CreateDiskRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateDiskRequest) GetTag() []*CreateDiskRequestTag {
	return s.Tag
}

func (s *CreateDiskRequest) SetCategory(v string) *CreateDiskRequest {
	s.Category = &v
	return s
}

func (s *CreateDiskRequest) SetDiskName(v string) *CreateDiskRequest {
	s.DiskName = &v
	return s
}

func (s *CreateDiskRequest) SetEncrypted(v bool) *CreateDiskRequest {
	s.Encrypted = &v
	return s
}

func (s *CreateDiskRequest) SetEnsRegionId(v string) *CreateDiskRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateDiskRequest) SetInstanceBillingCycle(v string) *CreateDiskRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *CreateDiskRequest) SetInstanceChargeType(v string) *CreateDiskRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateDiskRequest) SetKMSKeyId(v string) *CreateDiskRequest {
	s.KMSKeyId = &v
	return s
}

func (s *CreateDiskRequest) SetSize(v string) *CreateDiskRequest {
	s.Size = &v
	return s
}

func (s *CreateDiskRequest) SetSnapshotId(v string) *CreateDiskRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateDiskRequest) SetTag(v []*CreateDiskRequestTag) *CreateDiskRequest {
	s.Tag = v
	return s
}

func (s *CreateDiskRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDiskRequestTag struct {
	// The key of the tag. Valid values of N: **1*	- to **20**.
	//
	// 	- The key cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// 	- The key must be up to 64 characters in length.
	//
	// 	- The tag key cannot be an empty string.
	//
	// example:
	//
	// pro
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of a tag that is attached to the topics you want to query. This parameter is not required. If you configure this parameter, you must also configure the **Key*	- parameter.***	- If you include the Key and Value parameters in a request, this operation queries only the topics that use the specified tags. If you do not include these parameters in a request, this operation queries all topics that you can access.
	//
	// 	- Valid values: 1 to 20.
	//
	// 	- The value of this parameter can be an empty string.
	//
	// 	- The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// tagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDiskRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDiskRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDiskRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDiskRequestTag) SetKey(v string) *CreateDiskRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDiskRequestTag) SetValue(v string) *CreateDiskRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDiskRequestTag) Validate() error {
	return dara.Validate(s)
}

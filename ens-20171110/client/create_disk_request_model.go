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
	// The type of the cloud disk. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The name of the cloud disk.
	//
	// example:
	//
	// yourDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the cloud disk. Valid values:
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The edge zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The billing cycle of the instance. Valid values: Hour, Day, Month.
	//
	// example:
	//
	// Hour
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The billing method of the instance. Currently, only **PostPaid*	- (pay-as-you-go) is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The ID of the KMS key used by the cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fxxxxx
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The size of the cloud disk. Unit: GiB.
	//
	// example:
	//
	// 20
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot used to create the cloud disk.
	//
	// example:
	//
	// s-897654321****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The list of tags. You can add up to 20 tags in a single call.
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
	// The tag key of the instance. Valid values of N: **1*	- to **20**.
	//
	// example:
	//
	// pro
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the topic. **Key*	- and **Value*	- must be specified or left empty at the same time. If both are specified, topics are filtered by the tag. If neither is specified, all topics are queried.
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

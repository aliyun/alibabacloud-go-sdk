// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSnapshotRequest
	GetClientToken() *string
	SetDiskId(v string) *CreateSnapshotRequest
	GetDiskId() *string
	SetRegionId(v string) *CreateSnapshotRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateSnapshotRequest
	GetResourceGroupId() *string
	SetSnapshotName(v string) *CreateSnapshotRequest
	GetSnapshotName() *string
	SetTag(v []*CreateSnapshotRequestTag) *CreateSnapshotRequest
	GetTag() []*CreateSnapshotRequestTag
}

type CreateSnapshotRequest struct {
	// Ensures the idempotence of the request. You can generate the value from your client to ensure that the value is unique across different requests. **ClientToken*	- only supports ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp18kjxg9ebrhsgi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the Simple Application Server instance to which the specified disk belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the snapshot. The name must be 2 to 50 characters in length and must start with a letter or a Chinese character. It cannot start with `http://` or `https://`. It can contain Chinese characters, English letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-SnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The list of tags. You can specify up to 20 tags.
	Tag []*CreateSnapshotRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSnapshotRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSnapshotRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSnapshotRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *CreateSnapshotRequest) GetTag() []*CreateSnapshotRequestTag {
	return s.Tag
}

func (s *CreateSnapshotRequest) SetClientToken(v string) *CreateSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSnapshotRequest) SetDiskId(v string) *CreateSnapshotRequest {
	s.DiskId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRegionId(v string) *CreateSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetResourceGroupId(v string) *CreateSnapshotRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) SetTag(v []*CreateSnapshotRequestTag) *CreateSnapshotRequest {
	s.Tag = v
	return s
}

func (s *CreateSnapshotRequest) Validate() error {
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

type CreateSnapshotRequestTag struct {
	// The tag key specified for the snapshot. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string once specified. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value specified for the snapshot. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. It can be up to 64 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSnapshotRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequestTag) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateSnapshotRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateSnapshotRequestTag) SetKey(v string) *CreateSnapshotRequestTag {
	s.Key = &v
	return s
}

func (s *CreateSnapshotRequestTag) SetValue(v string) *CreateSnapshotRequestTag {
	s.Value = &v
	return s
}

func (s *CreateSnapshotRequestTag) Validate() error {
	return dara.Validate(s)
}

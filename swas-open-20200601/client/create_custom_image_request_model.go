// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCustomImageRequest
	GetClientToken() *string
	SetDataSnapshotId(v string) *CreateCustomImageRequest
	GetDataSnapshotId() *string
	SetDescription(v string) *CreateCustomImageRequest
	GetDescription() *string
	SetImageName(v string) *CreateCustomImageRequest
	GetImageName() *string
	SetInstanceId(v string) *CreateCustomImageRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateCustomImageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateCustomImageRequest
	GetResourceGroupId() *string
	SetSystemSnapshotId(v string) *CreateCustomImageRequest
	GetSystemSnapshotId() *string
	SetTag(v []*CreateCustomImageRequestTag) *CreateCustomImageRequest
	GetTag() []*CreateCustomImageRequestTag
}

type CreateCustomImageRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data disk snapshot.
	//
	// example:
	//
	// s-acscasca****
	DataSnapshotId *string `json:"DataSnapshotId,omitempty" xml:"DataSnapshotId,omitempty"`
	// The description of the image.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the custom image. The name must be 2 to 128 characters in length and must start with a letter, a Chinese character, or a digit. It can contain digits, colons (:), underscores (_), and hyphens (-). Default value: empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// customImage-test
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The instance ID of the specified Simple Application Server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query available regions.
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
	// The ID of the system disk snapshot.
	//
	// example:
	//
	// s-acscasca****
	SystemSnapshotId *string `json:"SystemSnapshotId,omitempty" xml:"SystemSnapshotId,omitempty"`
	// The list of tags. You can specify up to 20 tags.
	Tag []*CreateCustomImageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateCustomImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomImageRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomImageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCustomImageRequest) GetDataSnapshotId() *string {
	return s.DataSnapshotId
}

func (s *CreateCustomImageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateCustomImageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCustomImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCustomImageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCustomImageRequest) GetSystemSnapshotId() *string {
	return s.SystemSnapshotId
}

func (s *CreateCustomImageRequest) GetTag() []*CreateCustomImageRequestTag {
	return s.Tag
}

func (s *CreateCustomImageRequest) SetClientToken(v string) *CreateCustomImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomImageRequest) SetDataSnapshotId(v string) *CreateCustomImageRequest {
	s.DataSnapshotId = &v
	return s
}

func (s *CreateCustomImageRequest) SetDescription(v string) *CreateCustomImageRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomImageRequest) SetImageName(v string) *CreateCustomImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateCustomImageRequest) SetInstanceId(v string) *CreateCustomImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomImageRequest) SetRegionId(v string) *CreateCustomImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomImageRequest) SetResourceGroupId(v string) *CreateCustomImageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCustomImageRequest) SetSystemSnapshotId(v string) *CreateCustomImageRequest {
	s.SystemSnapshotId = &v
	return s
}

func (s *CreateCustomImageRequest) SetTag(v []*CreateCustomImageRequestTag) *CreateCustomImageRequest {
	s.Tag = v
	return s
}

func (s *CreateCustomImageRequest) Validate() error {
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

type CreateCustomImageRequestTag struct {
	// The key of tag N that is added to the custom image. Valid values of N: 1 to 20. The tag key cannot be an empty string. It can be up to 64 characters in length, and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the custom image. Valid values of N: 1 to 20. The tag value can be an empty string. It can be up to 64 characters in length, and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCustomImageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomImageRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCustomImageRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCustomImageRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCustomImageRequestTag) SetKey(v string) *CreateCustomImageRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCustomImageRequestTag) SetValue(v string) *CreateCustomImageRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCustomImageRequestTag) Validate() error {
	return dara.Validate(s)
}

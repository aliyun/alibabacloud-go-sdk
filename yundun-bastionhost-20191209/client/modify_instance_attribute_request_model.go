// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyInstanceAttributeRequest
	GetDescription() *string
	SetInstanceId(v string) *ModifyInstanceAttributeRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyInstanceAttributeRequest
	GetRegionId() *string
}

type ModifyInstanceAttributeRequest struct {
	// The description of the Bastionhost instance.
	//
	// > The description can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). Maximum 30 characters.
	//
	// example:
	//
	// Bastionhost demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Bastionhost instance.
	//
	// > Call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about regions and zones, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceAttributeRequest) SetDescription(v string) *ModifyInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetRegionId(v string) *ModifyInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}

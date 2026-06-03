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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

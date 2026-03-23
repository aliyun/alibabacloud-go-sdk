// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceDescription(v string) *ModifyRCInstanceDescriptionRequest
	GetInstanceDescription() *string
	SetInstanceId(v string) *ModifyRCInstanceDescriptionRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyRCInstanceDescriptionRequest
	GetRegionId() *string
}

type ModifyRCInstanceDescriptionRequest struct {
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRCInstanceDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceDescriptionRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *ModifyRCInstanceDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCInstanceDescriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCInstanceDescriptionRequest) SetInstanceDescription(v string) *ModifyRCInstanceDescriptionRequest {
	s.InstanceDescription = &v
	return s
}

func (s *ModifyRCInstanceDescriptionRequest) SetInstanceId(v string) *ModifyRCInstanceDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCInstanceDescriptionRequest) SetRegionId(v string) *ModifyRCInstanceDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCInstanceDescriptionRequest) Validate() error {
	return dara.Validate(s)
}

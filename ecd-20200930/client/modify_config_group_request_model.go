// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConfigGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyConfigGroupRequest
	GetDescription() *string
	SetGroupId(v string) *ModifyConfigGroupRequest
	GetGroupId() *string
	SetName(v string) *ModifyConfigGroupRequest
	GetName() *string
	SetRegionId(v string) *ModifyConfigGroupRequest
	GetRegionId() *string
}

type ModifyConfigGroupRequest struct {
	// The description of the configuration group.
	//
	// example:
	//
	// ScheduledTask
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the configuration group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the configuration group.
	//
	// example:
	//
	// ScheduledTask
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region. Set the value to `cn-shanghai`.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyConfigGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyConfigGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyConfigGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyConfigGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyConfigGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModifyConfigGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyConfigGroupRequest) SetDescription(v string) *ModifyConfigGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyConfigGroupRequest) SetGroupId(v string) *ModifyConfigGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyConfigGroupRequest) SetName(v string) *ModifyConfigGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyConfigGroupRequest) SetRegionId(v string) *ModifyConfigGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyConfigGroupRequest) Validate() error {
	return dara.Validate(s)
}

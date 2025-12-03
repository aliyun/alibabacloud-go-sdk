// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *ModifyAccessGroupRequest
	GetAccessGroupId() *string
	SetAccessGroupName(v string) *ModifyAccessGroupRequest
	GetAccessGroupName() *string
	SetDescription(v string) *ModifyAccessGroupRequest
	GetDescription() *string
	SetInputRegionId(v string) *ModifyAccessGroupRequest
	GetInputRegionId() *string
}

type ModifyAccessGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	// example:
	//
	// my-online-cluster-policy
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s ModifyAccessGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupRequest) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *ModifyAccessGroupRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *ModifyAccessGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAccessGroupRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ModifyAccessGroupRequest) SetAccessGroupId(v string) *ModifyAccessGroupRequest {
	s.AccessGroupId = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetAccessGroupName(v string) *ModifyAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetDescription(v string) *ModifyAccessGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetInputRegionId(v string) *ModifyAccessGroupRequest {
	s.InputRegionId = &v
	return s
}

func (s *ModifyAccessGroupRequest) Validate() error {
	return dara.Validate(s)
}

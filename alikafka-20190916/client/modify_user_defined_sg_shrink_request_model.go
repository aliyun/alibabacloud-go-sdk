// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserDefinedSgShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyUserDefinedSgShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyUserDefinedSgShrinkRequest
	GetRegionId() *string
	SetSgIdListShrink(v string) *ModifyUserDefinedSgShrinkRequest
	GetSgIdListShrink() *string
}

type ModifyUserDefinedSgShrinkRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SgIdListShrink *string `json:"SgIdList,omitempty" xml:"SgIdList,omitempty"`
}

func (s ModifyUserDefinedSgShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserDefinedSgShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserDefinedSgShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserDefinedSgShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserDefinedSgShrinkRequest) GetSgIdListShrink() *string {
	return s.SgIdListShrink
}

func (s *ModifyUserDefinedSgShrinkRequest) SetInstanceId(v string) *ModifyUserDefinedSgShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserDefinedSgShrinkRequest) SetRegionId(v string) *ModifyUserDefinedSgShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserDefinedSgShrinkRequest) SetSgIdListShrink(v string) *ModifyUserDefinedSgShrinkRequest {
	s.SgIdListShrink = &v
	return s
}

func (s *ModifyUserDefinedSgShrinkRequest) Validate() error {
	return dara.Validate(s)
}

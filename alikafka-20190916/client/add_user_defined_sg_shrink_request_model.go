// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserDefinedSgShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddUserDefinedSgShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *AddUserDefinedSgShrinkRequest
	GetRegionId() *string
	SetSgIdListShrink(v string) *AddUserDefinedSgShrinkRequest
	GetSgIdListShrink() *string
}

type AddUserDefinedSgShrinkRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SgIdListShrink *string `json:"SgIdList,omitempty" xml:"SgIdList,omitempty"`
}

func (s AddUserDefinedSgShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserDefinedSgShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddUserDefinedSgShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddUserDefinedSgShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddUserDefinedSgShrinkRequest) GetSgIdListShrink() *string {
	return s.SgIdListShrink
}

func (s *AddUserDefinedSgShrinkRequest) SetInstanceId(v string) *AddUserDefinedSgShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUserDefinedSgShrinkRequest) SetRegionId(v string) *AddUserDefinedSgShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserDefinedSgShrinkRequest) SetSgIdListShrink(v string) *AddUserDefinedSgShrinkRequest {
	s.SgIdListShrink = &v
	return s
}

func (s *AddUserDefinedSgShrinkRequest) Validate() error {
	return dara.Validate(s)
}

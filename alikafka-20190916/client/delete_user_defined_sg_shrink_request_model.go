// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDefinedSgShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteUserDefinedSgShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteUserDefinedSgShrinkRequest
	GetRegionId() *string
	SetSgIdListShrink(v string) *DeleteUserDefinedSgShrinkRequest
	GetSgIdListShrink() *string
}

type DeleteUserDefinedSgShrinkRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SgIdListShrink *string `json:"SgIdList,omitempty" xml:"SgIdList,omitempty"`
}

func (s DeleteUserDefinedSgShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDefinedSgShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserDefinedSgShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteUserDefinedSgShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteUserDefinedSgShrinkRequest) GetSgIdListShrink() *string {
	return s.SgIdListShrink
}

func (s *DeleteUserDefinedSgShrinkRequest) SetInstanceId(v string) *DeleteUserDefinedSgShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserDefinedSgShrinkRequest) SetRegionId(v string) *DeleteUserDefinedSgShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserDefinedSgShrinkRequest) SetSgIdListShrink(v string) *DeleteUserDefinedSgShrinkRequest {
	s.SgIdListShrink = &v
	return s
}

func (s *DeleteUserDefinedSgShrinkRequest) Validate() error {
	return dara.Validate(s)
}

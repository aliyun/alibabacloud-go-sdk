// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWordGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIdListShrink(v string) *DeleteWordGroupShrinkRequest
	GetGroupIdListShrink() *string
	SetRegionId(v string) *DeleteWordGroupShrinkRequest
	GetRegionId() *string
}

type DeleteWordGroupShrinkRequest struct {
	// List of keyword strategy IDs.
	GroupIdListShrink *string `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteWordGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWordGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteWordGroupShrinkRequest) GetGroupIdListShrink() *string {
	return s.GroupIdListShrink
}

func (s *DeleteWordGroupShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteWordGroupShrinkRequest) SetGroupIdListShrink(v string) *DeleteWordGroupShrinkRequest {
	s.GroupIdListShrink = &v
	return s
}

func (s *DeleteWordGroupShrinkRequest) SetRegionId(v string) *DeleteWordGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteWordGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

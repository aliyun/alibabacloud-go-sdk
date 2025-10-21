// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWordGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIdList(v []*int64) *DeleteWordGroupRequest
	GetGroupIdList() []*int64
	SetRegionId(v string) *DeleteWordGroupRequest
	GetRegionId() *string
}

type DeleteWordGroupRequest struct {
	// List of keyword strategy IDs.
	GroupIdList []*int64 `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty" type:"Repeated"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteWordGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWordGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteWordGroupRequest) GetGroupIdList() []*int64 {
	return s.GroupIdList
}

func (s *DeleteWordGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteWordGroupRequest) SetGroupIdList(v []*int64) *DeleteWordGroupRequest {
	s.GroupIdList = v
	return s
}

func (s *DeleteWordGroupRequest) SetRegionId(v string) *DeleteWordGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteWordGroupRequest) Validate() error {
	return dara.Validate(s)
}

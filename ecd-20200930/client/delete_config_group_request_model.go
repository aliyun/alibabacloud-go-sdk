// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIds(v []*string) *DeleteConfigGroupRequest
	GetGroupIds() []*string
	SetRegionId(v string) *DeleteConfigGroupRequest
	GetRegionId() *string
}

type DeleteConfigGroupRequest struct {
	// The IDs of the configuration groups that you want to delete.
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The ID of the region. Set the value to `cn-shanghai`.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteConfigGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigGroupRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *DeleteConfigGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteConfigGroupRequest) SetGroupIds(v []*string) *DeleteConfigGroupRequest {
	s.GroupIds = v
	return s
}

func (s *DeleteConfigGroupRequest) SetRegionId(v string) *DeleteConfigGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteConfigGroupRequest) Validate() error {
	return dara.Validate(s)
}

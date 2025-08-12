// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DeleteMonitorGroupRequest
	GetGroupId() *int64
	SetRegionId(v string) *DeleteMonitorGroupRequest
	GetRegionId() *string
}

type DeleteMonitorGroupRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId  *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMonitorGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteMonitorGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMonitorGroupRequest) SetGroupId(v int64) *DeleteMonitorGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMonitorGroupRequest) SetRegionId(v string) *DeleteMonitorGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMonitorGroupRequest) Validate() error {
	return dara.Validate(s)
}

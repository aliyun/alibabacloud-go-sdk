// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteMonitorGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteMonitorGroupRequest
	GetDryRun() *bool
	SetGroupId(v int64) *DeleteMonitorGroupRequest
	GetGroupId() *int64
	SetRegionId(v string) *DeleteMonitorGroupRequest
	GetRegionId() *string
}

type DeleteMonitorGroupRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. The dry run checks permissions and instance status. Valid values:
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, the name and description of the route table are modified.
	//
	// - **true**: sends a dry run request. Only the check is performed, and the name and description of the route table are not modified. The check items include whether required parameters are specified and the request format. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 246139436
	GroupId  *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMonitorGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteMonitorGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteMonitorGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteMonitorGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMonitorGroupRequest) SetClientToken(v string) *DeleteMonitorGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteMonitorGroupRequest) SetDryRun(v bool) *DeleteMonitorGroupRequest {
	s.DryRun = &v
	return s
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

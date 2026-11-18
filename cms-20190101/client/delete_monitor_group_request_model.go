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
	// Ensures the idempotence of the request. You can generate a parameter value on your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 6868E9A1-56E6-5FD7-9684-E2963BBAFB5F
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to execute a dry run on the request, including checks on permissions and instance status. Valid values:
	//
	// - **false*	- (default): Sends a normal request. If the check succeeds, the name and description of the routing table are directly modified.
	//
	// - **true**: Sends a check request to only perform the check, and the name and description of the routing table are not modified. The check items include whether required parameters are specified and the request format. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
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

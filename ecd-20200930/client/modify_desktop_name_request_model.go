// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *ModifyDesktopNameRequest
	GetDesktopId() *string
	SetDesktopIds(v []*string) *ModifyDesktopNameRequest
	GetDesktopIds() []*string
	SetNewDesktopName(v string) *ModifyDesktopNameRequest
	GetNewDesktopName() *string
	SetRegionId(v string) *ModifyDesktopNameRequest
	GetRegionId() *string
	SetUserAssignMode(v string) *ModifyDesktopNameRequest
	GetUserAssignMode() *string
}

type ModifyDesktopNameRequest struct {
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The list of cloud computer IDs.
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// The new name of the cloud computer. The name must meet the following requirements:
	//
	// - The name cannot exceed 64 characters in length.
	//
	// - The name must start with a letter or a Chinese character and cannot start with `http://` or `https://`.
	//
	// - The name can contain Chinese characters, letters, digits, colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// DemoComputer01
	NewDesktopName *string `json:"NewDesktopName,omitempty" xml:"NewDesktopName,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user assignment mode.
	//
	// example:
	//
	// ALL
	UserAssignMode *string `json:"UserAssignMode,omitempty" xml:"UserAssignMode,omitempty"`
}

func (s ModifyDesktopNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ModifyDesktopNameRequest) GetDesktopIds() []*string {
	return s.DesktopIds
}

func (s *ModifyDesktopNameRequest) GetNewDesktopName() *string {
	return s.NewDesktopName
}

func (s *ModifyDesktopNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDesktopNameRequest) GetUserAssignMode() *string {
	return s.UserAssignMode
}

func (s *ModifyDesktopNameRequest) SetDesktopId(v string) *ModifyDesktopNameRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetDesktopIds(v []*string) *ModifyDesktopNameRequest {
	s.DesktopIds = v
	return s
}

func (s *ModifyDesktopNameRequest) SetNewDesktopName(v string) *ModifyDesktopNameRequest {
	s.NewDesktopName = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetRegionId(v string) *ModifyDesktopNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetUserAssignMode(v string) *ModifyDesktopNameRequest {
	s.UserAssignMode = &v
	return s
}

func (s *ModifyDesktopNameRequest) Validate() error {
	return dara.Validate(s)
}

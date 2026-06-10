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
	// The ID of the cloud desktop.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId  *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// The new name of the cloud desktop. The name must meet the following requirements:
	//
	// - The name must be 1 to 64 characters in length.
	//
	// - The name must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// - The name can contain digits, letters, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// DemoComputer01
	NewDesktopName *string `json:"NewDesktopName,omitempty" xml:"NewDesktopName,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to obtain a list of regions that are supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

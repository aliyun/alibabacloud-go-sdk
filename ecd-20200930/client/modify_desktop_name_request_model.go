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
	SetNewDesktopName(v string) *ModifyDesktopNameRequest
	GetNewDesktopName() *string
	SetRegionId(v string) *ModifyDesktopNameRequest
	GetRegionId() *string
}

type ModifyDesktopNameRequest struct {
	// The ID of the cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The new name of the cloud computer. The name of the cloud computer must meet the following requirements:
	//
	// 	- The name must be 1 to 64 characters in length.
	//
	// 	- The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// 	- The name can only contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	NewDesktopName *string `json:"NewDesktopName,omitempty" xml:"NewDesktopName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ModifyDesktopNameRequest) GetNewDesktopName() *string {
	return s.NewDesktopName
}

func (s *ModifyDesktopNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDesktopNameRequest) SetDesktopId(v string) *ModifyDesktopNameRequest {
	s.DesktopId = &v
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

func (s *ModifyDesktopNameRequest) Validate() error {
	return dara.Validate(s)
}

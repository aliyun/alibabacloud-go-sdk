// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *StartDesktopsRequest
	GetDesktopId() []*string
	SetRegionId(v string) *StartDesktopsRequest
	GetRegionId() *string
}

type StartDesktopsRequest struct {
	// The cloud computer IDs. You can specify the IDs of 1 to 100 cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StartDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *StartDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDesktopsRequest) SetDesktopId(v []*string) *StartDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StartDesktopsRequest) SetRegionId(v string) *StartDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StartDesktopsRequest) Validate() error {
	return dara.Validate(s)
}

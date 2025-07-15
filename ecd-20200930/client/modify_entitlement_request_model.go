// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEntitlementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *ModifyEntitlementRequest
	GetDesktopId() *string
	SetEndUserId(v []*string) *ModifyEntitlementRequest
	GetEndUserId() []*string
	SetRegionId(v string) *ModifyEntitlementRequest
	GetRegionId() *string
}

type ModifyEntitlementRequest struct {
	// The ID of the cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The username IDs. End users specified by this parameter become the end users of the cloud computer, and the original end users of the cloud computer are removed. You can specify 1 to 100 IDs.
	//
	// example:
	//
	// alice
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyEntitlementRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEntitlementRequest) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ModifyEntitlementRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *ModifyEntitlementRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyEntitlementRequest) SetDesktopId(v string) *ModifyEntitlementRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyEntitlementRequest) SetEndUserId(v []*string) *ModifyEntitlementRequest {
	s.EndUserId = v
	return s
}

func (s *ModifyEntitlementRequest) SetRegionId(v string) *ModifyEntitlementRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEntitlementRequest) Validate() error {
	return dara.Validate(s)
}

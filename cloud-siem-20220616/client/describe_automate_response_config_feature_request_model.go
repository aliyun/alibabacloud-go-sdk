// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutomateResponseConfigFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoResponseType(v string) *DescribeAutomateResponseConfigFeatureRequest
	GetAutoResponseType() *string
	SetRegionId(v string) *DescribeAutomateResponseConfigFeatureRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAutomateResponseConfigFeatureRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAutomateResponseConfigFeatureRequest
	GetRoleType() *int32
}

type DescribeAutomateResponseConfigFeatureRequest struct {
	// The type of automated response. Valid values:
	//
	// - event: event
	//
	// - alert: alert
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The region where the Data Management center of threat analysis is deployed. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Assets in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can use this parameter to switch to the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of view. Valid values:
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts that are managed by the administrator account.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureRequest) GetAutoResponseType() *string {
	return s.AutoResponseType
}

func (s *DescribeAutomateResponseConfigFeatureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutomateResponseConfigFeatureRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAutomateResponseConfigFeatureRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAutomateResponseConfigFeatureRequest) SetAutoResponseType(v string) *DescribeAutomateResponseConfigFeatureRequest {
	s.AutoResponseType = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureRequest) SetRegionId(v string) *DescribeAutomateResponseConfigFeatureRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureRequest) SetRoleFor(v int64) *DescribeAutomateResponseConfigFeatureRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureRequest) SetRoleType(v int32) *DescribeAutomateResponseConfigFeatureRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureRequest) Validate() error {
	return dara.Validate(s)
}

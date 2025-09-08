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
	// The type of the automated response rule. Valid values:
	//
	// 	- event
	//
	// 	- alert
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
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

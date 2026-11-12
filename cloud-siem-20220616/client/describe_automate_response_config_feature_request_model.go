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
	// The automated response type. Valid values:
	//
	// - event: Event.
	//
	// - alert: Alert.
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The region where the data management center of Threat Analysis is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
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

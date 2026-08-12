// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeVersionConfigShrinkRequest
	GetRegionId() *string
	SetSdkRequestShrink(v string) *DescribeVersionConfigShrinkRequest
	GetSdkRequestShrink() *string
}

type DescribeVersionConfigShrinkRequest struct {
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-guangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request.
	SdkRequestShrink *string `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty"`
}

func (s DescribeVersionConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVersionConfigShrinkRequest) GetSdkRequestShrink() *string {
	return s.SdkRequestShrink
}

func (s *DescribeVersionConfigShrinkRequest) SetRegionId(v string) *DescribeVersionConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVersionConfigShrinkRequest) SetSdkRequestShrink(v string) *DescribeVersionConfigShrinkRequest {
	s.SdkRequestShrink = &v
	return s
}

func (s *DescribeVersionConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}

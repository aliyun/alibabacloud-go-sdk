// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSasTrialShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateSasTrialShrinkRequest
	GetRegionId() *string
	SetSdkRequestShrink(v string) *CreateSasTrialShrinkRequest
	GetSdkRequestShrink() *string
}

type CreateSasTrialShrinkRequest struct {
	// The region ID of the access control instance. You can call the DescribeRegions operation to query the region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request.
	SdkRequestShrink *string `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty"`
}

func (s CreateSasTrialShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSasTrialShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSasTrialShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSasTrialShrinkRequest) GetSdkRequestShrink() *string {
	return s.SdkRequestShrink
}

func (s *CreateSasTrialShrinkRequest) SetRegionId(v string) *CreateSasTrialShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSasTrialShrinkRequest) SetSdkRequestShrink(v string) *CreateSasTrialShrinkRequest {
	s.SdkRequestShrink = &v
	return s
}

func (s *CreateSasTrialShrinkRequest) Validate() error {
	return dara.Validate(s)
}

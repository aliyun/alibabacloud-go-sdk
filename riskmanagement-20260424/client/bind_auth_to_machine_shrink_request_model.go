// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAuthToMachineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *BindAuthToMachineShrinkRequest
	GetRegionId() *string
	SetSdkRequestShrink(v string) *BindAuthToMachineShrinkRequest
	GetSdkRequestShrink() *string
}

type BindAuthToMachineShrinkRequest struct {
	// The region ID of the Smart Access Gateway instance.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request.
	SdkRequestShrink *string `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty"`
}

func (s BindAuthToMachineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindAuthToMachineShrinkRequest) GetSdkRequestShrink() *string {
	return s.SdkRequestShrink
}

func (s *BindAuthToMachineShrinkRequest) SetRegionId(v string) *BindAuthToMachineShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *BindAuthToMachineShrinkRequest) SetSdkRequestShrink(v string) *BindAuthToMachineShrinkRequest {
	s.SdkRequestShrink = &v
	return s
}

func (s *BindAuthToMachineShrinkRequest) Validate() error {
	return dara.Validate(s)
}

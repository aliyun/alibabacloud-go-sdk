// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeSuspEventsShrinkRequest
	GetRegionId() *string
	SetSdkRequestShrink(v string) *DescribeSuspEventsShrinkRequest
	GetSdkRequestShrink() *string
}

type DescribeSuspEventsShrinkRequest struct {
	// example:
	//
	// cn-guangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SdkRequestShrink *string `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty"`
}

func (s DescribeSuspEventsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSuspEventsShrinkRequest) GetSdkRequestShrink() *string {
	return s.SdkRequestShrink
}

func (s *DescribeSuspEventsShrinkRequest) SetRegionId(v string) *DescribeSuspEventsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) SetSdkRequestShrink(v string) *DescribeSuspEventsShrinkRequest {
	s.SdkRequestShrink = &v
	return s
}

func (s *DescribeSuspEventsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

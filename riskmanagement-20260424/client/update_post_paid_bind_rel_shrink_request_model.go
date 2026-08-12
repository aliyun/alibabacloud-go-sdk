// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostPaidBindRelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdatePostPaidBindRelShrinkRequest
	GetRegionId() *string
	SetSdkRequestShrink(v string) *UpdatePostPaidBindRelShrinkRequest
	GetSdkRequestShrink() *string
}

type UpdatePostPaidBindRelShrinkRequest struct {
	// The region ID of the instance.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request parameters.
	SdkRequestShrink *string `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty"`
}

func (s UpdatePostPaidBindRelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostPaidBindRelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePostPaidBindRelShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePostPaidBindRelShrinkRequest) GetSdkRequestShrink() *string {
	return s.SdkRequestShrink
}

func (s *UpdatePostPaidBindRelShrinkRequest) SetRegionId(v string) *UpdatePostPaidBindRelShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePostPaidBindRelShrinkRequest) SetSdkRequestShrink(v string) *UpdatePostPaidBindRelShrinkRequest {
	s.SdkRequestShrink = &v
	return s
}

func (s *UpdatePostPaidBindRelShrinkRequest) Validate() error {
	return dara.Validate(s)
}

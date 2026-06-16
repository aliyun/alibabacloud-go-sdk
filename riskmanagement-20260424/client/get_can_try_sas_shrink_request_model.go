// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCanTrySasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetCanTrySasShrinkRequest
	GetRegionId() *string
	SetSdkRequestShrink(v string) *GetCanTrySasShrinkRequest
	GetSdkRequestShrink() *string
}

type GetCanTrySasShrinkRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SdkRequestShrink *string `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty"`
}

func (s GetCanTrySasShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCanTrySasShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCanTrySasShrinkRequest) GetSdkRequestShrink() *string {
	return s.SdkRequestShrink
}

func (s *GetCanTrySasShrinkRequest) SetRegionId(v string) *GetCanTrySasShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetCanTrySasShrinkRequest) SetSdkRequestShrink(v string) *GetCanTrySasShrinkRequest {
	s.SdkRequestShrink = &v
	return s
}

func (s *GetCanTrySasShrinkRequest) Validate() error {
	return dara.Validate(s)
}

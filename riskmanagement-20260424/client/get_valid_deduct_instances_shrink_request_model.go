// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidDeductInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetValidDeductInstancesShrinkRequest
	GetRegionId() *string
	SetSdkRequestShrink(v string) *GetValidDeductInstancesShrinkRequest
	GetSdkRequestShrink() *string
}

type GetValidDeductInstancesShrinkRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SdkRequestShrink *string `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty"`
}

func (s GetValidDeductInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetValidDeductInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetValidDeductInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetValidDeductInstancesShrinkRequest) GetSdkRequestShrink() *string {
	return s.SdkRequestShrink
}

func (s *GetValidDeductInstancesShrinkRequest) SetRegionId(v string) *GetValidDeductInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetValidDeductInstancesShrinkRequest) SetSdkRequestShrink(v string) *GetValidDeductInstancesShrinkRequest {
	s.SdkRequestShrink = &v
	return s
}

func (s *GetValidDeductInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

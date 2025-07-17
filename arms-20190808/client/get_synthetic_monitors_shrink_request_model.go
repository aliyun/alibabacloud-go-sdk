// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticMonitorsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *GetSyntheticMonitorsShrinkRequest
	GetFilterShrink() *string
	SetRegionId(v string) *GetSyntheticMonitorsShrinkRequest
	GetRegionId() *string
}

type GetSyntheticMonitorsShrinkRequest struct {
	// The query conditions.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSyntheticMonitorsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticMonitorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSyntheticMonitorsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetSyntheticMonitorsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSyntheticMonitorsShrinkRequest) SetFilterShrink(v string) *GetSyntheticMonitorsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetSyntheticMonitorsShrinkRequest) SetRegionId(v string) *GetSyntheticMonitorsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetSyntheticMonitorsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

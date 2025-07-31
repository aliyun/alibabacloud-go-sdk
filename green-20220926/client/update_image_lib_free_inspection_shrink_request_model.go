// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageLibFreeInspectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigShrink(v string) *UpdateImageLibFreeInspectionShrinkRequest
	GetConfigShrink() *string
	SetRegionId(v string) *UpdateImageLibFreeInspectionShrinkRequest
	GetRegionId() *string
}

type UpdateImageLibFreeInspectionShrinkRequest struct {
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateImageLibFreeInspectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageLibFreeInspectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageLibFreeInspectionShrinkRequest) GetConfigShrink() *string {
	return s.ConfigShrink
}

func (s *UpdateImageLibFreeInspectionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateImageLibFreeInspectionShrinkRequest) SetConfigShrink(v string) *UpdateImageLibFreeInspectionShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *UpdateImageLibFreeInspectionShrinkRequest) SetRegionId(v string) *UpdateImageLibFreeInspectionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateImageLibFreeInspectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}

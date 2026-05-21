// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoloWebLoginSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetHoloWebLoginSettingRequest
	GetRegionId() *string
}

type GetHoloWebLoginSettingRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHoloWebLoginSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHoloWebLoginSettingRequest) GoString() string {
	return s.String()
}

func (s *GetHoloWebLoginSettingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetHoloWebLoginSettingRequest) SetRegionId(v string) *GetHoloWebLoginSettingRequest {
	s.RegionId = &v
	return s
}

func (s *GetHoloWebLoginSettingRequest) Validate() error {
	return dara.Validate(s)
}

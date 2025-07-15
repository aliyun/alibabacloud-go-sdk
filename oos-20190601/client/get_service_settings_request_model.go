// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetServiceSettingsRequest
	GetRegionId() *string
}

type GetServiceSettingsRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServiceSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceSettingsRequest) SetRegionId(v string) *GetServiceSettingsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceSettingsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegionConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetZoneId(v string) *GetRegionConfigurationRequest
	GetZoneId() *string
}

type GetRegionConfigurationRequest struct {
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s GetRegionConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetRegionConfigurationRequest) SetZoneId(v string) *GetRegionConfigurationRequest {
	s.ZoneId = &v
	return s
}

func (s *GetRegionConfigurationRequest) Validate() error {
	return dara.Validate(s)
}

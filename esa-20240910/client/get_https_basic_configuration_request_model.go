// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpsBasicConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpsBasicConfigurationRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetHttpsBasicConfigurationRequest
	GetSiteId() *int64
}

type GetHttpsBasicConfigurationRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetHttpsBasicConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHttpsBasicConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetHttpsBasicConfigurationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpsBasicConfigurationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetHttpsBasicConfigurationRequest) SetConfigId(v int64) *GetHttpsBasicConfigurationRequest {
	s.ConfigId = &v
	return s
}

func (s *GetHttpsBasicConfigurationRequest) SetSiteId(v int64) *GetHttpsBasicConfigurationRequest {
	s.SiteId = &v
	return s
}

func (s *GetHttpsBasicConfigurationRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpsApplicationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetHttpsApplicationConfigurationRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetHttpsApplicationConfigurationRequest
	GetSiteId() *int64
}

type GetHttpsApplicationConfigurationRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetHttpsApplicationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHttpsApplicationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetHttpsApplicationConfigurationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpsApplicationConfigurationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetHttpsApplicationConfigurationRequest) SetConfigId(v int64) *GetHttpsApplicationConfigurationRequest {
	s.ConfigId = &v
	return s
}

func (s *GetHttpsApplicationConfigurationRequest) SetSiteId(v int64) *GetHttpsApplicationConfigurationRequest {
	s.SiteId = &v
	return s
}

func (s *GetHttpsApplicationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}

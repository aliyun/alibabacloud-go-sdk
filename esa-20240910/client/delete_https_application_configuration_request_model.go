// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpsApplicationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteHttpsApplicationConfigurationRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteHttpsApplicationConfigurationRequest
	GetSiteId() *int64
}

type DeleteHttpsApplicationConfigurationRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteHttpsApplicationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpsApplicationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteHttpsApplicationConfigurationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteHttpsApplicationConfigurationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteHttpsApplicationConfigurationRequest) SetConfigId(v int64) *DeleteHttpsApplicationConfigurationRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteHttpsApplicationConfigurationRequest) SetSiteId(v int64) *DeleteHttpsApplicationConfigurationRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteHttpsApplicationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}

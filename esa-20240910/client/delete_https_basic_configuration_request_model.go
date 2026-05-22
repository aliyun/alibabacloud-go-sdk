// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpsBasicConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteHttpsBasicConfigurationRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteHttpsBasicConfigurationRequest
	GetSiteId() *int64
}

type DeleteHttpsBasicConfigurationRequest struct {
	// ConfigId of the configuration, which can be obtained by calling the [ListHttpsBasicConfigurations](~~ListHttpsBasicConfigurations~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteHttpsBasicConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpsBasicConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteHttpsBasicConfigurationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteHttpsBasicConfigurationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteHttpsBasicConfigurationRequest) SetConfigId(v int64) *DeleteHttpsBasicConfigurationRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteHttpsBasicConfigurationRequest) SetSiteId(v int64) *DeleteHttpsBasicConfigurationRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteHttpsBasicConfigurationRequest) Validate() error {
	return dara.Validate(s)
}

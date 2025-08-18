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
	// ConfigId of the configuration, which can be obtained by calling the [listHttpsApplicationConfigurations](https://help.aliyun.com/document_detail/2869087.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
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

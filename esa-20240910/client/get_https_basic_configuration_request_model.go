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
	// ConfigId of the configuration, which can be obtained by calling the [ListHttpsBasicConfigurations](https://help.aliyun.com/document_detail/2867470.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
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

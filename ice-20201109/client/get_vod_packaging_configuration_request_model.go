// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodPackagingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationName(v string) *GetVodPackagingConfigurationRequest
	GetConfigurationName() *string
}

type GetVodPackagingConfigurationRequest struct {
	// The name of the packaging configuration.
	//
	// example:
	//
	// hls_3s
	ConfigurationName *string `json:"ConfigurationName,omitempty" xml:"ConfigurationName,omitempty"`
}

func (s GetVodPackagingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetVodPackagingConfigurationRequest) GetConfigurationName() *string {
	return s.ConfigurationName
}

func (s *GetVodPackagingConfigurationRequest) SetConfigurationName(v string) *GetVodPackagingConfigurationRequest {
	s.ConfigurationName = &v
	return s
}

func (s *GetVodPackagingConfigurationRequest) Validate() error {
	return dara.Validate(s)
}

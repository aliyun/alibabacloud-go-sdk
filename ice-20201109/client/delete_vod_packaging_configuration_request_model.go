// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodPackagingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationName(v string) *DeleteVodPackagingConfigurationRequest
	GetConfigurationName() *string
}

type DeleteVodPackagingConfigurationRequest struct {
	// The name of the packaging configuration.
	//
	// example:
	//
	// hls_3s
	ConfigurationName *string `json:"ConfigurationName,omitempty" xml:"ConfigurationName,omitempty"`
}

func (s DeleteVodPackagingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodPackagingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodPackagingConfigurationRequest) GetConfigurationName() *string {
	return s.ConfigurationName
}

func (s *DeleteVodPackagingConfigurationRequest) SetConfigurationName(v string) *DeleteVodPackagingConfigurationRequest {
	s.ConfigurationName = &v
	return s
}

func (s *DeleteVodPackagingConfigurationRequest) Validate() error {
	return dara.Validate(s)
}

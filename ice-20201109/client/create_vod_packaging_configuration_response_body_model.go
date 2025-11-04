// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPackagingConfiguration(v *VodPackagingConfiguration) *CreateVodPackagingConfigurationResponseBody
	GetPackagingConfiguration() *VodPackagingConfiguration
	SetRequestId(v string) *CreateVodPackagingConfigurationResponseBody
	GetRequestId() *string
}

type CreateVodPackagingConfigurationResponseBody struct {
	// The packaging configuration.
	PackagingConfiguration *VodPackagingConfiguration `json:"PackagingConfiguration,omitempty" xml:"PackagingConfiguration,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVodPackagingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingConfigurationResponseBody) GetPackagingConfiguration() *VodPackagingConfiguration {
	return s.PackagingConfiguration
}

func (s *CreateVodPackagingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVodPackagingConfigurationResponseBody) SetPackagingConfiguration(v *VodPackagingConfiguration) *CreateVodPackagingConfigurationResponseBody {
	s.PackagingConfiguration = v
	return s
}

func (s *CreateVodPackagingConfigurationResponseBody) SetRequestId(v string) *CreateVodPackagingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVodPackagingConfigurationResponseBody) Validate() error {
	if s.PackagingConfiguration != nil {
		if err := s.PackagingConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

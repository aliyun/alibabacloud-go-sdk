// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodPackagingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPackagingConfiguration(v *VodPackagingConfiguration) *GetVodPackagingConfigurationResponseBody
	GetPackagingConfiguration() *VodPackagingConfiguration
	SetRequestId(v string) *GetVodPackagingConfigurationResponseBody
	GetRequestId() *string
}

type GetVodPackagingConfigurationResponseBody struct {
	// The information about the packaging configuration.
	PackagingConfiguration *VodPackagingConfiguration `json:"PackagingConfiguration,omitempty" xml:"PackagingConfiguration,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVodPackagingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetVodPackagingConfigurationResponseBody) GetPackagingConfiguration() *VodPackagingConfiguration {
	return s.PackagingConfiguration
}

func (s *GetVodPackagingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVodPackagingConfigurationResponseBody) SetPackagingConfiguration(v *VodPackagingConfiguration) *GetVodPackagingConfigurationResponseBody {
	s.PackagingConfiguration = v
	return s
}

func (s *GetVodPackagingConfigurationResponseBody) SetRequestId(v string) *GetVodPackagingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVodPackagingConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

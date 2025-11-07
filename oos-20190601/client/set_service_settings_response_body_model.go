// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetServiceSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetServiceSettingsResponseBody
	GetRequestId() *string
	SetServiceSettings(v []*SetServiceSettingsResponseBodyServiceSettings) *SetServiceSettingsResponseBody
	GetServiceSettings() []*SetServiceSettingsResponseBodyServiceSettings
}

type SetServiceSettingsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CBEC8072-BEC2-478E-8EAE-E723BA79CF19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of service settings.
	ServiceSettings []*SetServiceSettingsResponseBodyServiceSettings `json:"ServiceSettings,omitempty" xml:"ServiceSettings,omitempty" type:"Repeated"`
}

func (s SetServiceSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetServiceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetServiceSettingsResponseBody) GetServiceSettings() []*SetServiceSettingsResponseBodyServiceSettings {
	return s.ServiceSettings
}

func (s *SetServiceSettingsResponseBody) SetRequestId(v string) *SetServiceSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetServiceSettingsResponseBody) SetServiceSettings(v []*SetServiceSettingsResponseBodyServiceSettings) *SetServiceSettingsResponseBody {
	s.ServiceSettings = v
	return s
}

func (s *SetServiceSettingsResponseBody) Validate() error {
	if s.ServiceSettings != nil {
		for _, item := range s.ServiceSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetServiceSettingsResponseBodyServiceSettings struct {
	// The name of OSS bucket to deliver.
	//
	// example:
	//
	// OssBucketName
	DeliveryOssBucketName *string `json:"DeliveryOssBucketName,omitempty" xml:"DeliveryOssBucketName,omitempty"`
	// Whether to enable OSS delivery.
	//
	// example:
	//
	// true
	DeliveryOssEnabled *bool `json:"DeliveryOssEnabled,omitempty" xml:"DeliveryOssEnabled,omitempty"`
	// The key prefix of OSS to deliver.
	//
	// example:
	//
	// oos/execution
	DeliveryOssKeyPrefix *string `json:"DeliveryOssKeyPrefix,omitempty" xml:"DeliveryOssKeyPrefix,omitempty"`
	// Whether to enable SLS delivery.
	//
	// example:
	//
	// false
	DeliverySlsEnabled *bool `json:"DeliverySlsEnabled,omitempty" xml:"DeliverySlsEnabled,omitempty"`
	// The name of SLS project to deliver.
	//
	// example:
	//
	// SlsProjectName
	DeliverySlsProjectName *string `json:"DeliverySlsProjectName,omitempty" xml:"DeliverySlsProjectName,omitempty"`
	// The id of RDC Enterprise.
	//
	// example:
	//
	// RdcEnterpriseId
	RdcEnterpriseId *string `json:"RdcEnterpriseId,omitempty" xml:"RdcEnterpriseId,omitempty"`
}

func (s SetServiceSettingsResponseBodyServiceSettings) String() string {
	return dara.Prettify(s)
}

func (s SetServiceSettingsResponseBodyServiceSettings) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsResponseBodyServiceSettings) GetDeliveryOssBucketName() *string {
	return s.DeliveryOssBucketName
}

func (s *SetServiceSettingsResponseBodyServiceSettings) GetDeliveryOssEnabled() *bool {
	return s.DeliveryOssEnabled
}

func (s *SetServiceSettingsResponseBodyServiceSettings) GetDeliveryOssKeyPrefix() *string {
	return s.DeliveryOssKeyPrefix
}

func (s *SetServiceSettingsResponseBodyServiceSettings) GetDeliverySlsEnabled() *bool {
	return s.DeliverySlsEnabled
}

func (s *SetServiceSettingsResponseBodyServiceSettings) GetDeliverySlsProjectName() *string {
	return s.DeliverySlsProjectName
}

func (s *SetServiceSettingsResponseBodyServiceSettings) GetRdcEnterpriseId() *string {
	return s.RdcEnterpriseId
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssBucketName(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssBucketName = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssEnabled(v bool) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssEnabled = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssKeyPrefix(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssKeyPrefix = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsEnabled(v bool) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsEnabled = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsProjectName(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsProjectName = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetRdcEnterpriseId(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.RdcEnterpriseId = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) Validate() error {
	return dara.Validate(s)
}

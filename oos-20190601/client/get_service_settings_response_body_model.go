// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceSettingsResponseBody
	GetRequestId() *string
	SetServiceSettings(v []*GetServiceSettingsResponseBodyServiceSettings) *GetServiceSettingsResponseBody
	GetServiceSettings() []*GetServiceSettingsResponseBodyServiceSettings
}

type GetServiceSettingsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9F755DC9-C0CF-4598-B2E3-2CC763F18CB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of service settings.
	ServiceSettings []*GetServiceSettingsResponseBodyServiceSettings `json:"ServiceSettings,omitempty" xml:"ServiceSettings,omitempty" type:"Repeated"`
}

func (s GetServiceSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceSettingsResponseBody) GetServiceSettings() []*GetServiceSettingsResponseBodyServiceSettings {
	return s.ServiceSettings
}

func (s *GetServiceSettingsResponseBody) SetRequestId(v string) *GetServiceSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceSettingsResponseBody) SetServiceSettings(v []*GetServiceSettingsResponseBodyServiceSettings) *GetServiceSettingsResponseBody {
	s.ServiceSettings = v
	return s
}

func (s *GetServiceSettingsResponseBody) Validate() error {
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

type GetServiceSettingsResponseBodyServiceSettings struct {
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
	// false
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

func (s GetServiceSettingsResponseBodyServiceSettings) String() string {
	return dara.Prettify(s)
}

func (s GetServiceSettingsResponseBodyServiceSettings) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsResponseBodyServiceSettings) GetDeliveryOssBucketName() *string {
	return s.DeliveryOssBucketName
}

func (s *GetServiceSettingsResponseBodyServiceSettings) GetDeliveryOssEnabled() *bool {
	return s.DeliveryOssEnabled
}

func (s *GetServiceSettingsResponseBodyServiceSettings) GetDeliveryOssKeyPrefix() *string {
	return s.DeliveryOssKeyPrefix
}

func (s *GetServiceSettingsResponseBodyServiceSettings) GetDeliverySlsEnabled() *bool {
	return s.DeliverySlsEnabled
}

func (s *GetServiceSettingsResponseBodyServiceSettings) GetDeliverySlsProjectName() *string {
	return s.DeliverySlsProjectName
}

func (s *GetServiceSettingsResponseBodyServiceSettings) GetRdcEnterpriseId() *string {
	return s.RdcEnterpriseId
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssBucketName(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssBucketName = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssEnabled(v bool) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssEnabled = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssKeyPrefix(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssKeyPrefix = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsEnabled(v bool) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsEnabled = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsProjectName(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsProjectName = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetRdcEnterpriseId(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.RdcEnterpriseId = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) Validate() error {
	return dara.Validate(s)
}

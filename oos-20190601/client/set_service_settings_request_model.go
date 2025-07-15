// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetServiceSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryOssBucketName(v string) *SetServiceSettingsRequest
	GetDeliveryOssBucketName() *string
	SetDeliveryOssEnabled(v bool) *SetServiceSettingsRequest
	GetDeliveryOssEnabled() *bool
	SetDeliveryOssKeyPrefix(v string) *SetServiceSettingsRequest
	GetDeliveryOssKeyPrefix() *string
	SetDeliverySlsEnabled(v bool) *SetServiceSettingsRequest
	GetDeliverySlsEnabled() *bool
	SetDeliverySlsProjectName(v string) *SetServiceSettingsRequest
	GetDeliverySlsProjectName() *string
	SetRdcEnterpriseId(v string) *SetServiceSettingsRequest
	GetRdcEnterpriseId() *string
	SetRegionId(v string) *SetServiceSettingsRequest
	GetRegionId() *string
}

type SetServiceSettingsRequest struct {
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetServiceSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s SetServiceSettingsRequest) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsRequest) GetDeliveryOssBucketName() *string {
	return s.DeliveryOssBucketName
}

func (s *SetServiceSettingsRequest) GetDeliveryOssEnabled() *bool {
	return s.DeliveryOssEnabled
}

func (s *SetServiceSettingsRequest) GetDeliveryOssKeyPrefix() *string {
	return s.DeliveryOssKeyPrefix
}

func (s *SetServiceSettingsRequest) GetDeliverySlsEnabled() *bool {
	return s.DeliverySlsEnabled
}

func (s *SetServiceSettingsRequest) GetDeliverySlsProjectName() *string {
	return s.DeliverySlsProjectName
}

func (s *SetServiceSettingsRequest) GetRdcEnterpriseId() *string {
	return s.RdcEnterpriseId
}

func (s *SetServiceSettingsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetServiceSettingsRequest) SetDeliveryOssBucketName(v string) *SetServiceSettingsRequest {
	s.DeliveryOssBucketName = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliveryOssEnabled(v bool) *SetServiceSettingsRequest {
	s.DeliveryOssEnabled = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliveryOssKeyPrefix(v string) *SetServiceSettingsRequest {
	s.DeliveryOssKeyPrefix = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliverySlsEnabled(v bool) *SetServiceSettingsRequest {
	s.DeliverySlsEnabled = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliverySlsProjectName(v string) *SetServiceSettingsRequest {
	s.DeliverySlsProjectName = &v
	return s
}

func (s *SetServiceSettingsRequest) SetRdcEnterpriseId(v string) *SetServiceSettingsRequest {
	s.RdcEnterpriseId = &v
	return s
}

func (s *SetServiceSettingsRequest) SetRegionId(v string) *SetServiceSettingsRequest {
	s.RegionId = &v
	return s
}

func (s *SetServiceSettingsRequest) Validate() error {
	return dara.Validate(s)
}

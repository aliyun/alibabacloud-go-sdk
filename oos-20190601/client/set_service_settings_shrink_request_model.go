// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetServiceSettingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryOssBucketName(v string) *SetServiceSettingsShrinkRequest
	GetDeliveryOssBucketName() *string
	SetDeliveryOssEnabled(v bool) *SetServiceSettingsShrinkRequest
	GetDeliveryOssEnabled() *bool
	SetDeliveryOssKeyPrefix(v string) *SetServiceSettingsShrinkRequest
	GetDeliveryOssKeyPrefix() *string
	SetDeliverySlsEnabled(v bool) *SetServiceSettingsShrinkRequest
	GetDeliverySlsEnabled() *bool
	SetDeliverySlsProjectName(v string) *SetServiceSettingsShrinkRequest
	GetDeliverySlsProjectName() *string
	SetRdFolderIdsShrink(v string) *SetServiceSettingsShrinkRequest
	GetRdFolderIdsShrink() *string
	SetRdcEnterpriseId(v string) *SetServiceSettingsShrinkRequest
	GetRdcEnterpriseId() *string
	SetRegionId(v string) *SetServiceSettingsShrinkRequest
	GetRegionId() *string
	SetServiceAccessRdEnabled(v bool) *SetServiceSettingsShrinkRequest
	GetServiceAccessRdEnabled() *bool
}

type SetServiceSettingsShrinkRequest struct {
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
	RdFolderIdsShrink      *string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty"`
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
	// if can be null:
	// true
	ServiceAccessRdEnabled *bool `json:"ServiceAccessRdEnabled,omitempty" xml:"ServiceAccessRdEnabled,omitempty"`
}

func (s SetServiceSettingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetServiceSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsShrinkRequest) GetDeliveryOssBucketName() *string {
	return s.DeliveryOssBucketName
}

func (s *SetServiceSettingsShrinkRequest) GetDeliveryOssEnabled() *bool {
	return s.DeliveryOssEnabled
}

func (s *SetServiceSettingsShrinkRequest) GetDeliveryOssKeyPrefix() *string {
	return s.DeliveryOssKeyPrefix
}

func (s *SetServiceSettingsShrinkRequest) GetDeliverySlsEnabled() *bool {
	return s.DeliverySlsEnabled
}

func (s *SetServiceSettingsShrinkRequest) GetDeliverySlsProjectName() *string {
	return s.DeliverySlsProjectName
}

func (s *SetServiceSettingsShrinkRequest) GetRdFolderIdsShrink() *string {
	return s.RdFolderIdsShrink
}

func (s *SetServiceSettingsShrinkRequest) GetRdcEnterpriseId() *string {
	return s.RdcEnterpriseId
}

func (s *SetServiceSettingsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetServiceSettingsShrinkRequest) GetServiceAccessRdEnabled() *bool {
	return s.ServiceAccessRdEnabled
}

func (s *SetServiceSettingsShrinkRequest) SetDeliveryOssBucketName(v string) *SetServiceSettingsShrinkRequest {
	s.DeliveryOssBucketName = &v
	return s
}

func (s *SetServiceSettingsShrinkRequest) SetDeliveryOssEnabled(v bool) *SetServiceSettingsShrinkRequest {
	s.DeliveryOssEnabled = &v
	return s
}

func (s *SetServiceSettingsShrinkRequest) SetDeliveryOssKeyPrefix(v string) *SetServiceSettingsShrinkRequest {
	s.DeliveryOssKeyPrefix = &v
	return s
}

func (s *SetServiceSettingsShrinkRequest) SetDeliverySlsEnabled(v bool) *SetServiceSettingsShrinkRequest {
	s.DeliverySlsEnabled = &v
	return s
}

func (s *SetServiceSettingsShrinkRequest) SetDeliverySlsProjectName(v string) *SetServiceSettingsShrinkRequest {
	s.DeliverySlsProjectName = &v
	return s
}

func (s *SetServiceSettingsShrinkRequest) SetRdFolderIdsShrink(v string) *SetServiceSettingsShrinkRequest {
	s.RdFolderIdsShrink = &v
	return s
}

func (s *SetServiceSettingsShrinkRequest) SetRdcEnterpriseId(v string) *SetServiceSettingsShrinkRequest {
	s.RdcEnterpriseId = &v
	return s
}

func (s *SetServiceSettingsShrinkRequest) SetRegionId(v string) *SetServiceSettingsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *SetServiceSettingsShrinkRequest) SetServiceAccessRdEnabled(v bool) *SetServiceSettingsShrinkRequest {
	s.ServiceAccessRdEnabled = &v
	return s
}

func (s *SetServiceSettingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

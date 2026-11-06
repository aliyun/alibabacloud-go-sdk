// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationVersion(v int64) *UpdateDataStorageRequest
	GetConfigurationVersion() *int64
	SetDataStorageRegionId(v string) *UpdateDataStorageRequest
	GetDataStorageRegionId() *string
	SetDeliveryStatus(v string) *UpdateDataStorageRequest
	GetDeliveryStatus() *string
	SetLang(v string) *UpdateDataStorageRequest
	GetLang() *string
	SetMultiAccountDeliveryEnabled(v bool) *UpdateDataStorageRequest
	GetMultiAccountDeliveryEnabled() *bool
	SetRegionId(v string) *UpdateDataStorageRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataStorageRequest
	GetRoleFor() *int64
	SetTargetProject(v string) *UpdateDataStorageRequest
	GetTargetProject() *string
}

type UpdateDataStorageRequest struct {
	// The configuration version obtained from the query when modifying multi-account delivery.
	//
	// example:
	//
	// 1
	ConfigurationVersion *int64 `json:"ConfigurationVersion,omitempty" xml:"ConfigurationVersion,omitempty"`
	// The region for log storage.
	//
	// example:
	//
	// cn-shanghai
	DataStorageRegionId *string `json:"DataStorageRegionId,omitempty" xml:"DataStorageRegionId,omitempty"`
	// The global delivery switch for log management. This parameter is not available for configuration. Valid values:
	//
	// - enable: Enables all delivery.
	//
	// - disable: Disables all delivery.
	//
	// example:
	//
	// enable
	DeliveryStatus *string `json:"DeliveryStatus,omitempty" xml:"DeliveryStatus,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required. If you do not specify this parameter, the API returns a 400 error: "The specified Lang can not be empty." Specify zh or en.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The global switch for multi-account centralized delivery. If this parameter is omitted, the original storage update behavior is maintained. Do not mix this parameter with the storage region or the original delivery status.
	//
	// example:
	//
	// true
	MultiAccountDeliveryEnabled *bool `json:"MultiAccountDeliveryEnabled,omitempty" xml:"MultiAccountDeliveryEnabled,omitempty"`
	// The region where the data management center for threat analysis is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The current target project obtained from the query when modifying multi-account delivery.
	//
	// example:
	//
	// aliyun-cloudsiem-data-1234567890123456-cn-shanghai
	TargetProject *string `json:"TargetProject,omitempty" xml:"TargetProject,omitempty"`
}

func (s UpdateDataStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageRequest) GetConfigurationVersion() *int64 {
	return s.ConfigurationVersion
}

func (s *UpdateDataStorageRequest) GetDataStorageRegionId() *string {
	return s.DataStorageRegionId
}

func (s *UpdateDataStorageRequest) GetDeliveryStatus() *string {
	return s.DeliveryStatus
}

func (s *UpdateDataStorageRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataStorageRequest) GetMultiAccountDeliveryEnabled() *bool {
	return s.MultiAccountDeliveryEnabled
}

func (s *UpdateDataStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataStorageRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataStorageRequest) GetTargetProject() *string {
	return s.TargetProject
}

func (s *UpdateDataStorageRequest) SetConfigurationVersion(v int64) *UpdateDataStorageRequest {
	s.ConfigurationVersion = &v
	return s
}

func (s *UpdateDataStorageRequest) SetDataStorageRegionId(v string) *UpdateDataStorageRequest {
	s.DataStorageRegionId = &v
	return s
}

func (s *UpdateDataStorageRequest) SetDeliveryStatus(v string) *UpdateDataStorageRequest {
	s.DeliveryStatus = &v
	return s
}

func (s *UpdateDataStorageRequest) SetLang(v string) *UpdateDataStorageRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataStorageRequest) SetMultiAccountDeliveryEnabled(v bool) *UpdateDataStorageRequest {
	s.MultiAccountDeliveryEnabled = &v
	return s
}

func (s *UpdateDataStorageRequest) SetRegionId(v string) *UpdateDataStorageRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataStorageRequest) SetRoleFor(v int64) *UpdateDataStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataStorageRequest) SetTargetProject(v string) *UpdateDataStorageRequest {
	s.TargetProject = &v
	return s
}

func (s *UpdateDataStorageRequest) Validate() error {
	return dara.Validate(s)
}

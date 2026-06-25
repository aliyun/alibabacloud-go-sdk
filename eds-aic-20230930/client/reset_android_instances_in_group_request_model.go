// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAndroidInstancesInGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *ResetAndroidInstancesInGroupRequest
	GetAndroidInstanceIds() []*string
	SetIgnoreParamValidation(v bool) *ResetAndroidInstancesInGroupRequest
	GetIgnoreParamValidation() *bool
	SetSaleMode(v string) *ResetAndroidInstancesInGroupRequest
	GetSaleMode() *string
	SetSettingResetType(v int32) *ResetAndroidInstancesInGroupRequest
	GetSettingResetType() *int32
}

type ResetAndroidInstancesInGroupRequest struct {
	// A list of instance IDs.
	AndroidInstanceIds    []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	IgnoreParamValidation *bool     `json:"IgnoreParamValidation,omitempty" xml:"IgnoreParamValidation,omitempty"`
	// The sale mode. This parameter is deprecated.
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// <props="china">Specifies whether to retain the property settings when you reset the instances. By default, the property settings are not retained. This parameter applies only to cloud phone matrix instances. Run the wya dump config command to view the details of the retained properties.<props="intl">This parameter is not supported on the international site (alibabacloud.com).
	//
	// example:
	//
	// 1
	SettingResetType *int32 `json:"SettingResetType,omitempty" xml:"SettingResetType,omitempty"`
}

func (s ResetAndroidInstancesInGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAndroidInstancesInGroupRequest) GoString() string {
	return s.String()
}

func (s *ResetAndroidInstancesInGroupRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *ResetAndroidInstancesInGroupRequest) GetIgnoreParamValidation() *bool {
	return s.IgnoreParamValidation
}

func (s *ResetAndroidInstancesInGroupRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *ResetAndroidInstancesInGroupRequest) GetSettingResetType() *int32 {
	return s.SettingResetType
}

func (s *ResetAndroidInstancesInGroupRequest) SetAndroidInstanceIds(v []*string) *ResetAndroidInstancesInGroupRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) SetIgnoreParamValidation(v bool) *ResetAndroidInstancesInGroupRequest {
	s.IgnoreParamValidation = &v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) SetSaleMode(v string) *ResetAndroidInstancesInGroupRequest {
	s.SaleMode = &v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) SetSettingResetType(v int32) *ResetAndroidInstancesInGroupRequest {
	s.SettingResetType = &v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) Validate() error {
	return dara.Validate(s)
}

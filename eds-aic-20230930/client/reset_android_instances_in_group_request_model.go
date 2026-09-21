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
	SetAutoPay(v bool) *ResetAndroidInstancesInGroupRequest
	GetAutoPay() *bool
	SetIgnoreParamValidation(v bool) *ResetAndroidInstancesInGroupRequest
	GetIgnoreParamValidation() *bool
	SetPromotionId(v string) *ResetAndroidInstancesInGroupRequest
	GetPromotionId() *string
	SetSaleMode(v string) *ResetAndroidInstancesInGroupRequest
	GetSaleMode() *string
	SetSettingResetType(v int32) *ResetAndroidInstancesInGroupRequest
	GetSettingResetType() *int32
	SetTargetDataDiskSize(v int32) *ResetAndroidInstancesInGroupRequest
	GetTargetDataDiskSize() *int32
}

type ResetAndroidInstancesInGroupRequest struct {
	// The list of instance IDs.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to enable automatic payment. Default value: false.
	//
	// example:
	//
	// true
	AutoPay               *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	IgnoreParamValidation *bool `json:"IgnoreParamValidation,omitempty" xml:"IgnoreParamValidation,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// **[Deprecated]*	- The sales mode. This parameter is deprecated.
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// <props="china">Specifies whether to retain attribute settings during the reset. If this parameter is not specified, attribute configurations are not retained by default. This parameter takes effect only for cloud phone matrix instances. Run the wya dump config command to view the details of retained attributes.
	//
	// <props="intl">This parameter is not supported on the international site.
	//
	// example:
	//
	// 1
	SettingResetType *int32 `json:"SettingResetType,omitempty" xml:"SettingResetType,omitempty"`
	// Specify this parameter when you need to reduce storage while resetting instances in a cloud phone matrix. This feature is currently available through a whitelist. This parameter applies only to instances in a cloud phone matrix.
	//
	// example:
	//
	// 10
	TargetDataDiskSize *int32 `json:"TargetDataDiskSize,omitempty" xml:"TargetDataDiskSize,omitempty"`
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

func (s *ResetAndroidInstancesInGroupRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ResetAndroidInstancesInGroupRequest) GetIgnoreParamValidation() *bool {
	return s.IgnoreParamValidation
}

func (s *ResetAndroidInstancesInGroupRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *ResetAndroidInstancesInGroupRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *ResetAndroidInstancesInGroupRequest) GetSettingResetType() *int32 {
	return s.SettingResetType
}

func (s *ResetAndroidInstancesInGroupRequest) GetTargetDataDiskSize() *int32 {
	return s.TargetDataDiskSize
}

func (s *ResetAndroidInstancesInGroupRequest) SetAndroidInstanceIds(v []*string) *ResetAndroidInstancesInGroupRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) SetAutoPay(v bool) *ResetAndroidInstancesInGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) SetIgnoreParamValidation(v bool) *ResetAndroidInstancesInGroupRequest {
	s.IgnoreParamValidation = &v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) SetPromotionId(v string) *ResetAndroidInstancesInGroupRequest {
	s.PromotionId = &v
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

func (s *ResetAndroidInstancesInGroupRequest) SetTargetDataDiskSize(v int32) *ResetAndroidInstancesInGroupRequest {
	s.TargetDataDiskSize = &v
	return s
}

func (s *ResetAndroidInstancesInGroupRequest) Validate() error {
	return dara.Validate(s)
}

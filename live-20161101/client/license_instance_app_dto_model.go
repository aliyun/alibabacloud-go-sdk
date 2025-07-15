// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLicenseInstanceAppDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *LicenseInstanceAppDTO
	GetAppId() *string
	SetBeginOn(v string) *LicenseInstanceAppDTO
	GetBeginOn() *string
	SetContractNo(v string) *LicenseInstanceAppDTO
	GetContractNo() *string
	SetCreationTime(v string) *LicenseInstanceAppDTO
	GetCreationTime() *string
	SetExpiredOn(v string) *LicenseInstanceAppDTO
	GetExpiredOn() *string
	SetInstanceId(v string) *LicenseInstanceAppDTO
	GetInstanceId() *string
	SetItemId(v string) *LicenseInstanceAppDTO
	GetItemId() *string
	SetLicenseConfigs(v []*LicenseInstanceAppDTOLicenseConfigs) *LicenseInstanceAppDTO
	GetLicenseConfigs() []*LicenseInstanceAppDTOLicenseConfigs
	SetModificationTime(v string) *LicenseInstanceAppDTO
	GetModificationTime() *string
	SetStatus(v string) *LicenseInstanceAppDTO
	GetStatus() *string
	SetUserId(v int64) *LicenseInstanceAppDTO
	GetUserId() *int64
}

type LicenseInstanceAppDTO struct {
	AppId            *string                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BeginOn          *string                                `json:"BeginOn,omitempty" xml:"BeginOn,omitempty"`
	ContractNo       *string                                `json:"ContractNo,omitempty" xml:"ContractNo,omitempty"`
	CreationTime     *string                                `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ExpiredOn        *string                                `json:"ExpiredOn,omitempty" xml:"ExpiredOn,omitempty"`
	InstanceId       *string                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemId           *string                                `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LicenseConfigs   []*LicenseInstanceAppDTOLicenseConfigs `json:"LicenseConfigs,omitempty" xml:"LicenseConfigs,omitempty" type:"Repeated"`
	ModificationTime *string                                `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Status           *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId           *int64                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LicenseInstanceAppDTO) String() string {
	return dara.Prettify(s)
}

func (s LicenseInstanceAppDTO) GoString() string {
	return s.String()
}

func (s *LicenseInstanceAppDTO) GetAppId() *string {
	return s.AppId
}

func (s *LicenseInstanceAppDTO) GetBeginOn() *string {
	return s.BeginOn
}

func (s *LicenseInstanceAppDTO) GetContractNo() *string {
	return s.ContractNo
}

func (s *LicenseInstanceAppDTO) GetCreationTime() *string {
	return s.CreationTime
}

func (s *LicenseInstanceAppDTO) GetExpiredOn() *string {
	return s.ExpiredOn
}

func (s *LicenseInstanceAppDTO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LicenseInstanceAppDTO) GetItemId() *string {
	return s.ItemId
}

func (s *LicenseInstanceAppDTO) GetLicenseConfigs() []*LicenseInstanceAppDTOLicenseConfigs {
	return s.LicenseConfigs
}

func (s *LicenseInstanceAppDTO) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *LicenseInstanceAppDTO) GetStatus() *string {
	return s.Status
}

func (s *LicenseInstanceAppDTO) GetUserId() *int64 {
	return s.UserId
}

func (s *LicenseInstanceAppDTO) SetAppId(v string) *LicenseInstanceAppDTO {
	s.AppId = &v
	return s
}

func (s *LicenseInstanceAppDTO) SetBeginOn(v string) *LicenseInstanceAppDTO {
	s.BeginOn = &v
	return s
}

func (s *LicenseInstanceAppDTO) SetContractNo(v string) *LicenseInstanceAppDTO {
	s.ContractNo = &v
	return s
}

func (s *LicenseInstanceAppDTO) SetCreationTime(v string) *LicenseInstanceAppDTO {
	s.CreationTime = &v
	return s
}

func (s *LicenseInstanceAppDTO) SetExpiredOn(v string) *LicenseInstanceAppDTO {
	s.ExpiredOn = &v
	return s
}

func (s *LicenseInstanceAppDTO) SetInstanceId(v string) *LicenseInstanceAppDTO {
	s.InstanceId = &v
	return s
}

func (s *LicenseInstanceAppDTO) SetItemId(v string) *LicenseInstanceAppDTO {
	s.ItemId = &v
	return s
}

func (s *LicenseInstanceAppDTO) SetLicenseConfigs(v []*LicenseInstanceAppDTOLicenseConfigs) *LicenseInstanceAppDTO {
	s.LicenseConfigs = v
	return s
}

func (s *LicenseInstanceAppDTO) SetModificationTime(v string) *LicenseInstanceAppDTO {
	s.ModificationTime = &v
	return s
}

func (s *LicenseInstanceAppDTO) SetStatus(v string) *LicenseInstanceAppDTO {
	s.Status = &v
	return s
}

func (s *LicenseInstanceAppDTO) SetUserId(v int64) *LicenseInstanceAppDTO {
	s.UserId = &v
	return s
}

func (s *LicenseInstanceAppDTO) Validate() error {
	return dara.Validate(s)
}

type LicenseInstanceAppDTOLicenseConfigs struct {
	BusinessType    *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	FeatureIds      *string `json:"FeatureIds,omitempty" xml:"FeatureIds,omitempty"`
	IsTrial         *bool   `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	SdkId           *int32  `json:"SdkId,omitempty" xml:"SdkId,omitempty"`
	SdkName         *string `json:"SdkName,omitempty" xml:"SdkName,omitempty"`
	Subscription    *string `json:"Subscription,omitempty" xml:"Subscription,omitempty"`
	SubscriptionImp *string `json:"SubscriptionImp,omitempty" xml:"SubscriptionImp,omitempty"`
	SubscriptionPkg *string `json:"SubscriptionPkg,omitempty" xml:"SubscriptionPkg,omitempty"`
}

func (s LicenseInstanceAppDTOLicenseConfigs) String() string {
	return dara.Prettify(s)
}

func (s LicenseInstanceAppDTOLicenseConfigs) GoString() string {
	return s.String()
}

func (s *LicenseInstanceAppDTOLicenseConfigs) GetBusinessType() *string {
	return s.BusinessType
}

func (s *LicenseInstanceAppDTOLicenseConfigs) GetFeatureIds() *string {
	return s.FeatureIds
}

func (s *LicenseInstanceAppDTOLicenseConfigs) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *LicenseInstanceAppDTOLicenseConfigs) GetSdkId() *int32 {
	return s.SdkId
}

func (s *LicenseInstanceAppDTOLicenseConfigs) GetSdkName() *string {
	return s.SdkName
}

func (s *LicenseInstanceAppDTOLicenseConfigs) GetSubscription() *string {
	return s.Subscription
}

func (s *LicenseInstanceAppDTOLicenseConfigs) GetSubscriptionImp() *string {
	return s.SubscriptionImp
}

func (s *LicenseInstanceAppDTOLicenseConfigs) GetSubscriptionPkg() *string {
	return s.SubscriptionPkg
}

func (s *LicenseInstanceAppDTOLicenseConfigs) SetBusinessType(v string) *LicenseInstanceAppDTOLicenseConfigs {
	s.BusinessType = &v
	return s
}

func (s *LicenseInstanceAppDTOLicenseConfigs) SetFeatureIds(v string) *LicenseInstanceAppDTOLicenseConfigs {
	s.FeatureIds = &v
	return s
}

func (s *LicenseInstanceAppDTOLicenseConfigs) SetIsTrial(v bool) *LicenseInstanceAppDTOLicenseConfigs {
	s.IsTrial = &v
	return s
}

func (s *LicenseInstanceAppDTOLicenseConfigs) SetSdkId(v int32) *LicenseInstanceAppDTOLicenseConfigs {
	s.SdkId = &v
	return s
}

func (s *LicenseInstanceAppDTOLicenseConfigs) SetSdkName(v string) *LicenseInstanceAppDTOLicenseConfigs {
	s.SdkName = &v
	return s
}

func (s *LicenseInstanceAppDTOLicenseConfigs) SetSubscription(v string) *LicenseInstanceAppDTOLicenseConfigs {
	s.Subscription = &v
	return s
}

func (s *LicenseInstanceAppDTOLicenseConfigs) SetSubscriptionImp(v string) *LicenseInstanceAppDTOLicenseConfigs {
	s.SubscriptionImp = &v
	return s
}

func (s *LicenseInstanceAppDTOLicenseConfigs) SetSubscriptionPkg(v string) *LicenseInstanceAppDTOLicenseConfigs {
	s.SubscriptionPkg = &v
	return s
}

func (s *LicenseInstanceAppDTOLicenseConfigs) Validate() error {
	return dara.Validate(s)
}

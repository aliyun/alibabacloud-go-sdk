// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAckAdmin(v string) *CreateAIDBClusterRequest
	GetAckAdmin() *string
	SetAutoRenew(v string) *CreateAIDBClusterRequest
	GetAutoRenew() *string
	SetAutoUseCoupon(v bool) *CreateAIDBClusterRequest
	GetAutoUseCoupon() *bool
	SetDBClusterDescription(v string) *CreateAIDBClusterRequest
	GetDBClusterDescription() *string
	SetDBClusterId(v string) *CreateAIDBClusterRequest
	GetDBClusterId() *string
	SetDBNodeClass(v string) *CreateAIDBClusterRequest
	GetDBNodeClass() *string
	SetExtension(v string) *CreateAIDBClusterRequest
	GetExtension() *string
	SetInferenceEngine(v string) *CreateAIDBClusterRequest
	GetInferenceEngine() *string
	SetKubeClusterId(v string) *CreateAIDBClusterRequest
	GetKubeClusterId() *string
	SetKubeConfig(v string) *CreateAIDBClusterRequest
	GetKubeConfig() *string
	SetKubeManagement(v string) *CreateAIDBClusterRequest
	GetKubeManagement() *string
	SetKubeType(v string) *CreateAIDBClusterRequest
	GetKubeType() *string
	SetKubernetesConfig(v string) *CreateAIDBClusterRequest
	GetKubernetesConfig() *string
	SetManagementMode(v string) *CreateAIDBClusterRequest
	GetManagementMode() *string
	SetModeName(v string) *CreateAIDBClusterRequest
	GetModeName() *string
	SetOwnerAccount(v string) *CreateAIDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAIDBClusterRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateAIDBClusterRequest
	GetPassword() *string
	SetPayType(v string) *CreateAIDBClusterRequest
	GetPayType() *string
	SetPeriod(v string) *CreateAIDBClusterRequest
	GetPeriod() *string
	SetPromotionCode(v string) *CreateAIDBClusterRequest
	GetPromotionCode() *string
	SetRegionId(v string) *CreateAIDBClusterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateAIDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAIDBClusterRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *CreateAIDBClusterRequest
	GetSecurityGroupId() *string
	SetStorageSpace(v int32) *CreateAIDBClusterRequest
	GetStorageSpace() *int32
	SetStorageType(v string) *CreateAIDBClusterRequest
	GetStorageType() *string
	SetTimeSlices(v []*CreateAIDBClusterRequestTimeSlices) *CreateAIDBClusterRequest
	GetTimeSlices() []*CreateAIDBClusterRequestTimeSlices
	SetUsedTime(v string) *CreateAIDBClusterRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreateAIDBClusterRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateAIDBClusterRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateAIDBClusterRequest
	GetZoneId() *string
}

type CreateAIDBClusterRequest struct {
	// example:
	//
	// yes
	AckAdmin *string `json:"AckAdmin,omitempty" xml:"AckAdmin,omitempty"`
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// false
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// example:
	//
	// pc-******************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// polar.pg.g4.6xlarge.gu4
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// maas
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// sglang
	InferenceEngine *string `json:"InferenceEngine,omitempty" xml:"InferenceEngine,omitempty"`
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxxxxx
	KubeClusterId *string `json:"KubeClusterId,omitempty" xml:"KubeClusterId,omitempty"`
	// example:
	//
	// xxx
	KubeConfig *string `json:"KubeConfig,omitempty" xml:"KubeConfig,omitempty"`
	// example:
	//
	// self_k8s
	KubeManagement *string `json:"KubeManagement,omitempty" xml:"KubeManagement,omitempty"`
	// aideploy
	//
	// example:
	//
	// aideploy
	KubeType *string `json:"KubeType,omitempty" xml:"KubeType,omitempty"`
	// example:
	//
	// xxx
	KubernetesConfig *string `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty"`
	// example:
	//
	// ack
	ManagementMode *string `json:"ManagementMode,omitempty" xml:"ManagementMode,omitempty"`
	// example:
	//
	// Qwen3-30B-A3B
	ModeName     *string `json:"ModeName,omitempty" xml:"ModeName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// xxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 5
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-bp**************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// 1024
	StorageSpace *int32 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// essdpl0
	StorageType *string                               `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	TimeSlices  []*CreateAIDBClusterRequestTimeSlices `json:"TimeSlices,omitempty" xml:"TimeSlices,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-**********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateAIDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterRequest) GetAckAdmin() *string {
	return s.AckAdmin
}

func (s *CreateAIDBClusterRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateAIDBClusterRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateAIDBClusterRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *CreateAIDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAIDBClusterRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *CreateAIDBClusterRequest) GetExtension() *string {
	return s.Extension
}

func (s *CreateAIDBClusterRequest) GetInferenceEngine() *string {
	return s.InferenceEngine
}

func (s *CreateAIDBClusterRequest) GetKubeClusterId() *string {
	return s.KubeClusterId
}

func (s *CreateAIDBClusterRequest) GetKubeConfig() *string {
	return s.KubeConfig
}

func (s *CreateAIDBClusterRequest) GetKubeManagement() *string {
	return s.KubeManagement
}

func (s *CreateAIDBClusterRequest) GetKubeType() *string {
	return s.KubeType
}

func (s *CreateAIDBClusterRequest) GetKubernetesConfig() *string {
	return s.KubernetesConfig
}

func (s *CreateAIDBClusterRequest) GetManagementMode() *string {
	return s.ManagementMode
}

func (s *CreateAIDBClusterRequest) GetModeName() *string {
	return s.ModeName
}

func (s *CreateAIDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAIDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAIDBClusterRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateAIDBClusterRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateAIDBClusterRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateAIDBClusterRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateAIDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAIDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAIDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAIDBClusterRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateAIDBClusterRequest) GetStorageSpace() *int32 {
	return s.StorageSpace
}

func (s *CreateAIDBClusterRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateAIDBClusterRequest) GetTimeSlices() []*CreateAIDBClusterRequestTimeSlices {
	return s.TimeSlices
}

func (s *CreateAIDBClusterRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateAIDBClusterRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateAIDBClusterRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateAIDBClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateAIDBClusterRequest) SetAckAdmin(v string) *CreateAIDBClusterRequest {
	s.AckAdmin = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetAutoRenew(v string) *CreateAIDBClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetAutoUseCoupon(v bool) *CreateAIDBClusterRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetDBClusterDescription(v string) *CreateAIDBClusterRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetDBClusterId(v string) *CreateAIDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetDBNodeClass(v string) *CreateAIDBClusterRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetExtension(v string) *CreateAIDBClusterRequest {
	s.Extension = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetInferenceEngine(v string) *CreateAIDBClusterRequest {
	s.InferenceEngine = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetKubeClusterId(v string) *CreateAIDBClusterRequest {
	s.KubeClusterId = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetKubeConfig(v string) *CreateAIDBClusterRequest {
	s.KubeConfig = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetKubeManagement(v string) *CreateAIDBClusterRequest {
	s.KubeManagement = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetKubeType(v string) *CreateAIDBClusterRequest {
	s.KubeType = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetKubernetesConfig(v string) *CreateAIDBClusterRequest {
	s.KubernetesConfig = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetManagementMode(v string) *CreateAIDBClusterRequest {
	s.ManagementMode = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetModeName(v string) *CreateAIDBClusterRequest {
	s.ModeName = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetOwnerAccount(v string) *CreateAIDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetOwnerId(v int64) *CreateAIDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetPassword(v string) *CreateAIDBClusterRequest {
	s.Password = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetPayType(v string) *CreateAIDBClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetPeriod(v string) *CreateAIDBClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetPromotionCode(v string) *CreateAIDBClusterRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetRegionId(v string) *CreateAIDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetResourceOwnerAccount(v string) *CreateAIDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetResourceOwnerId(v int64) *CreateAIDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetSecurityGroupId(v string) *CreateAIDBClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetStorageSpace(v int32) *CreateAIDBClusterRequest {
	s.StorageSpace = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetStorageType(v string) *CreateAIDBClusterRequest {
	s.StorageType = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetTimeSlices(v []*CreateAIDBClusterRequestTimeSlices) *CreateAIDBClusterRequest {
	s.TimeSlices = v
	return s
}

func (s *CreateAIDBClusterRequest) SetUsedTime(v string) *CreateAIDBClusterRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetVPCId(v string) *CreateAIDBClusterRequest {
	s.VPCId = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetVSwitchId(v string) *CreateAIDBClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateAIDBClusterRequest) SetZoneId(v string) *CreateAIDBClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateAIDBClusterRequest) Validate() error {
	if s.TimeSlices != nil {
		for _, item := range s.TimeSlices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAIDBClusterRequestTimeSlices struct {
	// example:
	//
	// 1758729600
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 1758733200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s CreateAIDBClusterRequestTimeSlices) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterRequestTimeSlices) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterRequestTimeSlices) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *CreateAIDBClusterRequestTimeSlices) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateAIDBClusterRequestTimeSlices) SetBeginTime(v int64) *CreateAIDBClusterRequestTimeSlices {
	s.BeginTime = &v
	return s
}

func (s *CreateAIDBClusterRequestTimeSlices) SetEndTime(v int64) *CreateAIDBClusterRequestTimeSlices {
	s.EndTime = &v
	return s
}

func (s *CreateAIDBClusterRequestTimeSlices) Validate() error {
	return dara.Validate(s)
}

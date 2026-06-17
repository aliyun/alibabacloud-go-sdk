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
	SetClientToken(v string) *CreateAIDBClusterRequest
	GetClientToken() *string
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
	SetModelName(v string) *CreateAIDBClusterRequest
	GetModelName() *string
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
	// Specifies whether the cluster is managed by an ACK cluster.
	//
	// example:
	//
	// yes
	AckAdmin *string `json:"AckAdmin,omitempty" xml:"AckAdmin,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - **true**: Auto-renewal is enabled.
	//
	// - **false**: Auto-renewal is disabled.
	//
	// Default value: **false**.
	//
	// > This parameter takes effect only when **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to automatically use a coupon. Valid values:
	//
	// example:
	//
	// false
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the cluster. You can use the description to perform a fuzzy search.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the PolarDB cluster that the application depends on.
	//
	// example:
	//
	// pc-******************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node specification.
	//
	// This parameter is required.
	//
	// example:
	//
	// polar.pg.g4.6xlarge.gu4
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The extension.
	//
	// example:
	//
	// maas
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The inference engine.
	//
	// example:
	//
	// sglang
	InferenceEngine *string `json:"InferenceEngine,omitempty" xml:"InferenceEngine,omitempty"`
	// The Container Service for Kubernetes (ACK) cluster ID.
	//
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxxxxx
	KubeClusterId *string `json:"KubeClusterId,omitempty" xml:"KubeClusterId,omitempty"`
	// The Kubernetes configuration.
	//
	// example:
	//
	// xxx
	KubeConfig *string `json:"KubeConfig,omitempty" xml:"KubeConfig,omitempty"`
	// The management mode of the Kubernetes cluster.
	//
	// example:
	//
	// self_k8s
	KubeManagement *string `json:"KubeManagement,omitempty" xml:"KubeManagement,omitempty"`
	// The type of the Kubernetes deployment.
	//
	// example:
	//
	// aideploy
	KubeType *string `json:"KubeType,omitempty" xml:"KubeType,omitempty"`
	// The Kubernetes configuration.
	//
	// example:
	//
	// xxx
	KubernetesConfig *string `json:"KubernetesConfig,omitempty" xml:"KubernetesConfig,omitempty"`
	// The management mode.
	//
	// example:
	//
	// ack
	ManagementMode *string `json:"ManagementMode,omitempty" xml:"ManagementMode,omitempty"`
	// example:
	//
	// Qwen3-30B-A3B
	ModelName    *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password.
	//
	// example:
	//
	// xxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. This parameter is required if **PayType*	- is set to **Prepaid**. Valid values:
	//
	// - **Year**
	//
	// - **Month**
	//
	// example:
	//
	// 5
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The coupon code. If you do not specify this parameter, the default coupon is used.
	//
	// - true (default): Use a coupon.
	//
	// - false: Do not use a coupon.
	//
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp**************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The storage space. Unit: GB.
	//
	// example:
	//
	// 1024
	StorageSpace *int32 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// The storage type.
	//
	// example:
	//
	// essdpl0
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The billing intervals for the pay-as-you-go cluster.
	TimeSlices []*CreateAIDBClusterRequestTimeSlices `json:"TimeSlices,omitempty" xml:"TimeSlices,omitempty" type:"Repeated"`
	// The subscription duration. This parameter is required if **PayType*	- is set to **Prepaid**.
	//
	// - If **Period*	- is set to **Month**, the value of **UsedTime*	- must be an integer from `[1-9]`.
	//
	// - If **Period*	- is set to **Year**, the value of **UsedTime*	- must be an integer from `[1-3]`.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-**********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The availability zone ID.
	//
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

func (s *CreateAIDBClusterRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *CreateAIDBClusterRequest) GetModelName() *string {
	return s.ModelName
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

func (s *CreateAIDBClusterRequest) SetClientToken(v string) *CreateAIDBClusterRequest {
	s.ClientToken = &v
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

func (s *CreateAIDBClusterRequest) SetModelName(v string) *CreateAIDBClusterRequest {
	s.ModelName = &v
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
	// The start time of the billing interval. The time is in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 1758729600
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end time of the billing interval, which must be later than the start time. The time is in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
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

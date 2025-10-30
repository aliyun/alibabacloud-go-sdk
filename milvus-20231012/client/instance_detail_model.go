// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBackup(v bool) *InstanceDetail
	GetAutoBackup() *bool
	SetComponents(v []*InstanceDetailComponents) *InstanceDetail
	GetComponents() []*InstanceDetailComponents
	SetConfiguration(v string) *InstanceDetail
	GetConfiguration() *string
	SetCreateTime(v string) *InstanceDetail
	GetCreateTime() *string
	SetDbVersion(v string) *InstanceDetail
	GetDbVersion() *string
	SetEncrypted(v bool) *InstanceDetail
	GetEncrypted() *bool
	SetExpireTime(v string) *InstanceDetail
	GetExpireTime() *string
	SetHa(v bool) *InstanceDetail
	GetHa() *bool
	SetInstanceId(v string) *InstanceDetail
	GetInstanceId() *string
	SetInstanceName(v string) *InstanceDetail
	GetInstanceName() *string
	SetKmsKeyId(v string) *InstanceDetail
	GetKmsKeyId() *string
	SetMultiZoneMode(v string) *InstanceDetail
	GetMultiZoneMode() *string
	SetOrderId(v string) *InstanceDetail
	GetOrderId() *string
	SetPaymentType(v string) *InstanceDetail
	GetPaymentType() *string
	SetRegionId(v string) *InstanceDetail
	GetRegionId() *string
	SetResourceGroupId(v string) *InstanceDetail
	GetResourceGroupId() *string
	SetRunningTime(v int64) *InstanceDetail
	GetRunningTime() *int64
	SetSecurityGroupIds(v []*string) *InstanceDetail
	GetSecurityGroupIds() []*string
	SetStatus(v string) *InstanceDetail
	GetStatus() *string
	SetTags(v []*InstanceDetailTags) *InstanceDetail
	GetTags() []*InstanceDetailTags
	SetVSwitchIds(v []*InstanceDetailVSwitchIds) *InstanceDetail
	GetVSwitchIds() []*InstanceDetailVSwitchIds
	SetVpcId(v string) *InstanceDetail
	GetVpcId() *string
	SetZoneId(v string) *InstanceDetail
	GetZoneId() *string
}

type InstanceDetail struct {
	AutoBackup    *bool                       `json:"autoBackup,omitempty" xml:"autoBackup,omitempty"`
	Components    []*InstanceDetailComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	Configuration *string                     `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DbVersion  *string `json:"dbVersion,omitempty" xml:"dbVersion,omitempty"`
	Encrypted  *bool   `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	Ha         *bool   `json:"ha,omitempty" xml:"ha,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// milvus-test
	InstanceName     *string                     `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	KmsKeyId         *string                     `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
	MultiZoneMode    *string                     `json:"multiZoneMode,omitempty" xml:"multiZoneMode,omitempty"`
	OrderId          *string                     `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PaymentType      *string                     `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	RegionId         *string                     `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ResourceGroupId  *string                     `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	RunningTime      *int64                      `json:"runningTime,omitempty" xml:"runningTime,omitempty"`
	SecurityGroupIds []*string                   `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
	Status           *string                     `json:"status,omitempty" xml:"status,omitempty"`
	Tags             []*InstanceDetailTags       `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	VSwitchIds       []*InstanceDetailVSwitchIds `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId            *string                     `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	ZoneId           *string                     `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s InstanceDetail) String() string {
	return dara.Prettify(s)
}

func (s InstanceDetail) GoString() string {
	return s.String()
}

func (s *InstanceDetail) GetAutoBackup() *bool {
	return s.AutoBackup
}

func (s *InstanceDetail) GetComponents() []*InstanceDetailComponents {
	return s.Components
}

func (s *InstanceDetail) GetConfiguration() *string {
	return s.Configuration
}

func (s *InstanceDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *InstanceDetail) GetDbVersion() *string {
	return s.DbVersion
}

func (s *InstanceDetail) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *InstanceDetail) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *InstanceDetail) GetHa() *bool {
	return s.Ha
}

func (s *InstanceDetail) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstanceDetail) GetInstanceName() *string {
	return s.InstanceName
}

func (s *InstanceDetail) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *InstanceDetail) GetMultiZoneMode() *string {
	return s.MultiZoneMode
}

func (s *InstanceDetail) GetOrderId() *string {
	return s.OrderId
}

func (s *InstanceDetail) GetPaymentType() *string {
	return s.PaymentType
}

func (s *InstanceDetail) GetRegionId() *string {
	return s.RegionId
}

func (s *InstanceDetail) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *InstanceDetail) GetRunningTime() *int64 {
	return s.RunningTime
}

func (s *InstanceDetail) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *InstanceDetail) GetStatus() *string {
	return s.Status
}

func (s *InstanceDetail) GetTags() []*InstanceDetailTags {
	return s.Tags
}

func (s *InstanceDetail) GetVSwitchIds() []*InstanceDetailVSwitchIds {
	return s.VSwitchIds
}

func (s *InstanceDetail) GetVpcId() *string {
	return s.VpcId
}

func (s *InstanceDetail) GetZoneId() *string {
	return s.ZoneId
}

func (s *InstanceDetail) SetAutoBackup(v bool) *InstanceDetail {
	s.AutoBackup = &v
	return s
}

func (s *InstanceDetail) SetComponents(v []*InstanceDetailComponents) *InstanceDetail {
	s.Components = v
	return s
}

func (s *InstanceDetail) SetConfiguration(v string) *InstanceDetail {
	s.Configuration = &v
	return s
}

func (s *InstanceDetail) SetCreateTime(v string) *InstanceDetail {
	s.CreateTime = &v
	return s
}

func (s *InstanceDetail) SetDbVersion(v string) *InstanceDetail {
	s.DbVersion = &v
	return s
}

func (s *InstanceDetail) SetEncrypted(v bool) *InstanceDetail {
	s.Encrypted = &v
	return s
}

func (s *InstanceDetail) SetExpireTime(v string) *InstanceDetail {
	s.ExpireTime = &v
	return s
}

func (s *InstanceDetail) SetHa(v bool) *InstanceDetail {
	s.Ha = &v
	return s
}

func (s *InstanceDetail) SetInstanceId(v string) *InstanceDetail {
	s.InstanceId = &v
	return s
}

func (s *InstanceDetail) SetInstanceName(v string) *InstanceDetail {
	s.InstanceName = &v
	return s
}

func (s *InstanceDetail) SetKmsKeyId(v string) *InstanceDetail {
	s.KmsKeyId = &v
	return s
}

func (s *InstanceDetail) SetMultiZoneMode(v string) *InstanceDetail {
	s.MultiZoneMode = &v
	return s
}

func (s *InstanceDetail) SetOrderId(v string) *InstanceDetail {
	s.OrderId = &v
	return s
}

func (s *InstanceDetail) SetPaymentType(v string) *InstanceDetail {
	s.PaymentType = &v
	return s
}

func (s *InstanceDetail) SetRegionId(v string) *InstanceDetail {
	s.RegionId = &v
	return s
}

func (s *InstanceDetail) SetResourceGroupId(v string) *InstanceDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *InstanceDetail) SetRunningTime(v int64) *InstanceDetail {
	s.RunningTime = &v
	return s
}

func (s *InstanceDetail) SetSecurityGroupIds(v []*string) *InstanceDetail {
	s.SecurityGroupIds = v
	return s
}

func (s *InstanceDetail) SetStatus(v string) *InstanceDetail {
	s.Status = &v
	return s
}

func (s *InstanceDetail) SetTags(v []*InstanceDetailTags) *InstanceDetail {
	s.Tags = v
	return s
}

func (s *InstanceDetail) SetVSwitchIds(v []*InstanceDetailVSwitchIds) *InstanceDetail {
	s.VSwitchIds = v
	return s
}

func (s *InstanceDetail) SetVpcId(v string) *InstanceDetail {
	s.VpcId = &v
	return s
}

func (s *InstanceDetail) SetZoneId(v string) *InstanceDetail {
	s.ZoneId = &v
	return s
}

func (s *InstanceDetail) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VSwitchIds != nil {
		for _, item := range s.VSwitchIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstanceDetailComponents struct {
	CuNum        *int32  `json:"cuNum,omitempty" xml:"cuNum,omitempty"`
	CuType       *string `json:"cuType,omitempty" xml:"cuType,omitempty"`
	DiskSizeType *string `json:"diskSizeType,omitempty" xml:"diskSizeType,omitempty"`
	Replica      *int32  `json:"replica,omitempty" xml:"replica,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InstanceDetailComponents) String() string {
	return dara.Prettify(s)
}

func (s InstanceDetailComponents) GoString() string {
	return s.String()
}

func (s *InstanceDetailComponents) GetCuNum() *int32 {
	return s.CuNum
}

func (s *InstanceDetailComponents) GetCuType() *string {
	return s.CuType
}

func (s *InstanceDetailComponents) GetDiskSizeType() *string {
	return s.DiskSizeType
}

func (s *InstanceDetailComponents) GetReplica() *int32 {
	return s.Replica
}

func (s *InstanceDetailComponents) GetType() *string {
	return s.Type
}

func (s *InstanceDetailComponents) SetCuNum(v int32) *InstanceDetailComponents {
	s.CuNum = &v
	return s
}

func (s *InstanceDetailComponents) SetCuType(v string) *InstanceDetailComponents {
	s.CuType = &v
	return s
}

func (s *InstanceDetailComponents) SetDiskSizeType(v string) *InstanceDetailComponents {
	s.DiskSizeType = &v
	return s
}

func (s *InstanceDetailComponents) SetReplica(v int32) *InstanceDetailComponents {
	s.Replica = &v
	return s
}

func (s *InstanceDetailComponents) SetType(v string) *InstanceDetailComponents {
	s.Type = &v
	return s
}

func (s *InstanceDetailComponents) Validate() error {
	return dara.Validate(s)
}

type InstanceDetailTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InstanceDetailTags) String() string {
	return dara.Prettify(s)
}

func (s InstanceDetailTags) GoString() string {
	return s.String()
}

func (s *InstanceDetailTags) GetKey() *string {
	return s.Key
}

func (s *InstanceDetailTags) GetValue() *string {
	return s.Value
}

func (s *InstanceDetailTags) SetKey(v string) *InstanceDetailTags {
	s.Key = &v
	return s
}

func (s *InstanceDetailTags) SetValue(v string) *InstanceDetailTags {
	s.Value = &v
	return s
}

func (s *InstanceDetailTags) Validate() error {
	return dara.Validate(s)
}

type InstanceDetailVSwitchIds struct {
	VswId  *string `json:"vswId,omitempty" xml:"vswId,omitempty"`
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s InstanceDetailVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s InstanceDetailVSwitchIds) GoString() string {
	return s.String()
}

func (s *InstanceDetailVSwitchIds) GetVswId() *string {
	return s.VswId
}

func (s *InstanceDetailVSwitchIds) GetZoneId() *string {
	return s.ZoneId
}

func (s *InstanceDetailVSwitchIds) SetVswId(v string) *InstanceDetailVSwitchIds {
	s.VswId = &v
	return s
}

func (s *InstanceDetailVSwitchIds) SetZoneId(v string) *InstanceDetailVSwitchIds {
	s.ZoneId = &v
	return s
}

func (s *InstanceDetailVSwitchIds) Validate() error {
	return dara.Validate(s)
}

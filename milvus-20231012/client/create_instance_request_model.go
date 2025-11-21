// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateInstanceRequest
	GetRegionId() *string
	SetAutoBackup(v bool) *CreateInstanceRequest
	GetAutoBackup() *bool
	SetComponents(v []*CreateInstanceRequestComponents) *CreateInstanceRequest
	GetComponents() []*CreateInstanceRequestComponents
	SetConfiguration(v string) *CreateInstanceRequest
	GetConfiguration() *string
	SetDbAdminPassword(v string) *CreateInstanceRequest
	GetDbAdminPassword() *string
	SetDbVersion(v string) *CreateInstanceRequest
	GetDbVersion() *string
	SetEncrypted(v bool) *CreateInstanceRequest
	GetEncrypted() *bool
	SetHa(v bool) *CreateInstanceRequest
	GetHa() *bool
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetKmsKeyId(v string) *CreateInstanceRequest
	GetKmsKeyId() *string
	SetMultiZoneMode(v string) *CreateInstanceRequest
	GetMultiZoneMode() *string
	SetPaymentDuration(v int32) *CreateInstanceRequest
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *CreateInstanceRequest
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *CreateInstanceRequest
	GetPaymentType() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest
	GetTags() []*CreateInstanceRequestTags
	SetVSwitchIds(v []*CreateInstanceRequestVSwitchIds) *CreateInstanceRequest
	GetVSwitchIds() []*CreateInstanceRequestVSwitchIds
	SetVpcId(v string) *CreateInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateInstanceRequest
	GetZoneId() *string
	SetClientToken(v string) *CreateInstanceRequest
	GetClientToken() *string
}

type CreateInstanceRequest struct {
	// example:
	//
	// cn-beijing
	RegionId      *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AutoBackup    *bool                              `json:"autoBackup,omitempty" xml:"autoBackup,omitempty"`
	Components    []*CreateInstanceRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	Configuration *string                            `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// example:
	//
	// test12
	DbAdminPassword *string `json:"dbAdminPassword,omitempty" xml:"dbAdminPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2.4
	DbVersion *string `json:"dbVersion,omitempty" xml:"dbVersion,omitempty"`
	Encrypted *bool   `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// example:
	//
	// true
	Ha *bool `json:"ha,omitempty" xml:"ha,omitempty"`
	// example:
	//
	// milvus-test
	InstanceName  *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	KmsKeyId      *string `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
	MultiZoneMode *string `json:"multiZoneMode,omitempty" xml:"multiZoneMode,omitempty"`
	// example:
	//
	// 1
	PaymentDuration *int32 `json:"paymentDuration,omitempty" xml:"paymentDuration,omitempty"`
	// example:
	//
	// month
	PaymentDurationUnit *string `json:"paymentDurationUnit,omitempty" xml:"paymentDurationUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// example:
	//
	// rg-xxx
	ResourceGroupId *string                      `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tags            []*CreateInstanceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// ["vsw-123xxx"]
	VSwitchIds []*CreateInstanceRequestVSwitchIds `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-123xxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-g
	ZoneId      *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceRequest) GetAutoBackup() *bool {
	return s.AutoBackup
}

func (s *CreateInstanceRequest) GetComponents() []*CreateInstanceRequestComponents {
	return s.Components
}

func (s *CreateInstanceRequest) GetConfiguration() *string {
	return s.Configuration
}

func (s *CreateInstanceRequest) GetDbAdminPassword() *string {
	return s.DbAdminPassword
}

func (s *CreateInstanceRequest) GetDbVersion() *string {
	return s.DbVersion
}

func (s *CreateInstanceRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateInstanceRequest) GetHa() *bool {
	return s.Ha
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateInstanceRequest) GetMultiZoneMode() *string {
	return s.MultiZoneMode
}

func (s *CreateInstanceRequest) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *CreateInstanceRequest) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *CreateInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetTags() []*CreateInstanceRequestTags {
	return s.Tags
}

func (s *CreateInstanceRequest) GetVSwitchIds() []*CreateInstanceRequestVSwitchIds {
	return s.VSwitchIds
}

func (s *CreateInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoBackup(v bool) *CreateInstanceRequest {
	s.AutoBackup = &v
	return s
}

func (s *CreateInstanceRequest) SetComponents(v []*CreateInstanceRequestComponents) *CreateInstanceRequest {
	s.Components = v
	return s
}

func (s *CreateInstanceRequest) SetConfiguration(v string) *CreateInstanceRequest {
	s.Configuration = &v
	return s
}

func (s *CreateInstanceRequest) SetDbAdminPassword(v string) *CreateInstanceRequest {
	s.DbAdminPassword = &v
	return s
}

func (s *CreateInstanceRequest) SetDbVersion(v string) *CreateInstanceRequest {
	s.DbVersion = &v
	return s
}

func (s *CreateInstanceRequest) SetEncrypted(v bool) *CreateInstanceRequest {
	s.Encrypted = &v
	return s
}

func (s *CreateInstanceRequest) SetHa(v bool) *CreateInstanceRequest {
	s.Ha = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetKmsKeyId(v string) *CreateInstanceRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateInstanceRequest) SetMultiZoneMode(v string) *CreateInstanceRequest {
	s.MultiZoneMode = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentDuration(v int32) *CreateInstanceRequest {
	s.PaymentDuration = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentDurationUnit(v string) *CreateInstanceRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchIds(v []*CreateInstanceRequestVSwitchIds) *CreateInstanceRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
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

type CreateInstanceRequestComponents struct {
	// This parameter is required.
	//
	// example:
	//
	// 8
	CuNum *int32 `json:"cuNum,omitempty" xml:"cuNum,omitempty"`
	// example:
	//
	// general
	CuType       *string `json:"cuType,omitempty" xml:"cuType,omitempty"`
	DiskSizeType *string `json:"diskSizeType,omitempty" xml:"diskSizeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// standalone
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateInstanceRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestComponents) GetCuNum() *int32 {
	return s.CuNum
}

func (s *CreateInstanceRequestComponents) GetCuType() *string {
	return s.CuType
}

func (s *CreateInstanceRequestComponents) GetDiskSizeType() *string {
	return s.DiskSizeType
}

func (s *CreateInstanceRequestComponents) GetReplica() *int32 {
	return s.Replica
}

func (s *CreateInstanceRequestComponents) GetType() *string {
	return s.Type
}

func (s *CreateInstanceRequestComponents) SetCuNum(v int32) *CreateInstanceRequestComponents {
	s.CuNum = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetCuType(v string) *CreateInstanceRequestComponents {
	s.CuType = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetDiskSizeType(v string) *CreateInstanceRequestComponents {
	s.DiskSizeType = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetReplica(v int32) *CreateInstanceRequestComponents {
	s.Replica = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetType(v string) *CreateInstanceRequestComponents {
	s.Type = &v
	return s
}

func (s *CreateInstanceRequestComponents) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTags) SetKey(v string) *CreateInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTags) SetValue(v string) *CreateInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestVSwitchIds struct {
	VswId  *string `json:"vswId,omitempty" xml:"vswId,omitempty"`
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CreateInstanceRequestVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestVSwitchIds) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestVSwitchIds) GetVswId() *string {
	return s.VswId
}

func (s *CreateInstanceRequestVSwitchIds) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequestVSwitchIds) SetVswId(v string) *CreateInstanceRequestVSwitchIds {
	s.VswId = &v
	return s
}

func (s *CreateInstanceRequestVSwitchIds) SetZoneId(v string) *CreateInstanceRequestVSwitchIds {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequestVSwitchIds) Validate() error {
	return dara.Validate(s)
}

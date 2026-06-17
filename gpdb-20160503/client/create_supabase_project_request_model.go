// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountPassword(v string) *CreateSupabaseProjectRequest
	GetAccountPassword() *string
	SetAutoScale(v bool) *CreateSupabaseProjectRequest
	GetAutoScale() *bool
	SetClientToken(v string) *CreateSupabaseProjectRequest
	GetClientToken() *string
	SetDiskPerformanceLevel(v string) *CreateSupabaseProjectRequest
	GetDiskPerformanceLevel() *string
	SetEngineVersion(v string) *CreateSupabaseProjectRequest
	GetEngineVersion() *string
	SetPayType(v string) *CreateSupabaseProjectRequest
	GetPayType() *string
	SetPeriod(v string) *CreateSupabaseProjectRequest
	GetPeriod() *string
	SetProjectName(v string) *CreateSupabaseProjectRequest
	GetProjectName() *string
	SetProjectSpec(v string) *CreateSupabaseProjectRequest
	GetProjectSpec() *string
	SetRegionId(v string) *CreateSupabaseProjectRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *CreateSupabaseProjectRequest
	GetSecurityIPList() *string
	SetStorageSize(v int64) *CreateSupabaseProjectRequest
	GetStorageSize() *int64
	SetUsedTime(v string) *CreateSupabaseProjectRequest
	GetUsedTime() *string
	SetVSwitchId(v string) *CreateSupabaseProjectRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateSupabaseProjectRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateSupabaseProjectRequest
	GetZoneId() *string
}

type CreateSupabaseProjectRequest struct {
	// The password of the initial account.
	//
	// - The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - The following special characters are supported: `!@#$%^&*()_+-=`
	//
	// - The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Pw123456
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AutoScale       *bool   `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// The client token that is used to ensure the idempotence of the request. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/327176.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88888888****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The performance level (PL) of the cloud disk. Default value: PL0. Valid values:
	//
	// - PL0
	//
	// - PL1
	//
	// example:
	//
	// PL0
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// > - If you do not specify this parameter, an instance of the Free type is created by default.
	//
	// > - If you select the subscription billing method, you can receive discounts when you purchase a one-year or longer subscription. We recommend that you select a billing method based on your business requirements.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// - **Month**: month.
	//
	// - **Year**: year.
	//
	// > This parameter is required when you create a subscription instance.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The project name. The naming rules are as follows:
	//
	// - The name must be 1 to 128 characters in length.
	//
	// - The name can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// - The name must start with a letter or an underscore (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// saas_iot_x86_modbustcp_lqt01
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The Supabase instance specification. The default specification for the Free type is 1C1G. The specifications for paid types are consistent with those available on the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1C1G
	ProjectSpec *string `json:"ProjectSpec,omitempty" xml:"ProjectSpec,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) to view the available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP address whitelist.
	//
	// The value 127.0.0.1 indicates that no external IP addresses are allowed to access the instance. After the instance is created, you can call [ModifySecurityIps](https://help.aliyun.com/document_detail/86928.html) to modify the IP address whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The storage size. Unit: GB. Default value: 1.
	//
	// example:
	//
	// 2
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The subscription duration. Valid values:
	//
	// - If **Period*	- is set to **Month**, the valid values are 1 to 11.
	//
	// - If **Period*	- is set to **Year**, the valid values are 1 to 3.
	//
	// > This parameter is required when you create a subscription instance.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The vSwitch ID.
	//
	// > - The **vSwitchId*	- parameter is required.
	//
	// > - The zone of the **vSwitch*	- must be the same as the value of **ZoneId**.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// >  - You can call [DescribeRdsVpcs](https://help.aliyun.com/document_detail/208327.html) to view the available VPC IDs.
	//
	// > - This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp*******************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// > You can call [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) to view the available zone IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateSupabaseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateSupabaseProjectRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateSupabaseProjectRequest) GetAutoScale() *bool {
	return s.AutoScale
}

func (s *CreateSupabaseProjectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSupabaseProjectRequest) GetDiskPerformanceLevel() *string {
	return s.DiskPerformanceLevel
}

func (s *CreateSupabaseProjectRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateSupabaseProjectRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateSupabaseProjectRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateSupabaseProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateSupabaseProjectRequest) GetProjectSpec() *string {
	return s.ProjectSpec
}

func (s *CreateSupabaseProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSupabaseProjectRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateSupabaseProjectRequest) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *CreateSupabaseProjectRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateSupabaseProjectRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateSupabaseProjectRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateSupabaseProjectRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateSupabaseProjectRequest) SetAccountPassword(v string) *CreateSupabaseProjectRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetAutoScale(v bool) *CreateSupabaseProjectRequest {
	s.AutoScale = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetClientToken(v string) *CreateSupabaseProjectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetDiskPerformanceLevel(v string) *CreateSupabaseProjectRequest {
	s.DiskPerformanceLevel = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetEngineVersion(v string) *CreateSupabaseProjectRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetPayType(v string) *CreateSupabaseProjectRequest {
	s.PayType = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetPeriod(v string) *CreateSupabaseProjectRequest {
	s.Period = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetProjectName(v string) *CreateSupabaseProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetProjectSpec(v string) *CreateSupabaseProjectRequest {
	s.ProjectSpec = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetRegionId(v string) *CreateSupabaseProjectRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetSecurityIPList(v string) *CreateSupabaseProjectRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetStorageSize(v int64) *CreateSupabaseProjectRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetUsedTime(v string) *CreateSupabaseProjectRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetVSwitchId(v string) *CreateSupabaseProjectRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetVpcId(v string) *CreateSupabaseProjectRequest {
	s.VpcId = &v
	return s
}

func (s *CreateSupabaseProjectRequest) SetZoneId(v string) *CreateSupabaseProjectRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateSupabaseProjectRequest) Validate() error {
	return dara.Validate(s)
}

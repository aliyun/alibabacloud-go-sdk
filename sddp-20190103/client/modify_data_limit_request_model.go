// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v int32) *ModifyDataLimitRequest
	GetAuditStatus() *int32
	SetAutoScan(v int32) *ModifyDataLimitRequest
	GetAutoScan() *int32
	SetEngineType(v string) *ModifyDataLimitRequest
	GetEngineType() *string
	SetFeatureType(v int32) *ModifyDataLimitRequest
	GetFeatureType() *int32
	SetId(v int64) *ModifyDataLimitRequest
	GetId() *int64
	SetLang(v string) *ModifyDataLimitRequest
	GetLang() *string
	SetLogStoreDay(v int32) *ModifyDataLimitRequest
	GetLogStoreDay() *int32
	SetModifyPassword(v bool) *ModifyDataLimitRequest
	GetModifyPassword() *bool
	SetPassword(v string) *ModifyDataLimitRequest
	GetPassword() *string
	SetPort(v int32) *ModifyDataLimitRequest
	GetPort() *int32
	SetResourceType(v int32) *ModifyDataLimitRequest
	GetResourceType() *int32
	SetSamplingSize(v int32) *ModifyDataLimitRequest
	GetSamplingSize() *int32
	SetSecurityGroupIdList(v []*string) *ModifyDataLimitRequest
	GetSecurityGroupIdList() []*string
	SetServiceRegionId(v string) *ModifyDataLimitRequest
	GetServiceRegionId() *string
	SetUserName(v string) *ModifyDataLimitRequest
	GetUserName() *string
	SetVSwitchIdList(v []*string) *ModifyDataLimitRequest
	GetVSwitchIdList() []*string
	SetVpcId(v string) *ModifyDataLimitRequest
	GetVpcId() *string
}

type ModifyDataLimitRequest struct {
	// Specifies whether to enable auditing. Valid values:
	//
	// - **0**: Disable auditing.
	//
	// - **1**: Enable auditing.
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Specifies whether to automatically trigger a rescan when a rule is changed. Valid values:
	//
	// - **0**: Do not automatically trigger a rescan.
	//
	// - **1**: Automatically trigger a rescan.
	//
	// > If a rescan is automatically triggered, a full scan is performed on the data asset.
	//
	// example:
	//
	// 1
	AutoScan *int32 `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	// The type of the database. Valid values:
	//
	// - **MySQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The unique ID of the data asset.
	//
	// > Call the [DescribeDataLimits](~~DescribeDataLimits~~) operation to query the ID of the data asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// - **zh**: Chinese (default)
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The retention period of raw logs after auditing is enabled. Unit: days. Valid values:
	//
	// - **30**
	//
	// - **90**
	//
	// - **180**
	//
	// - **365**
	//
	// example:
	//
	// 30
	LogStoreDay *int32 `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	// Specifies whether to modify the username and password that are used to connect to the database. Valid values:
	//
	// - **true**: Modify the username and password.
	//
	// - **false**: Do not modify the username and password.
	//
	// example:
	//
	// true
	ModifyPassword *bool `json:"ModifyPassword,omitempty" xml:"ModifyPassword,omitempty"`
	// The password for the ApsaraDB RDS database.
	//
	// example:
	//
	// ********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port used to connect to the database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The service to which the data asset belongs. Valid values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADS
	//
	// - **4**: OTS
	//
	// - **5**: RDS
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of sensitive data samples to return. This parameter takes effect only after data detection is enabled. Valid values:
	//
	// - **0**
	//
	// - **5**
	//
	// - **10**
	//
	// example:
	//
	// 0
	SamplingSize *int32 `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	// The security groups for agent-based auditing that uses PrivateLink.
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" xml:"SecurityGroupIdList,omitempty" type:"Repeated"`
	// The ID of the region in which the data asset is located. Valid values:
	//
	// - **cn-beijing**: China (Beijing)
	//
	// - **cn-zhangjiakou**: China (Zhangjiakou)
	//
	// - **cn-huhehaote**: China (Hohhot)
	//
	// - **cn-hangzhou**: China (Hangzhou)
	//
	// - **cn-shanghai**: China (Shanghai)
	//
	// - **cn-shenzhen**: China (Shenzhen)
	//
	// - **cn-hongkong**: China (Hong Kong)
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The username for the ApsaraDB RDS database.
	//
	// example:
	//
	// User01
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The vSwitches for agent-based auditing that uses PrivateLink.
	VSwitchIdList []*string `json:"VSwitchIdList,omitempty" xml:"VSwitchIdList,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the data asset belongs.
	//
	// example:
	//
	// vpc-2zevcqke6hh09c41****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyDataLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataLimitRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataLimitRequest) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *ModifyDataLimitRequest) GetAutoScan() *int32 {
	return s.AutoScan
}

func (s *ModifyDataLimitRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *ModifyDataLimitRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *ModifyDataLimitRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyDataLimitRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyDataLimitRequest) GetLogStoreDay() *int32 {
	return s.LogStoreDay
}

func (s *ModifyDataLimitRequest) GetModifyPassword() *bool {
	return s.ModifyPassword
}

func (s *ModifyDataLimitRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyDataLimitRequest) GetPort() *int32 {
	return s.Port
}

func (s *ModifyDataLimitRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *ModifyDataLimitRequest) GetSamplingSize() *int32 {
	return s.SamplingSize
}

func (s *ModifyDataLimitRequest) GetSecurityGroupIdList() []*string {
	return s.SecurityGroupIdList
}

func (s *ModifyDataLimitRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *ModifyDataLimitRequest) GetUserName() *string {
	return s.UserName
}

func (s *ModifyDataLimitRequest) GetVSwitchIdList() []*string {
	return s.VSwitchIdList
}

func (s *ModifyDataLimitRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyDataLimitRequest) SetAuditStatus(v int32) *ModifyDataLimitRequest {
	s.AuditStatus = &v
	return s
}

func (s *ModifyDataLimitRequest) SetAutoScan(v int32) *ModifyDataLimitRequest {
	s.AutoScan = &v
	return s
}

func (s *ModifyDataLimitRequest) SetEngineType(v string) *ModifyDataLimitRequest {
	s.EngineType = &v
	return s
}

func (s *ModifyDataLimitRequest) SetFeatureType(v int32) *ModifyDataLimitRequest {
	s.FeatureType = &v
	return s
}

func (s *ModifyDataLimitRequest) SetId(v int64) *ModifyDataLimitRequest {
	s.Id = &v
	return s
}

func (s *ModifyDataLimitRequest) SetLang(v string) *ModifyDataLimitRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDataLimitRequest) SetLogStoreDay(v int32) *ModifyDataLimitRequest {
	s.LogStoreDay = &v
	return s
}

func (s *ModifyDataLimitRequest) SetModifyPassword(v bool) *ModifyDataLimitRequest {
	s.ModifyPassword = &v
	return s
}

func (s *ModifyDataLimitRequest) SetPassword(v string) *ModifyDataLimitRequest {
	s.Password = &v
	return s
}

func (s *ModifyDataLimitRequest) SetPort(v int32) *ModifyDataLimitRequest {
	s.Port = &v
	return s
}

func (s *ModifyDataLimitRequest) SetResourceType(v int32) *ModifyDataLimitRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyDataLimitRequest) SetSamplingSize(v int32) *ModifyDataLimitRequest {
	s.SamplingSize = &v
	return s
}

func (s *ModifyDataLimitRequest) SetSecurityGroupIdList(v []*string) *ModifyDataLimitRequest {
	s.SecurityGroupIdList = v
	return s
}

func (s *ModifyDataLimitRequest) SetServiceRegionId(v string) *ModifyDataLimitRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *ModifyDataLimitRequest) SetUserName(v string) *ModifyDataLimitRequest {
	s.UserName = &v
	return s
}

func (s *ModifyDataLimitRequest) SetVSwitchIdList(v []*string) *ModifyDataLimitRequest {
	s.VSwitchIdList = v
	return s
}

func (s *ModifyDataLimitRequest) SetVpcId(v string) *ModifyDataLimitRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyDataLimitRequest) Validate() error {
	return dara.Validate(s)
}

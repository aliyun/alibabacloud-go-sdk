// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeAppInstanceAttributeResponseBody
	GetAppName() *string
	SetAppType(v string) *DescribeAppInstanceAttributeResponseBody
	GetAppType() *string
	SetBranchName(v string) *DescribeAppInstanceAttributeResponseBody
	GetBranchName() *string
	SetBranchingEnabled(v string) *DescribeAppInstanceAttributeResponseBody
	GetBranchingEnabled() *string
	SetComponents(v []*DescribeAppInstanceAttributeResponseBodyComponents) *DescribeAppInstanceAttributeResponseBody
	GetComponents() []*DescribeAppInstanceAttributeResponseBodyComponents
	SetDBInstanceName(v string) *DescribeAppInstanceAttributeResponseBody
	GetDBInstanceName() *string
	SetEipId(v string) *DescribeAppInstanceAttributeResponseBody
	GetEipId() *string
	SetEipStatus(v string) *DescribeAppInstanceAttributeResponseBody
	GetEipStatus() *string
	SetInstanceClass(v string) *DescribeAppInstanceAttributeResponseBody
	GetInstanceClass() *string
	SetInstanceLatestVersion(v string) *DescribeAppInstanceAttributeResponseBody
	GetInstanceLatestVersion() *string
	SetInstanceMinorVersion(v string) *DescribeAppInstanceAttributeResponseBody
	GetInstanceMinorVersion() *string
	SetInstanceName(v string) *DescribeAppInstanceAttributeResponseBody
	GetInstanceName() *string
	SetLoginToken(v string) *DescribeAppInstanceAttributeResponseBody
	GetLoginToken() *string
	SetMinorVersionDetail(v string) *DescribeAppInstanceAttributeResponseBody
	GetMinorVersionDetail() *string
	SetNatCreatedBy(v string) *DescribeAppInstanceAttributeResponseBody
	GetNatCreatedBy() *string
	SetNatGatewayId(v string) *DescribeAppInstanceAttributeResponseBody
	GetNatGatewayId() *string
	SetNatStatus(v string) *DescribeAppInstanceAttributeResponseBody
	GetNatStatus() *string
	SetPublicConnectionString(v string) *DescribeAppInstanceAttributeResponseBody
	GetPublicConnectionString() *string
	SetRegionId(v string) *DescribeAppInstanceAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeAppInstanceAttributeResponseBody
	GetRequestId() *string
	SetRetentionHours(v string) *DescribeAppInstanceAttributeResponseBody
	GetRetentionHours() *string
	SetSqlExtendMoInstanceId(v string) *DescribeAppInstanceAttributeResponseBody
	GetSqlExtendMoInstanceId() *string
	SetStatus(v string) *DescribeAppInstanceAttributeResponseBody
	GetStatus() *string
	SetUploadKey(v string) *DescribeAppInstanceAttributeResponseBody
	GetUploadKey() *string
	SetUploadKeyList(v []*DescribeAppInstanceAttributeResponseBodyUploadKeyList) *DescribeAppInstanceAttributeResponseBody
	GetUploadKeyList() []*DescribeAppInstanceAttributeResponseBodyUploadKeyList
	SetVSwitchId(v string) *DescribeAppInstanceAttributeResponseBody
	GetVSwitchId() *string
	SetVpcConnectionString(v string) *DescribeAppInstanceAttributeResponseBody
	GetVpcConnectionString() *string
	SetZoneId(v string) *DescribeAppInstanceAttributeResponseBody
	GetZoneId() *string
}

type DescribeAppInstanceAttributeResponseBody struct {
	// The name of the AI application.
	//
	// example:
	//
	// test-supabase
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type. Currently, only **supabase*	- is supported, which indicates [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html).
	//
	// example:
	//
	// supabase
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	BranchingEnabled *string `json:"BranchingEnabled,omitempty" xml:"BranchingEnabled,omitempty"`
	// The list of components.
	Components []*DescribeAppInstanceAttributeResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The ID of the RDS PostgreSQL database instance that the AI application is connected to.
	//
	// example:
	//
	// pgm-2ze49qv594vi****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The instance ID of the EIP.
	//
	// example:
	//
	// eip-wz9sfo01afag4hxc0utq0
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	// The activation status of the EIP.
	//
	// example:
	//
	// on
	EipStatus *string `json:"EipStatus,omitempty" xml:"EipStatus,omitempty"`
	// The instance class of the AI application.
	//
	// example:
	//
	// rdsai.supabase.basic
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The latest minor version of the RDS AI application instance.
	//
	// example:
	//
	// 20260903
	InstanceLatestVersion *string `json:"InstanceLatestVersion,omitempty" xml:"InstanceLatestVersion,omitempty"`
	// The minor version of the RDS AI application instance.
	//
	// example:
	//
	// 20241231
	InstanceMinorVersion *string `json:"InstanceMinorVersion,omitempty" xml:"InstanceMinorVersion,omitempty"`
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The minor version details of each component of the RDS AI application instance.
	//
	// example:
	//
	// {\\"studio\\":\\"20260602r1\\",\\"storage\\":\\"v1.48.9\\",\\"auth\\":\\"v2.188.1\\",\\"kong\\":\\"3.9.0\\",\\"rest\\":\\"v12.2.12\\",\\"meta\\":\\"v0.89.3\\",\\"realtime-dev\\":\\"v2.34.47\\",\\"postgres\\":\\"rds_postgres_1700_20260830\\"}
	MinorVersionDetail *string `json:"MinorVersionDetail,omitempty" xml:"MinorVersionDetail,omitempty"`
	// The creator of the NAT gateway.
	//
	// example:
	//
	// user
	NatCreatedBy *string `json:"NatCreatedBy,omitempty" xml:"NatCreatedBy,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-bp1l82hw87m2y77ci1hie
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The activation status of the NAT gateway.
	//
	// example:
	//
	// off
	NatStatus *string `json:"NatStatus,omitempty" xml:"NatStatus,omitempty"`
	// The public connection string of the AI application.
	//
	// example:
	//
	// 8.152. XXX.XXX:8000
	PublicConnectionString *string `json:"PublicConnectionString,omitempty" xml:"PublicConnectionString,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	RetentionHours *string `json:"RetentionHours,omitempty" xml:"RetentionHours,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	SqlExtendMoInstanceId *string `json:"SqlExtendMoInstanceId,omitempty" xml:"SqlExtendMoInstanceId,omitempty"`
	// The instance status. For more information, see [Instance status table](https://help.aliyun.com/document_detail/2623972.html).
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	UploadKey *string `json:"UploadKey,omitempty" xml:"UploadKey,omitempty"`
	// Reserved parameter.
	UploadKeyList []*DescribeAppInstanceAttributeResponseBodyUploadKeyList `json:"UploadKeyList,omitempty" xml:"UploadKeyList,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zeaepb8k4ku05ov2****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The internal connection string of the AI application.
	//
	// example:
	//
	// 172.16.XXX.XXX:8000
	VpcConnectionString *string `json:"VpcConnectionString,omitempty" xml:"VpcConnectionString,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAppInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppInstanceAttributeResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppInstanceAttributeResponseBody) GetAppType() *string {
	return s.AppType
}

func (s *DescribeAppInstanceAttributeResponseBody) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeAppInstanceAttributeResponseBody) GetBranchingEnabled() *string {
	return s.BranchingEnabled
}

func (s *DescribeAppInstanceAttributeResponseBody) GetComponents() []*DescribeAppInstanceAttributeResponseBodyComponents {
	return s.Components
}

func (s *DescribeAppInstanceAttributeResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAppInstanceAttributeResponseBody) GetEipId() *string {
	return s.EipId
}

func (s *DescribeAppInstanceAttributeResponseBody) GetEipStatus() *string {
	return s.EipStatus
}

func (s *DescribeAppInstanceAttributeResponseBody) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeAppInstanceAttributeResponseBody) GetInstanceLatestVersion() *string {
	return s.InstanceLatestVersion
}

func (s *DescribeAppInstanceAttributeResponseBody) GetInstanceMinorVersion() *string {
	return s.InstanceMinorVersion
}

func (s *DescribeAppInstanceAttributeResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAppInstanceAttributeResponseBody) GetLoginToken() *string {
	return s.LoginToken
}

func (s *DescribeAppInstanceAttributeResponseBody) GetMinorVersionDetail() *string {
	return s.MinorVersionDetail
}

func (s *DescribeAppInstanceAttributeResponseBody) GetNatCreatedBy() *string {
	return s.NatCreatedBy
}

func (s *DescribeAppInstanceAttributeResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeAppInstanceAttributeResponseBody) GetNatStatus() *string {
	return s.NatStatus
}

func (s *DescribeAppInstanceAttributeResponseBody) GetPublicConnectionString() *string {
	return s.PublicConnectionString
}

func (s *DescribeAppInstanceAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAppInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppInstanceAttributeResponseBody) GetRetentionHours() *string {
	return s.RetentionHours
}

func (s *DescribeAppInstanceAttributeResponseBody) GetSqlExtendMoInstanceId() *string {
	return s.SqlExtendMoInstanceId
}

func (s *DescribeAppInstanceAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppInstanceAttributeResponseBody) GetUploadKey() *string {
	return s.UploadKey
}

func (s *DescribeAppInstanceAttributeResponseBody) GetUploadKeyList() []*DescribeAppInstanceAttributeResponseBodyUploadKeyList {
	return s.UploadKeyList
}

func (s *DescribeAppInstanceAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAppInstanceAttributeResponseBody) GetVpcConnectionString() *string {
	return s.VpcConnectionString
}

func (s *DescribeAppInstanceAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAppInstanceAttributeResponseBody) SetAppName(v string) *DescribeAppInstanceAttributeResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetAppType(v string) *DescribeAppInstanceAttributeResponseBody {
	s.AppType = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetBranchName(v string) *DescribeAppInstanceAttributeResponseBody {
	s.BranchName = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetBranchingEnabled(v string) *DescribeAppInstanceAttributeResponseBody {
	s.BranchingEnabled = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetComponents(v []*DescribeAppInstanceAttributeResponseBodyComponents) *DescribeAppInstanceAttributeResponseBody {
	s.Components = v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetDBInstanceName(v string) *DescribeAppInstanceAttributeResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetEipId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.EipId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetEipStatus(v string) *DescribeAppInstanceAttributeResponseBody {
	s.EipStatus = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetInstanceClass(v string) *DescribeAppInstanceAttributeResponseBody {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetInstanceLatestVersion(v string) *DescribeAppInstanceAttributeResponseBody {
	s.InstanceLatestVersion = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetInstanceMinorVersion(v string) *DescribeAppInstanceAttributeResponseBody {
	s.InstanceMinorVersion = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetInstanceName(v string) *DescribeAppInstanceAttributeResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetLoginToken(v string) *DescribeAppInstanceAttributeResponseBody {
	s.LoginToken = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetMinorVersionDetail(v string) *DescribeAppInstanceAttributeResponseBody {
	s.MinorVersionDetail = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetNatCreatedBy(v string) *DescribeAppInstanceAttributeResponseBody {
	s.NatCreatedBy = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetNatGatewayId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetNatStatus(v string) *DescribeAppInstanceAttributeResponseBody {
	s.NatStatus = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetPublicConnectionString(v string) *DescribeAppInstanceAttributeResponseBody {
	s.PublicConnectionString = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetRegionId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetRequestId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetRetentionHours(v string) *DescribeAppInstanceAttributeResponseBody {
	s.RetentionHours = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetSqlExtendMoInstanceId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.SqlExtendMoInstanceId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetStatus(v string) *DescribeAppInstanceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetUploadKey(v string) *DescribeAppInstanceAttributeResponseBody {
	s.UploadKey = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetUploadKeyList(v []*DescribeAppInstanceAttributeResponseBodyUploadKeyList) *DescribeAppInstanceAttributeResponseBody {
	s.UploadKeyList = v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetVSwitchId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetVpcConnectionString(v string) *DescribeAppInstanceAttributeResponseBody {
	s.VpcConnectionString = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) SetZoneId(v string) *DescribeAppInstanceAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBody) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UploadKeyList != nil {
		for _, item := range s.UploadKeyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppInstanceAttributeResponseBodyComponents struct {
	// The component status.
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The component type.
	//
	// example:
	//
	// supabase
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAppInstanceAttributeResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstanceAttributeResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *DescribeAppInstanceAttributeResponseBodyComponents) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppInstanceAttributeResponseBodyComponents) GetType() *string {
	return s.Type
}

func (s *DescribeAppInstanceAttributeResponseBodyComponents) SetStatus(v string) *DescribeAppInstanceAttributeResponseBodyComponents {
	s.Status = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBodyComponents) SetType(v string) *DescribeAppInstanceAttributeResponseBodyComponents {
	s.Type = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBodyComponents) Validate() error {
	return dara.Validate(s)
}

type DescribeAppInstanceAttributeResponseBodyUploadKeyList struct {
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	IsSystemKey *bool `json:"IsSystemKey,omitempty" xml:"IsSystemKey,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	SlsStorageBytes *int64 `json:"SlsStorageBytes,omitempty" xml:"SlsStorageBytes,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	UploadKey *string `json:"UploadKey,omitempty" xml:"UploadKey,omitempty"`
}

func (s DescribeAppInstanceAttributeResponseBodyUploadKeyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstanceAttributeResponseBodyUploadKeyList) GoString() string {
	return s.String()
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) GetIsSystemKey() *bool {
	return s.IsSystemKey
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) GetRemark() *string {
	return s.Remark
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) GetSlsStorageBytes() *int64 {
	return s.SlsStorageBytes
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) GetUploadKey() *string {
	return s.UploadKey
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) SetIsSystemKey(v bool) *DescribeAppInstanceAttributeResponseBodyUploadKeyList {
	s.IsSystemKey = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) SetRemark(v string) *DescribeAppInstanceAttributeResponseBodyUploadKeyList {
	s.Remark = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) SetSlsStorageBytes(v int64) *DescribeAppInstanceAttributeResponseBodyUploadKeyList {
	s.SlsStorageBytes = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) SetStatus(v string) *DescribeAppInstanceAttributeResponseBodyUploadKeyList {
	s.Status = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) SetUploadKey(v string) *DescribeAppInstanceAttributeResponseBodyUploadKeyList {
	s.UploadKey = &v
	return s
}

func (s *DescribeAppInstanceAttributeResponseBodyUploadKeyList) Validate() error {
	return dara.Validate(s)
}

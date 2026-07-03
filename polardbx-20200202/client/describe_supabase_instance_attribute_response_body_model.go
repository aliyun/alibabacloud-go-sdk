// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) *DescribeSupabaseInstanceAttributeResponseBody
	GetAccessDeniedDetail() *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail
	SetData(v *DescribeSupabaseInstanceAttributeResponseBodyData) *DescribeSupabaseInstanceAttributeResponseBody
	GetData() *DescribeSupabaseInstanceAttributeResponseBodyData
	SetRequestId(v string) *DescribeSupabaseInstanceAttributeResponseBody
	GetRequestId() *string
}

type DescribeSupabaseInstanceAttributeResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The instance details.
	Data *DescribeSupabaseInstanceAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSupabaseInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstanceAttributeResponseBody) GetAccessDeniedDetail() *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeSupabaseInstanceAttributeResponseBody) GetData() *DescribeSupabaseInstanceAttributeResponseBodyData {
	return s.Data
}

func (s *DescribeSupabaseInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupabaseInstanceAttributeResponseBody) SetAccessDeniedDetail(v *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) *DescribeSupabaseInstanceAttributeResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBody) SetData(v *DescribeSupabaseInstanceAttributeResponseBodyData) *DescribeSupabaseInstanceAttributeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBody) SetRequestId(v string) *DescribeSupabaseInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the authentication principal.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of the permission denial.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeSupabaseInstanceAttributeResponseBodyData struct {
	// The list of endpoints.
	ConnAddrs []*DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2025-07-13 10:01:50 +0800
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxc-supabase-001
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Indicates whether the instance has expired.
	//
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The lock mode.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The minor engine version.
	//
	// example:
	//
	// polardb-2.4.0_5.4.19-20250116_xcluster5.4.20-20241213
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The node specifications.
	//
	// example:
	//
	// polarx.supabase.x2.small
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The list of nodes.
	Nodes []*DescribeSupabaseInstanceAttributeResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The name of the associated PolarDB-X instance.
	//
	// example:
	//
	// pxc-*********
	PolarDBXDBInstanceName *string `json:"PolarDBXDBInstanceName,omitempty" xml:"PolarDBXDBInstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Supabase class code.
	//
	// example:
	//
	// polarx.supabase.x4.xlarge
	SupabaseClassCode *string `json:"SupabaseClassCode,omitempty" xml:"SupabaseClassCode,omitempty"`
	// The multi-tenant mode.
	//
	// example:
	//
	// false
	TenantMode *bool `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	// The topology type.
	//
	// example:
	//
	// 1azone
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-********
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-zhangjiakou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSupabaseInstanceAttributeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstanceAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetConnAddrs() []*DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs {
	return s.ConnAddrs
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetNodes() []*DescribeSupabaseInstanceAttributeResponseBodyDataNodes {
	return s.Nodes
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetPolarDBXDBInstanceName() *string {
	return s.PolarDBXDBInstanceName
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetSupabaseClassCode() *string {
	return s.SupabaseClassCode
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetTenantMode() *bool {
	return s.TenantMode
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetTopologyType() *string {
	return s.TopologyType
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetConnAddrs(v []*DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.ConnAddrs = v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetCreateTime(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetDBInstanceName(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetEngineVersion(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.EngineVersion = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetExpired(v bool) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.Expired = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetLockMode(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.LockMode = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetMinorVersion(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.MinorVersion = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetNodeClass(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.NodeClass = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetNodeCount(v int32) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.NodeCount = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetNodes(v []*DescribeSupabaseInstanceAttributeResponseBodyDataNodes) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.Nodes = v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetPolarDBXDBInstanceName(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.PolarDBXDBInstanceName = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetRegionId(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetStatus(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetSupabaseClassCode(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.SupabaseClassCode = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetTenantMode(v bool) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.TenantMode = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetTopologyType(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.TopologyType = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetVSwitchId(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetVpcId(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) SetZoneId(v string) *DescribeSupabaseInstanceAttributeResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyData) Validate() error {
	if s.ConnAddrs != nil {
		for _, item := range s.ConnAddrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs struct {
	// The endpoint.
	//
	// example:
	//
	// pxsp-********.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The network type.
	//
	// example:
	//
	// 0
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-********
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) GetNetType() *string {
	return s.NetType
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) GetPort() *int32 {
	return s.Port
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) SetConnectionString(v string) *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) SetNetType(v string) *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs {
	s.NetType = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) SetPort(v int32) *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) SetVpcId(v string) *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs {
	s.VpcId = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataConnAddrs) Validate() error {
	return dara.Validate(s)
}

type DescribeSupabaseInstanceAttributeResponseBodyDataNodes struct {
	// The class code.
	//
	// example:
	//
	// polarx.supabase.x2.small
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 2005777
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The node name.
	//
	// example:
	//
	// pxsp-********
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSupabaseInstanceAttributeResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstanceAttributeResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataNodes) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataNodes) GetId() *string {
	return s.Id
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataNodes) GetName() *string {
	return s.Name
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataNodes) SetClassCode(v string) *DescribeSupabaseInstanceAttributeResponseBodyDataNodes {
	s.ClassCode = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataNodes) SetId(v string) *DescribeSupabaseInstanceAttributeResponseBodyDataNodes {
	s.Id = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataNodes) SetName(v string) *DescribeSupabaseInstanceAttributeResponseBodyDataNodes {
	s.Name = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataNodes) SetZoneId(v string) *DescribeSupabaseInstanceAttributeResponseBodyDataNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}

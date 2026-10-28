// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribePxfsInfoResponseBodyAccessDeniedDetail) *DescribePxfsInfoResponseBody
	GetAccessDeniedDetail() *DescribePxfsInfoResponseBodyAccessDeniedDetail
	SetData(v *DescribePxfsInfoResponseBodyData) *DescribePxfsInfoResponseBody
	GetData() *DescribePxfsInfoResponseBodyData
	SetRequestId(v string) *DescribePxfsInfoResponseBody
	GetRequestId() *string
}

type DescribePxfsInfoResponseBody struct {
	// The details about the access denial.
	AccessDeniedDetail *DescribePxfsInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The response data.
	Data *DescribePxfsInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePxfsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePxfsInfoResponseBody) GetAccessDeniedDetail() *DescribePxfsInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribePxfsInfoResponseBody) GetData() *DescribePxfsInfoResponseBodyData {
	return s.Data
}

func (s *DescribePxfsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePxfsInfoResponseBody) SetAccessDeniedDetail(v *DescribePxfsInfoResponseBodyAccessDeniedDetail) *DescribePxfsInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribePxfsInfoResponseBody) SetData(v *DescribePxfsInfoResponseBodyData) *DescribePxfsInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribePxfsInfoResponseBody) SetRequestId(v string) *DescribePxfsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePxfsInfoResponseBody) Validate() error {
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

type DescribePxfsInfoResponseBodyAccessDeniedDetail struct {
	// The authorization action. Same as described above.
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
	// The authorization action. Same as described above.
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
	// The type of missing permission.
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

func (s DescribePxfsInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribePxfsInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribePxfsInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribePxfsInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribePxfsInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribePxfsInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribePxfsInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribePxfsInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePxfsInfoResponseBodyData struct {
	// The instance information.
	Instance *DescribePxfsInfoResponseBodyDataInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
}

func (s DescribePxfsInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePxfsInfoResponseBodyData) GetInstance() *DescribePxfsInfoResponseBodyDataInstance {
	return s.Instance
}

func (s *DescribePxfsInfoResponseBodyData) SetInstance(v *DescribePxfsInfoResponseBodyDataInstance) *DescribePxfsInfoResponseBodyData {
	s.Instance = v
	return s
}

func (s *DescribePxfsInfoResponseBodyData) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePxfsInfoResponseBodyDataInstance struct {
	// The node specification.
	//
	// example:
	//
	// polarx.pxfs.x4.xlarge
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The list of connection addresses.
	ConnAddrs []*DescribePxfsInfoResponseBodyDataInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2026-09-03T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The PolarDB-X instance ID.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The associated instance ID.
	//
	// example:
	//
	// pxc-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 3
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The Agent file service instance ID.
	//
	// example:
	//
	// pxfs-xxxxxxxx
	PxfsInstanceId *string `json:"PxfsInstanceId,omitempty" xml:"PxfsInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage size.
	//
	// example:
	//
	// 500
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The storage type.
	//
	// example:
	//
	// cloud_essd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-xxxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The current version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribePxfsInfoResponseBodyDataInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsInfoResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetConnAddrs() []*DescribePxfsInfoResponseBodyDataInstanceConnAddrs {
	return s.ConnAddrs
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetPxfsInstanceId() *string {
	return s.PxfsInstanceId
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetVersion() *string {
	return s.Version
}

func (s *DescribePxfsInfoResponseBodyDataInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetClassCode(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.ClassCode = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetConnAddrs(v []*DescribePxfsInfoResponseBodyDataInstanceConnAddrs) *DescribePxfsInfoResponseBodyDataInstance {
	s.ConnAddrs = v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetCreateTime(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetDBInstanceName(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.DBInstanceName = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetInstanceId(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetNodeCount(v int32) *DescribePxfsInfoResponseBodyDataInstance {
	s.NodeCount = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetPxfsInstanceId(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.PxfsInstanceId = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetRegionId(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.RegionId = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetStatus(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.Status = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetStorageSize(v int64) *DescribePxfsInfoResponseBodyDataInstance {
	s.StorageSize = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetStorageType(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.StorageType = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetVPCId(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.VPCId = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetVSwitchId(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetVersion(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.Version = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) SetZoneId(v string) *DescribePxfsInfoResponseBodyDataInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstance) Validate() error {
	if s.ConnAddrs != nil {
		for _, item := range s.ConnAddrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePxfsInfoResponseBodyDataInstanceConnAddrs struct {
	// The connection string.
	//
	// example:
	//
	// agent-files.example.com:12049
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// agent-files.example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The port number.
	//
	// example:
	//
	// 12049
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The access protocol.
	//
	// example:
	//
	// NFS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Indicates whether TLS is enabled.
	//
	// example:
	//
	// true
	TlsEnabled *bool `json:"TlsEnabled,omitempty" xml:"TlsEnabled,omitempty"`
	// The transport protocol.
	//
	// example:
	//
	// TLS
	Transport *string `json:"Transport,omitempty" xml:"Transport,omitempty"`
	// The connection type.
	//
	// example:
	//
	// Intranet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-xxxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribePxfsInfoResponseBodyDataInstanceConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsInfoResponseBodyDataInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) GetPort() *int32 {
	return s.Port
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) GetTlsEnabled() *bool {
	return s.TlsEnabled
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) GetTransport() *string {
	return s.Transport
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) GetType() *string {
	return s.Type
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) SetConnectionString(v string) *DescribePxfsInfoResponseBodyDataInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) SetEndpoint(v string) *DescribePxfsInfoResponseBodyDataInstanceConnAddrs {
	s.Endpoint = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) SetPort(v int32) *DescribePxfsInfoResponseBodyDataInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) SetProtocol(v string) *DescribePxfsInfoResponseBodyDataInstanceConnAddrs {
	s.Protocol = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) SetTlsEnabled(v bool) *DescribePxfsInfoResponseBodyDataInstanceConnAddrs {
	s.TlsEnabled = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) SetTransport(v string) *DescribePxfsInfoResponseBodyDataInstanceConnAddrs {
	s.Transport = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) SetType(v string) *DescribePxfsInfoResponseBodyDataInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) SetVPCId(v string) *DescribePxfsInfoResponseBodyDataInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) SetVSwitchId(v string) *DescribePxfsInfoResponseBodyDataInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribePxfsInfoResponseBodyDataInstanceConnAddrs) Validate() error {
	return dara.Validate(s)
}

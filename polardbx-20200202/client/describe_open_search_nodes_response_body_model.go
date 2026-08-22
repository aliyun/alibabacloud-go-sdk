// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) *DescribeOpenSearchNodesResponseBody
	GetAccessDeniedDetail() *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail
	SetData(v *DescribeOpenSearchNodesResponseBodyData) *DescribeOpenSearchNodesResponseBody
	GetData() *DescribeOpenSearchNodesResponseBodyData
	SetRequestId(v string) *DescribeOpenSearchNodesResponseBody
	GetRequestId() *string
}

type DescribeOpenSearchNodesResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *DescribeOpenSearchNodesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// AE4F6C34-065F-45AA-B5DC-4B8D816F6305
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenSearchNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchNodesResponseBody) GetAccessDeniedDetail() *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeOpenSearchNodesResponseBody) GetData() *DescribeOpenSearchNodesResponseBodyData {
	return s.Data
}

func (s *DescribeOpenSearchNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenSearchNodesResponseBody) SetAccessDeniedDetail(v *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) *DescribeOpenSearchNodesResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeOpenSearchNodesResponseBody) SetData(v *DescribeOpenSearchNodesResponseBodyData) *DescribeOpenSearchNodesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenSearchNodesResponseBody) SetRequestId(v string) *DescribeOpenSearchNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBody) Validate() error {
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

type DescribeOpenSearchNodesResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The identity used for authentication in the request.
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

func (s DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchNodesResponseBodyData struct {
	// The query result object.
	Result []*DescribeOpenSearchNodesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeOpenSearchNodesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchNodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchNodesResponseBodyData) GetResult() []*DescribeOpenSearchNodesResponseBodyDataResult {
	return s.Result
}

func (s *DescribeOpenSearchNodesResponseBodyData) SetResult(v []*DescribeOpenSearchNodesResponseBodyDataResult) *DescribeOpenSearchNodesResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOpenSearchNodesResponseBodyDataResult struct {
	// The number of CPU cores of the node.
	//
	// example:
	//
	// 1
	CpuCores *int32 `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// The CPU usage (%).
	//
	// example:
	//
	// 35.6
	CpuPercent *string `json:"CpuPercent,omitempty" xml:"CpuPercent,omitempty"`
	// The total disk capacity of the node, in GB.
	//
	// example:
	//
	// 500
	DiskSizeGB *int32 `json:"DiskSizeGB,omitempty" xml:"DiskSizeGB,omitempty"`
	// The disk space usage of the node.
	//
	// example:
	//
	// 42.5
	DiskUsedPercent *string `json:"DiskUsedPercent,omitempty" xml:"DiskUsedPercent,omitempty"`
	// The total number of unresolved baseline check items.
	//
	// example:
	//
	// GREEN
	Health *string `json:"Health,omitempty" xml:"Health,omitempty"`
	// The JVM heap memory usage of the node.
	//
	// example:
	//
	// 38.2
	HeapPercent *string `json:"HeapPercent,omitempty" xml:"HeapPercent,omitempty"`
	// The IP address and port of the session host that initiated the session.
	//
	// example:
	//
	// 100.115.107.0/24
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The name of the host on which the node instance runs. You can log on to the host and run the `hostname` command to view the hostname.
	//
	// example:
	//
	// hb2h-ali-oceanbase-public-online-013
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The average system load of the node over the last 1 minute.
	//
	// example:
	//
	// 1.25
	LoadOneM *string `json:"LoadOneM,omitempty" xml:"LoadOneM,omitempty"`
	// The amount of memory used.
	//
	// example:
	//
	// 16
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
	// The node type to query. Valid values:
	//
	// - all: queries both dn and gms nodes.
	//
	// - gms: queries only gms nodes.
	//
	// - dn: queries only dn nodes.
	//
	// example:
	//
	// dn
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeOpenSearchNodesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchNodesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetCpuCores() *int32 {
	return s.CpuCores
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetCpuPercent() *string {
	return s.CpuPercent
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetDiskSizeGB() *int32 {
	return s.DiskSizeGB
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetDiskUsedPercent() *string {
	return s.DiskUsedPercent
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetHealth() *string {
	return s.Health
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetHeapPercent() *string {
	return s.HeapPercent
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetHost() *string {
	return s.Host
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetHostName() *string {
	return s.HostName
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetLoadOneM() *string {
	return s.LoadOneM
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetPort() *int32 {
	return s.Port
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetCpuCores(v int32) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.CpuCores = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetCpuPercent(v string) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.CpuPercent = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetDiskSizeGB(v int32) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.DiskSizeGB = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetDiskUsedPercent(v string) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.DiskUsedPercent = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetHealth(v string) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.Health = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetHeapPercent(v string) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.HeapPercent = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetHost(v string) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.Host = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetHostName(v string) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.HostName = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetLoadOneM(v string) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.LoadOneM = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetMemoryGB(v int32) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.MemoryGB = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetNodeType(v string) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.NodeType = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetPort(v int32) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.Port = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) SetZoneId(v string) *DescribeOpenSearchNodesResponseBodyDataResult {
	s.ZoneId = &v
	return s
}

func (s *DescribeOpenSearchNodesResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

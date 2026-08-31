// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) *DescribeOpenSearchTopologyResponseBody
	GetAccessDeniedDetail() *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail
	SetData(v *DescribeOpenSearchTopologyResponseBodyData) *DescribeOpenSearchTopologyResponseBody
	GetData() *DescribeOpenSearchTopologyResponseBodyData
	SetRequestId(v string) *DescribeOpenSearchTopologyResponseBody
	GetRequestId() *string
}

type DescribeOpenSearchTopologyResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *DescribeOpenSearchTopologyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenSearchTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchTopologyResponseBody) GetAccessDeniedDetail() *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeOpenSearchTopologyResponseBody) GetData() *DescribeOpenSearchTopologyResponseBodyData {
	return s.Data
}

func (s *DescribeOpenSearchTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenSearchTopologyResponseBody) SetAccessDeniedDetail(v *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) *DescribeOpenSearchTopologyResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBody) SetData(v *DescribeOpenSearchTopologyResponseBodyData) *DescribeOpenSearchTopologyResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBody) SetRequestId(v string) *DescribeOpenSearchTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBody) Validate() error {
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

type DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail struct {
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
	// The authentication principal type.
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
	// NoPermissionType
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

func (s DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchTopologyResponseBodyData struct {
	// The node IDs.
	Nodes []*DescribeOpenSearchTopologyResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The configuration of the automatic storage scaling feature for the instance.
	Storage *DescribeOpenSearchTopologyResponseBodyDataStorage `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
}

func (s DescribeOpenSearchTopologyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchTopologyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchTopologyResponseBodyData) GetNodes() []*DescribeOpenSearchTopologyResponseBodyDataNodes {
	return s.Nodes
}

func (s *DescribeOpenSearchTopologyResponseBodyData) GetStorage() *DescribeOpenSearchTopologyResponseBodyDataStorage {
	return s.Storage
}

func (s *DescribeOpenSearchTopologyResponseBodyData) SetNodes(v []*DescribeOpenSearchTopologyResponseBodyDataNodes) *DescribeOpenSearchTopologyResponseBodyData {
	s.Nodes = v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyData) SetStorage(v *DescribeOpenSearchTopologyResponseBodyDataStorage) *DescribeOpenSearchTopologyResponseBodyData {
	s.Storage = v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyData) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Storage != nil {
		if err := s.Storage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOpenSearchTopologyResponseBodyDataNodes struct {
	// The zone.
	//
	// example:
	//
	// t1222576965886205
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The total number of CPU cores in the cluster.
	//
	// example:
	//
	// 0.25
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The Origin Domain Name that you configured in the DCDN console, including IPv4 addresses, IPv6 addresses, common domain names, and OSS domain names.
	//
	// example:
	//
	// https://secnet-defense-vastip.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Indicates whether the current node is the primary node.
	IsLeader *bool `json:"IsLeader,omitempty" xml:"IsLeader,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 32
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
	// The node ID.
	//
	// example:
	//
	// pxc-c-jf0pivh2dt
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The role of a node in the three-node RDS cluster. Valid values:
	//
	// - **leader**: primary node
	//
	// - **follower**: secondary node
	//
	// - **logger**: logger node
	//
	// example:
	//
	// polarx_cn
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The node status. Valid values:
	//
	// - **0**: Running
	//
	// - **1**: Creating
	//
	// - **2**: Abnormal
	//
	// - **3**: Expired
	//
	// - **4**: Releasing
	//
	// - **5**: Released
	//
	// - **6**: Locked
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeOpenSearchTopologyResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchTopologyResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) GetHost() *string {
	return s.Host
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) GetIsLeader() *bool {
	return s.IsLeader
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) GetRole() *string {
	return s.Role
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) GetStatus() *string {
	return s.Status
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) SetAvailabilityZone(v string) *DescribeOpenSearchTopologyResponseBodyDataNodes {
	s.AvailabilityZone = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) SetCpu(v int32) *DescribeOpenSearchTopologyResponseBodyDataNodes {
	s.Cpu = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) SetHost(v string) *DescribeOpenSearchTopologyResponseBodyDataNodes {
	s.Host = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) SetIsLeader(v bool) *DescribeOpenSearchTopologyResponseBodyDataNodes {
	s.IsLeader = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) SetMemoryGB(v int32) *DescribeOpenSearchTopologyResponseBodyDataNodes {
	s.MemoryGB = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) SetNodeId(v string) *DescribeOpenSearchTopologyResponseBodyDataNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) SetRole(v string) *DescribeOpenSearchTopologyResponseBodyDataNodes {
	s.Role = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) SetStatus(v string) *DescribeOpenSearchTopologyResponseBodyDataNodes {
	s.Status = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchTopologyResponseBodyDataStorage struct {
	// The number of secondary nodes in the primary zone.
	//
	// > The **ReplicaCount*	- and **SlaveReplicaCount*	- parameters apply only to cloud-native instances. If the instance uses the cluster architecture, these parameters indicate the number of secondary nodes of a **single shard*	- in the primary and secondary zones.
	//
	// example:
	//
	// 1
	ReplicaCount *int32 `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	// The total storage capacity of the node. Unit: GB.
	//
	// example:
	//
	// 500
	StorageTotalGB *int32 `json:"StorageTotalGB,omitempty" xml:"StorageTotalGB,omitempty"`
	// The storage type.
	//
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeOpenSearchTopologyResponseBodyDataStorage) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchTopologyResponseBodyDataStorage) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchTopologyResponseBodyDataStorage) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *DescribeOpenSearchTopologyResponseBodyDataStorage) GetStorageTotalGB() *int32 {
	return s.StorageTotalGB
}

func (s *DescribeOpenSearchTopologyResponseBodyDataStorage) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeOpenSearchTopologyResponseBodyDataStorage) SetReplicaCount(v int32) *DescribeOpenSearchTopologyResponseBodyDataStorage {
	s.ReplicaCount = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataStorage) SetStorageTotalGB(v int32) *DescribeOpenSearchTopologyResponseBodyDataStorage {
	s.StorageTotalGB = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataStorage) SetStorageType(v string) *DescribeOpenSearchTopologyResponseBodyDataStorage {
	s.StorageType = &v
	return s
}

func (s *DescribeOpenSearchTopologyResponseBodyDataStorage) Validate() error {
	return dara.Validate(s)
}

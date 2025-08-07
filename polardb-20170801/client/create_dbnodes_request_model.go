// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDBNodesRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateDBNodesRequest
	GetDBClusterId() *string
	SetDBNode(v []*CreateDBNodesRequestDBNode) *CreateDBNodesRequest
	GetDBNode() []*CreateDBNodesRequestDBNode
	SetDBNodeType(v string) *CreateDBNodesRequest
	GetDBNodeType() *string
	SetEndpointBindList(v string) *CreateDBNodesRequest
	GetEndpointBindList() *string
	SetImciSwitch(v string) *CreateDBNodesRequest
	GetImciSwitch() *string
	SetOwnerAccount(v string) *CreateDBNodesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBNodesRequest
	GetOwnerId() *int64
	SetPlannedEndTime(v string) *CreateDBNodesRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *CreateDBNodesRequest
	GetPlannedStartTime() *string
	SetResourceGroupId(v string) *CreateDBNodesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDBNodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBNodesRequest
	GetResourceOwnerId() *int64
}

type CreateDBNodesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. The token is case-sensitive.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The details of the read-only node.
	//
	// This parameter is required.
	DBNode []*CreateDBNodesRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
	// The node type. Valid values:
	//
	// 	- RO
	//
	// 	- STANDBY
	//
	// 	- DLNode
	//
	// Enumerated values:
	//
	// 	- DLNode: AI node
	//
	// 	- STANDBY: standby node
	//
	// 	- RO: read-only node
	//
	// example:
	//
	// RO
	DBNodeType *string `json:"DBNodeType,omitempty" xml:"DBNodeType,omitempty"`
	// The ID of the cluster endpoint to which the read-only node is added. If you want to add the read-only node to multiple endpoints at the same time, separate the endpoint IDs with commas (,).
	//
	// > - You can call the [DescribeDBClusterEndpoints](https://help.aliyun.com/document_detail/98205.html) operation to query the details of cluster endpoints, including endpoint IDs.
	//
	// >- You can enter the ID of the default cluster endpoint or a custom cluster endpoint.
	//
	// >- If you leave this parameter empty, the read-only node is added to all cluster endpoints for which the **Automatically Associate New Nodes*	- feature is enabled. If you set `AutoAddNewNodes` to `Enable`, the Automatically Associate New Nodes feature is enabled.
	//
	// example:
	//
	// pe-****************,pe-****************
	EndpointBindList *string `json:"EndpointBindList,omitempty" xml:"EndpointBindList,omitempty"`
	// Specifies whether to enable the In-Memory Column Index (IMCI) feature. Default value: OFF. Valid values:
	//
	// 	- **ON**
	//
	// 	- **OFF**
	//
	// > This parameter is invalid for a PolarDB for Oracle or PolarDB for PostgreSQL cluster.
	//
	// example:
	//
	// ON
	ImciSwitch   *string `json:"ImciSwitch,omitempty" xml:"ImciSwitch,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest start time for upgrading the specifications within the scheduled time period. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// >- The value of this parameter must be at least 30 minutes later than the value of PlannedStartTime.
	//
	// >- If you specify `PlannedStartTime` but do not specify PlannedEndTime, the latest start time of the task is set to a value that is calculated by using the following formula: `PlannedEndTime value + 30 minutes`. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and you do not specify PlannedEndTime, the latest start time of the task is set to `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest start time of the scheduled task for adding the read-only node. The scheduled task specifies that the task is run in the required period. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// >- The earliest start time of the scheduled task can be a point in time within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can specify a point in time between `2021-01-14T09:00:00Z` and `2021-01-15T09:00:00Z`.
	//
	// >- If you leave this parameter empty, the task for adding the read-only node is immediately run by default.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDBNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBNodesRequest) GoString() string {
	return s.String()
}

func (s *CreateDBNodesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBNodesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBNodesRequest) GetDBNode() []*CreateDBNodesRequestDBNode {
	return s.DBNode
}

func (s *CreateDBNodesRequest) GetDBNodeType() *string {
	return s.DBNodeType
}

func (s *CreateDBNodesRequest) GetEndpointBindList() *string {
	return s.EndpointBindList
}

func (s *CreateDBNodesRequest) GetImciSwitch() *string {
	return s.ImciSwitch
}

func (s *CreateDBNodesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBNodesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBNodesRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *CreateDBNodesRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *CreateDBNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBNodesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBNodesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBNodesRequest) SetClientToken(v string) *CreateDBNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBNodesRequest) SetDBClusterId(v string) *CreateDBNodesRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBNodesRequest) SetDBNode(v []*CreateDBNodesRequestDBNode) *CreateDBNodesRequest {
	s.DBNode = v
	return s
}

func (s *CreateDBNodesRequest) SetDBNodeType(v string) *CreateDBNodesRequest {
	s.DBNodeType = &v
	return s
}

func (s *CreateDBNodesRequest) SetEndpointBindList(v string) *CreateDBNodesRequest {
	s.EndpointBindList = &v
	return s
}

func (s *CreateDBNodesRequest) SetImciSwitch(v string) *CreateDBNodesRequest {
	s.ImciSwitch = &v
	return s
}

func (s *CreateDBNodesRequest) SetOwnerAccount(v string) *CreateDBNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBNodesRequest) SetOwnerId(v int64) *CreateDBNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBNodesRequest) SetPlannedEndTime(v string) *CreateDBNodesRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *CreateDBNodesRequest) SetPlannedStartTime(v string) *CreateDBNodesRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *CreateDBNodesRequest) SetResourceGroupId(v string) *CreateDBNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBNodesRequest) SetResourceOwnerAccount(v string) *CreateDBNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBNodesRequest) SetResourceOwnerId(v int64) *CreateDBNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBNodesRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDBNodesRequestDBNode struct {
	// The specifications of the read-only node that you want to add, which must be the same as the specifications of the existing nodes. For more information, see the following topics:
	//
	// 	- PolarDB for MySQL: [Specifications of compute nodes](https://help.aliyun.com/document_detail/102542.html)
	//
	// 	- PolarDB for PostgreSQL (Compatible with Oracle): [Specifications of compute nodes](https://help.aliyun.com/document_detail/207921.html)
	//
	// 	- PolarDB for PostgreSQL: [Specifications of compute nodes](https://help.aliyun.com/document_detail/209380.html)
	//
	// >- You need to specify either DBNode.N.ZoneId or DBNode.N.TargetClass. N is an integer that starts from 1. The maximum value of N is equal to 16 minus the number of existing nodes.
	//
	// >- You can add multiple read-only nodes at the same time only to PolarDB for MySQL clusters, which can contain up to of 15 read-only nodes.
	//
	// >- This parameter is required for PolarDB for PostgreSQL (Compatible with Oracle) clusters or PolarDB for PostgreSQL clusters. This parameter is optional for PolarDB for MySQL clusters.
	//
	// example:
	//
	// polar.mysql.x4.medium
	TargetClass *string `json:"TargetClass,omitempty" xml:"TargetClass,omitempty"`
	// The zone ID of the node that you want to add, which must be the same as the zone ID of existing nodes. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the IDs of zones.
	//
	// >- You need to specify either DBNode.N.ZoneId or DBNode.N.TargetClass. N is an integer that starts from 1. The maximum value of N is equal to 16 minus the number of existing nodes.
	//
	// >- You can add multiple read-only nodes at the same time only to PolarDB for MySQL clusters, which can contain up to of 15 read-only nodes.
	//
	// >- This parameter is required for PolarDB for PostgreSQL (Compatible with Oracle) clusters or PolarDB for PostgreSQL clusters. This parameter is optional for PolarDB for MySQL clusters.
	//
	// example:
	//
	// cn-qingdao-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBNodesRequestDBNode) String() string {
	return dara.Prettify(s)
}

func (s CreateDBNodesRequestDBNode) GoString() string {
	return s.String()
}

func (s *CreateDBNodesRequestDBNode) GetTargetClass() *string {
	return s.TargetClass
}

func (s *CreateDBNodesRequestDBNode) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBNodesRequestDBNode) SetTargetClass(v string) *CreateDBNodesRequestDBNode {
	s.TargetClass = &v
	return s
}

func (s *CreateDBNodesRequestDBNode) SetZoneId(v string) *CreateDBNodesRequestDBNode {
	s.ZoneId = &v
	return s
}

func (s *CreateDBNodesRequestDBNode) Validate() error {
	return dara.Validate(s)
}

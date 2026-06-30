// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *CreateDBNodesRequest
	GetAutoUseCoupon() *bool
	SetClientToken(v string) *CreateDBNodesRequest
	GetClientToken() *string
	SetCloudProvider(v string) *CreateDBNodesRequest
	GetCloudProvider() *string
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
	SetPromotionCode(v string) *CreateDBNodesRequest
	GetPromotionCode() *string
	SetResourceGroupId(v string) *CreateDBNodesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDBNodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBNodesRequest
	GetResourceOwnerId() *int64
}

type CreateDBNodesRequest struct {
	// Specifies whether to automatically use coupons. Valid values:
	//
	// 	- true (default): Use coupons.
	//
	// 	- false: Do not use coupons.
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value. Make sure that the value is unique among different requests. The token is case-sensitive and can contain only ASCII characters. The token can be up to 64 characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cloud service provider to which the node belongs.
	//
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The information about the new node.
	//
	// This parameter is required.
	DBNode []*CreateDBNodesRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
	// The node type. Valid values:
	//
	// - RO
	//
	// - STANDBY
	//
	// - DLNode.
	//
	// example:
	//
	// RO
	DBNodeType *string `json:"DBNodeType,omitempty" xml:"DBNodeType,omitempty"`
	// The ID of the cluster endpoint to which you want to add the new node. If you want to add the node to multiple endpoints, separate the endpoint IDs with commas (,).
	//
	// > 	- You can call the [DescribeDBClusterEndpoints](https://help.aliyun.com/document_detail/98205.html) operation to query the details of cluster endpoints, including endpoint IDs.
	//
	// >	- You can specify the IDs of the default cluster endpoint and custom cluster endpoints.
	//
	// >	- If you leave this parameter empty, the new node is added to all cluster endpoints that have the **Automatically Associate New Nodes*	- feature enabled (the value of `AutoAddNewNodes` is `Enable`).
	//
	// example:
	//
	// pe-****************,pe-****************
	EndpointBindList *string `json:"EndpointBindList,omitempty" xml:"EndpointBindList,omitempty"`
	// Specifies whether to enable In-Memory Column Index (IMCI). Valid values:
	//
	// - **ON**: enabled.
	//
	// - **OFF**: disabled. This is the default value.
	//
	// > PolarDB for PostgreSQL (Compatible with Oracle) and PolarDB for PostgreSQL do not support this parameter.
	//
	// example:
	//
	// ON
	ImciSwitch   *string `json:"ImciSwitch,omitempty" xml:"ImciSwitch,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest time to start running the scheduled task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format (UTC).
	//
	// > 	- The latest time must be at least 30 minutes later than the start time.
	//
	// >	- If you specify `PlannedStartTime` but leave this parameter empty, the latest time defaults to `start time + 30 minutes`. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and leave this parameter empty, the task starts no later than `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest time to start running the scheduled task for adding nodes. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format (UTC).
	//
	// > 	- The start time must be within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can specify a start time within the range of `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// >	- If you leave this parameter empty, the task for adding nodes is immediately run.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// The coupon code. If you do not specify this parameter, the default coupon is used.
	//
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The resource group ID.
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

func (s *CreateDBNodesRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateDBNodesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBNodesRequest) GetCloudProvider() *string {
	return s.CloudProvider
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

func (s *CreateDBNodesRequest) GetPromotionCode() *string {
	return s.PromotionCode
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

func (s *CreateDBNodesRequest) SetAutoUseCoupon(v bool) *CreateDBNodesRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateDBNodesRequest) SetClientToken(v string) *CreateDBNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBNodesRequest) SetCloudProvider(v string) *CreateDBNodesRequest {
	s.CloudProvider = &v
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

func (s *CreateDBNodesRequest) SetPromotionCode(v string) *CreateDBNodesRequest {
	s.PromotionCode = &v
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
	if s.DBNode != nil {
		for _, item := range s.DBNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDBNodesRequestDBNode struct {
	// The specification of the new node. The specification must be the same as that of the existing nodes. For more information, see the following topics:
	//
	// 	- PolarDB for MySQL: [Compute node specifications](https://help.aliyun.com/document_detail/102542.html).
	//
	// 	- PolarDB for PostgreSQL (Compatible with Oracle): [Compute node specifications](https://help.aliyun.com/document_detail/207921.html).
	//
	// 	- PolarDB for PostgreSQL: [Compute node specifications](https://help.aliyun.com/document_detail/209380.html).
	//
	// >	- You must specify at least one of DBNode.N.ZoneId and DBNode.N.TargetClass. N is an integer that starts from 1. Maximum value of N = 16 - current number of nodes.
	//
	// >	- Only PolarDB for MySQL clusters support adding multiple read-only nodes at a time. You can add up to 15 read-only nodes.
	//
	// >	- This parameter is required for PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL clusters. This parameter is optional for PolarDB for MySQL clusters.
	//
	// example:
	//
	// polar.mysql.x4.medium
	TargetClass *string `json:"TargetClass,omitempty" xml:"TargetClass,omitempty"`
	// The zone of the new node. The zone must be the same as that of the existing nodes. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query zone IDs.
	//
	// >	- You must specify at least one of DBNode.N.ZoneId and DBNode.N.TargetClass. N is an integer that starts from 1. Maximum value of N = 16 - current number of nodes.
	//
	// >	- Only PolarDB for MySQL clusters support adding multiple read-only nodes at a time. You can add up to 15 read-only nodes.
	//
	// >	- This parameter is required for PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL clusters. This parameter is optional for PolarDB for MySQL clusters.
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

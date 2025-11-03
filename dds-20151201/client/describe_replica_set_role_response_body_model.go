// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicaSetRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeReplicaSetRoleResponseBody
	GetDBInstanceId() *string
	SetReplicaSets(v *DescribeReplicaSetRoleResponseBodyReplicaSets) *DescribeReplicaSetRoleResponseBody
	GetReplicaSets() *DescribeReplicaSetRoleResponseBodyReplicaSets
	SetRequestId(v string) *DescribeReplicaSetRoleResponseBody
	GetRequestId() *string
}

type DescribeReplicaSetRoleResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The details of the roles of the replica set instance.
	ReplicaSets *DescribeReplicaSetRoleResponseBodyReplicaSets `json:"ReplicaSets,omitempty" xml:"ReplicaSets,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DB4A0595-FCA9-437F-B2BB-25DBFC009D3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeReplicaSetRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicaSetRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReplicaSetRoleResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeReplicaSetRoleResponseBody) GetReplicaSets() *DescribeReplicaSetRoleResponseBodyReplicaSets {
	return s.ReplicaSets
}

func (s *DescribeReplicaSetRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReplicaSetRoleResponseBody) SetDBInstanceId(v string) *DescribeReplicaSetRoleResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBody) SetReplicaSets(v *DescribeReplicaSetRoleResponseBodyReplicaSets) *DescribeReplicaSetRoleResponseBody {
	s.ReplicaSets = v
	return s
}

func (s *DescribeReplicaSetRoleResponseBody) SetRequestId(v string) *DescribeReplicaSetRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBody) Validate() error {
	if s.ReplicaSets != nil {
		if err := s.ReplicaSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeReplicaSetRoleResponseBodyReplicaSets struct {
	ReplicaSet []*DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet `json:"ReplicaSet,omitempty" xml:"ReplicaSet,omitempty" type:"Repeated"`
}

func (s DescribeReplicaSetRoleResponseBodyReplicaSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicaSetRoleResponseBodyReplicaSets) GoString() string {
	return s.String()
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSets) GetReplicaSet() []*DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	return s.ReplicaSet
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSets) SetReplicaSet(v []*DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) *DescribeReplicaSetRoleResponseBodyReplicaSets {
	s.ReplicaSet = v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSets) Validate() error {
	if s.ReplicaSet != nil {
		for _, item := range s.ReplicaSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet struct {
	// The endpoint of the node.
	//
	// example:
	//
	// dds-bpxxxxxxxx.mongodb.rds.aliyuncs.com
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	// The port number that is used to connect to the node.
	//
	// example:
	//
	// 3717
	ConnectionPort *string `json:"ConnectionPort,omitempty" xml:"ConnectionPort,omitempty"`
	// The connection type of the node.
	//
	// example:
	//
	// SRV
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The remaining duration of the classic network endpoint. Unit: seconds.
	//
	// example:
	//
	// 1209582
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: the virtual private cloud (VPC)
	//
	// 	- **Classic**: the classic network
	//
	// 	- **Public**: the Internet
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The role of the node in the replica set instance. Valid values:
	//
	// 	- **Primary**
	//
	// 	- **Secondary**
	//
	// example:
	//
	// Primary
	ReplicaSetRole *string `json:"ReplicaSetRole,omitempty" xml:"ReplicaSetRole,omitempty"`
	// The role ID of the node.
	//
	// example:
	//
	// 651xxxxx
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) GoString() string {
	return s.String()
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) GetConnectionPort() *string {
	return s.ConnectionPort
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) GetReplicaSetRole() *string {
	return s.ReplicaSetRole
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) GetRoleId() *string {
	return s.RoleId
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetConnectionDomain(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetConnectionPort(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.ConnectionPort = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetConnectionType(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.ConnectionType = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetExpiredTime(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetNetworkType(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.NetworkType = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetReplicaSetRole(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.ReplicaSetRole = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) SetRoleId(v string) *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet {
	s.RoleId = &v
	return s
}

func (s *DescribeReplicaSetRoleResponseBodyReplicaSetsReplicaSet) Validate() error {
	return dara.Validate(s)
}

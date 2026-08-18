// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateContextDBPublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *AllocateContextDBPublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceName(v string) *AllocateContextDBPublicConnectionRequest
	GetDBInstanceName() *string
	SetNodeType(v string) *AllocateContextDBPublicConnectionRequest
	GetNodeType() *string
	SetPort(v string) *AllocateContextDBPublicConnectionRequest
	GetPort() *string
	SetRegionId(v string) *AllocateContextDBPublicConnectionRequest
	GetRegionId() *string
}

type AllocateContextDBPublicConnectionRequest struct {
	// The prefix of the public network connection string.
	//
	// example:
	//
	// pxc-**************
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The type of the target node. Valid values: service and dashboard.
	//
	// example:
	//
	// service
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocateContextDBPublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateContextDBPublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateContextDBPublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateContextDBPublicConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AllocateContextDBPublicConnectionRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *AllocateContextDBPublicConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateContextDBPublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateContextDBPublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateContextDBPublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateContextDBPublicConnectionRequest) SetDBInstanceName(v string) *AllocateContextDBPublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AllocateContextDBPublicConnectionRequest) SetNodeType(v string) *AllocateContextDBPublicConnectionRequest {
	s.NodeType = &v
	return s
}

func (s *AllocateContextDBPublicConnectionRequest) SetPort(v string) *AllocateContextDBPublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateContextDBPublicConnectionRequest) SetRegionId(v string) *AllocateContextDBPublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateContextDBPublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateContext0PublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *AllocateContext0PublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceName(v string) *AllocateContext0PublicConnectionRequest
	GetDBInstanceName() *string
	SetNodeType(v string) *AllocateContext0PublicConnectionRequest
	GetNodeType() *string
	SetPort(v string) *AllocateContext0PublicConnectionRequest
	GetPort() *string
	SetRegionId(v string) *AllocateContext0PublicConnectionRequest
	GetRegionId() *string
}

type AllocateContext0PublicConnectionRequest struct {
	// The prefix of the public network connection string.
	//
	// example:
	//
	// pxc-**************
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The target node type: service or dashboard.
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
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocateContext0PublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateContext0PublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateContext0PublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateContext0PublicConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AllocateContext0PublicConnectionRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *AllocateContext0PublicConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateContext0PublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateContext0PublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateContext0PublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateContext0PublicConnectionRequest) SetDBInstanceName(v string) *AllocateContext0PublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AllocateContext0PublicConnectionRequest) SetNodeType(v string) *AllocateContext0PublicConnectionRequest {
	s.NodeType = &v
	return s
}

func (s *AllocateContext0PublicConnectionRequest) SetPort(v string) *AllocateContext0PublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateContext0PublicConnectionRequest) SetRegionId(v string) *AllocateContext0PublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateContext0PublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}

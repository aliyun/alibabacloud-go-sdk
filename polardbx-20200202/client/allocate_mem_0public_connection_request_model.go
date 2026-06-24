// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateMem0PublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *AllocateMem0PublicConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceName(v string) *AllocateMem0PublicConnectionRequest
	GetDBInstanceName() *string
	SetPort(v string) *AllocateMem0PublicConnectionRequest
	GetPort() *string
	SetRegionId(v string) *AllocateMem0PublicConnectionRequest
	GetRegionId() *string
}

type AllocateMem0PublicConnectionRequest struct {
	// The prefix of the public network connection string.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-spsil01pww4hfz
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The port number.
	//
	// This parameter is required.
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
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocateMem0PublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateMem0PublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateMem0PublicConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateMem0PublicConnectionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AllocateMem0PublicConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateMem0PublicConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateMem0PublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateMem0PublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateMem0PublicConnectionRequest) SetDBInstanceName(v string) *AllocateMem0PublicConnectionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AllocateMem0PublicConnectionRequest) SetPort(v string) *AllocateMem0PublicConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateMem0PublicConnectionRequest) SetRegionId(v string) *AllocateMem0PublicConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateMem0PublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}

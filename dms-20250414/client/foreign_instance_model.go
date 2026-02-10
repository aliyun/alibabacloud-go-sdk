// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForeignInstance interface {
	dara.Model
	String() string
	GoString() string
	SetDataLinkName(v string) *ForeignInstance
	GetDataLinkName() *string
	SetHost(v string) *ForeignInstance
	GetHost() *string
	SetInstanceSource(v string) *ForeignInstance
	GetInstanceSource() *string
	SetInstanceType(v string) *ForeignInstance
	GetInstanceType() *string
	SetPort(v int32) *ForeignInstance
	GetPort() *int32
	SetProperties(v map[string]*string) *ForeignInstance
	GetProperties() map[string]*string
	SetRegionId(v string) *ForeignInstance
	GetRegionId() *string
	SetSid(v string) *ForeignInstance
	GetSid() *string
}

type ForeignInstance struct {
	// The name of the data link.
	//
	// example:
	//
	// dbl_mysql_2337
	DataLinkName *string `json:"DataLinkName,omitempty" xml:"DataLinkName,omitempty"`
	// The endpoint that is used to connect to the database instance.
	//
	// example:
	//
	// 127.0.0.1
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The source of the database instance. Valid values:
	//
	// 	- **PUBLIC_OWN:*	- a self-managed database instance that is deployed on the Internet.
	//
	// 	- **RDS**: an ApsaraDB RDS instance.
	//
	// 	- **ECS_OWN**: a self-managed database that is hosted on an Elastic Compute Service (ECS) instance.
	//
	// 	- **VPC_IDC**: a self-managed database instance that is deployed in the data center over a virtual private cloud (VPC).
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// mysql
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The port number that is used to connect to the database instance.
	//
	// example:
	//
	// 7890
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The information of the properties.
	Properties map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The system identifier (SID) of the database.
	//
	// >  The SID uniquely identifies an Oracle database. After a database is created, a SID is generated for the database.
	//
	// example:
	//
	// HXE
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
}

func (s ForeignInstance) String() string {
	return dara.Prettify(s)
}

func (s ForeignInstance) GoString() string {
	return s.String()
}

func (s *ForeignInstance) GetDataLinkName() *string {
	return s.DataLinkName
}

func (s *ForeignInstance) GetHost() *string {
	return s.Host
}

func (s *ForeignInstance) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *ForeignInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ForeignInstance) GetPort() *int32 {
	return s.Port
}

func (s *ForeignInstance) GetProperties() map[string]*string {
	return s.Properties
}

func (s *ForeignInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *ForeignInstance) GetSid() *string {
	return s.Sid
}

func (s *ForeignInstance) SetDataLinkName(v string) *ForeignInstance {
	s.DataLinkName = &v
	return s
}

func (s *ForeignInstance) SetHost(v string) *ForeignInstance {
	s.Host = &v
	return s
}

func (s *ForeignInstance) SetInstanceSource(v string) *ForeignInstance {
	s.InstanceSource = &v
	return s
}

func (s *ForeignInstance) SetInstanceType(v string) *ForeignInstance {
	s.InstanceType = &v
	return s
}

func (s *ForeignInstance) SetPort(v int32) *ForeignInstance {
	s.Port = &v
	return s
}

func (s *ForeignInstance) SetProperties(v map[string]*string) *ForeignInstance {
	s.Properties = v
	return s
}

func (s *ForeignInstance) SetRegionId(v string) *ForeignInstance {
	s.RegionId = &v
	return s
}

func (s *ForeignInstance) SetSid(v string) *ForeignInstance {
	s.Sid = &v
	return s
}

func (s *ForeignInstance) Validate() error {
	return dara.Validate(s)
}

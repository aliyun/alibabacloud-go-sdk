// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstancePublicConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *ReleaseInstancePublicConnectionRequest
	GetAddressType() *string
	SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest
	GetCurrentConnectionString() *string
	SetDBInstanceId(v string) *ReleaseInstancePublicConnectionRequest
	GetDBInstanceId() *string
}

type ReleaseInstancePublicConnectionRequest struct {
	// The type of the endpoint. Default value: primary. Valid values:
	//
	// 	- **primary**: primary endpoint.
	//
	// 	- **cluster**: cluster endpoint. This type of endpoints can be created only for instances that have multiple coordinator nodes.
	//
	// example:
	//
	// Intranet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The public endpoint of the instance.
	//
	// You can log on to the AnalyticDB for PostgreSQL console and go to the **Basic Information*	- page of the instance to view the **public endpoint*	- in the **Database Connection*	- section.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****.gpdb.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *ReleaseInstancePublicConnectionRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ReleaseInstancePublicConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ReleaseInstancePublicConnectionRequest) SetAddressType(v string) *ReleaseInstancePublicConnectionRequest {
	s.AddressType = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetDBInstanceId(v string) *ReleaseInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) Validate() error {
	return dara.Validate(s)
}

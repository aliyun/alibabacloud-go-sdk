// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceIPArrayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateDBInstanceIPArrayRequest
	GetDBInstanceId() *string
	SetIPArrayAttribute(v string) *CreateDBInstanceIPArrayRequest
	GetIPArrayAttribute() *string
	SetIPArrayName(v string) *CreateDBInstanceIPArrayRequest
	GetIPArrayName() *string
	SetSecurityIPList(v []*string) *CreateDBInstanceIPArrayRequest
	GetSecurityIPList() []*string
}

type CreateDBInstanceIPArrayRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The attribute of the IP whitelist group.
	//
	// example:
	//
	// taffyFish
	IPArrayAttribute *string `json:"IPArrayAttribute,omitempty" xml:"IPArrayAttribute,omitempty"`
	// The name of the IP whitelist group.
	//
	// This parameter is required.
	//
	// example:
	//
	// testarray
	IPArrayName *string `json:"IPArrayName,omitempty" xml:"IPArrayName,omitempty"`
	// The IP addresses in the IP whitelist group. A maximum of 1,000 IP addresses or CIDR blocks are allowed.
	//
	// This parameter is required.
	SecurityIPList []*string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty" type:"Repeated"`
}

func (s CreateDBInstanceIPArrayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceIPArrayRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceIPArrayRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceIPArrayRequest) GetIPArrayAttribute() *string {
	return s.IPArrayAttribute
}

func (s *CreateDBInstanceIPArrayRequest) GetIPArrayName() *string {
	return s.IPArrayName
}

func (s *CreateDBInstanceIPArrayRequest) GetSecurityIPList() []*string {
	return s.SecurityIPList
}

func (s *CreateDBInstanceIPArrayRequest) SetDBInstanceId(v string) *CreateDBInstanceIPArrayRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceIPArrayRequest) SetIPArrayAttribute(v string) *CreateDBInstanceIPArrayRequest {
	s.IPArrayAttribute = &v
	return s
}

func (s *CreateDBInstanceIPArrayRequest) SetIPArrayName(v string) *CreateDBInstanceIPArrayRequest {
	s.IPArrayName = &v
	return s
}

func (s *CreateDBInstanceIPArrayRequest) SetSecurityIPList(v []*string) *CreateDBInstanceIPArrayRequest {
	s.SecurityIPList = v
	return s
}

func (s *CreateDBInstanceIPArrayRequest) Validate() error {
	return dara.Validate(s)
}

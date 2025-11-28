// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceIPArrayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateDBInstanceIPArrayShrinkRequest
	GetDBInstanceId() *string
	SetIPArrayAttribute(v string) *CreateDBInstanceIPArrayShrinkRequest
	GetIPArrayAttribute() *string
	SetIPArrayName(v string) *CreateDBInstanceIPArrayShrinkRequest
	GetIPArrayName() *string
	SetSecurityIPListShrink(v string) *CreateDBInstanceIPArrayShrinkRequest
	GetSecurityIPListShrink() *string
}

type CreateDBInstanceIPArrayShrinkRequest struct {
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
	SecurityIPListShrink *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s CreateDBInstanceIPArrayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceIPArrayShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceIPArrayShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceIPArrayShrinkRequest) GetIPArrayAttribute() *string {
	return s.IPArrayAttribute
}

func (s *CreateDBInstanceIPArrayShrinkRequest) GetIPArrayName() *string {
	return s.IPArrayName
}

func (s *CreateDBInstanceIPArrayShrinkRequest) GetSecurityIPListShrink() *string {
	return s.SecurityIPListShrink
}

func (s *CreateDBInstanceIPArrayShrinkRequest) SetDBInstanceId(v string) *CreateDBInstanceIPArrayShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceIPArrayShrinkRequest) SetIPArrayAttribute(v string) *CreateDBInstanceIPArrayShrinkRequest {
	s.IPArrayAttribute = &v
	return s
}

func (s *CreateDBInstanceIPArrayShrinkRequest) SetIPArrayName(v string) *CreateDBInstanceIPArrayShrinkRequest {
	s.IPArrayName = &v
	return s
}

func (s *CreateDBInstanceIPArrayShrinkRequest) SetSecurityIPListShrink(v string) *CreateDBInstanceIPArrayShrinkRequest {
	s.SecurityIPListShrink = &v
	return s
}

func (s *CreateDBInstanceIPArrayShrinkRequest) Validate() error {
	return dara.Validate(s)
}

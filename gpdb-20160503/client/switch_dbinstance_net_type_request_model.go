// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceNetTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *SwitchDBInstanceNetTypeRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *SwitchDBInstanceNetTypeRequest
	GetDBInstanceId() *string
	SetPort(v string) *SwitchDBInstanceNetTypeRequest
	GetPort() *string
}

type SwitchDBInstanceNetTypeRequest struct {
	// The prefix of the custom endpoint.
	//
	// 	- The prefix can contain lowercase letters, digits, and hyphens (-) and must start with a lowercase letter.
	//
	// 	- The prefix can be up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1234
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/2361776.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The port number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s SwitchDBInstanceNetTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceNetTypeRequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *SwitchDBInstanceNetTypeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *SwitchDBInstanceNetTypeRequest) GetPort() *string {
	return s.Port
}

func (s *SwitchDBInstanceNetTypeRequest) SetConnectionStringPrefix(v string) *SwitchDBInstanceNetTypeRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetDBInstanceId(v string) *SwitchDBInstanceNetTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetPort(v string) *SwitchDBInstanceNetTypeRequest {
	s.Port = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) Validate() error {
	return dara.Validate(s)
}

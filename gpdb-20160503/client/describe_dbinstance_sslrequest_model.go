// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceSSLRequest
	GetDBInstanceId() *string
}

type DescribeDBInstanceSSLRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceSSLRequest) SetDBInstanceId(v string) *DescribeDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRebalanceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeRebalanceStatusRequest
	GetDBInstanceId() *string
}

type DescribeRebalanceStatusRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeRebalanceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRebalanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRebalanceStatusRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeRebalanceStatusRequest) SetDBInstanceId(v string) *DescribeRebalanceStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRebalanceStatusRequest) Validate() error {
	return dara.Validate(s)
}

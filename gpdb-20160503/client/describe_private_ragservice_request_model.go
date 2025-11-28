// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateRAGServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribePrivateRAGServiceRequest
	GetDBInstanceId() *string
}

type DescribePrivateRAGServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribePrivateRAGServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateRAGServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribePrivateRAGServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribePrivateRAGServiceRequest) SetDBInstanceId(v string) *DescribePrivateRAGServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribePrivateRAGServiceRequest) Validate() error {
	return dara.Validate(s)
}

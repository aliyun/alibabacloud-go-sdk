// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBVersionRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeDBVersionRequest
	GetOwnerId() *int64
}

type DescribeDBVersionRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDBVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBVersionRequest) SetDBInstanceId(v string) *DescribeDBVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBVersionRequest) SetOwnerId(v int64) *DescribeDBVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBVersionRequest) Validate() error {
	return dara.Validate(s)
}

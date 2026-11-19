// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDatabasesRequest
	GetDBInstanceId() *string
}

type DescribeDatabasesRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp14bf67a76d****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDatabasesRequest) SetDBInstanceId(v string) *DescribeDatabasesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDatabasesRequest) Validate() error {
	return dara.Validate(s)
}

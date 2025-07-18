// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSampleDataRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeSampleDataRequest
	GetOwnerId() *int64
}

type DescribeSampleDataRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeSampleDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSampleDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSampleDataRequest) SetDBInstanceId(v string) *DescribeSampleDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSampleDataRequest) SetOwnerId(v int64) *DescribeSampleDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSampleDataRequest) Validate() error {
	return dara.Validate(s)
}

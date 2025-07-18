// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisDimensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDiagnosisDimensionsRequest
	GetDBInstanceId() *string
}

type DescribeDiagnosisDimensionsRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDiagnosisDimensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisDimensionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDiagnosisDimensionsRequest) SetDBInstanceId(v string) *DescribeDiagnosisDimensionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsRequest) Validate() error {
	return dara.Validate(s)
}

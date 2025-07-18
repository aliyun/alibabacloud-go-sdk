// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportFeaturesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSupportFeaturesRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeSupportFeaturesRequest
	GetOwnerId() *int64
}

type DescribeSupportFeaturesRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeSupportFeaturesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportFeaturesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupportFeaturesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSupportFeaturesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSupportFeaturesRequest) SetDBInstanceId(v string) *DescribeSupportFeaturesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSupportFeaturesRequest) SetOwnerId(v int64) *DescribeSupportFeaturesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSupportFeaturesRequest) Validate() error {
	return dara.Validate(s)
}

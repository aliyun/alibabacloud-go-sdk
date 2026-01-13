// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesPrivateRAGServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeZonesPrivateRAGServiceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeZonesPrivateRAGServiceRequest
	GetRegionId() *string
}

type DescribeZonesPrivateRAGServiceRequest struct {
	// The ID of the instance.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the region.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation get a list of available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeZonesPrivateRAGServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesPrivateRAGServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesPrivateRAGServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeZonesPrivateRAGServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesPrivateRAGServiceRequest) SetDBInstanceId(v string) *DescribeZonesPrivateRAGServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeZonesPrivateRAGServiceRequest) SetRegionId(v string) *DescribeZonesPrivateRAGServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesPrivateRAGServiceRequest) Validate() error {
	return dara.Validate(s)
}

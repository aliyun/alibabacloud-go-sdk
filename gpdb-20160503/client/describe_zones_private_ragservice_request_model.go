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
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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

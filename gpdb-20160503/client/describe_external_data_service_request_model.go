// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExternalDataServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeExternalDataServiceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeExternalDataServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *DescribeExternalDataServiceRequest
	GetServiceId() *string
}

type DescribeExternalDataServiceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DescribeExternalDataServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExternalDataServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeExternalDataServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeExternalDataServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExternalDataServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeExternalDataServiceRequest) SetDBInstanceId(v string) *DescribeExternalDataServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeExternalDataServiceRequest) SetRegionId(v string) *DescribeExternalDataServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExternalDataServiceRequest) SetServiceId(v string) *DescribeExternalDataServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DescribeExternalDataServiceRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingDataServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeStreamingDataServiceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeStreamingDataServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *DescribeStreamingDataServiceRequest
	GetServiceId() *string
}

type DescribeStreamingDataServiceRequest struct {
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
	// cn-beijing
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

func (s DescribeStreamingDataServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingDataServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamingDataServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeStreamingDataServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStreamingDataServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeStreamingDataServiceRequest) SetDBInstanceId(v string) *DescribeStreamingDataServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeStreamingDataServiceRequest) SetRegionId(v string) *DescribeStreamingDataServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStreamingDataServiceRequest) SetServiceId(v string) *DescribeStreamingDataServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DescribeStreamingDataServiceRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeModelServiceRequest
	GetDBInstanceId() *string
	SetModelServiceId(v string) *DescribeModelServiceRequest
	GetModelServiceId() *string
}

type DescribeModelServiceRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The model service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mx-xxxxxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
}

func (s DescribeModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelServiceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeModelServiceRequest) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *DescribeModelServiceRequest) SetDBInstanceId(v string) *DescribeModelServiceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeModelServiceRequest) SetModelServiceId(v string) *DescribeModelServiceRequest {
	s.ModelServiceId = &v
	return s
}

func (s *DescribeModelServiceRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsSuperAccountInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstType(v string) *DescribeRdsSuperAccountInstancesRequest
	GetDbInstType() *string
	SetDrdsInstanceId(v string) *DescribeRdsSuperAccountInstancesRequest
	GetDrdsInstanceId() *string
	SetRdsInstance(v []*string) *DescribeRdsSuperAccountInstancesRequest
	GetRdsInstance() []*string
}

type DescribeRdsSuperAccountInstancesRequest struct {
	// The type of the ApsaraDB RDS for MySQL instances. Default value: **RDS**.
	//
	// example:
	//
	// RDS
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [\\"rm-*****************\\",\\"rm-*****************\\"]
	RdsInstance []*string `json:"RdsInstance,omitempty" xml:"RdsInstance,omitempty" type:"Repeated"`
}

func (s DescribeRdsSuperAccountInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesRequest) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeRdsSuperAccountInstancesRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeRdsSuperAccountInstancesRequest) GetRdsInstance() []*string {
	return s.RdsInstance
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetDbInstType(v string) *DescribeRdsSuperAccountInstancesRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetDrdsInstanceId(v string) *DescribeRdsSuperAccountInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesRequest) SetRdsInstance(v []*string) *DescribeRdsSuperAccountInstancesRequest {
	s.RdsInstance = v
	return s
}

func (s *DescribeRdsSuperAccountInstancesRequest) Validate() error {
	return dara.Validate(s)
}

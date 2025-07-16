// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceViaEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v string) *DescribeDBInstanceViaEndpointRequest
	GetEndpoint() *string
	SetRegionId(v string) *DescribeDBInstanceViaEndpointRequest
	GetRegionId() *string
}

type DescribeDBInstanceViaEndpointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hz*******.polarx.rds.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceViaEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeDBInstanceViaEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceViaEndpointRequest) SetEndpoint(v string) *DescribeDBInstanceViaEndpointRequest {
	s.Endpoint = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointRequest) SetRegionId(v string) *DescribeDBInstanceViaEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointRequest) Validate() error {
	return dara.Validate(s)
}

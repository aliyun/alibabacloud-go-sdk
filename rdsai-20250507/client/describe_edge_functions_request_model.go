// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeFunctionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeEdgeFunctionsRequest
	GetClientToken() *string
	SetEdgeFunctionName(v string) *DescribeEdgeFunctionsRequest
	GetEdgeFunctionName() *string
	SetInstanceName(v string) *DescribeEdgeFunctionsRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeEdgeFunctionsRequest
	GetRegionId() *string
}

type DescribeEdgeFunctionsRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// fc-xxxx
	//
	// example:
	//
	// ef-****
	EdgeFunctionName *string `json:"EdgeFunctionName,omitempty" xml:"EdgeFunctionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEdgeFunctionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeFunctionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEdgeFunctionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeEdgeFunctionsRequest) GetEdgeFunctionName() *string {
	return s.EdgeFunctionName
}

func (s *DescribeEdgeFunctionsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeEdgeFunctionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEdgeFunctionsRequest) SetClientToken(v string) *DescribeEdgeFunctionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeEdgeFunctionsRequest) SetEdgeFunctionName(v string) *DescribeEdgeFunctionsRequest {
	s.EdgeFunctionName = &v
	return s
}

func (s *DescribeEdgeFunctionsRequest) SetInstanceName(v string) *DescribeEdgeFunctionsRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeEdgeFunctionsRequest) SetRegionId(v string) *DescribeEdgeFunctionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEdgeFunctionsRequest) Validate() error {
	return dara.Validate(s)
}

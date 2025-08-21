// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoCcListCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAutoCcListCountRequest
	GetInstanceId() *string
	SetQueryType(v string) *DescribeAutoCcListCountRequest
	GetQueryType() *string
}

type DescribeAutoCcListCountRequest struct {
	// The ID of the instance.
	//
	// > You can call the **DescribeInstanceIds*	- operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mode of how an IP address is added to the whitelist or blacklist. Valid values:
	//
	// 	- **manual**: manually added
	//
	// 	- **auto**: automatically added
	//
	// example:
	//
	// manual
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
}

func (s DescribeAutoCcListCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcListCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcListCountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAutoCcListCountRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeAutoCcListCountRequest) SetInstanceId(v string) *DescribeAutoCcListCountRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAutoCcListCountRequest) SetQueryType(v string) *DescribeAutoCcListCountRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeAutoCcListCountRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *DescribeInstanceTypeRequest
	GetInstanceType() *string
}

type DescribeInstanceTypeRequest struct {
	// The instance specification name. For more information, see [Instance node specifications](https://help.aliyun.com/document_detail/194870.html).
	//
	// > If InstanceType is left empty, all instance specifications are returned.
	//
	// example:
	//
	// hbase.n2.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeInstanceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceTypeRequest) SetInstanceType(v string) *DescribeInstanceTypeRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceTypeRequest) Validate() error {
	return dara.Validate(s)
}

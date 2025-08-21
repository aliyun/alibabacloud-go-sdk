// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreatePrePaidInstanceResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCreatePrePaidInstanceResultRequest
	GetInstanceId() *string
}

type DescribeCreatePrePaidInstanceResultRequest struct {
	// The ID of the instance. You can call the CreateIntance operation to create an instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-6ecpqvkicnchxccozrp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCreatePrePaidInstanceResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreatePrePaidInstanceResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreatePrePaidInstanceResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCreatePrePaidInstanceResultRequest) SetInstanceId(v string) *DescribeCreatePrePaidInstanceResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCreatePrePaidInstanceResultRequest) Validate() error {
	return dara.Validate(s)
}

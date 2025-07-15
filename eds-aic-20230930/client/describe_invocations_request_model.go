// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeInvocationsRequest
	GetInstanceIds() []*string
	SetInvocationId(v string) *DescribeInvocationsRequest
	GetInvocationId() *string
}

type DescribeInvocationsRequest struct {
	// The IDs of the cloud phone instances. You can specify a maximum of 50 cloud phone instances.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the execution. You can retrieve the output of a command once by using either the execution ID or the cloud phone instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-4e98eeb5****
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeInvocationsRequest) GetInvocationId() *string {
	return s.InvocationId
}

func (s *DescribeInvocationsRequest) SetInstanceIds(v []*string) *DescribeInvocationsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeInvocationsRequest) SetInvocationId(v string) *DescribeInvocationsRequest {
	s.InvocationId = &v
	return s
}

func (s *DescribeInvocationsRequest) Validate() error {
	return dara.Validate(s)
}

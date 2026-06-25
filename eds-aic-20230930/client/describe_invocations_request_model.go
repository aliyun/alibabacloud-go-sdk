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
	// The IDs of the cloud phone instances. A single request can query the execution results for up to 50 instances.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the command execution. Use this ID and the cloud phone instance ID to query the result of a command execution.
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

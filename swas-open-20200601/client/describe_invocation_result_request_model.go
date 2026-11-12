// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInvocationResultRequest
	GetInstanceId() *string
	SetInvokeId(v string) *DescribeInvocationResultRequest
	GetInvokeId() *string
	SetRegionId(v string) *DescribeInvocationResultRequest
	GetRegionId() *string
}

type DescribeInvocationResultRequest struct {
	// The instance ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The command execution ID. You can call the [DescribeInvocations](https://help.aliyun.com/document_detail/439368.html) operation to query the InvokeId.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-bj02prjhw1n****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInvocationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInvocationResultRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInvocationResultRequest) SetInstanceId(v string) *DescribeInvocationResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationResultRequest) SetInvokeId(v string) *DescribeInvocationResultRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationResultRequest) SetRegionId(v string) *DescribeInvocationResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInvocationResultRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerId(v string) *DescribeInstanceLogRequest
	GetContainerId() *string
	SetInstanceId(v string) *DescribeInstanceLogRequest
	GetInstanceId() *string
}

type DescribeInstanceLogRequest struct {
	// The ID of the sidecar container. You can call the [DescribeApplicationInstances](https://help.aliyun.com/document_detail/2834847.html) to obtain the ID.
	//
	// example:
	//
	// sidecar-test-01
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******-d700e680-aa4d-4ec1-afc2-6566b5ff4d7a-85d44d4bfc-*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLogRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeInstanceLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceLogRequest) SetContainerId(v string) *DescribeInstanceLogRequest {
	s.ContainerId = &v
	return s
}

func (s *DescribeInstanceLogRequest) SetInstanceId(v string) *DescribeInstanceLogRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceLogRequest) Validate() error {
	return dara.Validate(s)
}

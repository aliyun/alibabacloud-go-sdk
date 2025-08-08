// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCurrentNodeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCurrentNodeInfoRequest
	GetInstanceId() *string
}

type DescribeCurrentNodeInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4****89
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCurrentNodeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCurrentNodeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCurrentNodeInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCurrentNodeInfoRequest) SetInstanceId(v string) *DescribeCurrentNodeInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCurrentNodeInfoRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeProjectInfoRequest
	GetInstanceId() *string
}

type DescribeProjectInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4****89
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeProjectInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeProjectInfoRequest) SetInstanceId(v string) *DescribeProjectInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProjectInfoRequest) Validate() error {
	return dara.Validate(s)
}

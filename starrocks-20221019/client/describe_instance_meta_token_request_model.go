// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMetaTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceMetaTokenRequest
	GetInstanceId() *string
}

type DescribeInstanceMetaTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceMetaTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMetaTokenRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMetaTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceMetaTokenRequest) SetInstanceId(v string) *DescribeInstanceMetaTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMetaTokenRequest) Validate() error {
	return dara.Validate(s)
}

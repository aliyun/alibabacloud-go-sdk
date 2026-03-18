// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemTimezoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSystemTimezoneRequest
	GetInstanceId() *string
}

type DescribeSystemTimezoneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeSystemTimezoneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemTimezoneRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemTimezoneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSystemTimezoneRequest) SetInstanceId(v string) *DescribeSystemTimezoneRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSystemTimezoneRequest) Validate() error {
	return dara.Validate(s)
}

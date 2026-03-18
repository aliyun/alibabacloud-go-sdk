// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeEventNamesRequest
	GetInstanceId() *string
}

type DescribeEventNamesRequest struct {
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeEventNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventNamesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEventNamesRequest) SetInstanceId(v string) *DescribeEventNamesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEventNamesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAccessControlListRequest
	GetInstanceId() *string
}

type DescribeAccessControlListRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeAccessControlListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAccessControlListRequest) SetInstanceId(v string) *DescribeAccessControlListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAccessControlListRequest) Validate() error {
	return dara.Validate(s)
}

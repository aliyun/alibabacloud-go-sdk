// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBuyVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeUserBuyVersionRequest
	GetInstanceId() *string
}

type DescribeUserBuyVersionRequest struct {
	// Instance ID. If the Instance ID is provided, the query will be based on this ID. If not provided, the latest instance will be queried by default.
	//
	// example:
	//
	// cfw_elasticity_public_cn-*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeUserBuyVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBuyVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserBuyVersionRequest) SetInstanceId(v string) *DescribeUserBuyVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserBuyVersionRequest) Validate() error {
	return dara.Validate(s)
}

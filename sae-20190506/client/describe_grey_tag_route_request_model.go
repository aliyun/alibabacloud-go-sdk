// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGreyTagRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGreyTagRouteId(v int64) *DescribeGreyTagRouteRequest
	GetGreyTagRouteId() *int64
}

type DescribeGreyTagRouteRequest struct {
	// The ID of the canary release rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
}

func (s DescribeGreyTagRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGreyTagRouteRequest) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteRequest) GetGreyTagRouteId() *int64 {
	return s.GreyTagRouteId
}

func (s *DescribeGreyTagRouteRequest) SetGreyTagRouteId(v int64) *DescribeGreyTagRouteRequest {
	s.GreyTagRouteId = &v
	return s
}

func (s *DescribeGreyTagRouteRequest) Validate() error {
	return dara.Validate(s)
}

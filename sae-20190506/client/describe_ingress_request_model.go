// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIngressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIngressId(v int64) *DescribeIngressRequest
	GetIngressId() *int64
}

type DescribeIngressRequest struct {
	// The ID of the routing rule to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 87
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
}

func (s DescribeIngressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIngressRequest) GoString() string {
	return s.String()
}

func (s *DescribeIngressRequest) GetIngressId() *int64 {
	return s.IngressId
}

func (s *DescribeIngressRequest) SetIngressId(v int64) *DescribeIngressRequest {
	s.IngressId = &v
	return s
}

func (s *DescribeIngressRequest) Validate() error {
	return dara.Validate(s)
}

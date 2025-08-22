// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserErStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeUserErStatusRequest
	GetOwnerId() *int64
}

type DescribeUserErStatusRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeUserErStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserErStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserErStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserErStatusRequest) SetOwnerId(v int64) *DescribeUserErStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserErStatusRequest) Validate() error {
	return dara.Validate(s)
}

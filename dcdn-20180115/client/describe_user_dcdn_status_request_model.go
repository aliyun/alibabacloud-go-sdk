// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDcdnStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeUserDcdnStatusRequest
	GetOwnerId() *int64
}

type DescribeUserDcdnStatusRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeUserDcdnStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDcdnStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserDcdnStatusRequest) SetOwnerId(v int64) *DescribeUserDcdnStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserDcdnStatusRequest) Validate() error {
	return dara.Validate(s)
}

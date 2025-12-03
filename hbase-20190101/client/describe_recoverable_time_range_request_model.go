// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecoverableTimeRangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeRecoverableTimeRangeRequest
	GetClusterId() *string
}

type DescribeRecoverableTimeRangeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeRecoverableTimeRangeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecoverableTimeRangeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableTimeRangeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRecoverableTimeRangeRequest) SetClusterId(v string) *DescribeRecoverableTimeRangeRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRecoverableTimeRangeRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegId(v string) *DescribePrivateStackRequest
	GetRegId() *string
}

type DescribePrivateStackRequest struct {
	// Region Code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribePrivateStackRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateStackRequest) GoString() string {
	return s.String()
}

func (s *DescribePrivateStackRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribePrivateStackRequest) SetRegId(v string) *DescribePrivateStackRequest {
	s.RegId = &v
	return s
}

func (s *DescribePrivateStackRequest) Validate() error {
	return dara.Validate(s)
}

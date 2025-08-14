// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAppKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegId(v string) *DescribeServiceAppKeyRequest
	GetRegId() *string
}

type DescribeServiceAppKeyRequest struct {
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeServiceAppKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAppKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceAppKeyRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeServiceAppKeyRequest) SetRegId(v string) *DescribeServiceAppKeyRequest {
	s.RegId = &v
	return s
}

func (s *DescribeServiceAppKeyRequest) Validate() error {
	return dara.Validate(s)
}

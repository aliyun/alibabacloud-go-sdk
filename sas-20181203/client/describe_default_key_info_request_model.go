// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultKeyInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeDefaultKeyInfoRequest
	GetSourceIp() *string
}

type DescribeDefaultKeyInfoRequest struct {
	// The source IP address.
	//
	// example:
	//
	// 58.246.73.***
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDefaultKeyInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultKeyInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefaultKeyInfoRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDefaultKeyInfoRequest) SetSourceIp(v string) *DescribeDefaultKeyInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDefaultKeyInfoRequest) Validate() error {
	return dara.Validate(s)
}

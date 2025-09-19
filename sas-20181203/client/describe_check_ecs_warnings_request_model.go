// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckEcsWarningsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeCheckEcsWarningsRequest
	GetSourceIp() *string
}

type DescribeCheckEcsWarningsRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.3.4
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeCheckEcsWarningsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckEcsWarningsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckEcsWarningsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCheckEcsWarningsRequest) SetSourceIp(v string) *DescribeCheckEcsWarningsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCheckEcsWarningsRequest) Validate() error {
	return dara.Validate(s)
}

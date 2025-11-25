// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedDictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeCustomizedDictRequest
	GetSourceIp() *string
}

type DescribeCustomizedDictRequest struct {
	// The IP address of the access source.
	//
	// example:
	//
	// 58.240.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeCustomizedDictRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedDictRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedDictRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCustomizedDictRequest) SetSourceIp(v string) *DescribeCustomizedDictRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCustomizedDictRequest) Validate() error {
	return dara.Validate(s)
}

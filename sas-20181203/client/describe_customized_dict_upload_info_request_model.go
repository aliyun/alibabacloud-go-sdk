// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedDictUploadInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeCustomizedDictUploadInfoRequest
	GetSourceIp() *string
}

type DescribeCustomizedDictUploadInfoRequest struct {
	// The source IP address.
	//
	// example:
	//
	// 106.11.43.***
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeCustomizedDictUploadInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedDictUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedDictUploadInfoRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCustomizedDictUploadInfoRequest) SetSourceIp(v string) *DescribeCustomizedDictUploadInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCustomizedDictUploadInfoRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckStatusListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *DescribeHealthCheckStatusListRequest
	GetListeners() *string
	SetSourceIp(v string) *DescribeHealthCheckStatusListRequest
	GetSourceIp() *string
}

type DescribeHealthCheckStatusListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"ddoscoo-cn-XXXXX","Protocol":"tcp","FrontendPort":80}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeHealthCheckStatusListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckStatusListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListRequest) GetListeners() *string {
	return s.Listeners
}

func (s *DescribeHealthCheckStatusListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeHealthCheckStatusListRequest) SetListeners(v string) *DescribeHealthCheckStatusListRequest {
	s.Listeners = &v
	return s
}

func (s *DescribeHealthCheckStatusListRequest) SetSourceIp(v string) *DescribeHealthCheckStatusListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHealthCheckStatusListRequest) Validate() error {
	return dara.Validate(s)
}

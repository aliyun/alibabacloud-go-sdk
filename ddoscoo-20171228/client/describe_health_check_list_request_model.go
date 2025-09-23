// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *DescribeHealthCheckListRequest
	GetListeners() *string
	SetSourceIp(v string) *DescribeHealthCheckListRequest
	GetSourceIp() *string
}

type DescribeHealthCheckListRequest struct {
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

func (s DescribeHealthCheckListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListRequest) GetListeners() *string {
	return s.Listeners
}

func (s *DescribeHealthCheckListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeHealthCheckListRequest) SetListeners(v string) *DescribeHealthCheckListRequest {
	s.Listeners = &v
	return s
}

func (s *DescribeHealthCheckListRequest) SetSourceIp(v string) *DescribeHealthCheckListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHealthCheckListRequest) Validate() error {
	return dara.Validate(s)
}

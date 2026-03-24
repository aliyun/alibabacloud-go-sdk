// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRumDnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *RumDnsResponse
	GetDomain() *string
	SetMessage(v string) *RumDnsResponse
	GetMessage() *string
	SetResult(v bool) *RumDnsResponse
	GetResult() *bool
}

type RumDnsResponse struct {
	// RUM上报域名
	//
	// example:
	//
	// rum
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// 初始化message（失败场景）
	//
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 域名DNS初始化状态
	//
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RumDnsResponse) String() string {
	return dara.Prettify(s)
}

func (s RumDnsResponse) GoString() string {
	return s.String()
}

func (s *RumDnsResponse) GetDomain() *string {
	return s.Domain
}

func (s *RumDnsResponse) GetMessage() *string {
	return s.Message
}

func (s *RumDnsResponse) GetResult() *bool {
	return s.Result
}

func (s *RumDnsResponse) SetDomain(v string) *RumDnsResponse {
	s.Domain = &v
	return s
}

func (s *RumDnsResponse) SetMessage(v string) *RumDnsResponse {
	s.Message = &v
	return s
}

func (s *RumDnsResponse) SetResult(v bool) *RumDnsResponse {
	s.Result = &v
	return s
}

func (s *RumDnsResponse) Validate() error {
	return dara.Validate(s)
}

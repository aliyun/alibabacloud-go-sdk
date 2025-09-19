// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListAuthorizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeWhiteListAuthorizeRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeWhiteListAuthorizeRequest
	GetSourceIp() *string
}

type DescribeWhiteListAuthorizeRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 180.119.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeWhiteListAuthorizeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListAuthorizeRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListAuthorizeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWhiteListAuthorizeRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWhiteListAuthorizeRequest) SetLang(v string) *DescribeWhiteListAuthorizeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhiteListAuthorizeRequest) SetSourceIp(v string) *DescribeWhiteListAuthorizeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhiteListAuthorizeRequest) Validate() error {
	return dara.Validate(s)
}

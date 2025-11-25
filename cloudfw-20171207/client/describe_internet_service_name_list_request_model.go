// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetServiceNameListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeInternetServiceNameListRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeInternetServiceNameListRequest
	GetSourceIp() *string
}

type DescribeInternetServiceNameListRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 140.240.17.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInternetServiceNameListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetServiceNameListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetServiceNameListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetServiceNameListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetServiceNameListRequest) SetLang(v string) *DescribeInternetServiceNameListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetServiceNameListRequest) SetSourceIp(v string) *DescribeInternetServiceNameListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetServiceNameListRequest) Validate() error {
	return dara.Validate(s)
}

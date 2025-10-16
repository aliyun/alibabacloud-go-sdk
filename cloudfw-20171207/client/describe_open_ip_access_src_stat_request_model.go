// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenIpAccessSrcStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeOpenIpAccessSrcStatRequest
	GetCurrentPage() *string
	SetDstIp(v string) *DescribeOpenIpAccessSrcStatRequest
	GetDstIp() *string
	SetLang(v string) *DescribeOpenIpAccessSrcStatRequest
	GetLang() *string
	SetPageSize(v string) *DescribeOpenIpAccessSrcStatRequest
	GetPageSize() *string
	SetSourceIp(v string) *DescribeOpenIpAccessSrcStatRequest
	GetSourceIp() *string
}

type DescribeOpenIpAccessSrcStatRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 47.100.102.XXX
	DstIp *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 47.100.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeOpenIpAccessSrcStatRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenIpAccessSrcStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenIpAccessSrcStatRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeOpenIpAccessSrcStatRequest) GetDstIp() *string {
	return s.DstIp
}

func (s *DescribeOpenIpAccessSrcStatRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOpenIpAccessSrcStatRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeOpenIpAccessSrcStatRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeOpenIpAccessSrcStatRequest) SetCurrentPage(v string) *DescribeOpenIpAccessSrcStatRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatRequest) SetDstIp(v string) *DescribeOpenIpAccessSrcStatRequest {
	s.DstIp = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatRequest) SetLang(v string) *DescribeOpenIpAccessSrcStatRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatRequest) SetPageSize(v string) *DescribeOpenIpAccessSrcStatRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatRequest) SetSourceIp(v string) *DescribeOpenIpAccessSrcStatRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatRequest) Validate() error {
	return dara.Validate(s)
}

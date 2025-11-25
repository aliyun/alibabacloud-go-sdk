// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVfwIPSConfigListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVfwIPSConfigListRequest
	GetLang() *string
	SetPageNo(v int32) *DescribeVfwIPSConfigListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeVfwIPSConfigListRequest
	GetPageSize() *int32
	SetVpcFirewallId(v string) *DescribeVfwIPSConfigListRequest
	GetVpcFirewallId() *string
}

type DescribeVfwIPSConfigListRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVfwIPSConfigListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVfwIPSConfigListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVfwIPSConfigListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVfwIPSConfigListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeVfwIPSConfigListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVfwIPSConfigListRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVfwIPSConfigListRequest) SetLang(v string) *DescribeVfwIPSConfigListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVfwIPSConfigListRequest) SetPageNo(v int32) *DescribeVfwIPSConfigListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeVfwIPSConfigListRequest) SetPageSize(v int32) *DescribeVfwIPSConfigListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVfwIPSConfigListRequest) SetVpcFirewallId(v string) *DescribeVfwIPSConfigListRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVfwIPSConfigListRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeInternetSlbRequest
	GetCurrentPage() *string
	SetInstanceId(v string) *DescribeInternetSlbRequest
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeInternetSlbRequest
	GetInstanceName() *string
	SetIpProtocol(v string) *DescribeInternetSlbRequest
	GetIpProtocol() *string
	SetLang(v string) *DescribeInternetSlbRequest
	GetLang() *string
	SetPageSize(v string) *DescribeInternetSlbRequest
	GetPageSize() *string
	SetPort(v string) *DescribeInternetSlbRequest
	GetPort() *string
	SetPublicIp(v string) *DescribeInternetSlbRequest
	GetPublicIp() *string
	SetRegionNo(v string) *DescribeInternetSlbRequest
	GetRegionNo() *string
	SetSourceIp(v string) *DescribeInternetSlbRequest
	GetSourceIp() *string
	SetTag(v string) *DescribeInternetSlbRequest
	GetTag() *string
}

type DescribeInternetSlbRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// lb-2ze8v2x5kd9qyvp2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
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
	// 63389
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 47.108.60.XXX
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 112.64.233.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeInternetSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetSlbRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetSlbRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeInternetSlbRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInternetSlbRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInternetSlbRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeInternetSlbRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetSlbRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInternetSlbRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeInternetSlbRequest) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeInternetSlbRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeInternetSlbRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetSlbRequest) GetTag() *string {
	return s.Tag
}

func (s *DescribeInternetSlbRequest) SetCurrentPage(v string) *DescribeInternetSlbRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetSlbRequest) SetInstanceId(v string) *DescribeInternetSlbRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInternetSlbRequest) SetInstanceName(v string) *DescribeInternetSlbRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInternetSlbRequest) SetIpProtocol(v string) *DescribeInternetSlbRequest {
	s.IpProtocol = &v
	return s
}

func (s *DescribeInternetSlbRequest) SetLang(v string) *DescribeInternetSlbRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetSlbRequest) SetPageSize(v string) *DescribeInternetSlbRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetSlbRequest) SetPort(v string) *DescribeInternetSlbRequest {
	s.Port = &v
	return s
}

func (s *DescribeInternetSlbRequest) SetPublicIp(v string) *DescribeInternetSlbRequest {
	s.PublicIp = &v
	return s
}

func (s *DescribeInternetSlbRequest) SetRegionNo(v string) *DescribeInternetSlbRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeInternetSlbRequest) SetSourceIp(v string) *DescribeInternetSlbRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetSlbRequest) SetTag(v string) *DescribeInternetSlbRequest {
	s.Tag = &v
	return s
}

func (s *DescribeInternetSlbRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclBackupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeAclBackupListRequest
	GetCurrentPage() *string
	SetLang(v string) *DescribeAclBackupListRequest
	GetLang() *string
	SetPageSize(v string) *DescribeAclBackupListRequest
	GetPageSize() *string
	SetSourceIp(v string) *DescribeAclBackupListRequest
	GetSourceIp() *string
}

type DescribeAclBackupListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 110.191.179.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAclBackupListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclBackupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclBackupListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeAclBackupListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAclBackupListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAclBackupListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAclBackupListRequest) SetCurrentPage(v string) *DescribeAclBackupListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAclBackupListRequest) SetLang(v string) *DescribeAclBackupListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAclBackupListRequest) SetPageSize(v string) *DescribeAclBackupListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAclBackupListRequest) SetSourceIp(v string) *DescribeAclBackupListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAclBackupListRequest) Validate() error {
	return dara.Validate(s)
}

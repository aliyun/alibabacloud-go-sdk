// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePageDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePageDocumentsRequest
	GetLang() *string
	SetPageName(v string) *DescribePageDocumentsRequest
	GetPageName() *string
	SetSourceCode(v string) *DescribePageDocumentsRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribePageDocumentsRequest
	GetSourceIp() *string
	SetTabName(v string) *DescribePageDocumentsRequest
	GetTabName() *string
}

type DescribePageDocumentsRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// overview
	PageName *string `json:"PageName,omitempty" xml:"PageName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yundun
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// example:
	//
	// 223.167.221.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// internet
	TabName *string `json:"TabName,omitempty" xml:"TabName,omitempty"`
}

func (s DescribePageDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePageDocumentsRequest) GoString() string {
	return s.String()
}

func (s *DescribePageDocumentsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePageDocumentsRequest) GetPageName() *string {
	return s.PageName
}

func (s *DescribePageDocumentsRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribePageDocumentsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribePageDocumentsRequest) GetTabName() *string {
	return s.TabName
}

func (s *DescribePageDocumentsRequest) SetLang(v string) *DescribePageDocumentsRequest {
	s.Lang = &v
	return s
}

func (s *DescribePageDocumentsRequest) SetPageName(v string) *DescribePageDocumentsRequest {
	s.PageName = &v
	return s
}

func (s *DescribePageDocumentsRequest) SetSourceCode(v string) *DescribePageDocumentsRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribePageDocumentsRequest) SetSourceIp(v string) *DescribePageDocumentsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePageDocumentsRequest) SetTabName(v string) *DescribePageDocumentsRequest {
	s.TabName = &v
	return s
}

func (s *DescribePageDocumentsRequest) Validate() error {
	return dara.Validate(s)
}

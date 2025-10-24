// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDnsContent(v *DescribeVerifyContentResponseBodyDnsContent) *DescribeVerifyContentResponseBody
	GetDnsContent() *DescribeVerifyContentResponseBodyDnsContent
	SetFileContent(v *DescribeVerifyContentResponseBodyFileContent) *DescribeVerifyContentResponseBody
	GetFileContent() *DescribeVerifyContentResponseBodyFileContent
	SetRequestId(v string) *DescribeVerifyContentResponseBody
	GetRequestId() *string
	SetVerifyResult(v bool) *DescribeVerifyContentResponseBody
	GetVerifyResult() *bool
}

type DescribeVerifyContentResponseBody struct {
	DnsContent  *DescribeVerifyContentResponseBodyDnsContent  `json:"DnsContent,omitempty" xml:"DnsContent,omitempty" type:"Struct"`
	FileContent *DescribeVerifyContentResponseBodyFileContent `json:"FileContent,omitempty" xml:"FileContent,omitempty" type:"Struct"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	VerifyResult *bool `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribeVerifyContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyContentResponseBody) GetDnsContent() *DescribeVerifyContentResponseBodyDnsContent {
	return s.DnsContent
}

func (s *DescribeVerifyContentResponseBody) GetFileContent() *DescribeVerifyContentResponseBodyFileContent {
	return s.FileContent
}

func (s *DescribeVerifyContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyContentResponseBody) GetVerifyResult() *bool {
	return s.VerifyResult
}

func (s *DescribeVerifyContentResponseBody) SetDnsContent(v *DescribeVerifyContentResponseBodyDnsContent) *DescribeVerifyContentResponseBody {
	s.DnsContent = v
	return s
}

func (s *DescribeVerifyContentResponseBody) SetFileContent(v *DescribeVerifyContentResponseBodyFileContent) *DescribeVerifyContentResponseBody {
	s.FileContent = v
	return s
}

func (s *DescribeVerifyContentResponseBody) SetRequestId(v string) *DescribeVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyContentResponseBody) SetVerifyResult(v bool) *DescribeVerifyContentResponseBody {
	s.VerifyResult = &v
	return s
}

func (s *DescribeVerifyContentResponseBody) Validate() error {
	if s.DnsContent != nil {
		if err := s.DnsContent.Validate(); err != nil {
			return err
		}
	}
	if s.FileContent != nil {
		if err := s.FileContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyContentResponseBodyDnsContent struct {
	// example:
	//
	// verification
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// example:
	//
	// TXT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// verify_0a246ca99d504ba087472d***
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVerifyContentResponseBodyDnsContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyContentResponseBodyDnsContent) GoString() string {
	return s.String()
}

func (s *DescribeVerifyContentResponseBodyDnsContent) GetRR() *string {
	return s.RR
}

func (s *DescribeVerifyContentResponseBodyDnsContent) GetType() *string {
	return s.Type
}

func (s *DescribeVerifyContentResponseBodyDnsContent) GetValue() *string {
	return s.Value
}

func (s *DescribeVerifyContentResponseBodyDnsContent) SetRR(v string) *DescribeVerifyContentResponseBodyDnsContent {
	s.RR = &v
	return s
}

func (s *DescribeVerifyContentResponseBodyDnsContent) SetType(v string) *DescribeVerifyContentResponseBodyDnsContent {
	s.Type = &v
	return s
}

func (s *DescribeVerifyContentResponseBodyDnsContent) SetValue(v string) *DescribeVerifyContentResponseBodyDnsContent {
	s.Value = &v
	return s
}

func (s *DescribeVerifyContentResponseBodyDnsContent) Validate() error {
	return dara.Validate(s)
}

type DescribeVerifyContentResponseBodyFileContent struct {
	// example:
	//
	// http://oss.xxx.com//xxx.html
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// xxx.html
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// aliyundemo.com
	TopDomain *string `json:"TopDomain,omitempty" xml:"TopDomain,omitempty"`
	// example:
	//
	// verify_0a246ca99d504ba08***
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// http://xxx.xxx.com//xxx.html
	VerifyPath *string `json:"VerifyPath,omitempty" xml:"VerifyPath,omitempty"`
}

func (s DescribeVerifyContentResponseBodyFileContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyContentResponseBodyFileContent) GoString() string {
	return s.String()
}

func (s *DescribeVerifyContentResponseBodyFileContent) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeVerifyContentResponseBodyFileContent) GetFileName() *string {
	return s.FileName
}

func (s *DescribeVerifyContentResponseBodyFileContent) GetTopDomain() *string {
	return s.TopDomain
}

func (s *DescribeVerifyContentResponseBodyFileContent) GetValue() *string {
	return s.Value
}

func (s *DescribeVerifyContentResponseBodyFileContent) GetVerifyPath() *string {
	return s.VerifyPath
}

func (s *DescribeVerifyContentResponseBodyFileContent) SetDownloadUrl(v string) *DescribeVerifyContentResponseBodyFileContent {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeVerifyContentResponseBodyFileContent) SetFileName(v string) *DescribeVerifyContentResponseBodyFileContent {
	s.FileName = &v
	return s
}

func (s *DescribeVerifyContentResponseBodyFileContent) SetTopDomain(v string) *DescribeVerifyContentResponseBodyFileContent {
	s.TopDomain = &v
	return s
}

func (s *DescribeVerifyContentResponseBodyFileContent) SetValue(v string) *DescribeVerifyContentResponseBodyFileContent {
	s.Value = &v
	return s
}

func (s *DescribeVerifyContentResponseBodyFileContent) SetVerifyPath(v string) *DescribeVerifyContentResponseBodyFileContent {
	s.VerifyPath = &v
	return s
}

func (s *DescribeVerifyContentResponseBodyFileContent) Validate() error {
	return dara.Validate(s)
}

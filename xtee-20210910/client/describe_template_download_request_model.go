// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateDownloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeTemplateDownloadRequest
	GetLang() *string
	SetRegId(v string) *DescribeTemplateDownloadRequest
	GetRegId() *string
}

type DescribeTemplateDownloadRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeTemplateDownloadRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateDownloadRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateDownloadRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTemplateDownloadRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTemplateDownloadRequest) SetLang(v string) *DescribeTemplateDownloadRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTemplateDownloadRequest) SetRegId(v string) *DescribeTemplateDownloadRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTemplateDownloadRequest) Validate() error {
	return dara.Validate(s)
}

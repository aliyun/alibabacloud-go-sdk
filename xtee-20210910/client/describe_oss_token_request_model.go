// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeOssTokenRequest
	GetLang() *string
	SetFileName(v string) *DescribeOssTokenRequest
	GetFileName() *string
	SetRegId(v string) *DescribeOssTokenRequest
	GetRegId() *string
	SetUploadType(v string) *DescribeOssTokenRequest
	GetUploadType() *string
}

type DescribeOssTokenRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// File name.
	//
	// example:
	//
	// test.csv
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Upload type
	//
	// example:
	//
	// COMMUNITY_SAMPLE
	UploadType *string `json:"uploadType,omitempty" xml:"uploadType,omitempty"`
}

func (s DescribeOssTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssTokenRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssTokenRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOssTokenRequest) GetFileName() *string {
	return s.FileName
}

func (s *DescribeOssTokenRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeOssTokenRequest) GetUploadType() *string {
	return s.UploadType
}

func (s *DescribeOssTokenRequest) SetLang(v string) *DescribeOssTokenRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssTokenRequest) SetFileName(v string) *DescribeOssTokenRequest {
	s.FileName = &v
	return s
}

func (s *DescribeOssTokenRequest) SetRegId(v string) *DescribeOssTokenRequest {
	s.RegId = &v
	return s
}

func (s *DescribeOssTokenRequest) SetUploadType(v string) *DescribeOssTokenRequest {
	s.UploadType = &v
	return s
}

func (s *DescribeOssTokenRequest) Validate() error {
	return dara.Validate(s)
}

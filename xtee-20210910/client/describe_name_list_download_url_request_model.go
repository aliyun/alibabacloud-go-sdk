// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNameListDownloadUrlRequest
	GetLang() *string
	SetRegId(v string) *DescribeNameListDownloadUrlRequest
	GetRegId() *string
	SetVariableId(v int64) *DescribeNameListDownloadUrlRequest
	GetVariableId() *int64
}

type DescribeNameListDownloadUrlRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
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
	// Variable ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 762
	VariableId *int64 `json:"variableId,omitempty" xml:"variableId,omitempty"`
}

func (s DescribeNameListDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeNameListDownloadUrlRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNameListDownloadUrlRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeNameListDownloadUrlRequest) GetVariableId() *int64 {
	return s.VariableId
}

func (s *DescribeNameListDownloadUrlRequest) SetLang(v string) *DescribeNameListDownloadUrlRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNameListDownloadUrlRequest) SetRegId(v string) *DescribeNameListDownloadUrlRequest {
	s.RegId = &v
	return s
}

func (s *DescribeNameListDownloadUrlRequest) SetVariableId(v int64) *DescribeNameListDownloadUrlRequest {
	s.VariableId = &v
	return s
}

func (s *DescribeNameListDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}

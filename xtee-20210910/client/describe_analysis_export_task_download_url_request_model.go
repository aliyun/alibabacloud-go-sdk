// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisExportTaskDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAnalysisExportTaskDownloadUrlRequest
	GetLang() *string
	SetRegId(v string) *DescribeAnalysisExportTaskDownloadUrlRequest
	GetRegId() *string
}

type DescribeAnalysisExportTaskDownloadUrlRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAnalysisExportTaskDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisExportTaskDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisExportTaskDownloadUrlRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAnalysisExportTaskDownloadUrlRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAnalysisExportTaskDownloadUrlRequest) SetLang(v string) *DescribeAnalysisExportTaskDownloadUrlRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAnalysisExportTaskDownloadUrlRequest) SetRegId(v string) *DescribeAnalysisExportTaskDownloadUrlRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAnalysisExportTaskDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}

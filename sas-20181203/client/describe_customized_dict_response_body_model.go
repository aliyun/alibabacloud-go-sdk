// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedDictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeCustomizedDictResponseBody
	GetCount() *int32
	SetFileName(v string) *DescribeCustomizedDictResponseBody
	GetFileName() *string
	SetFileSize(v int64) *DescribeCustomizedDictResponseBody
	GetFileSize() *int64
	SetLimit(v int32) *DescribeCustomizedDictResponseBody
	GetLimit() *int32
	SetOssUrl(v string) *DescribeCustomizedDictResponseBody
	GetOssUrl() *string
	SetRequestId(v string) *DescribeCustomizedDictResponseBody
	GetRequestId() *string
	SetUploadTime(v int64) *DescribeCustomizedDictResponseBody
	GetUploadTime() *int64
}

type DescribeCustomizedDictResponseBody struct {
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// test_dict.plain
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 40
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// 9
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// cloudtest01/661767e1-5ae3-4ec5-865f-03039436893a/sacc2*****
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-XXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1719919421
	UploadTime *int64 `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s DescribeCustomizedDictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedDictResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedDictResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCustomizedDictResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *DescribeCustomizedDictResponseBody) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeCustomizedDictResponseBody) GetLimit() *int32 {
	return s.Limit
}

func (s *DescribeCustomizedDictResponseBody) GetOssUrl() *string {
	return s.OssUrl
}

func (s *DescribeCustomizedDictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomizedDictResponseBody) GetUploadTime() *int64 {
	return s.UploadTime
}

func (s *DescribeCustomizedDictResponseBody) SetCount(v int32) *DescribeCustomizedDictResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCustomizedDictResponseBody) SetFileName(v string) *DescribeCustomizedDictResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeCustomizedDictResponseBody) SetFileSize(v int64) *DescribeCustomizedDictResponseBody {
	s.FileSize = &v
	return s
}

func (s *DescribeCustomizedDictResponseBody) SetLimit(v int32) *DescribeCustomizedDictResponseBody {
	s.Limit = &v
	return s
}

func (s *DescribeCustomizedDictResponseBody) SetOssUrl(v string) *DescribeCustomizedDictResponseBody {
	s.OssUrl = &v
	return s
}

func (s *DescribeCustomizedDictResponseBody) SetRequestId(v string) *DescribeCustomizedDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizedDictResponseBody) SetUploadTime(v int64) *DescribeCustomizedDictResponseBody {
	s.UploadTime = &v
	return s
}

func (s *DescribeCustomizedDictResponseBody) Validate() error {
	return dara.Validate(s)
}

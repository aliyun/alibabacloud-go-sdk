// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobGroupExportTaskProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeJobGroupExportTaskProgressResponseBody
	GetCode() *string
	SetFileHttpUrl(v string) *DescribeJobGroupExportTaskProgressResponseBody
	GetFileHttpUrl() *string
	SetHttpStatusCode(v int32) *DescribeJobGroupExportTaskProgressResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeJobGroupExportTaskProgressResponseBody
	GetMessage() *string
	SetProgress(v float32) *DescribeJobGroupExportTaskProgressResponseBody
	GetProgress() *float32
	SetRequestId(v string) *DescribeJobGroupExportTaskProgressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeJobGroupExportTaskProgressResponseBody
	GetSuccess() *bool
}

type DescribeJobGroupExportTaskProgressResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// File URL
	//
	// example:
	//
	// http://xxx.xx.com/xx
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Export progress
	//
	// example:
	//
	// 45.0
	Progress *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeJobGroupExportTaskProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupExportTaskProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) GetFileHttpUrl() *string {
	return s.FileHttpUrl
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) GetProgress() *float32 {
	return s.Progress
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) SetCode(v string) *DescribeJobGroupExportTaskProgressResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) SetFileHttpUrl(v string) *DescribeJobGroupExportTaskProgressResponseBody {
	s.FileHttpUrl = &v
	return s
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) SetHttpStatusCode(v int32) *DescribeJobGroupExportTaskProgressResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) SetMessage(v string) *DescribeJobGroupExportTaskProgressResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) SetProgress(v float32) *DescribeJobGroupExportTaskProgressResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) SetRequestId(v string) *DescribeJobGroupExportTaskProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) SetSuccess(v bool) *DescribeJobGroupExportTaskProgressResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobGroupExportTaskProgressResponseBody) Validate() error {
	return dara.Validate(s)
}

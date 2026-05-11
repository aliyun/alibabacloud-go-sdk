// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileHttpUrl(v string) *DescribeExportProgressResponseBody
	GetFileHttpUrl() *string
	SetRequestId(v string) *DescribeExportProgressResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeExportProgressResponseBody
	GetStatus() *string
}

type DescribeExportProgressResponseBody struct {
	// example:
	//
	// http://ssml-test.oss-cn-shanghai.aliyuncs.com/key
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
	// example:
	//
	// b19af5ce5314ac08108d1b33fe20e15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExportProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExportProgressResponseBody) GetFileHttpUrl() *string {
	return s.FileHttpUrl
}

func (s *DescribeExportProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExportProgressResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeExportProgressResponseBody) SetFileHttpUrl(v string) *DescribeExportProgressResponseBody {
	s.FileHttpUrl = &v
	return s
}

func (s *DescribeExportProgressResponseBody) SetRequestId(v string) *DescribeExportProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExportProgressResponseBody) SetStatus(v string) *DescribeExportProgressResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeExportProgressResponseBody) Validate() error {
	return dara.Validate(s)
}

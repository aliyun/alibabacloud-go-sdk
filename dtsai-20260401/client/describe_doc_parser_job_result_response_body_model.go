// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DescribeDocParserJobResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeDocParserJobResultResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DescribeDocParserJobResultResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeDocParserJobResultResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeDocParserJobResultResponseBody
	GetResult() *string
	SetSuccess(v bool) *DescribeDocParserJobResultResponseBody
	GetSuccess() *bool
}

type DescribeDocParserJobResultResponseBody struct {
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result         *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDocParserJobResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDocParserJobResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDocParserJobResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDocParserJobResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDocParserJobResultResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeDocParserJobResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDocParserJobResultResponseBody) SetErrorCode(v string) *DescribeDocParserJobResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetErrorMessage(v string) *DescribeDocParserJobResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetHttpStatusCode(v int32) *DescribeDocParserJobResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetRequestId(v string) *DescribeDocParserJobResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetResult(v string) *DescribeDocParserJobResultResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetSuccess(v bool) *DescribeDocParserJobResultResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) Validate() error {
	return dara.Validate(s)
}

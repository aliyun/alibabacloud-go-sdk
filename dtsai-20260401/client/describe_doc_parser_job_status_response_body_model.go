// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DescribeDocParserJobStatusResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeDocParserJobStatusResponseBody
	GetErrorMessage() *string
	SetFailureMessage(v string) *DescribeDocParserJobStatusResponseBody
	GetFailureMessage() *string
	SetHttpStatusCode(v int32) *DescribeDocParserJobStatusResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeDocParserJobStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeDocParserJobStatusResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DescribeDocParserJobStatusResponseBody
	GetSuccess() *bool
}

type DescribeDocParserJobStatusResponseBody struct {
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FailureMessage *string `json:"FailureMessage,omitempty" xml:"FailureMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDocParserJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDocParserJobStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDocParserJobStatusResponseBody) GetFailureMessage() *string {
	return s.FailureMessage
}

func (s *DescribeDocParserJobStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDocParserJobStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDocParserJobStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDocParserJobStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDocParserJobStatusResponseBody) SetErrorCode(v string) *DescribeDocParserJobStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDocParserJobStatusResponseBody) SetErrorMessage(v string) *DescribeDocParserJobStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDocParserJobStatusResponseBody) SetFailureMessage(v string) *DescribeDocParserJobStatusResponseBody {
	s.FailureMessage = &v
	return s
}

func (s *DescribeDocParserJobStatusResponseBody) SetHttpStatusCode(v int32) *DescribeDocParserJobStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDocParserJobStatusResponseBody) SetRequestId(v string) *DescribeDocParserJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDocParserJobStatusResponseBody) SetStatus(v string) *DescribeDocParserJobStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDocParserJobStatusResponseBody) SetSuccess(v bool) *DescribeDocParserJobStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDocParserJobStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

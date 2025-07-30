// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DescribeDocParserJobStatusResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDocParserJobStatusResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeDocParserJobStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDocParserJobStatusResponseBody
	GetErrMessage() *string
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
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// FDC111B1-ACBF-457D-9656-247FDEE9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDocParserJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobStatusResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDocParserJobStatusResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDocParserJobStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDocParserJobStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
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

func (s *DescribeDocParserJobStatusResponseBody) SetDynamicCode(v string) *DescribeDocParserJobStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDocParserJobStatusResponseBody) SetDynamicMessage(v string) *DescribeDocParserJobStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDocParserJobStatusResponseBody) SetErrCode(v string) *DescribeDocParserJobStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDocParserJobStatusResponseBody) SetErrMessage(v string) *DescribeDocParserJobStatusResponseBody {
	s.ErrMessage = &v
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

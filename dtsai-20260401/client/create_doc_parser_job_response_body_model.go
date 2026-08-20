// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocParserJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateDocParserJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDocParserJobResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateDocParserJobResponseBody
	GetHttpStatusCode() *int32
	SetJobId(v string) *CreateDocParserJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateDocParserJobResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateDocParserJobResponseBody
	GetResult() *string
	SetResultType(v string) *CreateDocParserJobResponseBody
	GetResultType() *string
	SetResultUrl(v string) *CreateDocParserJobResponseBody
	GetResultUrl() *string
	SetStatus(v string) *CreateDocParserJobResponseBody
	GetStatus() *string
	SetSuccess(v bool) *CreateDocParserJobResponseBody
	GetSuccess() *bool
}

type CreateDocParserJobResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The document parsing task ID.
	//
	// example:
	//
	// job_abc123
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A30D0930-xxxx-xxxx-xxxx-C2C661CC8B58
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     *string `json:"Result,omitempty" xml:"Result,omitempty"`
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	ResultUrl  *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDocParserJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDocParserJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDocParserJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDocParserJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateDocParserJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDocParserJobResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateDocParserJobResponseBody) GetResultType() *string {
	return s.ResultType
}

func (s *CreateDocParserJobResponseBody) GetResultUrl() *string {
	return s.ResultUrl
}

func (s *CreateDocParserJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateDocParserJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDocParserJobResponseBody) SetErrorCode(v string) *CreateDocParserJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetErrorMessage(v string) *CreateDocParserJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetHttpStatusCode(v int32) *CreateDocParserJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetJobId(v string) *CreateDocParserJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetRequestId(v string) *CreateDocParserJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetResult(v string) *CreateDocParserJobResponseBody {
	s.Result = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetResultType(v string) *CreateDocParserJobResponseBody {
	s.ResultType = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetResultUrl(v string) *CreateDocParserJobResponseBody {
	s.ResultUrl = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetStatus(v string) *CreateDocParserJobResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetSuccess(v bool) *CreateDocParserJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDocParserJobResponseBody) Validate() error {
	return dara.Validate(s)
}

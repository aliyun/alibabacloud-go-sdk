// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitScriptReviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitScriptReviewResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SubmitScriptReviewResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitScriptReviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitScriptReviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitScriptReviewResponseBody
	GetSuccess() *bool
}

type SubmitScriptReviewResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s SubmitScriptReviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitScriptReviewResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitScriptReviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitScriptReviewResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitScriptReviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitScriptReviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitScriptReviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitScriptReviewResponseBody) SetCode(v string) *SubmitScriptReviewResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitScriptReviewResponseBody) SetHttpStatusCode(v int32) *SubmitScriptReviewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitScriptReviewResponseBody) SetMessage(v string) *SubmitScriptReviewResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitScriptReviewResponseBody) SetRequestId(v string) *SubmitScriptReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitScriptReviewResponseBody) SetSuccess(v bool) *SubmitScriptReviewResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitScriptReviewResponseBody) Validate() error {
	return dara.Validate(s)
}

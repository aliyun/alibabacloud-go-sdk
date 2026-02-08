// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteGlobalQuestionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteGlobalQuestionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGlobalQuestionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGlobalQuestionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGlobalQuestionResponseBody
	GetSuccess() *bool
}

type DeleteGlobalQuestionResponseBody struct {
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
	// API message
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

func (s DeleteGlobalQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGlobalQuestionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGlobalQuestionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGlobalQuestionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGlobalQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGlobalQuestionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGlobalQuestionResponseBody) SetCode(v string) *DeleteGlobalQuestionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGlobalQuestionResponseBody) SetHttpStatusCode(v int32) *DeleteGlobalQuestionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGlobalQuestionResponseBody) SetMessage(v string) *DeleteGlobalQuestionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGlobalQuestionResponseBody) SetRequestId(v string) *DeleteGlobalQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGlobalQuestionResponseBody) SetSuccess(v bool) *DeleteGlobalQuestionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGlobalQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}

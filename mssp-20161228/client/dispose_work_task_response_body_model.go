// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisposeWorkTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisposeWorkTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DisposeWorkTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisposeWorkTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisposeWorkTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisposeWorkTaskResponseBody
	GetSuccess() *bool
}

type DisposeWorkTaskResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the returned result.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86786E4C-6416-55CF-9AB6-5E275B68801D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisposeWorkTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisposeWorkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisposeWorkTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisposeWorkTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisposeWorkTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisposeWorkTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisposeWorkTaskResponseBody) SetCode(v string) *DisposeWorkTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetHttpStatusCode(v int32) *DisposeWorkTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetMessage(v string) *DisposeWorkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetRequestId(v string) *DisposeWorkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetSuccess(v bool) *DisposeWorkTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAuthTokenResponseBody
	GetCode() *int32
	SetData(v string) *GetAuthTokenResponseBody
	GetData() *string
	SetMessage(v string) *GetAuthTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAuthTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAuthTokenResponseBody
	GetSuccess() *bool
}

type GetAuthTokenResponseBody struct {
	// Status code. 200 means success, other status codes are exceptions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned authentication token.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiJ9******
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Additional Information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the query was successful:
	//
	// - true: Success
	//
	// - false: Failure
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAuthTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAuthTokenResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAuthTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAuthTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAuthTokenResponseBody) SetCode(v int32) *GetAuthTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetData(v string) *GetAuthTokenResponseBody {
	s.Data = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetMessage(v string) *GetAuthTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetRequestId(v string) *GetAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetSuccess(v bool) *GetAuthTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetAuthTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

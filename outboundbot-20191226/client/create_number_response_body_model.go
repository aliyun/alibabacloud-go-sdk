// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateNumberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateNumberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateNumberResponseBody
	GetSuccess() *bool
}

type CreateNumberResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API operation.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNumberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateNumberResponseBody) SetCode(v string) *CreateNumberResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNumberResponseBody) SetHttpStatusCode(v int32) *CreateNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateNumberResponseBody) SetMessage(v string) *CreateNumberResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNumberResponseBody) SetRequestId(v string) *CreateNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNumberResponseBody) SetSuccess(v bool) *CreateNumberResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNumberResponseBody) Validate() error {
	return dara.Validate(s)
}

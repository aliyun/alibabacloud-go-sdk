// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteApiKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteApiKeyResponseBody
	GetSuccess() *bool
}

type DeleteApiKeyResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 43ED6679-672D-5EB6-B3AC-E0EF90129DFF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteApiKeyResponseBody) SetCode(v string) *DeleteApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApiKeyResponseBody) SetHttpStatusCode(v int32) *DeleteApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteApiKeyResponseBody) SetMessage(v string) *DeleteApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApiKeyResponseBody) SetRequestId(v string) *DeleteApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiKeyResponseBody) SetSuccess(v bool) *DeleteApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

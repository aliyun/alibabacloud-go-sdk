// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteStandardResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteStandardResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteStandardResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteStandardResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStandardResponseBody
	GetSuccess() *bool
}

type DeleteStandardResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteStandardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStandardResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteStandardResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteStandardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteStandardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStandardResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStandardResponseBody) SetCode(v string) *DeleteStandardResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStandardResponseBody) SetHttpStatusCode(v int32) *DeleteStandardResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteStandardResponseBody) SetMessage(v string) *DeleteStandardResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStandardResponseBody) SetRequestId(v string) *DeleteStandardResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStandardResponseBody) SetSuccess(v bool) *DeleteStandardResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStandardResponseBody) Validate() error {
	return dara.Validate(s)
}

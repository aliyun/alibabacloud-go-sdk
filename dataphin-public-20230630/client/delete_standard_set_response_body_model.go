// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteStandardSetResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteStandardSetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteStandardSetResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteStandardSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStandardSetResponseBody
	GetSuccess() *bool
}

type DeleteStandardSetResponseBody struct {
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

func (s DeleteStandardSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStandardSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteStandardSetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteStandardSetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteStandardSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStandardSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStandardSetResponseBody) SetCode(v string) *DeleteStandardSetResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStandardSetResponseBody) SetHttpStatusCode(v int32) *DeleteStandardSetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteStandardSetResponseBody) SetMessage(v string) *DeleteStandardSetResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStandardSetResponseBody) SetRequestId(v string) *DeleteStandardSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStandardSetResponseBody) SetSuccess(v bool) *DeleteStandardSetResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStandardSetResponseBody) Validate() error {
	return dara.Validate(s)
}

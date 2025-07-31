// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteResourceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteResourceResponseBody
	GetSuccess() *bool
}

type DeleteResourceResponseBody struct {
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

func (s DeleteResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteResourceResponseBody) SetCode(v string) *DeleteResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteResourceResponseBody) SetHttpStatusCode(v int32) *DeleteResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteResourceResponseBody) SetMessage(v string) *DeleteResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceResponseBody) SetSuccess(v bool) *DeleteResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

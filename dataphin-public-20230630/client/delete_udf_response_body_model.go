// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteUdfResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteUdfResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteUdfResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteUdfResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUdfResponseBody
	GetSuccess() *bool
}

type DeleteUdfResponseBody struct {
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

func (s DeleteUdfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUdfResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteUdfResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteUdfResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUdfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUdfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUdfResponseBody) SetCode(v string) *DeleteUdfResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUdfResponseBody) SetHttpStatusCode(v int32) *DeleteUdfResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteUdfResponseBody) SetMessage(v string) *DeleteUdfResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUdfResponseBody) SetRequestId(v string) *DeleteUdfResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUdfResponseBody) SetSuccess(v bool) *DeleteUdfResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUdfResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegisterLineageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRegisterLineageResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteRegisterLineageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteRegisterLineageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRegisterLineageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRegisterLineageResponseBody
	GetSuccess() *bool
}

type DeleteRegisterLineageResponseBody struct {
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

func (s DeleteRegisterLineageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegisterLineageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRegisterLineageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRegisterLineageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteRegisterLineageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRegisterLineageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRegisterLineageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRegisterLineageResponseBody) SetCode(v string) *DeleteRegisterLineageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRegisterLineageResponseBody) SetHttpStatusCode(v int32) *DeleteRegisterLineageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRegisterLineageResponseBody) SetMessage(v string) *DeleteRegisterLineageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRegisterLineageResponseBody) SetRequestId(v string) *DeleteRegisterLineageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRegisterLineageResponseBody) SetSuccess(v bool) *DeleteRegisterLineageResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRegisterLineageResponseBody) Validate() error {
	return dara.Validate(s)
}

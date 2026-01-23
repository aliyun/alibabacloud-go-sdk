// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardWordRootResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteStandardWordRootResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteStandardWordRootResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteStandardWordRootResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteStandardWordRootResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStandardWordRootResponseBody
	GetSuccess() *bool
}

type DeleteStandardWordRootResponseBody struct {
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

func (s DeleteStandardWordRootResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardWordRootResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStandardWordRootResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteStandardWordRootResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteStandardWordRootResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteStandardWordRootResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStandardWordRootResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStandardWordRootResponseBody) SetCode(v string) *DeleteStandardWordRootResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStandardWordRootResponseBody) SetHttpStatusCode(v int32) *DeleteStandardWordRootResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteStandardWordRootResponseBody) SetMessage(v string) *DeleteStandardWordRootResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStandardWordRootResponseBody) SetRequestId(v string) *DeleteStandardWordRootResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStandardWordRootResponseBody) SetSuccess(v bool) *DeleteStandardWordRootResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStandardWordRootResponseBody) Validate() error {
	return dara.Validate(s)
}

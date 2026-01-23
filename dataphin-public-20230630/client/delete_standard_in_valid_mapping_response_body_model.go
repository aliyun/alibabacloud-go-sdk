// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardInValidMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteStandardInValidMappingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteStandardInValidMappingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteStandardInValidMappingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteStandardInValidMappingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStandardInValidMappingResponseBody
	GetSuccess() *bool
}

type DeleteStandardInValidMappingResponseBody struct {
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

func (s DeleteStandardInValidMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardInValidMappingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStandardInValidMappingResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteStandardInValidMappingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteStandardInValidMappingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteStandardInValidMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStandardInValidMappingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStandardInValidMappingResponseBody) SetCode(v string) *DeleteStandardInValidMappingResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStandardInValidMappingResponseBody) SetHttpStatusCode(v int32) *DeleteStandardInValidMappingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteStandardInValidMappingResponseBody) SetMessage(v string) *DeleteStandardInValidMappingResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStandardInValidMappingResponseBody) SetRequestId(v string) *DeleteStandardInValidMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStandardInValidMappingResponseBody) SetSuccess(v bool) *DeleteStandardInValidMappingResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStandardInValidMappingResponseBody) Validate() error {
	return dara.Validate(s)
}

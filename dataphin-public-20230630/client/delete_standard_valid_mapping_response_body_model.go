// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardValidMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteStandardValidMappingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteStandardValidMappingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteStandardValidMappingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteStandardValidMappingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStandardValidMappingResponseBody
	GetSuccess() *bool
}

type DeleteStandardValidMappingResponseBody struct {
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

func (s DeleteStandardValidMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardValidMappingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStandardValidMappingResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteStandardValidMappingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteStandardValidMappingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteStandardValidMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStandardValidMappingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStandardValidMappingResponseBody) SetCode(v string) *DeleteStandardValidMappingResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStandardValidMappingResponseBody) SetHttpStatusCode(v int32) *DeleteStandardValidMappingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteStandardValidMappingResponseBody) SetMessage(v string) *DeleteStandardValidMappingResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStandardValidMappingResponseBody) SetRequestId(v string) *DeleteStandardValidMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStandardValidMappingResponseBody) SetSuccess(v bool) *DeleteStandardValidMappingResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStandardValidMappingResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardMappingToInvalidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateStandardMappingToInvalidResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateStandardMappingToInvalidResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateStandardMappingToInvalidResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateStandardMappingToInvalidResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateStandardMappingToInvalidResponseBody
	GetSuccess() *bool
}

type UpdateStandardMappingToInvalidResponseBody struct {
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

func (s UpdateStandardMappingToInvalidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardMappingToInvalidResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStandardMappingToInvalidResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateStandardMappingToInvalidResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateStandardMappingToInvalidResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateStandardMappingToInvalidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStandardMappingToInvalidResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateStandardMappingToInvalidResponseBody) SetCode(v string) *UpdateStandardMappingToInvalidResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateStandardMappingToInvalidResponseBody) SetHttpStatusCode(v int32) *UpdateStandardMappingToInvalidResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateStandardMappingToInvalidResponseBody) SetMessage(v string) *UpdateStandardMappingToInvalidResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateStandardMappingToInvalidResponseBody) SetRequestId(v string) *UpdateStandardMappingToInvalidResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStandardMappingToInvalidResponseBody) SetSuccess(v bool) *UpdateStandardMappingToInvalidResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateStandardMappingToInvalidResponseBody) Validate() error {
	return dara.Validate(s)
}

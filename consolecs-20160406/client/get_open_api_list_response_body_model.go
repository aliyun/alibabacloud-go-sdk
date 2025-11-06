// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenApiListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOpenApiListResponseBody
	GetCode() *string
	SetMessage(v string) *GetOpenApiListResponseBody
	GetMessage() *string
	SetOpenApiString(v string) *GetOpenApiListResponseBody
	GetOpenApiString() *string
	SetRequestId(v string) *GetOpenApiListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOpenApiListResponseBody
	GetSuccess() *bool
}

type GetOpenApiListResponseBody struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OpenApiString *string `json:"OpenApiString,omitempty" xml:"OpenApiString,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOpenApiListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpenApiListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenApiListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOpenApiListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOpenApiListResponseBody) GetOpenApiString() *string {
	return s.OpenApiString
}

func (s *GetOpenApiListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpenApiListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOpenApiListResponseBody) SetCode(v string) *GetOpenApiListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpenApiListResponseBody) SetMessage(v string) *GetOpenApiListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpenApiListResponseBody) SetOpenApiString(v string) *GetOpenApiListResponseBody {
	s.OpenApiString = &v
	return s
}

func (s *GetOpenApiListResponseBody) SetRequestId(v string) *GetOpenApiListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpenApiListResponseBody) SetSuccess(v bool) *GetOpenApiListResponseBody {
	s.Success = &v
	return s
}

func (s *GetOpenApiListResponseBody) Validate() error {
	return dara.Validate(s)
}

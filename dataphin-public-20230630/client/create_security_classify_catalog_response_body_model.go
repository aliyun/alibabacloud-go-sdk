// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityClassifyCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSecurityClassifyCatalogResponseBody
	GetCode() *string
	SetData(v string) *CreateSecurityClassifyCatalogResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateSecurityClassifyCatalogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateSecurityClassifyCatalogResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSecurityClassifyCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSecurityClassifyCatalogResponseBody
	GetSuccess() *bool
}

type CreateSecurityClassifyCatalogResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// /d1/d2/
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateSecurityClassifyCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyCatalogResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSecurityClassifyCatalogResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateSecurityClassifyCatalogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSecurityClassifyCatalogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSecurityClassifyCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecurityClassifyCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSecurityClassifyCatalogResponseBody) SetCode(v string) *CreateSecurityClassifyCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSecurityClassifyCatalogResponseBody) SetData(v string) *CreateSecurityClassifyCatalogResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSecurityClassifyCatalogResponseBody) SetHttpStatusCode(v int32) *CreateSecurityClassifyCatalogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSecurityClassifyCatalogResponseBody) SetMessage(v string) *CreateSecurityClassifyCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSecurityClassifyCatalogResponseBody) SetRequestId(v string) *CreateSecurityClassifyCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityClassifyCatalogResponseBody) SetSuccess(v bool) *CreateSecurityClassifyCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSecurityClassifyCatalogResponseBody) Validate() error {
	return dara.Validate(s)
}

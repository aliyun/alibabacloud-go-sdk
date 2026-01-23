// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityClassifyCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSecurityClassifyCatalogResponseBody
	GetCode() *string
	SetData(v string) *UpdateSecurityClassifyCatalogResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateSecurityClassifyCatalogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSecurityClassifyCatalogResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSecurityClassifyCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSecurityClassifyCatalogResponseBody
	GetSuccess() *bool
}

type UpdateSecurityClassifyCatalogResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// /d1/d3/
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

func (s UpdateSecurityClassifyCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyCatalogResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSecurityClassifyCatalogResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateSecurityClassifyCatalogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSecurityClassifyCatalogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSecurityClassifyCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecurityClassifyCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSecurityClassifyCatalogResponseBody) SetCode(v string) *UpdateSecurityClassifyCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogResponseBody) SetData(v string) *UpdateSecurityClassifyCatalogResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogResponseBody) SetHttpStatusCode(v int32) *UpdateSecurityClassifyCatalogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogResponseBody) SetMessage(v string) *UpdateSecurityClassifyCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogResponseBody) SetRequestId(v string) *UpdateSecurityClassifyCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogResponseBody) SetSuccess(v bool) *UpdateSecurityClassifyCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogResponseBody) Validate() error {
	return dara.Validate(s)
}

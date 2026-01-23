// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityClassifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSecurityClassifyResponseBody
	GetCode() *string
	SetData(v int64) *UpdateSecurityClassifyResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateSecurityClassifyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSecurityClassifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSecurityClassifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSecurityClassifyResponseBody
	GetSuccess() *bool
}

type UpdateSecurityClassifyResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s UpdateSecurityClassifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSecurityClassifyResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateSecurityClassifyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSecurityClassifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSecurityClassifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecurityClassifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSecurityClassifyResponseBody) SetCode(v string) *UpdateSecurityClassifyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSecurityClassifyResponseBody) SetData(v int64) *UpdateSecurityClassifyResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSecurityClassifyResponseBody) SetHttpStatusCode(v int32) *UpdateSecurityClassifyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSecurityClassifyResponseBody) SetMessage(v string) *UpdateSecurityClassifyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSecurityClassifyResponseBody) SetRequestId(v string) *UpdateSecurityClassifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityClassifyResponseBody) SetSuccess(v bool) *UpdateSecurityClassifyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSecurityClassifyResponseBody) Validate() error {
	return dara.Validate(s)
}

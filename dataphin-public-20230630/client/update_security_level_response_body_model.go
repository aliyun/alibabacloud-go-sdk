// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSecurityLevelResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateSecurityLevelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSecurityLevelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSecurityLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSecurityLevelResponseBody
	GetSuccess() *bool
}

type UpdateSecurityLevelResponseBody struct {
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

func (s UpdateSecurityLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityLevelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityLevelResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSecurityLevelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSecurityLevelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSecurityLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecurityLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSecurityLevelResponseBody) SetCode(v string) *UpdateSecurityLevelResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSecurityLevelResponseBody) SetHttpStatusCode(v int32) *UpdateSecurityLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSecurityLevelResponseBody) SetMessage(v string) *UpdateSecurityLevelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSecurityLevelResponseBody) SetRequestId(v string) *UpdateSecurityLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityLevelResponseBody) SetSuccess(v bool) *UpdateSecurityLevelResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSecurityLevelResponseBody) Validate() error {
	return dara.Validate(s)
}

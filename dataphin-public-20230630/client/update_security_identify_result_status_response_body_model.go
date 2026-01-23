// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityIdentifyResultStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSecurityIdentifyResultStatusResponseBody
	GetCode() *string
	SetData(v bool) *UpdateSecurityIdentifyResultStatusResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateSecurityIdentifyResultStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSecurityIdentifyResultStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSecurityIdentifyResultStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSecurityIdentifyResultStatusResponseBody
	GetSuccess() *bool
}

type UpdateSecurityIdentifyResultStatusResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s UpdateSecurityIdentifyResultStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityIdentifyResultStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) SetCode(v string) *UpdateSecurityIdentifyResultStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) SetData(v bool) *UpdateSecurityIdentifyResultStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) SetHttpStatusCode(v int32) *UpdateSecurityIdentifyResultStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) SetMessage(v string) *UpdateSecurityIdentifyResultStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) SetRequestId(v string) *UpdateSecurityIdentifyResultStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) SetSuccess(v bool) *UpdateSecurityIdentifyResultStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

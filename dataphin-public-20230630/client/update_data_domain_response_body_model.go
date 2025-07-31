// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDataDomainResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateDataDomainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDataDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDataDomainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataDomainResponseBody
	GetSuccess() *bool
}

type UpdateDataDomainResponseBody struct {
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDataDomainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDataDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDataDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataDomainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataDomainResponseBody) SetCode(v string) *UpdateDataDomainResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDataDomainResponseBody) SetHttpStatusCode(v int32) *UpdateDataDomainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataDomainResponseBody) SetMessage(v string) *UpdateDataDomainResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDataDomainResponseBody) SetRequestId(v string) *UpdateDataDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataDomainResponseBody) SetSuccess(v bool) *UpdateDataDomainResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

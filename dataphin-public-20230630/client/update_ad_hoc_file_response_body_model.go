// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdHocFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAdHocFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateAdHocFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateAdHocFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAdHocFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAdHocFileResponseBody
	GetSuccess() *bool
}

type UpdateAdHocFileResponseBody struct {
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

func (s UpdateAdHocFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdHocFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdHocFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAdHocFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateAdHocFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAdHocFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAdHocFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAdHocFileResponseBody) SetCode(v string) *UpdateAdHocFileResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAdHocFileResponseBody) SetHttpStatusCode(v int32) *UpdateAdHocFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAdHocFileResponseBody) SetMessage(v string) *UpdateAdHocFileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAdHocFileResponseBody) SetRequestId(v string) *UpdateAdHocFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAdHocFileResponseBody) SetSuccess(v bool) *UpdateAdHocFileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAdHocFileResponseBody) Validate() error {
	return dara.Validate(s)
}

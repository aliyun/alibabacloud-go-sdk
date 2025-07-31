// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateBizEntityResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateBizEntityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateBizEntityResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBizEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBizEntityResponseBody
	GetSuccess() *bool
}

type UpdateBizEntityResponseBody struct {
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

func (s UpdateBizEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateBizEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBizEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBizEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBizEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBizEntityResponseBody) SetCode(v string) *UpdateBizEntityResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBizEntityResponseBody) SetHttpStatusCode(v int32) *UpdateBizEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBizEntityResponseBody) SetMessage(v string) *UpdateBizEntityResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBizEntityResponseBody) SetRequestId(v string) *UpdateBizEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBizEntityResponseBody) SetSuccess(v bool) *UpdateBizEntityResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBizEntityResponseBody) Validate() error {
	return dara.Validate(s)
}

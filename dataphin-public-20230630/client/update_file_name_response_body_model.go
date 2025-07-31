// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFileNameResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateFileNameResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateFileNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateFileNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFileNameResponseBody
	GetSuccess() *bool
}

type UpdateFileNameResponseBody struct {
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

func (s UpdateFileNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFileNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateFileNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateFileNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFileNameResponseBody) SetCode(v string) *UpdateFileNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFileNameResponseBody) SetHttpStatusCode(v int32) *UpdateFileNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateFileNameResponseBody) SetMessage(v string) *UpdateFileNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFileNameResponseBody) SetRequestId(v string) *UpdateFileNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileNameResponseBody) SetSuccess(v bool) *UpdateFileNameResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFileNameResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDataServiceAppResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateDataServiceAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDataServiceAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDataServiceAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataServiceAppResponseBody
	GetSuccess() *bool
}

type UpdateDataServiceAppResponseBody struct {
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

func (s UpdateDataServiceAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDataServiceAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDataServiceAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDataServiceAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataServiceAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataServiceAppResponseBody) SetCode(v string) *UpdateDataServiceAppResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDataServiceAppResponseBody) SetHttpStatusCode(v int32) *UpdateDataServiceAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataServiceAppResponseBody) SetMessage(v string) *UpdateDataServiceAppResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDataServiceAppResponseBody) SetRequestId(v string) *UpdateDataServiceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataServiceAppResponseBody) SetSuccess(v bool) *UpdateDataServiceAppResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataServiceAppResponseBody) Validate() error {
	return dara.Validate(s)
}

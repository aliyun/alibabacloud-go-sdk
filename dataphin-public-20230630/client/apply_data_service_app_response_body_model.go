// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataServiceAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyDataServiceAppResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ApplyDataServiceAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ApplyDataServiceAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyDataServiceAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyDataServiceAppResponseBody
	GetSuccess() *bool
}

type ApplyDataServiceAppResponseBody struct {
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

func (s ApplyDataServiceAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceAppResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyDataServiceAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ApplyDataServiceAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyDataServiceAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyDataServiceAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyDataServiceAppResponseBody) SetCode(v string) *ApplyDataServiceAppResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyDataServiceAppResponseBody) SetHttpStatusCode(v int32) *ApplyDataServiceAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ApplyDataServiceAppResponseBody) SetMessage(v string) *ApplyDataServiceAppResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyDataServiceAppResponseBody) SetRequestId(v string) *ApplyDataServiceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyDataServiceAppResponseBody) SetSuccess(v bool) *ApplyDataServiceAppResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyDataServiceAppResponseBody) Validate() error {
	return dara.Validate(s)
}

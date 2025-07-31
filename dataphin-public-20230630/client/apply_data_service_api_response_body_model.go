// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyDataServiceApiResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ApplyDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ApplyDataServiceApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyDataServiceApiResponseBody
	GetSuccess() *bool
}

type ApplyDataServiceApiResponseBody struct {
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

func (s ApplyDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ApplyDataServiceApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyDataServiceApiResponseBody) SetCode(v string) *ApplyDataServiceApiResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyDataServiceApiResponseBody) SetHttpStatusCode(v int32) *ApplyDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ApplyDataServiceApiResponseBody) SetMessage(v string) *ApplyDataServiceApiResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyDataServiceApiResponseBody) SetRequestId(v string) *ApplyDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyDataServiceApiResponseBody) SetSuccess(v bool) *ApplyDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}

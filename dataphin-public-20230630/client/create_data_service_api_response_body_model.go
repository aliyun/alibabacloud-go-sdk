// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *CreateDataServiceApiResponseBody
	GetApiId() *int64
	SetCode(v string) *CreateDataServiceApiResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateDataServiceApiResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDataServiceApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDataServiceApiResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataServiceApiResponseBody
	GetSuccess() *bool
}

type CreateDataServiceApiResponseBody struct {
	// example:
	//
	// 10086
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
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

func (s CreateDataServiceApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiResponseBody) GetApiId() *int64 {
	return s.ApiId
}

func (s *CreateDataServiceApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDataServiceApiResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDataServiceApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDataServiceApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataServiceApiResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataServiceApiResponseBody) SetApiId(v int64) *CreateDataServiceApiResponseBody {
	s.ApiId = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) SetCode(v string) *CreateDataServiceApiResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) SetHttpStatusCode(v int32) *CreateDataServiceApiResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) SetMessage(v string) *CreateDataServiceApiResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) SetRequestId(v string) *CreateDataServiceApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) SetSuccess(v bool) *CreateDataServiceApiResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataServiceApiResponseBody) Validate() error {
	return dara.Validate(s)
}

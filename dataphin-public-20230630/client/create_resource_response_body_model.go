// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateResourceResponseBody
	GetCode() *string
	SetData(v int64) *CreateResourceResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateResourceResponseBody
	GetSuccess() *bool
}

type CreateResourceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10112101
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateResourceResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateResourceResponseBody) SetCode(v string) *CreateResourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateResourceResponseBody) SetData(v int64) *CreateResourceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateResourceResponseBody) SetHttpStatusCode(v int32) *CreateResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateResourceResponseBody) SetMessage(v string) *CreateResourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceResponseBody) SetSuccess(v bool) *CreateResourceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardWordRootResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStandardWordRootResponseBody
	GetCode() *string
	SetData(v string) *CreateStandardWordRootResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateStandardWordRootResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateStandardWordRootResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateStandardWordRootResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStandardWordRootResponseBody
	GetSuccess() *bool
}

type CreateStandardWordRootResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// average
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateStandardWordRootResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardWordRootResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStandardWordRootResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStandardWordRootResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateStandardWordRootResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateStandardWordRootResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateStandardWordRootResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStandardWordRootResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStandardWordRootResponseBody) SetCode(v string) *CreateStandardWordRootResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStandardWordRootResponseBody) SetData(v string) *CreateStandardWordRootResponseBody {
	s.Data = &v
	return s
}

func (s *CreateStandardWordRootResponseBody) SetHttpStatusCode(v int32) *CreateStandardWordRootResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateStandardWordRootResponseBody) SetMessage(v string) *CreateStandardWordRootResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStandardWordRootResponseBody) SetRequestId(v string) *CreateStandardWordRootResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStandardWordRootResponseBody) SetSuccess(v bool) *CreateStandardWordRootResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStandardWordRootResponseBody) Validate() error {
	return dara.Validate(s)
}

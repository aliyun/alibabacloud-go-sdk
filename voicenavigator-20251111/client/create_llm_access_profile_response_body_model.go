// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLlmAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateLlmAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *CreateLlmAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateLlmAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateLlmAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateLlmAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateLlmAccessProfileResponseBody
	GetRequestId() *string
}

type CreateLlmAccessProfileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 75BAAB9B-40B2-5FF5-A59A-7BCF8154C6EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLlmAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLlmAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLlmAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateLlmAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateLlmAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateLlmAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLlmAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateLlmAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLlmAccessProfileResponseBody) SetCode(v string) *CreateLlmAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetData(v string) *CreateLlmAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetHttpStatusCode(v int32) *CreateLlmAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetMessage(v string) *CreateLlmAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetParams(v []*string) *CreateLlmAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) SetRequestId(v string) *CreateLlmAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLlmAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}

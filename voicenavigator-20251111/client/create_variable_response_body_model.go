// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVariableResponseBody
	GetCode() *string
	SetData(v string) *CreateVariableResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateVariableResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateVariableResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateVariableResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateVariableResponseBody
	GetRequestId() *string
}

type CreateVariableResponseBody struct {
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
	// ED56B723-5802-5C32-865F-6E20B06D2455
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVariableResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVariableResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateVariableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateVariableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVariableResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVariableResponseBody) SetCode(v string) *CreateVariableResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVariableResponseBody) SetData(v string) *CreateVariableResponseBody {
	s.Data = &v
	return s
}

func (s *CreateVariableResponseBody) SetHttpStatusCode(v int32) *CreateVariableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateVariableResponseBody) SetMessage(v string) *CreateVariableResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVariableResponseBody) SetParams(v []*string) *CreateVariableResponseBody {
	s.Params = v
	return s
}

func (s *CreateVariableResponseBody) SetRequestId(v string) *CreateVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVariableResponseBody) Validate() error {
	return dara.Validate(s)
}

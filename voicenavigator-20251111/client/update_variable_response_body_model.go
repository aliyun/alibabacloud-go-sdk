// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVariableResponseBody
	GetCode() *string
	SetData(v string) *UpdateVariableResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateVariableResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateVariableResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateVariableResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateVariableResponseBody
	GetRequestId() *string
}

type UpdateVariableResponseBody struct {
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
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVariableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVariableResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVariableResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateVariableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateVariableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVariableResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVariableResponseBody) SetCode(v string) *UpdateVariableResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVariableResponseBody) SetData(v string) *UpdateVariableResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateVariableResponseBody) SetHttpStatusCode(v int32) *UpdateVariableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVariableResponseBody) SetMessage(v string) *UpdateVariableResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVariableResponseBody) SetParams(v []*string) *UpdateVariableResponseBody {
	s.Params = v
	return s
}

func (s *UpdateVariableResponseBody) SetRequestId(v string) *UpdateVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVariableResponseBody) Validate() error {
	return dara.Validate(s)
}

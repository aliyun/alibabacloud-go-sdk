// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSchemaPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSchemaPropertyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateSchemaPropertyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSchemaPropertyResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateSchemaPropertyResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateSchemaPropertyResponseBody
	GetRequestId() *string
}

type UpdateSchemaPropertyResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// A450574A-337F-43E2-BC59-9C6594C994C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSchemaPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSchemaPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSchemaPropertyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSchemaPropertyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSchemaPropertyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSchemaPropertyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateSchemaPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSchemaPropertyResponseBody) SetCode(v string) *UpdateSchemaPropertyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSchemaPropertyResponseBody) SetHttpStatusCode(v int32) *UpdateSchemaPropertyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSchemaPropertyResponseBody) SetMessage(v string) *UpdateSchemaPropertyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSchemaPropertyResponseBody) SetParams(v []*string) *UpdateSchemaPropertyResponseBody {
	s.Params = v
	return s
}

func (s *UpdateSchemaPropertyResponseBody) SetRequestId(v string) *UpdateSchemaPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSchemaPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}

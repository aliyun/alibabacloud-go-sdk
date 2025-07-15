// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSchemaPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddSchemaPropertyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddSchemaPropertyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddSchemaPropertyResponseBody
	GetMessage() *string
	SetParams(v []*string) *AddSchemaPropertyResponseBody
	GetParams() []*string
	SetRequestId(v string) *AddSchemaPropertyResponseBody
	GetRequestId() *string
}

type AddSchemaPropertyResponseBody struct {
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
	// BF268B34-09C2-43FD-BAC4-5D31EA633111
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSchemaPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSchemaPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *AddSchemaPropertyResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddSchemaPropertyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddSchemaPropertyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSchemaPropertyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *AddSchemaPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSchemaPropertyResponseBody) SetCode(v string) *AddSchemaPropertyResponseBody {
	s.Code = &v
	return s
}

func (s *AddSchemaPropertyResponseBody) SetHttpStatusCode(v int32) *AddSchemaPropertyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddSchemaPropertyResponseBody) SetMessage(v string) *AddSchemaPropertyResponseBody {
	s.Message = &v
	return s
}

func (s *AddSchemaPropertyResponseBody) SetParams(v []*string) *AddSchemaPropertyResponseBody {
	s.Params = v
	return s
}

func (s *AddSchemaPropertyResponseBody) SetRequestId(v string) *AddSchemaPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSchemaPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}

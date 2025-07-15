// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSchemaResponseBody
	GetCode() *string
	SetData(v string) *CreateSchemaResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateSchemaResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateSchemaResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateSchemaResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateSchemaResponseBody
	GetRequestId() *string
}

type CreateSchemaResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1ca2b084-6f0a-454b-9851-29768a9a5832
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 9F766284-F103-4298-8EC5-19F9F9BE5522
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchemaResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSchemaResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateSchemaResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSchemaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSchemaResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSchemaResponseBody) SetCode(v string) *CreateSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSchemaResponseBody) SetData(v string) *CreateSchemaResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSchemaResponseBody) SetHttpStatusCode(v int32) *CreateSchemaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSchemaResponseBody) SetMessage(v string) *CreateSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSchemaResponseBody) SetParams(v []*string) *CreateSchemaResponseBody {
	s.Params = v
	return s
}

func (s *CreateSchemaResponseBody) SetRequestId(v string) *CreateSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}

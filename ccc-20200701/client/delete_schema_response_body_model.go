// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSchemaResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteSchemaResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteSchemaResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteSchemaResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteSchemaResponseBody
	GetRequestId() *string
}

type DeleteSchemaResponseBody struct {
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
	// 678F7002-CA01-4ABF-A112-585AFBDF3A3B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchemaResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSchemaResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteSchemaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSchemaResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSchemaResponseBody) SetCode(v string) *DeleteSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSchemaResponseBody) SetHttpStatusCode(v int32) *DeleteSchemaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSchemaResponseBody) SetMessage(v string) *DeleteSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSchemaResponseBody) SetParams(v []*string) *DeleteSchemaResponseBody {
	s.Params = v
	return s
}

func (s *DeleteSchemaResponseBody) SetRequestId(v string) *DeleteSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}

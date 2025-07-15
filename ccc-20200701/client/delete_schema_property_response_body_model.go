// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemaPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSchemaPropertyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteSchemaPropertyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteSchemaPropertyResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteSchemaPropertyResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteSchemaPropertyResponseBody
	GetRequestId() *string
}

type DeleteSchemaPropertyResponseBody struct {
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
	// User 244715989906081477 does not exist in instance worldfirst01.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 7BEEA660-A45A-45E3-98CC-AFC65E715C23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSchemaPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemaPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchemaPropertyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSchemaPropertyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteSchemaPropertyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSchemaPropertyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteSchemaPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSchemaPropertyResponseBody) SetCode(v string) *DeleteSchemaPropertyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSchemaPropertyResponseBody) SetHttpStatusCode(v int32) *DeleteSchemaPropertyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSchemaPropertyResponseBody) SetMessage(v string) *DeleteSchemaPropertyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSchemaPropertyResponseBody) SetParams(v []*string) *DeleteSchemaPropertyResponseBody {
	s.Params = v
	return s
}

func (s *DeleteSchemaPropertyResponseBody) SetRequestId(v string) *DeleteSchemaPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSchemaPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}

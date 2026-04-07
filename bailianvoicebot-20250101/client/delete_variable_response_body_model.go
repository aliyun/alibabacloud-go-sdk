// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteVariableResponseBody
	GetCode() *string
	SetData(v string) *DeleteVariableResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteVariableResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteVariableResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteVariableResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteVariableResponseBody
	GetRequestId() *string
}

type DeleteVariableResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 82ea16d1-425c-4c03-9be5-cc91de9779ed
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance llm-rj6aqmctjcit4acy does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVariableResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteVariableResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteVariableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteVariableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVariableResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVariableResponseBody) SetCode(v string) *DeleteVariableResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVariableResponseBody) SetData(v string) *DeleteVariableResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteVariableResponseBody) SetHttpStatusCode(v int32) *DeleteVariableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteVariableResponseBody) SetMessage(v string) *DeleteVariableResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVariableResponseBody) SetParams(v []*string) *DeleteVariableResponseBody {
	s.Params = v
	return s
}

func (s *DeleteVariableResponseBody) SetRequestId(v string) *DeleteVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVariableResponseBody) Validate() error {
	return dara.Validate(s)
}

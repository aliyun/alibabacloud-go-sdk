// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateScriptResponseBody
	GetCode() *string
	SetData(v string) *CreateScriptResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateScriptResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateScriptResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateScriptResponseBody
	GetRequestId() *string
}

type CreateScriptResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0a88e269-01f5-49ac-97af-5830f0ccb271
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

func (s CreateScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateScriptResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScriptResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScriptResponseBody) SetCode(v string) *CreateScriptResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScriptResponseBody) SetData(v string) *CreateScriptResponseBody {
	s.Data = &v
	return s
}

func (s *CreateScriptResponseBody) SetHttpStatusCode(v int32) *CreateScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateScriptResponseBody) SetMessage(v string) *CreateScriptResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScriptResponseBody) SetParams(v []*string) *CreateScriptResponseBody {
	s.Params = v
	return s
}

func (s *CreateScriptResponseBody) SetRequestId(v string) *CreateScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScriptResponseBody) Validate() error {
	return dara.Validate(s)
}

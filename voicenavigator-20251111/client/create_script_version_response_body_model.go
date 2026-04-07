// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateScriptVersionResponseBody
	GetCode() *string
	SetData(v string) *CreateScriptVersionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateScriptVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateScriptVersionResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateScriptVersionResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateScriptVersionResponseBody
	GetRequestId() *string
}

type CreateScriptVersionResponseBody struct {
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
	// C00486CA-8025-5ED8-9AA5-E0AFBA87A098
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateScriptVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScriptVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateScriptVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateScriptVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateScriptVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScriptVersionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateScriptVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScriptVersionResponseBody) SetCode(v string) *CreateScriptVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScriptVersionResponseBody) SetData(v string) *CreateScriptVersionResponseBody {
	s.Data = &v
	return s
}

func (s *CreateScriptVersionResponseBody) SetHttpStatusCode(v int32) *CreateScriptVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateScriptVersionResponseBody) SetMessage(v string) *CreateScriptVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScriptVersionResponseBody) SetParams(v []*string) *CreateScriptVersionResponseBody {
	s.Params = v
	return s
}

func (s *CreateScriptVersionResponseBody) SetRequestId(v string) *CreateScriptVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScriptVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

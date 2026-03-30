// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateScriptResponseBody
	GetCode() *string
	SetData(v string) *UpdateScriptResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateScriptResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateScriptResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateScriptResponseBody
	GetRequestId() *string
}

type UpdateScriptResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// F132DDBA-66C4-5BD3-BF7E-9642FE877159
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
	// F132DDBA-66C4-5BD3-BF7E-9642FE877158
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateScriptResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateScriptResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScriptResponseBody) SetCode(v string) *UpdateScriptResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateScriptResponseBody) SetData(v string) *UpdateScriptResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateScriptResponseBody) SetHttpStatusCode(v int32) *UpdateScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateScriptResponseBody) SetMessage(v string) *UpdateScriptResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateScriptResponseBody) SetParams(v []*string) *UpdateScriptResponseBody {
	s.Params = v
	return s
}

func (s *UpdateScriptResponseBody) SetRequestId(v string) *UpdateScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScriptResponseBody) Validate() error {
	return dara.Validate(s)
}

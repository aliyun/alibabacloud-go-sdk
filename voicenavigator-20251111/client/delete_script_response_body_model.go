// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteScriptResponseBody
	GetCode() *string
	SetData(v string) *DeleteScriptResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteScriptResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteScriptResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteScriptResponseBody
	GetRequestId() *string
}

type DeleteScriptResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b090
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
	// D24E0148-6D40-550E-9471-B2C5A34C3D12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteScriptResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteScriptResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScriptResponseBody) SetCode(v string) *DeleteScriptResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScriptResponseBody) SetData(v string) *DeleteScriptResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteScriptResponseBody) SetHttpStatusCode(v int32) *DeleteScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteScriptResponseBody) SetMessage(v string) *DeleteScriptResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScriptResponseBody) SetParams(v []*string) *DeleteScriptResponseBody {
	s.Params = v
	return s
}

func (s *DeleteScriptResponseBody) SetRequestId(v string) *DeleteScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScriptResponseBody) Validate() error {
	return dara.Validate(s)
}

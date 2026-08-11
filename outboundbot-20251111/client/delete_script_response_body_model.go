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
	SetSuccess(v bool) *DeleteScriptResponseBody
	GetSuccess() *bool
}

type DeleteScriptResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b37
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=placeholder-instance-id.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *DeleteScriptResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *DeleteScriptResponseBody) SetSuccess(v bool) *DeleteScriptResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScriptResponseBody) Validate() error {
	return dara.Validate(s)
}

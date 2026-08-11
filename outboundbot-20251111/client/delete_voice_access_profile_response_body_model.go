// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVoiceAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteVoiceAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *DeleteVoiceAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteVoiceAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteVoiceAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteVoiceAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteVoiceAccessProfileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteVoiceAccessProfileResponseBody
	GetSuccess() *bool
}

type DeleteVoiceAccessProfileResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
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
	// Instance does not exist. Instance=ob-9876543210.
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

func (s DeleteVoiceAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVoiceAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVoiceAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteVoiceAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteVoiceAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteVoiceAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVoiceAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteVoiceAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVoiceAccessProfileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteVoiceAccessProfileResponseBody) SetCode(v string) *DeleteVoiceAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVoiceAccessProfileResponseBody) SetData(v string) *DeleteVoiceAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteVoiceAccessProfileResponseBody) SetHttpStatusCode(v int32) *DeleteVoiceAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteVoiceAccessProfileResponseBody) SetMessage(v string) *DeleteVoiceAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVoiceAccessProfileResponseBody) SetParams(v []*string) *DeleteVoiceAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *DeleteVoiceAccessProfileResponseBody) SetRequestId(v string) *DeleteVoiceAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVoiceAccessProfileResponseBody) SetSuccess(v bool) *DeleteVoiceAccessProfileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteVoiceAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}

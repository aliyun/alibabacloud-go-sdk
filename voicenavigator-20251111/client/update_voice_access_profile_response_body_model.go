// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVoiceAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVoiceAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *UpdateVoiceAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateVoiceAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateVoiceAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateVoiceAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateVoiceAccessProfileResponseBody
	GetRequestId() *string
}

type UpdateVoiceAccessProfileResponseBody struct {
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
	// 7401D698-0AAC-5909-B68E-88C68805FFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVoiceAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVoiceAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVoiceAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVoiceAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateVoiceAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateVoiceAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVoiceAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateVoiceAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVoiceAccessProfileResponseBody) SetCode(v string) *UpdateVoiceAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetData(v string) *UpdateVoiceAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetHttpStatusCode(v int32) *UpdateVoiceAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetMessage(v string) *UpdateVoiceAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetParams(v []*string) *UpdateVoiceAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) SetRequestId(v string) *UpdateVoiceAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVoiceAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}

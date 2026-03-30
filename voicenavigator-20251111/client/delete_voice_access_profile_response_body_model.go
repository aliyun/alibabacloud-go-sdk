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
}

type DeleteVoiceAccessProfileResponseBody struct {
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
	// 97E7ED13-6884-52D7-B97E-C780D7870680
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DeleteVoiceAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}

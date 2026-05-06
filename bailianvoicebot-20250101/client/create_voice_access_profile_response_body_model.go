// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVoiceAccessProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVoiceAccessProfileResponseBody
	GetCode() *string
	SetData(v string) *CreateVoiceAccessProfileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateVoiceAccessProfileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateVoiceAccessProfileResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateVoiceAccessProfileResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateVoiceAccessProfileResponseBody
	GetRequestId() *string
}

type CreateVoiceAccessProfileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance llm-zzu528i29ecnprcl does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// CF6D3484-19A1-5C77-863B-AC8B5754D37C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVoiceAccessProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVoiceAccessProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVoiceAccessProfileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVoiceAccessProfileResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateVoiceAccessProfileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateVoiceAccessProfileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVoiceAccessProfileResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateVoiceAccessProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVoiceAccessProfileResponseBody) SetCode(v string) *CreateVoiceAccessProfileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetData(v string) *CreateVoiceAccessProfileResponseBody {
	s.Data = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetHttpStatusCode(v int32) *CreateVoiceAccessProfileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetMessage(v string) *CreateVoiceAccessProfileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetParams(v []*string) *CreateVoiceAccessProfileResponseBody {
	s.Params = v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) SetRequestId(v string) *CreateVoiceAccessProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVoiceAccessProfileResponseBody) Validate() error {
	return dara.Validate(s)
}

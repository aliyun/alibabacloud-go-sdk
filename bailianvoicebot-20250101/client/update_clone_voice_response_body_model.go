// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloneVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCloneVoiceResponseBody
	GetCode() *string
	SetData(v string) *UpdateCloneVoiceResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateCloneVoiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCloneVoiceResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateCloneVoiceResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateCloneVoiceResponseBody
	GetRequestId() *string
}

type UpdateCloneVoiceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 8ee1160a-6999-478f-8df6-f33ef21f27d5
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance llm-xdne77rxe14ziszr
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCloneVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloneVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloneVoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCloneVoiceResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateCloneVoiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCloneVoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCloneVoiceResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateCloneVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloneVoiceResponseBody) SetCode(v string) *UpdateCloneVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetData(v string) *UpdateCloneVoiceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetHttpStatusCode(v int32) *UpdateCloneVoiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetMessage(v string) *UpdateCloneVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetParams(v []*string) *UpdateCloneVoiceResponseBody {
	s.Params = v
	return s
}

func (s *UpdateCloneVoiceResponseBody) SetRequestId(v string) *UpdateCloneVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloneVoiceResponseBody) Validate() error {
	return dara.Validate(s)
}

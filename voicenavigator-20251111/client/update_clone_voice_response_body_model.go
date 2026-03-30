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
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
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

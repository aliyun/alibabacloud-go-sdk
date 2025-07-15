// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAudioFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyAudioFileResponseBody
	GetCode() *string
	SetData(v string) *ModifyAudioFileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ModifyAudioFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyAudioFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAudioFileResponseBody
	GetRequestId() *string
}

type ModifyAudioFileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D2F84AAC-7C79-547F-8EE9-7B735F42B93F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAudioFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAudioFileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAudioFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyAudioFileResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyAudioFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyAudioFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAudioFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAudioFileResponseBody) SetCode(v string) *ModifyAudioFileResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAudioFileResponseBody) SetData(v string) *ModifyAudioFileResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyAudioFileResponseBody) SetHttpStatusCode(v int32) *ModifyAudioFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAudioFileResponseBody) SetMessage(v string) *ModifyAudioFileResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAudioFileResponseBody) SetRequestId(v string) *ModifyAudioFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAudioFileResponseBody) Validate() error {
	return dara.Validate(s)
}

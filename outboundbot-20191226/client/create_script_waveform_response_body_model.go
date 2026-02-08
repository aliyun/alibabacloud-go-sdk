// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptWaveformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateScriptWaveformResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateScriptWaveformResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateScriptWaveformResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateScriptWaveformResponseBody
	GetRequestId() *string
	SetScriptWaveformId(v string) *CreateScriptWaveformResponseBody
	GetScriptWaveformId() *string
	SetSuccess(v bool) *CreateScriptWaveformResponseBody
	GetSuccess() *bool
}

type CreateScriptWaveformResponseBody struct {
	// Status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Script information
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Script recording ID
	//
	// example:
	//
	// df8216aa-d8f6-4501-864f-f8334d946561
	ScriptWaveformId *string `json:"ScriptWaveformId,omitempty" xml:"ScriptWaveformId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateScriptWaveformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptWaveformResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScriptWaveformResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateScriptWaveformResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateScriptWaveformResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScriptWaveformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScriptWaveformResponseBody) GetScriptWaveformId() *string {
	return s.ScriptWaveformId
}

func (s *CreateScriptWaveformResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScriptWaveformResponseBody) SetCode(v string) *CreateScriptWaveformResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScriptWaveformResponseBody) SetHttpStatusCode(v int32) *CreateScriptWaveformResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateScriptWaveformResponseBody) SetMessage(v string) *CreateScriptWaveformResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScriptWaveformResponseBody) SetRequestId(v string) *CreateScriptWaveformResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScriptWaveformResponseBody) SetScriptWaveformId(v string) *CreateScriptWaveformResponseBody {
	s.ScriptWaveformId = &v
	return s
}

func (s *CreateScriptWaveformResponseBody) SetSuccess(v bool) *CreateScriptWaveformResponseBody {
	s.Success = &v
	return s
}

func (s *CreateScriptWaveformResponseBody) Validate() error {
	return dara.Validate(s)
}

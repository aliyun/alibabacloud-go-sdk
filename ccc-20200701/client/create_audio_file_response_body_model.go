// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAudioFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAudioFileResponseBody
	GetCode() *string
	SetData(v string) *CreateAudioFileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateAudioFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAudioFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAudioFileResponseBody
	GetRequestId() *string
}

type CreateAudioFileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2301b83f-1f9f-491e-9f97-2f832ed92f0b
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 94E29B9B-DBC6-5951-B3DD-C85C1BDF20ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAudioFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAudioFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAudioFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAudioFileResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateAudioFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAudioFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAudioFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAudioFileResponseBody) SetCode(v string) *CreateAudioFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAudioFileResponseBody) SetData(v string) *CreateAudioFileResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAudioFileResponseBody) SetHttpStatusCode(v int32) *CreateAudioFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAudioFileResponseBody) SetMessage(v string) *CreateAudioFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAudioFileResponseBody) SetRequestId(v string) *CreateAudioFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAudioFileResponseBody) Validate() error {
	return dara.Validate(s)
}

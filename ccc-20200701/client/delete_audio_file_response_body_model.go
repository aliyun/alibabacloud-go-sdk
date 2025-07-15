// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAudioFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAudioFileResponseBody
	GetCode() *string
	SetData(v string) *DeleteAudioFileResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteAudioFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAudioFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAudioFileResponseBody
	GetRequestId() *string
}

type DeleteAudioFileResponseBody struct {
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
	// FB77821C-912F-57FF-8834-6336A6479093
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAudioFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAudioFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAudioFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAudioFileResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteAudioFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAudioFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAudioFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAudioFileResponseBody) SetCode(v string) *DeleteAudioFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAudioFileResponseBody) SetData(v string) *DeleteAudioFileResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAudioFileResponseBody) SetHttpStatusCode(v int32) *DeleteAudioFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAudioFileResponseBody) SetMessage(v string) *DeleteAudioFileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAudioFileResponseBody) SetRequestId(v string) *DeleteAudioFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAudioFileResponseBody) Validate() error {
	return dara.Validate(s)
}

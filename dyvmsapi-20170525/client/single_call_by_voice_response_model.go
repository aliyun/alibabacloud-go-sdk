// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleCallByVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SingleCallByVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SingleCallByVoiceResponse
	GetStatusCode() *int32
	SetBody(v *SingleCallByVoiceResponseBody) *SingleCallByVoiceResponse
	GetBody() *SingleCallByVoiceResponseBody
}

type SingleCallByVoiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleCallByVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SingleCallByVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s SingleCallByVoiceResponse) GoString() string {
	return s.String()
}

func (s *SingleCallByVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SingleCallByVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SingleCallByVoiceResponse) GetBody() *SingleCallByVoiceResponseBody {
	return s.Body
}

func (s *SingleCallByVoiceResponse) SetHeaders(v map[string]*string) *SingleCallByVoiceResponse {
	s.Headers = v
	return s
}

func (s *SingleCallByVoiceResponse) SetStatusCode(v int32) *SingleCallByVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SingleCallByVoiceResponse) SetBody(v *SingleCallByVoiceResponseBody) *SingleCallByVoiceResponse {
	s.Body = v
	return s
}

func (s *SingleCallByVoiceResponse) Validate() error {
	return dara.Validate(s)
}

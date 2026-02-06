// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptWaveformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScriptWaveformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScriptWaveformResponse
	GetStatusCode() *int32
	SetBody(v *CreateScriptWaveformResponseBody) *CreateScriptWaveformResponse
	GetBody() *CreateScriptWaveformResponseBody
}

type CreateScriptWaveformResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScriptWaveformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScriptWaveformResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptWaveformResponse) GoString() string {
	return s.String()
}

func (s *CreateScriptWaveformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScriptWaveformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScriptWaveformResponse) GetBody() *CreateScriptWaveformResponseBody {
	return s.Body
}

func (s *CreateScriptWaveformResponse) SetHeaders(v map[string]*string) *CreateScriptWaveformResponse {
	s.Headers = v
	return s
}

func (s *CreateScriptWaveformResponse) SetStatusCode(v int32) *CreateScriptWaveformResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScriptWaveformResponse) SetBody(v *CreateScriptWaveformResponseBody) *CreateScriptWaveformResponse {
	s.Body = v
	return s
}

func (s *CreateScriptWaveformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

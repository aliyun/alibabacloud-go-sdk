// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptWaveformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScriptWaveformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScriptWaveformResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScriptWaveformResponseBody) *DeleteScriptWaveformResponse
	GetBody() *DeleteScriptWaveformResponseBody
}

type DeleteScriptWaveformResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScriptWaveformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScriptWaveformResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptWaveformResponse) GoString() string {
	return s.String()
}

func (s *DeleteScriptWaveformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScriptWaveformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScriptWaveformResponse) GetBody() *DeleteScriptWaveformResponseBody {
	return s.Body
}

func (s *DeleteScriptWaveformResponse) SetHeaders(v map[string]*string) *DeleteScriptWaveformResponse {
	s.Headers = v
	return s
}

func (s *DeleteScriptWaveformResponse) SetStatusCode(v int32) *DeleteScriptWaveformResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScriptWaveformResponse) SetBody(v *DeleteScriptWaveformResponseBody) *DeleteScriptWaveformResponse {
	s.Body = v
	return s
}

func (s *DeleteScriptWaveformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScriptWaveformsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryScriptWaveformsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryScriptWaveformsResponse
	GetStatusCode() *int32
	SetBody(v *QueryScriptWaveformsResponseBody) *QueryScriptWaveformsResponse
	GetBody() *QueryScriptWaveformsResponseBody
}

type QueryScriptWaveformsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScriptWaveformsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScriptWaveformsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptWaveformsResponse) GoString() string {
	return s.String()
}

func (s *QueryScriptWaveformsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryScriptWaveformsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryScriptWaveformsResponse) GetBody() *QueryScriptWaveformsResponseBody {
	return s.Body
}

func (s *QueryScriptWaveformsResponse) SetHeaders(v map[string]*string) *QueryScriptWaveformsResponse {
	s.Headers = v
	return s
}

func (s *QueryScriptWaveformsResponse) SetStatusCode(v int32) *QueryScriptWaveformsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScriptWaveformsResponse) SetBody(v *QueryScriptWaveformsResponseBody) *QueryScriptWaveformsResponse {
	s.Body = v
	return s
}

func (s *QueryScriptWaveformsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

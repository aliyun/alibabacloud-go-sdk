// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartCasterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartCasterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartCasterResponse
	GetStatusCode() *int32
	SetBody(v *RestartCasterResponseBody) *RestartCasterResponse
	GetBody() *RestartCasterResponseBody
}

type RestartCasterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartCasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartCasterResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartCasterResponse) GoString() string {
	return s.String()
}

func (s *RestartCasterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartCasterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartCasterResponse) GetBody() *RestartCasterResponseBody {
	return s.Body
}

func (s *RestartCasterResponse) SetHeaders(v map[string]*string) *RestartCasterResponse {
	s.Headers = v
	return s
}

func (s *RestartCasterResponse) SetStatusCode(v int32) *RestartCasterResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartCasterResponse) SetBody(v *RestartCasterResponseBody) *RestartCasterResponse {
	s.Body = v
	return s
}

func (s *RestartCasterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

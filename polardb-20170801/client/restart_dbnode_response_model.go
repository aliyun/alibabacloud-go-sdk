// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartDBNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartDBNodeResponse
	GetStatusCode() *int32
	SetBody(v *RestartDBNodeResponseBody) *RestartDBNodeResponse
	GetBody() *RestartDBNodeResponseBody
}

type RestartDBNodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDBNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDBNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartDBNodeResponse) GoString() string {
	return s.String()
}

func (s *RestartDBNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartDBNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartDBNodeResponse) GetBody() *RestartDBNodeResponseBody {
	return s.Body
}

func (s *RestartDBNodeResponse) SetHeaders(v map[string]*string) *RestartDBNodeResponse {
	s.Headers = v
	return s
}

func (s *RestartDBNodeResponse) SetStatusCode(v int32) *RestartDBNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDBNodeResponse) SetBody(v *RestartDBNodeResponseBody) *RestartDBNodeResponse {
	s.Body = v
	return s
}

func (s *RestartDBNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

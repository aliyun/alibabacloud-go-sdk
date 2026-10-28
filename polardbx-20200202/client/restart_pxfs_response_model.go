// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartPxfsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartPxfsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartPxfsResponse
	GetStatusCode() *int32
	SetBody(v *RestartPxfsResponseBody) *RestartPxfsResponse
	GetBody() *RestartPxfsResponseBody
}

type RestartPxfsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartPxfsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartPxfsResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartPxfsResponse) GoString() string {
	return s.String()
}

func (s *RestartPxfsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartPxfsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartPxfsResponse) GetBody() *RestartPxfsResponseBody {
	return s.Body
}

func (s *RestartPxfsResponse) SetHeaders(v map[string]*string) *RestartPxfsResponse {
	s.Headers = v
	return s
}

func (s *RestartPxfsResponse) SetStatusCode(v int32) *RestartPxfsResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartPxfsResponse) SetBody(v *RestartPxfsResponseBody) *RestartPxfsResponse {
	s.Body = v
	return s
}

func (s *RestartPxfsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

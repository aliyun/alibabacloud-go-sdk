// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBNodeZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartDBNodeZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartDBNodeZonalResponse
	GetStatusCode() *int32
	SetBody(v *RestartDBNodeZonalResponseBody) *RestartDBNodeZonalResponse
	GetBody() *RestartDBNodeZonalResponseBody
}

type RestartDBNodeZonalResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDBNodeZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDBNodeZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartDBNodeZonalResponse) GoString() string {
	return s.String()
}

func (s *RestartDBNodeZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartDBNodeZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartDBNodeZonalResponse) GetBody() *RestartDBNodeZonalResponseBody {
	return s.Body
}

func (s *RestartDBNodeZonalResponse) SetHeaders(v map[string]*string) *RestartDBNodeZonalResponse {
	s.Headers = v
	return s
}

func (s *RestartDBNodeZonalResponse) SetStatusCode(v int32) *RestartDBNodeZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDBNodeZonalResponse) SetBody(v *RestartDBNodeZonalResponseBody) *RestartDBNodeZonalResponse {
	s.Body = v
	return s
}

func (s *RestartDBNodeZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

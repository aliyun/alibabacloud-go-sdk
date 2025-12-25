// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWindowConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWindowConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWindowConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetWindowConfigResponseBody) *GetWindowConfigResponse
	GetBody() *GetWindowConfigResponseBody
}

type GetWindowConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWindowConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWindowConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWindowConfigResponse) GoString() string {
	return s.String()
}

func (s *GetWindowConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWindowConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWindowConfigResponse) GetBody() *GetWindowConfigResponseBody {
	return s.Body
}

func (s *GetWindowConfigResponse) SetHeaders(v map[string]*string) *GetWindowConfigResponse {
	s.Headers = v
	return s
}

func (s *GetWindowConfigResponse) SetStatusCode(v int32) *GetWindowConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWindowConfigResponse) SetBody(v *GetWindowConfigResponseBody) *GetWindowConfigResponse {
	s.Body = v
	return s
}

func (s *GetWindowConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

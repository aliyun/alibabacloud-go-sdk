// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenConsoleSlsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenConsoleSlsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenConsoleSlsResponse
	GetStatusCode() *int32
	SetBody(v *OpenConsoleSlsResponseBody) *OpenConsoleSlsResponse
	GetBody() *OpenConsoleSlsResponseBody
}

type OpenConsoleSlsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenConsoleSlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenConsoleSlsResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenConsoleSlsResponse) GoString() string {
	return s.String()
}

func (s *OpenConsoleSlsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenConsoleSlsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenConsoleSlsResponse) GetBody() *OpenConsoleSlsResponseBody {
	return s.Body
}

func (s *OpenConsoleSlsResponse) SetHeaders(v map[string]*string) *OpenConsoleSlsResponse {
	s.Headers = v
	return s
}

func (s *OpenConsoleSlsResponse) SetStatusCode(v int32) *OpenConsoleSlsResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenConsoleSlsResponse) SetBody(v *OpenConsoleSlsResponseBody) *OpenConsoleSlsResponse {
	s.Body = v
	return s
}

func (s *OpenConsoleSlsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

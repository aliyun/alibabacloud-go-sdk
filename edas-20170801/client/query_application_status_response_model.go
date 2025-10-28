// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApplicationStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryApplicationStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryApplicationStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryApplicationStatusResponseBody) *QueryApplicationStatusResponse
	GetBody() *QueryApplicationStatusResponseBody
}

type QueryApplicationStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryApplicationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryApplicationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryApplicationStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryApplicationStatusResponse) GetBody() *QueryApplicationStatusResponseBody {
	return s.Body
}

func (s *QueryApplicationStatusResponse) SetHeaders(v map[string]*string) *QueryApplicationStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryApplicationStatusResponse) SetStatusCode(v int32) *QueryApplicationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryApplicationStatusResponse) SetBody(v *QueryApplicationStatusResponseBody) *QueryApplicationStatusResponse {
	s.Body = v
	return s
}

func (s *QueryApplicationStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

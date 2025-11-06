// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartViewResponse
	GetStatusCode() *int32
	SetBody(v *StartViewResponseBody) *StartViewResponse
	GetBody() *StartViewResponseBody
}

type StartViewResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartViewResponse) String() string {
	return dara.Prettify(s)
}

func (s StartViewResponse) GoString() string {
	return s.String()
}

func (s *StartViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartViewResponse) GetBody() *StartViewResponseBody {
	return s.Body
}

func (s *StartViewResponse) SetHeaders(v map[string]*string) *StartViewResponse {
	s.Headers = v
	return s
}

func (s *StartViewResponse) SetStatusCode(v int32) *StartViewResponse {
	s.StatusCode = &v
	return s
}

func (s *StartViewResponse) SetBody(v *StartViewResponseBody) *StartViewResponse {
	s.Body = v
	return s
}

func (s *StartViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

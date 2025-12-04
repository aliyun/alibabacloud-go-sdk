// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLoggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartLoggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartLoggingResponse
	GetStatusCode() *int32
	SetBody(v *StartLoggingResponseBody) *StartLoggingResponse
	GetBody() *StartLoggingResponseBody
}

type StartLoggingResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartLoggingResponse) String() string {
	return dara.Prettify(s)
}

func (s StartLoggingResponse) GoString() string {
	return s.String()
}

func (s *StartLoggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartLoggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartLoggingResponse) GetBody() *StartLoggingResponseBody {
	return s.Body
}

func (s *StartLoggingResponse) SetHeaders(v map[string]*string) *StartLoggingResponse {
	s.Headers = v
	return s
}

func (s *StartLoggingResponse) SetStatusCode(v int32) *StartLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLoggingResponse) SetBody(v *StartLoggingResponseBody) *StartLoggingResponse {
	s.Body = v
	return s
}

func (s *StartLoggingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

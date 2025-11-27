// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FailoverTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FailoverTestResponse
	GetStatusCode() *int32
	SetBody(v *FailoverTestResponseBody) *FailoverTestResponse
	GetBody() *FailoverTestResponseBody
}

type FailoverTestResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FailoverTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FailoverTestResponse) String() string {
	return dara.Prettify(s)
}

func (s FailoverTestResponse) GoString() string {
	return s.String()
}

func (s *FailoverTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FailoverTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FailoverTestResponse) GetBody() *FailoverTestResponseBody {
	return s.Body
}

func (s *FailoverTestResponse) SetHeaders(v map[string]*string) *FailoverTestResponse {
	s.Headers = v
	return s
}

func (s *FailoverTestResponse) SetStatusCode(v int32) *FailoverTestResponse {
	s.StatusCode = &v
	return s
}

func (s *FailoverTestResponse) SetBody(v *FailoverTestResponseBody) *FailoverTestResponse {
	s.Body = v
	return s
}

func (s *FailoverTestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

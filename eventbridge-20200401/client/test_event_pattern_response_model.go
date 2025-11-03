// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestEventPatternResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestEventPatternResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestEventPatternResponse
	GetStatusCode() *int32
	SetBody(v *TestEventPatternResponseBody) *TestEventPatternResponse
	GetBody() *TestEventPatternResponseBody
}

type TestEventPatternResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestEventPatternResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestEventPatternResponse) String() string {
	return dara.Prettify(s)
}

func (s TestEventPatternResponse) GoString() string {
	return s.String()
}

func (s *TestEventPatternResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestEventPatternResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestEventPatternResponse) GetBody() *TestEventPatternResponseBody {
	return s.Body
}

func (s *TestEventPatternResponse) SetHeaders(v map[string]*string) *TestEventPatternResponse {
	s.Headers = v
	return s
}

func (s *TestEventPatternResponse) SetStatusCode(v int32) *TestEventPatternResponse {
	s.StatusCode = &v
	return s
}

func (s *TestEventPatternResponse) SetBody(v *TestEventPatternResponseBody) *TestEventPatternResponse {
	s.Body = v
	return s
}

func (s *TestEventPatternResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

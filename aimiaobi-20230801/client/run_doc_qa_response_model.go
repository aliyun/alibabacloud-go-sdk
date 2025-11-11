// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocQaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDocQaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDocQaResponse
	GetStatusCode() *int32
	SetBody(v *RunDocQaResponseBody) *RunDocQaResponse
	GetBody() *RunDocQaResponseBody
}

type RunDocQaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunDocQaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDocQaResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponse) GoString() string {
	return s.String()
}

func (s *RunDocQaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDocQaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDocQaResponse) GetBody() *RunDocQaResponseBody {
	return s.Body
}

func (s *RunDocQaResponse) SetHeaders(v map[string]*string) *RunDocQaResponse {
	s.Headers = v
	return s
}

func (s *RunDocQaResponse) SetStatusCode(v int32) *RunDocQaResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDocQaResponse) SetBody(v *RunDocQaResponseBody) *RunDocQaResponse {
	s.Body = v
	return s
}

func (s *RunDocQaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

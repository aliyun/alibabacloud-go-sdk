// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunImageTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunImageTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunImageTestResponse
	GetStatusCode() *int32
	SetBody(v *RunImageTestResponseBody) *RunImageTestResponse
	GetBody() *RunImageTestResponseBody
}

type RunImageTestResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunImageTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunImageTestResponse) String() string {
	return dara.Prettify(s)
}

func (s RunImageTestResponse) GoString() string {
	return s.String()
}

func (s *RunImageTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunImageTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunImageTestResponse) GetBody() *RunImageTestResponseBody {
	return s.Body
}

func (s *RunImageTestResponse) SetHeaders(v map[string]*string) *RunImageTestResponse {
	s.Headers = v
	return s
}

func (s *RunImageTestResponse) SetStatusCode(v int32) *RunImageTestResponse {
	s.StatusCode = &v
	return s
}

func (s *RunImageTestResponse) SetBody(v *RunImageTestResponseBody) *RunImageTestResponse {
	s.Body = v
	return s
}

func (s *RunImageTestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

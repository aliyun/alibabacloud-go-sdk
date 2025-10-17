// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckResultResponse
	GetStatusCode() *int32
	SetBody(v *CheckResultResponseBody) *CheckResultResponse
	GetBody() *CheckResultResponseBody
}

type CheckResultResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckResultResponse) GoString() string {
	return s.String()
}

func (s *CheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckResultResponse) GetBody() *CheckResultResponseBody {
	return s.Body
}

func (s *CheckResultResponse) SetHeaders(v map[string]*string) *CheckResultResponse {
	s.Headers = v
	return s
}

func (s *CheckResultResponse) SetStatusCode(v int32) *CheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckResultResponse) SetBody(v *CheckResultResponseBody) *CheckResultResponse {
	s.Body = v
	return s
}

func (s *CheckResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

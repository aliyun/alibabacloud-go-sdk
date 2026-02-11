// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *CheckServiceStatusResponseBody) *CheckServiceStatusResponse
	GetBody() *CheckServiceStatusResponseBody
}

type CheckServiceStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckServiceStatusResponse) GetBody() *CheckServiceStatusResponseBody {
	return s.Body
}

func (s *CheckServiceStatusResponse) SetHeaders(v map[string]*string) *CheckServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceStatusResponse) SetStatusCode(v int32) *CheckServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceStatusResponse) SetBody(v *CheckServiceStatusResponseBody) *CheckServiceStatusResponse {
	s.Body = v
	return s
}

func (s *CheckServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

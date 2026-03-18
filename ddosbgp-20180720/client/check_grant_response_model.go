// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckGrantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckGrantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckGrantResponse
	GetStatusCode() *int32
	SetBody(v *CheckGrantResponseBody) *CheckGrantResponse
	GetBody() *CheckGrantResponseBody
}

type CheckGrantResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckGrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckGrantResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckGrantResponse) GoString() string {
	return s.String()
}

func (s *CheckGrantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckGrantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckGrantResponse) GetBody() *CheckGrantResponseBody {
	return s.Body
}

func (s *CheckGrantResponse) SetHeaders(v map[string]*string) *CheckGrantResponse {
	s.Headers = v
	return s
}

func (s *CheckGrantResponse) SetStatusCode(v int32) *CheckGrantResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckGrantResponse) SetBody(v *CheckGrantResponseBody) *CheckGrantResponse {
	s.Body = v
	return s
}

func (s *CheckGrantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

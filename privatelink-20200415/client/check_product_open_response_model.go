// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckProductOpenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckProductOpenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckProductOpenResponse
	GetStatusCode() *int32
	SetBody(v *CheckProductOpenResponseBody) *CheckProductOpenResponse
	GetBody() *CheckProductOpenResponseBody
}

type CheckProductOpenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckProductOpenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckProductOpenResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckProductOpenResponse) GoString() string {
	return s.String()
}

func (s *CheckProductOpenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckProductOpenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckProductOpenResponse) GetBody() *CheckProductOpenResponseBody {
	return s.Body
}

func (s *CheckProductOpenResponse) SetHeaders(v map[string]*string) *CheckProductOpenResponse {
	s.Headers = v
	return s
}

func (s *CheckProductOpenResponse) SetStatusCode(v int32) *CheckProductOpenResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckProductOpenResponse) SetBody(v *CheckProductOpenResponseBody) *CheckProductOpenResponse {
	s.Body = v
	return s
}

func (s *CheckProductOpenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

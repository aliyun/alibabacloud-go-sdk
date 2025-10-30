// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCustVariableLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCustVariableLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCustVariableLimitResponse
	GetStatusCode() *int32
	SetBody(v *CheckCustVariableLimitResponseBody) *CheckCustVariableLimitResponse
	GetBody() *CheckCustVariableLimitResponseBody
}

type CheckCustVariableLimitResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCustVariableLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCustVariableLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCustVariableLimitResponse) GoString() string {
	return s.String()
}

func (s *CheckCustVariableLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCustVariableLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCustVariableLimitResponse) GetBody() *CheckCustVariableLimitResponseBody {
	return s.Body
}

func (s *CheckCustVariableLimitResponse) SetHeaders(v map[string]*string) *CheckCustVariableLimitResponse {
	s.Headers = v
	return s
}

func (s *CheckCustVariableLimitResponse) SetStatusCode(v int32) *CheckCustVariableLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCustVariableLimitResponse) SetBody(v *CheckCustVariableLimitResponseBody) *CheckCustVariableLimitResponse {
	s.Body = v
	return s
}

func (s *CheckCustVariableLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

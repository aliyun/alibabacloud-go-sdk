// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUsageVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUsageVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUsageVariableResponse
	GetStatusCode() *int32
	SetBody(v *CheckUsageVariableResponseBody) *CheckUsageVariableResponse
	GetBody() *CheckUsageVariableResponseBody
}

type CheckUsageVariableResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUsageVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUsageVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUsageVariableResponse) GoString() string {
	return s.String()
}

func (s *CheckUsageVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUsageVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUsageVariableResponse) GetBody() *CheckUsageVariableResponseBody {
	return s.Body
}

func (s *CheckUsageVariableResponse) SetHeaders(v map[string]*string) *CheckUsageVariableResponse {
	s.Headers = v
	return s
}

func (s *CheckUsageVariableResponse) SetStatusCode(v int32) *CheckUsageVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUsageVariableResponse) SetBody(v *CheckUsageVariableResponseBody) *CheckUsageVariableResponse {
	s.Body = v
	return s
}

func (s *CheckUsageVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

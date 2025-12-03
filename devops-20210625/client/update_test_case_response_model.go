// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTestCaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTestCaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTestCaseResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTestCaseResponseBody) *UpdateTestCaseResponse
	GetBody() *UpdateTestCaseResponseBody
}

type UpdateTestCaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTestCaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTestCaseResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponse) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTestCaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTestCaseResponse) GetBody() *UpdateTestCaseResponseBody {
	return s.Body
}

func (s *UpdateTestCaseResponse) SetHeaders(v map[string]*string) *UpdateTestCaseResponse {
	s.Headers = v
	return s
}

func (s *UpdateTestCaseResponse) SetStatusCode(v int32) *UpdateTestCaseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTestCaseResponse) SetBody(v *UpdateTestCaseResponseBody) *UpdateTestCaseResponse {
	s.Body = v
	return s
}

func (s *UpdateTestCaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

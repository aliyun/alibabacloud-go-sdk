// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTestCaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTestCaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTestCaseResponse
	GetStatusCode() *int32
	SetBody(v *CreateTestCaseResponseBody) *CreateTestCaseResponse
	GetBody() *CreateTestCaseResponseBody
}

type CreateTestCaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTestCaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTestCaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponse) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTestCaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTestCaseResponse) GetBody() *CreateTestCaseResponseBody {
	return s.Body
}

func (s *CreateTestCaseResponse) SetHeaders(v map[string]*string) *CreateTestCaseResponse {
	s.Headers = v
	return s
}

func (s *CreateTestCaseResponse) SetStatusCode(v int32) *CreateTestCaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTestCaseResponse) SetBody(v *CreateTestCaseResponseBody) *CreateTestCaseResponse {
	s.Body = v
	return s
}

func (s *CreateTestCaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

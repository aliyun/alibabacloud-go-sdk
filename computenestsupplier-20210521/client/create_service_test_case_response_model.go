// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTestCaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceTestCaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceTestCaseResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceTestCaseResponseBody) *CreateServiceTestCaseResponse
	GetBody() *CreateServiceTestCaseResponseBody
}

type CreateServiceTestCaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceTestCaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceTestCaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTestCaseResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceTestCaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceTestCaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceTestCaseResponse) GetBody() *CreateServiceTestCaseResponseBody {
	return s.Body
}

func (s *CreateServiceTestCaseResponse) SetHeaders(v map[string]*string) *CreateServiceTestCaseResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceTestCaseResponse) SetStatusCode(v int32) *CreateServiceTestCaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceTestCaseResponse) SetBody(v *CreateServiceTestCaseResponseBody) *CreateServiceTestCaseResponse {
	s.Body = v
	return s
}

func (s *CreateServiceTestCaseResponse) Validate() error {
	return dara.Validate(s)
}

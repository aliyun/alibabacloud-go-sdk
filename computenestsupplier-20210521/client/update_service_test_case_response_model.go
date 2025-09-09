// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceTestCaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceTestCaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceTestCaseResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceTestCaseResponseBody) *UpdateServiceTestCaseResponse
	GetBody() *UpdateServiceTestCaseResponseBody
}

type UpdateServiceTestCaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceTestCaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceTestCaseResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceTestCaseResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceTestCaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceTestCaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceTestCaseResponse) GetBody() *UpdateServiceTestCaseResponseBody {
	return s.Body
}

func (s *UpdateServiceTestCaseResponse) SetHeaders(v map[string]*string) *UpdateServiceTestCaseResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceTestCaseResponse) SetStatusCode(v int32) *UpdateServiceTestCaseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceTestCaseResponse) SetBody(v *UpdateServiceTestCaseResponseBody) *UpdateServiceTestCaseResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceTestCaseResponse) Validate() error {
	return dara.Validate(s)
}

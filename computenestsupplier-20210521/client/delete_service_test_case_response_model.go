// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceTestCaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceTestCaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceTestCaseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceTestCaseResponseBody) *DeleteServiceTestCaseResponse
	GetBody() *DeleteServiceTestCaseResponseBody
}

type DeleteServiceTestCaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceTestCaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceTestCaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceTestCaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceTestCaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceTestCaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceTestCaseResponse) GetBody() *DeleteServiceTestCaseResponseBody {
	return s.Body
}

func (s *DeleteServiceTestCaseResponse) SetHeaders(v map[string]*string) *DeleteServiceTestCaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceTestCaseResponse) SetStatusCode(v int32) *DeleteServiceTestCaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceTestCaseResponse) SetBody(v *DeleteServiceTestCaseResponseBody) *DeleteServiceTestCaseResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceTestCaseResponse) Validate() error {
	return dara.Validate(s)
}

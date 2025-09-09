// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTestCasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceTestCasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceTestCasesResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceTestCasesResponseBody) *ListServiceTestCasesResponse
	GetBody() *ListServiceTestCasesResponseBody
}

type ListServiceTestCasesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceTestCasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceTestCasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestCasesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceTestCasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceTestCasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceTestCasesResponse) GetBody() *ListServiceTestCasesResponseBody {
	return s.Body
}

func (s *ListServiceTestCasesResponse) SetHeaders(v map[string]*string) *ListServiceTestCasesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceTestCasesResponse) SetStatusCode(v int32) *ListServiceTestCasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceTestCasesResponse) SetBody(v *ListServiceTestCasesResponseBody) *ListServiceTestCasesResponse {
	s.Body = v
	return s
}

func (s *ListServiceTestCasesResponse) Validate() error {
	return dara.Validate(s)
}

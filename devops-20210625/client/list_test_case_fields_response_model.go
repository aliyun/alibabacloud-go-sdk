// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTestCaseFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTestCaseFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTestCaseFieldsResponse
	GetStatusCode() *int32
	SetBody(v *ListTestCaseFieldsResponseBody) *ListTestCaseFieldsResponse
	GetBody() *ListTestCaseFieldsResponseBody
}

type ListTestCaseFieldsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTestCaseFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTestCaseFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTestCaseFieldsResponse) GoString() string {
	return s.String()
}

func (s *ListTestCaseFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTestCaseFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTestCaseFieldsResponse) GetBody() *ListTestCaseFieldsResponseBody {
	return s.Body
}

func (s *ListTestCaseFieldsResponse) SetHeaders(v map[string]*string) *ListTestCaseFieldsResponse {
	s.Headers = v
	return s
}

func (s *ListTestCaseFieldsResponse) SetStatusCode(v int32) *ListTestCaseFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTestCaseFieldsResponse) SetBody(v *ListTestCaseFieldsResponseBody) *ListTestCaseFieldsResponse {
	s.Body = v
	return s
}

func (s *ListTestCaseFieldsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

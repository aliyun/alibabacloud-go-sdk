// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEmployeeFieldValuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchEmployeeFieldValuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchEmployeeFieldValuesResponse
	GetStatusCode() *int32
	SetBody(v *SearchEmployeeFieldValuesResponseBody) *SearchEmployeeFieldValuesResponse
	GetBody() *SearchEmployeeFieldValuesResponseBody
}

type SearchEmployeeFieldValuesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchEmployeeFieldValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchEmployeeFieldValuesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchEmployeeFieldValuesResponse) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchEmployeeFieldValuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchEmployeeFieldValuesResponse) GetBody() *SearchEmployeeFieldValuesResponseBody {
	return s.Body
}

func (s *SearchEmployeeFieldValuesResponse) SetHeaders(v map[string]*string) *SearchEmployeeFieldValuesResponse {
	s.Headers = v
	return s
}

func (s *SearchEmployeeFieldValuesResponse) SetStatusCode(v int32) *SearchEmployeeFieldValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchEmployeeFieldValuesResponse) SetBody(v *SearchEmployeeFieldValuesResponseBody) *SearchEmployeeFieldValuesResponse {
	s.Body = v
	return s
}

func (s *SearchEmployeeFieldValuesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

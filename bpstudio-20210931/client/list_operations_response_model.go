// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationsResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationsResponseBody) *ListOperationsResponse
	GetBody() *ListOperationsResponseBody
}

type ListOperationsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationsResponse) GetBody() *ListOperationsResponseBody {
	return s.Body
}

func (s *ListOperationsResponse) SetHeaders(v map[string]*string) *ListOperationsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationsResponse) SetStatusCode(v int32) *ListOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationsResponse) SetBody(v *ListOperationsResponseBody) *ListOperationsResponse {
	s.Body = v
	return s
}

func (s *ListOperationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

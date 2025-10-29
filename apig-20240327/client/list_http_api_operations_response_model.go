// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpApiOperationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHttpApiOperationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHttpApiOperationsResponse
	GetStatusCode() *int32
	SetBody(v *ListHttpApiOperationsResponseBody) *ListHttpApiOperationsResponse
	GetBody() *ListHttpApiOperationsResponseBody
}

type ListHttpApiOperationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHttpApiOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHttpApiOperationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApiOperationsResponse) GoString() string {
	return s.String()
}

func (s *ListHttpApiOperationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHttpApiOperationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHttpApiOperationsResponse) GetBody() *ListHttpApiOperationsResponseBody {
	return s.Body
}

func (s *ListHttpApiOperationsResponse) SetHeaders(v map[string]*string) *ListHttpApiOperationsResponse {
	s.Headers = v
	return s
}

func (s *ListHttpApiOperationsResponse) SetStatusCode(v int32) *ListHttpApiOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHttpApiOperationsResponse) SetBody(v *ListHttpApiOperationsResponseBody) *ListHttpApiOperationsResponse {
	s.Body = v
	return s
}

func (s *ListHttpApiOperationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

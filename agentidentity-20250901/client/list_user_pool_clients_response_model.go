// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPoolClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserPoolClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserPoolClientsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserPoolClientsResponseBody) *ListUserPoolClientsResponse
	GetBody() *ListUserPoolClientsResponseBody
}

type ListUserPoolClientsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserPoolClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserPoolClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolClientsResponse) GoString() string {
	return s.String()
}

func (s *ListUserPoolClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserPoolClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserPoolClientsResponse) GetBody() *ListUserPoolClientsResponseBody {
	return s.Body
}

func (s *ListUserPoolClientsResponse) SetHeaders(v map[string]*string) *ListUserPoolClientsResponse {
	s.Headers = v
	return s
}

func (s *ListUserPoolClientsResponse) SetStatusCode(v int32) *ListUserPoolClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserPoolClientsResponse) SetBody(v *ListUserPoolClientsResponseBody) *ListUserPoolClientsResponse {
	s.Body = v
	return s
}

func (s *ListUserPoolClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

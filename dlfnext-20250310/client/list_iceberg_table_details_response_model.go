// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIcebergTableDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIcebergTableDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIcebergTableDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListIcebergTableDetailsResponseBody) *ListIcebergTableDetailsResponse
	GetBody() *ListIcebergTableDetailsResponseBody
}

type ListIcebergTableDetailsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIcebergTableDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIcebergTableDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIcebergTableDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListIcebergTableDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIcebergTableDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIcebergTableDetailsResponse) GetBody() *ListIcebergTableDetailsResponseBody {
	return s.Body
}

func (s *ListIcebergTableDetailsResponse) SetHeaders(v map[string]*string) *ListIcebergTableDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListIcebergTableDetailsResponse) SetStatusCode(v int32) *ListIcebergTableDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIcebergTableDetailsResponse) SetBody(v *ListIcebergTableDetailsResponseBody) *ListIcebergTableDetailsResponse {
	s.Body = v
	return s
}

func (s *ListIcebergTableDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

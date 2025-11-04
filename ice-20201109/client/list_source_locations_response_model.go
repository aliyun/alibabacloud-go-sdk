// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourceLocationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSourceLocationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSourceLocationsResponse
	GetStatusCode() *int32
	SetBody(v *ListSourceLocationsResponseBody) *ListSourceLocationsResponse
	GetBody() *ListSourceLocationsResponseBody
}

type ListSourceLocationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSourceLocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSourceLocationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSourceLocationsResponse) GoString() string {
	return s.String()
}

func (s *ListSourceLocationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSourceLocationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSourceLocationsResponse) GetBody() *ListSourceLocationsResponseBody {
	return s.Body
}

func (s *ListSourceLocationsResponse) SetHeaders(v map[string]*string) *ListSourceLocationsResponse {
	s.Headers = v
	return s
}

func (s *ListSourceLocationsResponse) SetStatusCode(v int32) *ListSourceLocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSourceLocationsResponse) SetBody(v *ListSourceLocationsResponseBody) *ListSourceLocationsResponse {
	s.Body = v
	return s
}

func (s *ListSourceLocationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

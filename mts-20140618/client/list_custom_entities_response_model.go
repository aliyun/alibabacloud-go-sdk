// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomEntitiesResponseBody) *ListCustomEntitiesResponse
	GetBody() *ListCustomEntitiesResponseBody
}

type ListCustomEntitiesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomEntitiesResponse) GetBody() *ListCustomEntitiesResponseBody {
	return s.Body
}

func (s *ListCustomEntitiesResponse) SetHeaders(v map[string]*string) *ListCustomEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomEntitiesResponse) SetStatusCode(v int32) *ListCustomEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomEntitiesResponse) SetBody(v *ListCustomEntitiesResponseBody) *ListCustomEntitiesResponse {
	s.Body = v
	return s
}

func (s *ListCustomEntitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

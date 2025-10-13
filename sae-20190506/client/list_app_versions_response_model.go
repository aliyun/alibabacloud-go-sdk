// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppVersionsResponseBody) *ListAppVersionsResponse
	GetBody() *ListAppVersionsResponseBody
}

type ListAppVersionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListAppVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppVersionsResponse) GetBody() *ListAppVersionsResponseBody {
	return s.Body
}

func (s *ListAppVersionsResponse) SetHeaders(v map[string]*string) *ListAppVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListAppVersionsResponse) SetStatusCode(v int32) *ListAppVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppVersionsResponse) SetBody(v *ListAppVersionsResponseBody) *ListAppVersionsResponse {
	s.Body = v
	return s
}

func (s *ListAppVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

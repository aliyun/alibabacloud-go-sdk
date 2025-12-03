// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeRequestsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChangeRequestsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChangeRequestsResponse
	GetStatusCode() *int32
	SetBody(v *ListChangeRequestsResponseBody) *ListChangeRequestsResponse
	GetBody() *ListChangeRequestsResponseBody
}

type ListChangeRequestsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChangeRequestsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChangeRequestsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChangeRequestsResponse) GoString() string {
	return s.String()
}

func (s *ListChangeRequestsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChangeRequestsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChangeRequestsResponse) GetBody() *ListChangeRequestsResponseBody {
	return s.Body
}

func (s *ListChangeRequestsResponse) SetHeaders(v map[string]*string) *ListChangeRequestsResponse {
	s.Headers = v
	return s
}

func (s *ListChangeRequestsResponse) SetStatusCode(v int32) *ListChangeRequestsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChangeRequestsResponse) SetBody(v *ListChangeRequestsResponseBody) *ListChangeRequestsResponse {
	s.Body = v
	return s
}

func (s *ListChangeRequestsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

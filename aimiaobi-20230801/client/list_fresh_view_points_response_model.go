// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFreshViewPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFreshViewPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFreshViewPointsResponse
	GetStatusCode() *int32
	SetBody(v *ListFreshViewPointsResponseBody) *ListFreshViewPointsResponse
	GetBody() *ListFreshViewPointsResponseBody
}

type ListFreshViewPointsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFreshViewPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFreshViewPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFreshViewPointsResponse) GoString() string {
	return s.String()
}

func (s *ListFreshViewPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFreshViewPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFreshViewPointsResponse) GetBody() *ListFreshViewPointsResponseBody {
	return s.Body
}

func (s *ListFreshViewPointsResponse) SetHeaders(v map[string]*string) *ListFreshViewPointsResponse {
	s.Headers = v
	return s
}

func (s *ListFreshViewPointsResponse) SetStatusCode(v int32) *ListFreshViewPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFreshViewPointsResponse) SetBody(v *ListFreshViewPointsResponseBody) *ListFreshViewPointsResponse {
	s.Body = v
	return s
}

func (s *ListFreshViewPointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusinessAccessPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBusinessAccessPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBusinessAccessPointsResponse
	GetStatusCode() *int32
	SetBody(v *ListBusinessAccessPointsResponseBody) *ListBusinessAccessPointsResponse
	GetBody() *ListBusinessAccessPointsResponseBody
}

type ListBusinessAccessPointsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBusinessAccessPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBusinessAccessPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessAccessPointsResponse) GoString() string {
	return s.String()
}

func (s *ListBusinessAccessPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBusinessAccessPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBusinessAccessPointsResponse) GetBody() *ListBusinessAccessPointsResponseBody {
	return s.Body
}

func (s *ListBusinessAccessPointsResponse) SetHeaders(v map[string]*string) *ListBusinessAccessPointsResponse {
	s.Headers = v
	return s
}

func (s *ListBusinessAccessPointsResponse) SetStatusCode(v int32) *ListBusinessAccessPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBusinessAccessPointsResponse) SetBody(v *ListBusinessAccessPointsResponseBody) *ListBusinessAccessPointsResponse {
	s.Body = v
	return s
}

func (s *ListBusinessAccessPointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

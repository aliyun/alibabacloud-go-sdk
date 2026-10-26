// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectAuthorizationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProjectAuthorizationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProjectAuthorizationsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProjectAuthorizationsResponseBody) *DeleteProjectAuthorizationsResponse
	GetBody() *DeleteProjectAuthorizationsResponseBody
}

type DeleteProjectAuthorizationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectAuthorizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectAuthorizationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectAuthorizationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectAuthorizationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProjectAuthorizationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProjectAuthorizationsResponse) GetBody() *DeleteProjectAuthorizationsResponseBody {
	return s.Body
}

func (s *DeleteProjectAuthorizationsResponse) SetHeaders(v map[string]*string) *DeleteProjectAuthorizationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectAuthorizationsResponse) SetStatusCode(v int32) *DeleteProjectAuthorizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectAuthorizationsResponse) SetBody(v *DeleteProjectAuthorizationsResponseBody) *DeleteProjectAuthorizationsResponse {
	s.Body = v
	return s
}

func (s *DeleteProjectAuthorizationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

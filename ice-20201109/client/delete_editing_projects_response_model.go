// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEditingProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEditingProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEditingProjectsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEditingProjectsResponseBody) *DeleteEditingProjectsResponse
	GetBody() *DeleteEditingProjectsResponseBody
}

type DeleteEditingProjectsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEditingProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEditingProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEditingProjectsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEditingProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEditingProjectsResponse) GetBody() *DeleteEditingProjectsResponseBody {
	return s.Body
}

func (s *DeleteEditingProjectsResponse) SetHeaders(v map[string]*string) *DeleteEditingProjectsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEditingProjectsResponse) SetStatusCode(v int32) *DeleteEditingProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEditingProjectsResponse) SetBody(v *DeleteEditingProjectsResponseBody) *DeleteEditingProjectsResponse {
	s.Body = v
	return s
}

func (s *DeleteEditingProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

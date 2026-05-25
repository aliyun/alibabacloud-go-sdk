// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTlogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchTlogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchTlogResponse
	GetStatusCode() *int32
	SetBody(v *SearchTlogResponseBody) *SearchTlogResponse
	GetBody() *SearchTlogResponseBody
}

type SearchTlogResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTlogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTlogResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchTlogResponse) GoString() string {
	return s.String()
}

func (s *SearchTlogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchTlogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchTlogResponse) GetBody() *SearchTlogResponseBody {
	return s.Body
}

func (s *SearchTlogResponse) SetHeaders(v map[string]*string) *SearchTlogResponse {
	s.Headers = v
	return s
}

func (s *SearchTlogResponse) SetStatusCode(v int32) *SearchTlogResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTlogResponse) SetBody(v *SearchTlogResponseBody) *SearchTlogResponse {
	s.Body = v
	return s
}

func (s *SearchTlogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

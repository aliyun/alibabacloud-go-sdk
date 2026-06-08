// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceServerScopesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceServerScopesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceServerScopesResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceServerScopesResponseBody) *ListResourceServerScopesResponse
	GetBody() *ListResourceServerScopesResponseBody
}

type ListResourceServerScopesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceServerScopesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceServerScopesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServerScopesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceServerScopesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceServerScopesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceServerScopesResponse) GetBody() *ListResourceServerScopesResponseBody {
	return s.Body
}

func (s *ListResourceServerScopesResponse) SetHeaders(v map[string]*string) *ListResourceServerScopesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceServerScopesResponse) SetStatusCode(v int32) *ListResourceServerScopesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceServerScopesResponse) SetBody(v *ListResourceServerScopesResponseBody) *ListResourceServerScopesResponse {
	s.Body = v
	return s
}

func (s *ListResourceServerScopesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

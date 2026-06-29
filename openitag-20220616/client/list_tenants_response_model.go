// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTenantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTenantsResponse
	GetStatusCode() *int32
	SetBody(v *ListTenantsResponseBody) *ListTenantsResponse
	GetBody() *ListTenantsResponseBody
}

type ListTenantsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTenantsResponse) GoString() string {
	return s.String()
}

func (s *ListTenantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTenantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTenantsResponse) GetBody() *ListTenantsResponseBody {
	return s.Body
}

func (s *ListTenantsResponse) SetHeaders(v map[string]*string) *ListTenantsResponse {
	s.Headers = v
	return s
}

func (s *ListTenantsResponse) SetStatusCode(v int32) *ListTenantsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantsResponse) SetBody(v *ListTenantsResponseBody) *ListTenantsResponse {
	s.Body = v
	return s
}

func (s *ListTenantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

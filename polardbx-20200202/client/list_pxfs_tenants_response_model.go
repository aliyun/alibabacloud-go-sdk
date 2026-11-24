// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPxfsTenantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPxfsTenantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPxfsTenantsResponse
	GetStatusCode() *int32
	SetBody(v *ListPxfsTenantsResponseBody) *ListPxfsTenantsResponse
	GetBody() *ListPxfsTenantsResponseBody
}

type ListPxfsTenantsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPxfsTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPxfsTenantsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPxfsTenantsResponse) GoString() string {
	return s.String()
}

func (s *ListPxfsTenantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPxfsTenantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPxfsTenantsResponse) GetBody() *ListPxfsTenantsResponseBody {
	return s.Body
}

func (s *ListPxfsTenantsResponse) SetHeaders(v map[string]*string) *ListPxfsTenantsResponse {
	s.Headers = v
	return s
}

func (s *ListPxfsTenantsResponse) SetStatusCode(v int32) *ListPxfsTenantsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPxfsTenantsResponse) SetBody(v *ListPxfsTenantsResponseBody) *ListPxfsTenantsResponse {
	s.Body = v
	return s
}

func (s *ListPxfsTenantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

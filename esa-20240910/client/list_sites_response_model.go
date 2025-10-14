// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSitesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSitesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSitesResponse
	GetStatusCode() *int32
	SetBody(v *ListSitesResponseBody) *ListSitesResponse
	GetBody() *ListSitesResponseBody
}

type ListSitesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSitesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSitesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSitesResponse) GoString() string {
	return s.String()
}

func (s *ListSitesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSitesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSitesResponse) GetBody() *ListSitesResponseBody {
	return s.Body
}

func (s *ListSitesResponse) SetHeaders(v map[string]*string) *ListSitesResponse {
	s.Headers = v
	return s
}

func (s *ListSitesResponse) SetStatusCode(v int32) *ListSitesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSitesResponse) SetBody(v *ListSitesResponseBody) *ListSitesResponse {
	s.Body = v
	return s
}

func (s *ListSitesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

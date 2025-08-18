// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSiteRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSiteRoutesResponse
	GetStatusCode() *int32
	SetBody(v *ListSiteRoutesResponseBody) *ListSiteRoutesResponse
	GetBody() *ListSiteRoutesResponseBody
}

type ListSiteRoutesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSiteRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSiteRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSiteRoutesResponse) GoString() string {
	return s.String()
}

func (s *ListSiteRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSiteRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSiteRoutesResponse) GetBody() *ListSiteRoutesResponseBody {
	return s.Body
}

func (s *ListSiteRoutesResponse) SetHeaders(v map[string]*string) *ListSiteRoutesResponse {
	s.Headers = v
	return s
}

func (s *ListSiteRoutesResponse) SetStatusCode(v int32) *ListSiteRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSiteRoutesResponse) SetBody(v *ListSiteRoutesResponseBody) *ListSiteRoutesResponse {
	s.Body = v
	return s
}

func (s *ListSiteRoutesResponse) Validate() error {
	return dara.Validate(s)
}

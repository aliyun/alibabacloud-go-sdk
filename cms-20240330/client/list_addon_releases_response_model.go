// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonReleasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAddonReleasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAddonReleasesResponse
	GetStatusCode() *int32
	SetBody(v *ListAddonReleasesResponseBody) *ListAddonReleasesResponse
	GetBody() *ListAddonReleasesResponseBody
}

type ListAddonReleasesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddonReleasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddonReleasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAddonReleasesResponse) GoString() string {
	return s.String()
}

func (s *ListAddonReleasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAddonReleasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAddonReleasesResponse) GetBody() *ListAddonReleasesResponseBody {
	return s.Body
}

func (s *ListAddonReleasesResponse) SetHeaders(v map[string]*string) *ListAddonReleasesResponse {
	s.Headers = v
	return s
}

func (s *ListAddonReleasesResponse) SetStatusCode(v int32) *ListAddonReleasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddonReleasesResponse) SetBody(v *ListAddonReleasesResponseBody) *ListAddonReleasesResponse {
	s.Body = v
	return s
}

func (s *ListAddonReleasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

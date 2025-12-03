// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAppsAuthoritiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveAppsAuthoritiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveAppsAuthoritiesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveAppsAuthoritiesResponseBody) *RemoveAppsAuthoritiesResponse
	GetBody() *RemoveAppsAuthoritiesResponseBody
}

type RemoveAppsAuthoritiesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAppsAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAppsAuthoritiesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveAppsAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *RemoveAppsAuthoritiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveAppsAuthoritiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveAppsAuthoritiesResponse) GetBody() *RemoveAppsAuthoritiesResponseBody {
	return s.Body
}

func (s *RemoveAppsAuthoritiesResponse) SetHeaders(v map[string]*string) *RemoveAppsAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *RemoveAppsAuthoritiesResponse) SetStatusCode(v int32) *RemoveAppsAuthoritiesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAppsAuthoritiesResponse) SetBody(v *RemoveAppsAuthoritiesResponseBody) *RemoveAppsAuthoritiesResponse {
	s.Body = v
	return s
}

func (s *RemoveAppsAuthoritiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApisAuthoritiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveApisAuthoritiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveApisAuthoritiesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveApisAuthoritiesResponseBody) *RemoveApisAuthoritiesResponse
	GetBody() *RemoveApisAuthoritiesResponseBody
}

type RemoveApisAuthoritiesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveApisAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveApisAuthoritiesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveApisAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *RemoveApisAuthoritiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveApisAuthoritiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveApisAuthoritiesResponse) GetBody() *RemoveApisAuthoritiesResponseBody {
	return s.Body
}

func (s *RemoveApisAuthoritiesResponse) SetHeaders(v map[string]*string) *RemoveApisAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *RemoveApisAuthoritiesResponse) SetStatusCode(v int32) *RemoveApisAuthoritiesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveApisAuthoritiesResponse) SetBody(v *RemoveApisAuthoritiesResponseBody) *RemoveApisAuthoritiesResponse {
	s.Body = v
	return s
}

func (s *RemoveApisAuthoritiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReleaseVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReleaseVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReleaseVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListReleaseVersionsResponseBody) *ListReleaseVersionsResponse
	GetBody() *ListReleaseVersionsResponseBody
}

type ListReleaseVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReleaseVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReleaseVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReleaseVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListReleaseVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReleaseVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReleaseVersionsResponse) GetBody() *ListReleaseVersionsResponseBody {
	return s.Body
}

func (s *ListReleaseVersionsResponse) SetHeaders(v map[string]*string) *ListReleaseVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListReleaseVersionsResponse) SetStatusCode(v int32) *ListReleaseVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReleaseVersionsResponse) SetBody(v *ListReleaseVersionsResponseBody) *ListReleaseVersionsResponse {
	s.Body = v
	return s
}

func (s *ListReleaseVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppsResponseBody) *ListAppsResponse
	GetBody() *ListAppsResponseBody
}

type ListAppsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppsResponse) GetBody() *ListAppsResponseBody {
	return s.Body
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetStatusCode(v int32) *ListAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

func (s *ListAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

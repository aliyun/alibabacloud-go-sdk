// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublishedAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublishedAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListPublishedAppsResponseBody) *ListPublishedAppsResponse
	GetBody() *ListPublishedAppsResponseBody
}

type ListPublishedAppsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishedAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishedAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppsResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublishedAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublishedAppsResponse) GetBody() *ListPublishedAppsResponseBody {
	return s.Body
}

func (s *ListPublishedAppsResponse) SetHeaders(v map[string]*string) *ListPublishedAppsResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedAppsResponse) SetStatusCode(v int32) *ListPublishedAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedAppsResponse) SetBody(v *ListPublishedAppsResponseBody) *ListPublishedAppsResponse {
	s.Body = v
	return s
}

func (s *ListPublishedAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

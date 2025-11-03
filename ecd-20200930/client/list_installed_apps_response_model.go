// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstalledAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstalledAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstalledAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstalledAppsResponseBody) *ListInstalledAppsResponse
	GetBody() *ListInstalledAppsResponseBody
}

type ListInstalledAppsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstalledAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstalledAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstalledAppsResponse) GoString() string {
	return s.String()
}

func (s *ListInstalledAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstalledAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstalledAppsResponse) GetBody() *ListInstalledAppsResponseBody {
	return s.Body
}

func (s *ListInstalledAppsResponse) SetHeaders(v map[string]*string) *ListInstalledAppsResponse {
	s.Headers = v
	return s
}

func (s *ListInstalledAppsResponse) SetStatusCode(v int32) *ListInstalledAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstalledAppsResponse) SetBody(v *ListInstalledAppsResponseBody) *ListInstalledAppsResponse {
	s.Body = v
	return s
}

func (s *ListInstalledAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

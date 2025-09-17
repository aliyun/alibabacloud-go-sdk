// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRunningAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRunningAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRunningAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListRunningAppsResponseBody) *ListRunningAppsResponse
	GetBody() *ListRunningAppsResponseBody
}

type ListRunningAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRunningAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRunningAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRunningAppsResponse) GoString() string {
	return s.String()
}

func (s *ListRunningAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRunningAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRunningAppsResponse) GetBody() *ListRunningAppsResponseBody {
	return s.Body
}

func (s *ListRunningAppsResponse) SetHeaders(v map[string]*string) *ListRunningAppsResponse {
	s.Headers = v
	return s
}

func (s *ListRunningAppsResponse) SetStatusCode(v int32) *ListRunningAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRunningAppsResponse) SetBody(v *ListRunningAppsResponseBody) *ListRunningAppsResponse {
	s.Body = v
	return s
}

func (s *ListRunningAppsResponse) Validate() error {
	return dara.Validate(s)
}

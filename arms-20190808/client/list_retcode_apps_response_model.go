// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRetcodeAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRetcodeAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRetcodeAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListRetcodeAppsResponseBody) *ListRetcodeAppsResponse
	GetBody() *ListRetcodeAppsResponseBody
}

type ListRetcodeAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRetcodeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRetcodeAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRetcodeAppsResponse) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRetcodeAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRetcodeAppsResponse) GetBody() *ListRetcodeAppsResponseBody {
	return s.Body
}

func (s *ListRetcodeAppsResponse) SetHeaders(v map[string]*string) *ListRetcodeAppsResponse {
	s.Headers = v
	return s
}

func (s *ListRetcodeAppsResponse) SetStatusCode(v int32) *ListRetcodeAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRetcodeAppsResponse) SetBody(v *ListRetcodeAppsResponseBody) *ListRetcodeAppsResponse {
	s.Body = v
	return s
}

func (s *ListRetcodeAppsResponse) Validate() error {
	return dara.Validate(s)
}

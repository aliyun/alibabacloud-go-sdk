// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSparkAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSparkAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListSparkAppsResponseBody) *ListSparkAppsResponse
	GetBody() *ListSparkAppsResponseBody
}

type ListSparkAppsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSparkAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSparkAppsResponse) GoString() string {
	return s.String()
}

func (s *ListSparkAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSparkAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSparkAppsResponse) GetBody() *ListSparkAppsResponseBody {
	return s.Body
}

func (s *ListSparkAppsResponse) SetHeaders(v map[string]*string) *ListSparkAppsResponse {
	s.Headers = v
	return s
}

func (s *ListSparkAppsResponse) SetStatusCode(v int32) *ListSparkAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkAppsResponse) SetBody(v *ListSparkAppsResponseBody) *ListSparkAppsResponse {
	s.Body = v
	return s
}

func (s *ListSparkAppsResponse) Validate() error {
	return dara.Validate(s)
}

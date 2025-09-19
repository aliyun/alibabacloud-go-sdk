// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulGlobalConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVulGlobalConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVulGlobalConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListVulGlobalConfigResponseBody) *ListVulGlobalConfigResponse
	GetBody() *ListVulGlobalConfigResponseBody
}

type ListVulGlobalConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVulGlobalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVulGlobalConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVulGlobalConfigResponse) GoString() string {
	return s.String()
}

func (s *ListVulGlobalConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVulGlobalConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVulGlobalConfigResponse) GetBody() *ListVulGlobalConfigResponseBody {
	return s.Body
}

func (s *ListVulGlobalConfigResponse) SetHeaders(v map[string]*string) *ListVulGlobalConfigResponse {
	s.Headers = v
	return s
}

func (s *ListVulGlobalConfigResponse) SetStatusCode(v int32) *ListVulGlobalConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVulGlobalConfigResponse) SetBody(v *ListVulGlobalConfigResponseBody) *ListVulGlobalConfigResponse {
	s.Body = v
	return s
}

func (s *ListVulGlobalConfigResponse) Validate() error {
	return dara.Validate(s)
}

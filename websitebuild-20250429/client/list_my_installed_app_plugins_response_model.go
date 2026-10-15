// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyInstalledAppPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMyInstalledAppPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMyInstalledAppPluginsResponse
	GetStatusCode() *int32
	SetBody(v *ListMyInstalledAppPluginsResponseBody) *ListMyInstalledAppPluginsResponse
	GetBody() *ListMyInstalledAppPluginsResponseBody
}

type ListMyInstalledAppPluginsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyInstalledAppPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyInstalledAppPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMyInstalledAppPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListMyInstalledAppPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMyInstalledAppPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMyInstalledAppPluginsResponse) GetBody() *ListMyInstalledAppPluginsResponseBody {
	return s.Body
}

func (s *ListMyInstalledAppPluginsResponse) SetHeaders(v map[string]*string) *ListMyInstalledAppPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListMyInstalledAppPluginsResponse) SetStatusCode(v int32) *ListMyInstalledAppPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyInstalledAppPluginsResponse) SetBody(v *ListMyInstalledAppPluginsResponseBody) *ListMyInstalledAppPluginsResponse {
	s.Body = v
	return s
}

func (s *ListMyInstalledAppPluginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

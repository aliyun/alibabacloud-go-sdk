// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyAppPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMyAppPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMyAppPluginsResponse
	GetStatusCode() *int32
	SetBody(v *ListMyAppPluginsResponseBody) *ListMyAppPluginsResponse
	GetBody() *ListMyAppPluginsResponseBody
}

type ListMyAppPluginsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyAppPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyAppPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMyAppPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListMyAppPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMyAppPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMyAppPluginsResponse) GetBody() *ListMyAppPluginsResponseBody {
	return s.Body
}

func (s *ListMyAppPluginsResponse) SetHeaders(v map[string]*string) *ListMyAppPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListMyAppPluginsResponse) SetStatusCode(v int32) *ListMyAppPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyAppPluginsResponse) SetBody(v *ListMyAppPluginsResponseBody) *ListMyAppPluginsResponse {
	s.Body = v
	return s
}

func (s *ListMyAppPluginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

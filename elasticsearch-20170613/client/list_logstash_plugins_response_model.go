// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstashPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogstashPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogstashPluginsResponse
	GetStatusCode() *int32
	SetBody(v *ListLogstashPluginsResponseBody) *ListLogstashPluginsResponse
	GetBody() *ListLogstashPluginsResponseBody
}

type ListLogstashPluginsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogstashPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogstashPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListLogstashPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogstashPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogstashPluginsResponse) GetBody() *ListLogstashPluginsResponseBody {
	return s.Body
}

func (s *ListLogstashPluginsResponse) SetHeaders(v map[string]*string) *ListLogstashPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListLogstashPluginsResponse) SetStatusCode(v int32) *ListLogstashPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogstashPluginsResponse) SetBody(v *ListLogstashPluginsResponseBody) *ListLogstashPluginsResponse {
	s.Body = v
	return s
}

func (s *ListLogstashPluginsResponse) Validate() error {
	return dara.Validate(s)
}

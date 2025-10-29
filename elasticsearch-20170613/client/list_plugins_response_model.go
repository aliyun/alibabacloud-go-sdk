// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPluginsResponse
	GetStatusCode() *int32
	SetBody(v *ListPluginsResponseBody) *ListPluginsResponse
	GetBody() *ListPluginsResponseBody
}

type ListPluginsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPluginsResponse) GetBody() *ListPluginsResponseBody {
	return s.Body
}

func (s *ListPluginsResponse) SetHeaders(v map[string]*string) *ListPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListPluginsResponse) SetStatusCode(v int32) *ListPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPluginsResponse) SetBody(v *ListPluginsResponseBody) *ListPluginsResponse {
	s.Body = v
	return s
}

func (s *ListPluginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

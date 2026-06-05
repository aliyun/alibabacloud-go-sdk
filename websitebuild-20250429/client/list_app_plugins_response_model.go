// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppPluginsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppPluginsResponseBody) *ListAppPluginsResponse
	GetBody() *ListAppPluginsResponseBody
}

type ListAppPluginsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListAppPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppPluginsResponse) GetBody() *ListAppPluginsResponseBody {
	return s.Body
}

func (s *ListAppPluginsResponse) SetHeaders(v map[string]*string) *ListAppPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListAppPluginsResponse) SetStatusCode(v int32) *ListAppPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppPluginsResponse) SetBody(v *ListAppPluginsResponseBody) *ListAppPluginsResponse {
	s.Body = v
	return s
}

func (s *ListAppPluginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

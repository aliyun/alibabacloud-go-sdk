// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppPluginConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppPluginConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppPluginConfigsResponseBody) *ListAppPluginConfigsResponse
	GetBody() *ListAppPluginConfigsResponseBody
}

type ListAppPluginConfigsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppPluginConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppPluginConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAppPluginConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppPluginConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppPluginConfigsResponse) GetBody() *ListAppPluginConfigsResponseBody {
	return s.Body
}

func (s *ListAppPluginConfigsResponse) SetHeaders(v map[string]*string) *ListAppPluginConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAppPluginConfigsResponse) SetStatusCode(v int32) *ListAppPluginConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppPluginConfigsResponse) SetBody(v *ListAppPluginConfigsResponseBody) *ListAppPluginConfigsResponse {
	s.Body = v
	return s
}

func (s *ListAppPluginConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

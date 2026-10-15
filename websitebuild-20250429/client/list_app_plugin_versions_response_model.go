// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppPluginVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppPluginVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppPluginVersionsResponseBody) *ListAppPluginVersionsResponse
	GetBody() *ListAppPluginVersionsResponseBody
}

type ListAppPluginVersionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppPluginVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppPluginVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListAppPluginVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppPluginVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppPluginVersionsResponse) GetBody() *ListAppPluginVersionsResponseBody {
	return s.Body
}

func (s *ListAppPluginVersionsResponse) SetHeaders(v map[string]*string) *ListAppPluginVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListAppPluginVersionsResponse) SetStatusCode(v int32) *ListAppPluginVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppPluginVersionsResponse) SetBody(v *ListAppPluginVersionsResponseBody) *ListAppPluginVersionsResponse {
	s.Body = v
	return s
}

func (s *ListAppPluginVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

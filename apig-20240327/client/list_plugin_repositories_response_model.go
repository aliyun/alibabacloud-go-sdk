// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginRepositoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPluginRepositoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPluginRepositoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListPluginRepositoriesResponseBody) *ListPluginRepositoriesResponse
	GetBody() *ListPluginRepositoriesResponseBody
}

type ListPluginRepositoriesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPluginRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPluginRepositoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPluginRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListPluginRepositoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPluginRepositoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPluginRepositoriesResponse) GetBody() *ListPluginRepositoriesResponseBody {
	return s.Body
}

func (s *ListPluginRepositoriesResponse) SetHeaders(v map[string]*string) *ListPluginRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListPluginRepositoriesResponse) SetStatusCode(v int32) *ListPluginRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPluginRepositoriesResponse) SetBody(v *ListPluginRepositoriesResponseBody) *ListPluginRepositoriesResponse {
	s.Body = v
	return s
}

func (s *ListPluginRepositoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKibanaPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKibanaPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKibanaPluginsResponse
	GetStatusCode() *int32
	SetBody(v *ListKibanaPluginsResponseBody) *ListKibanaPluginsResponse
	GetBody() *ListKibanaPluginsResponseBody
}

type ListKibanaPluginsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKibanaPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKibanaPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKibanaPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListKibanaPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKibanaPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKibanaPluginsResponse) GetBody() *ListKibanaPluginsResponseBody {
	return s.Body
}

func (s *ListKibanaPluginsResponse) SetHeaders(v map[string]*string) *ListKibanaPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListKibanaPluginsResponse) SetStatusCode(v int32) *ListKibanaPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKibanaPluginsResponse) SetBody(v *ListKibanaPluginsResponseBody) *ListKibanaPluginsResponse {
	s.Body = v
	return s
}

func (s *ListKibanaPluginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

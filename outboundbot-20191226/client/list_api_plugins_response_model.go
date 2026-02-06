// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiPluginsResponse
	GetStatusCode() *int32
	SetBody(v *ListApiPluginsResponseBody) *ListApiPluginsResponse
	GetBody() *ListApiPluginsResponseBody
}

type ListApiPluginsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListApiPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiPluginsResponse) GetBody() *ListApiPluginsResponseBody {
	return s.Body
}

func (s *ListApiPluginsResponse) SetHeaders(v map[string]*string) *ListApiPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListApiPluginsResponse) SetStatusCode(v int32) *ListApiPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiPluginsResponse) SetBody(v *ListApiPluginsResponseBody) *ListApiPluginsResponse {
	s.Body = v
	return s
}

func (s *ListApiPluginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

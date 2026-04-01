// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserPluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserPluginResponse
	GetStatusCode() *int32
	SetBody(v *ListUserPluginResponseBody) *ListUserPluginResponse
	GetBody() *ListUserPluginResponseBody
}

type ListUserPluginResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserPluginResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserPluginResponse) GoString() string {
	return s.String()
}

func (s *ListUserPluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserPluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserPluginResponse) GetBody() *ListUserPluginResponseBody {
	return s.Body
}

func (s *ListUserPluginResponse) SetHeaders(v map[string]*string) *ListUserPluginResponse {
	s.Headers = v
	return s
}

func (s *ListUserPluginResponse) SetStatusCode(v int32) *ListUserPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserPluginResponse) SetBody(v *ListUserPluginResponseBody) *ListUserPluginResponse {
	s.Body = v
	return s
}

func (s *ListUserPluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

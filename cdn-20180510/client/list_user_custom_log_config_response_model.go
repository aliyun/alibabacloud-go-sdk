// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserCustomLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserCustomLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserCustomLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListUserCustomLogConfigResponseBody) *ListUserCustomLogConfigResponse
	GetBody() *ListUserCustomLogConfigResponseBody
}

type ListUserCustomLogConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserCustomLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserCustomLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserCustomLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ListUserCustomLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserCustomLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserCustomLogConfigResponse) GetBody() *ListUserCustomLogConfigResponseBody {
	return s.Body
}

func (s *ListUserCustomLogConfigResponse) SetHeaders(v map[string]*string) *ListUserCustomLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ListUserCustomLogConfigResponse) SetStatusCode(v int32) *ListUserCustomLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserCustomLogConfigResponse) SetBody(v *ListUserCustomLogConfigResponseBody) *ListUserCustomLogConfigResponse {
	s.Body = v
	return s
}

func (s *ListUserCustomLogConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

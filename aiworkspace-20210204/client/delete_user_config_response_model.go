// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserConfigResponseBody) *DeleteUserConfigResponse
	GetBody() *DeleteUserConfigResponseBody
}

type DeleteUserConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserConfigResponse) GetBody() *DeleteUserConfigResponseBody {
	return s.Body
}

func (s *DeleteUserConfigResponse) SetHeaders(v map[string]*string) *DeleteUserConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserConfigResponse) SetStatusCode(v int32) *DeleteUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserConfigResponse) SetBody(v *DeleteUserConfigResponseBody) *DeleteUserConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteUserConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

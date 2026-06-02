// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserAuthConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserAuthConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserAuthConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserAuthConfigResponseBody) *DeleteUserAuthConfigResponse
	GetBody() *DeleteUserAuthConfigResponseBody
}

type DeleteUserAuthConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserAuthConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserAuthConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserAuthConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserAuthConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserAuthConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserAuthConfigResponse) GetBody() *DeleteUserAuthConfigResponseBody {
	return s.Body
}

func (s *DeleteUserAuthConfigResponse) SetHeaders(v map[string]*string) *DeleteUserAuthConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserAuthConfigResponse) SetStatusCode(v int32) *DeleteUserAuthConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserAuthConfigResponse) SetBody(v *DeleteUserAuthConfigResponseBody) *DeleteUserAuthConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteUserAuthConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

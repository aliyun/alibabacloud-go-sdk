// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiMcpServerUserConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApiMcpServerUserConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApiMcpServerUserConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApiMcpServerUserConfigResponseBody) *UpdateApiMcpServerUserConfigResponse
	GetBody() *UpdateApiMcpServerUserConfigResponseBody
}

type UpdateApiMcpServerUserConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApiMcpServerUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApiMcpServerUserConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerUserConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerUserConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApiMcpServerUserConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApiMcpServerUserConfigResponse) GetBody() *UpdateApiMcpServerUserConfigResponseBody {
	return s.Body
}

func (s *UpdateApiMcpServerUserConfigResponse) SetHeaders(v map[string]*string) *UpdateApiMcpServerUserConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateApiMcpServerUserConfigResponse) SetStatusCode(v int32) *UpdateApiMcpServerUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApiMcpServerUserConfigResponse) SetBody(v *UpdateApiMcpServerUserConfigResponseBody) *UpdateApiMcpServerUserConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateApiMcpServerUserConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

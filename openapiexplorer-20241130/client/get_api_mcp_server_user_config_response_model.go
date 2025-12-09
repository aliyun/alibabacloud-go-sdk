// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiMcpServerUserConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiMcpServerUserConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiMcpServerUserConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetApiMcpServerUserConfigResponseBody) *GetApiMcpServerUserConfigResponse
	GetBody() *GetApiMcpServerUserConfigResponseBody
}

type GetApiMcpServerUserConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiMcpServerUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiMcpServerUserConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerUserConfigResponse) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerUserConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiMcpServerUserConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiMcpServerUserConfigResponse) GetBody() *GetApiMcpServerUserConfigResponseBody {
	return s.Body
}

func (s *GetApiMcpServerUserConfigResponse) SetHeaders(v map[string]*string) *GetApiMcpServerUserConfigResponse {
	s.Headers = v
	return s
}

func (s *GetApiMcpServerUserConfigResponse) SetStatusCode(v int32) *GetApiMcpServerUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiMcpServerUserConfigResponse) SetBody(v *GetApiMcpServerUserConfigResponseBody) *GetApiMcpServerUserConfigResponse {
	s.Body = v
	return s
}

func (s *GetApiMcpServerUserConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

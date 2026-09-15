// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpTemplateConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMcpTemplateConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMcpTemplateConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMcpTemplateConfigResponseBody) *UpdateMcpTemplateConfigResponse
	GetBody() *UpdateMcpTemplateConfigResponseBody
}

type UpdateMcpTemplateConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMcpTemplateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMcpTemplateConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcpTemplateConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMcpTemplateConfigResponse) GetBody() *UpdateMcpTemplateConfigResponseBody {
	return s.Body
}

func (s *UpdateMcpTemplateConfigResponse) SetHeaders(v map[string]*string) *UpdateMcpTemplateConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateMcpTemplateConfigResponse) SetStatusCode(v int32) *UpdateMcpTemplateConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMcpTemplateConfigResponse) SetBody(v *UpdateMcpTemplateConfigResponseBody) *UpdateMcpTemplateConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateMcpTemplateConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

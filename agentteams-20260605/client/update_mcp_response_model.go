// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMcpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMcpResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMcpResponseBody) *UpdateMcpResponse
	GetBody() *UpdateMcpResponseBody
}

type UpdateMcpResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMcpResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpResponse) GoString() string {
	return s.String()
}

func (s *UpdateMcpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMcpResponse) GetBody() *UpdateMcpResponseBody {
	return s.Body
}

func (s *UpdateMcpResponse) SetHeaders(v map[string]*string) *UpdateMcpResponse {
	s.Headers = v
	return s
}

func (s *UpdateMcpResponse) SetStatusCode(v int32) *UpdateMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMcpResponse) SetBody(v *UpdateMcpResponseBody) *UpdateMcpResponse {
	s.Body = v
	return s
}

func (s *UpdateMcpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

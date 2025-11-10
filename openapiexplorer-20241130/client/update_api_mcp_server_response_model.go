// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApiMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApiMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApiMcpServerResponseBody) *UpdateApiMcpServerResponse
	GetBody() *UpdateApiMcpServerResponseBody
}

type UpdateApiMcpServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApiMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApiMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerResponse) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApiMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApiMcpServerResponse) GetBody() *UpdateApiMcpServerResponseBody {
	return s.Body
}

func (s *UpdateApiMcpServerResponse) SetHeaders(v map[string]*string) *UpdateApiMcpServerResponse {
	s.Headers = v
	return s
}

func (s *UpdateApiMcpServerResponse) SetStatusCode(v int32) *UpdateApiMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApiMcpServerResponse) SetBody(v *UpdateApiMcpServerResponseBody) *UpdateApiMcpServerResponse {
	s.Body = v
	return s
}

func (s *UpdateApiMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

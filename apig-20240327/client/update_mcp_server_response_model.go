// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMcpServerResponseBody) *UpdateMcpServerResponse
	GetBody() *UpdateMcpServerResponseBody
}

type UpdateMcpServerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerResponse) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMcpServerResponse) GetBody() *UpdateMcpServerResponseBody {
	return s.Body
}

func (s *UpdateMcpServerResponse) SetHeaders(v map[string]*string) *UpdateMcpServerResponse {
	s.Headers = v
	return s
}

func (s *UpdateMcpServerResponse) SetStatusCode(v int32) *UpdateMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMcpServerResponse) SetBody(v *UpdateMcpServerResponseBody) *UpdateMcpServerResponse {
	s.Body = v
	return s
}

func (s *UpdateMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

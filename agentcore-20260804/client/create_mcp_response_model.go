// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcpResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcpResponseBody) *CreateMcpResponse
	GetBody() *CreateMcpResponseBody
}

type CreateMcpResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcpResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponse) GoString() string {
	return s.String()
}

func (s *CreateMcpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcpResponse) GetBody() *CreateMcpResponseBody {
	return s.Body
}

func (s *CreateMcpResponse) SetHeaders(v map[string]*string) *CreateMcpResponse {
	s.Headers = v
	return s
}

func (s *CreateMcpResponse) SetStatusCode(v int32) *CreateMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcpResponse) SetBody(v *CreateMcpResponseBody) *CreateMcpResponse {
	s.Body = v
	return s
}

func (s *CreateMcpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

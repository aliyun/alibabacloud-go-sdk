// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcpServerResponseBody) *CreateMcpServerResponse
	GetBody() *CreateMcpServerResponseBody
}

type CreateMcpServerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerResponse) GoString() string {
	return s.String()
}

func (s *CreateMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcpServerResponse) GetBody() *CreateMcpServerResponseBody {
	return s.Body
}

func (s *CreateMcpServerResponse) SetHeaders(v map[string]*string) *CreateMcpServerResponse {
	s.Headers = v
	return s
}

func (s *CreateMcpServerResponse) SetStatusCode(v int32) *CreateMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcpServerResponse) SetBody(v *CreateMcpServerResponseBody) *CreateMcpServerResponse {
	s.Body = v
	return s
}

func (s *CreateMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

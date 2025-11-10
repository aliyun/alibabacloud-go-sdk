// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApiMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApiMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *CreateApiMcpServerResponseBody) *CreateApiMcpServerResponse
	GetBody() *CreateApiMcpServerResponseBody
}

type CreateApiMcpServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApiMcpServerResponse) GoString() string {
	return s.String()
}

func (s *CreateApiMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApiMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApiMcpServerResponse) GetBody() *CreateApiMcpServerResponseBody {
	return s.Body
}

func (s *CreateApiMcpServerResponse) SetHeaders(v map[string]*string) *CreateApiMcpServerResponse {
	s.Headers = v
	return s
}

func (s *CreateApiMcpServerResponse) SetStatusCode(v int32) *CreateApiMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiMcpServerResponse) SetBody(v *CreateApiMcpServerResponseBody) *CreateApiMcpServerResponse {
	s.Body = v
	return s
}

func (s *CreateApiMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

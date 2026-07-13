// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcpResponse
	GetStatusCode() *int32
	SetBody(v *GetMcpResponseBody) *GetMcpResponse
	GetBody() *GetMcpResponseBody
}

type GetMcpResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponse) GoString() string {
	return s.String()
}

func (s *GetMcpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcpResponse) GetBody() *GetMcpResponseBody {
	return s.Body
}

func (s *GetMcpResponse) SetHeaders(v map[string]*string) *GetMcpResponse {
	s.Headers = v
	return s
}

func (s *GetMcpResponse) SetStatusCode(v int32) *GetMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpResponse) SetBody(v *GetMcpResponseBody) *GetMcpResponse {
	s.Body = v
	return s
}

func (s *GetMcpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

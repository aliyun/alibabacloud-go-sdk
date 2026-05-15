// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcuLvsIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcuLvsIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcuLvsIpResponse
	GetStatusCode() *int32
	SetBody(v *GetMcuLvsIpResponseBody) *GetMcuLvsIpResponse
	GetBody() *GetMcuLvsIpResponseBody
}

type GetMcuLvsIpResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcuLvsIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcuLvsIpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcuLvsIpResponse) GoString() string {
	return s.String()
}

func (s *GetMcuLvsIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcuLvsIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcuLvsIpResponse) GetBody() *GetMcuLvsIpResponseBody {
	return s.Body
}

func (s *GetMcuLvsIpResponse) SetHeaders(v map[string]*string) *GetMcuLvsIpResponse {
	s.Headers = v
	return s
}

func (s *GetMcuLvsIpResponse) SetStatusCode(v int32) *GetMcuLvsIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcuLvsIpResponse) SetBody(v *GetMcuLvsIpResponseBody) *GetMcuLvsIpResponse {
	s.Body = v
	return s
}

func (s *GetMcuLvsIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

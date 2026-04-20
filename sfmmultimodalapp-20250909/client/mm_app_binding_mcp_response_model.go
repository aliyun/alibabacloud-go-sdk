// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMmAppBindingMcpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MmAppBindingMcpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MmAppBindingMcpResponse
	GetStatusCode() *int32
	SetBody(v *MmAppBindingMcpResponseBody) *MmAppBindingMcpResponse
	GetBody() *MmAppBindingMcpResponseBody
}

type MmAppBindingMcpResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MmAppBindingMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MmAppBindingMcpResponse) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingMcpResponse) GoString() string {
	return s.String()
}

func (s *MmAppBindingMcpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MmAppBindingMcpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MmAppBindingMcpResponse) GetBody() *MmAppBindingMcpResponseBody {
	return s.Body
}

func (s *MmAppBindingMcpResponse) SetHeaders(v map[string]*string) *MmAppBindingMcpResponse {
	s.Headers = v
	return s
}

func (s *MmAppBindingMcpResponse) SetStatusCode(v int32) *MmAppBindingMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *MmAppBindingMcpResponse) SetBody(v *MmAppBindingMcpResponseBody) *MmAppBindingMcpResponse {
	s.Body = v
	return s
}

func (s *MmAppBindingMcpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

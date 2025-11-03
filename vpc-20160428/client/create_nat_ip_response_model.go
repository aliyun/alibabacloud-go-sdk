// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNatIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNatIpResponse
	GetStatusCode() *int32
	SetBody(v *CreateNatIpResponseBody) *CreateNatIpResponse
	GetBody() *CreateNatIpResponseBody
}

type CreateNatIpResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNatIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNatIpResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNatIpResponse) GoString() string {
	return s.String()
}

func (s *CreateNatIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNatIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNatIpResponse) GetBody() *CreateNatIpResponseBody {
	return s.Body
}

func (s *CreateNatIpResponse) SetHeaders(v map[string]*string) *CreateNatIpResponse {
	s.Headers = v
	return s
}

func (s *CreateNatIpResponse) SetStatusCode(v int32) *CreateNatIpResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNatIpResponse) SetBody(v *CreateNatIpResponseBody) *CreateNatIpResponse {
	s.Body = v
	return s
}

func (s *CreateNatIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

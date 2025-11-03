// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDhcpOptionsSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDhcpOptionsSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDhcpOptionsSetResponse
	GetStatusCode() *int32
	SetBody(v *CreateDhcpOptionsSetResponseBody) *CreateDhcpOptionsSetResponse
	GetBody() *CreateDhcpOptionsSetResponseBody
}

type CreateDhcpOptionsSetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDhcpOptionsSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDhcpOptionsSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDhcpOptionsSetResponse) GoString() string {
	return s.String()
}

func (s *CreateDhcpOptionsSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDhcpOptionsSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDhcpOptionsSetResponse) GetBody() *CreateDhcpOptionsSetResponseBody {
	return s.Body
}

func (s *CreateDhcpOptionsSetResponse) SetHeaders(v map[string]*string) *CreateDhcpOptionsSetResponse {
	s.Headers = v
	return s
}

func (s *CreateDhcpOptionsSetResponse) SetStatusCode(v int32) *CreateDhcpOptionsSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDhcpOptionsSetResponse) SetBody(v *CreateDhcpOptionsSetResponseBody) *CreateDhcpOptionsSetResponse {
	s.Body = v
	return s
}

func (s *CreateDhcpOptionsSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

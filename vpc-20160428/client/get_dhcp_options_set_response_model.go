// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDhcpOptionsSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDhcpOptionsSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDhcpOptionsSetResponse
	GetStatusCode() *int32
	SetBody(v *GetDhcpOptionsSetResponseBody) *GetDhcpOptionsSetResponse
	GetBody() *GetDhcpOptionsSetResponseBody
}

type GetDhcpOptionsSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDhcpOptionsSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDhcpOptionsSetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDhcpOptionsSetResponse) GoString() string {
	return s.String()
}

func (s *GetDhcpOptionsSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDhcpOptionsSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDhcpOptionsSetResponse) GetBody() *GetDhcpOptionsSetResponseBody {
	return s.Body
}

func (s *GetDhcpOptionsSetResponse) SetHeaders(v map[string]*string) *GetDhcpOptionsSetResponse {
	s.Headers = v
	return s
}

func (s *GetDhcpOptionsSetResponse) SetStatusCode(v int32) *GetDhcpOptionsSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDhcpOptionsSetResponse) SetBody(v *GetDhcpOptionsSetResponseBody) *GetDhcpOptionsSetResponse {
	s.Body = v
	return s
}

func (s *GetDhcpOptionsSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

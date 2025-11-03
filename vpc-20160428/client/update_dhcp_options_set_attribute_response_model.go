// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDhcpOptionsSetAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDhcpOptionsSetAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDhcpOptionsSetAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDhcpOptionsSetAttributeResponseBody) *UpdateDhcpOptionsSetAttributeResponse
	GetBody() *UpdateDhcpOptionsSetAttributeResponseBody
}

type UpdateDhcpOptionsSetAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDhcpOptionsSetAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDhcpOptionsSetAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDhcpOptionsSetAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateDhcpOptionsSetAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDhcpOptionsSetAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDhcpOptionsSetAttributeResponse) GetBody() *UpdateDhcpOptionsSetAttributeResponseBody {
	return s.Body
}

func (s *UpdateDhcpOptionsSetAttributeResponse) SetHeaders(v map[string]*string) *UpdateDhcpOptionsSetAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeResponse) SetStatusCode(v int32) *UpdateDhcpOptionsSetAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeResponse) SetBody(v *UpdateDhcpOptionsSetAttributeResponseBody) *UpdateDhcpOptionsSetAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

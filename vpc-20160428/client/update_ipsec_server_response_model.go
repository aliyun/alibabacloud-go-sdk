// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpsecServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIpsecServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIpsecServerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIpsecServerResponseBody) *UpdateIpsecServerResponse
	GetBody() *UpdateIpsecServerResponseBody
}

type UpdateIpsecServerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpsecServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpsecServerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpsecServerResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpsecServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIpsecServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIpsecServerResponse) GetBody() *UpdateIpsecServerResponseBody {
	return s.Body
}

func (s *UpdateIpsecServerResponse) SetHeaders(v map[string]*string) *UpdateIpsecServerResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpsecServerResponse) SetStatusCode(v int32) *UpdateIpsecServerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpsecServerResponse) SetBody(v *UpdateIpsecServerResponseBody) *UpdateIpsecServerResponse {
	s.Body = v
	return s
}

func (s *UpdateIpsecServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

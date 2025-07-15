// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpsecServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpsecServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpsecServerResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpsecServerResponseBody) *CreateIpsecServerResponse
	GetBody() *CreateIpsecServerResponseBody
}

type CreateIpsecServerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpsecServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpsecServerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpsecServerResponse) GoString() string {
	return s.String()
}

func (s *CreateIpsecServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpsecServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpsecServerResponse) GetBody() *CreateIpsecServerResponseBody {
	return s.Body
}

func (s *CreateIpsecServerResponse) SetHeaders(v map[string]*string) *CreateIpsecServerResponse {
	s.Headers = v
	return s
}

func (s *CreateIpsecServerResponse) SetStatusCode(v int32) *CreateIpsecServerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpsecServerResponse) SetBody(v *CreateIpsecServerResponseBody) *CreateIpsecServerResponse {
	s.Body = v
	return s
}

func (s *CreateIpsecServerResponse) Validate() error {
	return dara.Validate(s)
}

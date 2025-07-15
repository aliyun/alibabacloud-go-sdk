// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatIpCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNatIpCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNatIpCidrResponse
	GetStatusCode() *int32
	SetBody(v *CreateNatIpCidrResponseBody) *CreateNatIpCidrResponse
	GetBody() *CreateNatIpCidrResponseBody
}

type CreateNatIpCidrResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNatIpCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNatIpCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNatIpCidrResponse) GoString() string {
	return s.String()
}

func (s *CreateNatIpCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNatIpCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNatIpCidrResponse) GetBody() *CreateNatIpCidrResponseBody {
	return s.Body
}

func (s *CreateNatIpCidrResponse) SetHeaders(v map[string]*string) *CreateNatIpCidrResponse {
	s.Headers = v
	return s
}

func (s *CreateNatIpCidrResponse) SetStatusCode(v int32) *CreateNatIpCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNatIpCidrResponse) SetBody(v *CreateNatIpCidrResponseBody) *CreateNatIpCidrResponse {
	s.Body = v
	return s
}

func (s *CreateNatIpCidrResponse) Validate() error {
	return dara.Validate(s)
}

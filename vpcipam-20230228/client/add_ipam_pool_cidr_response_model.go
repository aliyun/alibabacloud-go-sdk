// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpamPoolCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddIpamPoolCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddIpamPoolCidrResponse
	GetStatusCode() *int32
	SetBody(v *AddIpamPoolCidrResponseBody) *AddIpamPoolCidrResponse
	GetBody() *AddIpamPoolCidrResponseBody
}

type AddIpamPoolCidrResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIpamPoolCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIpamPoolCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s AddIpamPoolCidrResponse) GoString() string {
	return s.String()
}

func (s *AddIpamPoolCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddIpamPoolCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddIpamPoolCidrResponse) GetBody() *AddIpamPoolCidrResponseBody {
	return s.Body
}

func (s *AddIpamPoolCidrResponse) SetHeaders(v map[string]*string) *AddIpamPoolCidrResponse {
	s.Headers = v
	return s
}

func (s *AddIpamPoolCidrResponse) SetStatusCode(v int32) *AddIpamPoolCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpamPoolCidrResponse) SetBody(v *AddIpamPoolCidrResponseBody) *AddIpamPoolCidrResponse {
	s.Body = v
	return s
}

func (s *AddIpamPoolCidrResponse) Validate() error {
	return dara.Validate(s)
}

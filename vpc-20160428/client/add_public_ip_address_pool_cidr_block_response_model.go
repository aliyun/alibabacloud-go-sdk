// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPublicIpAddressPoolCidrBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPublicIpAddressPoolCidrBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPublicIpAddressPoolCidrBlockResponse
	GetStatusCode() *int32
	SetBody(v *AddPublicIpAddressPoolCidrBlockResponseBody) *AddPublicIpAddressPoolCidrBlockResponse
	GetBody() *AddPublicIpAddressPoolCidrBlockResponseBody
}

type AddPublicIpAddressPoolCidrBlockResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPublicIpAddressPoolCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPublicIpAddressPoolCidrBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPublicIpAddressPoolCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *AddPublicIpAddressPoolCidrBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPublicIpAddressPoolCidrBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPublicIpAddressPoolCidrBlockResponse) GetBody() *AddPublicIpAddressPoolCidrBlockResponseBody {
	return s.Body
}

func (s *AddPublicIpAddressPoolCidrBlockResponse) SetHeaders(v map[string]*string) *AddPublicIpAddressPoolCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockResponse) SetStatusCode(v int32) *AddPublicIpAddressPoolCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockResponse) SetBody(v *AddPublicIpAddressPoolCidrBlockResponseBody) *AddPublicIpAddressPoolCidrBlockResponse {
	s.Body = v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

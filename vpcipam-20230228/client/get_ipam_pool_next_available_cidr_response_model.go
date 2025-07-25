// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpamPoolNextAvailableCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIpamPoolNextAvailableCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIpamPoolNextAvailableCidrResponse
	GetStatusCode() *int32
	SetBody(v *GetIpamPoolNextAvailableCidrResponseBody) *GetIpamPoolNextAvailableCidrResponse
	GetBody() *GetIpamPoolNextAvailableCidrResponseBody
}

type GetIpamPoolNextAvailableCidrResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpamPoolNextAvailableCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpamPoolNextAvailableCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIpamPoolNextAvailableCidrResponse) GoString() string {
	return s.String()
}

func (s *GetIpamPoolNextAvailableCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIpamPoolNextAvailableCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIpamPoolNextAvailableCidrResponse) GetBody() *GetIpamPoolNextAvailableCidrResponseBody {
	return s.Body
}

func (s *GetIpamPoolNextAvailableCidrResponse) SetHeaders(v map[string]*string) *GetIpamPoolNextAvailableCidrResponse {
	s.Headers = v
	return s
}

func (s *GetIpamPoolNextAvailableCidrResponse) SetStatusCode(v int32) *GetIpamPoolNextAvailableCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrResponse) SetBody(v *GetIpamPoolNextAvailableCidrResponseBody) *GetIpamPoolNextAvailableCidrResponse {
	s.Body = v
	return s
}

func (s *GetIpamPoolNextAvailableCidrResponse) Validate() error {
	return dara.Validate(s)
}

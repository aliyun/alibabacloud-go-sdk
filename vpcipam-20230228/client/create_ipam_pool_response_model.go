// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpamPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpamPoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpamPoolResponseBody) *CreateIpamPoolResponse
	GetBody() *CreateIpamPoolResponseBody
}

type CreateIpamPoolResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpamPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpamPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamPoolResponse) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpamPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpamPoolResponse) GetBody() *CreateIpamPoolResponseBody {
	return s.Body
}

func (s *CreateIpamPoolResponse) SetHeaders(v map[string]*string) *CreateIpamPoolResponse {
	s.Headers = v
	return s
}

func (s *CreateIpamPoolResponse) SetStatusCode(v int32) *CreateIpamPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpamPoolResponse) SetBody(v *CreateIpamPoolResponseBody) *CreateIpamPoolResponse {
	s.Body = v
	return s
}

func (s *CreateIpamPoolResponse) Validate() error {
	return dara.Validate(s)
}

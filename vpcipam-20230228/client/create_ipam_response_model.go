// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpamResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpamResponseBody) *CreateIpamResponse
	GetBody() *CreateIpamResponseBody
}

type CreateIpamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpamResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamResponse) GoString() string {
	return s.String()
}

func (s *CreateIpamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpamResponse) GetBody() *CreateIpamResponseBody {
	return s.Body
}

func (s *CreateIpamResponse) SetHeaders(v map[string]*string) *CreateIpamResponse {
	s.Headers = v
	return s
}

func (s *CreateIpamResponse) SetStatusCode(v int32) *CreateIpamResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpamResponse) SetBody(v *CreateIpamResponseBody) *CreateIpamResponse {
	s.Body = v
	return s
}

func (s *CreateIpamResponse) Validate() error {
	return dara.Validate(s)
}

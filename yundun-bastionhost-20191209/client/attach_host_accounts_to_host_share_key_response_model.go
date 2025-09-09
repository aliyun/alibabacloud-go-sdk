// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostAccountsToHostShareKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachHostAccountsToHostShareKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachHostAccountsToHostShareKeyResponse
	GetStatusCode() *int32
	SetBody(v *AttachHostAccountsToHostShareKeyResponseBody) *AttachHostAccountsToHostShareKeyResponse
	GetBody() *AttachHostAccountsToHostShareKeyResponseBody
}

type AttachHostAccountsToHostShareKeyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachHostAccountsToHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachHostAccountsToHostShareKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToHostShareKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachHostAccountsToHostShareKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachHostAccountsToHostShareKeyResponse) GetBody() *AttachHostAccountsToHostShareKeyResponseBody {
	return s.Body
}

func (s *AttachHostAccountsToHostShareKeyResponse) SetHeaders(v map[string]*string) *AttachHostAccountsToHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponse) SetStatusCode(v int32) *AttachHostAccountsToHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponse) SetBody(v *AttachHostAccountsToHostShareKeyResponseBody) *AttachHostAccountsToHostShareKeyResponse {
	s.Body = v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponse) Validate() error {
	return dara.Validate(s)
}

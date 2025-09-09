// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostGroupAccountsToUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachHostGroupAccountsToUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachHostGroupAccountsToUserResponse
	GetStatusCode() *int32
	SetBody(v *AttachHostGroupAccountsToUserResponseBody) *AttachHostGroupAccountsToUserResponse
	GetBody() *AttachHostGroupAccountsToUserResponseBody
}

type AttachHostGroupAccountsToUserResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachHostGroupAccountsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachHostGroupAccountsToUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachHostGroupAccountsToUserResponse) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachHostGroupAccountsToUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachHostGroupAccountsToUserResponse) GetBody() *AttachHostGroupAccountsToUserResponseBody {
	return s.Body
}

func (s *AttachHostGroupAccountsToUserResponse) SetHeaders(v map[string]*string) *AttachHostGroupAccountsToUserResponse {
	s.Headers = v
	return s
}

func (s *AttachHostGroupAccountsToUserResponse) SetStatusCode(v int32) *AttachHostGroupAccountsToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponse) SetBody(v *AttachHostGroupAccountsToUserResponseBody) *AttachHostGroupAccountsToUserResponse {
	s.Body = v
	return s
}

func (s *AttachHostGroupAccountsToUserResponse) Validate() error {
	return dara.Validate(s)
}

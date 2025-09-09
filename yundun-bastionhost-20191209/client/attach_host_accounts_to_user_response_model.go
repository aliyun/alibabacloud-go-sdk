// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostAccountsToUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachHostAccountsToUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachHostAccountsToUserResponse
	GetStatusCode() *int32
	SetBody(v *AttachHostAccountsToUserResponseBody) *AttachHostAccountsToUserResponse
	GetBody() *AttachHostAccountsToUserResponseBody
}

type AttachHostAccountsToUserResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachHostAccountsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachHostAccountsToUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToUserResponse) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachHostAccountsToUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachHostAccountsToUserResponse) GetBody() *AttachHostAccountsToUserResponseBody {
	return s.Body
}

func (s *AttachHostAccountsToUserResponse) SetHeaders(v map[string]*string) *AttachHostAccountsToUserResponse {
	s.Headers = v
	return s
}

func (s *AttachHostAccountsToUserResponse) SetStatusCode(v int32) *AttachHostAccountsToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachHostAccountsToUserResponse) SetBody(v *AttachHostAccountsToUserResponseBody) *AttachHostAccountsToUserResponse {
	s.Body = v
	return s
}

func (s *AttachHostAccountsToUserResponse) Validate() error {
	return dara.Validate(s)
}

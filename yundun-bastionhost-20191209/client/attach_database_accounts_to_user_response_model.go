// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDatabaseAccountsToUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachDatabaseAccountsToUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachDatabaseAccountsToUserResponse
	GetStatusCode() *int32
	SetBody(v *AttachDatabaseAccountsToUserResponseBody) *AttachDatabaseAccountsToUserResponse
	GetBody() *AttachDatabaseAccountsToUserResponseBody
}

type AttachDatabaseAccountsToUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachDatabaseAccountsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachDatabaseAccountsToUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserResponse) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachDatabaseAccountsToUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachDatabaseAccountsToUserResponse) GetBody() *AttachDatabaseAccountsToUserResponseBody {
	return s.Body
}

func (s *AttachDatabaseAccountsToUserResponse) SetHeaders(v map[string]*string) *AttachDatabaseAccountsToUserResponse {
	s.Headers = v
	return s
}

func (s *AttachDatabaseAccountsToUserResponse) SetStatusCode(v int32) *AttachDatabaseAccountsToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDatabaseAccountsToUserResponse) SetBody(v *AttachDatabaseAccountsToUserResponseBody) *AttachDatabaseAccountsToUserResponse {
	s.Body = v
	return s
}

func (s *AttachDatabaseAccountsToUserResponse) Validate() error {
	return dara.Validate(s)
}

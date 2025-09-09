// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostGroupAccountsFromUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachHostGroupAccountsFromUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachHostGroupAccountsFromUserResponse
	GetStatusCode() *int32
	SetBody(v *DetachHostGroupAccountsFromUserResponseBody) *DetachHostGroupAccountsFromUserResponse
	GetBody() *DetachHostGroupAccountsFromUserResponseBody
}

type DetachHostGroupAccountsFromUserResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachHostGroupAccountsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachHostGroupAccountsFromUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserResponse) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachHostGroupAccountsFromUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachHostGroupAccountsFromUserResponse) GetBody() *DetachHostGroupAccountsFromUserResponseBody {
	return s.Body
}

func (s *DetachHostGroupAccountsFromUserResponse) SetHeaders(v map[string]*string) *DetachHostGroupAccountsFromUserResponse {
	s.Headers = v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponse) SetStatusCode(v int32) *DetachHostGroupAccountsFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponse) SetBody(v *DetachHostGroupAccountsFromUserResponseBody) *DetachHostGroupAccountsFromUserResponse {
	s.Body = v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponse) Validate() error {
	return dara.Validate(s)
}

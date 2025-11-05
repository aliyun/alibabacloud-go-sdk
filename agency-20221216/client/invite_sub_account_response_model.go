// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteSubAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InviteSubAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InviteSubAccountResponse
	GetStatusCode() *int32
	SetBody(v *InviteSubAccountResponseBody) *InviteSubAccountResponse
	GetBody() *InviteSubAccountResponseBody
}

type InviteSubAccountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InviteSubAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InviteSubAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s InviteSubAccountResponse) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InviteSubAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InviteSubAccountResponse) GetBody() *InviteSubAccountResponseBody {
	return s.Body
}

func (s *InviteSubAccountResponse) SetHeaders(v map[string]*string) *InviteSubAccountResponse {
	s.Headers = v
	return s
}

func (s *InviteSubAccountResponse) SetStatusCode(v int32) *InviteSubAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *InviteSubAccountResponse) SetBody(v *InviteSubAccountResponseBody) *InviteSubAccountResponse {
	s.Body = v
	return s
}

func (s *InviteSubAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

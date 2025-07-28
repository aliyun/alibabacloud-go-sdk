// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemberAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMemberAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMemberAccountsResponse
	GetStatusCode() *int32
	SetBody(v *CreateMemberAccountsResponseBody) *CreateMemberAccountsResponse
	GetBody() *CreateMemberAccountsResponseBody
}

type CreateMemberAccountsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemberAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemberAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberAccountsResponse) GoString() string {
	return s.String()
}

func (s *CreateMemberAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMemberAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMemberAccountsResponse) GetBody() *CreateMemberAccountsResponseBody {
	return s.Body
}

func (s *CreateMemberAccountsResponse) SetHeaders(v map[string]*string) *CreateMemberAccountsResponse {
	s.Headers = v
	return s
}

func (s *CreateMemberAccountsResponse) SetStatusCode(v int32) *CreateMemberAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemberAccountsResponse) SetBody(v *CreateMemberAccountsResponseBody) *CreateMemberAccountsResponse {
	s.Body = v
	return s
}

func (s *CreateMemberAccountsResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMemberAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMemberAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMemberAccountResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMemberAccountResponseBody) *ModifyMemberAccountResponse
	GetBody() *ModifyMemberAccountResponseBody
}

type ModifyMemberAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMemberAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMemberAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMemberAccountResponse) GoString() string {
	return s.String()
}

func (s *ModifyMemberAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMemberAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMemberAccountResponse) GetBody() *ModifyMemberAccountResponseBody {
	return s.Body
}

func (s *ModifyMemberAccountResponse) SetHeaders(v map[string]*string) *ModifyMemberAccountResponse {
	s.Headers = v
	return s
}

func (s *ModifyMemberAccountResponse) SetStatusCode(v int32) *ModifyMemberAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMemberAccountResponse) SetBody(v *ModifyMemberAccountResponseBody) *ModifyMemberAccountResponse {
	s.Body = v
	return s
}

func (s *ModifyMemberAccountResponse) Validate() error {
	return dara.Validate(s)
}

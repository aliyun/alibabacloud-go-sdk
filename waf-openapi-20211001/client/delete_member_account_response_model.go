// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemberAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMemberAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMemberAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMemberAccountResponseBody) *DeleteMemberAccountResponse
	GetBody() *DeleteMemberAccountResponseBody
}

type DeleteMemberAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemberAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemberAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemberAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemberAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMemberAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMemberAccountResponse) GetBody() *DeleteMemberAccountResponseBody {
	return s.Body
}

func (s *DeleteMemberAccountResponse) SetHeaders(v map[string]*string) *DeleteMemberAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemberAccountResponse) SetStatusCode(v int32) *DeleteMemberAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemberAccountResponse) SetBody(v *DeleteMemberAccountResponseBody) *DeleteMemberAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteMemberAccountResponse) Validate() error {
	return dara.Validate(s)
}

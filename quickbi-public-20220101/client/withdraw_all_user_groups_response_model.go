// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawAllUserGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WithdrawAllUserGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WithdrawAllUserGroupsResponse
	GetStatusCode() *int32
	SetBody(v *WithdrawAllUserGroupsResponseBody) *WithdrawAllUserGroupsResponse
	GetBody() *WithdrawAllUserGroupsResponseBody
}

type WithdrawAllUserGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WithdrawAllUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WithdrawAllUserGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s WithdrawAllUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *WithdrawAllUserGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WithdrawAllUserGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WithdrawAllUserGroupsResponse) GetBody() *WithdrawAllUserGroupsResponseBody {
	return s.Body
}

func (s *WithdrawAllUserGroupsResponse) SetHeaders(v map[string]*string) *WithdrawAllUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *WithdrawAllUserGroupsResponse) SetStatusCode(v int32) *WithdrawAllUserGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *WithdrawAllUserGroupsResponse) SetBody(v *WithdrawAllUserGroupsResponseBody) *WithdrawAllUserGroupsResponse {
	s.Body = v
	return s
}

func (s *WithdrawAllUserGroupsResponse) Validate() error {
	return dara.Validate(s)
}

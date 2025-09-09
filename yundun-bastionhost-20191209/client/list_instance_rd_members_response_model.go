// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRdMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceRdMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceRdMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceRdMembersResponseBody) *ListInstanceRdMembersResponse
	GetBody() *ListInstanceRdMembersResponseBody
}

type ListInstanceRdMembersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceRdMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceRdMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRdMembersResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceRdMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceRdMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceRdMembersResponse) GetBody() *ListInstanceRdMembersResponseBody {
	return s.Body
}

func (s *ListInstanceRdMembersResponse) SetHeaders(v map[string]*string) *ListInstanceRdMembersResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceRdMembersResponse) SetStatusCode(v int32) *ListInstanceRdMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceRdMembersResponse) SetBody(v *ListInstanceRdMembersResponseBody) *ListInstanceRdMembersResponse {
	s.Body = v
	return s
}

func (s *ListInstanceRdMembersResponse) Validate() error {
	return dara.Validate(s)
}

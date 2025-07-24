// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMembersResponse
	GetStatusCode() *int32
	SetBody(v *AddMembersResponseBody) *AddMembersResponse
	GetBody() *AddMembersResponseBody
}

type AddMembersResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMembersResponse) GoString() string {
	return s.String()
}

func (s *AddMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMembersResponse) GetBody() *AddMembersResponseBody {
	return s.Body
}

func (s *AddMembersResponse) SetHeaders(v map[string]*string) *AddMembersResponse {
	s.Headers = v
	return s
}

func (s *AddMembersResponse) SetStatusCode(v int32) *AddMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMembersResponse) SetBody(v *AddMembersResponseBody) *AddMembersResponse {
	s.Body = v
	return s
}

func (s *AddMembersResponse) Validate() error {
	return dara.Validate(s)
}

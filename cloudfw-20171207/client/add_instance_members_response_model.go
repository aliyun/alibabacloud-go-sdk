// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddInstanceMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddInstanceMembersResponse
	GetStatusCode() *int32
	SetBody(v *AddInstanceMembersResponseBody) *AddInstanceMembersResponse
	GetBody() *AddInstanceMembersResponseBody
}

type AddInstanceMembersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddInstanceMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceMembersResponse) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddInstanceMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddInstanceMembersResponse) GetBody() *AddInstanceMembersResponseBody {
	return s.Body
}

func (s *AddInstanceMembersResponse) SetHeaders(v map[string]*string) *AddInstanceMembersResponse {
	s.Headers = v
	return s
}

func (s *AddInstanceMembersResponse) SetStatusCode(v int32) *AddInstanceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddInstanceMembersResponse) SetBody(v *AddInstanceMembersResponseBody) *AddInstanceMembersResponse {
	s.Body = v
	return s
}

func (s *AddInstanceMembersResponse) Validate() error {
	return dara.Validate(s)
}

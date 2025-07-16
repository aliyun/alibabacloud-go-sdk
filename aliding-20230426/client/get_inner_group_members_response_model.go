// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerGroupMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInnerGroupMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInnerGroupMembersResponse
	GetStatusCode() *int32
	SetBody(v *GetInnerGroupMembersResponseBody) *GetInnerGroupMembersResponse
	GetBody() *GetInnerGroupMembersResponseBody
}

type GetInnerGroupMembersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInnerGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInnerGroupMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInnerGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *GetInnerGroupMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInnerGroupMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInnerGroupMembersResponse) GetBody() *GetInnerGroupMembersResponseBody {
	return s.Body
}

func (s *GetInnerGroupMembersResponse) SetHeaders(v map[string]*string) *GetInnerGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *GetInnerGroupMembersResponse) SetStatusCode(v int32) *GetInnerGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInnerGroupMembersResponse) SetBody(v *GetInnerGroupMembersResponseBody) *GetInnerGroupMembersResponse {
	s.Body = v
	return s
}

func (s *GetInnerGroupMembersResponse) Validate() error {
	return dara.Validate(s)
}

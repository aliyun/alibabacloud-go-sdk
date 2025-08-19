// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateMembersResponse
	GetStatusCode() *int32
	SetBody(v *AssociateMembersResponseBody) *AssociateMembersResponse
	GetBody() *AssociateMembersResponseBody
}

type AssociateMembersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateMembersResponse) GoString() string {
	return s.String()
}

func (s *AssociateMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateMembersResponse) GetBody() *AssociateMembersResponseBody {
	return s.Body
}

func (s *AssociateMembersResponse) SetHeaders(v map[string]*string) *AssociateMembersResponse {
	s.Headers = v
	return s
}

func (s *AssociateMembersResponse) SetStatusCode(v int32) *AssociateMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateMembersResponse) SetBody(v *AssociateMembersResponseBody) *AssociateMembersResponse {
	s.Body = v
	return s
}

func (s *AssociateMembersResponse) Validate() error {
	return dara.Validate(s)
}

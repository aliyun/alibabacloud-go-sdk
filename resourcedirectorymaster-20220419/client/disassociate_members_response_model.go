// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateMembersResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateMembersResponseBody) *DisassociateMembersResponse
	GetBody() *DisassociateMembersResponseBody
}

type DisassociateMembersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateMembersResponse) GoString() string {
	return s.String()
}

func (s *DisassociateMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateMembersResponse) GetBody() *DisassociateMembersResponseBody {
	return s.Body
}

func (s *DisassociateMembersResponse) SetHeaders(v map[string]*string) *DisassociateMembersResponse {
	s.Headers = v
	return s
}

func (s *DisassociateMembersResponse) SetStatusCode(v int32) *DisassociateMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateMembersResponse) SetBody(v *DisassociateMembersResponseBody) *DisassociateMembersResponse {
	s.Body = v
	return s
}

func (s *DisassociateMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

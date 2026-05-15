// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MuteMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MuteMembersResponse
	GetStatusCode() *int32
	SetBody(v *MuteMembersResponseBody) *MuteMembersResponse
	GetBody() *MuteMembersResponseBody
}

type MuteMembersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MuteMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MuteMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s MuteMembersResponse) GoString() string {
	return s.String()
}

func (s *MuteMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MuteMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MuteMembersResponse) GetBody() *MuteMembersResponseBody {
	return s.Body
}

func (s *MuteMembersResponse) SetHeaders(v map[string]*string) *MuteMembersResponse {
	s.Headers = v
	return s
}

func (s *MuteMembersResponse) SetStatusCode(v int32) *MuteMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *MuteMembersResponse) SetBody(v *MuteMembersResponseBody) *MuteMembersResponse {
	s.Body = v
	return s
}

func (s *MuteMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

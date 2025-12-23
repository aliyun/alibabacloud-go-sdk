// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpamMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddIpamMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddIpamMembersResponse
	GetStatusCode() *int32
	SetBody(v *AddIpamMembersResponseBody) *AddIpamMembersResponse
	GetBody() *AddIpamMembersResponseBody
}

type AddIpamMembersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIpamMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIpamMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddIpamMembersResponse) GoString() string {
	return s.String()
}

func (s *AddIpamMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddIpamMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddIpamMembersResponse) GetBody() *AddIpamMembersResponseBody {
	return s.Body
}

func (s *AddIpamMembersResponse) SetHeaders(v map[string]*string) *AddIpamMembersResponse {
	s.Headers = v
	return s
}

func (s *AddIpamMembersResponse) SetStatusCode(v int32) *AddIpamMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpamMembersResponse) SetBody(v *AddIpamMembersResponseBody) *AddIpamMembersResponse {
	s.Body = v
	return s
}

func (s *AddIpamMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

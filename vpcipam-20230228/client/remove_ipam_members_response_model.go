// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpamMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveIpamMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveIpamMembersResponse
	GetStatusCode() *int32
	SetBody(v *RemoveIpamMembersResponseBody) *RemoveIpamMembersResponse
	GetBody() *RemoveIpamMembersResponseBody
}

type RemoveIpamMembersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveIpamMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveIpamMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpamMembersResponse) GoString() string {
	return s.String()
}

func (s *RemoveIpamMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveIpamMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveIpamMembersResponse) GetBody() *RemoveIpamMembersResponseBody {
	return s.Body
}

func (s *RemoveIpamMembersResponse) SetHeaders(v map[string]*string) *RemoveIpamMembersResponse {
	s.Headers = v
	return s
}

func (s *RemoveIpamMembersResponse) SetStatusCode(v int32) *RemoveIpamMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveIpamMembersResponse) SetBody(v *RemoveIpamMembersResponseBody) *RemoveIpamMembersResponse {
	s.Body = v
	return s
}

func (s *RemoveIpamMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

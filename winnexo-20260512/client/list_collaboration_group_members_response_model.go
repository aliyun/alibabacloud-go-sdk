// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCollaborationGroupMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCollaborationGroupMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListCollaborationGroupMembersResponseBody) *ListCollaborationGroupMembersResponse
	GetBody() *ListCollaborationGroupMembersResponseBody
}

type ListCollaborationGroupMembersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCollaborationGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCollaborationGroupMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCollaborationGroupMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCollaborationGroupMembersResponse) GetBody() *ListCollaborationGroupMembersResponseBody {
	return s.Body
}

func (s *ListCollaborationGroupMembersResponse) SetHeaders(v map[string]*string) *ListCollaborationGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *ListCollaborationGroupMembersResponse) SetStatusCode(v int32) *ListCollaborationGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCollaborationGroupMembersResponse) SetBody(v *ListCollaborationGroupMembersResponseBody) *ListCollaborationGroupMembersResponse {
	s.Body = v
	return s
}

func (s *ListCollaborationGroupMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

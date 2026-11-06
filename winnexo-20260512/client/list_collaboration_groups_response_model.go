// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCollaborationGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCollaborationGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListCollaborationGroupsResponseBody) *ListCollaborationGroupsResponse
	GetBody() *ListCollaborationGroupsResponseBody
}

type ListCollaborationGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCollaborationGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCollaborationGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCollaborationGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCollaborationGroupsResponse) GetBody() *ListCollaborationGroupsResponseBody {
	return s.Body
}

func (s *ListCollaborationGroupsResponse) SetHeaders(v map[string]*string) *ListCollaborationGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListCollaborationGroupsResponse) SetStatusCode(v int32) *ListCollaborationGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCollaborationGroupsResponse) SetBody(v *ListCollaborationGroupsResponseBody) *ListCollaborationGroupsResponse {
	s.Body = v
	return s
}

func (s *ListCollaborationGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

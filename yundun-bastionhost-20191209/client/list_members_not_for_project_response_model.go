// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersNotForProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMembersNotForProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMembersNotForProjectResponse
	GetStatusCode() *int32
	SetBody(v *ListMembersNotForProjectResponseBody) *ListMembersNotForProjectResponse
	GetBody() *ListMembersNotForProjectResponseBody
}

type ListMembersNotForProjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMembersNotForProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMembersNotForProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMembersNotForProjectResponse) GoString() string {
	return s.String()
}

func (s *ListMembersNotForProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMembersNotForProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMembersNotForProjectResponse) GetBody() *ListMembersNotForProjectResponseBody {
	return s.Body
}

func (s *ListMembersNotForProjectResponse) SetHeaders(v map[string]*string) *ListMembersNotForProjectResponse {
	s.Headers = v
	return s
}

func (s *ListMembersNotForProjectResponse) SetStatusCode(v int32) *ListMembersNotForProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMembersNotForProjectResponse) SetBody(v *ListMembersNotForProjectResponseBody) *ListMembersNotForProjectResponse {
	s.Body = v
	return s
}

func (s *ListMembersNotForProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachMembersFromProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachMembersFromProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachMembersFromProjectResponse
	GetStatusCode() *int32
	SetBody(v *DetachMembersFromProjectResponseBody) *DetachMembersFromProjectResponse
	GetBody() *DetachMembersFromProjectResponseBody
}

type DetachMembersFromProjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachMembersFromProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachMembersFromProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachMembersFromProjectResponse) GoString() string {
	return s.String()
}

func (s *DetachMembersFromProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachMembersFromProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachMembersFromProjectResponse) GetBody() *DetachMembersFromProjectResponseBody {
	return s.Body
}

func (s *DetachMembersFromProjectResponse) SetHeaders(v map[string]*string) *DetachMembersFromProjectResponse {
	s.Headers = v
	return s
}

func (s *DetachMembersFromProjectResponse) SetStatusCode(v int32) *DetachMembersFromProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachMembersFromProjectResponse) SetBody(v *DetachMembersFromProjectResponseBody) *DetachMembersFromProjectResponse {
	s.Body = v
	return s
}

func (s *DetachMembersFromProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

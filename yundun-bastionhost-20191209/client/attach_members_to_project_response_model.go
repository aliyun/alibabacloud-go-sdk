// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachMembersToProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachMembersToProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachMembersToProjectResponse
	GetStatusCode() *int32
	SetBody(v *AttachMembersToProjectResponseBody) *AttachMembersToProjectResponse
	GetBody() *AttachMembersToProjectResponseBody
}

type AttachMembersToProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachMembersToProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachMembersToProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachMembersToProjectResponse) GoString() string {
	return s.String()
}

func (s *AttachMembersToProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachMembersToProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachMembersToProjectResponse) GetBody() *AttachMembersToProjectResponseBody {
	return s.Body
}

func (s *AttachMembersToProjectResponse) SetHeaders(v map[string]*string) *AttachMembersToProjectResponse {
	s.Headers = v
	return s
}

func (s *AttachMembersToProjectResponse) SetStatusCode(v int32) *AttachMembersToProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachMembersToProjectResponse) SetBody(v *AttachMembersToProjectResponseBody) *AttachMembersToProjectResponse {
	s.Body = v
	return s
}

func (s *AttachMembersToProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

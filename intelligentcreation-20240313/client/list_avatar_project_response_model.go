// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvatarProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvatarProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvatarProjectResponse
	GetStatusCode() *int32
	SetBody(v *ListAvatarProjectResponseBody) *ListAvatarProjectResponse
	GetBody() *ListAvatarProjectResponseBody
}

type ListAvatarProjectResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvatarProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *ListAvatarProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvatarProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvatarProjectResponse) GetBody() *ListAvatarProjectResponseBody {
	return s.Body
}

func (s *ListAvatarProjectResponse) SetHeaders(v map[string]*string) *ListAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *ListAvatarProjectResponse) SetStatusCode(v int32) *ListAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvatarProjectResponse) SetBody(v *ListAvatarProjectResponseBody) *ListAvatarProjectResponse {
	s.Body = v
	return s
}

func (s *ListAvatarProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

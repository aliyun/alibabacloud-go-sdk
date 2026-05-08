// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvatarProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAvatarProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAvatarProjectResponse
	GetStatusCode() *int32
	SetBody(v *QueryAvatarProjectResponseBody) *QueryAvatarProjectResponse
	GetBody() *QueryAvatarProjectResponseBody
}

type QueryAvatarProjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAvatarProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAvatarProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAvatarProjectResponse) GetBody() *QueryAvatarProjectResponseBody {
	return s.Body
}

func (s *QueryAvatarProjectResponse) SetHeaders(v map[string]*string) *QueryAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *QueryAvatarProjectResponse) SetStatusCode(v int32) *QueryAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvatarProjectResponse) SetBody(v *QueryAvatarProjectResponseBody) *QueryAvatarProjectResponse {
	s.Body = v
	return s
}

func (s *QueryAvatarProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

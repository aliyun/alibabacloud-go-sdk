// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAvatarProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveAvatarProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveAvatarProjectResponse
	GetStatusCode() *int32
	SetBody(v *SaveAvatarProjectResponseBody) *SaveAvatarProjectResponse
	GetBody() *SaveAvatarProjectResponseBody
}

type SaveAvatarProjectResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveAvatarProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *SaveAvatarProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveAvatarProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveAvatarProjectResponse) GetBody() *SaveAvatarProjectResponseBody {
	return s.Body
}

func (s *SaveAvatarProjectResponse) SetHeaders(v map[string]*string) *SaveAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *SaveAvatarProjectResponse) SetStatusCode(v int32) *SaveAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAvatarProjectResponse) SetBody(v *SaveAvatarProjectResponseBody) *SaveAvatarProjectResponse {
	s.Body = v
	return s
}

func (s *SaveAvatarProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

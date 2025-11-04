// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvatarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAvatarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAvatarResponse
	GetStatusCode() *int32
	SetBody(v *GetAvatarResponseBody) *GetAvatarResponse
	GetBody() *GetAvatarResponseBody
}

type GetAvatarResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAvatarResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAvatarResponse) GoString() string {
	return s.String()
}

func (s *GetAvatarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAvatarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAvatarResponse) GetBody() *GetAvatarResponseBody {
	return s.Body
}

func (s *GetAvatarResponse) SetHeaders(v map[string]*string) *GetAvatarResponse {
	s.Headers = v
	return s
}

func (s *GetAvatarResponse) SetStatusCode(v int32) *GetAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAvatarResponse) SetBody(v *GetAvatarResponseBody) *GetAvatarResponse {
	s.Body = v
	return s
}

func (s *GetAvatarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

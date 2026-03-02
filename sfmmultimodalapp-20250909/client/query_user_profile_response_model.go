// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserProfileResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserProfileResponseBody) *QueryUserProfileResponse
	GetBody() *QueryUserProfileResponseBody
}

type QueryUserProfileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserProfileResponse) GoString() string {
	return s.String()
}

func (s *QueryUserProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserProfileResponse) GetBody() *QueryUserProfileResponseBody {
	return s.Body
}

func (s *QueryUserProfileResponse) SetHeaders(v map[string]*string) *QueryUserProfileResponse {
	s.Headers = v
	return s
}

func (s *QueryUserProfileResponse) SetStatusCode(v int32) *QueryUserProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserProfileResponse) SetBody(v *QueryUserProfileResponseBody) *QueryUserProfileResponse {
	s.Body = v
	return s
}

func (s *QueryUserProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

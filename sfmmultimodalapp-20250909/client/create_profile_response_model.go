// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProfileResponse
	GetStatusCode() *int32
	SetBody(v *CreateProfileResponseBody) *CreateProfileResponse
	GetBody() *CreateProfileResponseBody
}

type CreateProfileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProfileResponse) GoString() string {
	return s.String()
}

func (s *CreateProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProfileResponse) GetBody() *CreateProfileResponseBody {
	return s.Body
}

func (s *CreateProfileResponse) SetHeaders(v map[string]*string) *CreateProfileResponse {
	s.Headers = v
	return s
}

func (s *CreateProfileResponse) SetStatusCode(v int32) *CreateProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProfileResponse) SetBody(v *CreateProfileResponseBody) *CreateProfileResponse {
	s.Body = v
	return s
}

func (s *CreateProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLangfuseUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLangfuseUserResponse
	GetStatusCode() *int32
	SetBody(v *CreateLangfuseUserResponseBody) *CreateLangfuseUserResponse
	GetBody() *CreateLangfuseUserResponseBody
}

type CreateLangfuseUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLangfuseUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLangfuseUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseUserResponse) GoString() string {
	return s.String()
}

func (s *CreateLangfuseUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLangfuseUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLangfuseUserResponse) GetBody() *CreateLangfuseUserResponseBody {
	return s.Body
}

func (s *CreateLangfuseUserResponse) SetHeaders(v map[string]*string) *CreateLangfuseUserResponse {
	s.Headers = v
	return s
}

func (s *CreateLangfuseUserResponse) SetStatusCode(v int32) *CreateLangfuseUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLangfuseUserResponse) SetBody(v *CreateLangfuseUserResponseBody) *CreateLangfuseUserResponse {
	s.Body = v
	return s
}

func (s *CreateLangfuseUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

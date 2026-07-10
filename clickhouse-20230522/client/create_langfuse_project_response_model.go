// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLangfuseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLangfuseProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateLangfuseProjectResponseBody) *CreateLangfuseProjectResponse
	GetBody() *CreateLangfuseProjectResponseBody
}

type CreateLangfuseProjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLangfuseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLangfuseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateLangfuseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLangfuseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLangfuseProjectResponse) GetBody() *CreateLangfuseProjectResponseBody {
	return s.Body
}

func (s *CreateLangfuseProjectResponse) SetHeaders(v map[string]*string) *CreateLangfuseProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateLangfuseProjectResponse) SetStatusCode(v int32) *CreateLangfuseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLangfuseProjectResponse) SetBody(v *CreateLangfuseProjectResponseBody) *CreateLangfuseProjectResponse {
	s.Body = v
	return s
}

func (s *CreateLangfuseProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEditingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEditingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEditingProjectResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEditingProjectResponseBody) *UpdateEditingProjectResponse
	GetBody() *UpdateEditingProjectResponseBody
}

type UpdateEditingProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEditingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEditingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEditingProjectResponse) GetBody() *UpdateEditingProjectResponseBody {
	return s.Body
}

func (s *UpdateEditingProjectResponse) SetHeaders(v map[string]*string) *UpdateEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateEditingProjectResponse) SetStatusCode(v int32) *UpdateEditingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEditingProjectResponse) SetBody(v *UpdateEditingProjectResponseBody) *UpdateEditingProjectResponse {
	s.Body = v
	return s
}

func (s *UpdateEditingProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

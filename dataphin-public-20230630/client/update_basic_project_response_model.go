// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBasicProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBasicProjectResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBasicProjectResponseBody) *UpdateBasicProjectResponse
	GetBody() *UpdateBasicProjectResponseBody
}

type UpdateBasicProjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBasicProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBasicProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBasicProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBasicProjectResponse) GetBody() *UpdateBasicProjectResponseBody {
	return s.Body
}

func (s *UpdateBasicProjectResponse) SetHeaders(v map[string]*string) *UpdateBasicProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicProjectResponse) SetStatusCode(v int32) *UpdateBasicProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicProjectResponse) SetBody(v *UpdateBasicProjectResponseBody) *UpdateBasicProjectResponse {
	s.Body = v
	return s
}

func (s *UpdateBasicProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

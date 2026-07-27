// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBasicProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBasicProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateBasicProjectResponseBody) *CreateBasicProjectResponse
	GetBody() *CreateBasicProjectResponseBody
}

type CreateBasicProjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBasicProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBasicProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBasicProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBasicProjectResponse) GetBody() *CreateBasicProjectResponseBody {
	return s.Body
}

func (s *CreateBasicProjectResponse) SetHeaders(v map[string]*string) *CreateBasicProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicProjectResponse) SetStatusCode(v int32) *CreateBasicProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicProjectResponse) SetBody(v *CreateBasicProjectResponseBody) *CreateBasicProjectResponse {
	s.Body = v
	return s
}

func (s *CreateBasicProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

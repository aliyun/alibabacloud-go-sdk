// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelingProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelingProjectResponseBody) *CreateModelingProjectResponse
	GetBody() *CreateModelingProjectResponseBody
}

type CreateModelingProjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelingProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateModelingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelingProjectResponse) GetBody() *CreateModelingProjectResponseBody {
	return s.Body
}

func (s *CreateModelingProjectResponse) SetHeaders(v map[string]*string) *CreateModelingProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateModelingProjectResponse) SetStatusCode(v int32) *CreateModelingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelingProjectResponse) SetBody(v *CreateModelingProjectResponseBody) *CreateModelingProjectResponse {
	s.Body = v
	return s
}

func (s *CreateModelingProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEditingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddEditingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddEditingProjectResponse
	GetStatusCode() *int32
	SetBody(v *AddEditingProjectResponseBody) *AddEditingProjectResponse
	GetBody() *AddEditingProjectResponseBody
}

type AddEditingProjectResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEditingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *AddEditingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddEditingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddEditingProjectResponse) GetBody() *AddEditingProjectResponseBody {
	return s.Body
}

func (s *AddEditingProjectResponse) SetHeaders(v map[string]*string) *AddEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *AddEditingProjectResponse) SetStatusCode(v int32) *AddEditingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEditingProjectResponse) SetBody(v *AddEditingProjectResponseBody) *AddEditingProjectResponse {
	s.Body = v
	return s
}

func (s *AddEditingProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

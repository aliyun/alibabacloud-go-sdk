// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEditingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEditingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEditingProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateEditingProjectResponseBody) *CreateEditingProjectResponse
	GetBody() *CreateEditingProjectResponseBody
}

type CreateEditingProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEditingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEditingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEditingProjectResponse) GetBody() *CreateEditingProjectResponseBody {
	return s.Body
}

func (s *CreateEditingProjectResponse) SetHeaders(v map[string]*string) *CreateEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateEditingProjectResponse) SetStatusCode(v int32) *CreateEditingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEditingProjectResponse) SetBody(v *CreateEditingProjectResponseBody) *CreateEditingProjectResponse {
	s.Body = v
	return s
}

func (s *CreateEditingProjectResponse) Validate() error {
	return dara.Validate(s)
}

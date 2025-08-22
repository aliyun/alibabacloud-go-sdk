// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectResponseBody) *GetProjectResponse
	GetBody() *GetProjectResponseBody
}

type GetProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectResponse) GetBody() *GetProjectResponseBody {
	return s.Body
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

func (s *GetProjectResponse) Validate() error {
	return dara.Validate(s)
}

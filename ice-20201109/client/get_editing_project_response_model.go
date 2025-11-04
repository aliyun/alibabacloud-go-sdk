// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEditingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEditingProjectResponse
	GetStatusCode() *int32
	SetBody(v *GetEditingProjectResponseBody) *GetEditingProjectResponse
	GetBody() *GetEditingProjectResponseBody
}

type GetEditingProjectResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEditingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEditingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEditingProjectResponse) GetBody() *GetEditingProjectResponseBody {
	return s.Body
}

func (s *GetEditingProjectResponse) SetHeaders(v map[string]*string) *GetEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *GetEditingProjectResponse) SetStatusCode(v int32) *GetEditingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEditingProjectResponse) SetBody(v *GetEditingProjectResponseBody) *GetEditingProjectResponse {
	s.Body = v
	return s
}

func (s *GetEditingProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

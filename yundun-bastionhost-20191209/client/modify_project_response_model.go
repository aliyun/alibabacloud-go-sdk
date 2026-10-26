// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyProjectResponse
	GetStatusCode() *int32
	SetBody(v *ModifyProjectResponseBody) *ModifyProjectResponse
	GetBody() *ModifyProjectResponseBody
}

type ModifyProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyProjectResponse) GoString() string {
	return s.String()
}

func (s *ModifyProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyProjectResponse) GetBody() *ModifyProjectResponseBody {
	return s.Body
}

func (s *ModifyProjectResponse) SetHeaders(v map[string]*string) *ModifyProjectResponse {
	s.Headers = v
	return s
}

func (s *ModifyProjectResponse) SetStatusCode(v int32) *ModifyProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProjectResponse) SetBody(v *ModifyProjectResponseBody) *ModifyProjectResponse {
	s.Body = v
	return s
}

func (s *ModifyProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

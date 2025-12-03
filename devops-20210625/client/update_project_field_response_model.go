// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProjectFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProjectFieldResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProjectFieldResponseBody) *UpdateProjectFieldResponse
	GetBody() *UpdateProjectFieldResponseBody
}

type UpdateProjectFieldResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectFieldResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProjectFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProjectFieldResponse) GetBody() *UpdateProjectFieldResponseBody {
	return s.Body
}

func (s *UpdateProjectFieldResponse) SetHeaders(v map[string]*string) *UpdateProjectFieldResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectFieldResponse) SetStatusCode(v int32) *UpdateProjectFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectFieldResponse) SetBody(v *UpdateProjectFieldResponseBody) *UpdateProjectFieldResponse {
	s.Body = v
	return s
}

func (s *UpdateProjectFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

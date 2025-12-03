// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProjectLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProjectLabelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProjectLabelResponseBody) *UpdateProjectLabelResponse
	GetBody() *UpdateProjectLabelResponseBody
}

type UpdateProjectLabelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectLabelResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProjectLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProjectLabelResponse) GetBody() *UpdateProjectLabelResponseBody {
	return s.Body
}

func (s *UpdateProjectLabelResponse) SetHeaders(v map[string]*string) *UpdateProjectLabelResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectLabelResponse) SetStatusCode(v int32) *UpdateProjectLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectLabelResponse) SetBody(v *UpdateProjectLabelResponseBody) *UpdateProjectLabelResponse {
	s.Body = v
	return s
}

func (s *UpdateProjectLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

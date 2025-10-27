// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJobResponseBody) *UpdateJobResponse
	GetBody() *UpdateJobResponseBody
}

type UpdateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJobResponse) GetBody() *UpdateJobResponseBody {
	return s.Body
}

func (s *UpdateJobResponse) SetHeaders(v map[string]*string) *UpdateJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobResponse) SetStatusCode(v int32) *UpdateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobResponse) SetBody(v *UpdateJobResponseBody) *UpdateJobResponse {
	s.Body = v
	return s
}

func (s *UpdateJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

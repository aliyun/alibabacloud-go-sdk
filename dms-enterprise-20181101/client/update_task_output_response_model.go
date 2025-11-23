// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskOutputResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskOutputResponseBody) *UpdateTaskOutputResponse
	GetBody() *UpdateTaskOutputResponseBody
}

type UpdateTaskOutputResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskOutputResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskOutputResponse) GetBody() *UpdateTaskOutputResponseBody {
	return s.Body
}

func (s *UpdateTaskOutputResponse) SetHeaders(v map[string]*string) *UpdateTaskOutputResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskOutputResponse) SetStatusCode(v int32) *UpdateTaskOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskOutputResponse) SetBody(v *UpdateTaskOutputResponseBody) *UpdateTaskOutputResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

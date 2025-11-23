// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskConfigResponseBody) *UpdateTaskConfigResponse
	GetBody() *UpdateTaskConfigResponseBody
}

type UpdateTaskConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskConfigResponse) GetBody() *UpdateTaskConfigResponseBody {
	return s.Body
}

func (s *UpdateTaskConfigResponse) SetHeaders(v map[string]*string) *UpdateTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskConfigResponse) SetStatusCode(v int32) *UpdateTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskConfigResponse) SetBody(v *UpdateTaskConfigResponseBody) *UpdateTaskConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

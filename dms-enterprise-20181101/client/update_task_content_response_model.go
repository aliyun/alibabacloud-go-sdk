// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskContentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskContentResponseBody) *UpdateTaskContentResponse
	GetBody() *UpdateTaskContentResponseBody
}

type UpdateTaskContentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskContentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskContentResponse) GetBody() *UpdateTaskContentResponseBody {
	return s.Body
}

func (s *UpdateTaskContentResponse) SetHeaders(v map[string]*string) *UpdateTaskContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskContentResponse) SetStatusCode(v int32) *UpdateTaskContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskContentResponse) SetBody(v *UpdateTaskContentResponseBody) *UpdateTaskContentResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiCallTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAiCallTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAiCallTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAiCallTaskResponseBody) *UpdateAiCallTaskResponse
	GetBody() *UpdateAiCallTaskResponseBody
}

type UpdateAiCallTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAiCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAiCallTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiCallTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateAiCallTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAiCallTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAiCallTaskResponse) GetBody() *UpdateAiCallTaskResponseBody {
	return s.Body
}

func (s *UpdateAiCallTaskResponse) SetHeaders(v map[string]*string) *UpdateAiCallTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateAiCallTaskResponse) SetStatusCode(v int32) *UpdateAiCallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAiCallTaskResponse) SetBody(v *UpdateAiCallTaskResponseBody) *UpdateAiCallTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateAiCallTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

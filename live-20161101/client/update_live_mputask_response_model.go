// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveMPUTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveMPUTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveMPUTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveMPUTaskResponseBody) *UpdateLiveMPUTaskResponse
	GetBody() *UpdateLiveMPUTaskResponseBody
}

type UpdateLiveMPUTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveMPUTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveMPUTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveMPUTaskResponse) GetBody() *UpdateLiveMPUTaskResponseBody {
	return s.Body
}

func (s *UpdateLiveMPUTaskResponse) SetHeaders(v map[string]*string) *UpdateLiveMPUTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveMPUTaskResponse) SetStatusCode(v int32) *UpdateLiveMPUTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveMPUTaskResponse) SetBody(v *UpdateLiveMPUTaskResponseBody) *UpdateLiveMPUTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveMPUTaskResponse) Validate() error {
	return dara.Validate(s)
}

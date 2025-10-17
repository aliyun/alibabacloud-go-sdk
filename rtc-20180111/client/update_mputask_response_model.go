// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMPUTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMPUTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMPUTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMPUTaskResponseBody) *UpdateMPUTaskResponse
	GetBody() *UpdateMPUTaskResponseBody
}

type UpdateMPUTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMPUTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMPUTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMPUTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMPUTaskResponse) GetBody() *UpdateMPUTaskResponseBody {
	return s.Body
}

func (s *UpdateMPUTaskResponse) SetHeaders(v map[string]*string) *UpdateMPUTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateMPUTaskResponse) SetStatusCode(v int32) *UpdateMPUTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMPUTaskResponse) SetBody(v *UpdateMPUTaskResponseBody) *UpdateMPUTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateMPUTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

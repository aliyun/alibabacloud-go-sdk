// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFotaTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFotaTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFotaTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFotaTaskResponseBody) *UpdateFotaTaskResponse
	GetBody() *UpdateFotaTaskResponseBody
}

type UpdateFotaTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFotaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFotaTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFotaTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateFotaTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFotaTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFotaTaskResponse) GetBody() *UpdateFotaTaskResponseBody {
	return s.Body
}

func (s *UpdateFotaTaskResponse) SetHeaders(v map[string]*string) *UpdateFotaTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateFotaTaskResponse) SetStatusCode(v int32) *UpdateFotaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFotaTaskResponse) SetBody(v *UpdateFotaTaskResponseBody) *UpdateFotaTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateFotaTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

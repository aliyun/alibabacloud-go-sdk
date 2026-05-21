// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTempFileTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTempFileTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTempFileTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTempFileTaskResponseBody) *UpdateTempFileTaskResponse
	GetBody() *UpdateTempFileTaskResponseBody
}

type UpdateTempFileTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTempFileTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTempFileTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTempFileTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTempFileTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTempFileTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTempFileTaskResponse) GetBody() *UpdateTempFileTaskResponseBody {
	return s.Body
}

func (s *UpdateTempFileTaskResponse) SetHeaders(v map[string]*string) *UpdateTempFileTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTempFileTaskResponse) SetStatusCode(v int32) *UpdateTempFileTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTempFileTaskResponse) SetBody(v *UpdateTempFileTaskResponseBody) *UpdateTempFileTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateTempFileTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

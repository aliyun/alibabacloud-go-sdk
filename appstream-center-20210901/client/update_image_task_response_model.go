// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateImageTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateImageTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateImageTaskResponseBody) *UpdateImageTaskResponse
	GetBody() *UpdateImageTaskResponseBody
}

type UpdateImageTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateImageTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateImageTaskResponse) GetBody() *UpdateImageTaskResponseBody {
	return s.Body
}

func (s *UpdateImageTaskResponse) SetHeaders(v map[string]*string) *UpdateImageTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageTaskResponse) SetStatusCode(v int32) *UpdateImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageTaskResponse) SetBody(v *UpdateImageTaskResponseBody) *UpdateImageTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateImageTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

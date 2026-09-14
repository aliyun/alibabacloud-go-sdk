// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskAsyncResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskAsyncResponseBody) *UpdateTaskAsyncResponse
	GetBody() *UpdateTaskAsyncResponseBody
}

type UpdateTaskAsyncResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskAsyncResponse) GetBody() *UpdateTaskAsyncResponseBody {
	return s.Body
}

func (s *UpdateTaskAsyncResponse) SetHeaders(v map[string]*string) *UpdateTaskAsyncResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskAsyncResponse) SetStatusCode(v int32) *UpdateTaskAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskAsyncResponse) SetBody(v *UpdateTaskAsyncResponseBody) *UpdateTaskAsyncResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

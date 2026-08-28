// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRollbackTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRollbackTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRollbackTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRollbackTaskResponseBody) *CreateRollbackTaskResponse
	GetBody() *CreateRollbackTaskResponseBody
}

type CreateRollbackTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRollbackTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRollbackTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRollbackTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRollbackTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRollbackTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRollbackTaskResponse) GetBody() *CreateRollbackTaskResponseBody {
	return s.Body
}

func (s *CreateRollbackTaskResponse) SetHeaders(v map[string]*string) *CreateRollbackTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRollbackTaskResponse) SetStatusCode(v int32) *CreateRollbackTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRollbackTaskResponse) SetBody(v *CreateRollbackTaskResponseBody) *CreateRollbackTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRollbackTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComputeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComputeTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateComputeTaskResponseBody) *CreateComputeTaskResponse
	GetBody() *CreateComputeTaskResponseBody
}

type CreateComputeTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComputeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComputeTaskResponse) GetBody() *CreateComputeTaskResponseBody {
	return s.Body
}

func (s *CreateComputeTaskResponse) SetHeaders(v map[string]*string) *CreateComputeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeTaskResponse) SetStatusCode(v int32) *CreateComputeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeTaskResponse) SetBody(v *CreateComputeTaskResponseBody) *CreateComputeTaskResponse {
	s.Body = v
	return s
}

func (s *CreateComputeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComfyTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComfyTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComfyTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateComfyTaskResponseBody) *CreateComfyTaskResponse
	GetBody() *CreateComfyTaskResponseBody
}

type CreateComfyTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComfyTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComfyTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComfyTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateComfyTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComfyTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComfyTaskResponse) GetBody() *CreateComfyTaskResponseBody {
	return s.Body
}

func (s *CreateComfyTaskResponse) SetHeaders(v map[string]*string) *CreateComfyTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateComfyTaskResponse) SetStatusCode(v int32) *CreateComfyTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComfyTaskResponse) SetBody(v *CreateComfyTaskResponseBody) *CreateComfyTaskResponse {
	s.Body = v
	return s
}

func (s *CreateComfyTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

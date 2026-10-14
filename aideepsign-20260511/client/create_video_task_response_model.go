// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVideoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVideoTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateVideoTaskResponseBody) *CreateVideoTaskResponse
	GetBody() *CreateVideoTaskResponseBody
}

type CreateVideoTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVideoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVideoTaskResponse) GetBody() *CreateVideoTaskResponseBody {
	return s.Body
}

func (s *CreateVideoTaskResponse) SetHeaders(v map[string]*string) *CreateVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoTaskResponse) SetStatusCode(v int32) *CreateVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoTaskResponse) SetBody(v *CreateVideoTaskResponseBody) *CreateVideoTaskResponse {
	s.Body = v
	return s
}

func (s *CreateVideoTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

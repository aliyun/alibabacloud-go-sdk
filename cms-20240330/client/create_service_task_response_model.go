// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceTaskResponseBody) *CreateServiceTaskResponse
	GetBody() *CreateServiceTaskResponseBody
}

type CreateServiceTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceTaskResponse) GetBody() *CreateServiceTaskResponseBody {
	return s.Body
}

func (s *CreateServiceTaskResponse) SetHeaders(v map[string]*string) *CreateServiceTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceTaskResponse) SetStatusCode(v int32) *CreateServiceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceTaskResponse) SetBody(v *CreateServiceTaskResponseBody) *CreateServiceTaskResponse {
	s.Body = v
	return s
}

func (s *CreateServiceTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

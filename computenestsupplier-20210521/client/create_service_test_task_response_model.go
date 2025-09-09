// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTestTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceTestTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceTestTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceTestTaskResponseBody) *CreateServiceTestTaskResponse
	GetBody() *CreateServiceTestTaskResponseBody
}

type CreateServiceTestTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceTestTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceTestTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTestTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceTestTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceTestTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceTestTaskResponse) GetBody() *CreateServiceTestTaskResponseBody {
	return s.Body
}

func (s *CreateServiceTestTaskResponse) SetHeaders(v map[string]*string) *CreateServiceTestTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceTestTaskResponse) SetStatusCode(v int32) *CreateServiceTestTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceTestTaskResponse) SetBody(v *CreateServiceTestTaskResponseBody) *CreateServiceTestTaskResponse {
	s.Body = v
	return s
}

func (s *CreateServiceTestTaskResponse) Validate() error {
	return dara.Validate(s)
}

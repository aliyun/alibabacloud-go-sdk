// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskAttributeResponseBody) *UpdateTaskAttributeResponse
	GetBody() *UpdateTaskAttributeResponseBody
}

type UpdateTaskAttributeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskAttributeResponse) GetBody() *UpdateTaskAttributeResponseBody {
	return s.Body
}

func (s *UpdateTaskAttributeResponse) SetHeaders(v map[string]*string) *UpdateTaskAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskAttributeResponse) SetStatusCode(v int32) *UpdateTaskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskAttributeResponse) SetBody(v *UpdateTaskAttributeResponseBody) *UpdateTaskAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskAttributeResponse) Validate() error {
	return dara.Validate(s)
}

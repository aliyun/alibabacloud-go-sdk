// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelResponseBody) *UpdateModelResponse
	GetBody() *UpdateModelResponseBody
}

type UpdateModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelResponse) GetBody() *UpdateModelResponseBody {
	return s.Body
}

func (s *UpdateModelResponse) SetHeaders(v map[string]*string) *UpdateModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelResponse) SetStatusCode(v int32) *UpdateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelResponse) SetBody(v *UpdateModelResponseBody) *UpdateModelResponse {
	s.Body = v
	return s
}

func (s *UpdateModelResponse) Validate() error {
	return dara.Validate(s)
}

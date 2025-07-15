// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExecutionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExecutionResponseBody) *UpdateExecutionResponse
	GetBody() *UpdateExecutionResponseBody
}

type UpdateExecutionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutionResponse) GoString() string {
	return s.String()
}

func (s *UpdateExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExecutionResponse) GetBody() *UpdateExecutionResponseBody {
	return s.Body
}

func (s *UpdateExecutionResponse) SetHeaders(v map[string]*string) *UpdateExecutionResponse {
	s.Headers = v
	return s
}

func (s *UpdateExecutionResponse) SetStatusCode(v int32) *UpdateExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExecutionResponse) SetBody(v *UpdateExecutionResponseBody) *UpdateExecutionResponse {
	s.Body = v
	return s
}

func (s *UpdateExecutionResponse) Validate() error {
	return dara.Validate(s)
}

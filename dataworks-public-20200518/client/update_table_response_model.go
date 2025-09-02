// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTableResponseBody) *UpdateTableResponse
	GetBody() *UpdateTableResponseBody
}

type UpdateTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTableResponse) GetBody() *UpdateTableResponseBody {
	return s.Body
}

func (s *UpdateTableResponse) SetHeaders(v map[string]*string) *UpdateTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableResponse) SetStatusCode(v int32) *UpdateTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableResponse) SetBody(v *UpdateTableResponseBody) *UpdateTableResponse {
	s.Body = v
	return s
}

func (s *UpdateTableResponse) Validate() error {
	return dara.Validate(s)
}

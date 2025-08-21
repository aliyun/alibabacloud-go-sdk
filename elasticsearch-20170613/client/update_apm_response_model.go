// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApmResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApmResponseBody) *UpdateApmResponse
	GetBody() *UpdateApmResponseBody
}

type UpdateApmResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApmResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApmResponse) GoString() string {
	return s.String()
}

func (s *UpdateApmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApmResponse) GetBody() *UpdateApmResponseBody {
	return s.Body
}

func (s *UpdateApmResponse) SetHeaders(v map[string]*string) *UpdateApmResponse {
	s.Headers = v
	return s
}

func (s *UpdateApmResponse) SetStatusCode(v int32) *UpdateApmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApmResponse) SetBody(v *UpdateApmResponseBody) *UpdateApmResponse {
	s.Body = v
	return s
}

func (s *UpdateApmResponse) Validate() error {
	return dara.Validate(s)
}

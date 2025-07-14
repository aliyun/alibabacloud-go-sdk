// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIngressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIngressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIngressResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIngressResponseBody) *UpdateIngressResponse
	GetBody() *UpdateIngressResponseBody
}

type UpdateIngressResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIngressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIngressResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIngressResponse) GoString() string {
	return s.String()
}

func (s *UpdateIngressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIngressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIngressResponse) GetBody() *UpdateIngressResponseBody {
	return s.Body
}

func (s *UpdateIngressResponse) SetHeaders(v map[string]*string) *UpdateIngressResponse {
	s.Headers = v
	return s
}

func (s *UpdateIngressResponse) SetStatusCode(v int32) *UpdateIngressResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIngressResponse) SetBody(v *UpdateIngressResponseBody) *UpdateIngressResponse {
	s.Body = v
	return s
}

func (s *UpdateIngressResponse) Validate() error {
	return dara.Validate(s)
}

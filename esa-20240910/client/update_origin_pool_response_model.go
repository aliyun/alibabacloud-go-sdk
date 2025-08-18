// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOriginPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOriginPoolResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOriginPoolResponseBody) *UpdateOriginPoolResponse
	GetBody() *UpdateOriginPoolResponseBody
}

type UpdateOriginPoolResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOriginPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOriginPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginPoolResponse) GoString() string {
	return s.String()
}

func (s *UpdateOriginPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOriginPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOriginPoolResponse) GetBody() *UpdateOriginPoolResponseBody {
	return s.Body
}

func (s *UpdateOriginPoolResponse) SetHeaders(v map[string]*string) *UpdateOriginPoolResponse {
	s.Headers = v
	return s
}

func (s *UpdateOriginPoolResponse) SetStatusCode(v int32) *UpdateOriginPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOriginPoolResponse) SetBody(v *UpdateOriginPoolResponseBody) *UpdateOriginPoolResponse {
	s.Body = v
	return s
}

func (s *UpdateOriginPoolResponse) Validate() error {
	return dara.Validate(s)
}

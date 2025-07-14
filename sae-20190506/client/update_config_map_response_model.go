// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConfigMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConfigMapResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConfigMapResponseBody) *UpdateConfigMapResponse
	GetBody() *UpdateConfigMapResponseBody
}

type UpdateConfigMapResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConfigMapResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigMapResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConfigMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConfigMapResponse) GetBody() *UpdateConfigMapResponseBody {
	return s.Body
}

func (s *UpdateConfigMapResponse) SetHeaders(v map[string]*string) *UpdateConfigMapResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigMapResponse) SetStatusCode(v int32) *UpdateConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigMapResponse) SetBody(v *UpdateConfigMapResponseBody) *UpdateConfigMapResponse {
	s.Body = v
	return s
}

func (s *UpdateConfigMapResponse) Validate() error {
	return dara.Validate(s)
}

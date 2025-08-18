// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppInstanceGroupImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppInstanceGroupImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppInstanceGroupImageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppInstanceGroupImageResponseBody) *UpdateAppInstanceGroupImageResponse
	GetBody() *UpdateAppInstanceGroupImageResponseBody
}

type UpdateAppInstanceGroupImageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppInstanceGroupImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppInstanceGroupImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInstanceGroupImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppInstanceGroupImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppInstanceGroupImageResponse) GetBody() *UpdateAppInstanceGroupImageResponseBody {
	return s.Body
}

func (s *UpdateAppInstanceGroupImageResponse) SetHeaders(v map[string]*string) *UpdateAppInstanceGroupImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppInstanceGroupImageResponse) SetStatusCode(v int32) *UpdateAppInstanceGroupImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponse) SetBody(v *UpdateAppInstanceGroupImageResponseBody) *UpdateAppInstanceGroupImageResponse {
	s.Body = v
	return s
}

func (s *UpdateAppInstanceGroupImageResponse) Validate() error {
	return dara.Validate(s)
}

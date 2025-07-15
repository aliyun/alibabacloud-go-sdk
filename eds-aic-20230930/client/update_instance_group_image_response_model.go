// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceGroupImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceGroupImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceGroupImageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceGroupImageResponseBody) *UpdateInstanceGroupImageResponse
	GetBody() *UpdateInstanceGroupImageResponseBody
}

type UpdateInstanceGroupImageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceGroupImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceGroupImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceGroupImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceGroupImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceGroupImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceGroupImageResponse) GetBody() *UpdateInstanceGroupImageResponseBody {
	return s.Body
}

func (s *UpdateInstanceGroupImageResponse) SetHeaders(v map[string]*string) *UpdateInstanceGroupImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceGroupImageResponse) SetStatusCode(v int32) *UpdateInstanceGroupImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceGroupImageResponse) SetBody(v *UpdateInstanceGroupImageResponseBody) *UpdateInstanceGroupImageResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceGroupImageResponse) Validate() error {
	return dara.Validate(s)
}

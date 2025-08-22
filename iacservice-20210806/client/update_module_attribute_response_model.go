// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModuleAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModuleAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModuleAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModuleAttributeResponseBody) *UpdateModuleAttributeResponse
	GetBody() *UpdateModuleAttributeResponseBody
}

type UpdateModuleAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModuleAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateModuleAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModuleAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModuleAttributeResponse) GetBody() *UpdateModuleAttributeResponseBody {
	return s.Body
}

func (s *UpdateModuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateModuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateModuleAttributeResponse) SetStatusCode(v int32) *UpdateModuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModuleAttributeResponse) SetBody(v *UpdateModuleAttributeResponseBody) *UpdateModuleAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateModuleAttributeResponse) Validate() error {
	return dara.Validate(s)
}

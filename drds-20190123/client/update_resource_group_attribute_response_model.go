// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceGroupAttributeResponseBody) *UpdateResourceGroupAttributeResponse
	GetBody() *UpdateResourceGroupAttributeResponseBody
}

type UpdateResourceGroupAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceGroupAttributeResponse) GetBody() *UpdateResourceGroupAttributeResponseBody {
	return s.Body
}

func (s *UpdateResourceGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateResourceGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceGroupAttributeResponse) SetStatusCode(v int32) *UpdateResourceGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceGroupAttributeResponse) SetBody(v *UpdateResourceGroupAttributeResponseBody) *UpdateResourceGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceGroupAttributeResponse) Validate() error {
	return dara.Validate(s)
}

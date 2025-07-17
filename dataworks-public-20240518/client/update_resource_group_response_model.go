// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceGroupResponseBody) *UpdateResourceGroupResponse
	GetBody() *UpdateResourceGroupResponseBody
}

type UpdateResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceGroupResponse) GetBody() *UpdateResourceGroupResponseBody {
	return s.Body
}

func (s *UpdateResourceGroupResponse) SetHeaders(v map[string]*string) *UpdateResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceGroupResponse) SetStatusCode(v int32) *UpdateResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceGroupResponse) SetBody(v *UpdateResourceGroupResponseBody) *UpdateResourceGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}

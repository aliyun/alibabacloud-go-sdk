// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCasterResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCasterResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCasterResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCasterResourceGroupResponseBody) *UpdateCasterResourceGroupResponse
	GetBody() *UpdateCasterResourceGroupResponseBody
}

type UpdateCasterResourceGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCasterResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCasterResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCasterResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateCasterResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCasterResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCasterResourceGroupResponse) GetBody() *UpdateCasterResourceGroupResponseBody {
	return s.Body
}

func (s *UpdateCasterResourceGroupResponse) SetHeaders(v map[string]*string) *UpdateCasterResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateCasterResourceGroupResponse) SetStatusCode(v int32) *UpdateCasterResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCasterResourceGroupResponse) SetBody(v *UpdateCasterResourceGroupResponseBody) *UpdateCasterResourceGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateCasterResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}

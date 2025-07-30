// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyResourceGroupResponseBody) *ModifyResourceGroupResponse
	GetBody() *ModifyResourceGroupResponseBody
}

type ModifyResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyResourceGroupResponse) GetBody() *ModifyResourceGroupResponseBody {
	return s.Body
}

func (s *ModifyResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceGroupResponse) SetStatusCode(v int32) *ModifyResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResourceGroupResponse) SetBody(v *ModifyResourceGroupResponseBody) *ModifyResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}

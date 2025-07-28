// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefenseResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefenseResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefenseResourceGroupResponseBody) *ModifyDefenseResourceGroupResponse
	GetBody() *ModifyDefenseResourceGroupResponseBody
}

type ModifyDefenseResourceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefenseResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefenseResourceGroupResponse) GetBody() *ModifyDefenseResourceGroupResponseBody {
	return s.Body
}

func (s *ModifyDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseResourceGroupResponse) SetStatusCode(v int32) *ModifyDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseResourceGroupResponse) SetBody(v *ModifyDefenseResourceGroupResponseBody) *ModifyDefenseResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyDefenseResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}

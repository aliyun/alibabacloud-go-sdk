// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPolicyGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPolicyGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPolicyGroupResponseBody) *ModifyPolicyGroupResponse
	GetBody() *ModifyPolicyGroupResponseBody
}

type ModifyPolicyGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolicyGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPolicyGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPolicyGroupResponse) GetBody() *ModifyPolicyGroupResponseBody {
	return s.Body
}

func (s *ModifyPolicyGroupResponse) SetHeaders(v map[string]*string) *ModifyPolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyGroupResponse) SetStatusCode(v int32) *ModifyPolicyGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyGroupResponse) SetBody(v *ModifyPolicyGroupResponseBody) *ModifyPolicyGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyPolicyGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

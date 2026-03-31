// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyFromRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachPolicyFromRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachPolicyFromRoleResponse
	GetStatusCode() *int32
	SetBody(v *DetachPolicyFromRoleResponseBody) *DetachPolicyFromRoleResponse
	GetBody() *DetachPolicyFromRoleResponseBody
}

type DetachPolicyFromRoleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicyFromRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicyFromRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyFromRoleResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachPolicyFromRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachPolicyFromRoleResponse) GetBody() *DetachPolicyFromRoleResponseBody {
	return s.Body
}

func (s *DetachPolicyFromRoleResponse) SetHeaders(v map[string]*string) *DetachPolicyFromRoleResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicyFromRoleResponse) SetStatusCode(v int32) *DetachPolicyFromRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicyFromRoleResponse) SetBody(v *DetachPolicyFromRoleResponseBody) *DetachPolicyFromRoleResponse {
	s.Body = v
	return s
}

func (s *DetachPolicyFromRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

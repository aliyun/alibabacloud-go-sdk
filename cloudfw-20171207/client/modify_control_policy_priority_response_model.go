// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlPolicyPriorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyControlPolicyPriorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyControlPolicyPriorityResponse
	GetStatusCode() *int32
	SetBody(v *ModifyControlPolicyPriorityResponseBody) *ModifyControlPolicyPriorityResponse
	GetBody() *ModifyControlPolicyPriorityResponseBody
}

type ModifyControlPolicyPriorityResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyControlPolicyPriorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyControlPolicyPriorityResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlPolicyPriorityResponse) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPriorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyControlPolicyPriorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyControlPolicyPriorityResponse) GetBody() *ModifyControlPolicyPriorityResponseBody {
	return s.Body
}

func (s *ModifyControlPolicyPriorityResponse) SetHeaders(v map[string]*string) *ModifyControlPolicyPriorityResponse {
	s.Headers = v
	return s
}

func (s *ModifyControlPolicyPriorityResponse) SetStatusCode(v int32) *ModifyControlPolicyPriorityResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyControlPolicyPriorityResponse) SetBody(v *ModifyControlPolicyPriorityResponseBody) *ModifyControlPolicyPriorityResponse {
	s.Body = v
	return s
}

func (s *ModifyControlPolicyPriorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

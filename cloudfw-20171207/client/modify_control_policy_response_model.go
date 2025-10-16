// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyControlPolicyResponseBody) *ModifyControlPolicyResponse
	GetBody() *ModifyControlPolicyResponseBody
}

type ModifyControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyControlPolicyResponse) GetBody() *ModifyControlPolicyResponseBody {
	return s.Body
}

func (s *ModifyControlPolicyResponse) SetHeaders(v map[string]*string) *ModifyControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyControlPolicyResponse) SetStatusCode(v int32) *ModifyControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyControlPolicyResponse) SetBody(v *ModifyControlPolicyResponseBody) *ModifyControlPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

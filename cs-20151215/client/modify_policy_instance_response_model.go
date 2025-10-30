// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPolicyInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPolicyInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPolicyInstanceResponseBody) *ModifyPolicyInstanceResponse
	GetBody() *ModifyPolicyInstanceResponseBody
}

type ModifyPolicyInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolicyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolicyInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPolicyInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPolicyInstanceResponse) GetBody() *ModifyPolicyInstanceResponseBody {
	return s.Body
}

func (s *ModifyPolicyInstanceResponse) SetHeaders(v map[string]*string) *ModifyPolicyInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyInstanceResponse) SetStatusCode(v int32) *ModifyPolicyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyInstanceResponse) SetBody(v *ModifyPolicyInstanceResponseBody) *ModifyPolicyInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyPolicyInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

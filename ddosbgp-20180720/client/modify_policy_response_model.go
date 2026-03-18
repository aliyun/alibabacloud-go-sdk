// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPolicyResponseBody) *ModifyPolicyResponse
	GetBody() *ModifyPolicyResponseBody
}

type ModifyPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPolicyResponse) GetBody() *ModifyPolicyResponseBody {
	return s.Body
}

func (s *ModifyPolicyResponse) SetHeaders(v map[string]*string) *ModifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyResponse) SetStatusCode(v int32) *ModifyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyResponse) SetBody(v *ModifyPolicyResponseBody) *ModifyPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

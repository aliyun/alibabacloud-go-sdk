// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenterPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCenterPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCenterPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCenterPolicyResponseBody) *ModifyCenterPolicyResponse
	GetBody() *ModifyCenterPolicyResponseBody
}

type ModifyCenterPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCenterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCenterPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenterPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyCenterPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCenterPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCenterPolicyResponse) GetBody() *ModifyCenterPolicyResponseBody {
	return s.Body
}

func (s *ModifyCenterPolicyResponse) SetHeaders(v map[string]*string) *ModifyCenterPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyCenterPolicyResponse) SetStatusCode(v int32) *ModifyCenterPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCenterPolicyResponse) SetBody(v *ModifyCenterPolicyResponseBody) *ModifyCenterPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyCenterPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

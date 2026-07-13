// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtectionPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProtectionPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProtectionPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProtectionPolicyResponseBody) *DeleteProtectionPolicyResponse
	GetBody() *DeleteProtectionPolicyResponseBody
}

type DeleteProtectionPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProtectionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProtectionPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtectionPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteProtectionPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProtectionPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProtectionPolicyResponse) GetBody() *DeleteProtectionPolicyResponseBody {
	return s.Body
}

func (s *DeleteProtectionPolicyResponse) SetHeaders(v map[string]*string) *DeleteProtectionPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteProtectionPolicyResponse) SetStatusCode(v int32) *DeleteProtectionPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProtectionPolicyResponse) SetBody(v *DeleteProtectionPolicyResponseBody) *DeleteProtectionPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteProtectionPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppPolicyResponseBody) *ModifyAppPolicyResponse
	GetBody() *ModifyAppPolicyResponseBody
}

type ModifyAppPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppPolicyResponse) GetBody() *ModifyAppPolicyResponseBody {
	return s.Body
}

func (s *ModifyAppPolicyResponse) SetHeaders(v map[string]*string) *ModifyAppPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppPolicyResponse) SetStatusCode(v int32) *ModifyAppPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppPolicyResponse) SetBody(v *ModifyAppPolicyResponseBody) *ModifyAppPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyAppPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountCheckPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountCheckPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountCheckPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountCheckPolicyResponseBody) *ModifyAccountCheckPolicyResponse
	GetBody() *ModifyAccountCheckPolicyResponseBody
}

type ModifyAccountCheckPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountCheckPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountCheckPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountCheckPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountCheckPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountCheckPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountCheckPolicyResponse) GetBody() *ModifyAccountCheckPolicyResponseBody {
	return s.Body
}

func (s *ModifyAccountCheckPolicyResponse) SetHeaders(v map[string]*string) *ModifyAccountCheckPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountCheckPolicyResponse) SetStatusCode(v int32) *ModifyAccountCheckPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountCheckPolicyResponse) SetBody(v *ModifyAccountCheckPolicyResponseBody) *ModifyAccountCheckPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountCheckPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

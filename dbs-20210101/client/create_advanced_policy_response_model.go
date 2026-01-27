// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAdvancedPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAdvancedPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAdvancedPolicyResponseBody) *CreateAdvancedPolicyResponse
	GetBody() *CreateAdvancedPolicyResponseBody
}

type CreateAdvancedPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAdvancedPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAdvancedPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAdvancedPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAdvancedPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAdvancedPolicyResponse) GetBody() *CreateAdvancedPolicyResponseBody {
	return s.Body
}

func (s *CreateAdvancedPolicyResponse) SetHeaders(v map[string]*string) *CreateAdvancedPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAdvancedPolicyResponse) SetStatusCode(v int32) *CreateAdvancedPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAdvancedPolicyResponse) SetBody(v *CreateAdvancedPolicyResponseBody) *CreateAdvancedPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateAdvancedPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

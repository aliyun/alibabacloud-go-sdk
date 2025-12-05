// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyIPAclConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPolicyIPAclConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPolicyIPAclConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetPolicyIPAclConfigResponseBody) *SetPolicyIPAclConfigResponse
	GetBody() *SetPolicyIPAclConfigResponseBody
}

type SetPolicyIPAclConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPolicyIPAclConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPolicyIPAclConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyIPAclConfigResponse) GoString() string {
	return s.String()
}

func (s *SetPolicyIPAclConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPolicyIPAclConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPolicyIPAclConfigResponse) GetBody() *SetPolicyIPAclConfigResponseBody {
	return s.Body
}

func (s *SetPolicyIPAclConfigResponse) SetHeaders(v map[string]*string) *SetPolicyIPAclConfigResponse {
	s.Headers = v
	return s
}

func (s *SetPolicyIPAclConfigResponse) SetStatusCode(v int32) *SetPolicyIPAclConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPolicyIPAclConfigResponse) SetBody(v *SetPolicyIPAclConfigResponseBody) *SetPolicyIPAclConfigResponse {
	s.Body = v
	return s
}

func (s *SetPolicyIPAclConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

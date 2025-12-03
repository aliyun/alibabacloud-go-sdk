// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIntranetDomainPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIntranetDomainPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIntranetDomainPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIntranetDomainPolicyResponseBody) *ModifyIntranetDomainPolicyResponse
	GetBody() *ModifyIntranetDomainPolicyResponseBody
}

type ModifyIntranetDomainPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIntranetDomainPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIntranetDomainPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIntranetDomainPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyIntranetDomainPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIntranetDomainPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIntranetDomainPolicyResponse) GetBody() *ModifyIntranetDomainPolicyResponseBody {
	return s.Body
}

func (s *ModifyIntranetDomainPolicyResponse) SetHeaders(v map[string]*string) *ModifyIntranetDomainPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyIntranetDomainPolicyResponse) SetStatusCode(v int32) *ModifyIntranetDomainPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIntranetDomainPolicyResponse) SetBody(v *ModifyIntranetDomainPolicyResponseBody) *ModifyIntranetDomainPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyIntranetDomainPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

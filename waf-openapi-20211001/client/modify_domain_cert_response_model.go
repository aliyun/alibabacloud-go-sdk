// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDomainCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDomainCertResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDomainCertResponseBody) *ModifyDomainCertResponse
	GetBody() *ModifyDomainCertResponseBody
}

type ModifyDomainCertResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDomainCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDomainCertResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainCertResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDomainCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDomainCertResponse) GetBody() *ModifyDomainCertResponseBody {
	return s.Body
}

func (s *ModifyDomainCertResponse) SetHeaders(v map[string]*string) *ModifyDomainCertResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainCertResponse) SetStatusCode(v int32) *ModifyDomainCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDomainCertResponse) SetBody(v *ModifyDomainCertResponseBody) *ModifyDomainCertResponse {
	s.Body = v
	return s
}

func (s *ModifyDomainCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

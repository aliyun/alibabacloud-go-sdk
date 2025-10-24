// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyDomainOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyDomainOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyDomainOwnerResponse
	GetStatusCode() *int32
	SetBody(v *VerifyDomainOwnerResponseBody) *VerifyDomainOwnerResponse
	GetBody() *VerifyDomainOwnerResponseBody
}

type VerifyDomainOwnerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyDomainOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyDomainOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyDomainOwnerResponse) GetBody() *VerifyDomainOwnerResponseBody {
	return s.Body
}

func (s *VerifyDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyDomainOwnerResponse) SetStatusCode(v int32) *VerifyDomainOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyDomainOwnerResponse) SetBody(v *VerifyDomainOwnerResponseBody) *VerifyDomainOwnerResponse {
	s.Body = v
	return s
}

func (s *VerifyDomainOwnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

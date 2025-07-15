// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyLiveDomainOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyLiveDomainOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyLiveDomainOwnerResponse
	GetStatusCode() *int32
	SetBody(v *VerifyLiveDomainOwnerResponseBody) *VerifyLiveDomainOwnerResponse
	GetBody() *VerifyLiveDomainOwnerResponseBody
}

type VerifyLiveDomainOwnerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyLiveDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyLiveDomainOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyLiveDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyLiveDomainOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyLiveDomainOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyLiveDomainOwnerResponse) GetBody() *VerifyLiveDomainOwnerResponseBody {
	return s.Body
}

func (s *VerifyLiveDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyLiveDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyLiveDomainOwnerResponse) SetStatusCode(v int32) *VerifyLiveDomainOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyLiveDomainOwnerResponse) SetBody(v *VerifyLiveDomainOwnerResponseBody) *VerifyLiveDomainOwnerResponse {
	s.Body = v
	return s
}

func (s *VerifyLiveDomainOwnerResponse) Validate() error {
	return dara.Validate(s)
}

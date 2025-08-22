// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyDcdnDomainOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyDcdnDomainOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyDcdnDomainOwnerResponse
	GetStatusCode() *int32
	SetBody(v *VerifyDcdnDomainOwnerResponseBody) *VerifyDcdnDomainOwnerResponse
	GetBody() *VerifyDcdnDomainOwnerResponseBody
}

type VerifyDcdnDomainOwnerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyDcdnDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyDcdnDomainOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyDcdnDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyDcdnDomainOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyDcdnDomainOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyDcdnDomainOwnerResponse) GetBody() *VerifyDcdnDomainOwnerResponseBody {
	return s.Body
}

func (s *VerifyDcdnDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyDcdnDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyDcdnDomainOwnerResponse) SetStatusCode(v int32) *VerifyDcdnDomainOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyDcdnDomainOwnerResponse) SetBody(v *VerifyDcdnDomainOwnerResponseBody) *VerifyDcdnDomainOwnerResponse {
	s.Body = v
	return s
}

func (s *VerifyDcdnDomainOwnerResponse) Validate() error {
	return dara.Validate(s)
}

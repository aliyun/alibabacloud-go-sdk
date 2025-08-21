// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyVsDomainOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyVsDomainOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyVsDomainOwnerResponse
	GetStatusCode() *int32
	SetBody(v *VerifyVsDomainOwnerResponseBody) *VerifyVsDomainOwnerResponse
	GetBody() *VerifyVsDomainOwnerResponseBody
}

type VerifyVsDomainOwnerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyVsDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyVsDomainOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyVsDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyVsDomainOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyVsDomainOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyVsDomainOwnerResponse) GetBody() *VerifyVsDomainOwnerResponseBody {
	return s.Body
}

func (s *VerifyVsDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyVsDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyVsDomainOwnerResponse) SetStatusCode(v int32) *VerifyVsDomainOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyVsDomainOwnerResponse) SetBody(v *VerifyVsDomainOwnerResponseBody) *VerifyVsDomainOwnerResponse {
	s.Body = v
	return s
}

func (s *VerifyVsDomainOwnerResponse) Validate() error {
	return dara.Validate(s)
}

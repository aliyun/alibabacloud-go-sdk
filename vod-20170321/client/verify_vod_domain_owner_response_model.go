// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyVodDomainOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyVodDomainOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyVodDomainOwnerResponse
	GetStatusCode() *int32
	SetBody(v *VerifyVodDomainOwnerResponseBody) *VerifyVodDomainOwnerResponse
	GetBody() *VerifyVodDomainOwnerResponseBody
}

type VerifyVodDomainOwnerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyVodDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyVodDomainOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyVodDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyVodDomainOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyVodDomainOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyVodDomainOwnerResponse) GetBody() *VerifyVodDomainOwnerResponseBody {
	return s.Body
}

func (s *VerifyVodDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyVodDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyVodDomainOwnerResponse) SetStatusCode(v int32) *VerifyVodDomainOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyVodDomainOwnerResponse) SetBody(v *VerifyVodDomainOwnerResponseBody) *VerifyVodDomainOwnerResponse {
	s.Body = v
	return s
}

func (s *VerifyVodDomainOwnerResponse) Validate() error {
	return dara.Validate(s)
}

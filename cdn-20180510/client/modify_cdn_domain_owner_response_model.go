// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnDomainOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCdnDomainOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCdnDomainOwnerResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCdnDomainOwnerResponseBody) *ModifyCdnDomainOwnerResponse
	GetBody() *ModifyCdnDomainOwnerResponseBody
}

type ModifyCdnDomainOwnerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCdnDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCdnDomainOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCdnDomainOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCdnDomainOwnerResponse) GetBody() *ModifyCdnDomainOwnerResponseBody {
	return s.Body
}

func (s *ModifyCdnDomainOwnerResponse) SetHeaders(v map[string]*string) *ModifyCdnDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *ModifyCdnDomainOwnerResponse) SetStatusCode(v int32) *ModifyCdnDomainOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCdnDomainOwnerResponse) SetBody(v *ModifyCdnDomainOwnerResponseBody) *ModifyCdnDomainOwnerResponse {
	s.Body = v
	return s
}

func (s *ModifyCdnDomainOwnerResponse) Validate() error {
	return dara.Validate(s)
}

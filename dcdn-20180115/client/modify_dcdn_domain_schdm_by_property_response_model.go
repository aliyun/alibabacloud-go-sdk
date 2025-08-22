// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDCdnDomainSchdmByPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDCdnDomainSchdmByPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDCdnDomainSchdmByPropertyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDCdnDomainSchdmByPropertyResponseBody) *ModifyDCdnDomainSchdmByPropertyResponse
	GetBody() *ModifyDCdnDomainSchdmByPropertyResponseBody
}

type ModifyDCdnDomainSchdmByPropertyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDCdnDomainSchdmByPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDCdnDomainSchdmByPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDCdnDomainSchdmByPropertyResponse) GoString() string {
	return s.String()
}

func (s *ModifyDCdnDomainSchdmByPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDCdnDomainSchdmByPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDCdnDomainSchdmByPropertyResponse) GetBody() *ModifyDCdnDomainSchdmByPropertyResponseBody {
	return s.Body
}

func (s *ModifyDCdnDomainSchdmByPropertyResponse) SetHeaders(v map[string]*string) *ModifyDCdnDomainSchdmByPropertyResponse {
	s.Headers = v
	return s
}

func (s *ModifyDCdnDomainSchdmByPropertyResponse) SetStatusCode(v int32) *ModifyDCdnDomainSchdmByPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDCdnDomainSchdmByPropertyResponse) SetBody(v *ModifyDCdnDomainSchdmByPropertyResponseBody) *ModifyDCdnDomainSchdmByPropertyResponse {
	s.Body = v
	return s
}

func (s *ModifyDCdnDomainSchdmByPropertyResponse) Validate() error {
	return dara.Validate(s)
}

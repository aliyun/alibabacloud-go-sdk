// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnDomainSchdmByPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCdnDomainSchdmByPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCdnDomainSchdmByPropertyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCdnDomainSchdmByPropertyResponseBody) *ModifyCdnDomainSchdmByPropertyResponse
	GetBody() *ModifyCdnDomainSchdmByPropertyResponseBody
}

type ModifyCdnDomainSchdmByPropertyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCdnDomainSchdmByPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCdnDomainSchdmByPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnDomainSchdmByPropertyResponse) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainSchdmByPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCdnDomainSchdmByPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCdnDomainSchdmByPropertyResponse) GetBody() *ModifyCdnDomainSchdmByPropertyResponseBody {
	return s.Body
}

func (s *ModifyCdnDomainSchdmByPropertyResponse) SetHeaders(v map[string]*string) *ModifyCdnDomainSchdmByPropertyResponse {
	s.Headers = v
	return s
}

func (s *ModifyCdnDomainSchdmByPropertyResponse) SetStatusCode(v int32) *ModifyCdnDomainSchdmByPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCdnDomainSchdmByPropertyResponse) SetBody(v *ModifyCdnDomainSchdmByPropertyResponseBody) *ModifyCdnDomainSchdmByPropertyResponse {
	s.Body = v
	return s
}

func (s *ModifyCdnDomainSchdmByPropertyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveDomainSchdmByPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLiveDomainSchdmByPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLiveDomainSchdmByPropertyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLiveDomainSchdmByPropertyResponseBody) *ModifyLiveDomainSchdmByPropertyResponse
	GetBody() *ModifyLiveDomainSchdmByPropertyResponseBody
}

type ModifyLiveDomainSchdmByPropertyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLiveDomainSchdmByPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLiveDomainSchdmByPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveDomainSchdmByPropertyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveDomainSchdmByPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLiveDomainSchdmByPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLiveDomainSchdmByPropertyResponse) GetBody() *ModifyLiveDomainSchdmByPropertyResponseBody {
	return s.Body
}

func (s *ModifyLiveDomainSchdmByPropertyResponse) SetHeaders(v map[string]*string) *ModifyLiveDomainSchdmByPropertyResponse {
	s.Headers = v
	return s
}

func (s *ModifyLiveDomainSchdmByPropertyResponse) SetStatusCode(v int32) *ModifyLiveDomainSchdmByPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLiveDomainSchdmByPropertyResponse) SetBody(v *ModifyLiveDomainSchdmByPropertyResponseBody) *ModifyLiveDomainSchdmByPropertyResponse {
	s.Body = v
	return s
}

func (s *ModifyLiveDomainSchdmByPropertyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

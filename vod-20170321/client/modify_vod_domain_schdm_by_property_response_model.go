// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVodDomainSchdmByPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVodDomainSchdmByPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVodDomainSchdmByPropertyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVodDomainSchdmByPropertyResponseBody) *ModifyVodDomainSchdmByPropertyResponse
	GetBody() *ModifyVodDomainSchdmByPropertyResponseBody
}

type ModifyVodDomainSchdmByPropertyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVodDomainSchdmByPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVodDomainSchdmByPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVodDomainSchdmByPropertyResponse) GoString() string {
	return s.String()
}

func (s *ModifyVodDomainSchdmByPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVodDomainSchdmByPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVodDomainSchdmByPropertyResponse) GetBody() *ModifyVodDomainSchdmByPropertyResponseBody {
	return s.Body
}

func (s *ModifyVodDomainSchdmByPropertyResponse) SetHeaders(v map[string]*string) *ModifyVodDomainSchdmByPropertyResponse {
	s.Headers = v
	return s
}

func (s *ModifyVodDomainSchdmByPropertyResponse) SetStatusCode(v int32) *ModifyVodDomainSchdmByPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVodDomainSchdmByPropertyResponse) SetBody(v *ModifyVodDomainSchdmByPropertyResponseBody) *ModifyVodDomainSchdmByPropertyResponse {
	s.Body = v
	return s
}

func (s *ModifyVodDomainSchdmByPropertyResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPWByDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPWByDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPWByDomainResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPWByDomainResponseBody) *ModifyPWByDomainResponse
	GetBody() *ModifyPWByDomainResponseBody
}

type ModifyPWByDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPWByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPWByDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPWByDomainResponse) GoString() string {
	return s.String()
}

func (s *ModifyPWByDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPWByDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPWByDomainResponse) GetBody() *ModifyPWByDomainResponseBody {
	return s.Body
}

func (s *ModifyPWByDomainResponse) SetHeaders(v map[string]*string) *ModifyPWByDomainResponse {
	s.Headers = v
	return s
}

func (s *ModifyPWByDomainResponse) SetStatusCode(v int32) *ModifyPWByDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPWByDomainResponse) SetBody(v *ModifyPWByDomainResponseBody) *ModifyPWByDomainResponse {
	s.Body = v
	return s
}

func (s *ModifyPWByDomainResponse) Validate() error {
	return dara.Validate(s)
}

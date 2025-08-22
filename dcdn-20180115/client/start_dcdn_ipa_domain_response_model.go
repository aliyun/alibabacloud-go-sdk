// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDcdnIpaDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDcdnIpaDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDcdnIpaDomainResponse
	GetStatusCode() *int32
	SetBody(v *StartDcdnIpaDomainResponseBody) *StartDcdnIpaDomainResponse
	GetBody() *StartDcdnIpaDomainResponseBody
}

type StartDcdnIpaDomainResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDcdnIpaDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDcdnIpaDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDcdnIpaDomainResponse) GoString() string {
	return s.String()
}

func (s *StartDcdnIpaDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDcdnIpaDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDcdnIpaDomainResponse) GetBody() *StartDcdnIpaDomainResponseBody {
	return s.Body
}

func (s *StartDcdnIpaDomainResponse) SetHeaders(v map[string]*string) *StartDcdnIpaDomainResponse {
	s.Headers = v
	return s
}

func (s *StartDcdnIpaDomainResponse) SetStatusCode(v int32) *StartDcdnIpaDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDcdnIpaDomainResponse) SetBody(v *StartDcdnIpaDomainResponseBody) *StartDcdnIpaDomainResponse {
	s.Body = v
	return s
}

func (s *StartDcdnIpaDomainResponse) Validate() error {
	return dara.Validate(s)
}

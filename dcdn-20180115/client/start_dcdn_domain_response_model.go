// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDcdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDcdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDcdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *StartDcdnDomainResponseBody) *StartDcdnDomainResponse
	GetBody() *StartDcdnDomainResponseBody
}

type StartDcdnDomainResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDcdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *StartDcdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDcdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDcdnDomainResponse) GetBody() *StartDcdnDomainResponseBody {
	return s.Body
}

func (s *StartDcdnDomainResponse) SetHeaders(v map[string]*string) *StartDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *StartDcdnDomainResponse) SetStatusCode(v int32) *StartDcdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDcdnDomainResponse) SetBody(v *StartDcdnDomainResponseBody) *StartDcdnDomainResponse {
	s.Body = v
	return s
}

func (s *StartDcdnDomainResponse) Validate() error {
	return dara.Validate(s)
}

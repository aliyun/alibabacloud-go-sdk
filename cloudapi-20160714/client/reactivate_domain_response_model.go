// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReactivateDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReactivateDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReactivateDomainResponse
	GetStatusCode() *int32
	SetBody(v *ReactivateDomainResponseBody) *ReactivateDomainResponse
	GetBody() *ReactivateDomainResponseBody
}

type ReactivateDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReactivateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReactivateDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s ReactivateDomainResponse) GoString() string {
	return s.String()
}

func (s *ReactivateDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReactivateDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReactivateDomainResponse) GetBody() *ReactivateDomainResponseBody {
	return s.Body
}

func (s *ReactivateDomainResponse) SetHeaders(v map[string]*string) *ReactivateDomainResponse {
	s.Headers = v
	return s
}

func (s *ReactivateDomainResponse) SetStatusCode(v int32) *ReactivateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ReactivateDomainResponse) SetBody(v *ReactivateDomainResponseBody) *ReactivateDomainResponse {
	s.Body = v
	return s
}

func (s *ReactivateDomainResponse) Validate() error {
	return dara.Validate(s)
}

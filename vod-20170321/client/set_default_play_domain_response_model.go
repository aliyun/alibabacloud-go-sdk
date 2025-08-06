// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultPlayDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultPlayDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultPlayDomainResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultPlayDomainResponseBody) *SetDefaultPlayDomainResponse
	GetBody() *SetDefaultPlayDomainResponseBody
}

type SetDefaultPlayDomainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultPlayDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultPlayDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultPlayDomainResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultPlayDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultPlayDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultPlayDomainResponse) GetBody() *SetDefaultPlayDomainResponseBody {
	return s.Body
}

func (s *SetDefaultPlayDomainResponse) SetHeaders(v map[string]*string) *SetDefaultPlayDomainResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultPlayDomainResponse) SetStatusCode(v int32) *SetDefaultPlayDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultPlayDomainResponse) SetBody(v *SetDefaultPlayDomainResponseBody) *SetDefaultPlayDomainResponse {
	s.Body = v
	return s
}

func (s *SetDefaultPlayDomainResponse) Validate() error {
	return dara.Validate(s)
}

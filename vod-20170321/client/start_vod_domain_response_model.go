// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartVodDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartVodDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartVodDomainResponse
	GetStatusCode() *int32
	SetBody(v *StartVodDomainResponseBody) *StartVodDomainResponse
	GetBody() *StartVodDomainResponseBody
}

type StartVodDomainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartVodDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s StartVodDomainResponse) GoString() string {
	return s.String()
}

func (s *StartVodDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartVodDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartVodDomainResponse) GetBody() *StartVodDomainResponseBody {
	return s.Body
}

func (s *StartVodDomainResponse) SetHeaders(v map[string]*string) *StartVodDomainResponse {
	s.Headers = v
	return s
}

func (s *StartVodDomainResponse) SetStatusCode(v int32) *StartVodDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StartVodDomainResponse) SetBody(v *StartVodDomainResponseBody) *StartVodDomainResponse {
	s.Body = v
	return s
}

func (s *StartVodDomainResponse) Validate() error {
	return dara.Validate(s)
}

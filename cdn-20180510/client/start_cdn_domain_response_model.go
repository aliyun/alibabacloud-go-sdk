// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartCdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartCdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *StartCdnDomainResponseBody) *StartCdnDomainResponse
	GetBody() *StartCdnDomainResponseBody
}

type StartCdnDomainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s StartCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *StartCdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartCdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartCdnDomainResponse) GetBody() *StartCdnDomainResponseBody {
	return s.Body
}

func (s *StartCdnDomainResponse) SetHeaders(v map[string]*string) *StartCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *StartCdnDomainResponse) SetStatusCode(v int32) *StartCdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCdnDomainResponse) SetBody(v *StartCdnDomainResponseBody) *StartCdnDomainResponse {
	s.Body = v
	return s
}

func (s *StartCdnDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

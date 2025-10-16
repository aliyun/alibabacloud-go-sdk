// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrivateDnsDomainNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPrivateDnsDomainNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPrivateDnsDomainNameResponse
	GetStatusCode() *int32
	SetBody(v *AddPrivateDnsDomainNameResponseBody) *AddPrivateDnsDomainNameResponse
	GetBody() *AddPrivateDnsDomainNameResponseBody
}

type AddPrivateDnsDomainNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPrivateDnsDomainNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPrivateDnsDomainNameResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPrivateDnsDomainNameResponse) GoString() string {
	return s.String()
}

func (s *AddPrivateDnsDomainNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPrivateDnsDomainNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPrivateDnsDomainNameResponse) GetBody() *AddPrivateDnsDomainNameResponseBody {
	return s.Body
}

func (s *AddPrivateDnsDomainNameResponse) SetHeaders(v map[string]*string) *AddPrivateDnsDomainNameResponse {
	s.Headers = v
	return s
}

func (s *AddPrivateDnsDomainNameResponse) SetStatusCode(v int32) *AddPrivateDnsDomainNameResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPrivateDnsDomainNameResponse) SetBody(v *AddPrivateDnsDomainNameResponseBody) *AddPrivateDnsDomainNameResponse {
	s.Body = v
	return s
}

func (s *AddPrivateDnsDomainNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

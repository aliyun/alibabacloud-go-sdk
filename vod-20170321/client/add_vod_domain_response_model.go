// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVodDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddVodDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddVodDomainResponse
	GetStatusCode() *int32
	SetBody(v *AddVodDomainResponseBody) *AddVodDomainResponse
	GetBody() *AddVodDomainResponseBody
}

type AddVodDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddVodDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddVodDomainResponse) GoString() string {
	return s.String()
}

func (s *AddVodDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddVodDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddVodDomainResponse) GetBody() *AddVodDomainResponseBody {
	return s.Body
}

func (s *AddVodDomainResponse) SetHeaders(v map[string]*string) *AddVodDomainResponse {
	s.Headers = v
	return s
}

func (s *AddVodDomainResponse) SetStatusCode(v int32) *AddVodDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVodDomainResponse) SetBody(v *AddVodDomainResponseBody) *AddVodDomainResponse {
	s.Body = v
	return s
}

func (s *AddVodDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCdnDomainResponseBody) *ModifyCdnDomainResponse
	GetBody() *ModifyCdnDomainResponseBody
}

type ModifyCdnDomainResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCdnDomainResponse) GetBody() *ModifyCdnDomainResponseBody {
	return s.Body
}

func (s *ModifyCdnDomainResponse) SetHeaders(v map[string]*string) *ModifyCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *ModifyCdnDomainResponse) SetStatusCode(v int32) *ModifyCdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCdnDomainResponse) SetBody(v *ModifyCdnDomainResponseBody) *ModifyCdnDomainResponse {
	s.Body = v
	return s
}

func (s *ModifyCdnDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

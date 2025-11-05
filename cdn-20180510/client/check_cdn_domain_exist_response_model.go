// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCdnDomainExistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCdnDomainExistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCdnDomainExistResponse
	GetStatusCode() *int32
	SetBody(v *CheckCdnDomainExistResponseBody) *CheckCdnDomainExistResponse
	GetBody() *CheckCdnDomainExistResponseBody
}

type CheckCdnDomainExistResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCdnDomainExistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCdnDomainExistResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCdnDomainExistResponse) GoString() string {
	return s.String()
}

func (s *CheckCdnDomainExistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCdnDomainExistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCdnDomainExistResponse) GetBody() *CheckCdnDomainExistResponseBody {
	return s.Body
}

func (s *CheckCdnDomainExistResponse) SetHeaders(v map[string]*string) *CheckCdnDomainExistResponse {
	s.Headers = v
	return s
}

func (s *CheckCdnDomainExistResponse) SetStatusCode(v int32) *CheckCdnDomainExistResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCdnDomainExistResponse) SetBody(v *CheckCdnDomainExistResponseBody) *CheckCdnDomainExistResponse {
	s.Body = v
	return s
}

func (s *CheckCdnDomainExistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

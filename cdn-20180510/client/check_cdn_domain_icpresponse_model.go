// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCdnDomainICPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCdnDomainICPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCdnDomainICPResponse
	GetStatusCode() *int32
	SetBody(v *CheckCdnDomainICPResponseBody) *CheckCdnDomainICPResponse
	GetBody() *CheckCdnDomainICPResponseBody
}

type CheckCdnDomainICPResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCdnDomainICPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCdnDomainICPResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCdnDomainICPResponse) GoString() string {
	return s.String()
}

func (s *CheckCdnDomainICPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCdnDomainICPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCdnDomainICPResponse) GetBody() *CheckCdnDomainICPResponseBody {
	return s.Body
}

func (s *CheckCdnDomainICPResponse) SetHeaders(v map[string]*string) *CheckCdnDomainICPResponse {
	s.Headers = v
	return s
}

func (s *CheckCdnDomainICPResponse) SetStatusCode(v int32) *CheckCdnDomainICPResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCdnDomainICPResponse) SetBody(v *CheckCdnDomainICPResponseBody) *CheckCdnDomainICPResponse {
	s.Body = v
	return s
}

func (s *CheckCdnDomainICPResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckIpExistsInSecurityIpListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckIpExistsInSecurityIpListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckIpExistsInSecurityIpListResponse
	GetStatusCode() *int32
	SetBody(v *CheckIpExistsInSecurityIpListResponseBody) *CheckIpExistsInSecurityIpListResponse
	GetBody() *CheckIpExistsInSecurityIpListResponseBody
}

type CheckIpExistsInSecurityIpListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckIpExistsInSecurityIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckIpExistsInSecurityIpListResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckIpExistsInSecurityIpListResponse) GoString() string {
	return s.String()
}

func (s *CheckIpExistsInSecurityIpListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckIpExistsInSecurityIpListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckIpExistsInSecurityIpListResponse) GetBody() *CheckIpExistsInSecurityIpListResponseBody {
	return s.Body
}

func (s *CheckIpExistsInSecurityIpListResponse) SetHeaders(v map[string]*string) *CheckIpExistsInSecurityIpListResponse {
	s.Headers = v
	return s
}

func (s *CheckIpExistsInSecurityIpListResponse) SetStatusCode(v int32) *CheckIpExistsInSecurityIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckIpExistsInSecurityIpListResponse) SetBody(v *CheckIpExistsInSecurityIpListResponseBody) *CheckIpExistsInSecurityIpListResponse {
	s.Body = v
	return s
}

func (s *CheckIpExistsInSecurityIpListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

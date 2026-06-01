// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewMobileAgentPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewMobileAgentPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewMobileAgentPackageResponse
	GetStatusCode() *int32
	SetBody(v *RenewMobileAgentPackageResponseBody) *RenewMobileAgentPackageResponse
	GetBody() *RenewMobileAgentPackageResponseBody
}

type RenewMobileAgentPackageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewMobileAgentPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewMobileAgentPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewMobileAgentPackageResponse) GoString() string {
	return s.String()
}

func (s *RenewMobileAgentPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewMobileAgentPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewMobileAgentPackageResponse) GetBody() *RenewMobileAgentPackageResponseBody {
	return s.Body
}

func (s *RenewMobileAgentPackageResponse) SetHeaders(v map[string]*string) *RenewMobileAgentPackageResponse {
	s.Headers = v
	return s
}

func (s *RenewMobileAgentPackageResponse) SetStatusCode(v int32) *RenewMobileAgentPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewMobileAgentPackageResponse) SetBody(v *RenewMobileAgentPackageResponseBody) *RenewMobileAgentPackageResponse {
	s.Body = v
	return s
}

func (s *RenewMobileAgentPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

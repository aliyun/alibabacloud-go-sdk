// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMobileAgentPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMobileAgentPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMobileAgentPackageResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMobileAgentPackageResponseBody) *ModifyMobileAgentPackageResponse
	GetBody() *ModifyMobileAgentPackageResponseBody
}

type ModifyMobileAgentPackageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMobileAgentPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMobileAgentPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMobileAgentPackageResponse) GoString() string {
	return s.String()
}

func (s *ModifyMobileAgentPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMobileAgentPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMobileAgentPackageResponse) GetBody() *ModifyMobileAgentPackageResponseBody {
	return s.Body
}

func (s *ModifyMobileAgentPackageResponse) SetHeaders(v map[string]*string) *ModifyMobileAgentPackageResponse {
	s.Headers = v
	return s
}

func (s *ModifyMobileAgentPackageResponse) SetStatusCode(v int32) *ModifyMobileAgentPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMobileAgentPackageResponse) SetBody(v *ModifyMobileAgentPackageResponseBody) *ModifyMobileAgentPackageResponse {
	s.Body = v
	return s
}

func (s *ModifyMobileAgentPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

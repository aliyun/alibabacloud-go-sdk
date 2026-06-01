// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMobileAgentPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMobileAgentPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMobileAgentPackageResponse
	GetStatusCode() *int32
	SetBody(v *CreateMobileAgentPackageResponseBody) *CreateMobileAgentPackageResponse
	GetBody() *CreateMobileAgentPackageResponseBody
}

type CreateMobileAgentPackageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMobileAgentPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMobileAgentPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMobileAgentPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateMobileAgentPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMobileAgentPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMobileAgentPackageResponse) GetBody() *CreateMobileAgentPackageResponseBody {
	return s.Body
}

func (s *CreateMobileAgentPackageResponse) SetHeaders(v map[string]*string) *CreateMobileAgentPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateMobileAgentPackageResponse) SetStatusCode(v int32) *CreateMobileAgentPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMobileAgentPackageResponse) SetBody(v *CreateMobileAgentPackageResponseBody) *CreateMobileAgentPackageResponse {
	s.Body = v
	return s
}

func (s *CreateMobileAgentPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

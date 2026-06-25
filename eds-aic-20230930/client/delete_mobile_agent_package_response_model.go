// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMobileAgentPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMobileAgentPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMobileAgentPackageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMobileAgentPackageResponseBody) *DeleteMobileAgentPackageResponse
	GetBody() *DeleteMobileAgentPackageResponseBody
}

type DeleteMobileAgentPackageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMobileAgentPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMobileAgentPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMobileAgentPackageResponse) GoString() string {
	return s.String()
}

func (s *DeleteMobileAgentPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMobileAgentPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMobileAgentPackageResponse) GetBody() *DeleteMobileAgentPackageResponseBody {
	return s.Body
}

func (s *DeleteMobileAgentPackageResponse) SetHeaders(v map[string]*string) *DeleteMobileAgentPackageResponse {
	s.Headers = v
	return s
}

func (s *DeleteMobileAgentPackageResponse) SetStatusCode(v int32) *DeleteMobileAgentPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMobileAgentPackageResponse) SetBody(v *DeleteMobileAgentPackageResponseBody) *DeleteMobileAgentPackageResponse {
	s.Body = v
	return s
}

func (s *DeleteMobileAgentPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

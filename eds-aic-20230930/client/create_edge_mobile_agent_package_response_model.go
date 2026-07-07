// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeMobileAgentPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEdgeMobileAgentPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEdgeMobileAgentPackageResponse
	GetStatusCode() *int32
	SetBody(v *CreateEdgeMobileAgentPackageResponseBody) *CreateEdgeMobileAgentPackageResponse
	GetBody() *CreateEdgeMobileAgentPackageResponseBody
}

type CreateEdgeMobileAgentPackageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEdgeMobileAgentPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEdgeMobileAgentPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeMobileAgentPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateEdgeMobileAgentPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEdgeMobileAgentPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEdgeMobileAgentPackageResponse) GetBody() *CreateEdgeMobileAgentPackageResponseBody {
	return s.Body
}

func (s *CreateEdgeMobileAgentPackageResponse) SetHeaders(v map[string]*string) *CreateEdgeMobileAgentPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateEdgeMobileAgentPackageResponse) SetStatusCode(v int32) *CreateEdgeMobileAgentPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageResponse) SetBody(v *CreateEdgeMobileAgentPackageResponseBody) *CreateEdgeMobileAgentPackageResponse {
	s.Body = v
	return s
}

func (s *CreateEdgeMobileAgentPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

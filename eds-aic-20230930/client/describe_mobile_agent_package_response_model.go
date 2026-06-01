// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileAgentPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMobileAgentPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMobileAgentPackageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMobileAgentPackageResponseBody) *DescribeMobileAgentPackageResponse
	GetBody() *DescribeMobileAgentPackageResponseBody
}

type DescribeMobileAgentPackageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMobileAgentPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMobileAgentPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMobileAgentPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMobileAgentPackageResponse) GetBody() *DescribeMobileAgentPackageResponseBody {
	return s.Body
}

func (s *DescribeMobileAgentPackageResponse) SetHeaders(v map[string]*string) *DescribeMobileAgentPackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeMobileAgentPackageResponse) SetStatusCode(v int32) *DescribeMobileAgentPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMobileAgentPackageResponse) SetBody(v *DescribeMobileAgentPackageResponseBody) *DescribeMobileAgentPackageResponse {
	s.Body = v
	return s
}

func (s *DescribeMobileAgentPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

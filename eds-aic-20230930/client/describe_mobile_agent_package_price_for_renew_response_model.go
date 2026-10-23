// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileAgentPackagePriceForRenewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMobileAgentPackagePriceForRenewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMobileAgentPackagePriceForRenewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMobileAgentPackagePriceForRenewResponseBody) *DescribeMobileAgentPackagePriceForRenewResponse
	GetBody() *DescribeMobileAgentPackagePriceForRenewResponseBody
}

type DescribeMobileAgentPackagePriceForRenewResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMobileAgentPackagePriceForRenewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMobileAgentPackagePriceForRenewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceForRenewResponse) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceForRenewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMobileAgentPackagePriceForRenewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMobileAgentPackagePriceForRenewResponse) GetBody() *DescribeMobileAgentPackagePriceForRenewResponseBody {
	return s.Body
}

func (s *DescribeMobileAgentPackagePriceForRenewResponse) SetHeaders(v map[string]*string) *DescribeMobileAgentPackagePriceForRenewResponse {
	s.Headers = v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponse) SetStatusCode(v int32) *DescribeMobileAgentPackagePriceForRenewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponse) SetBody(v *DescribeMobileAgentPackagePriceForRenewResponseBody) *DescribeMobileAgentPackagePriceForRenewResponse {
	s.Body = v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

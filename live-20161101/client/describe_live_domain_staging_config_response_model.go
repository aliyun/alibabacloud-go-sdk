// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainStagingConfigResponseBody) *DescribeLiveDomainStagingConfigResponse
	GetBody() *DescribeLiveDomainStagingConfigResponseBody
}

type DescribeLiveDomainStagingConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainStagingConfigResponse) GetBody() *DescribeLiveDomainStagingConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainStagingConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponse) SetStatusCode(v int32) *DescribeLiveDomainStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponse) SetBody(v *DescribeLiveDomainStagingConfigResponseBody) *DescribeLiveDomainStagingConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

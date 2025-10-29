// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRecordUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainRecordUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainRecordUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainRecordUsageDataResponseBody) *DescribeLiveDomainRecordUsageDataResponse
	GetBody() *DescribeLiveDomainRecordUsageDataResponseBody
}

type DescribeLiveDomainRecordUsageDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainRecordUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainRecordUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRecordUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRecordUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainRecordUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainRecordUsageDataResponse) GetBody() *DescribeLiveDomainRecordUsageDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainRecordUsageDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainRecordUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponse) SetStatusCode(v int32) *DescribeLiveDomainRecordUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponse) SetBody(v *DescribeLiveDomainRecordUsageDataResponseBody) *DescribeLiveDomainRecordUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainRecordUsageDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

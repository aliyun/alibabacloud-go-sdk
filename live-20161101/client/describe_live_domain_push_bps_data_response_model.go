// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPushBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainPushBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainPushBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainPushBpsDataResponseBody) *DescribeLiveDomainPushBpsDataResponse
	GetBody() *DescribeLiveDomainPushBpsDataResponseBody
}

type DescribeLiveDomainPushBpsDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainPushBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainPushBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPushBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainPushBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainPushBpsDataResponse) GetBody() *DescribeLiveDomainPushBpsDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainPushBpsDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainPushBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponse) SetStatusCode(v int32) *DescribeLiveDomainPushBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponse) SetBody(v *DescribeLiveDomainPushBpsDataResponseBody) *DescribeLiveDomainPushBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeQpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealTimeQpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealTimeQpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealTimeQpsDataResponseBody) *DescribeDomainRealTimeQpsDataResponse
	GetBody() *DescribeDomainRealTimeQpsDataResponseBody
}

type DescribeDomainRealTimeQpsDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealTimeQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealTimeQpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeQpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealTimeQpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealTimeQpsDataResponse) GetBody() *DescribeDomainRealTimeQpsDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRealTimeQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponse) SetStatusCode(v int32) *DescribeDomainRealTimeQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponse) SetBody(v *DescribeDomainRealTimeQpsDataResponseBody) *DescribeDomainRealTimeQpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

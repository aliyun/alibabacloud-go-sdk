// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealTimeHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainRealTimeHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainRealTimeHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) *DescribeLiveDomainRealTimeHttpCodeDataResponse
	GetBody() *DescribeLiveDomainRealTimeHttpCodeDataResponseBody
}

type DescribeLiveDomainRealTimeHttpCodeDataResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainRealTimeHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) GetBody() *DescribeLiveDomainRealTimeHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainRealTimeHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) SetStatusCode(v int32) *DescribeLiveDomainRealTimeHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) SetBody(v *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) *DescribeLiveDomainRealTimeHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

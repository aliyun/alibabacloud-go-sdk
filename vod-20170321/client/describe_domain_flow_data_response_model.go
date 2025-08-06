// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainFlowDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainFlowDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainFlowDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainFlowDataResponseBody) *DescribeDomainFlowDataResponse
	GetBody() *DescribeDomainFlowDataResponseBody
}

type DescribeDomainFlowDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainFlowDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainFlowDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainFlowDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainFlowDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainFlowDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainFlowDataResponse) GetBody() *DescribeDomainFlowDataResponseBody {
	return s.Body
}

func (s *DescribeDomainFlowDataResponse) SetHeaders(v map[string]*string) *DescribeDomainFlowDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainFlowDataResponse) SetStatusCode(v int32) *DescribeDomainFlowDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainFlowDataResponse) SetBody(v *DescribeDomainFlowDataResponseBody) *DescribeDomainFlowDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainFlowDataResponse) Validate() error {
	return dara.Validate(s)
}

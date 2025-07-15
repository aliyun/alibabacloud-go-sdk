// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveProducerUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveProducerUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveProducerUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveProducerUsageDataResponseBody) *DescribeLiveProducerUsageDataResponse
	GetBody() *DescribeLiveProducerUsageDataResponseBody
}

type DescribeLiveProducerUsageDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveProducerUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveProducerUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveProducerUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveProducerUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveProducerUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveProducerUsageDataResponse) GetBody() *DescribeLiveProducerUsageDataResponseBody {
	return s.Body
}

func (s *DescribeLiveProducerUsageDataResponse) SetHeaders(v map[string]*string) *DescribeLiveProducerUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveProducerUsageDataResponse) SetStatusCode(v int32) *DescribeLiveProducerUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveProducerUsageDataResponse) SetBody(v *DescribeLiveProducerUsageDataResponseBody) *DescribeLiveProducerUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveProducerUsageDataResponse) Validate() error {
	return dara.Validate(s)
}

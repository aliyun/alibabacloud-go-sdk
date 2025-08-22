// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRealTimeDeliveryFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnRealTimeDeliveryFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnRealTimeDeliveryFieldResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnRealTimeDeliveryFieldResponseBody) *DescribeDcdnRealTimeDeliveryFieldResponse
	GetBody() *DescribeDcdnRealTimeDeliveryFieldResponseBody
}

type DescribeDcdnRealTimeDeliveryFieldResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnRealTimeDeliveryFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnRealTimeDeliveryFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRealTimeDeliveryFieldResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponse) GetBody() *DescribeDcdnRealTimeDeliveryFieldResponseBody {
	return s.Body
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponse) SetHeaders(v map[string]*string) *DescribeDcdnRealTimeDeliveryFieldResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponse) SetStatusCode(v int32) *DescribeDcdnRealTimeDeliveryFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponse) SetBody(v *DescribeDcdnRealTimeDeliveryFieldResponseBody) *DescribeDcdnRealTimeDeliveryFieldResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldResponse) Validate() error {
	return dara.Validate(s)
}

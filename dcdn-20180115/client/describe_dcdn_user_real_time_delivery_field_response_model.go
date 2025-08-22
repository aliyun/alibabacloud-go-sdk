// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserRealTimeDeliveryFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserRealTimeDeliveryFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserRealTimeDeliveryFieldResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserRealTimeDeliveryFieldResponseBody) *DescribeDcdnUserRealTimeDeliveryFieldResponse
	GetBody() *DescribeDcdnUserRealTimeDeliveryFieldResponseBody
}

type DescribeDcdnUserRealTimeDeliveryFieldResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserRealTimeDeliveryFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserRealTimeDeliveryFieldResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponse) GetBody() *DescribeDcdnUserRealTimeDeliveryFieldResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserRealTimeDeliveryFieldResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponse) SetStatusCode(v int32) *DescribeDcdnUserRealTimeDeliveryFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponse) SetBody(v *DescribeDcdnUserRealTimeDeliveryFieldResponseBody) *DescribeDcdnUserRealTimeDeliveryFieldResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldResponse) Validate() error {
	return dara.Validate(s)
}

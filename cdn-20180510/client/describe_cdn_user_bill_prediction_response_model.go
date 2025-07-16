// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserBillPredictionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnUserBillPredictionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnUserBillPredictionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnUserBillPredictionResponseBody) *DescribeCdnUserBillPredictionResponse
	GetBody() *DescribeCdnUserBillPredictionResponseBody
}

type DescribeCdnUserBillPredictionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnUserBillPredictionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnUserBillPredictionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillPredictionResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillPredictionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnUserBillPredictionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnUserBillPredictionResponse) GetBody() *DescribeCdnUserBillPredictionResponseBody {
	return s.Body
}

func (s *DescribeCdnUserBillPredictionResponse) SetHeaders(v map[string]*string) *DescribeCdnUserBillPredictionResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserBillPredictionResponse) SetStatusCode(v int32) *DescribeCdnUserBillPredictionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponse) SetBody(v *DescribeCdnUserBillPredictionResponseBody) *DescribeCdnUserBillPredictionResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnUserBillPredictionResponse) Validate() error {
	return dara.Validate(s)
}

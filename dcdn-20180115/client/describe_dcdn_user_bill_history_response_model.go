// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserBillHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserBillHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserBillHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserBillHistoryResponseBody) *DescribeDcdnUserBillHistoryResponse
	GetBody() *DescribeDcdnUserBillHistoryResponseBody
}

type DescribeDcdnUserBillHistoryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserBillHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserBillHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserBillHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserBillHistoryResponse) GetBody() *DescribeDcdnUserBillHistoryResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserBillHistoryResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserBillHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponse) SetStatusCode(v int32) *DescribeDcdnUserBillHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponse) SetBody(v *DescribeDcdnUserBillHistoryResponseBody) *DescribeDcdnUserBillHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponse) Validate() error {
	return dara.Validate(s)
}

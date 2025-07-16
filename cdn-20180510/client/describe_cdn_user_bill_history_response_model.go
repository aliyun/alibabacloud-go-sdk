// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserBillHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnUserBillHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnUserBillHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnUserBillHistoryResponseBody) *DescribeCdnUserBillHistoryResponse
	GetBody() *DescribeCdnUserBillHistoryResponseBody
}

type DescribeCdnUserBillHistoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnUserBillHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnUserBillHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnUserBillHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnUserBillHistoryResponse) GetBody() *DescribeCdnUserBillHistoryResponseBody {
	return s.Body
}

func (s *DescribeCdnUserBillHistoryResponse) SetHeaders(v map[string]*string) *DescribeCdnUserBillHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserBillHistoryResponse) SetStatusCode(v int32) *DescribeCdnUserBillHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponse) SetBody(v *DescribeCdnUserBillHistoryResponseBody) *DescribeCdnUserBillHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnUserBillHistoryResponse) Validate() error {
	return dara.Validate(s)
}

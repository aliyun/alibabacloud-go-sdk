// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarketRemainsQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemainsQuota(v int64) *DescribeMarketRemainsQuotaResponseBody
	GetRemainsQuota() *int64
	SetRequestId(v string) *DescribeMarketRemainsQuotaResponseBody
	GetRequestId() *string
}

type DescribeMarketRemainsQuotaResponseBody struct {
	// The remaining quota.
	//
	// example:
	//
	// 1000
	RemainsQuota *int64 `json:"RemainsQuota,omitempty" xml:"RemainsQuota,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E7FE7172-AA75-5880-B6F7-C00893E9BC06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMarketRemainsQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketRemainsQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMarketRemainsQuotaResponseBody) GetRemainsQuota() *int64 {
	return s.RemainsQuota
}

func (s *DescribeMarketRemainsQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMarketRemainsQuotaResponseBody) SetRemainsQuota(v int64) *DescribeMarketRemainsQuotaResponseBody {
	s.RemainsQuota = &v
	return s
}

func (s *DescribeMarketRemainsQuotaResponseBody) SetRequestId(v string) *DescribeMarketRemainsQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMarketRemainsQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

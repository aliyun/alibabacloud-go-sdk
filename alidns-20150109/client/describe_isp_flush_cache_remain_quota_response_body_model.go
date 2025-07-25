// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheRemainQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeIspFlushCacheRemainQuotaResponseBody
	GetRequestId() *string
	SetTelecomRemainQuota(v int32) *DescribeIspFlushCacheRemainQuotaResponseBody
	GetTelecomRemainQuota() *int32
}

type DescribeIspFlushCacheRemainQuotaResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TelecomRemainQuota *int32  `json:"TelecomRemainQuota,omitempty" xml:"TelecomRemainQuota,omitempty"`
}

func (s DescribeIspFlushCacheRemainQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheRemainQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheRemainQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIspFlushCacheRemainQuotaResponseBody) GetTelecomRemainQuota() *int32 {
	return s.TelecomRemainQuota
}

func (s *DescribeIspFlushCacheRemainQuotaResponseBody) SetRequestId(v string) *DescribeIspFlushCacheRemainQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIspFlushCacheRemainQuotaResponseBody) SetTelecomRemainQuota(v int32) *DescribeIspFlushCacheRemainQuotaResponseBody {
	s.TelecomRemainQuota = &v
	return s
}

func (s *DescribeIspFlushCacheRemainQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

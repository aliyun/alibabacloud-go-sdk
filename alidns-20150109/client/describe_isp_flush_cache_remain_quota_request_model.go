// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheRemainQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeIspFlushCacheRemainQuotaRequest
	GetLang() *string
}

type DescribeIspFlushCacheRemainQuotaRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeIspFlushCacheRemainQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheRemainQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheRemainQuotaRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeIspFlushCacheRemainQuotaRequest) SetLang(v string) *DescribeIspFlushCacheRemainQuotaRequest {
	s.Lang = &v
	return s
}

func (s *DescribeIspFlushCacheRemainQuotaRequest) Validate() error {
	return dara.Validate(s)
}

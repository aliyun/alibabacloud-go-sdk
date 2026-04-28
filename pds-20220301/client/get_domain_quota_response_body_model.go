// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSizeQuota(v int64) *GetDomainQuotaResponseBody
	GetSizeQuota() *int64
	SetSizeUsed(v int64) *GetDomainQuotaResponseBody
	GetSizeUsed() *int64
	SetUserCountQuota(v int64) *GetDomainQuotaResponseBody
	GetUserCountQuota() *int64
	SetUserCountUsed(v int64) *GetDomainQuotaResponseBody
	GetUserCountUsed() *int64
}

type GetDomainQuotaResponseBody struct {
	SizeQuota      *int64 `json:"size_quota,omitempty" xml:"size_quota,omitempty"`
	SizeUsed       *int64 `json:"size_used,omitempty" xml:"size_used,omitempty"`
	UserCountQuota *int64 `json:"user_count_quota,omitempty" xml:"user_count_quota,omitempty"`
	UserCountUsed  *int64 `json:"user_count_used,omitempty" xml:"user_count_used,omitempty"`
}

func (s GetDomainQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDomainQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainQuotaResponseBody) GetSizeQuota() *int64 {
	return s.SizeQuota
}

func (s *GetDomainQuotaResponseBody) GetSizeUsed() *int64 {
	return s.SizeUsed
}

func (s *GetDomainQuotaResponseBody) GetUserCountQuota() *int64 {
	return s.UserCountQuota
}

func (s *GetDomainQuotaResponseBody) GetUserCountUsed() *int64 {
	return s.UserCountUsed
}

func (s *GetDomainQuotaResponseBody) SetSizeQuota(v int64) *GetDomainQuotaResponseBody {
	s.SizeQuota = &v
	return s
}

func (s *GetDomainQuotaResponseBody) SetSizeUsed(v int64) *GetDomainQuotaResponseBody {
	s.SizeUsed = &v
	return s
}

func (s *GetDomainQuotaResponseBody) SetUserCountQuota(v int64) *GetDomainQuotaResponseBody {
	s.UserCountQuota = &v
	return s
}

func (s *GetDomainQuotaResponseBody) SetUserCountUsed(v int64) *GetDomainQuotaResponseBody {
	s.UserCountUsed = &v
	return s
}

func (s *GetDomainQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

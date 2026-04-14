// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidationQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetValidationQuotaResponseBody
	GetRequestId() *string
	SetTotalQuota(v int32) *GetValidationQuotaResponseBody
	GetTotalQuota() *int32
	SetUsedQuota(v int32) *GetValidationQuotaResponseBody
	GetUsedQuota() *int32
}

type GetValidationQuotaResponseBody struct {
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalQuota *int32 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// example:
	//
	// 1
	UsedQuota *int32 `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
}

func (s GetValidationQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetValidationQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetValidationQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetValidationQuotaResponseBody) GetTotalQuota() *int32 {
	return s.TotalQuota
}

func (s *GetValidationQuotaResponseBody) GetUsedQuota() *int32 {
	return s.UsedQuota
}

func (s *GetValidationQuotaResponseBody) SetRequestId(v string) *GetValidationQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetValidationQuotaResponseBody) SetTotalQuota(v int32) *GetValidationQuotaResponseBody {
	s.TotalQuota = &v
	return s
}

func (s *GetValidationQuotaResponseBody) SetUsedQuota(v int32) *GetValidationQuotaResponseBody {
	s.UsedQuota = &v
	return s
}

func (s *GetValidationQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

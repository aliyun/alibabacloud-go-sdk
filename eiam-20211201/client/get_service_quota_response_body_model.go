// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceQuotaResponseBody
	GetRequestId() *string
	SetServiceQuota(v *GetServiceQuotaResponseBodyServiceQuota) *GetServiceQuotaResponseBody
	GetServiceQuota() *GetServiceQuotaResponseBodyServiceQuota
}

type GetServiceQuotaResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceQuota *GetServiceQuotaResponseBodyServiceQuota `json:"ServiceQuota,omitempty" xml:"ServiceQuota,omitempty" type:"Struct"`
}

func (s GetServiceQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceQuotaResponseBody) GetServiceQuota() *GetServiceQuotaResponseBodyServiceQuota {
	return s.ServiceQuota
}

func (s *GetServiceQuotaResponseBody) SetRequestId(v string) *GetServiceQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceQuotaResponseBody) SetServiceQuota(v *GetServiceQuotaResponseBodyServiceQuota) *GetServiceQuotaResponseBody {
	s.ServiceQuota = v
	return s
}

func (s *GetServiceQuotaResponseBody) Validate() error {
	if s.ServiceQuota != nil {
		if err := s.ServiceQuota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceQuotaResponseBodyServiceQuota struct {
	// Quota 配额的唯一标识。
	//
	// example:
	//
	// instanceTrialNumber
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// Quota 配额的值。
	//
	// example:
	//
	// 5
	QuotaValue *int64 `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	// Quota 配额的当前用量。
	//
	// example:
	//
	// 1
	UsedQuotaValue *int64 `json:"UsedQuotaValue,omitempty" xml:"UsedQuotaValue,omitempty"`
}

func (s GetServiceQuotaResponseBodyServiceQuota) String() string {
	return dara.Prettify(s)
}

func (s GetServiceQuotaResponseBodyServiceQuota) GoString() string {
	return s.String()
}

func (s *GetServiceQuotaResponseBodyServiceQuota) GetQuotaType() *string {
	return s.QuotaType
}

func (s *GetServiceQuotaResponseBodyServiceQuota) GetQuotaValue() *int64 {
	return s.QuotaValue
}

func (s *GetServiceQuotaResponseBodyServiceQuota) GetUsedQuotaValue() *int64 {
	return s.UsedQuotaValue
}

func (s *GetServiceQuotaResponseBodyServiceQuota) SetQuotaType(v string) *GetServiceQuotaResponseBodyServiceQuota {
	s.QuotaType = &v
	return s
}

func (s *GetServiceQuotaResponseBodyServiceQuota) SetQuotaValue(v int64) *GetServiceQuotaResponseBodyServiceQuota {
	s.QuotaValue = &v
	return s
}

func (s *GetServiceQuotaResponseBodyServiceQuota) SetUsedQuotaValue(v int64) *GetServiceQuotaResponseBodyServiceQuota {
	s.UsedQuotaValue = &v
	return s
}

func (s *GetServiceQuotaResponseBodyServiceQuota) Validate() error {
	return dara.Validate(s)
}

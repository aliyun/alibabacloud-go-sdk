// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLogDeliveryQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GetUserLogDeliveryQuotaResponseBody
	GetBusinessType() *string
	SetFreeQuota(v int64) *GetUserLogDeliveryQuotaResponseBody
	GetFreeQuota() *int64
	SetRequestId(v string) *GetUserLogDeliveryQuotaResponseBody
	GetRequestId() *string
}

type GetUserLogDeliveryQuotaResponseBody struct {
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	FreeQuota    *int64  `json:"FreeQuota,omitempty" xml:"FreeQuota,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserLogDeliveryQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserLogDeliveryQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserLogDeliveryQuotaResponseBody) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetUserLogDeliveryQuotaResponseBody) GetFreeQuota() *int64 {
	return s.FreeQuota
}

func (s *GetUserLogDeliveryQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserLogDeliveryQuotaResponseBody) SetBusinessType(v string) *GetUserLogDeliveryQuotaResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetUserLogDeliveryQuotaResponseBody) SetFreeQuota(v int64) *GetUserLogDeliveryQuotaResponseBody {
	s.FreeQuota = &v
	return s
}

func (s *GetUserLogDeliveryQuotaResponseBody) SetRequestId(v string) *GetUserLogDeliveryQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserLogDeliveryQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

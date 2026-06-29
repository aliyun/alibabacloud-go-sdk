// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceQuotasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstanceQuotasResponseBody
	GetInstanceId() *string
	SetQuotas(v []*ListInstanceQuotasResponseBodyQuotas) *ListInstanceQuotasResponseBody
	GetQuotas() []*ListInstanceQuotasResponseBodyQuotas
	SetRequestId(v string) *ListInstanceQuotasResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListInstanceQuotasResponseBody
	GetStatus() *string
}

type ListInstanceQuotasResponseBody struct {
	// The plan instance ID.
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of plan instance quotas.
	Quotas []*ListInstanceQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The plan instance status. Valid values:
	//
	// - **online**: The instance is in normal service.
	//
	// - **offline**: The instance has expired but has not exceeded the grace period and is unavailable.
	//
	// - **disable**: The instance has been released.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceQuotasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceQuotasResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceQuotasResponseBody) GetQuotas() []*ListInstanceQuotasResponseBodyQuotas {
	return s.Quotas
}

func (s *ListInstanceQuotasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceQuotasResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceQuotasResponseBody) SetInstanceId(v string) *ListInstanceQuotasResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceQuotasResponseBody) SetQuotas(v []*ListInstanceQuotasResponseBodyQuotas) *ListInstanceQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListInstanceQuotasResponseBody) SetRequestId(v string) *ListInstanceQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceQuotasResponseBody) SetStatus(v string) *ListInstanceQuotasResponseBody {
	s.Status = &v
	return s
}

func (s *ListInstanceQuotasResponseBody) Validate() error {
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceQuotasResponseBodyQuotas struct {
	// The quota name.
	//
	// example:
	//
	// customHttpCert
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The quota value.
	//
	// example:
	//
	// 10
	QuotaValue *string `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	// The threshold type of the quota. Valid values:
	//
	// - **value**: enumeration type. The enumeration range of quota values.
	//
	// - **bool**: Boolean type. Indicates whether the quota is available.
	//
	// - **num**: numeric type. The maximum usage of the quota.
	//
	// - **range**: range type. The value range of the quota.
	//
	// - **custom**: custom type. Other types that do not fall into the preceding four threshold types.
	//
	// example:
	//
	// bool
	QuotaValueType *string `json:"QuotaValueType,omitempty" xml:"QuotaValueType,omitempty"`
}

func (s ListInstanceQuotasResponseBodyQuotas) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceQuotasResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListInstanceQuotasResponseBodyQuotas) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListInstanceQuotasResponseBodyQuotas) GetQuotaValue() *string {
	return s.QuotaValue
}

func (s *ListInstanceQuotasResponseBodyQuotas) GetQuotaValueType() *string {
	return s.QuotaValueType
}

func (s *ListInstanceQuotasResponseBodyQuotas) SetQuotaName(v string) *ListInstanceQuotasResponseBodyQuotas {
	s.QuotaName = &v
	return s
}

func (s *ListInstanceQuotasResponseBodyQuotas) SetQuotaValue(v string) *ListInstanceQuotasResponseBodyQuotas {
	s.QuotaValue = &v
	return s
}

func (s *ListInstanceQuotasResponseBodyQuotas) SetQuotaValueType(v string) *ListInstanceQuotasResponseBodyQuotas {
	s.QuotaValueType = &v
	return s
}

func (s *ListInstanceQuotasResponseBodyQuotas) Validate() error {
	return dara.Validate(s)
}

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
	// The plan ID.
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The quotas in the plan.
	Quotas []*ListInstanceQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The plan status. Valid values:
	//
	// 	- online: The plan is in service.
	//
	// 	- offline: The plan has expired within an allowable period. In this state, the plan is unavailable.
	//
	// 	- disable: The plan is released.
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
	return dara.Validate(s)
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
	// 	- value: enumerates the values of the quota.
	//
	// 	- bool: specifies whether the quota is available.
	//
	// 	- num: the upper limit of the quota.
	//
	// 	- range: the value range for the quota.
	//
	// 	- custom: other types than the preceding four quota threshold types.
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

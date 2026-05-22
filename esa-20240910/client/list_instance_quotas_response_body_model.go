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
	InstanceId *string                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Quotas     []*ListInstanceQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status     *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
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
	QuotaName      *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	QuotaValue     *string `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
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

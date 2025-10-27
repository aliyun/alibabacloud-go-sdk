// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLdpsNamespacedQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNamespacedQuotas(v []*GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) *GetLdpsNamespacedQuotaResponseBody
	GetNamespacedQuotas() []*GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas
	SetRequestId(v string) *GetLdpsNamespacedQuotaResponseBody
	GetRequestId() *string
}

type GetLdpsNamespacedQuotaResponseBody struct {
	NamespacedQuotas []*GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas `json:"NamespacedQuotas,omitempty" xml:"NamespacedQuotas,omitempty" type:"Repeated"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLdpsNamespacedQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLdpsNamespacedQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetLdpsNamespacedQuotaResponseBody) GetNamespacedQuotas() []*GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	return s.NamespacedQuotas
}

func (s *GetLdpsNamespacedQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLdpsNamespacedQuotaResponseBody) SetNamespacedQuotas(v []*GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) *GetLdpsNamespacedQuotaResponseBody {
	s.NamespacedQuotas = v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBody) SetRequestId(v string) *GetLdpsNamespacedQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBody) Validate() error {
	if s.NamespacedQuotas != nil {
		for _, item := range s.NamespacedQuotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas struct {
	CpuAmount    *string `json:"CpuAmount,omitempty" xml:"CpuAmount,omitempty"`
	MemoryAmount *string `json:"MemoryAmount,omitempty" xml:"MemoryAmount,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UsedCpu      *string `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
	UsedMemory   *string `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) String() string {
	return dara.Prettify(s)
}

func (s GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) GoString() string {
	return s.String()
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) GetCpuAmount() *string {
	return s.CpuAmount
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) GetMemoryAmount() *string {
	return s.MemoryAmount
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) GetName() *string {
	return s.Name
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) GetUsedCpu() *string {
	return s.UsedCpu
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) GetUsedMemory() *string {
	return s.UsedMemory
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) SetCpuAmount(v string) *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	s.CpuAmount = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) SetMemoryAmount(v string) *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	s.MemoryAmount = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) SetName(v string) *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	s.Name = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) SetUsedCpu(v string) *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	s.UsedCpu = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) SetUsedMemory(v string) *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	s.UsedMemory = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) Validate() error {
	return dara.Validate(s)
}

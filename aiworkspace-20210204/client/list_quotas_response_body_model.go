// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotas(v []*ListQuotasResponseBodyQuotas) *ListQuotasResponseBody
	GetQuotas() []*ListQuotasResponseBodyQuotas
	SetRequestId(v string) *ListQuotasResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListQuotasResponseBody
	GetTotalCount() *int64
}

type ListQuotasResponseBody struct {
	// The returned quotas.
	Quotas []*ListQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of quotas that meet the filter conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBody) GetQuotas() []*ListQuotasResponseBodyQuotas {
	return s.Quotas
}

func (s *ListQuotasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotasResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQuotasResponseBody) SetQuotas(v []*ListQuotasResponseBodyQuotas) *ListQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListQuotasResponseBody) SetRequestId(v string) *ListQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotasResponseBody) SetTotalCount(v int64) *ListQuotasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotasResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotas struct {
	// The alias of the quota.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// 1828233
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The billing method. Valid values:
	//
	// 	- isolate: subscription
	//
	// 	- share: pay-as-you-go
	//
	// example:
	//
	// isolate
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The quota name.
	//
	// example:
	//
	// quota-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The product code. Valid values:
	//
	// 	- PAI_isolate: CPU subscription resource groups of PAI
	//
	// 	- PAI_share: GPU pay-as-you-go resource groups of PAI
	//
	// example:
	//
	// MaxCompute_share
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota type. Valid value:
	//
	// PAI: indicates GPU resource groups of MaxCompute.
	//
	// example:
	//
	// MaxCompute
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The quota specifications.
	//
	// example:
	//
	// {\\"cu\\":\\"11500\\",\\"minCu\\":\\"2300\\",\\"parentId\\":\\"0\\"}
	Specs []*ListQuotasResponseBodyQuotasSpecs `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
}

func (s ListQuotasResponseBodyQuotas) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotas) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListQuotasResponseBodyQuotas) GetId() *string {
	return s.Id
}

func (s *ListQuotasResponseBodyQuotas) GetMode() *string {
	return s.Mode
}

func (s *ListQuotasResponseBodyQuotas) GetName() *string {
	return s.Name
}

func (s *ListQuotasResponseBodyQuotas) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotasResponseBodyQuotas) GetQuotaType() *string {
	return s.QuotaType
}

func (s *ListQuotasResponseBodyQuotas) GetSpecs() []*ListQuotasResponseBodyQuotasSpecs {
	return s.Specs
}

func (s *ListQuotasResponseBodyQuotas) SetDisplayName(v string) *ListQuotasResponseBodyQuotas {
	s.DisplayName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetId(v string) *ListQuotasResponseBodyQuotas {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetMode(v string) *ListQuotasResponseBodyQuotas {
	s.Mode = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetName(v string) *ListQuotasResponseBodyQuotas {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetProductCode(v string) *ListQuotasResponseBodyQuotas {
	s.ProductCode = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetQuotaType(v string) *ListQuotasResponseBodyQuotas {
	s.QuotaType = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetSpecs(v []*ListQuotasResponseBodyQuotasSpecs) *ListQuotasResponseBodyQuotas {
	s.Specs = v
	return s
}

func (s *ListQuotasResponseBodyQuotas) Validate() error {
	return dara.Validate(s)
}

type ListQuotasResponseBodyQuotasSpecs struct {
	// The specification name.
	//
	// example:
	//
	// cu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The specification type. The parameter can be left empty.
	//
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The specification value.
	//
	// example:
	//
	// 11500
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQuotasResponseBodyQuotasSpecs) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBodyQuotasSpecs) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotasSpecs) GetName() *string {
	return s.Name
}

func (s *ListQuotasResponseBodyQuotasSpecs) GetType() *string {
	return s.Type
}

func (s *ListQuotasResponseBodyQuotasSpecs) GetValue() *string {
	return s.Value
}

func (s *ListQuotasResponseBodyQuotasSpecs) SetName(v string) *ListQuotasResponseBodyQuotasSpecs {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyQuotasSpecs) SetType(v string) *ListQuotasResponseBodyQuotasSpecs {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyQuotasSpecs) SetValue(v string) *ListQuotasResponseBodyQuotasSpecs {
	s.Value = &v
	return s
}

func (s *ListQuotasResponseBodyQuotasSpecs) Validate() error {
	return dara.Validate(s)
}

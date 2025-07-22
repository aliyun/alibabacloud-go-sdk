// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanAllocateCreditAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEcid(v string) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetEcid() *string
	SetEcidAllocatedCreditAmount(v string) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetEcidAllocatedCreditAmount() *string
	SetEcidCreditAmount(v string) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetEcidCreditAmount() *string
	SetFundAccountEcid(v string) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetFundAccountEcid() *string
	SetFundAccountId(v int64) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetFundAccountId() *int64
	SetFundAccountName(v string) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetFundAccountName() *string
	SetFundAccountOwnerAccountId(v int64) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetFundAccountOwnerAccountId() *int64
	SetMaxCanAllocateCreditAmount(v string) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetMaxCanAllocateCreditAmount() *string
	SetMetadata(v interface{}) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetMetadata() interface{}
	SetMinCanAllocateCreditAmount(v string) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetMinCanAllocateCreditAmount() *string
	SetNbid(v string) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetNbid() *string
	SetRequestId(v string) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetRequestId() *string
	SetSite(v string) *GetFundAccountCanAllocateCreditAmountResponseBody
	GetSite() *string
}

type GetFundAccountCanAllocateCreditAmountResponseBody struct {
	// example:
	//
	// 2032123221
	Ecid *string `json:"Ecid,omitempty" xml:"Ecid,omitempty"`
	// example:
	//
	// 300
	EcidAllocatedCreditAmount *string `json:"EcidAllocatedCreditAmount,omitempty" xml:"EcidAllocatedCreditAmount,omitempty"`
	// example:
	//
	// 1000
	EcidCreditAmount *string `json:"EcidCreditAmount,omitempty" xml:"EcidCreditAmount,omitempty"`
	// example:
	//
	// 202321232
	FundAccountEcid *string `json:"FundAccountEcid,omitempty" xml:"FundAccountEcid,omitempty"`
	// example:
	//
	// 12332112
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 云某的名称
	FundAccountName *string `json:"FundAccountName,omitempty" xml:"FundAccountName,omitempty"`
	// example:
	//
	// 123433121
	FundAccountOwnerAccountId *int64 `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// 1500
	MaxCanAllocateCreditAmount *string `json:"MaxCanAllocateCreditAmount,omitempty" xml:"MaxCanAllocateCreditAmount,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 200
	MinCanAllocateCreditAmount *string `json:"MinCanAllocateCreditAmount,omitempty" xml:"MinCanAllocateCreditAmount,omitempty"`
	// example:
	//
	// 2684210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// CC706AAC-75A6-55B5-9AB7-7D171C6C7655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 26842
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s GetFundAccountCanAllocateCreditAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanAllocateCreditAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetEcid() *string {
	return s.Ecid
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetEcidAllocatedCreditAmount() *string {
	return s.EcidAllocatedCreditAmount
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetEcidCreditAmount() *string {
	return s.EcidCreditAmount
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetFundAccountEcid() *string {
	return s.FundAccountEcid
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetFundAccountName() *string {
	return s.FundAccountName
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetFundAccountOwnerAccountId() *int64 {
	return s.FundAccountOwnerAccountId
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetMaxCanAllocateCreditAmount() *string {
	return s.MaxCanAllocateCreditAmount
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetMinCanAllocateCreditAmount() *string {
	return s.MinCanAllocateCreditAmount
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetNbid() *string {
	return s.Nbid
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) GetSite() *string {
	return s.Site
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetEcid(v string) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.Ecid = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetEcidAllocatedCreditAmount(v string) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.EcidAllocatedCreditAmount = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetEcidCreditAmount(v string) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.EcidCreditAmount = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetFundAccountEcid(v string) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.FundAccountEcid = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetFundAccountId(v int64) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetFundAccountName(v string) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.FundAccountName = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetFundAccountOwnerAccountId(v int64) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetMaxCanAllocateCreditAmount(v string) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.MaxCanAllocateCreditAmount = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetMetadata(v interface{}) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetMinCanAllocateCreditAmount(v string) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.MinCanAllocateCreditAmount = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetNbid(v string) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.Nbid = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetRequestId(v string) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) SetSite(v string) *GetFundAccountCanAllocateCreditAmountResponseBody {
	s.Site = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFundAccountPayRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListFundAccountPayRelationResponseBody
	GetCurrentPage() *int32
	SetData(v []*ListFundAccountPayRelationResponseBodyData) *ListFundAccountPayRelationResponseBody
	GetData() []*ListFundAccountPayRelationResponseBodyData
	SetMetadata(v interface{}) *ListFundAccountPayRelationResponseBody
	GetMetadata() interface{}
	SetPageSize(v int32) *ListFundAccountPayRelationResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListFundAccountPayRelationResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListFundAccountPayRelationResponseBody
	GetTotalCount() *int32
}

type ListFundAccountPayRelationResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data list.
	Data []*ListFundAccountPayRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response metadata.
	//
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFundAccountPayRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFundAccountPayRelationResponseBody) GoString() string {
	return s.String()
}

func (s *ListFundAccountPayRelationResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFundAccountPayRelationResponseBody) GetData() []*ListFundAccountPayRelationResponseBodyData {
	return s.Data
}

func (s *ListFundAccountPayRelationResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *ListFundAccountPayRelationResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFundAccountPayRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFundAccountPayRelationResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFundAccountPayRelationResponseBody) SetCurrentPage(v int32) *ListFundAccountPayRelationResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBody) SetData(v []*ListFundAccountPayRelationResponseBodyData) *ListFundAccountPayRelationResponseBody {
	s.Data = v
	return s
}

func (s *ListFundAccountPayRelationResponseBody) SetMetadata(v interface{}) *ListFundAccountPayRelationResponseBody {
	s.Metadata = v
	return s
}

func (s *ListFundAccountPayRelationResponseBody) SetPageSize(v int32) *ListFundAccountPayRelationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBody) SetRequestId(v string) *ListFundAccountPayRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBody) SetTotalCount(v int32) *ListFundAccountPayRelationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFundAccountPayRelationResponseBodyData struct {
	// The account ID of the user associated with the payment relationship, that is, the account that uses this account for payment.
	//
	// example:
	//
	// 32812132121
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The user name.
	//
	// example:
	//
	// 云某的名称
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The enterprise entity ID of the user associated with the payment relationship.
	//
	// example:
	//
	// 213231232
	Ecid *string `json:"Ecid,omitempty" xml:"Ecid,omitempty"`
	// The time when the payment relationship takes effect.
	//
	// example:
	//
	// 2024-12-01 12:00:10
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The account ID.
	//
	// example:
	//
	// 123231213
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// The Alibaba Cloud account ID of the account owner.
	//
	// example:
	//
	// 312328912
	FundAccountOwnerAccountId *string `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// The time when the payment relationship expires.
	//
	// example:
	//
	// 2025-01-01 12:12:12
	IneffectiveTime *string `json:"IneffectiveTime,omitempty" xml:"IneffectiveTime,omitempty"`
	// The primary marketplace.
	//
	// example:
	//
	// 2684210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// The operator name.
	//
	// When the operator type is aliyun_pk, the operator name is the Alibaba Cloud nickname.
	//
	// When the operator type is system, the operator name is "Alibaba Cloud assistant".
	//
	// example:
	//
	// 云某的名称
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The authorized operator.
	//
	// When the operator type is aliyun_pk, operatorNo is the Alibaba Cloud account ID.
	//
	// example:
	//
	// 1232343423
	OperatorNo *string `json:"OperatorNo,omitempty" xml:"OperatorNo,omitempty"`
	// The type of the authorized operator.
	//
	// aliyun_pk: user.
	//
	// system: Alibaba Cloud system.
	//
	// example:
	//
	// aliyun_pk
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// The relationship type, which can be collection relationship or payment relationship.
	//
	// example:
	//
	// PAYMENT
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The site.
	//
	// example:
	//
	// 26842
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The relationship status.
	//
	// valid: valid.
	//
	// expired: invalid.
	//
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFundAccountPayRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFundAccountPayRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFundAccountPayRelationResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *ListFundAccountPayRelationResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *ListFundAccountPayRelationResponseBodyData) GetEcid() *string {
	return s.Ecid
}

func (s *ListFundAccountPayRelationResponseBodyData) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ListFundAccountPayRelationResponseBodyData) GetFundAccountId() *string {
	return s.FundAccountId
}

func (s *ListFundAccountPayRelationResponseBodyData) GetFundAccountOwnerAccountId() *string {
	return s.FundAccountOwnerAccountId
}

func (s *ListFundAccountPayRelationResponseBodyData) GetIneffectiveTime() *string {
	return s.IneffectiveTime
}

func (s *ListFundAccountPayRelationResponseBodyData) GetNbid() *string {
	return s.Nbid
}

func (s *ListFundAccountPayRelationResponseBodyData) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ListFundAccountPayRelationResponseBodyData) GetOperatorNo() *string {
	return s.OperatorNo
}

func (s *ListFundAccountPayRelationResponseBodyData) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ListFundAccountPayRelationResponseBodyData) GetRelationType() *string {
	return s.RelationType
}

func (s *ListFundAccountPayRelationResponseBodyData) GetSite() *string {
	return s.Site
}

func (s *ListFundAccountPayRelationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListFundAccountPayRelationResponseBodyData) SetAccountId(v string) *ListFundAccountPayRelationResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetAccountName(v string) *ListFundAccountPayRelationResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetEcid(v string) *ListFundAccountPayRelationResponseBodyData {
	s.Ecid = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetEffectiveTime(v string) *ListFundAccountPayRelationResponseBodyData {
	s.EffectiveTime = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetFundAccountId(v string) *ListFundAccountPayRelationResponseBodyData {
	s.FundAccountId = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetFundAccountOwnerAccountId(v string) *ListFundAccountPayRelationResponseBodyData {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetIneffectiveTime(v string) *ListFundAccountPayRelationResponseBodyData {
	s.IneffectiveTime = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetNbid(v string) *ListFundAccountPayRelationResponseBodyData {
	s.Nbid = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetOperatorName(v string) *ListFundAccountPayRelationResponseBodyData {
	s.OperatorName = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetOperatorNo(v string) *ListFundAccountPayRelationResponseBodyData {
	s.OperatorNo = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetOperatorType(v string) *ListFundAccountPayRelationResponseBodyData {
	s.OperatorType = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetRelationType(v string) *ListFundAccountPayRelationResponseBodyData {
	s.RelationType = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetSite(v string) *ListFundAccountPayRelationResponseBodyData {
	s.Site = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) SetStatus(v string) *ListFundAccountPayRelationResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListFundAccountPayRelationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

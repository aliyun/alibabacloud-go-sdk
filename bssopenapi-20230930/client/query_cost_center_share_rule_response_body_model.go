// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterShareRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryCostCenterShareRuleResponseBodyData) *QueryCostCenterShareRuleResponseBody
	GetData() []*QueryCostCenterShareRuleResponseBodyData
	SetMaxResults(v int32) *QueryCostCenterShareRuleResponseBody
	GetMaxResults() *int32
	SetMetadata(v interface{}) *QueryCostCenterShareRuleResponseBody
	GetMetadata() interface{}
	SetNextToken(v string) *QueryCostCenterShareRuleResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryCostCenterShareRuleResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryCostCenterShareRuleResponseBody
	GetTotalCount() *int32
}

type QueryCostCenterShareRuleResponseBody struct {
	Data []*QueryCostCenterShareRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6NH0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryCostCenterShareRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterShareRuleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostCenterShareRuleResponseBody) GetData() []*QueryCostCenterShareRuleResponseBodyData {
	return s.Data
}

func (s *QueryCostCenterShareRuleResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryCostCenterShareRuleResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *QueryCostCenterShareRuleResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryCostCenterShareRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCostCenterShareRuleResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryCostCenterShareRuleResponseBody) SetData(v []*QueryCostCenterShareRuleResponseBodyData) *QueryCostCenterShareRuleResponseBody {
	s.Data = v
	return s
}

func (s *QueryCostCenterShareRuleResponseBody) SetMaxResults(v int32) *QueryCostCenterShareRuleResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBody) SetMetadata(v interface{}) *QueryCostCenterShareRuleResponseBody {
	s.Metadata = v
	return s
}

func (s *QueryCostCenterShareRuleResponseBody) SetNextToken(v string) *QueryCostCenterShareRuleResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBody) SetRequestId(v string) *QueryCostCenterShareRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBody) SetTotalCount(v int32) *QueryCostCenterShareRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCostCenterShareRuleResponseBodyData struct {
	FromCostCenterShareRuleDetails []*QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails `json:"FromCostCenterShareRuleDetails,omitempty" xml:"FromCostCenterShareRuleDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 1529600453335198
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// example:
	//
	// 1826
	ShareRuleId *int64 `json:"ShareRuleId,omitempty" xml:"ShareRuleId,omitempty"`
	// example:
	//
	// test
	ShareRuleName *string `json:"ShareRuleName,omitempty" xml:"ShareRuleName,omitempty"`
	// example:
	//
	// AVERAGE
	ShareRuleType                *string                                                                 `json:"ShareRuleType,omitempty" xml:"ShareRuleType,omitempty"`
	ToCostCenterShareRuleDetails []*QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails `json:"ToCostCenterShareRuleDetails,omitempty" xml:"ToCostCenterShareRuleDetails,omitempty" type:"Repeated"`
}

func (s QueryCostCenterShareRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterShareRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCostCenterShareRuleResponseBodyData) GetFromCostCenterShareRuleDetails() []*QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails {
	return s.FromCostCenterShareRuleDetails
}

func (s *QueryCostCenterShareRuleResponseBodyData) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostCenterShareRuleResponseBodyData) GetShareRuleId() *int64 {
	return s.ShareRuleId
}

func (s *QueryCostCenterShareRuleResponseBodyData) GetShareRuleName() *string {
	return s.ShareRuleName
}

func (s *QueryCostCenterShareRuleResponseBodyData) GetShareRuleType() *string {
	return s.ShareRuleType
}

func (s *QueryCostCenterShareRuleResponseBodyData) GetToCostCenterShareRuleDetails() []*QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	return s.ToCostCenterShareRuleDetails
}

func (s *QueryCostCenterShareRuleResponseBodyData) SetFromCostCenterShareRuleDetails(v []*QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) *QueryCostCenterShareRuleResponseBodyData {
	s.FromCostCenterShareRuleDetails = v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyData) SetOwnerAccountId(v int64) *QueryCostCenterShareRuleResponseBodyData {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyData) SetShareRuleId(v int64) *QueryCostCenterShareRuleResponseBodyData {
	s.ShareRuleId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyData) SetShareRuleName(v string) *QueryCostCenterShareRuleResponseBodyData {
	s.ShareRuleName = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyData) SetShareRuleType(v string) *QueryCostCenterShareRuleResponseBodyData {
	s.ShareRuleType = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyData) SetToCostCenterShareRuleDetails(v []*QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) *QueryCostCenterShareRuleResponseBodyData {
	s.ToCostCenterShareRuleDetails = v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails struct {
	// example:
	//
	// 970354711219#
	CostCenterCode *string `json:"CostCenterCode,omitempty" xml:"CostCenterCode,omitempty"`
	// example:
	//
	// 637127
	CostCenterId   *int64  `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// example:
	//
	// 2025-07-16 13:49:59
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-07-16 13:49:59
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1529600453335198
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// example:
	//
	// 637537
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
	// example:
	//
	// 583059
	PrevCostCenterId *int64 `json:"PrevCostCenterId,omitempty" xml:"PrevCostCenterId,omitempty"`
	// example:
	//
	// 583050
	RootCostCenterId *int64 `json:"RootCostCenterId,omitempty" xml:"RootCostCenterId,omitempty"`
}

func (s QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) GoString() string {
	return s.String()
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) GetCostCenterCode() *string {
	return s.CostCenterCode
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) GetParentCostCenterId() *int64 {
	return s.ParentCostCenterId
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) GetPrevCostCenterId() *int64 {
	return s.PrevCostCenterId
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) GetRootCostCenterId() *int64 {
	return s.RootCostCenterId
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) SetCostCenterCode(v string) *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails {
	s.CostCenterCode = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) SetCostCenterId(v int64) *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) SetCostCenterName(v string) *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails {
	s.CostCenterName = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) SetGmtCreate(v string) *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails {
	s.GmtCreate = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) SetGmtModified(v string) *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails {
	s.GmtModified = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) SetOwnerAccountId(v int64) *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) SetParentCostCenterId(v int64) *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails {
	s.ParentCostCenterId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) SetPrevCostCenterId(v int64) *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails {
	s.PrevCostCenterId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) SetRootCostCenterId(v int64) *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails {
	s.RootCostCenterId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataFromCostCenterShareRuleDetails) Validate() error {
	return dara.Validate(s)
}

type QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails struct {
	// example:
	//
	// 970354711215#
	CostCenterCode *string `json:"CostCenterCode,omitempty" xml:"CostCenterCode,omitempty"`
	// example:
	//
	// 637127
	CostCenterId   *int64  `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// example:
	//
	// 2025-07-16 13:49:59
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-07-16 13:49:59
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1529600453335198
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// example:
	//
	// 637537
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
	// example:
	//
	// 583055
	PrevCostCenterId *int64 `json:"PrevCostCenterId,omitempty" xml:"PrevCostCenterId,omitempty"`
	// example:
	//
	// 583050
	RootCostCenterId *int64 `json:"RootCostCenterId,omitempty" xml:"RootCostCenterId,omitempty"`
	// example:
	//
	// 0.2
	ShareRatio *float64 `json:"ShareRatio,omitempty" xml:"ShareRatio,omitempty"`
}

func (s QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GoString() string {
	return s.String()
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GetCostCenterCode() *string {
	return s.CostCenterCode
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GetParentCostCenterId() *int64 {
	return s.ParentCostCenterId
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GetPrevCostCenterId() *int64 {
	return s.PrevCostCenterId
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GetRootCostCenterId() *int64 {
	return s.RootCostCenterId
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) GetShareRatio() *float64 {
	return s.ShareRatio
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) SetCostCenterCode(v string) *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	s.CostCenterCode = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) SetCostCenterId(v int64) *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) SetCostCenterName(v string) *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	s.CostCenterName = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) SetGmtCreate(v string) *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	s.GmtCreate = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) SetGmtModified(v string) *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	s.GmtModified = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) SetOwnerAccountId(v int64) *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) SetParentCostCenterId(v int64) *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	s.ParentCostCenterId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) SetPrevCostCenterId(v int64) *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	s.PrevCostCenterId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) SetRootCostCenterId(v int64) *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	s.RootCostCenterId = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) SetShareRatio(v float64) *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails {
	s.ShareRatio = &v
	return s
}

func (s *QueryCostCenterShareRuleResponseBodyDataToCostCenterShareRuleDetails) Validate() error {
	return dara.Validate(s)
}

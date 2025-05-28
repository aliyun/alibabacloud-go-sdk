// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataModuleMapListSpnTypeMapListValue struct {
	FilterModules []*DataModuleMapListSpnTypeMapListValueFilterModules `json:"FilterModules,omitempty" xml:"FilterModules,omitempty" type:"Repeated"`
	ShowModules   []*DataModuleMapListSpnTypeMapListValueShowModules   `json:"ShowModules,omitempty" xml:"ShowModules,omitempty" type:"Repeated"`
}

func (s DataModuleMapListSpnTypeMapListValue) String() string {
	return tea.Prettify(s)
}

func (s DataModuleMapListSpnTypeMapListValue) GoString() string {
	return s.String()
}

func (s *DataModuleMapListSpnTypeMapListValue) SetFilterModules(v []*DataModuleMapListSpnTypeMapListValueFilterModules) *DataModuleMapListSpnTypeMapListValue {
	s.FilterModules = v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValue) SetShowModules(v []*DataModuleMapListSpnTypeMapListValueShowModules) *DataModuleMapListSpnTypeMapListValue {
	s.ShowModules = v
	return s
}

type DataModuleMapListSpnTypeMapListValueFilterModules struct {
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DataModuleMapListSpnTypeMapListValueFilterModules) String() string {
	return tea.Prettify(s)
}

func (s DataModuleMapListSpnTypeMapListValueFilterModules) GoString() string {
	return s.String()
}

func (s *DataModuleMapListSpnTypeMapListValueFilterModules) SetModuleId(v int64) *DataModuleMapListSpnTypeMapListValueFilterModules {
	s.ModuleId = &v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValueFilterModules) SetModuleCode(v string) *DataModuleMapListSpnTypeMapListValueFilterModules {
	s.ModuleCode = &v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValueFilterModules) SetModuleName(v string) *DataModuleMapListSpnTypeMapListValueFilterModules {
	s.ModuleName = &v
	return s
}

type DataModuleMapListSpnTypeMapListValueShowModules struct {
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DataModuleMapListSpnTypeMapListValueShowModules) String() string {
	return tea.Prettify(s)
}

func (s DataModuleMapListSpnTypeMapListValueShowModules) GoString() string {
	return s.String()
}

func (s *DataModuleMapListSpnTypeMapListValueShowModules) SetModuleId(v int64) *DataModuleMapListSpnTypeMapListValueShowModules {
	s.ModuleId = &v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValueShowModules) SetModuleCode(v string) *DataModuleMapListSpnTypeMapListValueShowModules {
	s.ModuleCode = &v
	return s
}

func (s *DataModuleMapListSpnTypeMapListValueShowModules) SetModuleName(v string) *DataModuleMapListSpnTypeMapListValueShowModules {
	s.ModuleName = &v
	return s
}

type DataStepPriceMapValue struct {
	RightClose      *bool   `json:"RightClose,omitempty" xml:"RightClose,omitempty"`
	Min             *string `json:"Min,omitempty" xml:"Min,omitempty"`
	Max             *string `json:"Max,omitempty" xml:"Max,omitempty"`
	Currency        *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	LeftClose       *bool   `json:"LeftClose,omitempty" xml:"LeftClose,omitempty"`
	StepPriceValue  *string `json:"StepPriceValue,omitempty" xml:"StepPriceValue,omitempty"`
	PriceValueType  *string `json:"PriceValueType,omitempty" xml:"PriceValueType,omitempty"`
	PriceValue      *string `json:"PriceValue,omitempty" xml:"PriceValue,omitempty"`
	DeductCycleType *string `json:"DeductCycleType,omitempty" xml:"DeductCycleType,omitempty"`
}

func (s DataStepPriceMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataStepPriceMapValue) GoString() string {
	return s.String()
}

func (s *DataStepPriceMapValue) SetRightClose(v bool) *DataStepPriceMapValue {
	s.RightClose = &v
	return s
}

func (s *DataStepPriceMapValue) SetMin(v string) *DataStepPriceMapValue {
	s.Min = &v
	return s
}

func (s *DataStepPriceMapValue) SetMax(v string) *DataStepPriceMapValue {
	s.Max = &v
	return s
}

func (s *DataStepPriceMapValue) SetCurrency(v string) *DataStepPriceMapValue {
	s.Currency = &v
	return s
}

func (s *DataStepPriceMapValue) SetLeftClose(v bool) *DataStepPriceMapValue {
	s.LeftClose = &v
	return s
}

func (s *DataStepPriceMapValue) SetStepPriceValue(v string) *DataStepPriceMapValue {
	s.StepPriceValue = &v
	return s
}

func (s *DataStepPriceMapValue) SetPriceValueType(v string) *DataStepPriceMapValue {
	s.PriceValueType = &v
	return s
}

func (s *DataStepPriceMapValue) SetPriceValue(v string) *DataStepPriceMapValue {
	s.PriceValue = &v
	return s
}

func (s *DataStepPriceMapValue) SetDeductCycleType(v string) *DataStepPriceMapValue {
	s.DeductCycleType = &v
	return s
}

type AddCouponDeductTagRequest struct {
	CouponId       *string                                    `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIds []*AddCouponDeductTagRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid           *string                                    `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	Tags           []*AddCouponDeductTagRequestTags           `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AddCouponDeductTagRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCouponDeductTagRequest) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagRequest) SetCouponId(v string) *AddCouponDeductTagRequest {
	s.CouponId = &v
	return s
}

func (s *AddCouponDeductTagRequest) SetEcIdAccountIds(v []*AddCouponDeductTagRequestEcIdAccountIds) *AddCouponDeductTagRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *AddCouponDeductTagRequest) SetNbid(v string) *AddCouponDeductTagRequest {
	s.Nbid = &v
	return s
}

func (s *AddCouponDeductTagRequest) SetTags(v []*AddCouponDeductTagRequestTags) *AddCouponDeductTagRequest {
	s.Tags = v
	return s
}

type AddCouponDeductTagRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s AddCouponDeductTagRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s AddCouponDeductTagRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagRequestEcIdAccountIds) SetAccountIds(v []*int64) *AddCouponDeductTagRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *AddCouponDeductTagRequestEcIdAccountIds) SetEcId(v string) *AddCouponDeductTagRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type AddCouponDeductTagRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddCouponDeductTagRequestTags) String() string {
	return tea.Prettify(s)
}

func (s AddCouponDeductTagRequestTags) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagRequestTags) SetKey(v string) *AddCouponDeductTagRequestTags {
	s.Key = &v
	return s
}

func (s *AddCouponDeductTagRequestTags) SetValue(v string) *AddCouponDeductTagRequestTags {
	s.Value = &v
	return s
}

type AddCouponDeductTagShrinkRequest struct {
	CouponId             *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	TagsShrink           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddCouponDeductTagShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCouponDeductTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagShrinkRequest) SetCouponId(v string) *AddCouponDeductTagShrinkRequest {
	s.CouponId = &v
	return s
}

func (s *AddCouponDeductTagShrinkRequest) SetEcIdAccountIdsShrink(v string) *AddCouponDeductTagShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *AddCouponDeductTagShrinkRequest) SetNbid(v string) *AddCouponDeductTagShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *AddCouponDeductTagShrinkRequest) SetTagsShrink(v string) *AddCouponDeductTagShrinkRequest {
	s.TagsShrink = &v
	return s
}

type AddCouponDeductTagResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCouponDeductTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCouponDeductTagResponseBody) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagResponseBody) SetData(v bool) *AddCouponDeductTagResponseBody {
	s.Data = &v
	return s
}

func (s *AddCouponDeductTagResponseBody) SetRequestId(v string) *AddCouponDeductTagResponseBody {
	s.RequestId = &v
	return s
}

type AddCouponDeductTagResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCouponDeductTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCouponDeductTagResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCouponDeductTagResponse) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagResponse) SetHeaders(v map[string]*string) *AddCouponDeductTagResponse {
	s.Headers = v
	return s
}

func (s *AddCouponDeductTagResponse) SetStatusCode(v int32) *AddCouponDeductTagResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCouponDeductTagResponse) SetBody(v *AddCouponDeductTagResponseBody) *AddCouponDeductTagResponse {
	s.Body = v
	return s
}

type CancelFundAccountLowAvailableAmountAlarmRequest struct {
	// example:
	//
	// 123321123
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s CancelFundAccountLowAvailableAmountAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelFundAccountLowAvailableAmountAlarmRequest) GoString() string {
	return s.String()
}

func (s *CancelFundAccountLowAvailableAmountAlarmRequest) SetFundAccountId(v int64) *CancelFundAccountLowAvailableAmountAlarmRequest {
	s.FundAccountId = &v
	return s
}

type CancelFundAccountLowAvailableAmountAlarmResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// DFC1F7F9-3BA9-BA4D-2F2E653
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelFundAccountLowAvailableAmountAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelFundAccountLowAvailableAmountAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponseBody) SetData(v bool) *CancelFundAccountLowAvailableAmountAlarmResponseBody {
	s.Data = &v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponseBody) SetMetadata(v interface{}) *CancelFundAccountLowAvailableAmountAlarmResponseBody {
	s.Metadata = v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponseBody) SetRequestId(v string) *CancelFundAccountLowAvailableAmountAlarmResponseBody {
	s.RequestId = &v
	return s
}

type CancelFundAccountLowAvailableAmountAlarmResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelFundAccountLowAvailableAmountAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelFundAccountLowAvailableAmountAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelFundAccountLowAvailableAmountAlarmResponse) GoString() string {
	return s.String()
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponse) SetHeaders(v map[string]*string) *CancelFundAccountLowAvailableAmountAlarmResponse {
	s.Headers = v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponse) SetStatusCode(v int32) *CancelFundAccountLowAvailableAmountAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelFundAccountLowAvailableAmountAlarmResponse) SetBody(v *CancelFundAccountLowAvailableAmountAlarmResponseBody) *CancelFundAccountLowAvailableAmountAlarmResponse {
	s.Body = v
	return s
}

type CreateCostCenterRequest struct {
	// This parameter is required.
	CostCenterEntityList []*CreateCostCenterRequestCostCenterEntityList `json:"CostCenterEntityList,omitempty" xml:"CostCenterEntityList,omitempty" type:"Repeated"`
	// example:
	//
	// 2084210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s CreateCostCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCostCenterRequest) GoString() string {
	return s.String()
}

func (s *CreateCostCenterRequest) SetCostCenterEntityList(v []*CreateCostCenterRequestCostCenterEntityList) *CreateCostCenterRequest {
	s.CostCenterEntityList = v
	return s
}

func (s *CreateCostCenterRequest) SetNbid(v string) *CreateCostCenterRequest {
	s.Nbid = &v
	return s
}

type CreateCostCenterRequestCostCenterEntityList struct {
	// This parameter is required.
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
}

func (s CreateCostCenterRequestCostCenterEntityList) String() string {
	return tea.Prettify(s)
}

func (s CreateCostCenterRequestCostCenterEntityList) GoString() string {
	return s.String()
}

func (s *CreateCostCenterRequestCostCenterEntityList) SetCostCenterName(v string) *CreateCostCenterRequestCostCenterEntityList {
	s.CostCenterName = &v
	return s
}

func (s *CreateCostCenterRequestCostCenterEntityList) SetOwnerAccountId(v int64) *CreateCostCenterRequestCostCenterEntityList {
	s.OwnerAccountId = &v
	return s
}

func (s *CreateCostCenterRequestCostCenterEntityList) SetParentCostCenterId(v int64) *CreateCostCenterRequestCostCenterEntityList {
	s.ParentCostCenterId = &v
	return s
}

type CreateCostCenterShrinkRequest struct {
	// This parameter is required.
	CostCenterEntityListShrink *string `json:"CostCenterEntityList,omitempty" xml:"CostCenterEntityList,omitempty"`
	// example:
	//
	// 2084210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s CreateCostCenterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCostCenterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCostCenterShrinkRequest) SetCostCenterEntityListShrink(v string) *CreateCostCenterShrinkRequest {
	s.CostCenterEntityListShrink = &v
	return s
}

func (s *CreateCostCenterShrinkRequest) SetNbid(v string) *CreateCostCenterShrinkRequest {
	s.Nbid = &v
	return s
}

type CreateCostCenterResponseBody struct {
	CostCenterDtoList []*CreateCostCenterResponseBodyCostCenterDtoList `json:"CostCenterDtoList,omitempty" xml:"CostCenterDtoList,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// C1BD134E-D914-6AE0-1901-AEB2A99FA205
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCostCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCostCenterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCostCenterResponseBody) SetCostCenterDtoList(v []*CreateCostCenterResponseBodyCostCenterDtoList) *CreateCostCenterResponseBody {
	s.CostCenterDtoList = v
	return s
}

func (s *CreateCostCenterResponseBody) SetMetadata(v interface{}) *CreateCostCenterResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateCostCenterResponseBody) SetRequestId(v string) *CreateCostCenterResponseBody {
	s.RequestId = &v
	return s
}

type CreateCostCenterResponseBodyCostCenterDtoList struct {
	// example:
	//
	// 485938
	CostCenterId   *int64  `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// example:
	//
	// -1
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
}

func (s CreateCostCenterResponseBodyCostCenterDtoList) String() string {
	return tea.Prettify(s)
}

func (s CreateCostCenterResponseBodyCostCenterDtoList) GoString() string {
	return s.String()
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) SetCostCenterId(v int64) *CreateCostCenterResponseBodyCostCenterDtoList {
	s.CostCenterId = &v
	return s
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) SetCostCenterName(v string) *CreateCostCenterResponseBodyCostCenterDtoList {
	s.CostCenterName = &v
	return s
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) SetOwnerAccountId(v int64) *CreateCostCenterResponseBodyCostCenterDtoList {
	s.OwnerAccountId = &v
	return s
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) SetParentCostCenterId(v int64) *CreateCostCenterResponseBodyCostCenterDtoList {
	s.ParentCostCenterId = &v
	return s
}

type CreateCostCenterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCostCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCostCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCostCenterResponse) GoString() string {
	return s.String()
}

func (s *CreateCostCenterResponse) SetHeaders(v map[string]*string) *CreateCostCenterResponse {
	s.Headers = v
	return s
}

func (s *CreateCostCenterResponse) SetStatusCode(v int32) *CreateCostCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCostCenterResponse) SetBody(v *CreateCostCenterResponseBody) *CreateCostCenterResponse {
	s.Body = v
	return s
}

type CreateFundAccountPayRelationRequest struct {
	// This parameter is required.
	EcIdAccountIds []*CreateFundAccountPayRelationRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 12332112
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s CreateFundAccountPayRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFundAccountPayRelationRequest) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationRequest) SetEcIdAccountIds(v []*CreateFundAccountPayRelationRequestEcIdAccountIds) *CreateFundAccountPayRelationRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *CreateFundAccountPayRelationRequest) SetFundAccountId(v string) *CreateFundAccountPayRelationRequest {
	s.FundAccountId = &v
	return s
}

func (s *CreateFundAccountPayRelationRequest) SetNbid(v string) *CreateFundAccountPayRelationRequest {
	s.Nbid = &v
	return s
}

type CreateFundAccountPayRelationRequestEcIdAccountIds struct {
	// This parameter is required.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1501603440974415
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s CreateFundAccountPayRelationRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s CreateFundAccountPayRelationRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationRequestEcIdAccountIds) SetAccountIds(v []*int64) *CreateFundAccountPayRelationRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *CreateFundAccountPayRelationRequestEcIdAccountIds) SetEcId(v string) *CreateFundAccountPayRelationRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type CreateFundAccountPayRelationShrinkRequest struct {
	// This parameter is required.
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12332112
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s CreateFundAccountPayRelationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFundAccountPayRelationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationShrinkRequest) SetEcIdAccountIdsShrink(v string) *CreateFundAccountPayRelationShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *CreateFundAccountPayRelationShrinkRequest) SetFundAccountId(v string) *CreateFundAccountPayRelationShrinkRequest {
	s.FundAccountId = &v
	return s
}

func (s *CreateFundAccountPayRelationShrinkRequest) SetNbid(v string) *CreateFundAccountPayRelationShrinkRequest {
	s.Nbid = &v
	return s
}

type CreateFundAccountPayRelationResponseBody struct {
	Data []*CreateFundAccountPayRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFundAccountPayRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFundAccountPayRelationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationResponseBody) SetData(v []*CreateFundAccountPayRelationResponseBodyData) *CreateFundAccountPayRelationResponseBody {
	s.Data = v
	return s
}

func (s *CreateFundAccountPayRelationResponseBody) SetMetadata(v interface{}) *CreateFundAccountPayRelationResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateFundAccountPayRelationResponseBody) SetRequestId(v string) *CreateFundAccountPayRelationResponseBody {
	s.RequestId = &v
	return s
}

type CreateFundAccountPayRelationResponseBodyData struct {
	// example:
	//
	// 1501603440974415
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 12332112
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// Success
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// Successful
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateFundAccountPayRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateFundAccountPayRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationResponseBodyData) SetAccountId(v string) *CreateFundAccountPayRelationResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *CreateFundAccountPayRelationResponseBodyData) SetAccountName(v string) *CreateFundAccountPayRelationResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *CreateFundAccountPayRelationResponseBodyData) SetFundAccountId(v string) *CreateFundAccountPayRelationResponseBodyData {
	s.FundAccountId = &v
	return s
}

func (s *CreateFundAccountPayRelationResponseBodyData) SetResultCode(v string) *CreateFundAccountPayRelationResponseBodyData {
	s.ResultCode = &v
	return s
}

func (s *CreateFundAccountPayRelationResponseBodyData) SetResultMessage(v string) *CreateFundAccountPayRelationResponseBodyData {
	s.ResultMessage = &v
	return s
}

type CreateFundAccountPayRelationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFundAccountPayRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFundAccountPayRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFundAccountPayRelationResponse) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationResponse) SetHeaders(v map[string]*string) *CreateFundAccountPayRelationResponse {
	s.Headers = v
	return s
}

func (s *CreateFundAccountPayRelationResponse) SetStatusCode(v int32) *CreateFundAccountPayRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFundAccountPayRelationResponse) SetBody(v *CreateFundAccountPayRelationResponseBody) *CreateFundAccountPayRelationResponse {
	s.Body = v
	return s
}

type CreateFundAccountTransferRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cash
	FinanceType *string `json:"FinanceType,omitempty" xml:"FinanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123212323
	FromFundAccountId *int64 `json:"FromFundAccountId,omitempty" xml:"FromFundAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 转账的备注
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11112231
	ToFundAccountId *int64 `json:"ToFundAccountId,omitempty" xml:"ToFundAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// grant
	TransferType *string `json:"TransferType,omitempty" xml:"TransferType,omitempty"`
}

func (s CreateFundAccountTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFundAccountTransferRequest) GoString() string {
	return s.String()
}

func (s *CreateFundAccountTransferRequest) SetAmount(v string) *CreateFundAccountTransferRequest {
	s.Amount = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetCurrency(v string) *CreateFundAccountTransferRequest {
	s.Currency = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetFinanceType(v string) *CreateFundAccountTransferRequest {
	s.FinanceType = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetFromFundAccountId(v int64) *CreateFundAccountTransferRequest {
	s.FromFundAccountId = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetRemark(v string) *CreateFundAccountTransferRequest {
	s.Remark = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetToFundAccountId(v int64) *CreateFundAccountTransferRequest {
	s.ToFundAccountId = &v
	return s
}

func (s *CreateFundAccountTransferRequest) SetTransferType(v string) *CreateFundAccountTransferRequest {
	s.TransferType = &v
	return s
}

type CreateFundAccountTransferResponseBody struct {
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 1BB79-5B23-3EA-BB4F-352F93E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFundAccountTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFundAccountTransferResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFundAccountTransferResponseBody) SetMetadata(v interface{}) *CreateFundAccountTransferResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateFundAccountTransferResponseBody) SetRequestId(v string) *CreateFundAccountTransferResponseBody {
	s.RequestId = &v
	return s
}

type CreateFundAccountTransferResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFundAccountTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFundAccountTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFundAccountTransferResponse) GoString() string {
	return s.String()
}

func (s *CreateFundAccountTransferResponse) SetHeaders(v map[string]*string) *CreateFundAccountTransferResponse {
	s.Headers = v
	return s
}

func (s *CreateFundAccountTransferResponse) SetStatusCode(v int32) *CreateFundAccountTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFundAccountTransferResponse) SetBody(v *CreateFundAccountTransferResponseBody) *CreateFundAccountTransferResponse {
	s.Body = v
	return s
}

type CreateReportDefinitionRequest struct {
	// example:
	//
	// 2025-05
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// example:
	//
	// project
	McProject *string `json:"McProject,omitempty" xml:"McProject,omitempty"`
	// example:
	//
	// table
	McTableName *string `json:"McTableName,omitempty" xml:"McTableName,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// sh-bill
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// 1234567812345678
	OssBucketOwnerAccountId *int64 `json:"OssBucketOwnerAccountId,omitempty" xml:"OssBucketOwnerAccountId,omitempty"`
	// example:
	//
	// bill/
	OssBucketPath *string `json:"OssBucketPath,omitempty" xml:"OssBucketPath,omitempty"`
	// example:
	//
	// OSS
	ReportSourceType *string `json:"ReportSourceType,omitempty" xml:"ReportSourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BillingItemDetailForBillingPeriod
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s CreateReportDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReportDefinitionRequest) GoString() string {
	return s.String()
}

func (s *CreateReportDefinitionRequest) SetBeginBillingCycle(v string) *CreateReportDefinitionRequest {
	s.BeginBillingCycle = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetMcProject(v string) *CreateReportDefinitionRequest {
	s.McProject = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetMcTableName(v string) *CreateReportDefinitionRequest {
	s.McTableName = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetNbid(v string) *CreateReportDefinitionRequest {
	s.Nbid = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetOssBucketName(v string) *CreateReportDefinitionRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetOssBucketOwnerAccountId(v int64) *CreateReportDefinitionRequest {
	s.OssBucketOwnerAccountId = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetOssBucketPath(v string) *CreateReportDefinitionRequest {
	s.OssBucketPath = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetReportSourceType(v string) *CreateReportDefinitionRequest {
	s.ReportSourceType = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetReportType(v string) *CreateReportDefinitionRequest {
	s.ReportType = &v
	return s
}

type CreateReportDefinitionResponseBody struct {
	// example:
	//
	// 2025-05
	BeginBillingCycle *string     `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	Metadata          interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// sh-bill
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// 1234567812345678
	OssBucketOwnerAccountId *int64 `json:"OssBucketOwnerAccountId,omitempty" xml:"OssBucketOwnerAccountId,omitempty"`
	// example:
	//
	// bill/
	OssBucketPath *string `json:"OssBucketPath,omitempty" xml:"OssBucketPath,omitempty"`
	// example:
	//
	// OSS
	ReportSourceName *string `json:"ReportSourceName,omitempty" xml:"ReportSourceName,omitempty"`
	// example:
	//
	// OSS
	ReportSourceType *string `json:"ReportSourceType,omitempty" xml:"ReportSourceType,omitempty"`
	// example:
	//
	// 123123
	ReportTaskId *int64 `json:"ReportTaskId,omitempty" xml:"ReportTaskId,omitempty"`
	// example:
	//
	// BillingItemDetailForBillingPeriod
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// example:
	//
	// 340CAB45-0637-5875-9BE4-EFD5750F6BA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2025-05-21 10:36:31
	SubscribeCreateTime *string `json:"SubscribeCreateTime,omitempty" xml:"SubscribeCreateTime,omitempty"`
}

func (s CreateReportDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateReportDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReportDefinitionResponseBody) SetBeginBillingCycle(v string) *CreateReportDefinitionResponseBody {
	s.BeginBillingCycle = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetMetadata(v interface{}) *CreateReportDefinitionResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetOssBucketName(v string) *CreateReportDefinitionResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetOssBucketOwnerAccountId(v int64) *CreateReportDefinitionResponseBody {
	s.OssBucketOwnerAccountId = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetOssBucketPath(v string) *CreateReportDefinitionResponseBody {
	s.OssBucketPath = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportSourceName(v string) *CreateReportDefinitionResponseBody {
	s.ReportSourceName = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportSourceType(v string) *CreateReportDefinitionResponseBody {
	s.ReportSourceType = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportTaskId(v int64) *CreateReportDefinitionResponseBody {
	s.ReportTaskId = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetReportType(v string) *CreateReportDefinitionResponseBody {
	s.ReportType = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetRequestId(v string) *CreateReportDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReportDefinitionResponseBody) SetSubscribeCreateTime(v string) *CreateReportDefinitionResponseBody {
	s.SubscribeCreateTime = &v
	return s
}

type CreateReportDefinitionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReportDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReportDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReportDefinitionResponse) GoString() string {
	return s.String()
}

func (s *CreateReportDefinitionResponse) SetHeaders(v map[string]*string) *CreateReportDefinitionResponse {
	s.Headers = v
	return s
}

func (s *CreateReportDefinitionResponse) SetStatusCode(v int32) *CreateReportDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReportDefinitionResponse) SetBody(v *CreateReportDefinitionResponseBody) *CreateReportDefinitionResponse {
	s.Body = v
	return s
}

type DeleteCostCenterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 485938
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s DeleteCostCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCostCenterRequest) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterRequest) SetCostCenterId(v int64) *DeleteCostCenterRequest {
	s.CostCenterId = &v
	return s
}

func (s *DeleteCostCenterRequest) SetNbid(v string) *DeleteCostCenterRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteCostCenterRequest) SetOwnerAccountId(v int64) *DeleteCostCenterRequest {
	s.OwnerAccountId = &v
	return s
}

type DeleteCostCenterResponseBody struct {
	// example:
	//
	// 485938
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// example:
	//
	// C1BD134E-D914-6AE0-1901-AEB2A99FA205
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCostCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCostCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterResponseBody) SetCostCenterId(v int64) *DeleteCostCenterResponseBody {
	s.CostCenterId = &v
	return s
}

func (s *DeleteCostCenterResponseBody) SetIsSuccess(v bool) *DeleteCostCenterResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteCostCenterResponseBody) SetMetadata(v interface{}) *DeleteCostCenterResponseBody {
	s.Metadata = v
	return s
}

func (s *DeleteCostCenterResponseBody) SetOwnerAccountId(v int64) *DeleteCostCenterResponseBody {
	s.OwnerAccountId = &v
	return s
}

func (s *DeleteCostCenterResponseBody) SetRequestId(v string) *DeleteCostCenterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCostCenterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCostCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCostCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCostCenterResponse) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterResponse) SetHeaders(v map[string]*string) *DeleteCostCenterResponse {
	s.Headers = v
	return s
}

func (s *DeleteCostCenterResponse) SetStatusCode(v int32) *DeleteCostCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCostCenterResponse) SetBody(v *DeleteCostCenterResponseBody) *DeleteCostCenterResponse {
	s.Body = v
	return s
}

type DeleteCouponDeductTagRequest struct {
	CouponId       *string                                       `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIds []*DeleteCouponDeductTagRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid           *string                                       `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	TagKeys        []*string                                     `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s DeleteCouponDeductTagRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCouponDeductTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteCouponDeductTagRequest) SetCouponId(v string) *DeleteCouponDeductTagRequest {
	s.CouponId = &v
	return s
}

func (s *DeleteCouponDeductTagRequest) SetEcIdAccountIds(v []*DeleteCouponDeductTagRequestEcIdAccountIds) *DeleteCouponDeductTagRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *DeleteCouponDeductTagRequest) SetNbid(v string) *DeleteCouponDeductTagRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteCouponDeductTagRequest) SetTagKeys(v []*string) *DeleteCouponDeductTagRequest {
	s.TagKeys = v
	return s
}

type DeleteCouponDeductTagRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DeleteCouponDeductTagRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s DeleteCouponDeductTagRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *DeleteCouponDeductTagRequestEcIdAccountIds) SetAccountIds(v []*int64) *DeleteCouponDeductTagRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *DeleteCouponDeductTagRequestEcIdAccountIds) SetEcId(v string) *DeleteCouponDeductTagRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type DeleteCouponDeductTagShrinkRequest struct {
	CouponId             *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	TagKeysShrink        *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s DeleteCouponDeductTagShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCouponDeductTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteCouponDeductTagShrinkRequest) SetCouponId(v string) *DeleteCouponDeductTagShrinkRequest {
	s.CouponId = &v
	return s
}

func (s *DeleteCouponDeductTagShrinkRequest) SetEcIdAccountIdsShrink(v string) *DeleteCouponDeductTagShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *DeleteCouponDeductTagShrinkRequest) SetNbid(v string) *DeleteCouponDeductTagShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteCouponDeductTagShrinkRequest) SetTagKeysShrink(v string) *DeleteCouponDeductTagShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

type DeleteCouponDeductTagResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCouponDeductTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCouponDeductTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCouponDeductTagResponseBody) SetData(v bool) *DeleteCouponDeductTagResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCouponDeductTagResponseBody) SetRequestId(v string) *DeleteCouponDeductTagResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCouponDeductTagResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCouponDeductTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCouponDeductTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCouponDeductTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteCouponDeductTagResponse) SetHeaders(v map[string]*string) *DeleteCouponDeductTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteCouponDeductTagResponse) SetStatusCode(v int32) *DeleteCouponDeductTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCouponDeductTagResponse) SetBody(v *DeleteCouponDeductTagResponseBody) *DeleteCouponDeductTagResponse {
	s.Body = v
	return s
}

type DeleteReportDefinitionRequest struct {
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123123
	ReportTaskId *int64 `json:"ReportTaskId,omitempty" xml:"ReportTaskId,omitempty"`
}

func (s DeleteReportDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReportDefinitionRequest) GoString() string {
	return s.String()
}

func (s *DeleteReportDefinitionRequest) SetNbid(v string) *DeleteReportDefinitionRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteReportDefinitionRequest) SetReportTaskId(v int64) *DeleteReportDefinitionRequest {
	s.ReportTaskId = &v
	return s
}

type DeleteReportDefinitionResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReportDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteReportDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReportDefinitionResponseBody) SetData(v bool) *DeleteReportDefinitionResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteReportDefinitionResponseBody) SetMetadata(v interface{}) *DeleteReportDefinitionResponseBody {
	s.Metadata = v
	return s
}

func (s *DeleteReportDefinitionResponseBody) SetRequestId(v string) *DeleteReportDefinitionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteReportDefinitionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReportDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReportDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteReportDefinitionResponse) GoString() string {
	return s.String()
}

func (s *DeleteReportDefinitionResponse) SetHeaders(v map[string]*string) *DeleteReportDefinitionResponse {
	s.Headers = v
	return s
}

func (s *DeleteReportDefinitionResponse) SetStatusCode(v int32) *DeleteReportDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReportDefinitionResponse) SetBody(v *DeleteReportDefinitionResponseBody) *DeleteReportDefinitionResponse {
	s.Body = v
	return s
}

type DescribeCouponRequest struct {
	// example:
	//
	// 351430260343
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// example:
	//
	// 554863270150
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// example:
	//
	// CERTAIN
	CouponType *string `json:"CouponType,omitempty" xml:"CouponType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage    *int32                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds []*DescribeCouponRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1708423156000
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// example:
	//
	// 1684750028000
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// example:
	//
	// 1708423156000
	ExpireEndDate *int64 `json:"ExpireEndDate,omitempty" xml:"ExpireEndDate,omitempty"`
	// example:
	//
	// 1684750028000
	ExpireStartDate *int64 `json:"ExpireStartDate,omitempty" xml:"ExpireStartDate,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCouponRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponRequest) GoString() string {
	return s.String()
}

func (s *DescribeCouponRequest) SetCouponId(v int64) *DescribeCouponRequest {
	s.CouponId = &v
	return s
}

func (s *DescribeCouponRequest) SetCouponNo(v string) *DescribeCouponRequest {
	s.CouponNo = &v
	return s
}

func (s *DescribeCouponRequest) SetCouponType(v string) *DescribeCouponRequest {
	s.CouponType = &v
	return s
}

func (s *DescribeCouponRequest) SetCurrentPage(v int32) *DescribeCouponRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponRequest) SetEcIdAccountIds(v []*DescribeCouponRequestEcIdAccountIds) *DescribeCouponRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *DescribeCouponRequest) SetEffectiveEndTime(v int64) *DescribeCouponRequest {
	s.EffectiveEndTime = &v
	return s
}

func (s *DescribeCouponRequest) SetEffectiveStartTime(v int64) *DescribeCouponRequest {
	s.EffectiveStartTime = &v
	return s
}

func (s *DescribeCouponRequest) SetExpireEndDate(v int64) *DescribeCouponRequest {
	s.ExpireEndDate = &v
	return s
}

func (s *DescribeCouponRequest) SetExpireStartDate(v int64) *DescribeCouponRequest {
	s.ExpireStartDate = &v
	return s
}

func (s *DescribeCouponRequest) SetNbid(v string) *DescribeCouponRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeCouponRequest) SetPageSize(v int32) *DescribeCouponRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponRequest) SetStatus(v string) *DescribeCouponRequest {
	s.Status = &v
	return s
}

type DescribeCouponRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1501603440974415
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DescribeCouponRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *DescribeCouponRequestEcIdAccountIds) SetAccountIds(v []*int64) *DescribeCouponRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *DescribeCouponRequestEcIdAccountIds) SetEcId(v string) *DescribeCouponRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type DescribeCouponShrinkRequest struct {
	// example:
	//
	// 351430260343
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// example:
	//
	// 554863270150
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// example:
	//
	// CERTAIN
	CouponType *string `json:"CouponType,omitempty" xml:"CouponType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// example:
	//
	// 1708423156000
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// example:
	//
	// 1684750028000
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// example:
	//
	// 1708423156000
	ExpireEndDate *int64 `json:"ExpireEndDate,omitempty" xml:"ExpireEndDate,omitempty"`
	// example:
	//
	// 1684750028000
	ExpireStartDate *int64 `json:"ExpireStartDate,omitempty" xml:"ExpireStartDate,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCouponShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCouponShrinkRequest) SetCouponId(v int64) *DescribeCouponShrinkRequest {
	s.CouponId = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetCouponNo(v string) *DescribeCouponShrinkRequest {
	s.CouponNo = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetCouponType(v string) *DescribeCouponShrinkRequest {
	s.CouponType = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetCurrentPage(v int32) *DescribeCouponShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetEcIdAccountIdsShrink(v string) *DescribeCouponShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetEffectiveEndTime(v int64) *DescribeCouponShrinkRequest {
	s.EffectiveEndTime = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetEffectiveStartTime(v int64) *DescribeCouponShrinkRequest {
	s.EffectiveStartTime = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetExpireEndDate(v int64) *DescribeCouponShrinkRequest {
	s.ExpireEndDate = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetExpireStartDate(v int64) *DescribeCouponShrinkRequest {
	s.ExpireStartDate = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetNbid(v string) *DescribeCouponShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetPageSize(v int32) *DescribeCouponShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetStatus(v string) *DescribeCouponShrinkRequest {
	s.Status = &v
	return s
}

type DescribeCouponResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*DescribeCouponResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// C880B065-A781-4F19-B6DD-3E0E3B715C64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCouponResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCouponResponseBody) SetCurrentPage(v int32) *DescribeCouponResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponResponseBody) SetData(v []*DescribeCouponResponseBodyData) *DescribeCouponResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCouponResponseBody) SetPageSize(v int32) *DescribeCouponResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponResponseBody) SetRequestId(v string) *DescribeCouponResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCouponResponseBody) SetTotalCount(v int32) *DescribeCouponResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCouponResponseBodyData struct {
	// example:
	//
	// 9929.750000
	Amount        *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	CertainAmount *string `json:"CertainAmount,omitempty" xml:"CertainAmount,omitempty"`
	// example:
	//
	// 59243658
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// example:
	//
	// 731074910070
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// example:
	//
	// CERTAIN
	CouponType     *string `json:"CouponType,omitempty" xml:"CouponType,omitempty"`
	CouponTypeName *string `json:"CouponTypeName,omitempty" xml:"CouponTypeName,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 2021-03-06T15:12Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2021-03-02T15:12Z
	GmtCreate     *string   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ItemNames     []*string `json:"ItemNames,omitempty" xml:"ItemNames,omitempty" type:"Repeated"`
	MoneyLimit    *string   `json:"MoneyLimit,omitempty" xml:"MoneyLimit,omitempty"`
	OrderTimeRule *string   `json:"OrderTimeRule,omitempty" xml:"OrderTimeRule,omitempty"`
	// example:
	//
	// 100.00
	RemainAmount *string                                       `json:"RemainAmount,omitempty" xml:"RemainAmount,omitempty"`
	Remark       *string                                       `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ShareUidList []*DescribeCouponResponseBodyDataShareUidList `json:"ShareUidList,omitempty" xml:"ShareUidList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	ShowSetDeductTagButton *bool `json:"ShowSetDeductTagButton,omitempty" xml:"ShowSetDeductTagButton,omitempty"`
	// example:
	//
	// CHINA
	Site     *string `json:"Site,omitempty" xml:"Site,omitempty"`
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// example:
	//
	// 2021-03-02T15:12Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1902671110151254
	SuitAccount *string `json:"SuitAccount,omitempty" xml:"SuitAccount,omitempty"`
	// example:
	//
	// all
	SuitItemType *string `json:"SuitItemType,omitempty" xml:"SuitItemType,omitempty"`
	// example:
	//
	// UNIVERSAL
	UniversalType *string   `json:"UniversalType,omitempty" xml:"UniversalType,omitempty"`
	YhOrderTypes  []*string `json:"YhOrderTypes,omitempty" xml:"YhOrderTypes,omitempty" type:"Repeated"`
}

func (s DescribeCouponResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCouponResponseBodyData) SetAmount(v string) *DescribeCouponResponseBodyData {
	s.Amount = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCertainAmount(v string) *DescribeCouponResponseBodyData {
	s.CertainAmount = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCouponId(v int64) *DescribeCouponResponseBodyData {
	s.CouponId = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCouponNo(v string) *DescribeCouponResponseBodyData {
	s.CouponNo = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCouponType(v string) *DescribeCouponResponseBodyData {
	s.CouponType = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCouponTypeName(v string) *DescribeCouponResponseBodyData {
	s.CouponTypeName = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCurrency(v string) *DescribeCouponResponseBodyData {
	s.Currency = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetEndTime(v string) *DescribeCouponResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetGmtCreate(v string) *DescribeCouponResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetItemNames(v []*string) *DescribeCouponResponseBodyData {
	s.ItemNames = v
	return s
}

func (s *DescribeCouponResponseBodyData) SetMoneyLimit(v string) *DescribeCouponResponseBodyData {
	s.MoneyLimit = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetOrderTimeRule(v string) *DescribeCouponResponseBodyData {
	s.OrderTimeRule = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetRemainAmount(v string) *DescribeCouponResponseBodyData {
	s.RemainAmount = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetRemark(v string) *DescribeCouponResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetShareUidList(v []*DescribeCouponResponseBodyDataShareUidList) *DescribeCouponResponseBodyData {
	s.ShareUidList = v
	return s
}

func (s *DescribeCouponResponseBodyData) SetShowSetDeductTagButton(v bool) *DescribeCouponResponseBodyData {
	s.ShowSetDeductTagButton = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetSite(v string) *DescribeCouponResponseBodyData {
	s.Site = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetSiteName(v string) *DescribeCouponResponseBodyData {
	s.SiteName = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetStartTime(v string) *DescribeCouponResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetStatus(v string) *DescribeCouponResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetSuitAccount(v string) *DescribeCouponResponseBodyData {
	s.SuitAccount = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetSuitItemType(v string) *DescribeCouponResponseBodyData {
	s.SuitItemType = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetUniversalType(v string) *DescribeCouponResponseBodyData {
	s.UniversalType = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetYhOrderTypes(v []*string) *DescribeCouponResponseBodyData {
	s.YhOrderTypes = v
	return s
}

type DescribeCouponResponseBodyDataShareUidList struct {
	// example:
	//
	// 1902671110151254
	Uid      *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s DescribeCouponResponseBodyDataShareUidList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponResponseBodyDataShareUidList) GoString() string {
	return s.String()
}

func (s *DescribeCouponResponseBodyDataShareUidList) SetUid(v string) *DescribeCouponResponseBodyDataShareUidList {
	s.Uid = &v
	return s
}

func (s *DescribeCouponResponseBodyDataShareUidList) SetUserNick(v string) *DescribeCouponResponseBodyDataShareUidList {
	s.UserNick = &v
	return s
}

type DescribeCouponResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCouponResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCouponResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponResponse) GoString() string {
	return s.String()
}

func (s *DescribeCouponResponse) SetHeaders(v map[string]*string) *DescribeCouponResponse {
	s.Headers = v
	return s
}

func (s *DescribeCouponResponse) SetStatusCode(v int32) *DescribeCouponResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCouponResponse) SetBody(v *DescribeCouponResponseBody) *DescribeCouponResponse {
	s.Body = v
	return s
}

type DescribeCouponItemListRequest struct {
	// example:
	//
	// 59104570
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// example:
	//
	// 1
	CurrentPage    *int32                                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds []*DescribeCouponItemListRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Name           *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCouponItemListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponItemListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListRequest) SetCouponId(v int64) *DescribeCouponItemListRequest {
	s.CouponId = &v
	return s
}

func (s *DescribeCouponItemListRequest) SetCurrentPage(v int32) *DescribeCouponItemListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponItemListRequest) SetEcIdAccountIds(v []*DescribeCouponItemListRequestEcIdAccountIds) *DescribeCouponItemListRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *DescribeCouponItemListRequest) SetName(v string) *DescribeCouponItemListRequest {
	s.Name = &v
	return s
}

func (s *DescribeCouponItemListRequest) SetNbid(v string) *DescribeCouponItemListRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeCouponItemListRequest) SetPageSize(v int32) *DescribeCouponItemListRequest {
	s.PageSize = &v
	return s
}

type DescribeCouponItemListRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1004064243473974
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DescribeCouponItemListRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponItemListRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListRequestEcIdAccountIds) SetAccountIds(v []*int64) *DescribeCouponItemListRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *DescribeCouponItemListRequestEcIdAccountIds) SetEcId(v string) *DescribeCouponItemListRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type DescribeCouponItemListShrinkRequest struct {
	// example:
	//
	// 59104570
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// example:
	//
	// 1
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCouponItemListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponItemListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListShrinkRequest) SetCouponId(v int64) *DescribeCouponItemListShrinkRequest {
	s.CouponId = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) SetCurrentPage(v int32) *DescribeCouponItemListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) SetEcIdAccountIdsShrink(v string) *DescribeCouponItemListShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) SetName(v string) *DescribeCouponItemListShrinkRequest {
	s.Name = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) SetNbid(v string) *DescribeCouponItemListShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) SetPageSize(v int32) *DescribeCouponItemListShrinkRequest {
	s.PageSize = &v
	return s
}

type DescribeCouponItemListResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*DescribeCouponItemListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// EAE08A27-386C-579E-966D-8853EC3C5D0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCouponItemListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponItemListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListResponseBody) SetCurrentPage(v int32) *DescribeCouponItemListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponItemListResponseBody) SetData(v []*DescribeCouponItemListResponseBodyData) *DescribeCouponItemListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCouponItemListResponseBody) SetPageSize(v int32) *DescribeCouponItemListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponItemListResponseBody) SetRequestId(v string) *DescribeCouponItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCouponItemListResponseBody) SetTotalCount(v int32) *DescribeCouponItemListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCouponItemListResponseBodyData struct {
	// example:
	//
	// vm
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCouponItemListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponItemListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListResponseBodyData) SetCode(v string) *DescribeCouponItemListResponseBodyData {
	s.Code = &v
	return s
}

func (s *DescribeCouponItemListResponseBodyData) SetName(v string) *DescribeCouponItemListResponseBodyData {
	s.Name = &v
	return s
}

type DescribeCouponItemListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCouponItemListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCouponItemListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCouponItemListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListResponse) SetHeaders(v map[string]*string) *DescribeCouponItemListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCouponItemListResponse) SetStatusCode(v int32) *DescribeCouponItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCouponItemListResponse) SetBody(v *DescribeCouponItemListResponseBody) *DescribeCouponItemListResponse {
	s.Body = v
	return s
}

type DescribeUserSpnSummaryInfoRequest struct {
	EcIdAccountIds []*DescribeUserSpnSummaryInfoRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid           *string                                            `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s DescribeUserSpnSummaryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoRequest) SetEcIdAccountIds(v []*DescribeUserSpnSummaryInfoRequestEcIdAccountIds) *DescribeUserSpnSummaryInfoRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *DescribeUserSpnSummaryInfoRequest) SetNbid(v string) *DescribeUserSpnSummaryInfoRequest {
	s.Nbid = &v
	return s
}

type DescribeUserSpnSummaryInfoRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DescribeUserSpnSummaryInfoRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoRequestEcIdAccountIds) SetAccountIds(v []*int64) *DescribeUserSpnSummaryInfoRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *DescribeUserSpnSummaryInfoRequestEcIdAccountIds) SetEcId(v string) *DescribeUserSpnSummaryInfoRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type DescribeUserSpnSummaryInfoShrinkRequest struct {
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s DescribeUserSpnSummaryInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoShrinkRequest) SetEcIdAccountIdsShrink(v string) *DescribeUserSpnSummaryInfoShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoShrinkRequest) SetNbid(v string) *DescribeUserSpnSummaryInfoShrinkRequest {
	s.Nbid = &v
	return s
}

type DescribeUserSpnSummaryInfoResponseBody struct {
	InstanceFamilyList []*string                                                   `json:"InstanceFamilyList,omitempty" xml:"InstanceFamilyList,omitempty" type:"Repeated"`
	RegionList         []*DescribeUserSpnSummaryInfoResponseBodyRegionList         `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpnCodeAndTypeList []*DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList `json:"SpnCodeAndTypeList,omitempty" xml:"SpnCodeAndTypeList,omitempty" type:"Repeated"`
}

func (s DescribeUserSpnSummaryInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoResponseBody) SetInstanceFamilyList(v []*string) *DescribeUserSpnSummaryInfoResponseBody {
	s.InstanceFamilyList = v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBody) SetRegionList(v []*DescribeUserSpnSummaryInfoResponseBodyRegionList) *DescribeUserSpnSummaryInfoResponseBody {
	s.RegionList = v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBody) SetRequestId(v string) *DescribeUserSpnSummaryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBody) SetSpnCodeAndTypeList(v []*DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) *DescribeUserSpnSummaryInfoResponseBody {
	s.SpnCodeAndTypeList = v
	return s
}

type DescribeUserSpnSummaryInfoResponseBodyRegionList struct {
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeUserSpnSummaryInfoResponseBodyRegionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoResponseBodyRegionList) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoResponseBodyRegionList) SetRegionCode(v string) *DescribeUserSpnSummaryInfoResponseBodyRegionList {
	s.RegionCode = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBodyRegionList) SetRegionName(v string) *DescribeUserSpnSummaryInfoResponseBodyRegionList {
	s.RegionName = &v
	return s
}

type DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList struct {
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SpnCommodityCode *string `json:"SpnCommodityCode,omitempty" xml:"SpnCommodityCode,omitempty"`
	SpnType          *string `json:"SpnType,omitempty" xml:"SpnType,omitempty"`
	SpnTypeName      *string `json:"SpnTypeName,omitempty" xml:"SpnTypeName,omitempty"`
}

func (s DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) SetProductCode(v string) *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList {
	s.ProductCode = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) SetSpnCommodityCode(v string) *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList {
	s.SpnCommodityCode = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) SetSpnType(v string) *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList {
	s.SpnType = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) SetSpnTypeName(v string) *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList {
	s.SpnTypeName = &v
	return s
}

type DescribeUserSpnSummaryInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserSpnSummaryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserSpnSummaryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoResponse) SetHeaders(v map[string]*string) *DescribeUserSpnSummaryInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponse) SetStatusCode(v int32) *DescribeUserSpnSummaryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponse) SetBody(v *DescribeUserSpnSummaryInfoResponseBody) *DescribeUserSpnSummaryInfoResponse {
	s.Body = v
	return s
}

type GetFundAccountAvailableAmountRequest struct {
	// example:
	//
	// 12332112
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s GetFundAccountAvailableAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountAvailableAmountRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountAvailableAmountRequest) SetFundAccountId(v string) *GetFundAccountAvailableAmountRequest {
	s.FundAccountId = &v
	return s
}

type GetFundAccountAvailableAmountResponseBody struct {
	// example:
	//
	// 100
	AvailableAmount *string `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// example:
	//
	// 50
	AvailableCreditAmount *string `json:"AvailableCreditAmount,omitempty" xml:"AvailableCreditAmount,omitempty"`
	// example:
	//
	// 0
	BankAcceptanceAmount *string `json:"BankAcceptanceAmount,omitempty" xml:"BankAcceptanceAmount,omitempty"`
	// example:
	//
	// 50
	CashAmount *string `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	// example:
	//
	// 100
	CreditAmount *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	// example:
	//
	// 0
	CreditRefundAmount *string `json:"CreditRefundAmount,omitempty" xml:"CreditRefundAmount,omitempty"`
	CreditUser         *bool   `json:"CreditUser,omitempty" xml:"CreditUser,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 20
	CurrentMonthUnclearedAmount *string                                                      `json:"CurrentMonthUnclearedAmount,omitempty" xml:"CurrentMonthUnclearedAmount,omitempty"`
	ExtendLedgerList            []*GetFundAccountAvailableAmountResponseBodyExtendLedgerList `json:"ExtendLedgerList,omitempty" xml:"ExtendLedgerList,omitempty" type:"Repeated"`
	// example:
	//
	// 12332112
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 1344312434
	FundAccountOwnerAccountId *string `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// valid
	FundAccountStatus *string `json:"FundAccountStatus,omitempty" xml:"FundAccountStatus,omitempty"`
	// example:
	//
	// REDIRECT_USER
	FundAccountType *string `json:"FundAccountType,omitempty" xml:"FundAccountType,omitempty"`
	// example:
	//
	// 30
	HistoryMonthUnclearedAmount *string `json:"HistoryMonthUnclearedAmount,omitempty" xml:"HistoryMonthUnclearedAmount,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 0
	NegativeBillAmount     *string                                                            `json:"NegativeBillAmount,omitempty" xml:"NegativeBillAmount,omitempty"`
	OriginalCashAmountList []*GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList `json:"OriginalCashAmountList,omitempty" xml:"OriginalCashAmountList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	QuotaAmount *string `json:"QuotaAmount,omitempty" xml:"QuotaAmount,omitempty"`
	// example:
	//
	// 10
	QuotaConsumedAmount *string `json:"QuotaConsumedAmount,omitempty" xml:"QuotaConsumedAmount,omitempty"`
	// example:
	//
	// F96A2D13-7509-5DF9-A60E-E7E3A3CB68E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 50
	UnclearedAmount *string `json:"UnclearedAmount,omitempty" xml:"UnclearedAmount,omitempty"`
}

func (s GetFundAccountAvailableAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountAvailableAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountAvailableAmountResponseBody) SetAvailableAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.AvailableAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetAvailableCreditAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.AvailableCreditAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetBankAcceptanceAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.BankAcceptanceAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCashAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.CashAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCreditAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.CreditAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCreditRefundAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.CreditRefundAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCreditUser(v bool) *GetFundAccountAvailableAmountResponseBody {
	s.CreditUser = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCurrency(v string) *GetFundAccountAvailableAmountResponseBody {
	s.Currency = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetCurrentMonthUnclearedAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.CurrentMonthUnclearedAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetExtendLedgerList(v []*GetFundAccountAvailableAmountResponseBodyExtendLedgerList) *GetFundAccountAvailableAmountResponseBody {
	s.ExtendLedgerList = v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetFundAccountId(v string) *GetFundAccountAvailableAmountResponseBody {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetFundAccountOwnerAccountId(v string) *GetFundAccountAvailableAmountResponseBody {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetFundAccountStatus(v string) *GetFundAccountAvailableAmountResponseBody {
	s.FundAccountStatus = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetFundAccountType(v string) *GetFundAccountAvailableAmountResponseBody {
	s.FundAccountType = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetHistoryMonthUnclearedAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.HistoryMonthUnclearedAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetMetadata(v interface{}) *GetFundAccountAvailableAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetNegativeBillAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.NegativeBillAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetOriginalCashAmountList(v []*GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) *GetFundAccountAvailableAmountResponseBody {
	s.OriginalCashAmountList = v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetQuotaAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.QuotaAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetQuotaConsumedAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.QuotaConsumedAmount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetRequestId(v string) *GetFundAccountAvailableAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBody) SetUnclearedAmount(v string) *GetFundAccountAvailableAmountResponseBody {
	s.UnclearedAmount = &v
	return s
}

type GetFundAccountAvailableAmountResponseBodyExtendLedgerList struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 应付对冲账本
	LedgerName *string `json:"LedgerName,omitempty" xml:"LedgerName,omitempty"`
	// example:
	//
	// 50
	OriginalAmount *string `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
}

func (s GetFundAccountAvailableAmountResponseBodyExtendLedgerList) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountAvailableAmountResponseBodyExtendLedgerList) GoString() string {
	return s.String()
}

func (s *GetFundAccountAvailableAmountResponseBodyExtendLedgerList) SetCurrency(v string) *GetFundAccountAvailableAmountResponseBodyExtendLedgerList {
	s.Currency = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBodyExtendLedgerList) SetLedgerName(v string) *GetFundAccountAvailableAmountResponseBodyExtendLedgerList {
	s.LedgerName = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBodyExtendLedgerList) SetOriginalAmount(v string) *GetFundAccountAvailableAmountResponseBodyExtendLedgerList {
	s.OriginalAmount = &v
	return s
}

type GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList struct {
	// example:
	//
	// 10
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
}

func (s GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) GoString() string {
	return s.String()
}

func (s *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) SetAmount(v string) *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList {
	s.Amount = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList) SetCurrency(v string) *GetFundAccountAvailableAmountResponseBodyOriginalCashAmountList {
	s.Currency = &v
	return s
}

type GetFundAccountAvailableAmountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountAvailableAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountAvailableAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountAvailableAmountResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountAvailableAmountResponse) SetHeaders(v map[string]*string) *GetFundAccountAvailableAmountResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountAvailableAmountResponse) SetStatusCode(v int32) *GetFundAccountAvailableAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponse) SetBody(v *GetFundAccountAvailableAmountResponseBody) *GetFundAccountAvailableAmountResponse {
	s.Body = v
	return s
}

type GetFundAccountCanAllocateCreditAmountRequest struct {
	// example:
	//
	// 1233231
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s GetFundAccountCanAllocateCreditAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanAllocateCreditAmountRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanAllocateCreditAmountRequest) SetFundAccountId(v int64) *GetFundAccountCanAllocateCreditAmountRequest {
	s.FundAccountId = &v
	return s
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
	return tea.Prettify(s)
}

func (s GetFundAccountCanAllocateCreditAmountResponseBody) GoString() string {
	return s.String()
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

type GetFundAccountCanAllocateCreditAmountResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountCanAllocateCreditAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountCanAllocateCreditAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanAllocateCreditAmountResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanAllocateCreditAmountResponse) SetHeaders(v map[string]*string) *GetFundAccountCanAllocateCreditAmountResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponse) SetStatusCode(v int32) *GetFundAccountCanAllocateCreditAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponse) SetBody(v *GetFundAccountCanAllocateCreditAmountResponseBody) *GetFundAccountCanAllocateCreditAmountResponse {
	s.Body = v
	return s
}

type GetFundAccountCanRecycleAmountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 122321223
	RecycleFromFundAccountId *string `json:"RecycleFromFundAccountId,omitempty" xml:"RecycleFromFundAccountId,omitempty"`
}

func (s GetFundAccountCanRecycleAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanRecycleAmountRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanRecycleAmountRequest) SetCurrency(v string) *GetFundAccountCanRecycleAmountRequest {
	s.Currency = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountRequest) SetRecycleFromFundAccountId(v string) *GetFundAccountCanRecycleAmountRequest {
	s.RecycleFromFundAccountId = &v
	return s
}

type GetFundAccountCanRecycleAmountResponseBody struct {
	// example:
	//
	// 300
	AvailableAmount *string `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 1232122132
	RecycleFromFundAccountId *string                                                               `json:"RecycleFromFundAccountId,omitempty" xml:"RecycleFromFundAccountId,omitempty"`
	RecycleToFundAccountList []*GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList `json:"RecycleToFundAccountList,omitempty" xml:"RecycleToFundAccountList,omitempty" type:"Repeated"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	TransferAmount *string `json:"TransferAmount,omitempty" xml:"TransferAmount,omitempty"`
}

func (s GetFundAccountCanRecycleAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanRecycleAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetAvailableAmount(v string) *GetFundAccountCanRecycleAmountResponseBody {
	s.AvailableAmount = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetCurrency(v string) *GetFundAccountCanRecycleAmountResponseBody {
	s.Currency = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetMetadata(v interface{}) *GetFundAccountCanRecycleAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetRecycleFromFundAccountId(v string) *GetFundAccountCanRecycleAmountResponseBody {
	s.RecycleFromFundAccountId = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetRecycleToFundAccountList(v []*GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) *GetFundAccountCanRecycleAmountResponseBody {
	s.RecycleToFundAccountList = v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetRequestId(v string) *GetFundAccountCanRecycleAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetTransferAmount(v string) *GetFundAccountCanRecycleAmountResponseBody {
	s.TransferAmount = &v
	return s
}

type GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList struct {
	// example:
	//
	// 122323121
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 云某的账户
	FundAccountName *string `json:"FundAccountName,omitempty" xml:"FundAccountName,omitempty"`
	// example:
	//
	// 183221321
	FundAccountOwnerAccountId *string `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// 120
	MaxRecyclableAmount *string `json:"MaxRecyclableAmount,omitempty" xml:"MaxRecyclableAmount,omitempty"`
	// example:
	//
	// 120
	OriginalTransferRemainAmount *string `json:"OriginalTransferRemainAmount,omitempty" xml:"OriginalTransferRemainAmount,omitempty"`
}

func (s GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) SetFundAccountId(v string) *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) SetFundAccountName(v string) *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	s.FundAccountName = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) SetFundAccountOwnerAccountId(v string) *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) SetMaxRecyclableAmount(v string) *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	s.MaxRecyclableAmount = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) SetOriginalTransferRemainAmount(v string) *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	s.OriginalTransferRemainAmount = &v
	return s
}

type GetFundAccountCanRecycleAmountResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountCanRecycleAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountCanRecycleAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanRecycleAmountResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanRecycleAmountResponse) SetHeaders(v map[string]*string) *GetFundAccountCanRecycleAmountResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponse) SetStatusCode(v int32) *GetFundAccountCanRecycleAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponse) SetBody(v *GetFundAccountCanRecycleAmountResponseBody) *GetFundAccountCanRecycleAmountResponse {
	s.Body = v
	return s
}

type GetFundAccountCanTransferAmountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 123212
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s GetFundAccountCanTransferAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanTransferAmountRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanTransferAmountRequest) SetCurrency(v string) *GetFundAccountCanTransferAmountRequest {
	s.Currency = &v
	return s
}

func (s *GetFundAccountCanTransferAmountRequest) SetFundAccountId(v string) *GetFundAccountCanTransferAmountRequest {
	s.FundAccountId = &v
	return s
}

type GetFundAccountCanTransferAmountResponseBody struct {
	// example:
	//
	// 100
	AvailableAmount *string `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// example:
	//
	// 500
	CashAmount *string `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 2032121324
	FundAccountEcid *string `json:"FundAccountEcid,omitempty" xml:"FundAccountEcid,omitempty"`
	// example:
	//
	// 12332112
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 云某的账户
	FundAccountName *string `json:"FundAccountName,omitempty" xml:"FundAccountName,omitempty"`
	// example:
	//
	// 154738212323
	FundAccountOwnerAccountId *int64 `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// 100
	MaxTransferableAmount *string `json:"MaxTransferableAmount,omitempty" xml:"MaxTransferableAmount,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 2684210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 26842
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// example:
	//
	// 100
	TransferAmount *string `json:"TransferAmount,omitempty" xml:"TransferAmount,omitempty"`
}

func (s GetFundAccountCanTransferAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanTransferAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetAvailableAmount(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.AvailableAmount = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetCashAmount(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.CashAmount = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetCurrency(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.Currency = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetFundAccountEcid(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.FundAccountEcid = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetFundAccountId(v int64) *GetFundAccountCanTransferAmountResponseBody {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetFundAccountName(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.FundAccountName = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetFundAccountOwnerAccountId(v int64) *GetFundAccountCanTransferAmountResponseBody {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetMaxTransferableAmount(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.MaxTransferableAmount = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetMetadata(v interface{}) *GetFundAccountCanTransferAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetNbid(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.Nbid = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetRequestId(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetSite(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.Site = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponseBody) SetTransferAmount(v string) *GetFundAccountCanTransferAmountResponseBody {
	s.TransferAmount = &v
	return s
}

type GetFundAccountCanTransferAmountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountCanTransferAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountCanTransferAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanTransferAmountResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanTransferAmountResponse) SetHeaders(v map[string]*string) *GetFundAccountCanTransferAmountResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountCanTransferAmountResponse) SetStatusCode(v int32) *GetFundAccountCanTransferAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponse) SetBody(v *GetFundAccountCanTransferAmountResponseBody) *GetFundAccountCanTransferAmountResponse {
	s.Body = v
	return s
}

type GetFundAccountCanWithdrawAmountRequest struct {
	// example:
	//
	// 123212232
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s GetFundAccountCanWithdrawAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanWithdrawAmountRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanWithdrawAmountRequest) SetFundAccountId(v int64) *GetFundAccountCanWithdrawAmountRequest {
	s.FundAccountId = &v
	return s
}

type GetFundAccountCanWithdrawAmountResponseBody struct {
	// example:
	//
	// 400
	CanOriginalWithdrawAmount *string `json:"CanOriginalWithdrawAmount,omitempty" xml:"CanOriginalWithdrawAmount,omitempty"`
	// example:
	//
	// 500
	CanWithdrawAmount *string `json:"CanWithdrawAmount,omitempty" xml:"CanWithdrawAmount,omitempty"`
	// example:
	//
	// 100
	CannotOriginalWithdrawAmount *string `json:"CannotOriginalWithdrawAmount,omitempty" xml:"CannotOriginalWithdrawAmount,omitempty"`
	// example:
	//
	// 1000
	CashAmount *string `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	// example:
	//
	// 100
	CreditMemoAmount *string `json:"CreditMemoAmount,omitempty" xml:"CreditMemoAmount,omitempty"`
	// example:
	//
	// 200
	CurrentMonthUnclearedAmount *string `json:"CurrentMonthUnclearedAmount,omitempty" xml:"CurrentMonthUnclearedAmount,omitempty"`
	// example:
	//
	// 100
	HistoryMonthUnclearedAmount *string `json:"HistoryMonthUnclearedAmount,omitempty" xml:"HistoryMonthUnclearedAmount,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 100
	PayAsYouGoReversedAmount *string `json:"PayAsYouGoReversedAmount,omitempty" xml:"PayAsYouGoReversedAmount,omitempty"`
	// example:
	//
	// DF58589C-A06C-4224-8615-7797E6474FA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	TransferAmount *string `json:"TransferAmount,omitempty" xml:"TransferAmount,omitempty"`
}

func (s GetFundAccountCanWithdrawAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanWithdrawAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCanOriginalWithdrawAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CanOriginalWithdrawAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCanWithdrawAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CanWithdrawAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCannotOriginalWithdrawAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CannotOriginalWithdrawAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCashAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CashAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCreditMemoAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CreditMemoAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetCurrentMonthUnclearedAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.CurrentMonthUnclearedAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetHistoryMonthUnclearedAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.HistoryMonthUnclearedAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetMetadata(v interface{}) *GetFundAccountCanWithdrawAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetPayAsYouGoReversedAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.PayAsYouGoReversedAmount = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetRequestId(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponseBody) SetTransferAmount(v string) *GetFundAccountCanWithdrawAmountResponseBody {
	s.TransferAmount = &v
	return s
}

type GetFundAccountCanWithdrawAmountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountCanWithdrawAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountCanWithdrawAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountCanWithdrawAmountResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanWithdrawAmountResponse) SetHeaders(v map[string]*string) *GetFundAccountCanWithdrawAmountResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponse) SetStatusCode(v int32) *GetFundAccountCanWithdrawAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponse) SetBody(v *GetFundAccountCanWithdrawAmountResponseBody) *GetFundAccountCanWithdrawAmountResponse {
	s.Body = v
	return s
}

type GetFundAccountLowAvailableAmountAlarmRequest struct {
	// example:
	//
	// 12332112
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s GetFundAccountLowAvailableAmountAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountLowAvailableAmountAlarmRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountLowAvailableAmountAlarmRequest) SetFundAccountId(v int64) *GetFundAccountLowAvailableAmountAlarmRequest {
	s.FundAccountId = &v
	return s
}

type GetFundAccountLowAvailableAmountAlarmResponseBody struct {
	AlarmEnabled *bool `json:"AlarmEnabled,omitempty" xml:"AlarmEnabled,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ThresholdAmount *string `json:"ThresholdAmount,omitempty" xml:"ThresholdAmount,omitempty"`
}

func (s GetFundAccountLowAvailableAmountAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountLowAvailableAmountAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) SetAlarmEnabled(v bool) *GetFundAccountLowAvailableAmountAlarmResponseBody {
	s.AlarmEnabled = &v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) SetMetadata(v interface{}) *GetFundAccountLowAvailableAmountAlarmResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) SetRequestId(v string) *GetFundAccountLowAvailableAmountAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponseBody) SetThresholdAmount(v string) *GetFundAccountLowAvailableAmountAlarmResponseBody {
	s.ThresholdAmount = &v
	return s
}

type GetFundAccountLowAvailableAmountAlarmResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountLowAvailableAmountAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountLowAvailableAmountAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountLowAvailableAmountAlarmResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountLowAvailableAmountAlarmResponse) SetHeaders(v map[string]*string) *GetFundAccountLowAvailableAmountAlarmResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponse) SetStatusCode(v int32) *GetFundAccountLowAvailableAmountAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountLowAvailableAmountAlarmResponse) SetBody(v *GetFundAccountLowAvailableAmountAlarmResponseBody) *GetFundAccountLowAvailableAmountAlarmResponse {
	s.Body = v
	return s
}

type GetFundAccountTransactionDetailsRequest struct {
	// example:
	//
	// 2023212312321
	BillNumber *string `json:"BillNumber,omitempty" xml:"BillNumber,omitempty"`
	// example:
	//
	// 20250312334312322
	ChannelTransactionNumber *string `json:"ChannelTransactionNumber,omitempty" xml:"ChannelTransactionNumber,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1735664561000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 123221232
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1735664461000
	StartTime              *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TransactionChannelList []*string `json:"TransactionChannelList,omitempty" xml:"TransactionChannelList,omitempty" type:"Repeated"`
	// example:
	//
	// IN
	TransactionDirection *string `json:"TransactionDirection,omitempty" xml:"TransactionDirection,omitempty"`
	// example:
	//
	// 543231231
	TransactionNumber *int64 `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	// example:
	//
	// CHARGE
	TransactionType     *string   `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
	TransactionTypeList []*string `json:"TransactionTypeList,omitempty" xml:"TransactionTypeList,omitempty" type:"Repeated"`
}

func (s GetFundAccountTransactionDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountTransactionDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsRequest) SetBillNumber(v string) *GetFundAccountTransactionDetailsRequest {
	s.BillNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetChannelTransactionNumber(v string) *GetFundAccountTransactionDetailsRequest {
	s.ChannelTransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetCurrentPage(v int32) *GetFundAccountTransactionDetailsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetEndTime(v int64) *GetFundAccountTransactionDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetFundAccountId(v int64) *GetFundAccountTransactionDetailsRequest {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetPageSize(v int32) *GetFundAccountTransactionDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetStartTime(v int64) *GetFundAccountTransactionDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetTransactionChannelList(v []*string) *GetFundAccountTransactionDetailsRequest {
	s.TransactionChannelList = v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetTransactionDirection(v string) *GetFundAccountTransactionDetailsRequest {
	s.TransactionDirection = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetTransactionNumber(v int64) *GetFundAccountTransactionDetailsRequest {
	s.TransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetTransactionType(v string) *GetFundAccountTransactionDetailsRequest {
	s.TransactionType = &v
	return s
}

func (s *GetFundAccountTransactionDetailsRequest) SetTransactionTypeList(v []*string) *GetFundAccountTransactionDetailsRequest {
	s.TransactionTypeList = v
	return s
}

type GetFundAccountTransactionDetailsShrinkRequest struct {
	// example:
	//
	// 2023212312321
	BillNumber *string `json:"BillNumber,omitempty" xml:"BillNumber,omitempty"`
	// example:
	//
	// 20250312334312322
	ChannelTransactionNumber *string `json:"ChannelTransactionNumber,omitempty" xml:"ChannelTransactionNumber,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1735664561000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 123221232
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1735664461000
	StartTime                    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TransactionChannelListShrink *string `json:"TransactionChannelList,omitempty" xml:"TransactionChannelList,omitempty"`
	// example:
	//
	// IN
	TransactionDirection *string `json:"TransactionDirection,omitempty" xml:"TransactionDirection,omitempty"`
	// example:
	//
	// 543231231
	TransactionNumber *int64 `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	// example:
	//
	// CHARGE
	TransactionType           *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
	TransactionTypeListShrink *string `json:"TransactionTypeList,omitempty" xml:"TransactionTypeList,omitempty"`
}

func (s GetFundAccountTransactionDetailsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountTransactionDetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetBillNumber(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.BillNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetChannelTransactionNumber(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.ChannelTransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetCurrentPage(v int32) *GetFundAccountTransactionDetailsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetEndTime(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetFundAccountId(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetPageSize(v int32) *GetFundAccountTransactionDetailsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetStartTime(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionChannelListShrink(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionChannelListShrink = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionDirection(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionDirection = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionNumber(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionType(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionType = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionTypeListShrink(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionTypeListShrink = &v
	return s
}

type GetFundAccountTransactionDetailsResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*GetFundAccountTransactionDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// D6E068C3-25BC-455A-85FE-45F0B22ECB1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetFundAccountTransactionDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountTransactionDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetCurrentPage(v int32) *GetFundAccountTransactionDetailsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetData(v []*GetFundAccountTransactionDetailsResponseBodyData) *GetFundAccountTransactionDetailsResponseBody {
	s.Data = v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetMetadata(v interface{}) *GetFundAccountTransactionDetailsResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetPageSize(v int32) *GetFundAccountTransactionDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetRequestId(v string) *GetFundAccountTransactionDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBody) SetTotalCount(v int32) *GetFundAccountTransactionDetailsResponseBody {
	s.TotalCount = &v
	return s
}

type GetFundAccountTransactionDetailsResponseBodyData struct {
	// example:
	//
	// 5
	Balance *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	// example:
	//
	// 2323203243
	BillNumber *string `json:"BillNumber,omitempty" xml:"BillNumber,omitempty"`
	// example:
	//
	// 20244389232
	ChannelTransactionNumber *string `json:"ChannelTransactionNumber,omitempty" xml:"ChannelTransactionNumber,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 23473943
	FundAccountEcid *string `json:"FundAccountEcid,omitempty" xml:"FundAccountEcid,omitempty"`
	// example:
	//
	// 1232121
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 云某的名称
	FundAccountName *string `json:"FundAccountName,omitempty" xml:"FundAccountName,omitempty"`
	// example:
	//
	// 32343231
	FundAccountOwnerAccountId *int64 `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// ACCT_BOOK
	FundType *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	// example:
	//
	// 2684210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 订单备注
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 26842
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// example:
	//
	// 186****3975
	TransactionAccount *string `json:"TransactionAccount,omitempty" xml:"TransactionAccount,omitempty"`
	// example:
	//
	// 10
	TransactionAmount *string `json:"TransactionAmount,omitempty" xml:"TransactionAmount,omitempty"`
	// example:
	//
	// ALIPAY
	TransactionChannel *string `json:"TransactionChannel,omitempty" xml:"TransactionChannel,omitempty"`
	// example:
	//
	// IN
	TransactionDirection *string `json:"TransactionDirection,omitempty" xml:"TransactionDirection,omitempty"`
	// example:
	//
	// 5423121
	TransactionNumber *int64 `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	// example:
	//
	// 2024-12-01 12:00:00
	TransactionTime *string `json:"TransactionTime,omitempty" xml:"TransactionTime,omitempty"`
	// example:
	//
	// CHARGE
	TransactionType *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
}

func (s GetFundAccountTransactionDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountTransactionDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetBalance(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.Balance = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetBillNumber(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.BillNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetChannelTransactionNumber(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.ChannelTransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetCurrency(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetFundAccountEcid(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.FundAccountEcid = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetFundAccountId(v int64) *GetFundAccountTransactionDetailsResponseBodyData {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetFundAccountName(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.FundAccountName = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetFundAccountOwnerAccountId(v int64) *GetFundAccountTransactionDetailsResponseBodyData {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetFundType(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.FundType = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetNbid(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.Nbid = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetRemark(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetSite(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.Site = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionAccount(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionAccount = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionAmount(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionAmount = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionChannel(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionChannel = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionDirection(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionDirection = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionNumber(v int64) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionTime(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponseBodyData) SetTransactionType(v string) *GetFundAccountTransactionDetailsResponseBodyData {
	s.TransactionType = &v
	return s
}

type GetFundAccountTransactionDetailsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountTransactionDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountTransactionDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFundAccountTransactionDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsResponse) SetHeaders(v map[string]*string) *GetFundAccountTransactionDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountTransactionDetailsResponse) SetStatusCode(v int32) *GetFundAccountTransactionDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponse) SetBody(v *GetFundAccountTransactionDetailsResponseBody) *GetFundAccountTransactionDetailsResponse {
	s.Body = v
	return s
}

type GetSavingPlanDeductableCommodityRequest struct {
	EcIdAccountIds []*GetSavingPlanDeductableCommodityRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid           *string                                                  `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s GetSavingPlanDeductableCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityRequest) SetEcIdAccountIds(v []*GetSavingPlanDeductableCommodityRequestEcIdAccountIds) *GetSavingPlanDeductableCommodityRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *GetSavingPlanDeductableCommodityRequest) SetNbid(v string) *GetSavingPlanDeductableCommodityRequest {
	s.Nbid = &v
	return s
}

type GetSavingPlanDeductableCommodityRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s GetSavingPlanDeductableCommodityRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityRequestEcIdAccountIds) SetAccountIds(v []*int64) *GetSavingPlanDeductableCommodityRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *GetSavingPlanDeductableCommodityRequestEcIdAccountIds) SetEcId(v string) *GetSavingPlanDeductableCommodityRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type GetSavingPlanDeductableCommodityShrinkRequest struct {
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s GetSavingPlanDeductableCommodityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityShrinkRequest) SetEcIdAccountIdsShrink(v string) *GetSavingPlanDeductableCommodityShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityShrinkRequest) SetNbid(v string) *GetSavingPlanDeductableCommodityShrinkRequest {
	s.Nbid = &v
	return s
}

type GetSavingPlanDeductableCommodityResponseBody struct {
	Data      []*GetSavingPlanDeductableCommodityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBody) SetData(v []*GetSavingPlanDeductableCommodityResponseBodyData) *GetSavingPlanDeductableCommodityResponseBody {
	s.Data = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBody) SetRequestId(v string) *GetSavingPlanDeductableCommodityResponseBody {
	s.RequestId = &v
	return s
}

type GetSavingPlanDeductableCommodityResponseBodyData struct {
	ActivityId            *int64                                                            `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	CommodityCode         *string                                                           `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityId           *int64                                                            `json:"CommodityId,omitempty" xml:"CommodityId,omitempty"`
	CommodityName         *string                                                           `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	CycleList             []*GetSavingPlanDeductableCommodityResponseBodyDataCycleList      `json:"CycleList,omitempty" xml:"CycleList,omitempty" type:"Repeated"`
	FilterModules         []*GetSavingPlanDeductableCommodityResponseBodyDataFilterModules  `json:"FilterModules,omitempty" xml:"FilterModules,omitempty" type:"Repeated"`
	ItemCode              *string                                                           `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	ItemId                *int64                                                            `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName              *string                                                           `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	ModuleMapList         []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList  `json:"ModuleMapList,omitempty" xml:"ModuleMapList,omitempty" type:"Repeated"`
	PayModeList           []*GetSavingPlanDeductableCommodityResponseBodyDataPayModeList    `json:"PayModeList,omitempty" xml:"PayModeList,omitempty" type:"Repeated"`
	PricingModules        []*GetSavingPlanDeductableCommodityResponseBodyDataPricingModules `json:"PricingModules,omitempty" xml:"PricingModules,omitempty" type:"Repeated"`
	SpnCommodityCode      *string                                                           `json:"SpnCommodityCode,omitempty" xml:"SpnCommodityCode,omitempty"`
	SpnCommodityName      *string                                                           `json:"SpnCommodityName,omitempty" xml:"SpnCommodityName,omitempty"`
	SpnDiscountConfigType *string                                                           `json:"SpnDiscountConfigType,omitempty" xml:"SpnDiscountConfigType,omitempty"`
	StepPriceMap          map[string][]*DataStepPriceMapValue                               `json:"StepPriceMap,omitempty" xml:"StepPriceMap,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetActivityId(v int64) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetCommodityCode(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetCommodityId(v int64) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.CommodityId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetCommodityName(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetCycleList(v []*GetSavingPlanDeductableCommodityResponseBodyDataCycleList) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.CycleList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetFilterModules(v []*GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.FilterModules = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetItemCode(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.ItemCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetItemId(v int64) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.ItemId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetItemName(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.ItemName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetModuleMapList(v []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.ModuleMapList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetPayModeList(v []*GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.PayModeList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetPricingModules(v []*GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.PricingModules = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetSpnCommodityCode(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.SpnCommodityCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetSpnCommodityName(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.SpnCommodityName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetSpnDiscountConfigType(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.SpnDiscountConfigType = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetStepPriceMap(v map[string][]*DataStepPriceMapValue) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.StepPriceMap = v
	return s
}

type GetSavingPlanDeductableCommodityResponseBodyDataCycleList struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataCycleList) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataCycleList) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataCycleList) SetCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataCycleList {
	s.Code = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataCycleList) SetName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataCycleList {
	s.Name = &v
	return s
}

type GetSavingPlanDeductableCommodityResponseBodyDataFilterModules struct {
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) SetModuleCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) SetModuleId(v int64) *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules {
	s.ModuleId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) SetModuleName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules {
	s.ModuleName = &v
	return s
}

type GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList struct {
	FilterModules   []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules   `json:"FilterModules,omitempty" xml:"FilterModules,omitempty" type:"Repeated"`
	ModuleCode      *string                                                                         `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleId        *int64                                                                          `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName      *string                                                                         `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ShowModules     []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules     `json:"ShowModules,omitempty" xml:"ShowModules,omitempty" type:"Repeated"`
	SpnTypeList     []*string                                                                       `json:"SpnTypeList,omitempty" xml:"SpnTypeList,omitempty" type:"Repeated"`
	SpnTypeMapList  []map[string]*DataModuleMapListSpnTypeMapListValue                              `json:"SpnTypeMapList,omitempty" xml:"SpnTypeMapList,omitempty" type:"Repeated"`
	SpnTypeNameList []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList `json:"SpnTypeNameList,omitempty" xml:"SpnTypeNameList,omitempty" type:"Repeated"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetFilterModules(v []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.FilterModules = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetModuleCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetModuleId(v int64) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.ModuleId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetModuleName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.ModuleName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetShowModules(v []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.ShowModules = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetSpnTypeList(v []*string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.SpnTypeList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetSpnTypeMapList(v []map[string]*DataModuleMapListSpnTypeMapListValue) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.SpnTypeMapList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetSpnTypeNameList(v []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.SpnTypeNameList = v
	return s
}

type GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules struct {
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) SetModuleCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) SetModuleId(v int64) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules {
	s.ModuleId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) SetModuleName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules {
	s.ModuleName = &v
	return s
}

type GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules struct {
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) SetModuleCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) SetModuleId(v int64) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules {
	s.ModuleId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) SetModuleName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules {
	s.ModuleName = &v
	return s
}

type GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) SetCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList {
	s.Code = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) SetName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList {
	s.Name = &v
	return s
}

type GetSavingPlanDeductableCommodityResponseBodyDataPayModeList struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) SetCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList {
	s.Code = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) SetName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList {
	s.Name = &v
	return s
}

type GetSavingPlanDeductableCommodityResponseBodyDataPricingModules struct {
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) SetModuleCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) SetModuleId(v int64) *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules {
	s.ModuleId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) SetModuleName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules {
	s.ModuleName = &v
	return s
}

type GetSavingPlanDeductableCommodityResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavingPlanDeductableCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponse) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponse) SetHeaders(v map[string]*string) *GetSavingPlanDeductableCommodityResponse {
	s.Headers = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponse) SetStatusCode(v int32) *GetSavingPlanDeductableCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponse) SetBody(v *GetSavingPlanDeductableCommodityResponseBody) *GetSavingPlanDeductableCommodityResponse {
	s.Body = v
	return s
}

type GetSavingPlanShareAccountsRequest struct {
	CurrentPage     *int32                                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds  []*GetSavingPlanShareAccountsRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid            *string                                            `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	PageSize        *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpnInstanceCode *string                                            `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
}

func (s GetSavingPlanShareAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanShareAccountsRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsRequest) SetCurrentPage(v int32) *GetSavingPlanShareAccountsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSavingPlanShareAccountsRequest) SetEcIdAccountIds(v []*GetSavingPlanShareAccountsRequestEcIdAccountIds) *GetSavingPlanShareAccountsRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *GetSavingPlanShareAccountsRequest) SetNbid(v string) *GetSavingPlanShareAccountsRequest {
	s.Nbid = &v
	return s
}

func (s *GetSavingPlanShareAccountsRequest) SetPageSize(v int32) *GetSavingPlanShareAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *GetSavingPlanShareAccountsRequest) SetSpnInstanceCode(v string) *GetSavingPlanShareAccountsRequest {
	s.SpnInstanceCode = &v
	return s
}

type GetSavingPlanShareAccountsRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s GetSavingPlanShareAccountsRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanShareAccountsRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsRequestEcIdAccountIds) SetAccountIds(v []*int64) *GetSavingPlanShareAccountsRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *GetSavingPlanShareAccountsRequestEcIdAccountIds) SetEcId(v string) *GetSavingPlanShareAccountsRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type GetSavingPlanShareAccountsShrinkRequest struct {
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpnInstanceCode      *string `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
}

func (s GetSavingPlanShareAccountsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanShareAccountsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsShrinkRequest) SetCurrentPage(v int32) *GetSavingPlanShareAccountsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSavingPlanShareAccountsShrinkRequest) SetEcIdAccountIdsShrink(v string) *GetSavingPlanShareAccountsShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *GetSavingPlanShareAccountsShrinkRequest) SetNbid(v string) *GetSavingPlanShareAccountsShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *GetSavingPlanShareAccountsShrinkRequest) SetPageSize(v int32) *GetSavingPlanShareAccountsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSavingPlanShareAccountsShrinkRequest) SetSpnInstanceCode(v string) *GetSavingPlanShareAccountsShrinkRequest {
	s.SpnInstanceCode = &v
	return s
}

type GetSavingPlanShareAccountsResponseBody struct {
	Data      []*GetSavingPlanShareAccountsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSavingPlanShareAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanShareAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsResponseBody) SetData(v []*GetSavingPlanShareAccountsResponseBodyData) *GetSavingPlanShareAccountsResponseBody {
	s.Data = v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBody) SetRequestId(v string) *GetSavingPlanShareAccountsResponseBody {
	s.RequestId = &v
	return s
}

type GetSavingPlanShareAccountsResponseBodyData struct {
	AccountId     *string                                                    `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AliUid        *int64                                                     `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ShareTimeList []*GetSavingPlanShareAccountsResponseBodyDataShareTimeList `json:"ShareTimeList,omitempty" xml:"ShareTimeList,omitempty" type:"Repeated"`
}

func (s GetSavingPlanShareAccountsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanShareAccountsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsResponseBodyData) SetAccountId(v string) *GetSavingPlanShareAccountsResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBodyData) SetAliUid(v int64) *GetSavingPlanShareAccountsResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBodyData) SetShareTimeList(v []*GetSavingPlanShareAccountsResponseBodyDataShareTimeList) *GetSavingPlanShareAccountsResponseBodyData {
	s.ShareTimeList = v
	return s
}

type GetSavingPlanShareAccountsResponseBodyDataShareTimeList struct {
	ShareEndTime   *string `json:"ShareEndTime,omitempty" xml:"ShareEndTime,omitempty"`
	ShareStartTime *string `json:"ShareStartTime,omitempty" xml:"ShareStartTime,omitempty"`
}

func (s GetSavingPlanShareAccountsResponseBodyDataShareTimeList) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanShareAccountsResponseBodyDataShareTimeList) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsResponseBodyDataShareTimeList) SetShareEndTime(v string) *GetSavingPlanShareAccountsResponseBodyDataShareTimeList {
	s.ShareEndTime = &v
	return s
}

func (s *GetSavingPlanShareAccountsResponseBodyDataShareTimeList) SetShareStartTime(v string) *GetSavingPlanShareAccountsResponseBodyDataShareTimeList {
	s.ShareStartTime = &v
	return s
}

type GetSavingPlanShareAccountsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavingPlanShareAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSavingPlanShareAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanShareAccountsResponse) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsResponse) SetHeaders(v map[string]*string) *GetSavingPlanShareAccountsResponse {
	s.Headers = v
	return s
}

func (s *GetSavingPlanShareAccountsResponse) SetStatusCode(v int32) *GetSavingPlanShareAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavingPlanShareAccountsResponse) SetBody(v *GetSavingPlanShareAccountsResponseBody) *GetSavingPlanShareAccountsResponse {
	s.Body = v
	return s
}

type GetSavingPlanUserDeductRuleRequest struct {
	CurrentPage     *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds  []*GetSavingPlanUserDeductRuleRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid            *string                                             `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	PageSize        *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpnInstanceCode *string                                             `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
}

func (s GetSavingPlanUserDeductRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleRequest) SetCurrentPage(v int32) *GetSavingPlanUserDeductRuleRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequest) SetEcIdAccountIds(v []*GetSavingPlanUserDeductRuleRequestEcIdAccountIds) *GetSavingPlanUserDeductRuleRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequest) SetNbid(v string) *GetSavingPlanUserDeductRuleRequest {
	s.Nbid = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequest) SetPageSize(v int32) *GetSavingPlanUserDeductRuleRequest {
	s.PageSize = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequest) SetSpnInstanceCode(v string) *GetSavingPlanUserDeductRuleRequest {
	s.SpnInstanceCode = &v
	return s
}

type GetSavingPlanUserDeductRuleRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s GetSavingPlanUserDeductRuleRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleRequestEcIdAccountIds) SetAccountIds(v []*int64) *GetSavingPlanUserDeductRuleRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *GetSavingPlanUserDeductRuleRequestEcIdAccountIds) SetEcId(v string) *GetSavingPlanUserDeductRuleRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type GetSavingPlanUserDeductRuleShrinkRequest struct {
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpnInstanceCode      *string `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
}

func (s GetSavingPlanUserDeductRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) SetCurrentPage(v int32) *GetSavingPlanUserDeductRuleShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) SetEcIdAccountIdsShrink(v string) *GetSavingPlanUserDeductRuleShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) SetNbid(v string) *GetSavingPlanUserDeductRuleShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) SetPageSize(v int32) *GetSavingPlanUserDeductRuleShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleShrinkRequest) SetSpnInstanceCode(v string) *GetSavingPlanUserDeductRuleShrinkRequest {
	s.SpnInstanceCode = &v
	return s
}

type GetSavingPlanUserDeductRuleResponseBody struct {
	Data      []*GetSavingPlanUserDeductRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSavingPlanUserDeductRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleResponseBody) SetData(v []*GetSavingPlanUserDeductRuleResponseBodyData) *GetSavingPlanUserDeductRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBody) SetRequestId(v string) *GetSavingPlanUserDeductRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetSavingPlanUserDeductRuleResponseBodyData struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	ModuleCode    *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleName    *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	SkipDeduct    *bool   `json:"SkipDeduct,omitempty" xml:"SkipDeduct,omitempty"`
}

func (s GetSavingPlanUserDeductRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) SetCommodityCode(v string) *GetSavingPlanUserDeductRuleResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) SetCommodityName(v string) *GetSavingPlanUserDeductRuleResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) SetModuleCode(v string) *GetSavingPlanUserDeductRuleResponseBodyData {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) SetModuleName(v string) *GetSavingPlanUserDeductRuleResponseBodyData {
	s.ModuleName = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) SetSkipDeduct(v bool) *GetSavingPlanUserDeductRuleResponseBodyData {
	s.SkipDeduct = &v
	return s
}

type GetSavingPlanUserDeductRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavingPlanUserDeductRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSavingPlanUserDeductRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleResponse) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleResponse) SetHeaders(v map[string]*string) *GetSavingPlanUserDeductRuleResponse {
	s.Headers = v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponse) SetStatusCode(v int32) *GetSavingPlanUserDeductRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponse) SetBody(v *GetSavingPlanUserDeductRuleResponseBody) *GetSavingPlanUserDeductRuleResponse {
	s.Body = v
	return s
}

type ListCouponDeductTagRequest struct {
	CouponId       *string                                     `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIds []*ListCouponDeductTagRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid           *string                                     `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s ListCouponDeductTagRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCouponDeductTagRequest) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagRequest) SetCouponId(v string) *ListCouponDeductTagRequest {
	s.CouponId = &v
	return s
}

func (s *ListCouponDeductTagRequest) SetEcIdAccountIds(v []*ListCouponDeductTagRequestEcIdAccountIds) *ListCouponDeductTagRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *ListCouponDeductTagRequest) SetNbid(v string) *ListCouponDeductTagRequest {
	s.Nbid = &v
	return s
}

type ListCouponDeductTagRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s ListCouponDeductTagRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s ListCouponDeductTagRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagRequestEcIdAccountIds) SetAccountIds(v []*int64) *ListCouponDeductTagRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *ListCouponDeductTagRequestEcIdAccountIds) SetEcId(v string) *ListCouponDeductTagRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type ListCouponDeductTagShrinkRequest struct {
	CouponId             *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s ListCouponDeductTagShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCouponDeductTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagShrinkRequest) SetCouponId(v string) *ListCouponDeductTagShrinkRequest {
	s.CouponId = &v
	return s
}

func (s *ListCouponDeductTagShrinkRequest) SetEcIdAccountIdsShrink(v string) *ListCouponDeductTagShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *ListCouponDeductTagShrinkRequest) SetNbid(v string) *ListCouponDeductTagShrinkRequest {
	s.Nbid = &v
	return s
}

type ListCouponDeductTagResponseBody struct {
	Data      []*ListCouponDeductTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCouponDeductTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCouponDeductTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagResponseBody) SetData(v []*ListCouponDeductTagResponseBodyData) *ListCouponDeductTagResponseBody {
	s.Data = v
	return s
}

func (s *ListCouponDeductTagResponseBody) SetRequestId(v string) *ListCouponDeductTagResponseBody {
	s.RequestId = &v
	return s
}

type ListCouponDeductTagResponseBodyData struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCouponDeductTagResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCouponDeductTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagResponseBodyData) SetKey(v string) *ListCouponDeductTagResponseBodyData {
	s.Key = &v
	return s
}

func (s *ListCouponDeductTagResponseBodyData) SetValue(v string) *ListCouponDeductTagResponseBodyData {
	s.Value = &v
	return s
}

type ListCouponDeductTagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCouponDeductTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCouponDeductTagResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCouponDeductTagResponse) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagResponse) SetHeaders(v map[string]*string) *ListCouponDeductTagResponse {
	s.Headers = v
	return s
}

func (s *ListCouponDeductTagResponse) SetStatusCode(v int32) *ListCouponDeductTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCouponDeductTagResponse) SetBody(v *ListCouponDeductTagResponseBody) *ListCouponDeductTagResponse {
	s.Body = v
	return s
}

type ListFundAccountRequest struct {
	// example:
	//
	// 2084210001
	Nbid            *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	QueryOnlyInUse  *bool   `json:"QueryOnlyInUse,omitempty" xml:"QueryOnlyInUse,omitempty"`
	QueryOnlyManage *bool   `json:"QueryOnlyManage,omitempty" xml:"QueryOnlyManage,omitempty"`
}

func (s ListFundAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFundAccountRequest) GoString() string {
	return s.String()
}

func (s *ListFundAccountRequest) SetNbid(v string) *ListFundAccountRequest {
	s.Nbid = &v
	return s
}

func (s *ListFundAccountRequest) SetQueryOnlyInUse(v bool) *ListFundAccountRequest {
	s.QueryOnlyInUse = &v
	return s
}

func (s *ListFundAccountRequest) SetQueryOnlyManage(v bool) *ListFundAccountRequest {
	s.QueryOnlyManage = &v
	return s
}

type ListFundAccountResponseBody struct {
	Data []*ListFundAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFundAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFundAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListFundAccountResponseBody) SetData(v []*ListFundAccountResponseBodyData) *ListFundAccountResponseBody {
	s.Data = v
	return s
}

func (s *ListFundAccountResponseBody) SetMetadata(v interface{}) *ListFundAccountResponseBody {
	s.Metadata = v
	return s
}

func (s *ListFundAccountResponseBody) SetRequestId(v string) *ListFundAccountResponseBody {
	s.RequestId = &v
	return s
}

type ListFundAccountResponseBodyData struct {
	// example:
	//
	// 2024-12-30 12:00:00
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// example:
	//
	// 1232121
	FundAccountAdminAccountId *string `json:"FundAccountAdminAccountId,omitempty" xml:"FundAccountAdminAccountId,omitempty"`
	// example:
	//
	// 云某的账户
	FundAccountAdminAccountName *string `json:"FundAccountAdminAccountName,omitempty" xml:"FundAccountAdminAccountName,omitempty"`
	// example:
	//
	// 1022231
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 云某的账户
	FundAccountName *string `json:"FundAccountName,omitempty" xml:"FundAccountName,omitempty"`
	// example:
	//
	// 132123211
	FundAccountOwnerAccountId *string `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// VALID
	FundAccountStatus *string `json:"FundAccountStatus,omitempty" xml:"FundAccountStatus,omitempty"`
	// example:
	//
	// DIRECT_USER
	FundAccountType *string `json:"FundAccountType,omitempty" xml:"FundAccountType,omitempty"`
	// example:
	//
	// 2684210001
	Nbid        *string   `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	Permissions []*string `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// example:
	//
	// 26842
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s ListFundAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFundAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFundAccountResponseBodyData) SetCreateDate(v string) *ListFundAccountResponseBodyData {
	s.CreateDate = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountAdminAccountId(v string) *ListFundAccountResponseBodyData {
	s.FundAccountAdminAccountId = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountAdminAccountName(v string) *ListFundAccountResponseBodyData {
	s.FundAccountAdminAccountName = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountId(v string) *ListFundAccountResponseBodyData {
	s.FundAccountId = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountName(v string) *ListFundAccountResponseBodyData {
	s.FundAccountName = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountOwnerAccountId(v string) *ListFundAccountResponseBodyData {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountStatus(v string) *ListFundAccountResponseBodyData {
	s.FundAccountStatus = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetFundAccountType(v string) *ListFundAccountResponseBodyData {
	s.FundAccountType = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetNbid(v string) *ListFundAccountResponseBodyData {
	s.Nbid = &v
	return s
}

func (s *ListFundAccountResponseBodyData) SetPermissions(v []*string) *ListFundAccountResponseBodyData {
	s.Permissions = v
	return s
}

func (s *ListFundAccountResponseBodyData) SetSite(v string) *ListFundAccountResponseBodyData {
	s.Site = &v
	return s
}

type ListFundAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFundAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFundAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFundAccountResponse) GoString() string {
	return s.String()
}

func (s *ListFundAccountResponse) SetHeaders(v map[string]*string) *ListFundAccountResponse {
	s.Headers = v
	return s
}

func (s *ListFundAccountResponse) SetStatusCode(v int32) *ListFundAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFundAccountResponse) SetBody(v *ListFundAccountResponseBody) *ListFundAccountResponse {
	s.Body = v
	return s
}

type ListFundAccountPayRelationRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12323123
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFundAccountPayRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFundAccountPayRelationRequest) GoString() string {
	return s.String()
}

func (s *ListFundAccountPayRelationRequest) SetCurrentPage(v int32) *ListFundAccountPayRelationRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListFundAccountPayRelationRequest) SetFundAccountId(v string) *ListFundAccountPayRelationRequest {
	s.FundAccountId = &v
	return s
}

func (s *ListFundAccountPayRelationRequest) SetNbid(v string) *ListFundAccountPayRelationRequest {
	s.Nbid = &v
	return s
}

func (s *ListFundAccountPayRelationRequest) SetPageSize(v int32) *ListFundAccountPayRelationRequest {
	s.PageSize = &v
	return s
}

func (s *ListFundAccountPayRelationRequest) SetStatus(v string) *ListFundAccountPayRelationRequest {
	s.Status = &v
	return s
}

type ListFundAccountPayRelationResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*ListFundAccountPayRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFundAccountPayRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFundAccountPayRelationResponseBody) GoString() string {
	return s.String()
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

type ListFundAccountPayRelationResponseBodyData struct {
	// example:
	//
	// 32812132121
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 云某的名称
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 213231232
	Ecid *string `json:"Ecid,omitempty" xml:"Ecid,omitempty"`
	// example:
	//
	// 2024-12-01 12:00:10
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// example:
	//
	// 123231213
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 312328912
	FundAccountOwnerAccountId *string `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// 2025-01-01 12:12:12
	IneffectiveTime *string `json:"IneffectiveTime,omitempty" xml:"IneffectiveTime,omitempty"`
	// example:
	//
	// 2684210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 云某的名称
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// example:
	//
	// 1232343423
	OperatorNo *string `json:"OperatorNo,omitempty" xml:"OperatorNo,omitempty"`
	// example:
	//
	// aliyun_pk
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// example:
	//
	// PAYMENT
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// example:
	//
	// 26842
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFundAccountPayRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFundAccountPayRelationResponseBodyData) GoString() string {
	return s.String()
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

type ListFundAccountPayRelationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFundAccountPayRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFundAccountPayRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFundAccountPayRelationResponse) GoString() string {
	return s.String()
}

func (s *ListFundAccountPayRelationResponse) SetHeaders(v map[string]*string) *ListFundAccountPayRelationResponse {
	s.Headers = v
	return s
}

func (s *ListFundAccountPayRelationResponse) SetStatusCode(v int32) *ListFundAccountPayRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFundAccountPayRelationResponse) SetBody(v *ListFundAccountPayRelationResponseBody) *ListFundAccountPayRelationResponse {
	s.Body = v
	return s
}

type ListReportDefinitionsRequest struct {
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s ListReportDefinitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListReportDefinitionsRequest) GoString() string {
	return s.String()
}

func (s *ListReportDefinitionsRequest) SetNbid(v string) *ListReportDefinitionsRequest {
	s.Nbid = &v
	return s
}

type ListReportDefinitionsResponseBody struct {
	Metadata          interface{}                                           `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	ReportDefinitions []*ListReportDefinitionsResponseBodyReportDefinitions `json:"ReportDefinitions,omitempty" xml:"ReportDefinitions,omitempty" type:"Repeated"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListReportDefinitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListReportDefinitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListReportDefinitionsResponseBody) SetMetadata(v interface{}) *ListReportDefinitionsResponseBody {
	s.Metadata = v
	return s
}

func (s *ListReportDefinitionsResponseBody) SetReportDefinitions(v []*ListReportDefinitionsResponseBodyReportDefinitions) *ListReportDefinitionsResponseBody {
	s.ReportDefinitions = v
	return s
}

func (s *ListReportDefinitionsResponseBody) SetRequestId(v string) *ListReportDefinitionsResponseBody {
	s.RequestId = &v
	return s
}

type ListReportDefinitionsResponseBodyReportDefinitions struct {
	// example:
	//
	// 2025-05
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// example:
	//
	// oss-bill
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// 1234567812345678
	OssBucketOwnerAccountId *int64 `json:"OssBucketOwnerAccountId,omitempty" xml:"OssBucketOwnerAccountId,omitempty"`
	// example:
	//
	// bill/
	OssBucketPath *string `json:"OssBucketPath,omitempty" xml:"OssBucketPath,omitempty"`
	// example:
	//
	// OSS
	ReportSourceName *string `json:"ReportSourceName,omitempty" xml:"ReportSourceName,omitempty"`
	// example:
	//
	// OSS
	ReportSourceType *string `json:"ReportSourceType,omitempty" xml:"ReportSourceType,omitempty"`
	// example:
	//
	// 123321
	ReportTaskId *int64 `json:"ReportTaskId,omitempty" xml:"ReportTaskId,omitempty"`
	// example:
	//
	// BillingItemDetailForBillingPeriod
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// example:
	//
	// 2025-05-21 10:36:31
	SubscribeCreateTime *string `json:"SubscribeCreateTime,omitempty" xml:"SubscribeCreateTime,omitempty"`
}

func (s ListReportDefinitionsResponseBodyReportDefinitions) String() string {
	return tea.Prettify(s)
}

func (s ListReportDefinitionsResponseBodyReportDefinitions) GoString() string {
	return s.String()
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetBeginBillingCycle(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.BeginBillingCycle = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetOssBucketName(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.OssBucketName = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetOssBucketOwnerAccountId(v int64) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.OssBucketOwnerAccountId = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetOssBucketPath(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.OssBucketPath = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetReportSourceName(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.ReportSourceName = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetReportSourceType(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.ReportSourceType = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetReportTaskId(v int64) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.ReportTaskId = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetReportType(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.ReportType = &v
	return s
}

func (s *ListReportDefinitionsResponseBodyReportDefinitions) SetSubscribeCreateTime(v string) *ListReportDefinitionsResponseBodyReportDefinitions {
	s.SubscribeCreateTime = &v
	return s
}

type ListReportDefinitionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReportDefinitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReportDefinitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListReportDefinitionsResponse) GoString() string {
	return s.String()
}

func (s *ListReportDefinitionsResponse) SetHeaders(v map[string]*string) *ListReportDefinitionsResponse {
	s.Headers = v
	return s
}

func (s *ListReportDefinitionsResponse) SetStatusCode(v int32) *ListReportDefinitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReportDefinitionsResponse) SetBody(v *ListReportDefinitionsResponseBody) *ListReportDefinitionsResponse {
	s.Body = v
	return s
}

type ModifyCostCenterRequest struct {
	// This parameter is required.
	CostCenterEntityList []*ModifyCostCenterRequestCostCenterEntityList `json:"CostCenterEntityList,omitempty" xml:"CostCenterEntityList,omitempty" type:"Repeated"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s ModifyCostCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostCenterRequest) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterRequest) SetCostCenterEntityList(v []*ModifyCostCenterRequestCostCenterEntityList) *ModifyCostCenterRequest {
	s.CostCenterEntityList = v
	return s
}

func (s *ModifyCostCenterRequest) SetNbid(v string) *ModifyCostCenterRequest {
	s.Nbid = &v
	return s
}

type ModifyCostCenterRequestCostCenterEntityList struct {
	// This parameter is required.
	//
	// example:
	//
	// 485938
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// This parameter is required.
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s ModifyCostCenterRequestCostCenterEntityList) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostCenterRequestCostCenterEntityList) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterRequestCostCenterEntityList) SetCostCenterId(v int64) *ModifyCostCenterRequestCostCenterEntityList {
	s.CostCenterId = &v
	return s
}

func (s *ModifyCostCenterRequestCostCenterEntityList) SetCostCenterName(v string) *ModifyCostCenterRequestCostCenterEntityList {
	s.CostCenterName = &v
	return s
}

func (s *ModifyCostCenterRequestCostCenterEntityList) SetOwnerAccountId(v int64) *ModifyCostCenterRequestCostCenterEntityList {
	s.OwnerAccountId = &v
	return s
}

type ModifyCostCenterShrinkRequest struct {
	// This parameter is required.
	CostCenterEntityListShrink *string `json:"CostCenterEntityList,omitempty" xml:"CostCenterEntityList,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s ModifyCostCenterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostCenterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterShrinkRequest) SetCostCenterEntityListShrink(v string) *ModifyCostCenterShrinkRequest {
	s.CostCenterEntityListShrink = &v
	return s
}

func (s *ModifyCostCenterShrinkRequest) SetNbid(v string) *ModifyCostCenterShrinkRequest {
	s.Nbid = &v
	return s
}

type ModifyCostCenterResponseBody struct {
	CostCenterOperateDto []*ModifyCostCenterResponseBodyCostCenterOperateDto `json:"CostCenterOperateDto,omitempty" xml:"CostCenterOperateDto,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCostCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostCenterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterResponseBody) SetCostCenterOperateDto(v []*ModifyCostCenterResponseBodyCostCenterOperateDto) *ModifyCostCenterResponseBody {
	s.CostCenterOperateDto = v
	return s
}

func (s *ModifyCostCenterResponseBody) SetMetadata(v interface{}) *ModifyCostCenterResponseBody {
	s.Metadata = v
	return s
}

func (s *ModifyCostCenterResponseBody) SetRequestId(v string) *ModifyCostCenterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCostCenterResponseBodyCostCenterOperateDto struct {
	// example:
	//
	// 485938
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s ModifyCostCenterResponseBodyCostCenterOperateDto) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostCenterResponseBodyCostCenterOperateDto) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterResponseBodyCostCenterOperateDto) SetCostCenterId(v int64) *ModifyCostCenterResponseBodyCostCenterOperateDto {
	s.CostCenterId = &v
	return s
}

func (s *ModifyCostCenterResponseBodyCostCenterOperateDto) SetIsSuccess(v bool) *ModifyCostCenterResponseBodyCostCenterOperateDto {
	s.IsSuccess = &v
	return s
}

func (s *ModifyCostCenterResponseBodyCostCenterOperateDto) SetOwnerAccountId(v int64) *ModifyCostCenterResponseBodyCostCenterOperateDto {
	s.OwnerAccountId = &v
	return s
}

type ModifyCostCenterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCostCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCostCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostCenterResponse) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterResponse) SetHeaders(v map[string]*string) *ModifyCostCenterResponse {
	s.Headers = v
	return s
}

func (s *ModifyCostCenterResponse) SetStatusCode(v int32) *ModifyCostCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCostCenterResponse) SetBody(v *ModifyCostCenterResponseBody) *ModifyCostCenterResponse {
	s.Body = v
	return s
}

type QueryCostCenterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage    *int32                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds []*QueryCostCenterRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
}

func (s QueryCostCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterRequest) GoString() string {
	return s.String()
}

func (s *QueryCostCenterRequest) SetCurrentPage(v int32) *QueryCostCenterRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryCostCenterRequest) SetEcIdAccountIds(v []*QueryCostCenterRequestEcIdAccountIds) *QueryCostCenterRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *QueryCostCenterRequest) SetNbid(v string) *QueryCostCenterRequest {
	s.Nbid = &v
	return s
}

func (s *QueryCostCenterRequest) SetOwnerAccountId(v int64) *QueryCostCenterRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterRequest) SetPageSize(v int32) *QueryCostCenterRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCostCenterRequest) SetParentCostCenterId(v int64) *QueryCostCenterRequest {
	s.ParentCostCenterId = &v
	return s
}

type QueryCostCenterRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1004064243473974
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s QueryCostCenterRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *QueryCostCenterRequestEcIdAccountIds) SetAccountIds(v []*int64) *QueryCostCenterRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *QueryCostCenterRequestEcIdAccountIds) SetEcId(v string) *QueryCostCenterRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type QueryCostCenterShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
}

func (s QueryCostCenterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCostCenterShrinkRequest) SetCurrentPage(v int32) *QueryCostCenterShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) SetEcIdAccountIdsShrink(v string) *QueryCostCenterShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) SetNbid(v string) *QueryCostCenterShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) SetOwnerAccountId(v int64) *QueryCostCenterShrinkRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) SetPageSize(v int32) *QueryCostCenterShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) SetParentCostCenterId(v int64) *QueryCostCenterShrinkRequest {
	s.ParentCostCenterId = &v
	return s
}

type QueryCostCenterResponseBody struct {
	CostCenterDtoList []*QueryCostCenterResponseBodyCostCenterDtoList `json:"CostCenterDtoList,omitempty" xml:"CostCenterDtoList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryCostCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResponseBody) SetCostCenterDtoList(v []*QueryCostCenterResponseBodyCostCenterDtoList) *QueryCostCenterResponseBody {
	s.CostCenterDtoList = v
	return s
}

func (s *QueryCostCenterResponseBody) SetCurrentPage(v int32) *QueryCostCenterResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryCostCenterResponseBody) SetMetadata(v interface{}) *QueryCostCenterResponseBody {
	s.Metadata = v
	return s
}

func (s *QueryCostCenterResponseBody) SetPageSize(v int32) *QueryCostCenterResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryCostCenterResponseBody) SetRequestId(v string) *QueryCostCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostCenterResponseBody) SetTotalCount(v int32) *QueryCostCenterResponseBody {
	s.TotalCount = &v
	return s
}

type QueryCostCenterResponseBodyCostCenterDtoList struct {
	// example:
	//
	// 15945703968#
	CostCenterCode *string `json:"CostCenterCode,omitempty" xml:"CostCenterCode,omitempty"`
	// example:
	//
	// 485938
	CostCenterId   *int64  `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// example:
	//
	// loose
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// example:
	//
	// -1
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
	// example:
	//
	// 485996
	PrevCostCenterId *int64 `json:"PrevCostCenterId,omitempty" xml:"PrevCostCenterId,omitempty"`
}

func (s QueryCostCenterResponseBodyCostCenterDtoList) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterResponseBodyCostCenterDtoList) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetCostCenterCode(v string) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.CostCenterCode = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetCostCenterId(v int64) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetCostCenterName(v string) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.CostCenterName = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetLevel(v int32) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.Level = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetOwnerAccountId(v int64) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetParentCostCenterId(v int64) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.ParentCostCenterId = &v
	return s
}

func (s *QueryCostCenterResponseBodyCostCenterDtoList) SetPrevCostCenterId(v int64) *QueryCostCenterResponseBodyCostCenterDtoList {
	s.PrevCostCenterId = &v
	return s
}

type QueryCostCenterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCostCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCostCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterResponse) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResponse) SetHeaders(v map[string]*string) *QueryCostCenterResponse {
	s.Headers = v
	return s
}

func (s *QueryCostCenterResponse) SetStatusCode(v int32) *QueryCostCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCostCenterResponse) SetBody(v *QueryCostCenterResponseBody) *QueryCostCenterResponse {
	s.Body = v
	return s
}

type QueryCostCenterResourceRequest struct {
	// example:
	//
	// 123456
	CostCenterId   *int64                                          `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	EcIdAccountIds []*QueryCostCenterResourceRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// CAESEgoQCg4KCmd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1234567812345678
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s QueryCostCenterResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterResourceRequest) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResourceRequest) SetCostCenterId(v int64) *QueryCostCenterResourceRequest {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostCenterResourceRequest) SetEcIdAccountIds(v []*QueryCostCenterResourceRequestEcIdAccountIds) *QueryCostCenterResourceRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *QueryCostCenterResourceRequest) SetMaxResults(v int32) *QueryCostCenterResourceRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCostCenterResourceRequest) SetNbid(v string) *QueryCostCenterResourceRequest {
	s.Nbid = &v
	return s
}

func (s *QueryCostCenterResourceRequest) SetNextToken(v string) *QueryCostCenterResourceRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCostCenterResourceRequest) SetOwnerAccountId(v int64) *QueryCostCenterResourceRequest {
	s.OwnerAccountId = &v
	return s
}

type QueryCostCenterResourceRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1501603440974415
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s QueryCostCenterResourceRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterResourceRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResourceRequestEcIdAccountIds) SetAccountIds(v []*int64) *QueryCostCenterResourceRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *QueryCostCenterResourceRequestEcIdAccountIds) SetEcId(v string) *QueryCostCenterResourceRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type QueryCostCenterResourceResponseBody struct {
	CostCenterResourceDtoList []*QueryCostCenterResourceResponseBodyCostCenterResourceDtoList `json:"CostCenterResourceDtoList,omitempty" xml:"CostCenterResourceDtoList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryCostCenterResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterResourceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResourceResponseBody) SetCostCenterResourceDtoList(v []*QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) *QueryCostCenterResourceResponseBody {
	s.CostCenterResourceDtoList = v
	return s
}

func (s *QueryCostCenterResourceResponseBody) SetMaxResults(v int32) *QueryCostCenterResourceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryCostCenterResourceResponseBody) SetMetadata(v interface{}) *QueryCostCenterResourceResponseBody {
	s.Metadata = v
	return s
}

func (s *QueryCostCenterResourceResponseBody) SetNextToken(v string) *QueryCostCenterResourceResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCostCenterResourceResponseBody) SetRequestId(v string) *QueryCostCenterResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBody) SetTotalCount(v int32) *QueryCostCenterResourceResponseBody {
	s.TotalCount = &v
	return s
}

type QueryCostCenterResourceResponseBodyCostCenterResourceDtoList struct {
	// example:
	//
	// test
	ApportionItemCode *string `json:"ApportionItemCode,omitempty" xml:"ApportionItemCode,omitempty"`
	// example:
	//
	// test
	ApportionItemName *string `json:"ApportionItemName,omitempty" xml:"ApportionItemName,omitempty"`
	// example:
	//
	// otsbag
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// example:
	//
	// code
	CostCenterCode *string `json:"CostCenterCode,omitempty" xml:"CostCenterCode,omitempty"`
	// example:
	//
	// 2025-05-18 12:12:25
	CostCenterCreateTime *string `json:"CostCenterCreateTime,omitempty" xml:"CostCenterCreateTime,omitempty"`
	// example:
	//
	// 123456
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// example:
	//
	// test
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// example:
	//
	// 2025-05-18 16:12:25
	CostCenterUpdateTime *string `json:"CostCenterUpdateTime,omitempty" xml:"CostCenterUpdateTime,omitempty"`
	// example:
	//
	// 1234567812345678
	OwnerAccountId   *int64  `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	OwnerAccountName *string `json:"OwnerAccountName,omitempty" xml:"OwnerAccountName,omitempty"`
	// example:
	//
	// 123456
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
	// example:
	//
	// rds
	PipCode       *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	PipName       *string `json:"PipName,omitempty" xml:"PipName,omitempty"`
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// OSSBAG-cn-v0h1s4hma018
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// testResource
	ResourceNick *string `json:"ResourceNick,omitempty" xml:"ResourceNick,omitempty"`
	// example:
	//
	// MANUAL_ALLOCATE
	ResourceSource *string `json:"ResourceSource,omitempty" xml:"ResourceSource,omitempty"`
	// example:
	//
	// tag
	ResourceTag *string `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty"`
	// example:
	//
	// FPT_ossbag_absolute_Storage_bj
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 2025-05-18 16:12:25
	ResourceUpdateTime *string `json:"ResourceUpdateTime,omitempty" xml:"ResourceUpdateTime,omitempty"`
	// example:
	//
	// 1234567812345678
	ResourceUserId *int64 `json:"ResourceUserId,omitempty" xml:"ResourceUserId,omitempty"`
	// example:
	//
	// test@test.aliyun.com
	ResourceUserName *string `json:"ResourceUserName,omitempty" xml:"ResourceUserName,omitempty"`
	// example:
	//
	// -1
	RootCostCenterId *int64 `json:"RootCostCenterId,omitempty" xml:"RootCostCenterId,omitempty"`
}

func (s QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetApportionItemCode(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ApportionItemCode = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetApportionItemName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ApportionItemName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCommodityCode(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CommodityCode = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCommodityName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CommodityName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCostCenterCode(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CostCenterCode = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCostCenterCreateTime(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CostCenterCreateTime = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCostCenterId(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCostCenterName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CostCenterName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCostCenterUpdateTime(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CostCenterUpdateTime = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetOwnerAccountId(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetOwnerAccountName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.OwnerAccountName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetParentCostCenterId(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ParentCostCenterId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetPipCode(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.PipCode = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetPipName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.PipName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceGroup(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceGroup = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceId(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceNick(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceNick = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceSource(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceSource = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceTag(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceTag = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceType(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceType = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceUpdateTime(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceUpdateTime = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceUserId(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceUserId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceUserName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceUserName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetRootCostCenterId(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.RootCostCenterId = &v
	return s
}

type QueryCostCenterResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCostCenterResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCostCenterResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCostCenterResourceResponse) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResourceResponse) SetHeaders(v map[string]*string) *QueryCostCenterResourceResponse {
	s.Headers = v
	return s
}

func (s *QueryCostCenterResourceResponse) SetStatusCode(v int32) *QueryCostCenterResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCostCenterResourceResponse) SetBody(v *QueryCostCenterResourceResponseBody) *QueryCostCenterResourceResponse {
	s.Body = v
	return s
}

type SetFundAccountCreditAmountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 500
	CreditAmount *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 1232312
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
}

func (s SetFundAccountCreditAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s SetFundAccountCreditAmountRequest) GoString() string {
	return s.String()
}

func (s *SetFundAccountCreditAmountRequest) SetCreditAmount(v string) *SetFundAccountCreditAmountRequest {
	s.CreditAmount = &v
	return s
}

func (s *SetFundAccountCreditAmountRequest) SetCurrency(v string) *SetFundAccountCreditAmountRequest {
	s.Currency = &v
	return s
}

func (s *SetFundAccountCreditAmountRequest) SetFundAccountId(v int64) *SetFundAccountCreditAmountRequest {
	s.FundAccountId = &v
	return s
}

type SetFundAccountCreditAmountResponseBody struct {
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFundAccountCreditAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetFundAccountCreditAmountResponseBody) GoString() string {
	return s.String()
}

func (s *SetFundAccountCreditAmountResponseBody) SetMetadata(v interface{}) *SetFundAccountCreditAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *SetFundAccountCreditAmountResponseBody) SetRequestId(v string) *SetFundAccountCreditAmountResponseBody {
	s.RequestId = &v
	return s
}

type SetFundAccountCreditAmountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFundAccountCreditAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFundAccountCreditAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s SetFundAccountCreditAmountResponse) GoString() string {
	return s.String()
}

func (s *SetFundAccountCreditAmountResponse) SetHeaders(v map[string]*string) *SetFundAccountCreditAmountResponse {
	s.Headers = v
	return s
}

func (s *SetFundAccountCreditAmountResponse) SetStatusCode(v int32) *SetFundAccountCreditAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFundAccountCreditAmountResponse) SetBody(v *SetFundAccountCreditAmountResponseBody) *SetFundAccountCreditAmountResponse {
	s.Body = v
	return s
}

type SetFundAccountLowAvailableAmountAlarmRequest struct {
	// example:
	//
	// 12321213
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 100
	ThresholdAmount *string `json:"ThresholdAmount,omitempty" xml:"ThresholdAmount,omitempty"`
}

func (s SetFundAccountLowAvailableAmountAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s SetFundAccountLowAvailableAmountAlarmRequest) GoString() string {
	return s.String()
}

func (s *SetFundAccountLowAvailableAmountAlarmRequest) SetFundAccountId(v int64) *SetFundAccountLowAvailableAmountAlarmRequest {
	s.FundAccountId = &v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmRequest) SetThresholdAmount(v string) *SetFundAccountLowAvailableAmountAlarmRequest {
	s.ThresholdAmount = &v
	return s
}

type SetFundAccountLowAvailableAmountAlarmResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 08108BF5-1AA3-518E-9986-95A3616E8DA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFundAccountLowAvailableAmountAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetFundAccountLowAvailableAmountAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *SetFundAccountLowAvailableAmountAlarmResponseBody) SetData(v bool) *SetFundAccountLowAvailableAmountAlarmResponseBody {
	s.Data = &v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmResponseBody) SetMetadata(v interface{}) *SetFundAccountLowAvailableAmountAlarmResponseBody {
	s.Metadata = v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmResponseBody) SetRequestId(v string) *SetFundAccountLowAvailableAmountAlarmResponseBody {
	s.RequestId = &v
	return s
}

type SetFundAccountLowAvailableAmountAlarmResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFundAccountLowAvailableAmountAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFundAccountLowAvailableAmountAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s SetFundAccountLowAvailableAmountAlarmResponse) GoString() string {
	return s.String()
}

func (s *SetFundAccountLowAvailableAmountAlarmResponse) SetHeaders(v map[string]*string) *SetFundAccountLowAvailableAmountAlarmResponse {
	s.Headers = v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmResponse) SetStatusCode(v int32) *SetFundAccountLowAvailableAmountAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFundAccountLowAvailableAmountAlarmResponse) SetBody(v *SetFundAccountLowAvailableAmountAlarmResponseBody) *SetFundAccountLowAvailableAmountAlarmResponse {
	s.Body = v
	return s
}

type SetSavingPlanUserDeductRuleRequest struct {
	EcIdAccountIds  []*SetSavingPlanUserDeductRuleRequestEcIdAccountIds  `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	Nbid            *string                                              `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	SpnInstanceCode *string                                              `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
	UserDeductRules []*SetSavingPlanUserDeductRuleRequestUserDeductRules `json:"UserDeductRules,omitempty" xml:"UserDeductRules,omitempty" type:"Repeated"`
}

func (s SetSavingPlanUserDeductRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleRequest) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleRequest) SetEcIdAccountIds(v []*SetSavingPlanUserDeductRuleRequestEcIdAccountIds) *SetSavingPlanUserDeductRuleRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequest) SetNbid(v string) *SetSavingPlanUserDeductRuleRequest {
	s.Nbid = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequest) SetSpnInstanceCode(v string) *SetSavingPlanUserDeductRuleRequest {
	s.SpnInstanceCode = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequest) SetUserDeductRules(v []*SetSavingPlanUserDeductRuleRequestUserDeductRules) *SetSavingPlanUserDeductRuleRequest {
	s.UserDeductRules = v
	return s
}

type SetSavingPlanUserDeductRuleRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	EcId       *string  `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s SetSavingPlanUserDeductRuleRequestEcIdAccountIds) String() string {
	return tea.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleRequestEcIdAccountIds) SetAccountIds(v []*int64) *SetSavingPlanUserDeductRuleRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequestEcIdAccountIds) SetEcId(v string) *SetSavingPlanUserDeductRuleRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

type SetSavingPlanUserDeductRuleRequestUserDeductRules struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ModuleCode    *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	SkipDeduct    *bool   `json:"SkipDeduct,omitempty" xml:"SkipDeduct,omitempty"`
}

func (s SetSavingPlanUserDeductRuleRequestUserDeductRules) String() string {
	return tea.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleRequestUserDeductRules) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleRequestUserDeductRules) SetCommodityCode(v string) *SetSavingPlanUserDeductRuleRequestUserDeductRules {
	s.CommodityCode = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequestUserDeductRules) SetModuleCode(v string) *SetSavingPlanUserDeductRuleRequestUserDeductRules {
	s.ModuleCode = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleRequestUserDeductRules) SetSkipDeduct(v bool) *SetSavingPlanUserDeductRuleRequestUserDeductRules {
	s.SkipDeduct = &v
	return s
}

type SetSavingPlanUserDeductRuleShrinkRequest struct {
	EcIdAccountIdsShrink  *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                  *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	SpnInstanceCode       *string `json:"SpnInstanceCode,omitempty" xml:"SpnInstanceCode,omitempty"`
	UserDeductRulesShrink *string `json:"UserDeductRules,omitempty" xml:"UserDeductRules,omitempty"`
}

func (s SetSavingPlanUserDeductRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) SetEcIdAccountIdsShrink(v string) *SetSavingPlanUserDeductRuleShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) SetNbid(v string) *SetSavingPlanUserDeductRuleShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) SetSpnInstanceCode(v string) *SetSavingPlanUserDeductRuleShrinkRequest {
	s.SpnInstanceCode = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleShrinkRequest) SetUserDeductRulesShrink(v string) *SetSavingPlanUserDeductRuleShrinkRequest {
	s.UserDeductRulesShrink = &v
	return s
}

type SetSavingPlanUserDeductRuleResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSavingPlanUserDeductRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleResponseBody) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleResponseBody) SetData(v bool) *SetSavingPlanUserDeductRuleResponseBody {
	s.Data = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleResponseBody) SetRequestId(v string) *SetSavingPlanUserDeductRuleResponseBody {
	s.RequestId = &v
	return s
}

type SetSavingPlanUserDeductRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSavingPlanUserDeductRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSavingPlanUserDeductRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleResponse) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleResponse) SetHeaders(v map[string]*string) *SetSavingPlanUserDeductRuleResponse {
	s.Headers = v
	return s
}

func (s *SetSavingPlanUserDeductRuleResponse) SetStatusCode(v int32) *SetSavingPlanUserDeductRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleResponse) SetBody(v *SetSavingPlanUserDeductRuleResponseBody) *SetSavingPlanUserDeductRuleResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":                 tea.String("business.aliyuncs.com"),
		"cn-shanghai":                 tea.String("business.aliyuncs.com"),
		"ap-southeast-1":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-1":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"cn-beijing":                  tea.String("business.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("business.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("business.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("business.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("business.aliyuncs.com"),
		"cn-chengdu":                  tea.String("business.aliyuncs.com"),
		"cn-edge-1":                   tea.String("business.aliyuncs.com"),
		"cn-fujian":                   tea.String("business.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("business.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("business.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("business.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("business.aliyuncs.com"),
		"cn-hongkong":                 tea.String("business.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("business.aliyuncs.com"),
		"cn-huhehaote":                tea.String("business.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("business.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("business.aliyuncs.com"),
		"cn-qingdao":                  tea.String("business.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("business.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("business.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("business.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("business.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("business.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("business.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("business.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("business.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("business.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("business.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("business.aliyuncs.com"),
		"cn-wuhan":                    tea.String("business.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("business.aliyuncs.com"),
		"cn-yushanfang":               tea.String("business.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("business.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("business.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("business.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("business.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("business.aliyuncs.com"),
		"eu-central-1":                tea.String("business.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":                   tea.String("business.ap-southeast-1.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("business.ap-southeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("business.ap-southeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"us-east-1":                   tea.String("business.ap-southeast-1.aliyuncs.com"),
		"us-west-1":                   tea.String("business.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("bssopenapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加优惠券抵扣标签
//
// @param tmpReq - AddCouponDeductTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCouponDeductTagResponse
func (client *Client) AddCouponDeductTagWithOptions(tmpReq *AddCouponDeductTagRequest, runtime *util.RuntimeOptions) (_result *AddCouponDeductTagResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddCouponDeductTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CouponId)) {
		query["CouponId"] = request.CouponId
	}

	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIdsShrink)) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCouponDeductTag"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCouponDeductTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加优惠券抵扣标签
//
// @param request - AddCouponDeductTagRequest
//
// @return AddCouponDeductTagResponse
func (client *Client) AddCouponDeductTag(request *AddCouponDeductTagRequest) (_result *AddCouponDeductTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCouponDeductTagResponse{}
	_body, _err := client.AddCouponDeductTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消资金账户低额预警
//
// @param request - CancelFundAccountLowAvailableAmountAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelFundAccountLowAvailableAmountAlarmResponse
func (client *Client) CancelFundAccountLowAvailableAmountAlarmWithOptions(request *CancelFundAccountLowAvailableAmountAlarmRequest, runtime *util.RuntimeOptions) (_result *CancelFundAccountLowAvailableAmountAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelFundAccountLowAvailableAmountAlarm"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消资金账户低额预警
//
// @param request - CancelFundAccountLowAvailableAmountAlarmRequest
//
// @return CancelFundAccountLowAvailableAmountAlarmResponse
func (client *Client) CancelFundAccountLowAvailableAmountAlarm(request *CancelFundAccountLowAvailableAmountAlarmRequest) (_result *CancelFundAccountLowAvailableAmountAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.CancelFundAccountLowAvailableAmountAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建财务单元
//
// @param tmpReq - CreateCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCostCenterResponse
func (client *Client) CreateCostCenterWithOptions(tmpReq *CreateCostCenterRequest, runtime *util.RuntimeOptions) (_result *CreateCostCenterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCostCenterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CostCenterEntityList)) {
		request.CostCenterEntityListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CostCenterEntityList, tea.String("CostCenterEntityList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CostCenterEntityListShrink)) {
		query["CostCenterEntityList"] = request.CostCenterEntityListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCostCenter"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCostCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建财务单元
//
// @param request - CreateCostCenterRequest
//
// @return CreateCostCenterResponse
func (client *Client) CreateCostCenter(request *CreateCostCenterRequest) (_result *CreateCostCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCostCenterResponse{}
	_body, _err := client.CreateCostCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资金账户付款关系
//
// @param tmpReq - CreateFundAccountPayRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFundAccountPayRelationResponse
func (client *Client) CreateFundAccountPayRelationWithOptions(tmpReq *CreateFundAccountPayRelationRequest, runtime *util.RuntimeOptions) (_result *CreateFundAccountPayRelationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateFundAccountPayRelationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIdsShrink)) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFundAccountPayRelation"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFundAccountPayRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资金账户付款关系
//
// @param request - CreateFundAccountPayRelationRequest
//
// @return CreateFundAccountPayRelationResponse
func (client *Client) CreateFundAccountPayRelation(request *CreateFundAccountPayRelationRequest) (_result *CreateFundAccountPayRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFundAccountPayRelationResponse{}
	_body, _err := client.CreateFundAccountPayRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资金账户划拨/回收
//
// @param request - CreateFundAccountTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFundAccountTransferResponse
func (client *Client) CreateFundAccountTransferWithOptions(request *CreateFundAccountTransferRequest, runtime *util.RuntimeOptions) (_result *CreateFundAccountTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.Currency)) {
		body["Currency"] = request.Currency
	}

	if !tea.BoolValue(util.IsUnset(request.FinanceType)) {
		body["FinanceType"] = request.FinanceType
	}

	if !tea.BoolValue(util.IsUnset(request.FromFundAccountId)) {
		body["FromFundAccountId"] = request.FromFundAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ToFundAccountId)) {
		body["ToFundAccountId"] = request.ToFundAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.TransferType)) {
		body["TransferType"] = request.TransferType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFundAccountTransfer"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFundAccountTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资金账户划拨/回收
//
// @param request - CreateFundAccountTransferRequest
//
// @return CreateFundAccountTransferResponse
func (client *Client) CreateFundAccountTransfer(request *CreateFundAccountTransferRequest) (_result *CreateFundAccountTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFundAccountTransferResponse{}
	_body, _err := client.CreateFundAccountTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建账单订阅
//
// @param request - CreateReportDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReportDefinitionResponse
func (client *Client) CreateReportDefinitionWithOptions(request *CreateReportDefinitionRequest, runtime *util.RuntimeOptions) (_result *CreateReportDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginBillingCycle)) {
		query["BeginBillingCycle"] = request.BeginBillingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketOwnerAccountId)) {
		query["OssBucketOwnerAccountId"] = request.OssBucketOwnerAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketPath)) {
		query["OssBucketPath"] = request.OssBucketPath
	}

	if !tea.BoolValue(util.IsUnset(request.ReportType)) {
		query["ReportType"] = request.ReportType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.McProject)) {
		body["McProject"] = request.McProject
	}

	if !tea.BoolValue(util.IsUnset(request.McTableName)) {
		body["McTableName"] = request.McTableName
	}

	if !tea.BoolValue(util.IsUnset(request.ReportSourceType)) {
		body["ReportSourceType"] = request.ReportSourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateReportDefinition"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateReportDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建账单订阅
//
// @param request - CreateReportDefinitionRequest
//
// @return CreateReportDefinitionResponse
func (client *Client) CreateReportDefinition(request *CreateReportDefinitionRequest) (_result *CreateReportDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateReportDefinitionResponse{}
	_body, _err := client.CreateReportDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除财务单元
//
// @param request - DeleteCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCostCenterResponse
func (client *Client) DeleteCostCenterWithOptions(request *DeleteCostCenterRequest, runtime *util.RuntimeOptions) (_result *DeleteCostCenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CostCenterId)) {
		query["CostCenterId"] = request.CostCenterId
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccountId)) {
		query["OwnerAccountId"] = request.OwnerAccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCostCenter"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCostCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除财务单元
//
// @param request - DeleteCostCenterRequest
//
// @return DeleteCostCenterResponse
func (client *Client) DeleteCostCenter(request *DeleteCostCenterRequest) (_result *DeleteCostCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCostCenterResponse{}
	_body, _err := client.DeleteCostCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除优惠券的抵扣标签
//
// @param tmpReq - DeleteCouponDeductTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCouponDeductTagResponse
func (client *Client) DeleteCouponDeductTagWithOptions(tmpReq *DeleteCouponDeductTagRequest, runtime *util.RuntimeOptions) (_result *DeleteCouponDeductTagResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteCouponDeductTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagKeys)) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, tea.String("TagKeys"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CouponId)) {
		query["CouponId"] = request.CouponId
	}

	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIdsShrink)) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeysShrink)) {
		query["TagKeys"] = request.TagKeysShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCouponDeductTag"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCouponDeductTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除优惠券的抵扣标签
//
// @param request - DeleteCouponDeductTagRequest
//
// @return DeleteCouponDeductTagResponse
func (client *Client) DeleteCouponDeductTag(request *DeleteCouponDeductTagRequest) (_result *DeleteCouponDeductTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCouponDeductTagResponse{}
	_body, _err := client.DeleteCouponDeductTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消账单订阅
//
// @param request - DeleteReportDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReportDefinitionResponse
func (client *Client) DeleteReportDefinitionWithOptions(request *DeleteReportDefinitionRequest, runtime *util.RuntimeOptions) (_result *DeleteReportDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	if !tea.BoolValue(util.IsUnset(request.ReportTaskId)) {
		query["ReportTaskId"] = request.ReportTaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteReportDefinition"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteReportDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消账单订阅
//
// @param request - DeleteReportDefinitionRequest
//
// @return DeleteReportDefinitionResponse
func (client *Client) DeleteReportDefinition(request *DeleteReportDefinitionRequest) (_result *DeleteReportDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteReportDefinitionResponse{}
	_body, _err := client.DeleteReportDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询优惠券列表
//
// @param tmpReq - DescribeCouponRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCouponResponse
func (client *Client) DescribeCouponWithOptions(tmpReq *DescribeCouponRequest, runtime *util.RuntimeOptions) (_result *DescribeCouponResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCouponShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCoupon"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCouponResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询优惠券列表
//
// @param request - DescribeCouponRequest
//
// @return DescribeCouponResponse
func (client *Client) DescribeCoupon(request *DescribeCouponRequest) (_result *DescribeCouponResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCouponResponse{}
	_body, _err := client.DescribeCouponWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询优惠券可用商品列表
//
// @param tmpReq - DescribeCouponItemListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCouponItemListResponse
func (client *Client) DescribeCouponItemListWithOptions(tmpReq *DescribeCouponItemListRequest, runtime *util.RuntimeOptions) (_result *DescribeCouponItemListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCouponItemListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCouponItemList"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCouponItemListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询优惠券可用商品列表
//
// @param request - DescribeCouponItemListRequest
//
// @return DescribeCouponItemListResponse
func (client *Client) DescribeCouponItemList(request *DescribeCouponItemListRequest) (_result *DescribeCouponItemListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCouponItemListResponse{}
	_body, _err := client.DescribeCouponItemListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取客户使用SPN的概述信息
//
// @param tmpReq - DescribeUserSpnSummaryInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserSpnSummaryInfoResponse
func (client *Client) DescribeUserSpnSummaryInfoWithOptions(tmpReq *DescribeUserSpnSummaryInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeUserSpnSummaryInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeUserSpnSummaryInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIdsShrink)) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserSpnSummaryInfo"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserSpnSummaryInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取客户使用SPN的概述信息
//
// @param request - DescribeUserSpnSummaryInfoRequest
//
// @return DescribeUserSpnSummaryInfoResponse
func (client *Client) DescribeUserSpnSummaryInfo(request *DescribeUserSpnSummaryInfoRequest) (_result *DescribeUserSpnSummaryInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserSpnSummaryInfoResponse{}
	_body, _err := client.DescribeUserSpnSummaryInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户可用金
//
// @param request - GetFundAccountAvailableAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountAvailableAmountResponse
func (client *Client) GetFundAccountAvailableAmountWithOptions(request *GetFundAccountAvailableAmountRequest, runtime *util.RuntimeOptions) (_result *GetFundAccountAvailableAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFundAccountAvailableAmount"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFundAccountAvailableAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户可用金
//
// @param request - GetFundAccountAvailableAmountRequest
//
// @return GetFundAccountAvailableAmountResponse
func (client *Client) GetFundAccountAvailableAmount(request *GetFundAccountAvailableAmountRequest) (_result *GetFundAccountAvailableAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFundAccountAvailableAmountResponse{}
	_body, _err := client.GetFundAccountAvailableAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户可分配信控额度
//
// @param request - GetFundAccountCanAllocateCreditAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanAllocateCreditAmountResponse
func (client *Client) GetFundAccountCanAllocateCreditAmountWithOptions(request *GetFundAccountCanAllocateCreditAmountRequest, runtime *util.RuntimeOptions) (_result *GetFundAccountCanAllocateCreditAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFundAccountCanAllocateCreditAmount"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFundAccountCanAllocateCreditAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户可分配信控额度
//
// @param request - GetFundAccountCanAllocateCreditAmountRequest
//
// @return GetFundAccountCanAllocateCreditAmountResponse
func (client *Client) GetFundAccountCanAllocateCreditAmount(request *GetFundAccountCanAllocateCreditAmountRequest) (_result *GetFundAccountCanAllocateCreditAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFundAccountCanAllocateCreditAmountResponse{}
	_body, _err := client.GetFundAccountCanAllocateCreditAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户可回收金额
//
// @param request - GetFundAccountCanRecycleAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanRecycleAmountResponse
func (client *Client) GetFundAccountCanRecycleAmountWithOptions(request *GetFundAccountCanRecycleAmountRequest, runtime *util.RuntimeOptions) (_result *GetFundAccountCanRecycleAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Currency)) {
		body["Currency"] = request.Currency
	}

	if !tea.BoolValue(util.IsUnset(request.RecycleFromFundAccountId)) {
		body["RecycleFromFundAccountId"] = request.RecycleFromFundAccountId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFundAccountCanRecycleAmount"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFundAccountCanRecycleAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户可回收金额
//
// @param request - GetFundAccountCanRecycleAmountRequest
//
// @return GetFundAccountCanRecycleAmountResponse
func (client *Client) GetFundAccountCanRecycleAmount(request *GetFundAccountCanRecycleAmountRequest) (_result *GetFundAccountCanRecycleAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFundAccountCanRecycleAmountResponse{}
	_body, _err := client.GetFundAccountCanRecycleAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户的可转出金额
//
// @param request - GetFundAccountCanTransferAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanTransferAmountResponse
func (client *Client) GetFundAccountCanTransferAmountWithOptions(request *GetFundAccountCanTransferAmountRequest, runtime *util.RuntimeOptions) (_result *GetFundAccountCanTransferAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Currency)) {
		body["Currency"] = request.Currency
	}

	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFundAccountCanTransferAmount"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFundAccountCanTransferAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户的可转出金额
//
// @param request - GetFundAccountCanTransferAmountRequest
//
// @return GetFundAccountCanTransferAmountResponse
func (client *Client) GetFundAccountCanTransferAmount(request *GetFundAccountCanTransferAmountRequest) (_result *GetFundAccountCanTransferAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFundAccountCanTransferAmountResponse{}
	_body, _err := client.GetFundAccountCanTransferAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户可提现金额
//
// @param request - GetFundAccountCanWithdrawAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanWithdrawAmountResponse
func (client *Client) GetFundAccountCanWithdrawAmountWithOptions(request *GetFundAccountCanWithdrawAmountRequest, runtime *util.RuntimeOptions) (_result *GetFundAccountCanWithdrawAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFundAccountCanWithdrawAmount"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFundAccountCanWithdrawAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户可提现金额
//
// @param request - GetFundAccountCanWithdrawAmountRequest
//
// @return GetFundAccountCanWithdrawAmountResponse
func (client *Client) GetFundAccountCanWithdrawAmount(request *GetFundAccountCanWithdrawAmountRequest) (_result *GetFundAccountCanWithdrawAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFundAccountCanWithdrawAmountResponse{}
	_body, _err := client.GetFundAccountCanWithdrawAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户低额预警
//
// @param request - GetFundAccountLowAvailableAmountAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) GetFundAccountLowAvailableAmountAlarmWithOptions(request *GetFundAccountLowAvailableAmountAlarmRequest, runtime *util.RuntimeOptions) (_result *GetFundAccountLowAvailableAmountAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFundAccountLowAvailableAmountAlarm"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户低额预警
//
// @param request - GetFundAccountLowAvailableAmountAlarmRequest
//
// @return GetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) GetFundAccountLowAvailableAmountAlarm(request *GetFundAccountLowAvailableAmountAlarmRequest) (_result *GetFundAccountLowAvailableAmountAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.GetFundAccountLowAvailableAmountAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户收支明细
//
// @param tmpReq - GetFundAccountTransactionDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountTransactionDetailsResponse
func (client *Client) GetFundAccountTransactionDetailsWithOptions(tmpReq *GetFundAccountTransactionDetailsRequest, runtime *util.RuntimeOptions) (_result *GetFundAccountTransactionDetailsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetFundAccountTransactionDetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TransactionChannelList)) {
		request.TransactionChannelListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransactionChannelList, tea.String("TransactionChannelList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TransactionTypeList)) {
		request.TransactionTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransactionTypeList, tea.String("TransactionTypeList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillNumber)) {
		body["BillNumber"] = request.BillNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelTransactionNumber)) {
		body["ChannelTransactionNumber"] = request.ChannelTransactionNumber
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionChannelListShrink)) {
		body["TransactionChannelList"] = request.TransactionChannelListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionDirection)) {
		body["TransactionDirection"] = request.TransactionDirection
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionNumber)) {
		body["TransactionNumber"] = request.TransactionNumber
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionType)) {
		body["TransactionType"] = request.TransactionType
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionTypeListShrink)) {
		body["TransactionTypeList"] = request.TransactionTypeListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFundAccountTransactionDetails"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFundAccountTransactionDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户收支明细
//
// @param request - GetFundAccountTransactionDetailsRequest
//
// @return GetFundAccountTransactionDetailsResponse
func (client *Client) GetFundAccountTransactionDetails(request *GetFundAccountTransactionDetailsRequest) (_result *GetFundAccountTransactionDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFundAccountTransactionDetailsResponse{}
	_body, _err := client.GetFundAccountTransactionDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节省计划及可抵扣商品信息
//
// @param tmpReq - GetSavingPlanDeductableCommodityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavingPlanDeductableCommodityResponse
func (client *Client) GetSavingPlanDeductableCommodityWithOptions(tmpReq *GetSavingPlanDeductableCommodityRequest, runtime *util.RuntimeOptions) (_result *GetSavingPlanDeductableCommodityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSavingPlanDeductableCommodityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIdsShrink)) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSavingPlanDeductableCommodity"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSavingPlanDeductableCommodityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节省计划及可抵扣商品信息
//
// @param request - GetSavingPlanDeductableCommodityRequest
//
// @return GetSavingPlanDeductableCommodityResponse
func (client *Client) GetSavingPlanDeductableCommodity(request *GetSavingPlanDeductableCommodityRequest) (_result *GetSavingPlanDeductableCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSavingPlanDeductableCommodityResponse{}
	_body, _err := client.GetSavingPlanDeductableCommodityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节省计划实例共享账号信息
//
// @param tmpReq - GetSavingPlanShareAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavingPlanShareAccountsResponse
func (client *Client) GetSavingPlanShareAccountsWithOptions(tmpReq *GetSavingPlanShareAccountsRequest, runtime *util.RuntimeOptions) (_result *GetSavingPlanShareAccountsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSavingPlanShareAccountsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIdsShrink)) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SpnInstanceCode)) {
		query["SpnInstanceCode"] = request.SpnInstanceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSavingPlanShareAccounts"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSavingPlanShareAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节省计划实例共享账号信息
//
// @param request - GetSavingPlanShareAccountsRequest
//
// @return GetSavingPlanShareAccountsResponse
func (client *Client) GetSavingPlanShareAccounts(request *GetSavingPlanShareAccountsRequest) (_result *GetSavingPlanShareAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSavingPlanShareAccountsResponse{}
	_body, _err := client.GetSavingPlanShareAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节省计划实例客户自定义规则
//
// @param tmpReq - GetSavingPlanUserDeductRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavingPlanUserDeductRuleResponse
func (client *Client) GetSavingPlanUserDeductRuleWithOptions(tmpReq *GetSavingPlanUserDeductRuleRequest, runtime *util.RuntimeOptions) (_result *GetSavingPlanUserDeductRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSavingPlanUserDeductRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIdsShrink)) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SpnInstanceCode)) {
		query["SpnInstanceCode"] = request.SpnInstanceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSavingPlanUserDeductRule"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSavingPlanUserDeductRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节省计划实例客户自定义规则
//
// @param request - GetSavingPlanUserDeductRuleRequest
//
// @return GetSavingPlanUserDeductRuleResponse
func (client *Client) GetSavingPlanUserDeductRule(request *GetSavingPlanUserDeductRuleRequest) (_result *GetSavingPlanUserDeductRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSavingPlanUserDeductRuleResponse{}
	_body, _err := client.GetSavingPlanUserDeductRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询优惠券设置的抵扣标签
//
// @param tmpReq - ListCouponDeductTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCouponDeductTagResponse
func (client *Client) ListCouponDeductTagWithOptions(tmpReq *ListCouponDeductTagRequest, runtime *util.RuntimeOptions) (_result *ListCouponDeductTagResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCouponDeductTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CouponId)) {
		query["CouponId"] = request.CouponId
	}

	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIdsShrink)) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCouponDeductTag"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCouponDeductTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询优惠券设置的抵扣标签
//
// @param request - ListCouponDeductTagRequest
//
// @return ListCouponDeductTagResponse
func (client *Client) ListCouponDeductTag(request *ListCouponDeductTagRequest) (_result *ListCouponDeductTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCouponDeductTagResponse{}
	_body, _err := client.ListCouponDeductTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户列表
//
// @param request - ListFundAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFundAccountResponse
func (client *Client) ListFundAccountWithOptions(request *ListFundAccountRequest, runtime *util.RuntimeOptions) (_result *ListFundAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryOnlyInUse)) {
		body["QueryOnlyInUse"] = request.QueryOnlyInUse
	}

	if !tea.BoolValue(util.IsUnset(request.QueryOnlyManage)) {
		body["QueryOnlyManage"] = request.QueryOnlyManage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFundAccount"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFundAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户列表
//
// @param request - ListFundAccountRequest
//
// @return ListFundAccountResponse
func (client *Client) ListFundAccount(request *ListFundAccountRequest) (_result *ListFundAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFundAccountResponse{}
	_body, _err := client.ListFundAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户的付款关系
//
// @param request - ListFundAccountPayRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFundAccountPayRelationResponse
func (client *Client) ListFundAccountPayRelationWithOptions(request *ListFundAccountPayRelationRequest, runtime *util.RuntimeOptions) (_result *ListFundAccountPayRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFundAccountPayRelation"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFundAccountPayRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户的付款关系
//
// @param request - ListFundAccountPayRelationRequest
//
// @return ListFundAccountPayRelationResponse
func (client *Client) ListFundAccountPayRelation(request *ListFundAccountPayRelationRequest) (_result *ListFundAccountPayRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFundAccountPayRelationResponse{}
	_body, _err := client.ListFundAccountPayRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看已订阅的报告列表
//
// @param request - ListReportDefinitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReportDefinitionsResponse
func (client *Client) ListReportDefinitionsWithOptions(request *ListReportDefinitionsRequest, runtime *util.RuntimeOptions) (_result *ListReportDefinitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListReportDefinitions"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListReportDefinitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看已订阅的报告列表
//
// @param request - ListReportDefinitionsRequest
//
// @return ListReportDefinitionsResponse
func (client *Client) ListReportDefinitions(request *ListReportDefinitionsRequest) (_result *ListReportDefinitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListReportDefinitionsResponse{}
	_body, _err := client.ListReportDefinitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改财务单元
//
// @param tmpReq - ModifyCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCostCenterResponse
func (client *Client) ModifyCostCenterWithOptions(tmpReq *ModifyCostCenterRequest, runtime *util.RuntimeOptions) (_result *ModifyCostCenterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyCostCenterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CostCenterEntityList)) {
		request.CostCenterEntityListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CostCenterEntityList, tea.String("CostCenterEntityList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CostCenterEntityListShrink)) {
		query["CostCenterEntityList"] = request.CostCenterEntityListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCostCenter"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyCostCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改财务单元
//
// @param request - ModifyCostCenterRequest
//
// @return ModifyCostCenterResponse
func (client *Client) ModifyCostCenter(request *ModifyCostCenterRequest) (_result *ModifyCostCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCostCenterResponse{}
	_body, _err := client.ModifyCostCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询财务单元
//
// @param tmpReq - QueryCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterResponse
func (client *Client) QueryCostCenterWithOptions(tmpReq *QueryCostCenterRequest, runtime *util.RuntimeOptions) (_result *QueryCostCenterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryCostCenterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIdsShrink)) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccountId)) {
		query["OwnerAccountId"] = request.OwnerAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCostCenterId)) {
		query["ParentCostCenterId"] = request.ParentCostCenterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCostCenter"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCostCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询财务单元
//
// @param request - QueryCostCenterRequest
//
// @return QueryCostCenterResponse
func (client *Client) QueryCostCenter(request *QueryCostCenterRequest) (_result *QueryCostCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCostCenterResponse{}
	_body, _err := client.QueryCostCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询财务单元下资源信息
//
// @param request - QueryCostCenterResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterResourceResponse
func (client *Client) QueryCostCenterResourceWithOptions(request *QueryCostCenterResourceRequest, runtime *util.RuntimeOptions) (_result *QueryCostCenterResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIds)) {
		query["EcIdAccountIds"] = request.EcIdAccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CostCenterId)) {
		body["CostCenterId"] = request.CostCenterId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccountId)) {
		body["OwnerAccountId"] = request.OwnerAccountId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCostCenterResource"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCostCenterResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询财务单元下资源信息
//
// @param request - QueryCostCenterResourceRequest
//
// @return QueryCostCenterResourceResponse
func (client *Client) QueryCostCenterResource(request *QueryCostCenterResourceRequest) (_result *QueryCostCenterResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCostCenterResourceResponse{}
	_body, _err := client.QueryCostCenterResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置资金账户的信控限额
//
// @param request - SetFundAccountCreditAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFundAccountCreditAmountResponse
func (client *Client) SetFundAccountCreditAmountWithOptions(request *SetFundAccountCreditAmountRequest, runtime *util.RuntimeOptions) (_result *SetFundAccountCreditAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreditAmount)) {
		body["CreditAmount"] = request.CreditAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Currency)) {
		body["Currency"] = request.Currency
	}

	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetFundAccountCreditAmount"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetFundAccountCreditAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置资金账户的信控限额
//
// @param request - SetFundAccountCreditAmountRequest
//
// @return SetFundAccountCreditAmountResponse
func (client *Client) SetFundAccountCreditAmount(request *SetFundAccountCreditAmountRequest) (_result *SetFundAccountCreditAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetFundAccountCreditAmountResponse{}
	_body, _err := client.SetFundAccountCreditAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置资金账户低额预警
//
// @param request - SetFundAccountLowAvailableAmountAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) SetFundAccountLowAvailableAmountAlarmWithOptions(request *SetFundAccountLowAvailableAmountAlarmRequest, runtime *util.RuntimeOptions) (_result *SetFundAccountLowAvailableAmountAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FundAccountId)) {
		body["FundAccountId"] = request.FundAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdAmount)) {
		body["ThresholdAmount"] = request.ThresholdAmount
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetFundAccountLowAvailableAmountAlarm"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置资金账户低额预警
//
// @param request - SetFundAccountLowAvailableAmountAlarmRequest
//
// @return SetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) SetFundAccountLowAvailableAmountAlarm(request *SetFundAccountLowAvailableAmountAlarmRequest) (_result *SetFundAccountLowAvailableAmountAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.SetFundAccountLowAvailableAmountAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置节省计划用户级抵扣规则
//
// @param tmpReq - SetSavingPlanUserDeductRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSavingPlanUserDeductRuleResponse
func (client *Client) SetSavingPlanUserDeductRuleWithOptions(tmpReq *SetSavingPlanUserDeductRuleRequest, runtime *util.RuntimeOptions) (_result *SetSavingPlanUserDeductRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetSavingPlanUserDeductRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EcIdAccountIds)) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, tea.String("EcIdAccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserDeductRules)) {
		request.UserDeductRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserDeductRules, tea.String("UserDeductRules"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcIdAccountIdsShrink)) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Nbid)) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpnInstanceCode)) {
		body["SpnInstanceCode"] = request.SpnInstanceCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserDeductRulesShrink)) {
		body["UserDeductRules"] = request.UserDeductRulesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSavingPlanUserDeductRule"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSavingPlanUserDeductRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置节省计划用户级抵扣规则
//
// @param request - SetSavingPlanUserDeductRuleRequest
//
// @return SetSavingPlanUserDeductRuleResponse
func (client *Client) SetSavingPlanUserDeductRule(request *SetSavingPlanUserDeductRuleRequest) (_result *SetSavingPlanUserDeductRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSavingPlanUserDeductRuleResponse{}
	_body, _err := client.SetSavingPlanUserDeductRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

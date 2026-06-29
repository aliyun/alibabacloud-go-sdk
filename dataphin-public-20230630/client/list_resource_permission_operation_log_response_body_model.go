// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePermissionOperationLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListResourcePermissionOperationLogResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListResourcePermissionOperationLogResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListResourcePermissionOperationLogResponseBody
	GetMessage() *string
	SetPageResult(v *ListResourcePermissionOperationLogResponseBodyPageResult) *ListResourcePermissionOperationLogResponseBody
	GetPageResult() *ListResourcePermissionOperationLogResponseBodyPageResult
	SetRequestId(v string) *ListResourcePermissionOperationLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListResourcePermissionOperationLogResponseBody
	GetSuccess() *bool
}

type ListResourcePermissionOperationLogResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paging query result.
	PageResult *ListResourcePermissionOperationLogResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListResourcePermissionOperationLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListResourcePermissionOperationLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListResourcePermissionOperationLogResponseBody) GetPageResult() *ListResourcePermissionOperationLogResponseBodyPageResult {
	return s.PageResult
}

func (s *ListResourcePermissionOperationLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourcePermissionOperationLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListResourcePermissionOperationLogResponseBody) SetCode(v string) *ListResourcePermissionOperationLogResponseBody {
	s.Code = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) SetHttpStatusCode(v int32) *ListResourcePermissionOperationLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) SetMessage(v string) *ListResourcePermissionOperationLogResponseBody {
	s.Message = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) SetPageResult(v *ListResourcePermissionOperationLogResponseBodyPageResult) *ListResourcePermissionOperationLogResponseBody {
	s.PageResult = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) SetRequestId(v string) *ListResourcePermissionOperationLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) SetSuccess(v bool) *ListResourcePermissionOperationLogResponseBody {
	s.Success = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcePermissionOperationLogResponseBodyPageResult struct {
	// The paginated records.
	Data []*ListResourcePermissionOperationLogResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 121
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResult) GetData() []*ListResourcePermissionOperationLogResponseBodyPageResultData {
	return s.Data
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResult) SetData(v []*ListResourcePermissionOperationLogResponseBodyPageResultData) *ListResourcePermissionOperationLogResponseBodyPageResult {
	s.Data = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResult) SetTotalCount(v int64) *ListResourcePermissionOperationLogResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResult) Validate() error {
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

type ListResourcePermissionOperationLogResponseBodyPageResultData struct {
	// The operator.
	Account *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// The authorization scope of the table. Valid values:
	//
	// - selectTable: specified table
	//
	// - projectAllTable: all tables in the project
	//
	// - bizUnitAllLogicTable: all logical tables in the business unit.
	//
	// example:
	//
	// selectTable
	AuthScope *string `json:"AuthScope,omitempty" xml:"AuthScope,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// 123133
	OperateId *int64 `json:"OperateId,omitempty" xml:"OperateId,omitempty"`
	// The operation time.
	//
	// example:
	//
	// 1710012121000
	OperateTime *int64 `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// The operation type. Valid values:
	//
	// - APPLY: Apply for permissions.
	//
	// - GRANT: Grant permissions.
	//
	// example:
	//
	// APPLY
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The validity period settings.
	Period *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The reason.
	//
	// example:
	//
	// xx测试
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The permission resource.
	ResourceInfo *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty" type:"Struct"`
	// The authorized object.
	TargetAccount *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount `json:"TargetAccount,omitempty" xml:"TargetAccount,omitempty" type:"Struct"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultData) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultData) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) GetAccount() *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount {
	return s.Account
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) GetAuthScope() *string {
	return s.AuthScope
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) GetOperateId() *int64 {
	return s.OperateId
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) GetOperateTime() *int64 {
	return s.OperateTime
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) GetOperateType() *string {
	return s.OperateType
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) GetPeriod() *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod {
	return s.Period
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) GetReason() *string {
	return s.Reason
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) GetResourceInfo() *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	return s.ResourceInfo
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) GetTargetAccount() *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount {
	return s.TargetAccount
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetAccount(v *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.Account = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetAuthScope(v string) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.AuthScope = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetOperateId(v int64) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.OperateId = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetOperateTime(v int64) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.OperateTime = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetOperateType(v string) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.OperateType = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetPeriod(v *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.Period = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetReason(v string) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.Reason = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetResourceInfo(v *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.ResourceInfo = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetTargetAccount(v *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.TargetAccount = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) Validate() error {
	if s.Account != nil {
		if err := s.Account.Validate(); err != nil {
			return err
		}
	}
	if s.Period != nil {
		if err := s.Period.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceInfo != nil {
		if err := s.ResourceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TargetAccount != nil {
		if err := s.TargetAccount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataAccount struct {
	// The account ID.
	//
	// - Individual account: the userId on the Dataphin side.
	//
	// - Production account: the UserId obtained by calling the GetProjectProduceUser operation.
	//
	// - User group: the user group ID obtained by calling the ListUserGroup operation.
	//
	// example:
	//
	// 1212131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The account name.
	//
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The authorization account type. Valid values:
	//
	// - PERSONAL: individual account
	//
	// - PRODUCE: production account
	//
	// - USER_GROUP: user group.
	//
	// example:
	//
	// PERSONAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) GetId() *string {
	return s.Id
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) GetName() *string {
	return s.Name
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) GetType() *string {
	return s.Type
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) SetId(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) SetName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) SetType(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount {
	s.Type = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) Validate() error {
	return dara.Validate(s)
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod struct {
	// The expiration time.
	//
	// example:
	//
	// 1712000000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The validity period type. Valid values:
	//
	// - CUSTOM: custom
	//
	// - LONG_TERM: permanently valid
	//
	// - DAYS_30: valid for 30 days
	//
	// - DAYS_90: valid for 90 days
	//
	// - DAYS_180: valid for 180 days.
	//
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) GetType() *string {
	return s.Type
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) SetEndTime(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod {
	s.EndTime = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) SetType(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod {
	s.Type = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) Validate() error {
	return dara.Validate(s)
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo struct {
	// The business unit information.
	BizUnitInfo *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo `json:"BizUnitInfo,omitempty" xml:"BizUnitInfo,omitempty" type:"Struct"`
	// The display name of the resource.
	//
	// example:
	//
	// tb1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The resource environment. Valid values:
	//
	// - DEV
	//
	// - PROD.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The permission resource ID.
	//
	// example:
	//
	// a.tb1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The permission resource name.
	//
	// example:
	//
	// tb1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The project information.
	ProjectInfo *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// The resource type. Valid values:
	//
	// - PHYSICAL_TABLE: physical table
	//
	// - PHYSICAL_VIEW: physical view
	//
	// - LOGICAL_TABLE: fact logical table
	//
	// - LOGICAL_VIEW: fact logical view
	//
	// - REALTIME_LOGICAL_TABLE: real-time meta table
	//
	// - REALTIME_MIRROR_TABLE: real-time meta table
	//
	// - DATASOURCE: datasource.
	//
	// example:
	//
	// PHYSICAL_TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) GetBizUnitInfo() *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo {
	return s.BizUnitInfo
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) GetEnv() *string {
	return s.Env
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) GetId() *string {
	return s.Id
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) GetName() *string {
	return s.Name
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) GetProjectInfo() *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo {
	return s.ProjectInfo
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) GetType() *string {
	return s.Type
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetBizUnitInfo(v *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.BizUnitInfo = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetDisplayName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetEnv(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetId(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetProjectInfo(v *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.ProjectInfo = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetType(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.Type = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) Validate() error {
	if s.BizUnitInfo != nil {
		if err := s.BizUnitInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ProjectInfo != nil {
		if err := s.ProjectInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo struct {
	// The display name.
	//
	// example:
	//
	// xx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The environment identifier. Valid values:
	//
	// - DEV
	//
	// - PROD.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Id
	//
	// example:
	//
	// 121323
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name.
	//
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) GetEnv() *string {
	return s.Env
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) GetId() *string {
	return s.Id
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) GetName() *string {
	return s.Name
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) SetDisplayName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) SetEnv(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) SetId(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) SetName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) Validate() error {
	return dara.Validate(s)
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo struct {
	// The display name.
	//
	// example:
	//
	// xx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The environment identifier. Valid values:
	//
	// - DEV
	//
	// - PROD.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 1123131
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name.
	//
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) GetEnv() *string {
	return s.Env
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) GetId() *int64 {
	return s.Id
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) GetName() *string {
	return s.Name
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) SetDisplayName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) SetEnv(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) SetId(v int64) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) SetName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) Validate() error {
	return dara.Validate(s)
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount struct {
	// The account ID.
	//
	// example:
	//
	// 1212131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The account name.
	//
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The authorized account type.
	//
	// example:
	//
	// PERSONAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) GetId() *string {
	return s.Id
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) GetName() *string {
	return s.Name
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) GetType() *string {
	return s.Type
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) SetId(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) SetName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) SetType(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount {
	s.Type = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) Validate() error {
	return dara.Validate(s)
}

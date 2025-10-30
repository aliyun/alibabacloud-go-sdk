// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListResourcePermissionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListResourcePermissionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListResourcePermissionsResponseBody
	GetMessage() *string
	SetPageResult(v *ListResourcePermissionsResponseBodyPageResult) *ListResourcePermissionsResponseBody
	GetPageResult() *ListResourcePermissionsResponseBodyPageResult
	SetRequestId(v string) *ListResourcePermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListResourcePermissionsResponseBody
	GetSuccess() *bool
}

type ListResourcePermissionsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListResourcePermissionsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListResourcePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListResourcePermissionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListResourcePermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListResourcePermissionsResponseBody) GetPageResult() *ListResourcePermissionsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListResourcePermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourcePermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListResourcePermissionsResponseBody) SetCode(v string) *ListResourcePermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListResourcePermissionsResponseBody) SetHttpStatusCode(v int32) *ListResourcePermissionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListResourcePermissionsResponseBody) SetMessage(v string) *ListResourcePermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListResourcePermissionsResponseBody) SetPageResult(v *ListResourcePermissionsResponseBodyPageResult) *ListResourcePermissionsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListResourcePermissionsResponseBody) SetRequestId(v string) *ListResourcePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcePermissionsResponseBody) SetSuccess(v bool) *ListResourcePermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListResourcePermissionsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcePermissionsResponseBodyPageResult struct {
	Data []*ListResourcePermissionsResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 121
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResult) GetData() []*ListResourcePermissionsResponseBodyPageResultData {
	return s.Data
}

func (s *ListResourcePermissionsResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListResourcePermissionsResponseBodyPageResult) SetData(v []*ListResourcePermissionsResponseBodyPageResultData) *ListResourcePermissionsResponseBodyPageResult {
	s.Data = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResult) SetTotalCount(v int64) *ListResourcePermissionsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResult) Validate() error {
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

type ListResourcePermissionsResponseBodyPageResultData struct {
	// example:
	//
	// selectTable
	AuthScope            *string                                                                  `json:"AuthScope,omitempty" xml:"AuthScope,omitempty"`
	Period               *ListResourcePermissionsResponseBodyPageResultDataPeriod                 `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	PermissionPeriodList []*ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList `json:"PermissionPeriodList,omitempty" xml:"PermissionPeriodList,omitempty" type:"Repeated"`
	// example:
	//
	// 12123111
	RecordId      *string                                                         `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	ResourceInfo  *ListResourcePermissionsResponseBodyPageResultDataResourceInfo  `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty" type:"Struct"`
	TargetAccount *ListResourcePermissionsResponseBodyPageResultDataTargetAccount `json:"TargetAccount,omitempty" xml:"TargetAccount,omitempty" type:"Struct"`
}

func (s ListResourcePermissionsResponseBodyPageResultData) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultData) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultData) GetAuthScope() *string {
	return s.AuthScope
}

func (s *ListResourcePermissionsResponseBodyPageResultData) GetPeriod() *ListResourcePermissionsResponseBodyPageResultDataPeriod {
	return s.Period
}

func (s *ListResourcePermissionsResponseBodyPageResultData) GetPermissionPeriodList() []*ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList {
	return s.PermissionPeriodList
}

func (s *ListResourcePermissionsResponseBodyPageResultData) GetRecordId() *string {
	return s.RecordId
}

func (s *ListResourcePermissionsResponseBodyPageResultData) GetResourceInfo() *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	return s.ResourceInfo
}

func (s *ListResourcePermissionsResponseBodyPageResultData) GetTargetAccount() *ListResourcePermissionsResponseBodyPageResultDataTargetAccount {
	return s.TargetAccount
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetAuthScope(v string) *ListResourcePermissionsResponseBodyPageResultData {
	s.AuthScope = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetPeriod(v *ListResourcePermissionsResponseBodyPageResultDataPeriod) *ListResourcePermissionsResponseBodyPageResultData {
	s.Period = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetPermissionPeriodList(v []*ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) *ListResourcePermissionsResponseBodyPageResultData {
	s.PermissionPeriodList = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetRecordId(v string) *ListResourcePermissionsResponseBodyPageResultData {
	s.RecordId = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetResourceInfo(v *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) *ListResourcePermissionsResponseBodyPageResultData {
	s.ResourceInfo = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetTargetAccount(v *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) *ListResourcePermissionsResponseBodyPageResultData {
	s.TargetAccount = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) Validate() error {
	if s.Period != nil {
		if err := s.Period.Validate(); err != nil {
			return err
		}
	}
	if s.PermissionPeriodList != nil {
		for _, item := range s.PermissionPeriodList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type ListResourcePermissionsResponseBodyPageResultDataPeriod struct {
	// example:
	//
	// 1712000000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataPeriod) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPeriod) GetType() *string {
	return s.Type
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPeriod) SetEndTime(v string) *ListResourcePermissionsResponseBodyPageResultDataPeriod {
	s.EndTime = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPeriod) SetType(v string) *ListResourcePermissionsResponseBodyPageResultDataPeriod {
	s.Type = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPeriod) Validate() error {
	return dara.Validate(s)
}

type ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList struct {
	Period *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// example:
	//
	// SELECT
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) GetPeriod() *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod {
	return s.Period
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) GetPermissionType() *string {
	return s.PermissionType
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) SetPeriod(v *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList {
	s.Period = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) SetPermissionType(v string) *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList {
	s.PermissionType = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) Validate() error {
	if s.Period != nil {
		if err := s.Period.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod struct {
	// example:
	//
	// 1712000000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) GetType() *string {
	return s.Type
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) SetEndTime(v string) *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod {
	s.EndTime = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) SetType(v string) *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod {
	s.Type = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) Validate() error {
	return dara.Validate(s)
}

type ListResourcePermissionsResponseBodyPageResultDataResourceInfo struct {
	BizUnitInfo *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo `json:"BizUnitInfo,omitempty" xml:"BizUnitInfo,omitempty" type:"Struct"`
	// example:
	//
	// tb1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// a.tb1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// tb1
	Name        *string                                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectInfo *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// example:
	//
	// PHYSICAL_TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) GetBizUnitInfo() *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo {
	return s.BizUnitInfo
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) GetEnv() *string {
	return s.Env
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) GetId() *string {
	return s.Id
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) GetName() *string {
	return s.Name
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) GetProjectInfo() *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo {
	return s.ProjectInfo
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) GetType() *string {
	return s.Type
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetBizUnitInfo(v *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.BizUnitInfo = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetDisplayName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetEnv(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetId(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetProjectInfo(v *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.ProjectInfo = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetType(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.Type = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) Validate() error {
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

type ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo struct {
	// example:
	//
	// xx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) GetEnv() *string {
	return s.Env
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) GetId() *string {
	return s.Id
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) GetName() *string {
	return s.Name
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) SetDisplayName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) SetEnv(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) SetId(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) SetName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) Validate() error {
	return dara.Validate(s)
}

type ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo struct {
	// example:
	//
	// xx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 1123131
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) GetEnv() *string {
	return s.Env
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) GetId() *int64 {
	return s.Id
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) GetName() *string {
	return s.Name
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) SetDisplayName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) SetEnv(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) SetId(v int64) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) SetName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) Validate() error {
	return dara.Validate(s)
}

type ListResourcePermissionsResponseBodyPageResultDataTargetAccount struct {
	// example:
	//
	// 1212131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// PERSONAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataTargetAccount) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataTargetAccount) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) GetId() *string {
	return s.Id
}

func (s *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) GetName() *string {
	return s.Name
}

func (s *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) GetType() *string {
	return s.Type
}

func (s *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) SetId(v string) *ListResourcePermissionsResponseBodyPageResultDataTargetAccount {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) SetName(v string) *ListResourcePermissionsResponseBodyPageResultDataTargetAccount {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) SetType(v string) *ListResourcePermissionsResponseBodyPageResultDataTargetAccount {
	s.Type = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) Validate() error {
	return dara.Validate(s)
}

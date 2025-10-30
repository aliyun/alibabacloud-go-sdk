// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceWithConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataSourceWithConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListDataSourceWithConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataSourceWithConfigResponseBody
	GetMessage() *string
	SetPageResult(v *ListDataSourceWithConfigResponseBodyPageResult) *ListDataSourceWithConfigResponseBody
	GetPageResult() *ListDataSourceWithConfigResponseBodyPageResult
	SetRequestId(v string) *ListDataSourceWithConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataSourceWithConfigResponseBody
	GetSuccess() *bool
}

type ListDataSourceWithConfigResponseBody struct {
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
	Message    *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListDataSourceWithConfigResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataSourceWithConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataSourceWithConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataSourceWithConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataSourceWithConfigResponseBody) GetPageResult() *ListDataSourceWithConfigResponseBodyPageResult {
	return s.PageResult
}

func (s *ListDataSourceWithConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceWithConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataSourceWithConfigResponseBody) SetCode(v string) *ListDataSourceWithConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) SetHttpStatusCode(v int32) *ListDataSourceWithConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) SetMessage(v string) *ListDataSourceWithConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) SetPageResult(v *ListDataSourceWithConfigResponseBodyPageResult) *ListDataSourceWithConfigResponseBody {
	s.PageResult = v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) SetRequestId(v string) *ListDataSourceWithConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) SetSuccess(v bool) *ListDataSourceWithConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataSourceWithConfigResponseBodyPageResult struct {
	DataSourceList []*ListDataSourceWithConfigResponseBodyPageResultDataSourceList `json:"DataSourceList,omitempty" xml:"DataSourceList,omitempty" type:"Repeated"`
	// example:
	//
	// 39
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataSourceWithConfigResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResult) GetDataSourceList() []*ListDataSourceWithConfigResponseBodyPageResultDataSourceList {
	return s.DataSourceList
}

func (s *ListDataSourceWithConfigResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDataSourceWithConfigResponseBodyPageResult) SetDataSourceList(v []*ListDataSourceWithConfigResponseBodyPageResultDataSourceList) *ListDataSourceWithConfigResponseBodyPageResult {
	s.DataSourceList = v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResult) SetTotalCount(v int64) *ListDataSourceWithConfigResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResult) Validate() error {
	if s.DataSourceList != nil {
		for _, item := range s.DataSourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourceWithConfigResponseBodyPageResultDataSourceList struct {
	// 开发环境数据源信息
	DevDataSourceInfo *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo `json:"DevDataSourceInfo,omitempty" xml:"DevDataSourceInfo,omitempty" type:"Struct"`
	// 生产环境数据源
	ProdDataSourceInfo *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo `json:"ProdDataSourceInfo,omitempty" xml:"ProdDataSourceInfo,omitempty" type:"Struct"`
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceList) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceList) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceList) GetDevDataSourceInfo() *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	return s.DevDataSourceInfo
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceList) GetProdDataSourceInfo() *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	return s.ProdDataSourceInfo
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceList) SetDevDataSourceInfo(v *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) *ListDataSourceWithConfigResponseBodyPageResultDataSourceList {
	s.DevDataSourceInfo = v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceList) SetProdDataSourceInfo(v *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) *ListDataSourceWithConfigResponseBodyPageResultDataSourceList {
	s.ProdDataSourceInfo = v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceList) Validate() error {
	if s.DevDataSourceInfo != nil {
		if err := s.DevDataSourceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ProdDataSourceInfo != nil {
		if err := s.ProdDataSourceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo struct {
	ConfigItemList []*ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// example:
	//
	// 1710209552704
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 212211111
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 12313123131
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1710209552704
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 212211111
	Owner     *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// ALL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetConfigItemList() []*ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList {
	return s.ConfigItemList
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetDescription() *string {
	return s.Description
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetEnv() *string {
	return s.Env
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetId() *int64 {
	return s.Id
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetName() *string {
	return s.Name
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetOwner() *string {
	return s.Owner
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetScope() *string {
	return s.Scope
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GetType() *string {
	return s.Type
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetConfigItemList(v []*ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.ConfigItemList = v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetCreateTime(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetCreator(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Creator = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetCreatorName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.CreatorName = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetDescription(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Description = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetEnv(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Env = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetId(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Id = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetModifyTime(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.ModifyTime = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Name = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetOwner(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Owner = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetOwnerName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.OwnerName = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetScope(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Scope = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetType(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Type = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) Validate() error {
	if s.ConfigItemList != nil {
		for _, item := range s.ConfigItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList struct {
	// example:
	//
	// param1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) GetKey() *string {
	return s.Key
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) GetValue() *string {
	return s.Value
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) SetKey(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList {
	s.Key = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) SetValue(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList {
	s.Value = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) Validate() error {
	return dara.Validate(s)
}

type ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo struct {
	ConfigItemList []*ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// example:
	//
	// 1708303959000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 212211111
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 300000028799
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1708303959000
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 212211111
	Owner     *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// ALL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetConfigItemList() []*ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList {
	return s.ConfigItemList
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetDescription() *string {
	return s.Description
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetEnv() *string {
	return s.Env
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetId() *int64 {
	return s.Id
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetName() *string {
	return s.Name
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetOwner() *string {
	return s.Owner
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetScope() *string {
	return s.Scope
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GetType() *string {
	return s.Type
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetConfigItemList(v []*ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.ConfigItemList = v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetCreateTime(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetCreator(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Creator = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetCreatorName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.CreatorName = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetDescription(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Description = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetEnv(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Env = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetId(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Id = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetModifyTime(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.ModifyTime = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Name = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetOwner(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Owner = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetOwnerName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.OwnerName = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetScope(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Scope = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetType(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Type = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) Validate() error {
	if s.ConfigItemList != nil {
		for _, item := range s.ConfigItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList struct {
	// example:
	//
	// param1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) GetKey() *string {
	return s.Key
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) GetValue() *string {
	return s.Value
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) SetKey(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList {
	s.Key = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) SetValue(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList {
	s.Value = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) Validate() error {
	return dara.Validate(s)
}

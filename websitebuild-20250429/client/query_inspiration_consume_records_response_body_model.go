// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationConsumeRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryInspirationConsumeRecordsResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *QueryInspirationConsumeRecordsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *QueryInspirationConsumeRecordsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *QueryInspirationConsumeRecordsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *QueryInspirationConsumeRecordsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *QueryInspirationConsumeRecordsResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *QueryInspirationConsumeRecordsResponseBodyModule) *QueryInspirationConsumeRecordsResponseBody
	GetModule() *QueryInspirationConsumeRecordsResponseBodyModule
	SetRequestId(v string) *QueryInspirationConsumeRecordsResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *QueryInspirationConsumeRecordsResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *QueryInspirationConsumeRecordsResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *QueryInspirationConsumeRecordsResponseBody
	GetSynchro() *bool
}

type QueryInspirationConsumeRecordsResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// ish-intelligence-store-platform-admin-web
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                           `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                     `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *QueryInspirationConsumeRecordsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s QueryInspirationConsumeRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationConsumeRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetModule() *QueryInspirationConsumeRecordsResponseBodyModule {
	return s.Module
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *QueryInspirationConsumeRecordsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetAccessDeniedDetail(v string) *QueryInspirationConsumeRecordsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetAllowRetry(v bool) *QueryInspirationConsumeRecordsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetAppName(v string) *QueryInspirationConsumeRecordsResponseBody {
	s.AppName = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetDynamicCode(v string) *QueryInspirationConsumeRecordsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetDynamicMessage(v string) *QueryInspirationConsumeRecordsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetErrorArgs(v []interface{}) *QueryInspirationConsumeRecordsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetModule(v *QueryInspirationConsumeRecordsResponseBodyModule) *QueryInspirationConsumeRecordsResponseBody {
	s.Module = v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetRequestId(v string) *QueryInspirationConsumeRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetRootErrorCode(v string) *QueryInspirationConsumeRecordsResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetRootErrorMsg(v string) *QueryInspirationConsumeRecordsResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) SetSynchro(v bool) *QueryInspirationConsumeRecordsResponseBody {
	s.Synchro = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInspirationConsumeRecordsResponseBodyModule struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                                                  `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryInspirationConsumeRecordsResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Next           *QueryInspirationConsumeRecordsResponseBodyModuleNext   `json:"Next,omitempty" xml:"Next,omitempty" type:"Struct"`
	// example:
	//
	// False
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// False
	PrePage     *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	ResultLimit *bool `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryInspirationConsumeRecordsResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationConsumeRecordsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) GetData() []*QueryInspirationConsumeRecordsResponseBodyModuleData {
	return s.Data
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) GetNext() *QueryInspirationConsumeRecordsResponseBodyModuleNext {
	return s.Next
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) GetResultLimit() *bool {
	return s.ResultLimit
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) SetCurrentPageNum(v int32) *QueryInspirationConsumeRecordsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) SetData(v []*QueryInspirationConsumeRecordsResponseBodyModuleData) *QueryInspirationConsumeRecordsResponseBodyModule {
	s.Data = v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) SetNext(v *QueryInspirationConsumeRecordsResponseBodyModuleNext) *QueryInspirationConsumeRecordsResponseBodyModule {
	s.Next = v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) SetNextPage(v bool) *QueryInspirationConsumeRecordsResponseBodyModule {
	s.NextPage = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) SetPageSize(v int32) *QueryInspirationConsumeRecordsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) SetPrePage(v bool) *QueryInspirationConsumeRecordsResponseBodyModule {
	s.PrePage = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) SetResultLimit(v bool) *QueryInspirationConsumeRecordsResponseBodyModule {
	s.ResultLimit = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) SetTotalItemNum(v int32) *QueryInspirationConsumeRecordsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) SetTotalPageNum(v int32) *QueryInspirationConsumeRecordsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModule) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Next != nil {
		if err := s.Next.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInspirationConsumeRecordsResponseBodyModuleData struct {
	// example:
	//
	// 120
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 2026-06-01 12:00:00
	ConsumeTime *string `json:"ConsumeTime,omitempty" xml:"ConsumeTime,omitempty"`
	// example:
	//
	// {\\"MD5\\":\\"296f6c01e7fea2697ffe1cf41082b774\\",\\"driver\\":\\"vhd\\",\\"flag\\":\\"12845825\\",\\"imds_support\\":\\"v1\\",\\"io_optimized\\":true,\\"nvme_supported\\":true,\\"uefi_preferred\\":false}
	MetaData  *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s QueryInspirationConsumeRecordsResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationConsumeRecordsResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleData) GetAmount() *int64 {
	return s.Amount
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleData) GetConsumeTime() *string {
	return s.ConsumeTime
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleData) GetMetaData() *string {
	return s.MetaData
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleData) GetSceneName() *string {
	return s.SceneName
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleData) SetAmount(v int64) *QueryInspirationConsumeRecordsResponseBodyModuleData {
	s.Amount = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleData) SetConsumeTime(v string) *QueryInspirationConsumeRecordsResponseBodyModuleData {
	s.ConsumeTime = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleData) SetMetaData(v string) *QueryInspirationConsumeRecordsResponseBodyModuleData {
	s.MetaData = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleData) SetSceneName(v string) *QueryInspirationConsumeRecordsResponseBodyModuleData {
	s.SceneName = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

type QueryInspirationConsumeRecordsResponseBodyModuleNext struct {
	// example:
	//
	// 2
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 2026-06-01 12:00:00
	ConsumeTime *string `json:"ConsumeTime,omitempty" xml:"ConsumeTime,omitempty"`
	// example:
	//
	// {\\"MD5\\":\\"1042e65a2b7cdd3059b6a873ee1a3260\\",\\"driver\\":\\"vhd\\",\\"flag\\":\\"12845825\\",\\"imds_support\\":\\"v1\\",\\"io_optimized\\":true,\\"nvme_supported\\":true,\\"uefi_preferred\\":false}
	MetaData  *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s QueryInspirationConsumeRecordsResponseBodyModuleNext) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationConsumeRecordsResponseBodyModuleNext) GoString() string {
	return s.String()
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleNext) GetAmount() *int64 {
	return s.Amount
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleNext) GetConsumeTime() *string {
	return s.ConsumeTime
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleNext) GetMetaData() *string {
	return s.MetaData
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleNext) GetSceneName() *string {
	return s.SceneName
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleNext) SetAmount(v int64) *QueryInspirationConsumeRecordsResponseBodyModuleNext {
	s.Amount = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleNext) SetConsumeTime(v string) *QueryInspirationConsumeRecordsResponseBodyModuleNext {
	s.ConsumeTime = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleNext) SetMetaData(v string) *QueryInspirationConsumeRecordsResponseBodyModuleNext {
	s.MetaData = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleNext) SetSceneName(v string) *QueryInspirationConsumeRecordsResponseBodyModuleNext {
	s.SceneName = &v
	return s
}

func (s *QueryInspirationConsumeRecordsResponseBodyModuleNext) Validate() error {
	return dara.Validate(s)
}

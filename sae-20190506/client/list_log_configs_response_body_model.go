// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLogConfigsResponseBody
	GetCode() *string
	SetData(v *ListLogConfigsResponseBodyData) *ListLogConfigsResponseBody
	GetData() *ListLogConfigsResponseBodyData
	SetErrorCode(v string) *ListLogConfigsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListLogConfigsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLogConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLogConfigsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListLogConfigsResponseBody
	GetTraceId() *string
}

type ListLogConfigsResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The request is successful.
	//
	// - **3xx**: The request is redirected.
	//
	// - **4xx**: A request error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the file logs.
	Data *ListLogConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is not returned if the request is successful.
	//
	// - This parameter is returned if the request fails. For more information, see the **Error codes*	- list in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the list of application logs was obtained. Valid values:
	//
	// - **true**: The list was obtained.
	//
	// - **false**: The list failed to be obtained.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID of the request. You can use the trace ID to query the details of the request.
	//
	// example:
	//
	// ac1d5e2c15671581252413581d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListLogConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLogConfigsResponseBody) GetData() *ListLogConfigsResponseBodyData {
	return s.Data
}

func (s *ListLogConfigsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListLogConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLogConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLogConfigsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListLogConfigsResponseBody) SetCode(v string) *ListLogConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLogConfigsResponseBody) SetData(v *ListLogConfigsResponseBodyData) *ListLogConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListLogConfigsResponseBody) SetErrorCode(v string) *ListLogConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListLogConfigsResponseBody) SetMessage(v string) *ListLogConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLogConfigsResponseBody) SetRequestId(v string) *ListLogConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogConfigsResponseBody) SetSuccess(v bool) *ListLogConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListLogConfigsResponseBody) SetTraceId(v string) *ListLogConfigsResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListLogConfigsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLogConfigsResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The log configurations.
	LogConfigs []*ListLogConfigsResponseBodyDataLogConfigs `json:"LogConfigs,omitempty" xml:"LogConfigs,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListLogConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLogConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLogConfigsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListLogConfigsResponseBodyData) GetLogConfigs() []*ListLogConfigsResponseBodyDataLogConfigs {
	return s.LogConfigs
}

func (s *ListLogConfigsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLogConfigsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListLogConfigsResponseBodyData) SetCurrentPage(v int32) *ListLogConfigsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListLogConfigsResponseBodyData) SetLogConfigs(v []*ListLogConfigsResponseBodyDataLogConfigs) *ListLogConfigsResponseBodyData {
	s.LogConfigs = v
	return s
}

func (s *ListLogConfigsResponseBodyData) SetPageSize(v int32) *ListLogConfigsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListLogConfigsResponseBodyData) SetTotalSize(v int32) *ListLogConfigsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListLogConfigsResponseBodyData) Validate() error {
	if s.LogConfigs != nil {
		for _, item := range s.LogConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLogConfigsResponseBodyDataLogConfigs struct {
	// The name of the Simple Log Service configuration.
	//
	// example:
	//
	// sae-1f240907a6faf58c653f09e81b7e****
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// The time when the log configuration was created.
	//
	// example:
	//
	// 2019-08-29 17:18:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The path of the log file. This is the log source.
	//
	// example:
	//
	// /root/logs/hsf/hsf.log
	LogDir *string `json:"LogDir,omitempty" xml:"LogDir,omitempty"`
	// The log type. Only **file_log*	- is supported.
	//
	// example:
	//
	// file_log
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the Logstore in Simple Log Service.
	//
	// example:
	//
	// sae-1f240907a6faf58c653f09e81b7e****
	SlsLogStore *string `json:"SlsLogStore,omitempty" xml:"SlsLogStore,omitempty"`
	// The ID of the Simple Log Service project.
	//
	// example:
	//
	// sae-56f77b65-788d-442a-9885-7f20d91f****
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// The storage class for Simple Log Service.
	//
	// example:
	//
	// sls
	StoreType *string `json:"StoreType,omitempty" xml:"StoreType,omitempty"`
}

func (s ListLogConfigsResponseBodyDataLogConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListLogConfigsResponseBodyDataLogConfigs) GoString() string {
	return s.String()
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) GetConfigName() *string {
	return s.ConfigName
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) GetLogDir() *string {
	return s.LogDir
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) GetLogType() *string {
	return s.LogType
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) GetSlsLogStore() *string {
	return s.SlsLogStore
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) GetSlsProject() *string {
	return s.SlsProject
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) GetStoreType() *string {
	return s.StoreType
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetConfigName(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.ConfigName = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetCreateTime(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.CreateTime = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetLogDir(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.LogDir = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetLogType(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.LogType = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetRegionId(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.RegionId = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetSlsLogStore(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.SlsLogStore = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetSlsProject(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.SlsProject = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetStoreType(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.StoreType = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) Validate() error {
	return dara.Validate(s)
}

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
	// Indicates whether the logging configurations of an application were obtained. Valid values:
	//
	// 	- **true**: indicates that the configurations were obtained.
	//
	// 	- **false**: indicates that the configurations could not be obtained.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The logging configurations.
	Data *ListLogConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the request was successful.
	//
	// 	- **3xx**: indicates that the request was redirected.
	//
	// 	- **4xx**: indicates that the request was invalid.
	//
	// 	- **5xx**: indicates that a server error occurred.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the trace. It can be used to query the details of a request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned message.
	//
	// 	- **success*	- is returned when the request succeeds.
	//
	// 	- An error code is returned when the request fails.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The logging configurations.
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
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The details of the logging configuration.
	LogConfigs []*ListLogConfigsResponseBodyDataLogConfigs `json:"LogConfigs,omitempty" xml:"LogConfigs,omitempty" type:"Repeated"`
	// The error code.
	//
	// 	- The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// 	- The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The number of entries returned on each page.
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
	// The path of logs.
	//
	// example:
	//
	// sae-1f240907a6faf58c653f09e81b7e****
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// The storage type of logs.
	//
	// example:
	//
	// 2019-08-29 17:18:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The path of the log file (log source).
	//
	// example:
	//
	// /root/logs/hsf/hsf.log
	LogDir *string `json:"LogDir,omitempty" xml:"LogDir,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// file_log
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The number of the returned page.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the configuration was created.
	//
	// example:
	//
	// sae-1f240907a6faf58c653f09e81b7e****
	SlsLogStore *string `json:"SlsLogStore,omitempty" xml:"SlsLogStore,omitempty"`
	// The type of the log. Set this value to **file_log**.
	//
	// example:
	//
	// sae-56f77b65-788d-442a-9885-7f20d91f****
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// The ID of the Log Service project.
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

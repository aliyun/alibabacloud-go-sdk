// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataSourceLogsResponseBodyData) *ListDataSourceLogsResponseBody
	GetData() *ListDataSourceLogsResponseBodyData
	SetRequestId(v string) *ListDataSourceLogsResponseBody
	GetRequestId() *string
}

type ListDataSourceLogsResponseBody struct {
	// The data returned.
	Data *ListDataSourceLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourceLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsResponseBody) GetData() *ListDataSourceLogsResponseBodyData {
	return s.Data
}

func (s *ListDataSourceLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceLogsResponseBody) SetData(v *ListDataSourceLogsResponseBodyData) *ListDataSourceLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListDataSourceLogsResponseBody) SetRequestId(v string) *ListDataSourceLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceLogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataSourceLogsResponseBodyData struct {
	// The ID of the cloud account.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code that is used for multi-cloud environments. Valid values:
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The ID of the data source. The value is obtained after the threat analysis feature calculates the MD5 hash value of a parameter.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The logs of the data source.
	DataSourceInstanceLogs []*ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs `json:"DataSourceInstanceLogs,omitempty" xml:"DataSourceInstanceLogs,omitempty" type:"Repeated"`
	// The name of the data source.
	//
	// example:
	//
	// waf kafka
	DataSourceInstanceName *string `json:"DataSourceInstanceName,omitempty" xml:"DataSourceInstanceName,omitempty"`
	// The remarks of the data source.
	//
	// example:
	//
	// waf kafka
	DataSourceInstanceRemark *string `json:"DataSourceInstanceRemark,omitempty" xml:"DataSourceInstanceRemark,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListDataSourceLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *ListDataSourceLogsResponseBodyData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListDataSourceLogsResponseBodyData) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *ListDataSourceLogsResponseBodyData) GetDataSourceInstanceLogs() []*ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	return s.DataSourceInstanceLogs
}

func (s *ListDataSourceLogsResponseBodyData) GetDataSourceInstanceName() *string {
	return s.DataSourceInstanceName
}

func (s *ListDataSourceLogsResponseBodyData) GetDataSourceInstanceRemark() *string {
	return s.DataSourceInstanceRemark
}

func (s *ListDataSourceLogsResponseBodyData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *ListDataSourceLogsResponseBodyData) SetAccountId(v string) *ListDataSourceLogsResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetCloudCode(v string) *ListDataSourceLogsResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetDataSourceInstanceId(v string) *ListDataSourceLogsResponseBodyData {
	s.DataSourceInstanceId = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetDataSourceInstanceLogs(v []*ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) *ListDataSourceLogsResponseBodyData {
	s.DataSourceInstanceLogs = v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetDataSourceInstanceName(v string) *ListDataSourceLogsResponseBodyData {
	s.DataSourceInstanceName = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetDataSourceInstanceRemark(v string) *ListDataSourceLogsResponseBodyData {
	s.DataSourceInstanceRemark = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) SetSubUserId(v int64) *ListDataSourceLogsResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyData) Validate() error {
	if s.DataSourceInstanceLogs != nil {
		for _, item := range s.DataSourceInstanceLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs struct {
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_waf_xxxxx
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The ID of the log. The value is obtained after the threat analysis feature calculates the MD5 hash value of a parameter.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	LogInstanceId *string `json:"LogInstanceId,omitempty" xml:"LogInstanceId,omitempty"`
	// The display code of the log.
	//
	// example:
	//
	// ${siem.prod.cloud_siem_waf_xxxxx}
	LogMdsCode *string `json:"LogMdsCode,omitempty" xml:"LogMdsCode,omitempty"`
	// The parameters of the log.
	LogParams []*ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams `json:"LogParams,omitempty" xml:"LogParams,omitempty" type:"Repeated"`
	// Indicates whether the task for which logs are collected is enabled. Valid values:
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) GetLogCode() *string {
	return s.LogCode
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) GetLogInstanceId() *string {
	return s.LogInstanceId
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) GetLogMdsCode() *string {
	return s.LogMdsCode
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) GetLogParams() []*ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams {
	return s.LogParams
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) SetLogCode(v string) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	s.LogCode = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) SetLogInstanceId(v string) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	s.LogInstanceId = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) SetLogMdsCode(v string) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	s.LogMdsCode = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) SetLogParams(v []*ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	s.LogParams = v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) SetTaskStatus(v int32) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs {
	s.TaskStatus = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogs) Validate() error {
	if s.LogParams != nil {
		for _, item := range s.LogParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams struct {
	// The parameter code of the log.
	//
	// example:
	//
	// region_code
	ParaCode *string `json:"ParaCode,omitempty" xml:"ParaCode,omitempty"`
	// The parameter value of the log.
	//
	// example:
	//
	// ap-guangzhou
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
}

func (s ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) GetParaCode() *string {
	return s.ParaCode
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) GetParaValue() *string {
	return s.ParaValue
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) SetParaCode(v string) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams {
	s.ParaCode = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) SetParaValue(v string) *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams {
	s.ParaValue = &v
	return s
}

func (s *ListDataSourceLogsResponseBodyDataDataSourceInstanceLogsLogParams) Validate() error {
	return dara.Validate(s)
}

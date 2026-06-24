// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataLimitId(v int64) *CreateScanTaskRequest
	GetDataLimitId() *int64
	SetFeatureType(v int32) *CreateScanTaskRequest
	GetFeatureType() *int32
	SetIntervalDay(v int32) *CreateScanTaskRequest
	GetIntervalDay() *int32
	SetLang(v string) *CreateScanTaskRequest
	GetLang() *string
	SetOssScanPath(v string) *CreateScanTaskRequest
	GetOssScanPath() *string
	SetResourceType(v int64) *CreateScanTaskRequest
	GetResourceType() *int64
	SetRunHour(v int32) *CreateScanTaskRequest
	GetRunHour() *int32
	SetRunMinute(v int32) *CreateScanTaskRequest
	GetRunMinute() *int32
	SetScanRange(v int32) *CreateScanTaskRequest
	GetScanRange() *int32
	SetScanRangeContent(v string) *CreateScanTaskRequest
	GetScanRangeContent() *string
	SetSourceIp(v string) *CreateScanTaskRequest
	GetSourceIp() *string
	SetTaskName(v string) *CreateScanTaskRequest
	GetTaskName() *string
	SetTaskUserName(v string) *CreateScanTaskRequest
	GetTaskUserName() *string
}

type CreateScanTaskRequest struct {
	// The unique ID of the data asset. The asset can be an instance, a database, or a bucket. Call the [DescribeDataLimits](~~DescribeDataLimits~~) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataLimitId *int64 `json:"DataLimitId,omitempty" xml:"DataLimitId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The interval in days between two consecutive custom scan tasks. The value must be between 1 and 2147483648.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	IntervalDay *int32 `json:"IntervalDay,omitempty" xml:"IntervalDay,omitempty"`
	// The language of the request and response.
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The scan scope for OSS assets. You can specify a prefix, a suffix, or a regular expression to match objects.
	//
	// example:
	//
	// /test/test
	OssScanPath *string `json:"OssScanPath,omitempty" xml:"OssScanPath,omitempty"`
	// The type of resource to query. Valid values:
	//
	// - **1**: MaxCompute.
	//
	// - **2**: OSS.
	//
	// - **3**: AnalyticDB.
	//
	// - **4**: Tablestore.
	//
	// - **5**: RDS.
	//
	// - **6**: a self-managed database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The hour at which the next scan task runs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	RunHour *int32 `json:"RunHour,omitempty" xml:"RunHour,omitempty"`
	// The minute at which the next scan task runs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	RunMinute *int32 `json:"RunMinute,omitempty" xml:"RunMinute,omitempty"`
	// The matching rule for the scan scope of the custom scan task. This parameter takes effect only when you configure the **ScanRangeContent*	- parameter. Valid values:
	//
	// - **0**: full match.
	//
	// - **1**: prefix match.
	//
	// - **2**: suffix match.
	//
	// - **3**: regular expression match.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ScanRange *int32 `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	// The content to match for the scan of structured data assets. This parameter is used with the ScanRange parameter.
	//
	// > If you set ScanRange to 0, the scan matches the exact value of this parameter. If you set ScanRange to 1, the scan matches items that have the prefix specified by this parameter. For example, if you set this parameter to \\`test/abc\\`, file paths that start with \\`test/abc\\` are matched. If you set ScanRange to 2, the scan matches items that have the suffix specified by this parameter. If you set ScanRange to 3, the scan matches items that match the regular expression specified by this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// datamask/
	ScanRangeContent *string `json:"ScanRangeContent,omitempty" xml:"ScanRangeContent,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 106.11.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The name of the scan task.
	//
	// This parameter is required.
	//
	// example:
	//
	// scan-test-sample****
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The account that creates the scan task.
	//
	// example:
	//
	// demo
	TaskUserName *string `json:"TaskUserName,omitempty" xml:"TaskUserName,omitempty"`
}

func (s CreateScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScanTaskRequest) GetDataLimitId() *int64 {
	return s.DataLimitId
}

func (s *CreateScanTaskRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *CreateScanTaskRequest) GetIntervalDay() *int32 {
	return s.IntervalDay
}

func (s *CreateScanTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateScanTaskRequest) GetOssScanPath() *string {
	return s.OssScanPath
}

func (s *CreateScanTaskRequest) GetResourceType() *int64 {
	return s.ResourceType
}

func (s *CreateScanTaskRequest) GetRunHour() *int32 {
	return s.RunHour
}

func (s *CreateScanTaskRequest) GetRunMinute() *int32 {
	return s.RunMinute
}

func (s *CreateScanTaskRequest) GetScanRange() *int32 {
	return s.ScanRange
}

func (s *CreateScanTaskRequest) GetScanRangeContent() *string {
	return s.ScanRangeContent
}

func (s *CreateScanTaskRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateScanTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateScanTaskRequest) GetTaskUserName() *string {
	return s.TaskUserName
}

func (s *CreateScanTaskRequest) SetDataLimitId(v int64) *CreateScanTaskRequest {
	s.DataLimitId = &v
	return s
}

func (s *CreateScanTaskRequest) SetFeatureType(v int32) *CreateScanTaskRequest {
	s.FeatureType = &v
	return s
}

func (s *CreateScanTaskRequest) SetIntervalDay(v int32) *CreateScanTaskRequest {
	s.IntervalDay = &v
	return s
}

func (s *CreateScanTaskRequest) SetLang(v string) *CreateScanTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateScanTaskRequest) SetOssScanPath(v string) *CreateScanTaskRequest {
	s.OssScanPath = &v
	return s
}

func (s *CreateScanTaskRequest) SetResourceType(v int64) *CreateScanTaskRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateScanTaskRequest) SetRunHour(v int32) *CreateScanTaskRequest {
	s.RunHour = &v
	return s
}

func (s *CreateScanTaskRequest) SetRunMinute(v int32) *CreateScanTaskRequest {
	s.RunMinute = &v
	return s
}

func (s *CreateScanTaskRequest) SetScanRange(v int32) *CreateScanTaskRequest {
	s.ScanRange = &v
	return s
}

func (s *CreateScanTaskRequest) SetScanRangeContent(v string) *CreateScanTaskRequest {
	s.ScanRangeContent = &v
	return s
}

func (s *CreateScanTaskRequest) SetSourceIp(v string) *CreateScanTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateScanTaskRequest) SetTaskName(v string) *CreateScanTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateScanTaskRequest) SetTaskUserName(v string) *CreateScanTaskRequest {
	s.TaskUserName = &v
	return s
}

func (s *CreateScanTaskRequest) Validate() error {
	return dara.Validate(s)
}

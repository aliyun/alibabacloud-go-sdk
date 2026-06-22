// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevels(v []*int32) *ListFileProtectEventRequest
	GetAlertLevels() []*int32
	SetCurrentPage(v int32) *ListFileProtectEventRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListFileProtectEventRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListFileProtectEventRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListFileProtectEventRequest
	GetInstanceName() *string
	SetInternetIp(v string) *ListFileProtectEventRequest
	GetInternetIp() *string
	SetIntranetIp(v string) *ListFileProtectEventRequest
	GetIntranetIp() *string
	SetOperation(v string) *ListFileProtectEventRequest
	GetOperation() *string
	SetPageSize(v string) *ListFileProtectEventRequest
	GetPageSize() *string
	SetRuleName(v string) *ListFileProtectEventRequest
	GetRuleName() *string
	SetStartTime(v int64) *ListFileProtectEventRequest
	GetStartTime() *int64
	SetStatus(v string) *ListFileProtectEventRequest
	GetStatus() *string
	SetUuid(v string) *ListFileProtectEventRequest
	GetUuid() *string
}

type ListFileProtectEventRequest struct {
	// The list of alert notification levels.
	AlertLevels []*int32 `json:"AlertLevels,omitempty" xml:"AlertLevels,omitempty" type:"Repeated"`
	// The page number of the current page in a paging query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The timestamp of the end time.
	//
	// example:
	//
	// 1683257937775
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the asset instance.
	//
	// example:
	//
	// i-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server to query.
	//
	// example:
	//
	// ca_cpm_****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the asset to query.
	//
	// example:
	//
	// 120.27.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the asset to query.
	//
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The type of operation performed on the file. Valid values:
	//
	// - **DELETE**: deletes the file.
	//
	// - **WRITE**: writes to the file.
	//
	// - **READ**: reads the file.
	//
	// - **RENAME**: renames the file.
	//
	// - **CHOWN**: changes the file owner and associated group.
	//
	// example:
	//
	// READ
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The maximum number of entries to return on each page in a paging query.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule name.
	//
	// example:
	//
	// test-rule-1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The timestamp of the start time.
	//
	// example:
	//
	// 1656038740435
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event status. Valid values:
	//
	// - 0: Unhandled.
	//
	// - 1: Manually handled.
	//
	// - 2: Whitelisted.
	//
	// - 3: Ignored.
	//
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server to query.
	//
	// >You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// 4fe8e1cd-3c37-4851-b9de-124da32c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListFileProtectEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectEventRequest) GoString() string {
	return s.String()
}

func (s *ListFileProtectEventRequest) GetAlertLevels() []*int32 {
	return s.AlertLevels
}

func (s *ListFileProtectEventRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFileProtectEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListFileProtectEventRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFileProtectEventRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListFileProtectEventRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListFileProtectEventRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListFileProtectEventRequest) GetOperation() *string {
	return s.Operation
}

func (s *ListFileProtectEventRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListFileProtectEventRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListFileProtectEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListFileProtectEventRequest) GetStatus() *string {
	return s.Status
}

func (s *ListFileProtectEventRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListFileProtectEventRequest) SetAlertLevels(v []*int32) *ListFileProtectEventRequest {
	s.AlertLevels = v
	return s
}

func (s *ListFileProtectEventRequest) SetCurrentPage(v int32) *ListFileProtectEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectEventRequest) SetEndTime(v int64) *ListFileProtectEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListFileProtectEventRequest) SetInstanceId(v string) *ListFileProtectEventRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFileProtectEventRequest) SetInstanceName(v string) *ListFileProtectEventRequest {
	s.InstanceName = &v
	return s
}

func (s *ListFileProtectEventRequest) SetInternetIp(v string) *ListFileProtectEventRequest {
	s.InternetIp = &v
	return s
}

func (s *ListFileProtectEventRequest) SetIntranetIp(v string) *ListFileProtectEventRequest {
	s.IntranetIp = &v
	return s
}

func (s *ListFileProtectEventRequest) SetOperation(v string) *ListFileProtectEventRequest {
	s.Operation = &v
	return s
}

func (s *ListFileProtectEventRequest) SetPageSize(v string) *ListFileProtectEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectEventRequest) SetRuleName(v string) *ListFileProtectEventRequest {
	s.RuleName = &v
	return s
}

func (s *ListFileProtectEventRequest) SetStartTime(v int64) *ListFileProtectEventRequest {
	s.StartTime = &v
	return s
}

func (s *ListFileProtectEventRequest) SetStatus(v string) *ListFileProtectEventRequest {
	s.Status = &v
	return s
}

func (s *ListFileProtectEventRequest) SetUuid(v string) *ListFileProtectEventRequest {
	s.Uuid = &v
	return s
}

func (s *ListFileProtectEventRequest) Validate() error {
	return dara.Validate(s)
}

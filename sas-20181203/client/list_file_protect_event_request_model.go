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
	// The severities of alerts.
	AlertLevels []*int32 `json:"AlertLevels,omitempty" xml:"AlertLevels,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end timestamp of the query.
	//
	// example:
	//
	// 1683195595204
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// i-bp1fu4aqltf1huhc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// ca_cpm_****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 120.27.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// Type of operation on a file. eg:
	//
	// - **DELETE**: delete the file.
	//
	// - **WRITE**: write the file.
	//
	// - **READ**: read the file.
	//
	// - **RENAME**: rename the file.
	//
	// - **CHOWN**: set the file owner and file association group operations.
	//
	// example:
	//
	// READ
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test-rule-1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The start timestamp of the query.
	//
	// example:
	//
	// 1683080489594
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- 0: unhandled
	//
	// 	- 1: handled
	//
	// 	- 2: added to the whitelist
	//
	// 	- 3: ignored
	//
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUID of the server.
	//
	// example:
	//
	// inet-ecs-4e876cb0-09f7-43b8-82ef-4bc7a937***
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

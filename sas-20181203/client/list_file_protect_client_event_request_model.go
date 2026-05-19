// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectClientEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevels(v []*int32) *ListFileProtectClientEventRequest
	GetAlertLevels() []*int32
	SetCurrentPage(v int32) *ListFileProtectClientEventRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListFileProtectClientEventRequest
	GetEndTime() *int64
	SetFilePath(v string) *ListFileProtectClientEventRequest
	GetFilePath() *string
	SetInstanceId(v string) *ListFileProtectClientEventRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListFileProtectClientEventRequest
	GetInstanceName() *string
	SetInternetIp(v string) *ListFileProtectClientEventRequest
	GetInternetIp() *string
	SetIntranetIp(v string) *ListFileProtectClientEventRequest
	GetIntranetIp() *string
	SetOperation(v string) *ListFileProtectClientEventRequest
	GetOperation() *string
	SetPageSize(v int32) *ListFileProtectClientEventRequest
	GetPageSize() *int32
	SetProcPath(v string) *ListFileProtectClientEventRequest
	GetProcPath() *string
	SetRuleName(v string) *ListFileProtectClientEventRequest
	GetRuleName() *string
	SetStartTime(v int64) *ListFileProtectClientEventRequest
	GetStartTime() *int64
	SetStatus(v string) *ListFileProtectClientEventRequest
	GetStatus() *string
	SetUuid(v string) *ListFileProtectClientEventRequest
	GetUuid() *string
}

type ListFileProtectClientEventRequest struct {
	AlertLevels []*int32 `json:"AlertLevels,omitempty" xml:"AlertLevels,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1650470399999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// /etc/pam****
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// i-bp1fu4aqltf1huhc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ca_cpm_****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 120.27.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// example:
	//
	// READ
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// c:/windows/system32/i****
	ProcPath *string `json:"ProcPath,omitempty" xml:"ProcPath,omitempty"`
	// example:
	//
	// text-001
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1650470399999
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 4fe8e1cd-3c37-4851-b9de-124da32c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListFileProtectClientEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientEventRequest) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientEventRequest) GetAlertLevels() []*int32 {
	return s.AlertLevels
}

func (s *ListFileProtectClientEventRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFileProtectClientEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListFileProtectClientEventRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ListFileProtectClientEventRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFileProtectClientEventRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListFileProtectClientEventRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListFileProtectClientEventRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListFileProtectClientEventRequest) GetOperation() *string {
	return s.Operation
}

func (s *ListFileProtectClientEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileProtectClientEventRequest) GetProcPath() *string {
	return s.ProcPath
}

func (s *ListFileProtectClientEventRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListFileProtectClientEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListFileProtectClientEventRequest) GetStatus() *string {
	return s.Status
}

func (s *ListFileProtectClientEventRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListFileProtectClientEventRequest) SetAlertLevels(v []*int32) *ListFileProtectClientEventRequest {
	s.AlertLevels = v
	return s
}

func (s *ListFileProtectClientEventRequest) SetCurrentPage(v int32) *ListFileProtectClientEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetEndTime(v int64) *ListFileProtectClientEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetFilePath(v string) *ListFileProtectClientEventRequest {
	s.FilePath = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetInstanceId(v string) *ListFileProtectClientEventRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetInstanceName(v string) *ListFileProtectClientEventRequest {
	s.InstanceName = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetInternetIp(v string) *ListFileProtectClientEventRequest {
	s.InternetIp = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetIntranetIp(v string) *ListFileProtectClientEventRequest {
	s.IntranetIp = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetOperation(v string) *ListFileProtectClientEventRequest {
	s.Operation = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetPageSize(v int32) *ListFileProtectClientEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetProcPath(v string) *ListFileProtectClientEventRequest {
	s.ProcPath = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetRuleName(v string) *ListFileProtectClientEventRequest {
	s.RuleName = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetStartTime(v int64) *ListFileProtectClientEventRequest {
	s.StartTime = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetStatus(v string) *ListFileProtectClientEventRequest {
	s.Status = &v
	return s
}

func (s *ListFileProtectClientEventRequest) SetUuid(v string) *ListFileProtectClientEventRequest {
	s.Uuid = &v
	return s
}

func (s *ListFileProtectClientEventRequest) Validate() error {
	return dara.Validate(s)
}

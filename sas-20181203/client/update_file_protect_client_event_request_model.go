// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectClientEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevels(v []*int32) *UpdateFileProtectClientEventRequest
	GetAlertLevels() []*int32
	SetEndTime(v int64) *UpdateFileProtectClientEventRequest
	GetEndTime() *int64
	SetExcludeIdList(v []*int64) *UpdateFileProtectClientEventRequest
	GetExcludeIdList() []*int64
	SetFilePath(v string) *UpdateFileProtectClientEventRequest
	GetFilePath() *string
	SetIdList(v []*int64) *UpdateFileProtectClientEventRequest
	GetIdList() []*int64
	SetInstanceId(v string) *UpdateFileProtectClientEventRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateFileProtectClientEventRequest
	GetInstanceName() *string
	SetInternetIp(v string) *UpdateFileProtectClientEventRequest
	GetInternetIp() *string
	SetIntranetIp(v string) *UpdateFileProtectClientEventRequest
	GetIntranetIp() *string
	SetNewStatus(v int32) *UpdateFileProtectClientEventRequest
	GetNewStatus() *int32
	SetOperation(v string) *UpdateFileProtectClientEventRequest
	GetOperation() *string
	SetProcPath(v string) *UpdateFileProtectClientEventRequest
	GetProcPath() *string
	SetRemark(v []*string) *UpdateFileProtectClientEventRequest
	GetRemark() []*string
	SetRuleName(v string) *UpdateFileProtectClientEventRequest
	GetRuleName() *string
	SetSelectAll(v bool) *UpdateFileProtectClientEventRequest
	GetSelectAll() *bool
	SetStartTime(v int64) *UpdateFileProtectClientEventRequest
	GetStartTime() *int64
	SetStatus(v string) *UpdateFileProtectClientEventRequest
	GetStatus() *string
	SetUuid(v string) *UpdateFileProtectClientEventRequest
	GetUuid() *string
}

type UpdateFileProtectClientEventRequest struct {
	AlertLevels []*int32 `json:"AlertLevels,omitempty" xml:"AlertLevels,omitempty" type:"Repeated"`
	// example:
	//
	// 1650470399999
	EndTime       *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExcludeIdList []*int64 `json:"ExcludeIdList,omitempty" xml:"ExcludeIdList,omitempty" type:"Repeated"`
	// example:
	//
	// /etc/pam****
	FilePath *string  `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	IdList   []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// example:
	//
	// i-bp1fu4aqltf1huhc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 120.27.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	NewStatus *int32 `json:"NewStatus,omitempty" xml:"NewStatus,omitempty"`
	// example:
	//
	// READ
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// /root/1111/****
	ProcPath *string   `json:"ProcPath,omitempty" xml:"ProcPath,omitempty"`
	Remark   []*string `json:"Remark,omitempty" xml:"Remark,omitempty" type:"Repeated"`
	// example:
	//
	// tetsRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	SelectAll *bool `json:"SelectAll,omitempty" xml:"SelectAll,omitempty"`
	// example:
	//
	// 1649260800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ae1527a9-2308-46ab-b10a-48ae7ff7****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UpdateFileProtectClientEventRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectClientEventRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectClientEventRequest) GetAlertLevels() []*int32 {
	return s.AlertLevels
}

func (s *UpdateFileProtectClientEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *UpdateFileProtectClientEventRequest) GetExcludeIdList() []*int64 {
	return s.ExcludeIdList
}

func (s *UpdateFileProtectClientEventRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *UpdateFileProtectClientEventRequest) GetIdList() []*int64 {
	return s.IdList
}

func (s *UpdateFileProtectClientEventRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateFileProtectClientEventRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateFileProtectClientEventRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *UpdateFileProtectClientEventRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *UpdateFileProtectClientEventRequest) GetNewStatus() *int32 {
	return s.NewStatus
}

func (s *UpdateFileProtectClientEventRequest) GetOperation() *string {
	return s.Operation
}

func (s *UpdateFileProtectClientEventRequest) GetProcPath() *string {
	return s.ProcPath
}

func (s *UpdateFileProtectClientEventRequest) GetRemark() []*string {
	return s.Remark
}

func (s *UpdateFileProtectClientEventRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateFileProtectClientEventRequest) GetSelectAll() *bool {
	return s.SelectAll
}

func (s *UpdateFileProtectClientEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateFileProtectClientEventRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateFileProtectClientEventRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UpdateFileProtectClientEventRequest) SetAlertLevels(v []*int32) *UpdateFileProtectClientEventRequest {
	s.AlertLevels = v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetEndTime(v int64) *UpdateFileProtectClientEventRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetExcludeIdList(v []*int64) *UpdateFileProtectClientEventRequest {
	s.ExcludeIdList = v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetFilePath(v string) *UpdateFileProtectClientEventRequest {
	s.FilePath = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetIdList(v []*int64) *UpdateFileProtectClientEventRequest {
	s.IdList = v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetInstanceId(v string) *UpdateFileProtectClientEventRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetInstanceName(v string) *UpdateFileProtectClientEventRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetInternetIp(v string) *UpdateFileProtectClientEventRequest {
	s.InternetIp = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetIntranetIp(v string) *UpdateFileProtectClientEventRequest {
	s.IntranetIp = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetNewStatus(v int32) *UpdateFileProtectClientEventRequest {
	s.NewStatus = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetOperation(v string) *UpdateFileProtectClientEventRequest {
	s.Operation = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetProcPath(v string) *UpdateFileProtectClientEventRequest {
	s.ProcPath = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetRemark(v []*string) *UpdateFileProtectClientEventRequest {
	s.Remark = v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetRuleName(v string) *UpdateFileProtectClientEventRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetSelectAll(v bool) *UpdateFileProtectClientEventRequest {
	s.SelectAll = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetStartTime(v int64) *UpdateFileProtectClientEventRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetStatus(v string) *UpdateFileProtectClientEventRequest {
	s.Status = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) SetUuid(v string) *UpdateFileProtectClientEventRequest {
	s.Uuid = &v
	return s
}

func (s *UpdateFileProtectClientEventRequest) Validate() error {
	return dara.Validate(s)
}

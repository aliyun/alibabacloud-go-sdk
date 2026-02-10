// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevels(v []*int32) *UpdateFileProtectRemarkRequest
	GetAlertLevels() []*int32
	SetEndTime(v int64) *UpdateFileProtectRemarkRequest
	GetEndTime() *int64
	SetId(v int64) *UpdateFileProtectRemarkRequest
	GetId() *int64
	SetIdList(v []*int64) *UpdateFileProtectRemarkRequest
	GetIdList() []*int64
	SetInstanceId(v string) *UpdateFileProtectRemarkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateFileProtectRemarkRequest
	GetInstanceName() *string
	SetInternetIp(v string) *UpdateFileProtectRemarkRequest
	GetInternetIp() *string
	SetIntranetIp(v string) *UpdateFileProtectRemarkRequest
	GetIntranetIp() *string
	SetOperation(v string) *UpdateFileProtectRemarkRequest
	GetOperation() *string
	SetRemark(v []*string) *UpdateFileProtectRemarkRequest
	GetRemark() []*string
	SetRuleName(v string) *UpdateFileProtectRemarkRequest
	GetRuleName() *string
	SetSelectAllAcrossPages(v bool) *UpdateFileProtectRemarkRequest
	GetSelectAllAcrossPages() *bool
	SetStartTime(v int64) *UpdateFileProtectRemarkRequest
	GetStartTime() *int64
	SetUuid(v string) *UpdateFileProtectRemarkRequest
	GetUuid() *string
}

type UpdateFileProtectRemarkRequest struct {
	// Alert notification level list.
	AlertLevels []*int32 `json:"AlertLevels,omitempty" xml:"AlertLevels,omitempty" type:"Repeated"`
	// End time timestamp.
	//
	// example:
	//
	// 1649040221
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// 1764
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Event ID list.
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// Asset instance ID.
	//
	// example:
	//
	// i-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Asset instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Public IP.
	//
	// example:
	//
	// 101.132.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Private IP.
	//
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// File operation type. Values:
	//
	// - **DELETE**: File deletion operation.
	//
	// - **WRITE**: File write operation.
	//
	// - **READ**: File read operation.
	//
	// - **RENAME**: File rename operation.
	//
	// - **CHOWN**: Set file owner and associated group operation.
	//
	// example:
	//
	// READ
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The remarks.
	Remark []*string `json:"Remark,omitempty" xml:"Remark,omitempty" type:"Repeated"`
	// Rule name.
	//
	// example:
	//
	// test-rule-1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Cross-page select all indicator. Values:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// true
	SelectAllAcrossPages *bool `json:"SelectAllAcrossPages,omitempty" xml:"SelectAllAcrossPages,omitempty"`
	// Start time timestamp.
	//
	// example:
	//
	// 1651290987000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Server UUID.
	//
	// example:
	//
	// 5d55af3c-35f3-4d4d-8ccc-8c5443b0****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UpdateFileProtectRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectRemarkRequest) GetAlertLevels() []*int32 {
	return s.AlertLevels
}

func (s *UpdateFileProtectRemarkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *UpdateFileProtectRemarkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateFileProtectRemarkRequest) GetIdList() []*int64 {
	return s.IdList
}

func (s *UpdateFileProtectRemarkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateFileProtectRemarkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateFileProtectRemarkRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *UpdateFileProtectRemarkRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *UpdateFileProtectRemarkRequest) GetOperation() *string {
	return s.Operation
}

func (s *UpdateFileProtectRemarkRequest) GetRemark() []*string {
	return s.Remark
}

func (s *UpdateFileProtectRemarkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateFileProtectRemarkRequest) GetSelectAllAcrossPages() *bool {
	return s.SelectAllAcrossPages
}

func (s *UpdateFileProtectRemarkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateFileProtectRemarkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UpdateFileProtectRemarkRequest) SetAlertLevels(v []*int32) *UpdateFileProtectRemarkRequest {
	s.AlertLevels = v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetEndTime(v int64) *UpdateFileProtectRemarkRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetId(v int64) *UpdateFileProtectRemarkRequest {
	s.Id = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetIdList(v []*int64) *UpdateFileProtectRemarkRequest {
	s.IdList = v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetInstanceId(v string) *UpdateFileProtectRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetInstanceName(v string) *UpdateFileProtectRemarkRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetInternetIp(v string) *UpdateFileProtectRemarkRequest {
	s.InternetIp = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetIntranetIp(v string) *UpdateFileProtectRemarkRequest {
	s.IntranetIp = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetOperation(v string) *UpdateFileProtectRemarkRequest {
	s.Operation = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetRemark(v []*string) *UpdateFileProtectRemarkRequest {
	s.Remark = v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetRuleName(v string) *UpdateFileProtectRemarkRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetSelectAllAcrossPages(v bool) *UpdateFileProtectRemarkRequest {
	s.SelectAllAcrossPages = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetStartTime(v int64) *UpdateFileProtectRemarkRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetUuid(v string) *UpdateFileProtectRemarkRequest {
	s.Uuid = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectEventStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertLevels(v []*int32) *UpdateFileProtectEventStatusRequest
	GetAlertLevels() []*int32
	SetEndTime(v int64) *UpdateFileProtectEventStatusRequest
	GetEndTime() *int64
	SetId(v []*int64) *UpdateFileProtectEventStatusRequest
	GetId() []*int64
	SetInstanceId(v string) *UpdateFileProtectEventStatusRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateFileProtectEventStatusRequest
	GetInstanceName() *string
	SetInternetIp(v string) *UpdateFileProtectEventStatusRequest
	GetInternetIp() *string
	SetIntranetIp(v string) *UpdateFileProtectEventStatusRequest
	GetIntranetIp() *string
	SetOperation(v string) *UpdateFileProtectEventStatusRequest
	GetOperation() *string
	SetRuleName(v string) *UpdateFileProtectEventStatusRequest
	GetRuleName() *string
	SetSelectAllAcrossPages(v bool) *UpdateFileProtectEventStatusRequest
	GetSelectAllAcrossPages() *bool
	SetStartTime(v int64) *UpdateFileProtectEventStatusRequest
	GetStartTime() *int64
	SetStatus(v int32) *UpdateFileProtectEventStatusRequest
	GetStatus() *int32
	SetUuid(v string) *UpdateFileProtectEventStatusRequest
	GetUuid() *string
}

type UpdateFileProtectEventStatusRequest struct {
	// The list of alert notification levels.
	AlertLevels []*int32 `json:"AlertLevels,omitempty" xml:"AlertLevels,omitempty" type:"Repeated"`
	// The end timestamp of the query.
	//
	// example:
	//
	// 1649040221
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of event IDs.
	Id []*int64 `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
	// The ID of the asset instance.
	//
	// example:
	//
	// i-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset instance.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 120.27.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The internal IP address.
	//
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The type of operation performed on the file. Valid values:
	//
	// - **DELETE**: Deletes a file.
	//
	// - **WRITE**: Writes to a file.
	//
	// - **READ**: Reads a file.
	//
	// - **RENAME**: Renames a file.
	//
	// - **CHOWN**: Changes the file owner and associated file group.
	//
	// example:
	//
	// READ
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The rule name.
	//
	// example:
	//
	// tetsRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies whether to select all items across pages. Valid values:
	//
	// - **true**: Selected.
	//
	// - **false**: Not selected.
	SelectAllAcrossPages *bool `json:"SelectAllAcrossPages,omitempty" xml:"SelectAllAcrossPages,omitempty"`
	// The start timestamp of the query.
	//
	// example:
	//
	// 1680919232000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event handling status.
	//
	// - **0**: Unhandled
	//
	// - **1**: Manually handled
	//
	// - **2**: Added to whitelist
	//
	// - **3**: Ignored
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// eb2c782e-64f2-4590-a86c-d90164df****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UpdateFileProtectEventStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectEventStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectEventStatusRequest) GetAlertLevels() []*int32 {
	return s.AlertLevels
}

func (s *UpdateFileProtectEventStatusRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *UpdateFileProtectEventStatusRequest) GetId() []*int64 {
	return s.Id
}

func (s *UpdateFileProtectEventStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateFileProtectEventStatusRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateFileProtectEventStatusRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *UpdateFileProtectEventStatusRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *UpdateFileProtectEventStatusRequest) GetOperation() *string {
	return s.Operation
}

func (s *UpdateFileProtectEventStatusRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateFileProtectEventStatusRequest) GetSelectAllAcrossPages() *bool {
	return s.SelectAllAcrossPages
}

func (s *UpdateFileProtectEventStatusRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateFileProtectEventStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateFileProtectEventStatusRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UpdateFileProtectEventStatusRequest) SetAlertLevels(v []*int32) *UpdateFileProtectEventStatusRequest {
	s.AlertLevels = v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetEndTime(v int64) *UpdateFileProtectEventStatusRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetId(v []*int64) *UpdateFileProtectEventStatusRequest {
	s.Id = v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetInstanceId(v string) *UpdateFileProtectEventStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetInstanceName(v string) *UpdateFileProtectEventStatusRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetInternetIp(v string) *UpdateFileProtectEventStatusRequest {
	s.InternetIp = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetIntranetIp(v string) *UpdateFileProtectEventStatusRequest {
	s.IntranetIp = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetOperation(v string) *UpdateFileProtectEventStatusRequest {
	s.Operation = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetRuleName(v string) *UpdateFileProtectEventStatusRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetSelectAllAcrossPages(v bool) *UpdateFileProtectEventStatusRequest {
	s.SelectAllAcrossPages = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetStartTime(v int64) *UpdateFileProtectEventStatusRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetStatus(v int32) *UpdateFileProtectEventStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) SetUuid(v string) *UpdateFileProtectEventStatusRequest {
	s.Uuid = &v
	return s
}

func (s *UpdateFileProtectEventStatusRequest) Validate() error {
	return dara.Validate(s)
}

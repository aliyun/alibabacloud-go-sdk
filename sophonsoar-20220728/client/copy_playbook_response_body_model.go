// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyPlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CopyPlaybookResponseBodyData) *CopyPlaybookResponseBody
	GetData() *CopyPlaybookResponseBodyData
	SetPage(v *CopyPlaybookResponseBodyPage) *CopyPlaybookResponseBody
	GetPage() *CopyPlaybookResponseBodyPage
	SetRequestId(v string) *CopyPlaybookResponseBody
	GetRequestId() *string
}

type CopyPlaybookResponseBody struct {
	// The response parameters.
	Data *CopyPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The pagination information.
	Page *CopyPlaybookResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2EC05B06-****-5F3E-****-3B1FAD76087A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyPlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *CopyPlaybookResponseBody) GetData() *CopyPlaybookResponseBodyData {
	return s.Data
}

func (s *CopyPlaybookResponseBody) GetPage() *CopyPlaybookResponseBodyPage {
	return s.Page
}

func (s *CopyPlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyPlaybookResponseBody) SetData(v *CopyPlaybookResponseBodyData) *CopyPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *CopyPlaybookResponseBody) SetPage(v *CopyPlaybookResponseBodyPage) *CopyPlaybookResponseBody {
	s.Page = v
	return s
}

func (s *CopyPlaybookResponseBody) SetRequestId(v string) *CopyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyPlaybookResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyPlaybookResponseBodyData struct {
	// The status of the playbook. Valid values:
	//
	// 	- 1: enabled.
	//
	// 	- 0: disabled.
	//
	// example:
	//
	// 1
	Active *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	// The description of the playbook.
	//
	// example:
	//
	// This is a action of processing for WAF
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the new playbook.
	//
	// example:
	//
	// 11111
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The number of playbook execution failures.
	//
	// example:
	//
	// 1
	FailNum *int32 `json:"FailNum,omitempty" xml:"FailNum,omitempty"`
	// The failure rate of playbook execution.
	//
	// example:
	//
	// 0.5
	FailRate *float64 `json:"FailRate,omitempty" xml:"FailRate,omitempty"`
	// The time when the playbook was created. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1655951601000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the playbook was modified. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1638270967000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The number of historical versions of the playbook.
	//
	// example:
	//
	// 1
	HistoryMd5 *int32 `json:"HistoryMd5,omitempty" xml:"HistoryMd5,omitempty"`
	// The input parameters of the playbook.
	//
	// example:
	//
	// [{\\"name\\":\\"1\\",\\"dataType\\":\\"String\\",\\"required\\":false,\\"isArray\\":false,\\"example\\":\\"\\",\\"description\\":\\"\\",\\"id\\":0,\\"typeName\\":\\"String\\",\\"dataClass\\":\\"normal\\"}]
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The time when the playbook was last run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1725258397847
	LastRuntime *int64 `json:"LastRuntime,omitempty" xml:"LastRuntime,omitempty"`
	// The online version MD5 of the playbook.
	//
	// example:
	//
	// 037046****1b00c4717963818ccbf2xx
	LogicReleaseTaskflowMd5 *string `json:"LogicReleaseTaskflowMd5,omitempty" xml:"LogicReleaseTaskflowMd5,omitempty"`
	// The output parameters of the playbook.
	//
	// example:
	//
	// []
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	// The type of the playbook. Valid values:
	//
	// 	- preset: predefined playbook.
	//
	// 	- user: custom playbook.
	//
	// example:
	//
	// user
	OwnType *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The permission to operate the playbook. Valid values:
	//
	// 	- 1: view.
	//
	// 	- 2: edit.
	//
	// example:
	//
	// 1
	Permission *int32 `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The status of the playbook.
	//
	// example:
	//
	// 1
	PlaybookStatus *int32 `json:"PlaybookStatus,omitempty" xml:"PlaybookStatus,omitempty"`
	// The UUID of the new playbook.
	//
	// example:
	//
	// 9e38111e-9794-4784-9ca8-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The number of successful playbook executions.
	//
	// example:
	//
	// 1
	SuccNum *int32 `json:"SuccNum,omitempty" xml:"SuccNum,omitempty"`
	// The ID of the user to which the playbook belongs.
	//
	// example:
	//
	// 13760*****8718726
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CopyPlaybookResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CopyPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *CopyPlaybookResponseBodyData) GetActive() *int32 {
	return s.Active
}

func (s *CopyPlaybookResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CopyPlaybookResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CopyPlaybookResponseBodyData) GetFailNum() *int32 {
	return s.FailNum
}

func (s *CopyPlaybookResponseBodyData) GetFailRate() *float64 {
	return s.FailRate
}

func (s *CopyPlaybookResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CopyPlaybookResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CopyPlaybookResponseBodyData) GetHistoryMd5() *int32 {
	return s.HistoryMd5
}

func (s *CopyPlaybookResponseBodyData) GetInputParams() *string {
	return s.InputParams
}

func (s *CopyPlaybookResponseBodyData) GetLastRuntime() *int64 {
	return s.LastRuntime
}

func (s *CopyPlaybookResponseBodyData) GetLogicReleaseTaskflowMd5() *string {
	return s.LogicReleaseTaskflowMd5
}

func (s *CopyPlaybookResponseBodyData) GetOutputParams() *string {
	return s.OutputParams
}

func (s *CopyPlaybookResponseBodyData) GetOwnType() *string {
	return s.OwnType
}

func (s *CopyPlaybookResponseBodyData) GetPermission() *int32 {
	return s.Permission
}

func (s *CopyPlaybookResponseBodyData) GetPlaybookStatus() *int32 {
	return s.PlaybookStatus
}

func (s *CopyPlaybookResponseBodyData) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *CopyPlaybookResponseBodyData) GetSuccNum() *int32 {
	return s.SuccNum
}

func (s *CopyPlaybookResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *CopyPlaybookResponseBodyData) SetActive(v int32) *CopyPlaybookResponseBodyData {
	s.Active = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetDescription(v string) *CopyPlaybookResponseBodyData {
	s.Description = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetDisplayName(v string) *CopyPlaybookResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetFailNum(v int32) *CopyPlaybookResponseBodyData {
	s.FailNum = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetFailRate(v float64) *CopyPlaybookResponseBodyData {
	s.FailRate = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetGmtCreate(v int64) *CopyPlaybookResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetGmtModified(v int64) *CopyPlaybookResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetHistoryMd5(v int32) *CopyPlaybookResponseBodyData {
	s.HistoryMd5 = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetInputParams(v string) *CopyPlaybookResponseBodyData {
	s.InputParams = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetLastRuntime(v int64) *CopyPlaybookResponseBodyData {
	s.LastRuntime = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetLogicReleaseTaskflowMd5(v string) *CopyPlaybookResponseBodyData {
	s.LogicReleaseTaskflowMd5 = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetOutputParams(v string) *CopyPlaybookResponseBodyData {
	s.OutputParams = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetOwnType(v string) *CopyPlaybookResponseBodyData {
	s.OwnType = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetPermission(v int32) *CopyPlaybookResponseBodyData {
	s.Permission = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetPlaybookStatus(v int32) *CopyPlaybookResponseBodyData {
	s.PlaybookStatus = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetPlaybookUuid(v string) *CopyPlaybookResponseBodyData {
	s.PlaybookUuid = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetSuccNum(v int32) *CopyPlaybookResponseBodyData {
	s.SuccNum = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) SetTenantId(v string) *CopyPlaybookResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *CopyPlaybookResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CopyPlaybookResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CopyPlaybookResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s CopyPlaybookResponseBodyPage) GoString() string {
	return s.String()
}

func (s *CopyPlaybookResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *CopyPlaybookResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CopyPlaybookResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *CopyPlaybookResponseBodyPage) SetPageNumber(v int32) *CopyPlaybookResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *CopyPlaybookResponseBodyPage) SetPageSize(v int32) *CopyPlaybookResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *CopyPlaybookResponseBodyPage) SetTotalCount(v int32) *CopyPlaybookResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *CopyPlaybookResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

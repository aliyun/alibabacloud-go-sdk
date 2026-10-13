// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetMigrateApplicationResponseBody
	GetCode() *int64
	SetData(v *GetMigrateApplicationResponseBodyData) *GetMigrateApplicationResponseBody
	GetData() *GetMigrateApplicationResponseBodyData
	SetMessage(v string) *GetMigrateApplicationResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetMigrateApplicationResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *GetMigrateApplicationResponseBody
	GetTraceId() *string
}

type GetMigrateApplicationResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetMigrateApplicationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// “”
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// 3b52155517851187240212426d0099
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetMigrateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMigrateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrateApplicationResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetMigrateApplicationResponseBody) GetData() *GetMigrateApplicationResponseBodyData {
	return s.Data
}

func (s *GetMigrateApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMigrateApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMigrateApplicationResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetMigrateApplicationResponseBody) SetCode(v int64) *GetMigrateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetMigrateApplicationResponseBody) SetData(v *GetMigrateApplicationResponseBodyData) *GetMigrateApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetMigrateApplicationResponseBody) SetMessage(v string) *GetMigrateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetMigrateApplicationResponseBody) SetSuccess(v bool) *GetMigrateApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *GetMigrateApplicationResponseBody) SetTraceId(v string) *GetMigrateApplicationResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetMigrateApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMigrateApplicationResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 4
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of results.
	Result []*GetMigrateApplicationResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalSize *int64 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetMigrateApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMigrateApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMigrateApplicationResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetMigrateApplicationResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetMigrateApplicationResponseBodyData) GetResult() []*GetMigrateApplicationResponseBodyDataResult {
	return s.Result
}

func (s *GetMigrateApplicationResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *GetMigrateApplicationResponseBodyData) SetCurrentPage(v int64) *GetMigrateApplicationResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyData) SetPageSize(v int64) *GetMigrateApplicationResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyData) SetResult(v []*GetMigrateApplicationResponseBodyDataResult) *GetMigrateApplicationResponseBodyData {
	s.Result = v
	return s
}

func (s *GetMigrateApplicationResponseBodyData) SetTotalSize(v int64) *GetMigrateApplicationResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMigrateApplicationResponseBodyDataResult struct {
	// The detailed status of the application.
	//
	// example:
	//
	// "{\\"a259bc63-c1f4-48ba-89b1-d6a4591e946c\\":{\\"app_name\\":\\"prodoam1\\",\\"cluster_name\\":\\"hthoamtest\\",\\"cluster_id\\":\\"f68945a8-47af-4644-a951-0486d674f26e\\",\\"error_msg\\":\\"\\",\\"is_multilingual_app\\":false,\\"state\\":\\"success\\",\\"clusterId\\":\\"f68945a8-47af-4644-a951-0486d674f26e\\",\\"app_id\\":\\"a259bc63-c1f4-48ba-89b1-d6a4591e946c\\",\\"logical_region_id\\":\\"cn-hangzhou\\",\\"k8sNamespace\\":\\"default\\"},\\"c5a258ac-f02d-4035-b59f-cd7f4a34156a\\":{\\"app_name\\":\\"testoam1\\",\\"cluster_name\\":\\"hthoamtest\\",\\"cluster_id\\":\\"f68945a8-47af-4644-a951-0486d674f26e\\",\\"error_msg\\":\\"\\",\\"is_multilingual_app\\":false,\\"state\\":\\"success\\",\\"clusterId\\":\\"f68945a8-47af-4644-a951-0486d674f26e\\",\\"app_id\\":\\"c5a258ac-f02d-4035-b59f-cd7f4a34156a\\",\\"logical_region_id\\":\\"cn-hangzhou\\",\\"k8sNamespace\\":\\"default\\"},\\"d132c827-db2f-444e-bc9e-9182d8112806\\":{\\"app_name\\":\\"oam1\\",\\"cluster_name\\":\\"hthoamtest\\",\\"cluster_id\\":\\"f68945a8-47af-4644-a951-0486d674f26e\\",\\"error_msg\\":\\"\\",\\"is_multilingual_app\\":false,\\"state\\":\\"success\\",\\"clusterId\\":\\"f68945a8-47af-4644-a951-0486d674f26e\\",\\"app_id\\":\\"d132c827-db2f-444e-bc9e-9182d8112806\\",\\"logical_region_id\\":\\"cn-hangzhou\\",\\"k8sNamespace\\":\\"default\\"}}"
	AppDetailState *string `json:"appDetailState,omitempty" xml:"appDetailState,omitempty"`
	// Indicates whether binary resources are cleaned up. Valid values: 0: not cleaned up. 1: cleaned up.
	//
	// example:
	//
	// 1
	BinaryCleaned *int64 `json:"binaryCleaned,omitempty" xml:"binaryCleaned,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1791512615778
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The additional information.
	ExtraInfo *GetMigrateApplicationResponseBodyDataResultExtraInfo `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	// The migration ID.
	//
	// example:
	//
	// "e7c2e7d9-5e87-46a2-99e9-2be0b577a3fd"
	MigrationId *string `json:"migrationId,omitempty" xml:"migrationId,omitempty"`
	// The operation type of the migration task. Valid values: 1: import. 2: export. 3: clone.
	//
	// example:
	//
	// 2
	MigrationType *int64 `json:"migrationType,omitempty" xml:"migrationType,omitempty"`
	// The executor.
	//
	// example:
	//
	// lvwan
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// regionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The application migration status. Valid values: failed: Failed. finished: Succeeded.
	//
	// example:
	//
	// finished
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The task progress percentage, stored as a floating-point number. For example, if the API returns 0.33, the progress bar displays 33%.
	//
	// example:
	//
	// 0.33
	StatePercent *string `json:"statePercent,omitempty" xml:"statePercent,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1791512615778
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetMigrateApplicationResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetMigrateApplicationResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetAppDetailState() *string {
	return s.AppDetailState
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetBinaryCleaned() *int64 {
	return s.BinaryCleaned
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetExtraInfo() *GetMigrateApplicationResponseBodyDataResultExtraInfo {
	return s.ExtraInfo
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetMigrationId() *string {
	return s.MigrationId
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetMigrationType() *int64 {
	return s.MigrationType
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetOwner() *string {
	return s.Owner
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetState() *string {
	return s.State
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetStatePercent() *string {
	return s.StatePercent
}

func (s *GetMigrateApplicationResponseBodyDataResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetAppDetailState(v string) *GetMigrateApplicationResponseBodyDataResult {
	s.AppDetailState = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetBinaryCleaned(v int64) *GetMigrateApplicationResponseBodyDataResult {
	s.BinaryCleaned = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetCreateTime(v int64) *GetMigrateApplicationResponseBodyDataResult {
	s.CreateTime = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetExtraInfo(v *GetMigrateApplicationResponseBodyDataResultExtraInfo) *GetMigrateApplicationResponseBodyDataResult {
	s.ExtraInfo = v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetMigrationId(v string) *GetMigrateApplicationResponseBodyDataResult {
	s.MigrationId = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetMigrationType(v int64) *GetMigrateApplicationResponseBodyDataResult {
	s.MigrationType = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetOwner(v string) *GetMigrateApplicationResponseBodyDataResult {
	s.Owner = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetRegionId(v string) *GetMigrateApplicationResponseBodyDataResult {
	s.RegionId = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetState(v string) *GetMigrateApplicationResponseBodyDataResult {
	s.State = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetStatePercent(v string) *GetMigrateApplicationResponseBodyDataResult {
	s.StatePercent = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) SetUpdateTime(v int64) *GetMigrateApplicationResponseBodyDataResult {
	s.UpdateTime = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResult) Validate() error {
	if s.ExtraInfo != nil {
		if err := s.ExtraInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMigrateApplicationResponseBodyDataResultExtraInfo struct {
	// appJsonFile
	//
	// example:
	//
	// "https://edas-hz-secure.oss-cn-hangzhou.aliyuncs.com/edas-app-migration/1057288613002791/e7c2e7d9-5e87-46a2-99e9-2be0b577a3fd/edas-app.json?Expires=1791517039&OSSAccessKeyId=LTAIWLbjtAGvdMWX&Signature=HyZ2b3u5mlP56T8QgkARwWE5cQc%3D"
	AppJsonFile *string `json:"appJsonFile,omitempty" xml:"appJsonFile,omitempty"`
	// The error message.
	//
	// example:
	//
	// “”
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
}

func (s GetMigrateApplicationResponseBodyDataResultExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMigrateApplicationResponseBodyDataResultExtraInfo) GoString() string {
	return s.String()
}

func (s *GetMigrateApplicationResponseBodyDataResultExtraInfo) GetAppJsonFile() *string {
	return s.AppJsonFile
}

func (s *GetMigrateApplicationResponseBodyDataResultExtraInfo) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetMigrateApplicationResponseBodyDataResultExtraInfo) SetAppJsonFile(v string) *GetMigrateApplicationResponseBodyDataResultExtraInfo {
	s.AppJsonFile = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResultExtraInfo) SetErrorMsg(v string) *GetMigrateApplicationResponseBodyDataResultExtraInfo {
	s.ErrorMsg = &v
	return s
}

func (s *GetMigrateApplicationResponseBodyDataResultExtraInfo) Validate() error {
	return dara.Validate(s)
}
